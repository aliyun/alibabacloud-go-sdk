// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFbInstagramPagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetFbInstagramPagesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetFbInstagramPagesResponseBody
	GetCode() *string
	SetData(v []*GetFbInstagramPagesResponseBodyData) *GetFbInstagramPagesResponseBody
	GetData() []*GetFbInstagramPagesResponseBodyData
	SetMessage(v string) *GetFbInstagramPagesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFbInstagramPagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFbInstagramPagesResponseBody
	GetSuccess() *bool
}

type GetFbInstagramPagesResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetFbInstagramPagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// hgfh77-gfh55***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFbInstagramPagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFbInstagramPagesResponseBody) GoString() string {
	return s.String()
}

func (s *GetFbInstagramPagesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetFbInstagramPagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFbInstagramPagesResponseBody) GetData() []*GetFbInstagramPagesResponseBodyData {
	return s.Data
}

func (s *GetFbInstagramPagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFbInstagramPagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFbInstagramPagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFbInstagramPagesResponseBody) SetAccessDeniedDetail(v string) *GetFbInstagramPagesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetFbInstagramPagesResponseBody) SetCode(v string) *GetFbInstagramPagesResponseBody {
	s.Code = &v
	return s
}

func (s *GetFbInstagramPagesResponseBody) SetData(v []*GetFbInstagramPagesResponseBodyData) *GetFbInstagramPagesResponseBody {
	s.Data = v
	return s
}

func (s *GetFbInstagramPagesResponseBody) SetMessage(v string) *GetFbInstagramPagesResponseBody {
	s.Message = &v
	return s
}

func (s *GetFbInstagramPagesResponseBody) SetRequestId(v string) *GetFbInstagramPagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFbInstagramPagesResponseBody) SetSuccess(v bool) *GetFbInstagramPagesResponseBody {
	s.Success = &v
	return s
}

func (s *GetFbInstagramPagesResponseBody) Validate() error {
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

type GetFbInstagramPagesResponseBodyData struct {
	// example:
	//
	// 1245454654654
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// Alice
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// Connected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// example:
	//
	// 12534653543
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// example:
	//
	// iwhalecloud
	PageName *string `json:"PageName,omitempty" xml:"PageName,omitempty"`
}

func (s GetFbInstagramPagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFbInstagramPagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFbInstagramPagesResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *GetFbInstagramPagesResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *GetFbInstagramPagesResponseBodyData) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *GetFbInstagramPagesResponseBodyData) GetPageId() *string {
	return s.PageId
}

func (s *GetFbInstagramPagesResponseBodyData) GetPageName() *string {
	return s.PageName
}

func (s *GetFbInstagramPagesResponseBodyData) SetAccountId(v string) *GetFbInstagramPagesResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetFbInstagramPagesResponseBodyData) SetAccountName(v string) *GetFbInstagramPagesResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *GetFbInstagramPagesResponseBodyData) SetConnectionStatus(v string) *GetFbInstagramPagesResponseBodyData {
	s.ConnectionStatus = &v
	return s
}

func (s *GetFbInstagramPagesResponseBodyData) SetPageId(v string) *GetFbInstagramPagesResponseBodyData {
	s.PageId = &v
	return s
}

func (s *GetFbInstagramPagesResponseBodyData) SetPageName(v string) *GetFbInstagramPagesResponseBodyData {
	s.PageName = &v
	return s
}

func (s *GetFbInstagramPagesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
