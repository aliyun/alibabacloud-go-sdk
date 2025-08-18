// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCompressionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCompressionRuleResponseBody
	GetRequestId() *string
}

type UpdateCompressionRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// CE8EB0DE-3267-53D1-AB83-F36377D63FD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCompressionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompressionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCompressionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCompressionRuleResponseBody) SetRequestId(v string) *UpdateCompressionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCompressionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
