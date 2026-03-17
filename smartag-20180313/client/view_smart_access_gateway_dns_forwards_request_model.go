// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayDnsForwardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ViewSmartAccessGatewayDnsForwardsRequest
	GetRegionId() *string
	SetSagInsId(v string) *ViewSmartAccessGatewayDnsForwardsRequest
	GetSagInsId() *string
	SetSagSn(v string) *ViewSmartAccessGatewayDnsForwardsRequest
	GetSagSn() *string
}

type ViewSmartAccessGatewayDnsForwardsRequest struct {
	// The ID of the region in which the SAG instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance.
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

func (s ViewSmartAccessGatewayDnsForwardsRequest) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayDnsForwardsRequest) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayDnsForwardsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ViewSmartAccessGatewayDnsForwardsRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *ViewSmartAccessGatewayDnsForwardsRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *ViewSmartAccessGatewayDnsForwardsRequest) SetRegionId(v string) *ViewSmartAccessGatewayDnsForwardsRequest {
	s.RegionId = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsRequest) SetSagInsId(v string) *ViewSmartAccessGatewayDnsForwardsRequest {
	s.SagInsId = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsRequest) SetSagSn(v string) *ViewSmartAccessGatewayDnsForwardsRequest {
	s.SagSn = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsRequest) Validate() error {
	return dara.Validate(s)
}
