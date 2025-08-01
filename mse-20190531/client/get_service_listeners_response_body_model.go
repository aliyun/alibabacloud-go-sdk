// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceListenersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetServiceListenersResponseBodyData) *GetServiceListenersResponseBody
	GetData() []*GetServiceListenersResponseBodyData
	SetErrorCode(v string) *GetServiceListenersResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *GetServiceListenersResponseBody
	GetHttpCode() *string
	SetMessage(v string) *GetServiceListenersResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetServiceListenersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetServiceListenersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetServiceListenersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetServiceListenersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetServiceListenersResponseBody
	GetTotalCount() *int32
}

type GetServiceListenersResponseBody struct {
	// The returned data.
	Data []*GetServiceListenersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// 	- If the request is successful, a success message is returned.
	//
	// 	- If the request fails, an error message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 54973C90-F379-4372-9AA5-053A3F7****
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
	// The number of listeners that are queried.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetServiceListenersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListenersResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceListenersResponseBody) GetData() []*GetServiceListenersResponseBodyData {
	return s.Data
}

func (s *GetServiceListenersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetServiceListenersResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *GetServiceListenersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceListenersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetServiceListenersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetServiceListenersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceListenersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceListenersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetServiceListenersResponseBody) SetData(v []*GetServiceListenersResponseBodyData) *GetServiceListenersResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceListenersResponseBody) SetErrorCode(v string) *GetServiceListenersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetServiceListenersResponseBody) SetHttpCode(v string) *GetServiceListenersResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetServiceListenersResponseBody) SetMessage(v string) *GetServiceListenersResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceListenersResponseBody) SetPageNumber(v int32) *GetServiceListenersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetServiceListenersResponseBody) SetPageSize(v int32) *GetServiceListenersResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetServiceListenersResponseBody) SetRequestId(v string) *GetServiceListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceListenersResponseBody) SetSuccess(v bool) *GetServiceListenersResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceListenersResponseBody) SetTotalCount(v int32) *GetServiceListenersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetServiceListenersResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetServiceListenersResponseBodyData struct {
	// The IP address of the listener.
	//
	// example:
	//
	// 119.23.84.102
	Addr *string `json:"Addr,omitempty" xml:"Addr,omitempty"`
	// The listener client version.
	//
	// example:
	//
	// Nacos-Java-Client:v2.1.1
	Agent *string `json:"Agent,omitempty" xml:"Agent,omitempty"`
	// The application name of the listener.
	//
	// example:
	//
	// app
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The name of the cluster to which the monitored service belongs.
	//
	// example:
	//
	// DEFAULT
	Cluster *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	// The IP address of the monitored service.
	//
	// example:
	//
	// 172.16.1.5
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 0ba53825-b183-414f-a6a0-288e4f1c467e
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The port number of the monitored service.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the monitored service.
	//
	// example:
	//
	// zeekr-orderboss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetServiceListenersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListenersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceListenersResponseBodyData) GetAddr() *string {
	return s.Addr
}

func (s *GetServiceListenersResponseBodyData) GetAgent() *string {
	return s.Agent
}

func (s *GetServiceListenersResponseBodyData) GetApp() *string {
	return s.App
}

func (s *GetServiceListenersResponseBodyData) GetCluster() *string {
	return s.Cluster
}

func (s *GetServiceListenersResponseBodyData) GetIP() *string {
	return s.IP
}

func (s *GetServiceListenersResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetServiceListenersResponseBodyData) GetPort() *string {
	return s.Port
}

func (s *GetServiceListenersResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceListenersResponseBodyData) SetAddr(v string) *GetServiceListenersResponseBodyData {
	s.Addr = &v
	return s
}

func (s *GetServiceListenersResponseBodyData) SetAgent(v string) *GetServiceListenersResponseBodyData {
	s.Agent = &v
	return s
}

func (s *GetServiceListenersResponseBodyData) SetApp(v string) *GetServiceListenersResponseBodyData {
	s.App = &v
	return s
}

func (s *GetServiceListenersResponseBodyData) SetCluster(v string) *GetServiceListenersResponseBodyData {
	s.Cluster = &v
	return s
}

func (s *GetServiceListenersResponseBodyData) SetIP(v string) *GetServiceListenersResponseBodyData {
	s.IP = &v
	return s
}

func (s *GetServiceListenersResponseBodyData) SetNamespaceId(v string) *GetServiceListenersResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *GetServiceListenersResponseBodyData) SetPort(v string) *GetServiceListenersResponseBodyData {
	s.Port = &v
	return s
}

func (s *GetServiceListenersResponseBodyData) SetServiceName(v string) *GetServiceListenersResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetServiceListenersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
