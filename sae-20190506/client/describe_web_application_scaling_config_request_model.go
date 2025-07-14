// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebApplicationScalingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DescribeWebApplicationScalingConfigRequest
	GetNamespaceId() *string
}

type DescribeWebApplicationScalingConfigRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:sae-test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeWebApplicationScalingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebApplicationScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebApplicationScalingConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeWebApplicationScalingConfigRequest) SetNamespaceId(v string) *DescribeWebApplicationScalingConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeWebApplicationScalingConfigRequest) Validate() error {
	return dara.Validate(s)
}
