// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInterceptionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteInterceptionRuleResponseBody
	GetRequestId() *string
}

type DeleteInterceptionRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53CXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInterceptionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInterceptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInterceptionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInterceptionRuleResponseBody) SetRequestId(v string) *DeleteInterceptionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInterceptionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
