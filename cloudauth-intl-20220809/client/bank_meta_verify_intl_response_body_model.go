// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBankMetaVerifyIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BankMetaVerifyIntlResponseBody
	GetCode() *string
	SetMessage(v string) *BankMetaVerifyIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *BankMetaVerifyIntlResponseBody
	GetRequestId() *string
	SetResultObject(v *BankMetaVerifyIntlResponseBodyResultObject) *BankMetaVerifyIntlResponseBody
	GetResultObject() *BankMetaVerifyIntlResponseBodyResultObject
}

type BankMetaVerifyIntlResponseBody struct {
	// The response code. A value of 200 indicates success. Other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information.
	ResultObject *BankMetaVerifyIntlResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s BankMetaVerifyIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *BankMetaVerifyIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BankMetaVerifyIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BankMetaVerifyIntlResponseBody) GetResultObject() *BankMetaVerifyIntlResponseBodyResultObject {
	return s.ResultObject
}

func (s *BankMetaVerifyIntlResponseBody) SetCode(v string) *BankMetaVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBody) SetMessage(v string) *BankMetaVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBody) SetRequestId(v string) *BankMetaVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBody) SetResultObject(v *BankMetaVerifyIntlResponseBodyResultObject) *BankMetaVerifyIntlResponseBody {
	s.ResultObject = v
	return s
}

func (s *BankMetaVerifyIntlResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BankMetaVerifyIntlResponseBodyResultObject struct {
	// The verification result code. Valid values:
	//
	// - 1: Verification consistent (billable).
	//
	// - 2: Verification inconsistent (billable).
	//
	// - 3: No record found (not billable).
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The verification details. Valid values:
	//
	// - **101**: Verification passed.
	//
	// - **201**: Authentication information inconsistent. The cardholder information is incorrect.
	//
	// - **202**: Authentication information inconsistent. The bank card has not enabled authenticated payment.
	//
	// - **203**: Authentication information inconsistent. The bank card has expired.
	//
	// - **204**: Authentication information inconsistent. The bank card is restricted.
	//
	// - **205**: Authentication information inconsistent. The card has been confiscated.
	//
	// - **206**: Authentication information inconsistent. The bank card is invalid.
	//
	// - **207**: Authentication information inconsistent. No issuing bank found for this card.
	//
	// - **208**: Authentication information inconsistent. The card is not initialized or is a dormant card.
	//
	// - **209**: Authentication information inconsistent. The card is a fraudulent or retained card.
	//
	// - **210**: Authentication information inconsistent. The card has been reported lost.
	//
	// - **211**: Authentication information inconsistent. The number of incorrect password attempts has exceeded the limit.
	//
	// - **212**: Authentication information inconsistent. The issuing bank does not support this transaction.
	//
	// - **213**: Authentication information inconsistent. The card status is abnormal or the card is invalid.
	//
	// - **214**: Authentication information inconsistent. No phone number is registered with the card.
	//
	// - **215**: Authentication information inconsistent. The password, expiration date, or CVN2 is incorrect.
	//
	// - **216**: Authentication information inconsistent. Other card exceptions.
	//
	// - **301**: Verification unavailable. The bank card does not support this service.
	//
	// - **302**: Verification unavailable. Verification failed or the bank rejected the verification. Contact the issuing bank.
	//
	// - **303**: Verification unavailable. The bank card does not currently support phone number verification.
	//
	// - **304**: Verification unavailable. The bank card number is incorrect.
	//
	// - **305**: Verification unavailable. Other reasons.
	//
	// - **306**: Verification unavailable. The number of verification attempts has exceeded the limit.
	//
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s BankMetaVerifyIntlResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyIntlResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) SetBizCode(v string) *BankMetaVerifyIntlResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) SetSubCode(v string) *BankMetaVerifyIntlResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
