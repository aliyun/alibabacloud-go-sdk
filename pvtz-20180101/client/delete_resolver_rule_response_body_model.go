// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResolverRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResolverRuleResponseBody
	GetRequestId() *string
}

type DeleteResolverRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0C9959BE-3A6A-4803-8DCE-973B42ACD599
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResolverRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResolverRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResolverRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResolverRuleResponseBody) SetRequestId(v string) *DeleteResolverRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResolverRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
