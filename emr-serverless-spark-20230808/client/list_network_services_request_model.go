// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListNetworkServicesRequest
	GetRegionId() *string
}

type ListNetworkServicesRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListNetworkServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkServicesRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNetworkServicesRequest) SetRegionId(v string) *ListNetworkServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListNetworkServicesRequest) Validate() error {
	return dara.Validate(s)
}
