// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnsServiceClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListAnsServiceClustersResponseBodyData) *ListAnsServiceClustersResponseBody
	GetData() *ListAnsServiceClustersResponseBodyData
	SetErrorCode(v string) *ListAnsServiceClustersResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListAnsServiceClustersResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListAnsServiceClustersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAnsServiceClustersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAnsServiceClustersResponseBody
	GetSuccess() *bool
}

type ListAnsServiceClustersResponseBody struct {
	// The data returned.
	Data *ListAnsServiceClustersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
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
}

func (s ListAnsServiceClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAnsServiceClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnsServiceClustersResponseBody) GetData() *ListAnsServiceClustersResponseBodyData {
	return s.Data
}

func (s *ListAnsServiceClustersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAnsServiceClustersResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListAnsServiceClustersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAnsServiceClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAnsServiceClustersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAnsServiceClustersResponseBody) SetData(v *ListAnsServiceClustersResponseBodyData) *ListAnsServiceClustersResponseBody {
	s.Data = v
	return s
}

func (s *ListAnsServiceClustersResponseBody) SetErrorCode(v string) *ListAnsServiceClustersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAnsServiceClustersResponseBody) SetHttpCode(v string) *ListAnsServiceClustersResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAnsServiceClustersResponseBody) SetMessage(v string) *ListAnsServiceClustersResponseBody {
	s.Message = &v
	return s
}

func (s *ListAnsServiceClustersResponseBody) SetRequestId(v string) *ListAnsServiceClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnsServiceClustersResponseBody) SetSuccess(v bool) *ListAnsServiceClustersResponseBody {
	s.Success = &v
	return s
}

func (s *ListAnsServiceClustersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAnsServiceClustersResponseBodyData struct {
	// The information about the associated application for which Microservices Governance is enabled when the Source parameter is set to governance.
	AppDetail *ListAnsServiceClustersResponseBodyDataAppDetail `json:"AppDetail,omitempty" xml:"AppDetail,omitempty" type:"Struct"`
	// The cluster information.
	Clusters []*ListAnsServiceClustersResponseBodyDataClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// Indicates whether the service is a temporary service. Valid values:
	//
	// 	- `true`: yes
	//
	// 	- `false`: no
	//
	// example:
	//
	// true
	Ephemeral *bool `json:"Ephemeral,omitempty" xml:"Ephemeral,omitempty"`
	// The service group to which the service belongs.
	//
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The metadata of the service.
	//
	// example:
	//
	// 111
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// nacos.test.3
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The protection threshold.
	//
	// example:
	//
	// 0
	ProtectThreshold *float32 `json:"ProtectThreshold,omitempty" xml:"ProtectThreshold,omitempty"`
	// The election mode.
	//
	// example:
	//
	// none
	SelectorType *string `json:"SelectorType,omitempty" xml:"SelectorType,omitempty"`
	// The source type of the service. Valid values:
	//
	// 	- console: The service was registered in the console.
	//
	// 	- sdk: The service was registered by using the SDK.
	//
	// 	- governance: The service was registered on Microservices Governance.
	//
	// example:
	//
	// console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListAnsServiceClustersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAnsServiceClustersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAnsServiceClustersResponseBodyData) GetAppDetail() *ListAnsServiceClustersResponseBodyDataAppDetail {
	return s.AppDetail
}

func (s *ListAnsServiceClustersResponseBodyData) GetClusters() []*ListAnsServiceClustersResponseBodyDataClusters {
	return s.Clusters
}

func (s *ListAnsServiceClustersResponseBodyData) GetEphemeral() *bool {
	return s.Ephemeral
}

func (s *ListAnsServiceClustersResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAnsServiceClustersResponseBodyData) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *ListAnsServiceClustersResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAnsServiceClustersResponseBodyData) GetProtectThreshold() *float32 {
	return s.ProtectThreshold
}

func (s *ListAnsServiceClustersResponseBodyData) GetSelectorType() *string {
	return s.SelectorType
}

