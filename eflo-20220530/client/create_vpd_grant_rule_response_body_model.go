// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpdGrantRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateVpdGrantRuleResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *CreateVpdGrantRuleResponseBody
	GetCode() *int32
	SetContent(v *CreateVpdGrantRuleResponseBodyContent) *CreateVpdGrantRuleResponseBody
	GetContent() *CreateVpdGrantRuleResponseBodyContent
	SetMessage(v string) *CreateVpdGrantRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVpdGrantRuleResponseBody
	GetRequestId() *string
}

type CreateVpdGrantRuleResponseBody struct {
	// The details about the access denial.
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
	Content *CreateVpdGrantRuleResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// DBAD15D6-3F47-5B36-8A92-57C2919D13D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpdGrantRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpdGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpdGrantRuleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateVpdGrantRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateVpdGrantRuleResponseBody) GetContent() *CreateVpdGrantRuleResponseBodyContent {
	return s.Content
}

func (s *CreateVpdGrantRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVpdGrantRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpdGrantRuleResponseBody) SetAccessDeniedDetail(v string) *CreateVpdGrantRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateVpdGrantRuleResponseBody) SetCode(v int32) *CreateVpdGrantRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVpdGrantRuleResponseBody) SetContent(v *CreateVpdGrantRuleResponseBodyContent) *CreateVpdGrantRuleResponseBody {
	s.Content = v
	return s
}

func (s *CreateVpdGrantRuleResponseBody) SetMessage(v string) *CreateVpdGrantRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVpdGrantRuleResponseBody) SetRequestId(v string) *CreateVpdGrantRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpdGrantRuleResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVpdGrantRuleResponseBodyContent struct {
	// Authorized resource primary key ID
	//
	// example:
	//
	// grant-rule-hnevjkmw
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
}

func (s CreateVpdGrantRuleResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateVpdGrantRuleResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateVpdGrantRuleResponseBodyContent) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *CreateVpdGrantRuleResponseBodyContent) SetGrantRuleId(v string) *CreateVpdGrantRuleResponseBodyContent {
	s.GrantRuleId = &v
	return s
}

func (s *CreateVpdGrantRuleResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
