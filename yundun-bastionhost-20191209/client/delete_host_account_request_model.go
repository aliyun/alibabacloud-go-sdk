// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountId(v string) *DeleteHostAccountRequest
	GetHostAccountId() *string
	SetInstanceId(v string) *DeleteHostAccountRequest
	GetInstanceId() *string
	SetProjectId(v int64) *DeleteHostAccountRequest
	GetProjectId() *int64
	SetRegionId(v string) *DeleteHostAccountRequest
	GetRegionId() *string
}

type DeleteHostAccountRequest struct {
	// The ID of the host account that you want to delete.
	//
	// > You can call the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The ID of the Bastionhost instance where the host account that you want to delete resides.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance where the host account resides.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteHostAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostAccountRequest) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *DeleteHostAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteHostAccountRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteHostAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHostAccountRequest) SetHostAccountId(v string) *DeleteHostAccountRequest {
	s.HostAccountId = &v
	return s
}

func (s *DeleteHostAccountRequest) SetInstanceId(v string) *DeleteHostAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHostAccountRequest) SetProjectId(v int64) *DeleteHostAccountRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteHostAccountRequest) SetRegionId(v string) *DeleteHostAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHostAccountRequest) Validate() error {
	return dara.Validate(s)
}
