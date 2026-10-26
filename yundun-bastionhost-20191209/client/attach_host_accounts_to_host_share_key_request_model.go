// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostAccountsToHostShareKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountIds(v string) *AttachHostAccountsToHostShareKeyRequest
	GetHostAccountIds() *string
	SetHostShareKeyId(v string) *AttachHostAccountsToHostShareKeyRequest
	GetHostShareKeyId() *string
	SetInstanceId(v string) *AttachHostAccountsToHostShareKeyRequest
	GetInstanceId() *string
	SetProjectId(v int64) *AttachHostAccountsToHostShareKeyRequest
	GetProjectId() *int64
	SetRegionId(v string) *AttachHostAccountsToHostShareKeyRequest
	GetRegionId() *string
}

type AttachHostAccountsToHostShareKeyRequest struct {
	// The list of host account IDs.
	//
	// > This parameter is required. You can call the [ListHostAccounts](https://help.aliyun.com/document_detail/462937.html) operation to obtain this parameter.
	//
	// example:
	//
	// ["1","2","3"]
	HostAccountIds *string `json:"HostAccountIds,omitempty" xml:"HostAccountIds,omitempty"`
	// The ID of the host shared key.
	//
	// > This parameter is required. You can call the [ListHostShareKeys](https://help.aliyun.com/document_detail/462973.html) operation to obtain this parameter.
	//
	// example:
	//
	// 10267
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The instance ID of the bastion host. You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
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
	// The region ID of the bastion host. For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachHostAccountsToHostShareKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToHostShareKeyRequest) GetHostAccountIds() *string {
	return s.HostAccountIds
}

func (s *AttachHostAccountsToHostShareKeyRequest) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *AttachHostAccountsToHostShareKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachHostAccountsToHostShareKeyRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AttachHostAccountsToHostShareKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachHostAccountsToHostShareKeyRequest) SetHostAccountIds(v string) *AttachHostAccountsToHostShareKeyRequest {
	s.HostAccountIds = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyRequest) SetHostShareKeyId(v string) *AttachHostAccountsToHostShareKeyRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyRequest) SetInstanceId(v string) *AttachHostAccountsToHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyRequest) SetProjectId(v int64) *AttachHostAccountsToHostShareKeyRequest {
	s.ProjectId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyRequest) SetRegionId(v string) *AttachHostAccountsToHostShareKeyRequest {
	s.RegionId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyRequest) Validate() error {
	return dara.Validate(s)
}
