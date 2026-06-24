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
	// Specifies whether to return instances that are already connected. Valid values:
	//
	// - true (default): The returned instance list includes instances that are already connected.
	//
	// - false: The returned instance list does not include instances that are already connected.
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
