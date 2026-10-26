// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostAccountsFromHostShareKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountIds(v string) *DetachHostAccountsFromHostShareKeyRequest
	GetHostAccountIds() *string
	SetHostShareKeyId(v string) *DetachHostAccountsFromHostShareKeyRequest
	GetHostShareKeyId() *string
	SetInstanceId(v string) *DetachHostAccountsFromHostShareKeyRequest
	GetInstanceId() *string
	SetProjectId(v int64) *DetachHostAccountsFromHostShareKeyRequest
	GetProjectId() *int64
	SetRegionId(v string) *DetachHostAccountsFromHostShareKeyRequest
	GetRegionId() *string
}

type DetachHostAccountsFromHostShareKeyRequest struct {
	// The list of host account IDs.
	//
	// > You can call [ListHostAccountsForHostShareKey](https://help.aliyun.com/document_detail/462975.html) to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2","3"]
	HostAccountIds *string `json:"HostAccountIds,omitempty" xml:"HostAccountIds,omitempty"`
	// The ID of the host shared key.
	//
	// > You can call [ListHostShareKeys](https://help.aliyun.com/document_detail/462973.html) to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can invoke [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) to obtain this parameter.
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
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachHostAccountsFromHostShareKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromHostShareKeyRequest) GetHostAccountIds() *string {
	return s.HostAccountIds
}

func (s *DetachHostAccountsFromHostShareKeyRequest) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *DetachHostAccountsFromHostShareKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachHostAccountsFromHostShareKeyRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DetachHostAccountsFromHostShareKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetHostAccountIds(v string) *DetachHostAccountsFromHostShareKeyRequest {
	s.HostAccountIds = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetHostShareKeyId(v string) *DetachHostAccountsFromHostShareKeyRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetInstanceId(v string) *DetachHostAccountsFromHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetProjectId(v int64) *DetachHostAccountsFromHostShareKeyRequest {
	s.ProjectId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetRegionId(v string) *DetachHostAccountsFromHostShareKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyRequest) Validate() error {
	return dara.Validate(s)
}
