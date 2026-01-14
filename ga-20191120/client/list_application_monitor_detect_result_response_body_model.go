// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationMonitorDetectResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationMonitorDetectResultList(v []*ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) *ListApplicationMonitorDetectResultResponseBody
	GetApplicationMonitorDetectResultList() []*ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList
	SetPageNumber(v int32) *ListApplicationMonitorDetectResultResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApplicationMonitorDetectResultResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListApplicationMonitorDetectResultResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApplicationMonitorDetectResultResponseBody
	GetTotalCount() *int32
}

type ListApplicationMonitorDetectResultResponseBody struct {
	// Details about the diagnostic result of the origin probing task.
	ApplicationMonitorDetectResultList []*ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList `json:"ApplicationMonitorDetectResultList,omitempty" xml:"ApplicationMonitorDetectResultList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationMonitorDetectResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMonitorDetectResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorDetectResultResponseBody) GetApplicationMonitorDetectResultList() []*ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	return s.ApplicationMonitorDetectResultList
}

func (s *ListApplicationMonitorDetectResultResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationMonitorDetectResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationMonitorDetectResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationMonitorDetectResultResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApplicationMonitorDetectResultResponseBody) SetApplicationMonitorDetectResultList(v []*ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) *ListApplicationMonitorDetectResultResponseBody {
	s.ApplicationMonitorDetectResultList = v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBody) SetPageNumber(v int32) *ListApplicationMonitorDetectResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBody) SetPageSize(v int32) *ListApplicationMonitorDetectResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBody) SetRequestId(v string) *ListApplicationMonitorDetectResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBody) SetTotalCount(v int32) *ListApplicationMonitorDetectResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBody) Validate() error {
	if s.ApplicationMonitorDetectResultList != nil {
		for _, item := range s.ApplicationMonitorDetectResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList struct {
	// The ID of the GA instance on which the origin probing task runs.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The response content returned by the origin probing task.
	//
	// example:
	//
	// 502 BadGateway
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the diagnostic result. Valid values:
	//
	// 	- **All forward nodes work well.:*	- The origin server is normal.
	//
	// 	- **Endpoint network error.:*	- The origin server is abnormal. You must check whether the origin server is running as expected.
	//
	// 	- **Public network error.:*	- An Internet error occurred, which is a network error that occurred when the client connected to the acceleration region.
	//
	// 	- **Ga internal error.:*	- An internal error occurred. For example, an exception occurred when GA processed a request.
	//
	// 	- **Ga has been deleted.:*	- The current GA instance is deleted.
	//
	// 	- **Ga state is not stable.:*	- The current GA instance is in an unstable state, such as the Configuring state.
	//
	// 	- **Ga has no listener configuration.:*	- No listener is configured for the current GA instance.
	//
	// 	- **Missing endpoint configuration.:*	- No endpoint is configured.
	//
	// 	- **Missing acceleration region configuration.:*	- No acceleration region is configured.
	//
	// 	- **Missing endpointgroup configuration.:*	- No endpoint group is configured.
	//
	// example:
	//
	// All forward nodes work well
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The time when the diagnosis of the origin probing task ends.
	//
	// example:
	//
	// 1663205460
	DetectTime *string `json:"DetectTime,omitempty" xml:"DetectTime,omitempty"`
	// The diagnostic result of the origin probing task. Valid values:
	//
	// 	- **success:*	- The origin probing task succeeded.
	//
	// 	- **failed:*	- The origin probing task failed.
	//
	// example:
	//
	// success
	DiagStatus *string `json:"DiagStatus,omitempty" xml:"DiagStatus,omitempty"`
	// The ID of the listener on which the origin probing task runs.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listener port.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The network transmission protocol that is used by the listener. Valid values:
	//
	// 	- **tcp:*	- TCP.
	//
	// 	- **udp:*	- UDP.
	//
	// 	- **http:*	- HTTP.
	//
	// 	- **https:*	- HTTPS.
	//
	// >  UDP listeners do not support probing.
	//
	// example:
	//
	// http
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The error code returned by the origin probing task.
	//
	// example:
	//
	// 502
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The origin probing task ID.
	//
	// example:
	//
	// sm-bp1fpdjfju9k8yr1y****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) GetContent() *string {
	return s.Content
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) GetDetail() *string {
	return s.Detail
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) GetDetectTime() *string {
	return s.DetectTime
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) GetDiagStatus() *string {
	return s.DiagStatus
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) GetPort() *string {
	return s.Port
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) GetProtocol() *string {
	return s.Protocol
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) GetTaskId() *string {
	return s.TaskId
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetAcceleratorId(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.AcceleratorId = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetContent(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.Content = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetDetail(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.Detail = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetDetectTime(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.DetectTime = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetDiagStatus(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.DiagStatus = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetListenerId(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.ListenerId = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetPort(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.Port = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetProtocol(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.Protocol = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetStatusCode(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetTaskId(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.TaskId = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) Validate() error {
	return dara.Validate(s)
}
