// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArmsTopNMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetArmsTopNMetricResponseBody
	GetCode() *string
	SetData(v []*GetArmsTopNMetricResponseBodyData) *GetArmsTopNMetricResponseBody
	GetData() []*GetArmsTopNMetricResponseBodyData
	SetMessage(v string) *GetArmsTopNMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetArmsTopNMetricResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetArmsTopNMetricResponseBody
	GetSuccess() *bool
}

type GetArmsTopNMetricResponseBody struct {
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
	Data []*GetArmsTopNMetricResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 3B763F98-0BA2-5C23-B6B8-558568D2C1C2
	//
	// example:
	//
	// 3B763F98-0BA2-5C23-B6B8-558568D2****
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

func (s GetArmsTopNMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArmsTopNMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetArmsTopNMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetArmsTopNMetricResponseBody) GetData() []*GetArmsTopNMetricResponseBodyData {
	return s.Data
}

func (s *GetArmsTopNMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetArmsTopNMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArmsTopNMetricResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetArmsTopNMetricResponseBody) SetCode(v string) *GetArmsTopNMetricResponseBody {
	s.Code = &v
	return s
}

func (s *GetArmsTopNMetricResponseBody) SetData(v []*GetArmsTopNMetricResponseBodyData) *GetArmsTopNMetricResponseBody {
	s.Data = v
	return s
}

func (s *GetArmsTopNMetricResponseBody) SetMessage(v string) *GetArmsTopNMetricResponseBody {
	s.Message = &v
	return s
}

func (s *GetArmsTopNMetricResponseBody) SetRequestId(v string) *GetArmsTopNMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArmsTopNMetricResponseBody) SetSuccess(v bool) *GetArmsTopNMetricResponseBody {
	s.Success = &v
	return s
}

func (s *GetArmsTopNMetricResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetArmsTopNMetricResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 100
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The number of errors.
	//
	// example:
	//
	// 0
	Error *int64 `json:"Error,omitempty" xml:"Error,omitempty"`
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
	// The average response time. Unit: milliseconds.
	//
	// example:
	//
	// 100
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
}

func (s GetArmsTopNMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetArmsTopNMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetArmsTopNMetricResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetArmsTopNMetricResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *GetArmsTopNMetricResponseBodyData) GetError() *int64 {
	return s.Error
}

func (s *GetArmsTopNMetricResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetArmsTopNMetricResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetArmsTopNMetricResponseBodyData) GetRt() *int64 {
	return s.Rt
}

func (s *GetArmsTopNMetricResponseBodyData) SetAppId(v string) *GetArmsTopNMetricResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetArmsTopNMetricResponseBodyData) SetCount(v int64) *GetArmsTopNMetricResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetArmsTopNMetricResponseBodyData) SetError(v int64) *GetArmsTopNMetricResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetArmsTopNMetricResponseBodyData) SetName(v string) *GetArmsTopNMetricResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetArmsTopNMetricResponseBodyData) SetRegionId(v string) *GetArmsTopNMetricResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetArmsTopNMetricResponseBodyData) SetRt(v int64) *GetArmsTopNMetricResponseBodyData {
	s.Rt = &v
	return s
}

func (s *GetArmsTopNMetricResponseBodyData) Validate() error {
	return dara.Validate(s)
}
