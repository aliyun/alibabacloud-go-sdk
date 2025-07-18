// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingDataServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyStreamingDataServiceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *ModifyStreamingDataServiceRequest
	GetRegionId() *string
	SetServiceDescription(v string) *ModifyStreamingDataServiceRequest
	GetServiceDescription() *string
	SetServiceId(v string) *ModifyStreamingDataServiceRequest
	GetServiceId() *string
	SetServiceSpec(v string) *ModifyStreamingDataServiceRequest
	GetServiceSpec() *string
}

type ModifyStreamingDataServiceRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
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
	// The description of the service.
	//
	// example:
	//
	// test-adbpgss
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The specifications of the service. Unit: capacity units (CUs). Valid values:
	//
	// 	- 2
	//
	// 	- 4
	//
	// 	- 8
	//
	// 	- 16
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ServiceSpec *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
}

func (s ModifyStreamingDataServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingDataServiceRequest) GoString() string {
	return s.String()
}

func (s *ModifyStreamingDataServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyStreamingDataServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyStreamingDataServiceRequest) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *ModifyStreamingDataServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ModifyStreamingDataServiceRequest) GetServiceSpec() *string {
	return s.ServiceSpec
}

func (s *ModifyStreamingDataServiceRequest) SetDBInstanceId(v string) *ModifyStreamingDataServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyStreamingDataServiceRequest) SetRegionId(v string) *ModifyStreamingDataServiceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyStreamingDataServiceRequest) SetServiceDescription(v string) *ModifyStreamingDataServiceRequest {
	s.ServiceDescription = &v
	return s
}

func (s *ModifyStreamingDataServiceRequest) SetServiceId(v string) *ModifyStreamingDataServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *ModifyStreamingDataServiceRequest) SetServiceSpec(v string) *ModifyStreamingDataServiceRequest {
	s.ServiceSpec = &v
	return s
}

func (s *ModifyStreamingDataServiceRequest) Validate() error {
	return dara.Validate(s)
}
