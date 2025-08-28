// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotlineTransferRegisterFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHotlineTransferRegisterFileResponseBody
	GetCode() *string
	SetData(v *ListHotlineTransferRegisterFileResponseBodyData) *ListHotlineTransferRegisterFileResponseBody
	GetData() *ListHotlineTransferRegisterFileResponseBodyData
	SetMessage(v string) *ListHotlineTransferRegisterFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotlineTransferRegisterFileResponseBody
	GetRequestId() *string
}

type ListHotlineTransferRegisterFileResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *ListHotlineTransferRegisterFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHotlineTransferRegisterFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineTransferRegisterFileResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferRegisterFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHotlineTransferRegisterFileResponseBody) GetData() *ListHotlineTransferRegisterFileResponseBodyData {
	return s.Data
}

func (s *ListHotlineTransferRegisterFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotlineTransferRegisterFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotlineTransferRegisterFileResponseBody) SetCode(v string) *ListHotlineTransferRegisterFileResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBody) SetData(v *ListHotlineTransferRegisterFileResponseBodyData) *ListHotlineTransferRegisterFileResponseBody {
	s.Data = v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBody) SetMessage(v string) *ListHotlineTransferRegisterFileResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBody) SetRequestId(v string) *ListHotlineTransferRegisterFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHotlineTransferRegisterFileResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The registration file.
	Values []*ListHotlineTransferRegisterFileResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListHotlineTransferRegisterFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineTransferRegisterFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) GetValues() []*ListHotlineTransferRegisterFileResponseBodyDataValues {
	return s.Values
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) SetPageNo(v int32) *ListHotlineTransferRegisterFileResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) SetPageSize(v int32) *ListHotlineTransferRegisterFileResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) SetTotal(v int64) *ListHotlineTransferRegisterFileResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) SetValues(v []*ListHotlineTransferRegisterFileResponseBodyDataValues) *ListHotlineTransferRegisterFileResponseBodyData {
	s.Values = v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListHotlineTransferRegisterFileResponseBodyDataValues struct {
	// The authenticity of the commitment.
	//
	// example:
	//
	// true
	Agree *string `json:"Agree,omitempty" xml:"Agree,omitempty"`
	// The enterprise name.
	//
	// example:
	//
	// A**
	CorpName *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	// The China 400 number.
	//
	// example:
	//
	// 400****
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// The ID card number of the handler.
	//
	// example:
	//
	// 5****************9
	MngOpIdentityCard *string `json:"MngOpIdentityCard,omitempty" xml:"MngOpIdentityCard,omitempty"`
	// The email address of the handler.
	//
	// example:
	//
	// username@example.com
	MngOpMail *string `json:"MngOpMail,omitempty" xml:"MngOpMail,omitempty"`
	// The mobile phone number of the handler.
	//
	// example:
	//
	// 150****0000
	MngOpMobile *string `json:"MngOpMobile,omitempty" xml:"MngOpMobile,omitempty"`
	// The name of the handler.
	//
	// example:
	//
	// A***
	MngOpName *string `json:"MngOpName,omitempty" xml:"MngOpName,omitempty"`
	// The qualification ID.
	//
	// example:
	//
	// 1234****
	QualificationId *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// The unique code of the query operation.
	//
	// example:
	//
	// 1
	ResUniqueCode *int64 `json:"ResUniqueCode,omitempty" xml:"ResUniqueCode,omitempty"`
}

func (s ListHotlineTransferRegisterFileResponseBodyDataValues) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineTransferRegisterFileResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) GetAgree() *string {
	return s.Agree
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) GetCorpName() *string {
	return s.CorpName
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) GetHotlineNumber() *string {
	return s.HotlineNumber
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) GetMngOpIdentityCard() *string {
	return s.MngOpIdentityCard
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) GetMngOpMail() *string {
	return s.MngOpMail
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) GetMngOpMobile() *string {
	return s.MngOpMobile
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) GetMngOpName() *string {
	return s.MngOpName
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) GetQualificationId() *string {
	return s.QualificationId
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) GetResUniqueCode() *int64 {
	return s.ResUniqueCode
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetAgree(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.Agree = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetCorpName(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.CorpName = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetHotlineNumber(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.HotlineNumber = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetMngOpIdentityCard(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.MngOpIdentityCard = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetMngOpMail(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.MngOpMail = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetMngOpMobile(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.MngOpMobile = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetMngOpName(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.MngOpName = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetQualificationId(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.QualificationId = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetResUniqueCode(v int64) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.ResUniqueCode = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) Validate() error {
	return dara.Validate(s)
}
