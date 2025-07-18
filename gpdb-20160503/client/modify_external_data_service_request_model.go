// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExternalDataServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyExternalDataServiceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *ModifyExternalDataServiceRequest
	GetRegionId() *string
	SetServiceDescription(v string) *ModifyExternalDataServiceRequest
	GetServiceDescription() *string
	SetServiceId(v string) *ModifyExternalDataServiceRequest
	GetServiceId() *string
	SetServiceSpec(v string) *ModifyExternalDataServiceRequest
	GetServiceSpec() *string
}

type ModifyExternalDataServiceRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) API to view available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service description.
	//
	// example:
	//
	// pxf test
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// Service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
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

func (s ModifyExternalDataServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyExternalDataServiceRequest) GoString() string {
	return s.String()
}

func (s *ModifyExternalDataServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyExternalDataServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyExternalDataServiceRequest) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *ModifyExternalDataServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ModifyExternalDataServiceRequest) GetServiceSpec() *string {
	return s.ServiceSpec
}

func (s *ModifyExternalDataServiceRequest) SetDBInstanceId(v string) *ModifyExternalDataServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyExternalDataServiceRequest) SetRegionId(v string) *ModifyExternalDataServiceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyExternalDataServiceRequest) SetServiceDescription(v string) *ModifyExternalDataServiceRequest {
	s.ServiceDescription = &v
	return s
}

func (s *ModifyExternalDataServiceRequest) SetServiceId(v string) *ModifyExternalDataServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *ModifyExternalDataServiceRequest) SetServiceSpec(v string) *ModifyExternalDataServiceRequest {
	s.ServiceSpec = &v
	return s
}

func (s *ModifyExternalDataServiceRequest) Validate() error {
	return dara.Validate(s)
}
