// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmartAccessGatewayDnsForwardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteSmartAccessGatewayDnsForwardRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteSmartAccessGatewayDnsForwardRequest
	GetRegionId() *string
	SetSagInsId(v string) *DeleteSmartAccessGatewayDnsForwardRequest
	GetSagInsId() *string
	SetSagSn(v string) *DeleteSmartAccessGatewayDnsForwardRequest
	GetSagSn() *string
}

type DeleteSmartAccessGatewayDnsForwardRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sagv3dnsforward-nc7qabskj17werc7su
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-v9un1ccz22owd76lf8
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// The serial number (SN) of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sagf4dkqh78
	SagSn *string `json:"SagSn,omitempty" xml:"SagSn,omitempty"`
}

func (s DeleteSmartAccessGatewayDnsForwardRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmartAccessGatewayDnsForwardRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmartAccessGatewayDnsForwardRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteSmartAccessGatewayDnsForwardRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSmartAccessGatewayDnsForwardRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *DeleteSmartAccessGatewayDnsForwardRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *DeleteSmartAccessGatewayDnsForwardRequest) SetInstanceId(v string) *DeleteSmartAccessGatewayDnsForwardRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSmartAccessGatewayDnsForwardRequest) SetRegionId(v string) *DeleteSmartAccessGatewayDnsForwardRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSmartAccessGatewayDnsForwardRequest) SetSagInsId(v string) *DeleteSmartAccessGatewayDnsForwardRequest {
	s.SagInsId = &v
	return s
}

func (s *DeleteSmartAccessGatewayDnsForwardRequest) SetSagSn(v string) *DeleteSmartAccessGatewayDnsForwardRequest {
	s.SagSn = &v
	return s
}

func (s *DeleteSmartAccessGatewayDnsForwardRequest) Validate() error {
	return dara.Validate(s)
}
