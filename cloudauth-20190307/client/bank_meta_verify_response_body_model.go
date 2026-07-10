// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBankMetaVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BankMetaVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *BankMetaVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *BankMetaVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *BankMetaVerifyResponseBodyResultObject) *BankMetaVerifyResponseBody
	GetResultObject() *BankMetaVerifyResponseBodyResultObject
}

type BankMetaVerifyResponseBody struct {
	// The return code. A value of 200 indicates success. Other values indicate failure.
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
	// 473469C7-A***B-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information.
	ResultObject *BankMetaVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s BankMetaVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *BankMetaVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BankMetaVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BankMetaVerifyResponseBody) GetResultObject() *BankMetaVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *BankMetaVerifyResponseBody) SetCode(v string) *BankMetaVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *BankMetaVerifyResponseBody) SetMessage(v string) *BankMetaVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *BankMetaVerifyResponseBody) SetRequestId(v string) *BankMetaVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *BankMetaVerifyResponseBody) SetResultObject(v *BankMetaVerifyResponseBodyResultObject) *BankMetaVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *BankMetaVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BankMetaVerifyResponseBodyResultObject struct {
	// The verification result. Valid values:
	//
	// - 1: Consistent (billable).
	//
	// - 2: Inconsistent (billable).
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
	// - **201**: Verification information is inconsistent. The cardholder information is incorrect.
	//
	// - **202**: Verification information is inconsistent. The bank card has not enabled authenticated payment.
	//
	// - **203**: Verification information is inconsistent. The bank card has expired.
	//
	// - **204**: Verification information is inconsistent. The bank card is restricted.
	//
	// - **205**: Verification information is inconsistent. The card has been confiscated.
	//
	// - **206**: Verification information is inconsistent. The bank card is invalid.
	//
	// - **207**: Verification information is inconsistent. No issuing bank found for this card.
	//
	// - **208**: Verification information is inconsistent. The card has not been initialized or is a dormant card.
	//
	// - **209**: Verification information is inconsistent. The card is a fraudulent or retained card.
	//
	// - **210**: Verification information is inconsistent. The card has been reported lost.
	//
	// - **211**: Verification information is inconsistent. The number of incorrect password attempts has exceeded the limit.
	//
	// - **212**: Verification information is inconsistent. The issuing bank does not support this transaction.
	//
	// - **213**: Verification information is inconsistent. The card status is abnormal or the card is invalid.
	//
	// - **214**: Verification information is inconsistent. No phone number is registered with the card.
	//
	// - **215**: Verification information is inconsistent. The password, expiration date, or CVN2 is incorrect.
	//
	// - **216**: Verification information is inconsistent. Other card exceptions.
	//
	// - **301**: Unable to verify. The bank card does not support this service.
	//
	// - **302**: Unable to verify. Verification failed or the bank rejected the verification. Contact the issuing bank.
	//
	// - **303**: Unable to verify. The bank card does not support phone number verification.
	//
	// - **304**: Unable to verify. The bank card number is incorrect.
	//
	// - **305**: Unable to verify. Other reasons.
	//
	// - **306**: Unable to verify. The number of verification attempts has exceeded the limit.
	//
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s BankMetaVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *BankMetaVerifyResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *BankMetaVerifyResponseBodyResultObject) SetBizCode(v string) *BankMetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *BankMetaVerifyResponseBodyResultObject) SetSubCode(v string) *BankMetaVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *BankMetaVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
