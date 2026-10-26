// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachReviewersToReviewRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachReviewersToReviewRuleResponseBody
	GetRequestId() *string
}

type AttachReviewersToReviewRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachReviewersToReviewRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachReviewersToReviewRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AttachReviewersToReviewRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachReviewersToReviewRuleResponseBody) SetRequestId(v string) *AttachReviewersToReviewRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachReviewersToReviewRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
