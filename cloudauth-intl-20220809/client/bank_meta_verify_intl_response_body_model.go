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
	// Return code: 200 for success, others for failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information
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
	// Verification result code.
	//
	// - 1: Consistent (charged)
	//
	// - 2: Inconsistent (charged)
	//
	// - 3: No record found (not charged)
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Verification details:
	//
	// - **101**: Verification passed.
	//
	// - **201**: Authentication information does not match, cardholder information is incorrect.
	//
	// - **202**: Authentication information does not match, bank card has not been activated for authenticated payments.
	//
	// - **203**: Authentication information does not match, bank card has expired.
	//
	// - **204**: Authentication information does not match, bank card is a restricted card.
	//
	// - **205**: Authentication information does not match, this card has been confiscated.
	//
	// - **206**: Authentication information does not match, bank card is invalid.
	//
	// - **207**: Authentication information does not match, this card has no corresponding issuing bank.
	//
	// - **208**: Authentication information does not match, this card is uninitialized or dormant.
	//
	// - **209**: Authentication information does not match, this card is a cheating card or swallowed card.
	//
	// - **210**: Authentication information does not match, this card has been reported lost.
	//
	// - **211**: Authentication information does not match, password error limit exceeded.
	//
	// - **212**: Authentication information does not match, issuing bank does not support this transaction.
	//
	// - **213**: Authentication information does not match, card status is abnormal or card is invalid.
	//
	// - **214**: Authentication information does not match, no phone number reserved.
	//
	// - **215**: Authentication information does not match, entered password, expiration date, or CVN2 is incorrect.
	//
	// - **216**: Authentication information does not match, other card anomalies.
	//
	// - **301**: Unable to verify, bank card does not support this service.
	//
	// - **302**: Unable to verify, verification failed or bank refused verification, please contact the issuing bank.
	//
	// - **303**: Unable to verify, bank card does not currently support phone number verification.
	//
	// - **304**: Unable to verify, bank card number is incorrect.
	//
	// - **305**: Unable to verify, other reasons.
	//
	// - **306**: Unable to verify, verification attempt limit exceeded.
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
