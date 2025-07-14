// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNameSpaceShortId(v string) *DescribeNamespaceResourcesRequest
	GetNameSpaceShortId() *string
	SetNamespaceId(v string) *DescribeNamespaceResourcesRequest
	GetNamespaceId() *string
}

type DescribeNamespaceResourcesRequest struct {
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// cn-shanghai:test
	//
	// example:
	//
	// cn-shanghai:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeNamespaceResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResourcesRequest) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *DescribeNamespaceResourcesRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeNamespaceResourcesRequest) SetNameSpaceShortId(v string) *DescribeNamespaceResourcesRequest {
	s.NameSpaceShortId = &v
	return s
}

func (s *DescribeNamespaceResourcesRequest) SetNamespaceId(v string) *DescribeNamespaceResourcesRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeNamespaceResourcesRequest) Validate() error {
	return dara.Validate(s)
}
