// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigTrackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListConfigTrackResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListConfigTrackResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListConfigTrackResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *ListConfigTrackResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListConfigTrackResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListConfigTrackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListConfigTrackResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListConfigTrackResponseBody
	GetTotalCount() *int64
	SetTraces(v []*ListConfigTrackResponseBodyTraces) *ListConfigTrackResponseBody
	GetTraces() []*ListConfigTrackResponseBodyTraces
}

type ListConfigTrackResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0CE3ABD2-1E04-561F-A9B4-0423D50****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The track data.
	Traces []*ListConfigTrackResponseBodyTraces `json:"Traces,omitempty" xml:"Traces,omitempty" type:"Repeated"`
}

func (s ListConfigTrackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConfigTrackResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigTrackResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListConfigTrackResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListConfigTrackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConfigTrackResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListConfigTrackResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListConfigTrackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConfigTrackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListConfigTrackResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListConfigTrackResponseBody) GetTraces() []*ListConfigTrackResponseBodyTraces {
	return s.Traces
}

func (s *ListConfigTrackResponseBody) SetErrorCode(v string) *ListConfigTrackResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListConfigTrackResponseBody) SetHttpCode(v string) *ListConfigTrackResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListConfigTrackResponseBody) SetMessage(v string) *ListConfigTrackResponseBody {
	s.Message = &v
	return s
}

func (s *ListConfigTrackResponseBody) SetPageNumber(v int64) *ListConfigTrackResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListConfigTrackResponseBody) SetPageSize(v int64) *ListConfigTrackResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListConfigTrackResponseBody) SetRequestId(v string) *ListConfigTrackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigTrackResponseBody) SetSuccess(v bool) *ListConfigTrackResponseBody {
	s.Success = &v
	return s
}

func (s *ListConfigTrackResponseBody) SetTotalCount(v int64) *ListConfigTrackResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListConfigTrackResponseBody) SetTraces(v []*ListConfigTrackResponseBodyTraces) *ListConfigTrackResponseBody {
	s.Traces = v
	return s
}

func (s *ListConfigTrackResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListConfigTrackResponseBodyTraces struct {
	// Indicates whether the request is sent from the client. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Client *bool `json:"Client,omitempty" xml:"Client,omitempty"`
	// The ID of the configuration.
	//
	// example:
	//
	// eir-server.properties
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The response latency. Unit: milliseconds.
	//
	// example:
	//
	// 0
	Delay *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The event. Valid values:
	//
	// 	- pull: configuration acquisition events
	//
	// 	- persist: persistence events
	//
	// example:
	//
	// pull
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The name of the configuration group.
	//
	// example:
	//
	// DEFAULT_GROUP
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The logging time.
	//
	// example:
	//
	// 2022-11-28 15:09:15
	LogDate *string `json:"LogDate,omitempty" xml:"LogDate,omitempty"`
	// The MD5 value.
	//
	// example:
	//
	// d21c9091c60daa0ff7ee2f420141e5a0
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// Indicates whether messages are pushed by a server. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Push *bool `json:"Push,omitempty" xml:"Push,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.1.2
	RequestIp *string `json:"RequestIp,omitempty" xml:"RequestIp,omitempty"`
	// The response node.
	//
	// example:
	//
	// mse-1973b9a0-1670834*****-reg-center-0-2
	ResponseIp *string `json:"ResponseIp,omitempty" xml:"ResponseIp,omitempty"`
	// The result.
	//
	// example:
	//
	// ok
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The timestamp that indicates the time when the metric value is collected.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 1659666529
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
	// The release type. Valid values:
	//
	// 	- beta: beta release
	//
	// 	- tag: canary release
	//
	// 	- batch: batch release
	//
	// example:
	//
	// beta
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListConfigTrackResponseBodyTraces) String() string {
	return dara.Prettify(s)
}

func (s ListConfigTrackResponseBodyTraces) GoString() string {
	return s.String()
}

func (s *ListConfigTrackResponseBodyTraces) GetClient() *bool {
	return s.Client
}

func (s *ListConfigTrackResponseBodyTraces) GetDataId() *string {
	return s.DataId
}

func (s *ListConfigTrackResponseBodyTraces) GetDelay() *string {
	return s.Delay
}

func (s *ListConfigTrackResponseBodyTraces) GetEvent() *string {
	return s.Event
}

func (s *ListConfigTrackResponseBodyTraces) GetGroup() *string {
	return s.Group
}

func (s *ListConfigTrackResponseBodyTraces) GetLogDate() *string {
	return s.LogDate
}

func (s *ListConfigTrackResponseBodyTraces) GetMd5() *string {
	return s.Md5
}

func (s *ListConfigTrackResponseBodyTraces) GetPush() *bool {
	return s.Push
}

func (s *ListConfigTrackResponseBodyTraces) GetRequestIp() *string {
	return s.RequestIp
}

func (s *ListConfigTrackResponseBodyTraces) GetResponseIp() *string {
	return s.ResponseIp
}

func (s *ListConfigTrackResponseBodyTraces) GetResult() *string {
	return s.Result
}

func (s *ListConfigTrackResponseBodyTraces) GetTs() *string {
	return s.Ts
}

func (s *ListConfigTrackResponseBodyTraces) GetType() *string {
	return s.Type
}

func (s *ListConfigTrackResponseBodyTraces) SetClient(v bool) *ListConfigTrackResponseBodyTraces {
	s.Client = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) SetDataId(v string) *ListConfigTrackResponseBodyTraces {
	s.DataId = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) SetDelay(v string) *ListConfigTrackResponseBodyTraces {
	s.Delay = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) SetEvent(v string) *ListConfigTrackResponseBodyTraces {
	s.Event = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) SetGroup(v string) *ListConfigTrackResponseBodyTraces {
	s.Group = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) SetLogDate(v string) *ListConfigTrackResponseBodyTraces {
	s.LogDate = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) SetMd5(v string) *ListConfigTrackResponseBodyTraces {
	s.Md5 = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) SetPush(v bool) *ListConfigTrackResponseBodyTraces {
	s.Push = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) SetRequestIp(v string) *ListConfigTrackResponseBodyTraces {
	s.RequestIp = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) SetResponseIp(v string) *ListConfigTrackResponseBodyTraces {
	s.ResponseIp = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) SetResult(v string) *ListConfigTrackResponseBodyTraces {
	s.Result = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) SetTs(v string) *ListConfigTrackResponseBodyTraces {
	s.Ts = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) SetType(v string) *ListConfigTrackResponseBodyTraces {
	s.Type = &v
	return s
}

func (s *ListConfigTrackResponseBodyTraces) Validate() error {
	return dara.Validate(s)
}
