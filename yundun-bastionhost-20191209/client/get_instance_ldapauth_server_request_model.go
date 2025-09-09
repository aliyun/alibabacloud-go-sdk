// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLDAPAuthServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceLDAPAuthServerRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetInstanceLDAPAuthServerRequest
	GetRegionId() *string
}

type GetInstanceLDAPAuthServerRequest struct {
	// The ID of the bastion host.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInstanceLDAPAuthServerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLDAPAuthServerRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceLDAPAuthServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceLDAPAuthServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceLDAPAuthServerRequest) SetInstanceId(v string) *GetInstanceLDAPAuthServerRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceLDAPAuthServerRequest) SetRegionId(v string) *GetInstanceLDAPAuthServerRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceLDAPAuthServerRequest) Validate() error {
	return dara.Validate(s)
}
