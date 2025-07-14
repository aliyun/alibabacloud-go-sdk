// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DescribeWebApplicationRequest
	GetNamespaceId() *string
}

type DescribeWebApplicationRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeWebApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebApplicationRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebApplicationRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeWebApplicationRequest) SetNamespaceId(v string) *DescribeWebApplicationRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeWebApplicationRequest) Validate() error {
	return dara.Validate(s)
}
