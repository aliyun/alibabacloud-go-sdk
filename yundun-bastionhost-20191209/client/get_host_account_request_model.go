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
	SetRegionId(v string) *GetHostAccountRequest
	GetRegionId() *string
}

type GetHostAccountRequest struct {
	// The ID of the host account that you want to query.
	//
	// > You can call the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to query the ID of the host account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The ID of the bastion host in which you want to query the details of the host account.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host in which you want to query the details of the host account.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *GetHostAccountRequest) SetRegionId(v string) *GetHostAccountRequest {
	s.RegionId = &v
	return s
}

func (s *GetHostAccountRequest) Validate() error {
	return dara.Validate(s)
}
