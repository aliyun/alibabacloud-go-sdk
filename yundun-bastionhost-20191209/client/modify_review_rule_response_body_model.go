// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReviewRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyReviewRuleResponseBody
	GetRequestId() *string
}

type ModifyReviewRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7E68165E-1191-5CC2-B54B-5EF7390A5400
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyReviewRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyReviewRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyReviewRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyReviewRuleResponseBody) SetRequestId(v string) *ModifyReviewRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyReviewRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
