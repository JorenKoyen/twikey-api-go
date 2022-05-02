package twikey

import (
	"os"
	"testing"
)

func TestDocumentFeed(t *testing.T) {
	if os.Getenv("TWIKEY_API_KEY") == "" {
		t.Skip("No TWIKEY_API_KEY available")
	}

	c := newTestClient()
	invite, err := c.DocumentInvite(&InviteRequest{
		Template:       getEnv("CT", "1"),
		CustomerNumber: "123",
		Amount:         "123.10",
		Email:          "john@doe.com",
		Firstname:      "John",
		Lastname:       "Doe",
		Language:       "en",
		Address:        "Abbey road",
		City:           "Liverpool",
		Zip:            "1526",
		Country:        "BE",
		Mobile:         "",
		CompanyName:    "",
		Coc:            "",
		Iban:           "",
		Bic:            "",
		MandateNumber:  "",
		ContractNumber: "",
	})
	if err != nil {
		t.Fatal(err)
	}

	if invite == nil || invite.Url == "" {
		t.Error("No valid invite retrieved")
	}

	mandateimport, err := c.DocumentSign(&InviteRequest{
		Template:       getEnv("CT", "1"),
		CustomerNumber: "123",
		Amount:         "123.10",
		Email:          "john@doe.com",
		Firstname:      "John",
		Lastname:       "Doe",
		Language:       "en",
		Address:        "Abbey road",
		City:           "Liverpool",
		Zip:            "1526",
		Country:        "BE",
		Mobile:         "",
		CompanyName:    "",
		Coc:            "",
		Iban:           "BE09363107700857",
		MandateNumber:  "",
		ContractNumber: "",
		Method:         "import",
	})
	if err != nil {
		t.Fatal(err)
	}

	if mandateimport == nil {
		t.Error("No valid import")
	} else {
		t.Log("Imported", mandateimport.MndtId)
	}

	t.Run("DocumentFeed", func(t *testing.T) {
		err := c.DocumentFeed(func(mandate *Mndt, eventTime string) {
			t.Log("Document created   ", mandate.MndtId, " @ ", eventTime)
		}, func(originalMandateNumber string, mandate *Mndt, reason *AmdmntRsn, eventTime string) {
			t.Log("Document updated   ", originalMandateNumber, reason.Rsn, " @ ", eventTime)
		}, func(mandateNumber string, reason *CxlRsn, eventTime string) {
			t.Log("Document cancelled ", mandateNumber, reason.Rsn, " @ ", eventTime)
		})
		if err != nil {
			return
		}
	})
}

func TestDocumentDetail(t *testing.T) {
	if os.Getenv("TWIKEY_API_KEY") == "" {
		t.Skip("No TWIKEY_API_KEY available")
	}

	if os.Getenv("MNDTNUMBER") == "" {
		t.Skip("No MNDTNUMBER available")
	}

	c := newTestClient()
	mndt, err := c.DocumentDetail(os.Getenv("MNDTNUMBER"))
	if err != nil {
		t.Fatal(err, os.Getenv("MNDTNUMBER"))
	}

	if mndt == nil {
		t.Error("No valid mandate retrieved")
	} else {
		t.Log("Got Mandate", mndt.MndtId)
	}
}
