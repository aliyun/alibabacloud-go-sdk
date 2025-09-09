// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceADAuthServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceADAuthServerRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetInstanceADAuthServerRequest
	GetRegionId() *string
}

type GetInstanceADAuthServerRequest struct {
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInstanceADAuthServerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceADAuthServerRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceADAuthServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceADAuthServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceADAuthServerRequest) SetInstanceId(v string) *GetInstanceADAuthServerRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceADAuthServerRequest) SetRegionId(v string) *GetInstanceADAuthServerRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceADAuthServerRequest) Validate() error {
	return dara.Validate(s)
}
