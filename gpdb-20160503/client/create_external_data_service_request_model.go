// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalDataServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateExternalDataServiceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *CreateExternalDataServiceRequest
	GetRegionId() *string
	SetServiceDescription(v string) *CreateExternalDataServiceRequest
	GetServiceDescription() *string
	SetServiceName(v string) *CreateExternalDataServiceRequest
	GetServiceName() *string
	SetServiceSpec(v string) *CreateExternalDataServiceRequest
	GetServiceSpec() *string
}

type CreateExternalDataServiceRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Region ID, you can view available region IDs through the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) interface.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service description.
	//
	// example:
	//
	// pxf test
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// Service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-adbpgss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// Service specification (in CU), value:
	//
	// - 8
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	ServiceSpec *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
}

func (s CreateExternalDataServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalDataServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateExternalDataServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateExternalDataServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateExternalDataServiceRequest) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *CreateExternalDataServiceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateExternalDataServiceRequest) GetServiceSpec() *string {
	return s.ServiceSpec
}

func (s *CreateExternalDataServiceRequest) SetDBInstanceId(v string) *CreateExternalDataServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateExternalDataServiceRequest) SetRegionId(v string) *CreateExternalDataServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateExternalDataServiceRequest) SetServiceDescription(v string) *CreateExternalDataServiceRequest {
	s.ServiceDescription = &v
	return s
}

func (s *CreateExternalDataServiceRequest) SetServiceName(v string) *CreateExternalDataServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateExternalDataServiceRequest) SetServiceSpec(v string) *CreateExternalDataServiceRequest {
	s.ServiceSpec = &v
	return s
}

func (s *CreateExternalDataServiceRequest) Validate() error {
	return dara.Validate(s)
}
