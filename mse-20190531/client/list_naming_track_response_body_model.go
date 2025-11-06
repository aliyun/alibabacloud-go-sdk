// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamingTrackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListNamingTrackResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListNamingTrackResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListNamingTrackResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *ListNamingTrackResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListNamingTrackResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListNamingTrackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNamingTrackResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListNamingTrackResponseBody
	GetTotalCount() *int64
	SetTraces(v []*ListNamingTrackResponseBodyTraces) *ListNamingTrackResponseBody
	GetTraces() []*ListNamingTrackResponseBodyTraces
}

type ListNamingTrackResponseBody struct {
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
	// The number of the returned page.
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
	// 9916CBED-B2D5-5685-9129-4592FE1*****
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
	// The total number of returned entries.
	//
	// example:
	//
	// 7
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The data information.
	Traces []*ListNamingTrackResponseBodyTraces `json:"Traces,omitempty" xml:"Traces,omitempty" type:"Repeated"`
}

func (s ListNamingTrackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNamingTrackResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamingTrackResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListNamingTrackResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListNamingTrackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNamingTrackResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListNamingTrackResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListNamingTrackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNamingTrackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNamingTrackResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListNamingTrackResponseBody) GetTraces() []*ListNamingTrackResponseBodyTraces {
	return s.Traces
}

func (s *ListNamingTrackResponseBody) SetErrorCode(v string) *ListNamingTrackResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNamingTrackResponseBody) SetHttpCode(v string) *ListNamingTrackResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListNamingTrackResponseBody) SetMessage(v string) *ListNamingTrackResponseBody {
	s.Message = &v
	return s
}

func (s *ListNamingTrackResponseBody) SetPageNumber(v int64) *ListNamingTrackResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNamingTrackResponseBody) SetPageSize(v int64) *ListNamingTrackResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNamingTrackResponseBody) SetRequestId(v string) *ListNamingTrackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamingTrackResponseBody) SetSuccess(v bool) *ListNamingTrackResponseBody {
	s.Success = &v
	return s
}

func (s *ListNamingTrackResponseBody) SetTotalCount(v int64) *ListNamingTrackResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNamingTrackResponseBody) SetTraces(v []*ListNamingTrackResponseBodyTraces) *ListNamingTrackResponseBody {
	s.Traces = v
	return s
}

func (s *ListNamingTrackResponseBody) Validate() error {
	if s.Traces != nil {
		for _, item := range s.Traces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNamingTrackResponseBodyTraces struct {
	// The IP address of the client.
	//
	// example:
	//
	// 120.40.32.235
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The group.
	//
	// example:
	//
	// prod
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 1
	InstanceSize *string `json:"InstanceSize,omitempty" xml:"InstanceSize,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// mse-197*****-167083******-reg-center-0-0
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The push time.
	//
	// example:
	//
	// 2022-12-16 11:48:07
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// The total push time.
	//
	// example:
	//
	// 628ms
	PushTimeAll *string `json:"PushTimeAll,omitempty" xml:"PushTimeAll,omitempty"`
	// The push time for the network.
	//
	// example:
	//
	// 37ms
	PushTimeNetwork *string `json:"PushTimeNetwork,omitempty" xml:"PushTimeNetwork,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// fpx-xms-baseinfo
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
	// The duration that is specified in the service-level agreement (SLA).
	//
	// example:
	//
	// 628ms
	SlaTime *string `json:"SlaTime,omitempty" xml:"SlaTime,omitempty"`
}

func (s ListNamingTrackResponseBodyTraces) String() string {
	return dara.Prettify(s)
}

func (s ListNamingTrackResponseBodyTraces) GoString() string {
	return s.String()
}

func (s *ListNamingTrackResponseBodyTraces) GetClientIp() *string {
	return s.ClientIp
}

func (s *ListNamingTrackResponseBodyTraces) GetGroup() *string {
	return s.Group
}

func (s *ListNamingTrackResponseBodyTraces) GetInstanceSize() *string {
	return s.InstanceSize
}

func (s *ListNamingTrackResponseBodyTraces) GetNodeName() *string {
	return s.NodeName
}

func (s *ListNamingTrackResponseBodyTraces) GetPushTime() *string {
	return s.PushTime
}

func (s *ListNamingTrackResponseBodyTraces) GetPushTimeAll() *string {
	return s.PushTimeAll
}

func (s *ListNamingTrackResponseBodyTraces) GetPushTimeNetwork() *string {
	return s.PushTimeNetwork
}

func (s *ListNamingTrackResponseBodyTraces) GetServerName() *string {
	return s.ServerName
}

func (s *ListNamingTrackResponseBodyTraces) GetSlaTime() *string {
	return s.SlaTime
}

func (s *ListNamingTrackResponseBodyTraces) SetClientIp(v string) *ListNamingTrackResponseBodyTraces {
	s.ClientIp = &v
	return s
}

func (s *ListNamingTrackResponseBodyTraces) SetGroup(v string) *ListNamingTrackResponseBodyTraces {
	s.Group = &v
	return s
}

func (s *ListNamingTrackResponseBodyTraces) SetInstanceSize(v string) *ListNamingTrackResponseBodyTraces {
	s.InstanceSize = &v
	return s
}

func (s *ListNamingTrackResponseBodyTraces) SetNodeName(v string) *ListNamingTrackResponseBodyTraces {
	s.NodeName = &v
	return s
}

func (s *ListNamingTrackResponseBodyTraces) SetPushTime(v string) *ListNamingTrackResponseBodyTraces {
	s.PushTime = &v
	return s
}

func (s *ListNamingTrackResponseBodyTraces) SetPushTimeAll(v string) *ListNamingTrackResponseBodyTraces {
	s.PushTimeAll = &v
	return s
}

func (s *ListNamingTrackResponseBodyTraces) SetPushTimeNetwork(v string) *ListNamingTrackResponseBodyTraces {
	s.PushTimeNetwork = &v
	return s
}

func (s *ListNamingTrackResponseBodyTraces) SetServerName(v string) *ListNamingTrackResponseBodyTraces {
	s.ServerName = &v
	return s
}

func (s *ListNamingTrackResponseBodyTraces) SetSlaTime(v string) *ListNamingTrackResponseBodyTraces {
	s.SlaTime = &v
	return s
}

func (s *ListNamingTrackResponseBodyTraces) Validate() error {
	return dara.Validate(s)
}
