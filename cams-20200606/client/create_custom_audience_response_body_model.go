// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAudienceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateCustomAudienceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateCustomAudienceResponseBody
	GetCode() *string
	SetCustomAudienceId(v string) *CreateCustomAudienceResponseBody
	GetCustomAudienceId() *string
	SetMessage(v string) *CreateCustomAudienceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCustomAudienceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCustomAudienceResponseBody
	GetSuccess() *bool
}

type CreateCustomAudienceResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 393**
	CustomAudienceId *string `json:"CustomAudienceId,omitempty" xml:"CustomAudienceId,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// dd**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCustomAudienceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAudienceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomAudienceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateCustomAudienceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCustomAudienceResponseBody) GetCustomAudienceId() *string {
	return s.CustomAudienceId
}

func (s *CreateCustomAudienceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCustomAudienceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomAudienceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCustomAudienceResponseBody) SetAccessDeniedDetail(v string) *CreateCustomAudienceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateCustomAudienceResponseBody) SetCode(v string) *CreateCustomAudienceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomAudienceResponseBody) SetCustomAudienceId(v string) *CreateCustomAudienceResponseBody {
	s.CustomAudienceId = &v
	return s
}

func (s *CreateCustomAudienceResponseBody) SetMessage(v string) *CreateCustomAudienceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCustomAudienceResponseBody) SetRequestId(v string) *CreateCustomAudienceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomAudienceResponseBody) SetSuccess(v bool) *CreateCustomAudienceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomAudienceResponseBody) Validate() error {
	return dara.Validate(s)
}
