// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountId(v string) *GetHostAccountRequest
	GetHostAccountId() *string
	SetInstanceId(v string) *GetHostAccountRequest
	GetInstanceId() *string
	SetProjectId(v int64) *GetHostAccountRequest
	GetProjectId() *int64
	SetRegionId(v string) *GetHostAccountRequest
	GetRegionId() *string
}

type GetHostAccountRequest struct {
	// The ID of the host account to query.
	//
	// > You can call the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The instance ID of the bastion host to which the host account to query belongs.
	//
	// >You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
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
	// The region ID of the bastion host instance to which the host account to query belongs.
	//
	// >For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHostAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHostAccountRequest) GoString() string {
	return s.String()
}

func (s *GetHostAccountRequest) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *GetHostAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHostAccountRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetHostAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetHostAccountRequest) SetHostAccountId(v string) *GetHostAccountRequest {
	s.HostAccountId = &v
	return s
}

func (s *GetHostAccountRequest) SetInstanceId(v string) *GetHostAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHostAccountRequest) SetProjectId(v int64) *GetHostAccountRequest {
	s.ProjectId = &v
	return s
}

func (s *GetHostAccountRequest) SetRegionId(v string) *GetHostAccountRequest {
	s.RegionId = &v
	return s
}

func (s *GetHostAccountRequest) Validate() error {
	return dara.Validate(s)
}
