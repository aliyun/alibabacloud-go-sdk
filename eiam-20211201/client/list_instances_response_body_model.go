// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody
	GetInstances() []*ListInstancesResponseBodyInstances
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListInstancesResponseBody
	GetTotalCount() *int64
}

type ListInstancesResponseBody struct {
	// The information of instances.
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetInstances() []*ListInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int64) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstances struct {
	// The time when the instance was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1550115455000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default endpoint of the instance.
	DefaultEndpoint *ListInstancesResponseBodyInstancesDefaultEndpoint `json:"DefaultEndpoint,omitempty" xml:"DefaultEndpoint,omitempty" type:"Struct"`
	// The description of the instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_eypq6ljgyeuwmlw672sulxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- creating
	//
	// 	- running
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListInstancesResponseBodyInstances) GetDefaultEndpoint() *ListInstancesResponseBodyInstancesDefaultEndpoint {
	return s.DefaultEndpoint
}

func (s *ListInstancesResponseBodyInstances) GetDescription() *string {
	return s.Description
}

func (s *ListInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstances) SetCreateTime(v int64) *ListInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDefaultEndpoint(v *ListInstancesResponseBodyInstancesDefaultEndpoint) *ListInstancesResponseBodyInstances {
	s.DefaultEndpoint = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDescription(v string) *ListInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesDefaultEndpoint struct {
	// The endpoint of the instance.
	//
	// example:
	//
	// example-xxx.aliyunidaas.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The status of the endpoint. Valid values:
	//
	// 	- resolved
	//
	// 	- unresolved
	//
	// example:
	//
	// resolved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstancesDefaultEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesDefaultEndpoint) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesDefaultEndpoint) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListInstancesResponseBodyInstancesDefaultEndpoint) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstancesDefaultEndpoint) SetEndpoint(v string) *ListInstancesResponseBodyInstancesDefaultEndpoint {
	s.Endpoint = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDefaultEndpoint) SetStatus(v string) *ListInstancesResponseBodyInstancesDefaultEndpoint {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDefaultEndpoint) Validate() error {
	return dara.Validate(s)
}
