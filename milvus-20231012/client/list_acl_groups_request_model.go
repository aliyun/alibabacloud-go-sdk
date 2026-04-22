// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAclGroupsRequest
	GetInstanceId() *string
	SetRegionId(v string) *ListAclGroupsRequest
	GetRegionId() *string
}

type ListAclGroupsRequest struct {
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListAclGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAclGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAclGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAclGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAclGroupsRequest) SetInstanceId(v string) *ListAclGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAclGroupsRequest) SetRegionId(v string) *ListAclGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAclGroupsRequest) Validate() error {
	return dara.Validate(s)
}
