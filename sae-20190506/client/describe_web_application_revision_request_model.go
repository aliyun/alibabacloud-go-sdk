// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebApplicationRevisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DescribeWebApplicationRevisionRequest
	GetNamespaceId() *string
}

type DescribeWebApplicationRevisionRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeWebApplicationRevisionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebApplicationRevisionRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebApplicationRevisionRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeWebApplicationRevisionRequest) SetNamespaceId(v string) *DescribeWebApplicationRevisionRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeWebApplicationRevisionRequest) Validate() error {
	return dara.Validate(s)
}
