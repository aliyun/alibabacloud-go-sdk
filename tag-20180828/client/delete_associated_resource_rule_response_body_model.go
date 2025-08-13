// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAssociatedResourceRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAssociatedResourceRuleResponseBody
	GetRequestId() *string
}

type DeleteAssociatedResourceRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BB532282-94F5-5F56-877F-32D5E2A04F3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAssociatedResourceRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAssociatedResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAssociatedResourceRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAssociatedResourceRuleResponseBody) SetRequestId(v string) *DeleteAssociatedResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAssociatedResourceRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
