// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvailabilityMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAvailabilityMetricResponseBody
	GetCode() *string
	SetData(v []*GetAvailabilityMetricResponseBodyData) *GetAvailabilityMetricResponseBody
	GetData() []*GetAvailabilityMetricResponseBodyData
	SetMessage(v string) *GetAvailabilityMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAvailabilityMetricResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAvailabilityMetricResponseBody
	GetSuccess() *bool
}

type GetAvailabilityMetricResponseBody struct {
	// The HTTP status code. The following limits are imposed on the ID:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data entries returned.
	Data []*GetAvailabilityMetricResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The additional information that is returned. The following limits are imposed on the ID:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3A92C4EA-4C53-5A1C-8AEB-F2DB11982D5F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the list of applications was obtained. The following limits are imposed on the ID:
	//
	// 	- **true**: The namespaces were obtained.
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAvailabilityMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAvailabilityMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetAvailabilityMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAvailabilityMetricResponseBody) GetData() []*GetAvailabilityMetricResponseBodyData {
	return s.Data
}

func (s *GetAvailabilityMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAvailabilityMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAvailabilityMetricResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAvailabilityMetricResponseBody) SetCode(v string) *GetAvailabilityMetricResponseBody {
	s.Code = &v
	return s
}

func (s *GetAvailabilityMetricResponseBody) SetData(v []*GetAvailabilityMetricResponseBodyData) *GetAvailabilityMetricResponseBody {
	s.Data = v
	return s
}

func (s *GetAvailabilityMetricResponseBody) SetMessage(v string) *GetAvailabilityMetricResponseBody {
	s.Message = &v
	return s
}

func (s *GetAvailabilityMetricResponseBody) SetRequestId(v string) *GetAvailabilityMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAvailabilityMetricResponseBody) SetSuccess(v bool) *GetAvailabilityMetricResponseBody {
	s.Success = &v
	return s
}

func (s *GetAvailabilityMetricResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAvailabilityMetricResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Indicates whether an auto scaling policy is enabled. Valid values:
	//
	// 	- **1**: An auto scaling policy is enabled.
	//
	// 	- **0**: No auto scaling policy is enabled.
	//
	// example:
	//
	// 0
	EnableAutoscale *int64 `json:"EnableAutoscale,omitempty" xml:"EnableAutoscale,omitempty"`
	// The number of abnormal instances.
	//
	// example:
	//
	// 0
	ErrorInstances *int64 `json:"ErrorInstances,omitempty" xml:"ErrorInstances,omitempty"`
	// The expected number of instances.
	//
	// example:
	//
	// 0
	Instances *int64 `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// The application name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The current number of instances.
	//
	// example:
	//
	// 1
	Runnings *int64 `json:"Runnings,omitempty" xml:"Runnings,omitempty"`
}

func (s GetAvailabilityMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAvailabilityMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAvailabilityMetricResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetAvailabilityMetricResponseBodyData) GetEnableAutoscale() *int64 {
	return s.EnableAutoscale
}

func (s *GetAvailabilityMetricResponseBodyData) GetErrorInstances() *int64 {
	return s.ErrorInstances
}

func (s *GetAvailabilityMetricResponseBodyData) GetInstances() *int64 {
	return s.Instances
}

func (s *GetAvailabilityMetricResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetAvailabilityMetricResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAvailabilityMetricResponseBodyData) GetRunnings() *int64 {
	return s.Runnings
}

func (s *GetAvailabilityMetricResponseBodyData) SetAppId(v string) *GetAvailabilityMetricResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetAvailabilityMetricResponseBodyData) SetEnableAutoscale(v int64) *GetAvailabilityMetricResponseBodyData {
	s.EnableAutoscale = &v
	return s
}

func (s *GetAvailabilityMetricResponseBodyData) SetErrorInstances(v int64) *GetAvailabilityMetricResponseBodyData {
	s.ErrorInstances = &v
	return s
}

func (s *GetAvailabilityMetricResponseBodyData) SetInstances(v int64) *GetAvailabilityMetricResponseBodyData {
	s.Instances = &v
	return s
}

func (s *GetAvailabilityMetricResponseBodyData) SetName(v string) *GetAvailabilityMetricResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAvailabilityMetricResponseBodyData) SetRegionId(v string) *GetAvailabilityMetricResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetAvailabilityMetricResponseBodyData) SetRunnings(v int64) *GetAvailabilityMetricResponseBodyData {
	s.Runnings = &v
	return s
}

func (s *GetAvailabilityMetricResponseBodyData) Validate() error {
	return dara.Validate(s)
}
