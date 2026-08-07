// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstancesResponseBody
	GetCode() *string
	SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody
	GetData() *ListInstancesResponseBodyData
	SetHttpStatusCode(v int32) *ListInstancesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstancesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListInstancesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstancesResponseBody
	GetSuccess() *bool
}

type ListInstancesResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The paged query result.
	Data *ListInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance does not exist. Instance=placeholder-instance-id.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstancesResponseBody) GetData() *ListInstancesResponseBodyData {
	return s.Data
}

func (s *ListInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetParams(v []*string) *ListInstancesResponseBody {
	s.Params = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesResponseBodyData struct {
	// The list of instances.
	Instances []*ListInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) GetInstances() []*ListInstancesResponseBodyDataInstances {
	return s.Instances
}

func (s *ListInstancesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesResponseBodyData) SetInstances(v []*ListInstancesResponseBodyDataInstances) *ListInstancesResponseBodyData {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageNumber(v int32) *ListInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageSize(v int32) *ListInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetTotalCount(v int32) *ListInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBodyData) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyDataInstances struct {
	// The number of concurrent connections.
	//
	// example:
	//
	// 10
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1769653616000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The instance description.
	//
	// example:
	//
	// Ask about customer satisfaction
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// fc21ccd7-6f0a-4261-ac63-d079bfe3bc2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// Satisfaction survey
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The service mode.
	//
	// example:
	//
	// STANDARD
	ServiceMode *string `json:"ServiceMode,omitempty" xml:"ServiceMode,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1308144684576765
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1769653616000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListInstancesResponseBodyDataInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataInstances) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *ListInstancesResponseBodyDataInstances) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListInstancesResponseBodyDataInstances) GetDescription() *string {
	return s.Description
}

func (s *ListInstancesResponseBodyDataInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyDataInstances) GetName() *string {
	return s.Name
}

func (s *ListInstancesResponseBodyDataInstances) GetServiceMode() *string {
	return s.ServiceMode
}

func (s *ListInstancesResponseBodyDataInstances) GetTenantId() *string {
	return s.TenantId
}

func (s *ListInstancesResponseBodyDataInstances) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListInstancesResponseBodyDataInstances) SetConcurrency(v int32) *ListInstancesResponseBodyDataInstances {
	s.Concurrency = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetCreatedTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.CreatedTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetDescription(v string) *ListInstancesResponseBodyDataInstances {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceId(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetName(v string) *ListInstancesResponseBodyDataInstances {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetServiceMode(v string) *ListInstancesResponseBodyDataInstances {
	s.ServiceMode = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetTenantId(v string) *ListInstancesResponseBodyDataInstances {
	s.TenantId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetUpdatedTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.UpdatedTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) Validate() error {
	return dara.Validate(s)
}
