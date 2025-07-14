// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScaleAppMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetScaleAppMetricResponseBody
	GetCode() *string
	SetData(v []*GetScaleAppMetricResponseBodyData) *GetScaleAppMetricResponseBody
	GetData() []*GetScaleAppMetricResponseBodyData
	SetMessage(v string) *GetScaleAppMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetScaleAppMetricResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetScaleAppMetricResponseBody
	GetSuccess() *bool
}

type GetScaleAppMetricResponseBody struct {
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
	// The details of applications.
	Data []*GetScaleAppMetricResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 3B763F98-0BA2-5C23-B6B8-558568D2C1C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the microservice list was obtained. The following limits are imposed on the ID:
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

func (s GetScaleAppMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScaleAppMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetScaleAppMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetScaleAppMetricResponseBody) GetData() []*GetScaleAppMetricResponseBodyData {
	return s.Data
}

func (s *GetScaleAppMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScaleAppMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScaleAppMetricResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetScaleAppMetricResponseBody) SetCode(v string) *GetScaleAppMetricResponseBody {
	s.Code = &v
	return s
}

func (s *GetScaleAppMetricResponseBody) SetData(v []*GetScaleAppMetricResponseBodyData) *GetScaleAppMetricResponseBody {
	s.Data = v
	return s
}

func (s *GetScaleAppMetricResponseBody) SetMessage(v string) *GetScaleAppMetricResponseBody {
	s.Message = &v
	return s
}

func (s *GetScaleAppMetricResponseBody) SetRequestId(v string) *GetScaleAppMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScaleAppMetricResponseBody) SetSuccess(v bool) *GetScaleAppMetricResponseBody {
	s.Success = &v
	return s
}

func (s *GetScaleAppMetricResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetScaleAppMetricResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The maximum number of instances.
	//
	// example:
	//
	// 10
	MaxReplicas *int64 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
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
	// 10
	Runnings *int64 `json:"Runnings,omitempty" xml:"Runnings,omitempty"`
}

func (s GetScaleAppMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScaleAppMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScaleAppMetricResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetScaleAppMetricResponseBodyData) GetMaxReplicas() *int64 {
	return s.MaxReplicas
}

func (s *GetScaleAppMetricResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetScaleAppMetricResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetScaleAppMetricResponseBodyData) GetRunnings() *int64 {
	return s.Runnings
}

func (s *GetScaleAppMetricResponseBodyData) SetAppId(v string) *GetScaleAppMetricResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetScaleAppMetricResponseBodyData) SetMaxReplicas(v int64) *GetScaleAppMetricResponseBodyData {
	s.MaxReplicas = &v
	return s
}

func (s *GetScaleAppMetricResponseBodyData) SetName(v string) *GetScaleAppMetricResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetScaleAppMetricResponseBodyData) SetRegionId(v string) *GetScaleAppMetricResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetScaleAppMetricResponseBodyData) SetRunnings(v int64) *GetScaleAppMetricResponseBodyData {
	s.Runnings = &v
	return s
}

func (s *GetScaleAppMetricResponseBodyData) Validate() error {
	return dara.Validate(s)
}
