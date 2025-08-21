// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConnectableClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlreadySetItems(v bool) *DescribeConnectableClustersRequest
	GetAlreadySetItems() *bool
}

type DescribeConnectableClustersRequest struct {
	// The ID of the instance that can communicate with each other.
	//
	// example:
	//
	// true
	AlreadySetItems *bool `json:"alreadySetItems,omitempty" xml:"alreadySetItems,omitempty"`
}

func (s DescribeConnectableClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConnectableClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeConnectableClustersRequest) GetAlreadySetItems() *bool {
	return s.AlreadySetItems
}

func (s *DescribeConnectableClustersRequest) SetAlreadySetItems(v bool) *DescribeConnectableClustersRequest {
	s.AlreadySetItems = &v
	return s
}

func (s *DescribeConnectableClustersRequest) Validate() error {
	return dara.Validate(s)
}
