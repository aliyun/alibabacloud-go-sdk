// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindInstagramPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindInstagramPageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *BindInstagramPageResponseBody
	GetCode() *string
	SetData(v []*BindInstagramPageResponseBodyData) *BindInstagramPageResponseBody
	GetData() []*BindInstagramPageResponseBodyData
	SetMessage(v string) *BindInstagramPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindInstagramPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindInstagramPageResponseBody
	GetSuccess() *bool
}

type BindInstagramPageResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// ok
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*BindInstagramPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// fdsfdsf-22fk***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindInstagramPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindInstagramPageResponseBody) GoString() string {
	return s.String()
}

func (s *BindInstagramPageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindInstagramPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindInstagramPageResponseBody) GetData() []*BindInstagramPageResponseBodyData {
	return s.Data
}

func (s *BindInstagramPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindInstagramPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindInstagramPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindInstagramPageResponseBody) SetAccessDeniedDetail(v string) *BindInstagramPageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindInstagramPageResponseBody) SetCode(v string) *BindInstagramPageResponseBody {
	s.Code = &v
	return s
}

func (s *BindInstagramPageResponseBody) SetData(v []*BindInstagramPageResponseBodyData) *BindInstagramPageResponseBody {
	s.Data = v
	return s
}

func (s *BindInstagramPageResponseBody) SetMessage(v string) *BindInstagramPageResponseBody {
	s.Message = &v
	return s
}

func (s *BindInstagramPageResponseBody) SetRequestId(v string) *BindInstagramPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindInstagramPageResponseBody) SetSuccess(v bool) *BindInstagramPageResponseBody {
	s.Success = &v
	return s
}

func (s *BindInstagramPageResponseBody) Validate() error {
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

type BindInstagramPageResponseBodyData struct {
	// example:
	//
	// 13243543543
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// connected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// example:
	//
	// 1234322333
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// example:
	//
	// iwhalecloud
	PageName *string `json:"PageName,omitempty" xml:"PageName,omitempty"`
}

func (s BindInstagramPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindInstagramPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindInstagramPageResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *BindInstagramPageResponseBodyData) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *BindInstagramPageResponseBodyData) GetPageId() *string {
	return s.PageId
}

func (s *BindInstagramPageResponseBodyData) GetPageName() *string {
	return s.PageName
}

func (s *BindInstagramPageResponseBodyData) SetAccountId(v string) *BindInstagramPageResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *BindInstagramPageResponseBodyData) SetConnectionStatus(v string) *BindInstagramPageResponseBodyData {
	s.ConnectionStatus = &v
	return s
}

func (s *BindInstagramPageResponseBodyData) SetPageId(v string) *BindInstagramPageResponseBodyData {
	s.PageId = &v
	return s
}

func (s *BindInstagramPageResponseBodyData) SetPageName(v string) *BindInstagramPageResponseBodyData {
	s.PageName = &v
	return s
}

func (s *BindInstagramPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
