// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWarningEventMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWarningEventMetricResponseBody
	GetCode() *string
	SetData(v []*GetWarningEventMetricResponseBodyData) *GetWarningEventMetricResponseBody
	GetData() []*GetWarningEventMetricResponseBodyData
	SetMessage(v string) *GetWarningEventMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWarningEventMetricResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWarningEventMetricResponseBody
	GetSuccess() *bool
}

type GetWarningEventMetricResponseBody struct {
	// The number of Warning events.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 3B763F98-0BA2-5C23-B6B8-558568D2C1C2
	Data []*GetWarningEventMetricResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 3B763F98-0BA2-5C23-B6B8-558568D2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWarningEventMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWarningEventMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetWarningEventMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWarningEventMetricResponseBody) GetData() []*GetWarningEventMetricResponseBodyData {
	return s.Data
}

func (s *GetWarningEventMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWarningEventMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWarningEventMetricResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWarningEventMetricResponseBody) SetCode(v string) *GetWarningEventMetricResponseBody {
	s.Code = &v
	return s
}

func (s *GetWarningEventMetricResponseBody) SetData(v []*GetWarningEventMetricResponseBodyData) *GetWarningEventMetricResponseBody {
	s.Data = v
	return s
}

func (s *GetWarningEventMetricResponseBody) SetMessage(v string) *GetWarningEventMetricResponseBody {
	s.Message = &v
	return s
}

func (s *GetWarningEventMetricResponseBody) SetRequestId(v string) *GetWarningEventMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWarningEventMetricResponseBody) SetSuccess(v bool) *GetWarningEventMetricResponseBody {
	s.Success = &v
	return s
}

func (s *GetWarningEventMetricResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWarningEventMetricResponseBodyData struct {
	// The details of the application.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application ID.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The application name.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// 10
	WarningCount *int64 `json:"WarningCount,omitempty" xml:"WarningCount,omitempty"`
}

func (s GetWarningEventMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWarningEventMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWarningEventMetricResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetWarningEventMetricResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetWarningEventMetricResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWarningEventMetricResponseBodyData) GetWarningCount() *int64 {
	return s.WarningCount
}

func (s *GetWarningEventMetricResponseBodyData) SetAppId(v string) *GetWarningEventMetricResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetWarningEventMetricResponseBodyData) SetName(v string) *GetWarningEventMetricResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetWarningEventMetricResponseBodyData) SetRegionId(v string) *GetWarningEventMetricResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetWarningEventMetricResponseBodyData) SetWarningCount(v int64) *GetWarningEventMetricResponseBodyData {
	s.WarningCount = &v
	return s
}

func (s *GetWarningEventMetricResponseBodyData) Validate() error {
	return dara.Validate(s)
}
