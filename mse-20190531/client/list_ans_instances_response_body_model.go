// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnsInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAnsInstancesResponseBodyData) *ListAnsInstancesResponseBody
	GetData() []*ListAnsInstancesResponseBodyData
	SetErrorCode(v string) *ListAnsInstancesResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListAnsInstancesResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListAnsInstancesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListAnsInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAnsInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAnsInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAnsInstancesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListAnsInstancesResponseBody
	GetTotalCount() *int32
}

type ListAnsInstancesResponseBody struct {
	// The details of the data.
	Data []*ListAnsInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of the page to return.
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
	// 52BA6DA6-A702-4362-A32F-DFF79655****
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
	// The total number of returned instances.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAnsInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAnsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnsInstancesResponseBody) GetData() []*ListAnsInstancesResponseBodyData {
	return s.Data
}

func (s *ListAnsInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAnsInstancesResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListAnsInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAnsInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAnsInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAnsInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAnsInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAnsInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAnsInstancesResponseBody) SetData(v []*ListAnsInstancesResponseBodyData) *ListAnsInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListAnsInstancesResponseBody) SetErrorCode(v string) *ListAnsInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetHttpCode(v string) *ListAnsInstancesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetMessage(v string) *ListAnsInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetPageNumber(v int32) *ListAnsInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetPageSize(v int32) *ListAnsInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetRequestId(v string) *ListAnsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetSuccess(v bool) *ListAnsInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetTotalCount(v int32) *ListAnsInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAnsInstancesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAnsInstancesResponseBodyData struct {
	// The name of the application.
	//
	// example:
	//
	// app
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// DEFAULT
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The reference key.
	//
	// example:
	//
	// 30.5.XX.XX:unknown:DEFAULT
	DatumKey *string `json:"DatumKey,omitempty" xml:"DatumKey,omitempty"`
	// The default key.
	//
	// example:
	//
	// 30.5.XX.XX:unknown
	DefaultKey *string `json:"DefaultKey,omitempty" xml:"DefaultKey,omitempty"`
	// The effective status of the instance. Valid values:
	//
	// 	- `true`: The instance takes effect.
	//
	// 	- `false`: The instance does not take effect.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether the information about the ephemeral node is obtained. Valid values:
	//
	// 	- `true`: yes
	//
	// 	- `false`: no
	//
	// example:
	//
	// true
	Ephemeral *bool `json:"Ephemeral,omitempty" xml:"Ephemeral,omitempty"`
	// The number of counted failures.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The health status of the instance. Valid values:
	//
	// 	- `true`: The instance is healthy.
	//
	// 	- `false`: The instance is unhealthy.
	//
	// example:
	//
	// true
	Healthy *bool `json:"Healthy,omitempty" xml:"Healthy,omitempty"`
	// The heartbeat interval of the instance. Unit: seconds.
	//
	// example:
	//
	// 5000
	InstanceHeartBeatInterval *int32 `json:"InstanceHeartBeatInterval,omitempty" xml:"InstanceHeartBeatInterval,omitempty"`
	// The timeout period of the instance heartbeat.
	//
	// example:
	//
	// 15000
	InstanceHeartBeatTimeOut *int32 `json:"InstanceHeartBeatTimeOut,omitempty" xml:"InstanceHeartBeatTimeOut,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 30.5.XX.XX#0#DEFAULT#DEFAULT_GROUP@@consumers:com.alibaba.edas.IHelloService
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 30.5.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The timeout period for removing an IP address.
	//
	// example:
	//
	// 30000
	IpDeleteTimeout *int32 `json:"IpDeleteTimeout,omitempty" xml:"IpDeleteTimeout,omitempty"`
	// The last heartbeat time.
	//
	// example:
	//
	// 20201010
	LastBeat *int64 `json:"LastBeat,omitempty" xml:"LastBeat,omitempty"`
	// Indicates whether the instance was marked. Valid values:
	//
	// 	- `true`: The instance marking was successful.
	//
	// 	- `false`: The instance marking failed.
	//
	// example:
	//
	// true
	Marked *bool `json:"Marked,omitempty" xml:"Marked,omitempty"`
	// The metadata.
	//
	// example:
	//
	// [int]
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The number of counted successes.
	//
	// example:
	//
	// 0
	OkCount *int32 `json:"OkCount,omitempty" xml:"OkCount,omitempty"`
	// The port number.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// DEFAULT_GROUP@@consumers:com.alibaba.edas.IHelloService::
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The weight.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListAnsInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAnsInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAnsInstancesResponseBodyData) GetApp() *string {
	return s.App
}

func (s *ListAnsInstancesResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListAnsInstancesResponseBodyData) GetDatumKey() *string {
	return s.DatumKey
}

func (s *ListAnsInstancesResponseBodyData) GetDefaultKey() *string {
	return s.DefaultKey
}

func (s *ListAnsInstancesResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListAnsInstancesResponseBodyData) GetEphemeral() *bool {
	return s.Ephemeral
}

func (s *ListAnsInstancesResponseBodyData) GetFailCount() *int32 {
	return s.FailCount
}

func (s *ListAnsInstancesResponseBodyData) GetHealthy() *bool {
	return s.Healthy
}

func (s *ListAnsInstancesResponseBodyData) GetInstanceHeartBeatInterval() *int32 {
	return s.InstanceHeartBeatInterval
}

func (s *ListAnsInstancesResponseBodyData) GetInstanceHeartBeatTimeOut() *int32 {
	return s.InstanceHeartBeatTimeOut
}

func (s *ListAnsInstancesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAnsInstancesResponseBodyData) GetIp() *string {
	return s.Ip
}

func (s *ListAnsInstancesResponseBodyData) GetIpDeleteTimeout() *int32 {
	return s.IpDeleteTimeout
}

func (s *ListAnsInstancesResponseBodyData) GetLastBeat() *int64 {
	return s.LastBeat
}

func (s *ListAnsInstancesResponseBodyData) GetMarked() *bool {
	return s.Marked
}

func (s *ListAnsInstancesResponseBodyData) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *ListAnsInstancesResponseBodyData) GetOkCount() *int32 {
	return s.OkCount
}

func (s *ListAnsInstancesResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *ListAnsInstancesResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListAnsInstancesResponseBodyData) GetWeight() *int32 {
	return s.Weight
}

func (s *ListAnsInstancesResponseBodyData) SetApp(v string) *ListAnsInstancesResponseBodyData {
	s.App = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetClusterName(v string) *ListAnsInstancesResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetDatumKey(v string) *ListAnsInstancesResponseBodyData {
	s.DatumKey = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetDefaultKey(v string) *ListAnsInstancesResponseBodyData {
	s.DefaultKey = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetEnabled(v bool) *ListAnsInstancesResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetEphemeral(v bool) *ListAnsInstancesResponseBodyData {
	s.Ephemeral = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetFailCount(v int32) *ListAnsInstancesResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetHealthy(v bool) *ListAnsInstancesResponseBodyData {
	s.Healthy = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetInstanceHeartBeatInterval(v int32) *ListAnsInstancesResponseBodyData {
	s.InstanceHeartBeatInterval = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetInstanceHeartBeatTimeOut(v int32) *ListAnsInstancesResponseBodyData {
	s.InstanceHeartBeatTimeOut = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetInstanceId(v string) *ListAnsInstancesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetIp(v string) *ListAnsInstancesResponseBodyData {
	s.Ip = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetIpDeleteTimeout(v int32) *ListAnsInstancesResponseBodyData {
	s.IpDeleteTimeout = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetLastBeat(v int64) *ListAnsInstancesResponseBodyData {
	s.LastBeat = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetMarked(v bool) *ListAnsInstancesResponseBodyData {
	s.Marked = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetMetadata(v map[string]interface{}) *ListAnsInstancesResponseBodyData {
	s.Metadata = v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetOkCount(v int32) *ListAnsInstancesResponseBodyData {
	s.OkCount = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetPort(v int32) *ListAnsInstancesResponseBodyData {
	s.Port = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetServiceName(v string) *ListAnsInstancesResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetWeight(v int32) *ListAnsInstancesResponseBodyData {
	s.Weight = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
