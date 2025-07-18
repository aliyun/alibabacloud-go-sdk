// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamingDataServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteStreamingDataServiceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DeleteStreamingDataServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *DeleteStreamingDataServiceRequest
	GetServiceId() *string
}

type DeleteStreamingDataServiceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteStreamingDataServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamingDataServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteStreamingDataServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteStreamingDataServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteStreamingDataServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DeleteStreamingDataServiceRequest) SetDBInstanceId(v string) *DeleteStreamingDataServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteStreamingDataServiceRequest) SetRegionId(v string) *DeleteStreamingDataServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStreamingDataServiceRequest) SetServiceId(v string) *DeleteStreamingDataServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DeleteStreamingDataServiceRequest) Validate() error {
	return dara.Validate(s)
}
