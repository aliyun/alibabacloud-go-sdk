// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateDnsEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceId(v string) *DeletePrivateDnsEndpointRequest
	GetAccessInstanceId() *string
	SetRegionNo(v string) *DeletePrivateDnsEndpointRequest
	GetRegionNo() *string
}

type DeletePrivateDnsEndpointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pd-12345
	AccessInstanceId *string `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DeletePrivateDnsEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateDnsEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivateDnsEndpointRequest) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *DeletePrivateDnsEndpointRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DeletePrivateDnsEndpointRequest) SetAccessInstanceId(v string) *DeletePrivateDnsEndpointRequest {
	s.AccessInstanceId = &v
	return s
}

func (s *DeletePrivateDnsEndpointRequest) SetRegionNo(v string) *DeletePrivateDnsEndpointRequest {
	s.RegionNo = &v
	return s
}

func (s *DeletePrivateDnsEndpointRequest) Validate() error {
	return dara.Validate(s)
}
