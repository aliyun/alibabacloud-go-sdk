// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenArmsServiceSecondVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *OpenArmsServiceSecondVersionRequest
	GetRegionId() *string
	SetType(v string) *OpenArmsServiceSecondVersionRequest
	GetType() *string
}

type OpenArmsServiceSecondVersionRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- `arms`: ARMS
	//
	// 	- `arms_app`: Application Monitoring
	//
	// 	- `arms_web`: Browser Monitoring
	//
	// 	- `prometheus_monitor`: Managed Service for Prometheus
	//
	// 	- `synthetic_post`: Synthetic Monitoring
	//
	// This parameter is required.
	//
	// example:
	//
	// arms
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OpenArmsServiceSecondVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenArmsServiceSecondVersionRequest) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceSecondVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenArmsServiceSecondVersionRequest) GetType() *string {
	return s.Type
}

func (s *OpenArmsServiceSecondVersionRequest) SetRegionId(v string) *OpenArmsServiceSecondVersionRequest {
	s.RegionId = &v
	return s
}

func (s *OpenArmsServiceSecondVersionRequest) SetType(v string) *OpenArmsServiceSecondVersionRequest {
	s.Type = &v
	return s
}

func (s *OpenArmsServiceSecondVersionRequest) Validate() error {
	return dara.Validate(s)
}
