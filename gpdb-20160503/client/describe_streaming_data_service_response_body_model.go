// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamingDataServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeStreamingDataServiceResponseBody
	GetCreateTime() *string
	SetModifyTime(v string) *DescribeStreamingDataServiceResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *DescribeStreamingDataServiceResponseBody
	GetRequestId() *string
	SetServiceDescription(v string) *DescribeStreamingDataServiceResponseBody
	GetServiceDescription() *string
	SetServiceId(v string) *DescribeStreamingDataServiceResponseBody
	GetServiceId() *string
	SetServiceIp(v string) *DescribeStreamingDataServiceResponseBody
	GetServiceIp() *string
	SetServiceManaged(v bool) *DescribeStreamingDataServiceResponseBody
	GetServiceManaged() *bool
	SetServiceName(v string) *DescribeStreamingDataServiceResponseBody
	GetServiceName() *string
	SetServiceOwnerId(v string) *DescribeStreamingDataServiceResponseBody
	GetServiceOwnerId() *string
	SetServicePort(v int32) *DescribeStreamingDataServiceResponseBody
	GetServicePort() *int32
	SetServiceSpec(v string) *DescribeStreamingDataServiceResponseBody
	GetServiceSpec() *string
	SetStatus(v string) *DescribeStreamingDataServiceResponseBody
	GetStatus() *string
}

type DescribeStreamingDataServiceResponseBody struct {
	// The time when the service was created.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the service was last modified.
	//
	// example:
	//
	// 2019-09-08T17:00:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// test-adbpgss
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// The service ID.
	//
	// example:
	//
	// 1
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The IP address of the service.
	//
	// example:
	//
	// 192.168.0.1
	ServiceIp *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	// The service is managed by other aliyun product or not.
	//
	// example:
	//
	// False
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// test-adbpgss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service account uid of the aliyun product
	//
	// example:
	//
	// 123456
	ServiceOwnerId *string `json:"ServiceOwnerId,omitempty" xml:"ServiceOwnerId,omitempty"`
	// The port number of the service.
	//
	// example:
	//
	// 5432
	ServicePort *int32 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The specifications of the service.
	//
	// example:
	//
	// 2
	ServiceSpec *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
	// The status of the service. Valid values:
	//
	// 	- Init
	//
	// 	- Running
	//
	// 	- Exception
	//
	// 	- Paused
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeStreamingDataServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamingDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamingDataServiceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeStreamingDataServiceResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeStreamingDataServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStreamingDataServiceResponseBody) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *DescribeStreamingDataServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeStreamingDataServiceResponseBody) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *DescribeStreamingDataServiceResponseBody) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeStreamingDataServiceResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeStreamingDataServiceResponseBody) GetServiceOwnerId() *string {
	return s.ServiceOwnerId
}

func (s *DescribeStreamingDataServiceResponseBody) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *DescribeStreamingDataServiceResponseBody) GetServiceSpec() *string {
	return s.ServiceSpec
}

func (s *DescribeStreamingDataServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeStreamingDataServiceResponseBody) SetCreateTime(v string) *DescribeStreamingDataServiceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeStreamingDataServiceResponseBody) SetModifyTime(v string) *DescribeStreamingDataServiceResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeStreamingDataServiceResponseBody) SetRequestId(v string) *DescribeStreamingDataServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamingDataServiceResponseBody) SetServiceDescription(v string) *DescribeStreamingDataServiceResponseBody {
	s.ServiceDescription = &v
	return s
}

func (s *DescribeStreamingDataServiceResponseBody) SetServiceId(v string) *DescribeStreamingDataServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DescribeStreamingDataServiceResponseBody) SetServiceIp(v string) *DescribeStreamingDataServiceResponseBody {
	s.ServiceIp = &v
	return s
}

func (s *DescribeStreamingDataServiceResponseBody) SetServiceManaged(v bool) *DescribeStreamingDataServiceResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeStreamingDataServiceResponseBody) SetServiceName(v string) *DescribeStreamingDataServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeStreamingDataServiceResponseBody) SetServiceOwnerId(v string) *DescribeStreamingDataServiceResponseBody {
	s.ServiceOwnerId = &v
	return s
}

func (s *DescribeStreamingDataServiceResponseBody) SetServicePort(v int32) *DescribeStreamingDataServiceResponseBody {
	s.ServicePort = &v
	return s
}

func (s *DescribeStreamingDataServiceResponseBody) SetServiceSpec(v string) *DescribeStreamingDataServiceResponseBody {
	s.ServiceSpec = &v
	return s
}

func (s *DescribeStreamingDataServiceResponseBody) SetStatus(v string) *DescribeStreamingDataServiceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeStreamingDataServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
