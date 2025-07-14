// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCustomDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DescribeWebCustomDomainRequest
	GetNamespaceId() *string
}

type DescribeWebCustomDomainRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeWebCustomDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebCustomDomainRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeWebCustomDomainRequest) SetNamespaceId(v string) *DescribeWebCustomDomainRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeWebCustomDomainRequest) Validate() error {
	return dara.Validate(s)
}
