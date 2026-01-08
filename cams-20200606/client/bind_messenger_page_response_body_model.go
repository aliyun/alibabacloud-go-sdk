// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindMessengerPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindMessengerPageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *BindMessengerPageResponseBody
	GetCode() *string
	SetData(v []*BindMessengerPageResponseBodyData) *BindMessengerPageResponseBody
	GetData() []*BindMessengerPageResponseBodyData
	SetMessage(v string) *BindMessengerPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindMessengerPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindMessengerPageResponseBody
	GetSuccess() *bool
}

type BindMessengerPageResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// ok
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*BindMessengerPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DSFDS-8FSDFS**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindMessengerPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindMessengerPageResponseBody) GoString() string {
	return s.String()
}

func (s *BindMessengerPageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindMessengerPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindMessengerPageResponseBody) GetData() []*BindMessengerPageResponseBodyData {
	return s.Data
}

func (s *BindMessengerPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindMessengerPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindMessengerPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindMessengerPageResponseBody) SetAccessDeniedDetail(v string) *BindMessengerPageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindMessengerPageResponseBody) SetCode(v string) *BindMessengerPageResponseBody {
	s.Code = &v
	return s
}

func (s *BindMessengerPageResponseBody) SetData(v []*BindMessengerPageResponseBodyData) *BindMessengerPageResponseBody {
	s.Data = v
	return s
}

func (s *BindMessengerPageResponseBody) SetMessage(v string) *BindMessengerPageResponseBody {
	s.Message = &v
	return s
}

func (s *BindMessengerPageResponseBody) SetRequestId(v string) *BindMessengerPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindMessengerPageResponseBody) SetSuccess(v bool) *BindMessengerPageResponseBody {
	s.Success = &v
	return s
}

func (s *BindMessengerPageResponseBody) Validate() error {
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

type BindMessengerPageResponseBodyData struct {
	// example:
	//
	// CONNECTED
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// example:
	//
	// 1654543543543
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// example:
	//
	// iwhalecloud
	PageName *string `json:"PageName,omitempty" xml:"PageName,omitempty"`
}

func (s BindMessengerPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindMessengerPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindMessengerPageResponseBodyData) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *BindMessengerPageResponseBodyData) GetPageId() *string {
	return s.PageId
}

func (s *BindMessengerPageResponseBodyData) GetPageName() *string {
	return s.PageName
}

func (s *BindMessengerPageResponseBodyData) SetConnectionStatus(v string) *BindMessengerPageResponseBodyData {
	s.ConnectionStatus = &v
	return s
}

func (s *BindMessengerPageResponseBodyData) SetPageId(v string) *BindMessengerPageResponseBodyData {
	s.PageId = &v
	return s
}

func (s *BindMessengerPageResponseBodyData) SetPageName(v string) *BindMessengerPageResponseBodyData {
	s.PageName = &v
	return s
}

func (s *BindMessengerPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
