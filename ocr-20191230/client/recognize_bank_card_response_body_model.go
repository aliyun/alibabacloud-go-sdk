// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeBankCardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeBankCardResponseBodyData) *RecognizeBankCardResponseBody
	GetData() *RecognizeBankCardResponseBodyData
	SetRequestId(v string) *RecognizeBankCardResponseBody
	GetRequestId() *string
}

type RecognizeBankCardResponseBody struct {
	Data *RecognizeBankCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D9C7521-0367-42EE-9646-FD066CCADB26
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBankCardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBankCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponseBody) GetData() *RecognizeBankCardResponseBodyData {
	return s.Data
}

func (s *RecognizeBankCardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeBankCardResponseBody) SetData(v *RecognizeBankCardResponseBodyData) *RecognizeBankCardResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeBankCardResponseBody) SetRequestId(v string) *RecognizeBankCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeBankCardResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeBankCardResponseBodyData struct {
	BankName *string `json:"BankName,omitempty" xml:"BankName,omitempty"`
	// example:
	//
	// 6212262315007683105
	CardNumber *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	CardType   *string `json:"CardType,omitempty" xml:"CardType,omitempty"`
	// example:
	//
	// 07/26
	ValidDate *string `json:"ValidDate,omitempty" xml:"ValidDate,omitempty"`
}

func (s RecognizeBankCardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBankCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponseBodyData) GetBankName() *string {
	return s.BankName
}

func (s *RecognizeBankCardResponseBodyData) GetCardNumber() *string {
	return s.CardNumber
}

func (s *RecognizeBankCardResponseBodyData) GetCardType() *string {
	return s.CardType
}

func (s *RecognizeBankCardResponseBodyData) GetValidDate() *string {
	return s.ValidDate
}

func (s *RecognizeBankCardResponseBodyData) SetBankName(v string) *RecognizeBankCardResponseBodyData {
	s.BankName = &v
	return s
}

func (s *RecognizeBankCardResponseBodyData) SetCardNumber(v string) *RecognizeBankCardResponseBodyData {
	s.CardNumber = &v
	return s
}

func (s *RecognizeBankCardResponseBodyData) SetCardType(v string) *RecognizeBankCardResponseBodyData {
	s.CardType = &v
	return s
}

func (s *RecognizeBankCardResponseBodyData) SetValidDate(v string) *RecognizeBankCardResponseBodyData {
	s.ValidDate = &v
	return s
}

func (s *RecognizeBankCardResponseBodyData) Validate() error {
	return dara.Validate(s)
}
