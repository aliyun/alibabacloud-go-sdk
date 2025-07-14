// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNameSpaceShortId(v string) *DescribeNamespaceRequest
	GetNameSpaceShortId() *string
	SetNamespaceId(v string) *DescribeNamespaceRequest
	GetNamespaceId() *string
}

type DescribeNamespaceRequest struct {
	// The short ID of the namespace. You do not need to specify a region ID. The value of this parameter can be up to 20 characters in length and can contain only lowercase letters and digits.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The ID of the namespace. The information about the default namespace cannot be queried or modified. The default namespace cannot be deleted.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceRequest) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *DescribeNamespaceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeNamespaceRequest) SetNameSpaceShortId(v string) *DescribeNamespaceRequest {
	s.NameSpaceShortId = &v
	return s
}

func (s *DescribeNamespaceRequest) SetNamespaceId(v string) *DescribeNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
