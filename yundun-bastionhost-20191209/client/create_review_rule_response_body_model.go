// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReviewRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateReviewRuleResponseBody
	GetRequestId() *string
	SetReviewRuleId(v string) *CreateReviewRuleResponseBody
	GetReviewRuleId() *string
}

type CreateReviewRuleResponseBody struct {
	// The unique identifier that Alibaba Cloud generated for the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the audit policy.
	//
	// example:
	//
	// 1
	ReviewRuleId *string `json:"ReviewRuleId,omitempty" xml:"ReviewRuleId,omitempty"`
}

func (s CreateReviewRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReviewRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReviewRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateReviewRuleResponseBody) GetReviewRuleId() *string {
	return s.ReviewRuleId
}

func (s *CreateReviewRuleResponseBody) SetRequestId(v string) *CreateReviewRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReviewRuleResponseBody) SetReviewRuleId(v string) *CreateReviewRuleResponseBody {
	s.ReviewRuleId = &v
	return s
}

func (s *CreateReviewRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
