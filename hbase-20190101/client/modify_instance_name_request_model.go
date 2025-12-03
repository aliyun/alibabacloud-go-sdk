// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyInstanceNameRequest
	GetClientToken() *string
	SetClusterId(v string) *ModifyInstanceNameRequest
	GetClusterId() *string
	SetClusterName(v string) *ModifyInstanceNameRequest
	GetClusterName() *string
	SetRegionId(v string) *ModifyInstanceNameRequest
	GetRegionId() *string
	SetZoneId(v string) *ModifyInstanceNameRequest
	GetZoneId() *string
}

type ModifyInstanceNameRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testhbaseone
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyInstanceNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceNameRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyInstanceNameRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ModifyInstanceNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceNameRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyInstanceNameRequest) SetClientToken(v string) *ModifyInstanceNameRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetClusterId(v string) *ModifyInstanceNameRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetClusterName(v string) *ModifyInstanceNameRequest {
	s.ClusterName = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetRegionId(v string) *ModifyInstanceNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetZoneId(v string) *ModifyInstanceNameRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyInstanceNameRequest) Validate() error {
	return dara.Validate(s)
}
