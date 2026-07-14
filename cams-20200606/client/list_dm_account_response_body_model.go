// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDmAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListDmAccountResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListDmAccountResponseBody
	GetCode() *string
	SetData(v []*ListDmAccountResponseBodyData) *ListDmAccountResponseBody
	GetData() []*ListDmAccountResponseBodyData
	SetMessage(v string) *ListDmAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDmAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDmAccountResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListDmAccountResponseBody
	GetTotal() *int64
}

type ListDmAccountResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code. Valid values:
	//
	// - OK: The request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListDmAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xis-sx***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDmAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDmAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListDmAccountResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListDmAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDmAccountResponseBody) GetData() []*ListDmAccountResponseBodyData {
	return s.Data
}

func (s *ListDmAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDmAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDmAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDmAccountResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListDmAccountResponseBody) SetAccessDeniedDetail(v string) *ListDmAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListDmAccountResponseBody) SetCode(v string) *ListDmAccountResponseBody {
	s.Code = &v
	return s
}

func (s *ListDmAccountResponseBody) SetData(v []*ListDmAccountResponseBodyData) *ListDmAccountResponseBody {
	s.Data = v
	return s
}

func (s *ListDmAccountResponseBody) SetMessage(v string) *ListDmAccountResponseBody {
	s.Message = &v
	return s
}

func (s *ListDmAccountResponseBody) SetRequestId(v string) *ListDmAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDmAccountResponseBody) SetSuccess(v bool) *ListDmAccountResponseBody {
	s.Success = &v
	return s
}

func (s *ListDmAccountResponseBody) SetTotal(v int64) *ListDmAccountResponseBody {
	s.Total = &v
	return s
}

func (s *ListDmAccountResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDmAccountResponseBodyData struct {
	// The account name.
	//
	// example:
	//
	// send1@xx.xx.asia
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The time when the account was created.
	//
	// example:
	//
	// 1743579634000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the email address.
	//
	// example:
	//
	// a@b.com
	MailAddressId *string `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	// The type of the sender address. Valid values:
	//
	// - batch: batch emails
	//
	// - trigger: triggered emails
	//
	// example:
	//
	// trigger
	Sendtype *string `json:"Sendtype,omitempty" xml:"Sendtype,omitempty"`
}

func (s ListDmAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDmAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDmAccountResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *ListDmAccountResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDmAccountResponseBodyData) GetMailAddressId() *string {
	return s.MailAddressId
}

func (s *ListDmAccountResponseBodyData) GetSendtype() *string {
	return s.Sendtype
}

func (s *ListDmAccountResponseBodyData) SetAccountName(v string) *ListDmAccountResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListDmAccountResponseBodyData) SetCreateTime(v int64) *ListDmAccountResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListDmAccountResponseBodyData) SetMailAddressId(v string) *ListDmAccountResponseBodyData {
	s.MailAddressId = &v
	return s
}

func (s *ListDmAccountResponseBodyData) SetSendtype(v string) *ListDmAccountResponseBodyData {
	s.Sendtype = &v
	return s
}

func (s *ListDmAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