func (s *ListAnsServiceClustersResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *ListAnsServiceClustersResponseBodyData) SetAppDetail(v *ListAnsServiceClustersResponseBodyDataAppDetail) *ListAnsServiceClustersResponseBodyData {
	s.AppDetail = v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetClusters(v []*ListAnsServiceClustersResponseBodyDataClusters) *ListAnsServiceClustersResponseBodyData {
	s.Clusters = v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetEphemeral(v bool) *ListAnsServiceClustersResponseBodyData {
	s.Ephemeral = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetGroupName(v string) *ListAnsServiceClustersResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetMetadata(v map[string]interface{}) *ListAnsServiceClustersResponseBodyData {
	s.Metadata = v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetName(v string) *ListAnsServiceClustersResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetProtectThreshold(v float32) *ListAnsServiceClustersResponseBodyData {
	s.ProtectThreshold = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetSelectorType(v string) *ListAnsServiceClustersResponseBodyData {
	s.SelectorType = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetSource(v string) *ListAnsServiceClustersResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) Validate() error {
	if s.AppDetail != nil {
		if err := s.AppDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Clusters != nil {
		for _, item := range s.Clusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAnsServiceClustersResponseBodyDataAppDetail struct {
	// The ID of the application for which Microservices Governance is enabled.
	//
	// example:
	//
	// hkhon1po62@904cba2c0xxxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application for which Microservices Governance is enabled.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The health check interval. Unit: seconds.
	//
	// example:
	//
	// 2
	CheckInternal *int32 `json:"CheckInternal,omitempty" xml:"CheckInternal,omitempty"`
	// The path of the health check. This parameter is required only when the CheckType parameter is set to http.
	//
	// example:
	//
	// /health
	CheckPath *string `json:"CheckPath,omitempty" xml:"CheckPath,omitempty"`
	// The timeout period of the health check response. Unit: seconds.
	//
	// example:
	//
	// 5
	CheckTimeout *int32 `json:"CheckTimeout,omitempty" xml:"CheckTimeout,omitempty"`
	// The type of the health check. Valid values:
	//
	// 	- connection: connection status check
	//
	// 	- tcp: TCP connection check
	//
	// 	- http: HTTP connection check
	//
	// example:
	//
	// connection
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The maximum number of health check retries when the instance state changes from unhealthy to healthy.
	//
	// example:
	//
	// 2
	HealthyCheckTimes *int32 `json:"HealthyCheckTimes,omitempty" xml:"HealthyCheckTimes,omitempty"`
	// The port number of the application for which Microservices Governance is enabled.
	//
	// example:
	//
	// 6001
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The maximum number of health check retries when the instance state changes from healthy to unhealthy.
	//
	// example:
	//
	// 2
	UnhealthyCheckTimes *int32 `json:"UnhealthyCheckTimes,omitempty" xml:"UnhealthyCheckTimes,omitempty"`
}

func (s ListAnsServiceClustersResponseBodyDataAppDetail) String() string {
	return dara.Prettify(s)
}

func (s ListAnsServiceClustersResponseBodyDataAppDetail) GoString() string {
	return s.String()
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) GetAppId() *string {
	return s.AppId
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) GetAppName() *string {
	return s.AppName
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) GetCheckInternal() *int32 {
	return s.CheckInternal
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) GetCheckPath() *string {
	return s.CheckPath
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) GetCheckTimeout() *int32 {
	return s.CheckTimeout
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) GetCheckType() *string {
	return s.CheckType
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) GetHealthyCheckTimes() *int32 {
	return s.HealthyCheckTimes
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) GetPort() *int32 {
	return s.Port
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) GetUnhealthyCheckTimes() *int32 {
	return s.UnhealthyCheckTimes
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) SetAppId(v string) *ListAnsServiceClustersResponseBodyDataAppDetail {
	s.AppId = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) SetAppName(v string) *ListAnsServiceClustersResponseBodyDataAppDetail {
	s.AppName = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) SetCheckInternal(v int32) *ListAnsServiceClustersResponseBodyDataAppDetail {
	s.CheckInternal = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) SetCheckPath(v string) *ListAnsServiceClustersResponseBodyDataAppDetail {
	s.CheckPath = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) SetCheckTimeout(v int32) *ListAnsServiceClustersResponseBodyDataAppDetail {
	s.CheckTimeout = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) SetCheckType(v string) *ListAnsServiceClustersResponseBodyDataAppDetail {
	s.CheckType = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) SetHealthyCheckTimes(v int32) *ListAnsServiceClustersResponseBodyDataAppDetail {
	s.HealthyCheckTimes = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) SetPort(v int32) *ListAnsServiceClustersResponseBodyDataAppDetail {
	s.Port = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) SetUnhealthyCheckTimes(v int32) *ListAnsServiceClustersResponseBodyDataAppDetail {
	s.UnhealthyCheckTimes = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataAppDetail) Validate() error {
	return dara.Validate(s)
}

type ListAnsServiceClustersResponseBodyDataClusters struct {
	// The default port used for a health check.
	//
	// example:
	//
	// 80
	DefaultCheckPort *int32 `json:"DefaultCheckPort,omitempty" xml:"DefaultCheckPort,omitempty"`
	// The default port.
	//
	// example:
	//
	// 80
	DefaultPort *int32 `json:"DefaultPort,omitempty" xml:"DefaultPort,omitempty"`
	// The type of the health check.
	//
	// example:
	//
	// Heartbeat Reporting
	HealthCheckerType *string `json:"HealthCheckerType,omitempty" xml:"HealthCheckerType,omitempty"`
	// The metadata of the cluster.
	//
	// example:
	//
	// 111
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The full name of the service.
	//
	// example:
	//
	// DEFAULT_GROUP@@nacos.test.3
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// Indicates whether an end-to-end health check was initiated by the server. This parameter is valid only if the service is not a temporary service.
	//
	// example:
	//
	// true
	UseIPPort4Check *bool `json:"UseIPPort4Check,omitempty" xml:"UseIPPort4Check,omitempty"`
}

func (s ListAnsServiceClustersResponseBodyDataClusters) String() string {
	return dara.Prettify(s)
}

func (s ListAnsServiceClustersResponseBodyDataClusters) GoString() string {
	return s.String()
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) GetDefaultCheckPort() *int32 {
	return s.DefaultCheckPort
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) GetDefaultPort() *int32 {
	return s.DefaultPort
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) GetHealthCheckerType() *string {
	return s.HealthCheckerType
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) GetName() *string {
	return s.Name
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) GetUseIPPort4Check() *bool {
	return s.UseIPPort4Check
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetDefaultCheckPort(v int32) *ListAnsServiceClustersResponseBodyDataClusters {
	s.DefaultCheckPort = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetDefaultPort(v int32) *ListAnsServiceClustersResponseBodyDataClusters {
	s.DefaultPort = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetHealthCheckerType(v string) *ListAnsServiceClustersResponseBodyDataClusters {
	s.HealthCheckerType = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetMetadata(v map[string]interface{}) *ListAnsServiceClustersResponseBodyDataClusters {
	s.Metadata = v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetName(v string) *ListAnsServiceClustersResponseBodyDataClusters {
	s.Name = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetServiceName(v string) *ListAnsServiceClustersResponseBodyDataClusters {
	s.ServiceName = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetUseIPPort4Check(v bool) *ListAnsServiceClustersResponseBodyDataClusters {
	s.UseIPPort4Check = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) Validate() error {
	return dara.Validate(s)
}
