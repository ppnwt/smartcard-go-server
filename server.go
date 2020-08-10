package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ebfe/scard"
	"github.com/gogetth/sscard"
)

func main(){
	//Establish a PC/SC context
	context, err := scard.EstablishContext()
	if err != nil {
		fmt.Println("Error EstablishContext:", err)
		return
	}

	// Release the PC/SC context (when needed)
	readers, err := context.ListReaders()
	if err != nil {
		fmt.Println("Error ListReaders:", err)
		return
	}

	// Use the first reader
	reader := readers[0]
	fmt.Println("Using reader:", reader)

	// Connect to the card
	card, err := context.Connect(reader, scard.ShareShared, scard.ProtocolAny)
	if err != nil {
		fmt.Println("Error Connect:", err)
		return
	}

	// Disconnect (when needed)
	defer card.Disconnect(scard.LeaveCard)

		// Send select APDU
	selectRsp, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardSelect)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Println("resp sscard.APDUThaiIDCardSelect: ", selectRsp)

	// CID
	cid, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardCID)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("cid: _%s_\n", string(cid))

	// FULLNAME EN
	fullnameEN, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardFullnameEn)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("fullnameEN: _%s_\n", string(fullnameEN))

	// FULLNAME TH
	fullnameTH, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardFullnameTh)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("fullnameTH: _%s_\n", string(fullnameTH))

	// DOB
	birth, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardBirth)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("birth: _%s_\n", string(birth))

	// GENDER
	gender, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardGender)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("gender: _%s_\n", string(gender))

	// ISSUER
	issuer, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardIssuer)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("issuer: _%s_\n", string(issuer))

	// ISSUE DATE
	issueDate, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardIssuedate)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("issueDate: _%s_\n", string(issueDate))
	
	// ISSUE DATE EXP
	issueExp, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardExpiredate)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("issueExp: _%s_\n", string(issueExp))

	// ADDRESS
	address, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardAddress)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("address: _%s_\n", string(address))

	// send response

	r := gin.Default()
	r.GET("/smartcard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "card success",
			"cid": string(cid),
			"fullname_en": string(fullnameEN),
			"fullname_th": string(fullnameTH),
			"birth": string(birth),
			"gender": string(gender),
			"issuer": string(issuer),
			"issueDate": string(issueDate),
			"issueExp": string(issueExp),
			"addressTH": string(address),
		})
	})
	r.Run()
}