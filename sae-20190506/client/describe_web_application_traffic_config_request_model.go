// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebApplicationTrafficConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DescribeWebApplicationTrafficConfigRequest
	GetNamespaceId() *string
}

type DescribeWebApplicationTrafficConfigRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeWebApplicationTrafficConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebApplicationTrafficConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebApplicationTrafficConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeWebApplicationTrafficConfigRequest) SetNamespaceId(v string) *DescribeWebApplicationTrafficConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeWebApplicationTrafficConfigRequest) Validate() error {
	return dara.Validate(s)
}
