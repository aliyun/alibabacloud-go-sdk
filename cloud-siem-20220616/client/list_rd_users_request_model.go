// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRdUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListRdUsersRequest
	GetRegionId() *string
}

type ListRdUsersRequest struct {
	// The region where the Data Management center of Threat Analysis is located. Select the region of the Management Center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRdUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRdUsersRequest) GoString() string {
	return s.String()
}

func (s *ListRdUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRdUsersRequest) SetRegionId(v string) *ListRdUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListRdUsersRequest) Validate() error {
	return dara.Validate(s)
}
