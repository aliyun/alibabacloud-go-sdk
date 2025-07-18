// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExternalDataServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteExternalDataServiceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DeleteExternalDataServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *DeleteExternalDataServiceRequest
	GetServiceId() *string
}

type DeleteExternalDataServiceRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) API to view available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteExternalDataServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExternalDataServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteExternalDataServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteExternalDataServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteExternalDataServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DeleteExternalDataServiceRequest) SetDBInstanceId(v string) *DeleteExternalDataServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteExternalDataServiceRequest) SetRegionId(v string) *DeleteExternalDataServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteExternalDataServiceRequest) SetServiceId(v string) *DeleteExternalDataServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DeleteExternalDataServiceRequest) Validate() error {
	return dara.Validate(s)
}
