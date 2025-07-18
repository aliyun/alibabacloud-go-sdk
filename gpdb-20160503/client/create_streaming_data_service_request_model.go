// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamingDataServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateStreamingDataServiceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *CreateStreamingDataServiceRequest
	GetRegionId() *string
	SetServiceDescription(v string) *CreateStreamingDataServiceRequest
	GetServiceDescription() *string
	SetServiceName(v string) *CreateStreamingDataServiceRequest
	GetServiceName() *string
	SetServiceSpec(v string) *CreateStreamingDataServiceRequest
	GetServiceSpec() *string
}

type CreateStreamingDataServiceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent list of regions.
	//
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
	// The name of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-adbpgss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The specifications of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ServiceSpec *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
}

func (s CreateStreamingDataServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamingDataServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamingDataServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateStreamingDataServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStreamingDataServiceRequest) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *CreateStreamingDataServiceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateStreamingDataServiceRequest) GetServiceSpec() *string {
	return s.ServiceSpec
}

func (s *CreateStreamingDataServiceRequest) SetDBInstanceId(v string) *CreateStreamingDataServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateStreamingDataServiceRequest) SetRegionId(v string) *CreateStreamingDataServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStreamingDataServiceRequest) SetServiceDescription(v string) *CreateStreamingDataServiceRequest {
	s.ServiceDescription = &v
	return s
}

func (s *CreateStreamingDataServiceRequest) SetServiceName(v string) *CreateStreamingDataServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateStreamingDataServiceRequest) SetServiceSpec(v string) *CreateStreamingDataServiceRequest {
	s.ServiceSpec = &v
	return s
}

func (s *CreateStreamingDataServiceRequest) Validate() error {
	return dara.Validate(s)
}
