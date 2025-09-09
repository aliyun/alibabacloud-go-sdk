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
	SetRegionId(v string) *AttachHostAccountsToHostShareKeyRequest
	GetRegionId() *string
}

type AttachHostAccountsToHostShareKeyRequest struct {
	// The host account IDs.
	//
	// >  You must specify this parameter. You can call the [ListHostAccounts](https://help.aliyun.com/document_detail/462937.html) operation to query the host account IDs.
	//
	// example:
	//
	// ["1","2","3"]
	HostAccountIds *string `json:"HostAccountIds,omitempty" xml:"HostAccountIds,omitempty"`
	// The shared key ID.
	//
	// >  You must specify this parameter. You can call the [ListHostShareKeys](https://help.aliyun.com/document_detail/462973.html) operation to query the shared key ID.
	//
	// example:
	//
	// 10267
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the bastion host. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host. For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *AttachHostAccountsToHostShareKeyRequest) SetRegionId(v string) *AttachHostAccountsToHostShareKeyRequest {
	s.RegionId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyRequest) Validate() error {
	return dara.Validate(s)
}
