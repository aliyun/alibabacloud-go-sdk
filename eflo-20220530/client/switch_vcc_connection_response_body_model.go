// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchVccConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SwitchVccConnectionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *SwitchVccConnectionResponseBody
	GetCode() *int32
	SetContent(v interface{}) *SwitchVccConnectionResponseBody
	GetContent() interface{}
	SetMessage(v string) *SwitchVccConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *SwitchVccConnectionResponseBody
	GetRequestId() *string
}

type SwitchVccConnectionResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Code number
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response content
	//
	// example:
	//
	// {}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// Return message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF3EFCD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchVccConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchVccConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchVccConnectionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SwitchVccConnectionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SwitchVccConnectionResponseBody) GetContent() interface{} {
	return s.Content
}

func (s *SwitchVccConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SwitchVccConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchVccConnectionResponseBody) SetAccessDeniedDetail(v string) *SwitchVccConnectionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SwitchVccConnectionResponseBody) SetCode(v int32) *SwitchVccConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *SwitchVccConnectionResponseBody) SetContent(v interface{}) *SwitchVccConnectionResponseBody {
	s.Content = v
	return s
}

func (s *SwitchVccConnectionResponseBody) SetMessage(v string) *SwitchVccConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *SwitchVccConnectionResponseBody) SetRequestId(v string) *SwitchVccConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchVccConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
