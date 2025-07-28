// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindResolverRuleVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindResolverRuleVpcResponseBody
	GetRequestId() *string
}

type BindResolverRuleVpcResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 12FE6E98-3885-423E-B18B-88CC17052A31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindResolverRuleVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindResolverRuleVpcResponseBody) GoString() string {
	return s.String()
}

func (s *BindResolverRuleVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindResolverRuleVpcResponseBody) SetRequestId(v string) *BindResolverRuleVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindResolverRuleVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
