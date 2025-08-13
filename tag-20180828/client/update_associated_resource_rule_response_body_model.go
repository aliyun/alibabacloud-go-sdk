// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssociatedResourceRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAssociatedResourceRuleResponseBody
	GetRequestId() *string
}

type UpdateAssociatedResourceRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 94E16BB6-3FB6-1297-B5B2-ED2250F437CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAssociatedResourceRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssociatedResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAssociatedResourceRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAssociatedResourceRuleResponseBody) SetRequestId(v string) *UpdateAssociatedResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAssociatedResourceRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
