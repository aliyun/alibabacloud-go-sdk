// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkAccessEndpointAvailableZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNaeRegionId(v string) *ListNetworkAccessEndpointAvailableZonesRequest
	GetNaeRegionId() *string
}

type ListNetworkAccessEndpointAvailableZonesRequest struct {
	// 专属网络端点支持的地域
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	NaeRegionId *string `json:"NaeRegionId,omitempty" xml:"NaeRegionId,omitempty"`
}

func (s ListNetworkAccessEndpointAvailableZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableZonesRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableZonesRequest) GetNaeRegionId() *string {
	return s.NaeRegionId
}

func (s *ListNetworkAccessEndpointAvailableZonesRequest) SetNaeRegionId(v string) *ListNetworkAccessEndpointAvailableZonesRequest {
	s.NaeRegionId = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableZonesRequest) Validate() error {
	return dara.Validate(s)
}
