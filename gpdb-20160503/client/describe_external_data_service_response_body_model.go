// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExternalDataServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeExternalDataServiceResponseBody
	GetCreateTime() *string
	SetModifyTime(v string) *DescribeExternalDataServiceResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *DescribeExternalDataServiceResponseBody
	GetRequestId() *string
	SetServiceDescription(v string) *DescribeExternalDataServiceResponseBody
	GetServiceDescription() *string
	SetServiceId(v string) *DescribeExternalDataServiceResponseBody
	GetServiceId() *string
	SetServiceName(v string) *DescribeExternalDataServiceResponseBody
	GetServiceName() *string
	SetServiceSpec(v string) *DescribeExternalDataServiceResponseBody
	GetServiceSpec() *string
	SetStatus(v string) *DescribeExternalDataServiceResponseBody
	GetStatus() *string
}

type DescribeExternalDataServiceResponseBody struct {
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
	// The name of the service.
	//
	// example:
	//
	// test-adbpgss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The specifications of the service. Unit: compute units (CUs).
	//
	// example:
	//
	// 2
	ServiceSpec *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
	// The status of the operation.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExternalDataServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExternalDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExternalDataServiceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeExternalDataServiceResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeExternalDataServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExternalDataServiceResponseBody) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *DescribeExternalDataServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeExternalDataServiceResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeExternalDataServiceResponseBody) GetServiceSpec() *string {
	return s.ServiceSpec
}

func (s *DescribeExternalDataServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeExternalDataServiceResponseBody) SetCreateTime(v string) *DescribeExternalDataServiceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeExternalDataServiceResponseBody) SetModifyTime(v string) *DescribeExternalDataServiceResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeExternalDataServiceResponseBody) SetRequestId(v string) *DescribeExternalDataServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExternalDataServiceResponseBody) SetServiceDescription(v string) *DescribeExternalDataServiceResponseBody {
	s.ServiceDescription = &v
	return s
}

func (s *DescribeExternalDataServiceResponseBody) SetServiceId(v string) *DescribeExternalDataServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DescribeExternalDataServiceResponseBody) SetServiceName(v string) *DescribeExternalDataServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeExternalDataServiceResponseBody) SetServiceSpec(v string) *DescribeExternalDataServiceResponseBody {
	s.ServiceSpec = &v
	return s
}

func (s *DescribeExternalDataServiceResponseBody) SetStatus(v string) *DescribeExternalDataServiceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeExternalDataServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
