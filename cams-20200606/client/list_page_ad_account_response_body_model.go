// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPageAdAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListPageAdAccountResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListPageAdAccountResponseBody
	GetCode() *string
	SetData(v []*ListPageAdAccountResponseBodyData) *ListPageAdAccountResponseBody
	GetData() []*ListPageAdAccountResponseBodyData
	SetMessage(v string) *ListPageAdAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPageAdAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPageAdAccountResponseBody
	GetSuccess() *bool
}

type ListPageAdAccountResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, refer to [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data object.
	Data []*ListPageAdAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39**
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
}

func (s ListPageAdAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPageAdAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListPageAdAccountResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListPageAdAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPageAdAccountResponseBody) GetData() []*ListPageAdAccountResponseBodyData {
	return s.Data
}

func (s *ListPageAdAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPageAdAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPageAdAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPageAdAccountResponseBody) SetAccessDeniedDetail(v string) *ListPageAdAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListPageAdAccountResponseBody) SetCode(v string) *ListPageAdAccountResponseBody {
	s.Code = &v
	return s
}

func (s *ListPageAdAccountResponseBody) SetData(v []*ListPageAdAccountResponseBodyData) *ListPageAdAccountResponseBody {
	s.Data = v
	return s
}

func (s *ListPageAdAccountResponseBody) SetMessage(v string) *ListPageAdAccountResponseBody {
	s.Message = &v
	return s
}

func (s *ListPageAdAccountResponseBody) SetRequestId(v string) *ListPageAdAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPageAdAccountResponseBody) SetSuccess(v bool) *ListPageAdAccountResponseBody {
	s.Success = &v
	return s
}

func (s *ListPageAdAccountResponseBody) Validate() error {
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

type ListPageAdAccountResponseBodyData struct {
	// The Meta ad account ID.
	//
	// example:
	//
	// 29393****
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// The name of the ad account.
	//
	// example:
	//
	// Alibaba
	AdAccountName *string `json:"AdAccountName,omitempty" xml:"AdAccountName,omitempty"`
	// The currency.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The PageId of Messenger.
	//
	// example:
	//
	// 2030***
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
}

func (s ListPageAdAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPageAdAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPageAdAccountResponseBodyData) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *ListPageAdAccountResponseBodyData) GetAdAccountName() *string {
	return s.AdAccountName
}

func (s *ListPageAdAccountResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *ListPageAdAccountResponseBodyData) GetPageId() *string {
	return s.PageId
}

func (s *ListPageAdAccountResponseBodyData) SetAdAccountId(v string) *ListPageAdAccountResponseBodyData {
	s.AdAccountId = &v
	return s
}

func (s *ListPageAdAccountResponseBodyData) SetAdAccountName(v string) *ListPageAdAccountResponseBodyData {
	s.AdAccountName = &v
	return s
}

func (s *ListPageAdAccountResponseBodyData) SetCurrency(v string) *ListPageAdAccountResponseBodyData {
	s.Currency = &v
	return s
}

func (s *ListPageAdAccountResponseBodyData) SetPageId(v string) *ListPageAdAccountResponseBodyData {
	s.PageId = &v
	return s
}

func (s *ListPageAdAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
