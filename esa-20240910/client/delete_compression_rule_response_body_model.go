// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCompressionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCompressionRuleResponseBody
	GetRequestId() *string
}

type DeleteCompressionRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCompressionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCompressionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCompressionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCompressionRuleResponseBody) SetRequestId(v string) *DeleteCompressionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCompressionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
