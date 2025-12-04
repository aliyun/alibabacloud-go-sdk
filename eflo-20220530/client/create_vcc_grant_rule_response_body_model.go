// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVccGrantRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateVccGrantRuleResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *CreateVccGrantRuleResponseBody
	GetCode() *int32
	SetContent(v *CreateVccGrantRuleResponseBodyContent) *CreateVccGrantRuleResponseBody
	GetContent() *CreateVccGrantRuleResponseBodyContent
	SetMessage(v string) *CreateVccGrantRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVccGrantRuleResponseBody
	GetRequestId() *string
}

type CreateVccGrantRuleResponseBody struct {
	// The details about the access denial. This parameter is returned only if Resource Access Management (RAM) permission verification failed.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// {}
	Content *CreateVccGrantRuleResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVccGrantRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVccGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVccGrantRuleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateVccGrantRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateVccGrantRuleResponseBody) GetContent() *CreateVccGrantRuleResponseBodyContent {
	return s.Content
}

func (s *CreateVccGrantRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVccGrantRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVccGrantRuleResponseBody) SetAccessDeniedDetail(v string) *CreateVccGrantRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateVccGrantRuleResponseBody) SetCode(v int32) *CreateVccGrantRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVccGrantRuleResponseBody) SetContent(v *CreateVccGrantRuleResponseBodyContent) *CreateVccGrantRuleResponseBody {
	s.Content = v
	return s
}

func (s *CreateVccGrantRuleResponseBody) SetMessage(v string) *CreateVccGrantRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVccGrantRuleResponseBody) SetRequestId(v string) *CreateVccGrantRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVccGrantRuleResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVccGrantRuleResponseBodyContent struct {
	// Authorized resource primary key ID
	//
	// example:
	//
	// grant-rule-8rgvqazb
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
}

func (s CreateVccGrantRuleResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateVccGrantRuleResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateVccGrantRuleResponseBodyContent) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *CreateVccGrantRuleResponseBodyContent) SetGrantRuleId(v string) *CreateVccGrantRuleResponseBodyContent {
	s.GrantRuleId = &v
	return s
}

func (s *CreateVccGrantRuleResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
