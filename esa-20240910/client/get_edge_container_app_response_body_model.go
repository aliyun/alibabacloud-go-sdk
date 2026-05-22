// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v *GetEdgeContainerAppResponseBodyApp) *GetEdgeContainerAppResponseBody
	GetApp() *GetEdgeContainerAppResponseBodyApp
	SetRequestId(v string) *GetEdgeContainerAppResponseBody
	GetRequestId() *string
}

type GetEdgeContainerAppResponseBody struct {
	// The basic information about the application.
	App *GetEdgeContainerAppResponseBodyApp `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 156A6B-677B1A-4297B7-9187B7-2B44792
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEdgeContainerAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResponseBody) GetApp() *GetEdgeContainerAppResponseBodyApp {
	return s.App
}

func (s *GetEdgeContainerAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerAppResponseBody) SetApp(v *GetEdgeContainerAppResponseBodyApp) *GetEdgeContainerAppResponseBody {
	s.App = v
	return s
}

func (s *GetEdgeContainerAppResponseBody) SetRequestId(v string) *GetEdgeContainerAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerAppResponseBody) Validate() error {
	if s.App != nil {
		if err := s.App.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEdgeContainerAppResponseBodyApp struct {
	// The application ID.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2023-07-25T05:58:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The domain name that is associated with the application. If no domain name is associated with the application, the value is an empty string.
	//
	// example:
	//
	// www.1feel.cn
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The type of the gateway. Valid values:
	//
	// 	- l7: Layer 7 gateway.
	//
	// 	- l4: Layer 4 gateway.
	//
	// example:
	//
	// l7
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// The information about health checks.
	HealthCheck *GetEdgeContainerAppResponseBodyAppHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
	// The application name.
	//
	// example:
	//
	// test-app1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether QUIC is enabled.
	//
	// example:
	//
	// false
	QuicCid *string `json:"QuicCid,omitempty" xml:"QuicCid,omitempty"`
	// The remarks about the application.
	//
	// example:
	//
	// test app
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The server port. Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	ServicePort *int32 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- creating: The application is being created.
	//
	// 	- failed: The application failed to be created.
	//
	// 	- created: The application is created.
	//
	// example:
	//
	// created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The backend port, which is also the service port of the application. Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The time when the application was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-26T02:35:58Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The number of versions of the application.
	//
	// example:
	//
	// 1
	VersionCount *int32 `json:"VersionCount,omitempty" xml:"VersionCount,omitempty"`
}

func (s GetEdgeContainerAppResponseBodyApp) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResponseBodyApp) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResponseBodyApp) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerAppResponseBodyApp) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEdgeContainerAppResponseBodyApp) GetDomainName() *string {
	return s.DomainName
}

func (s *GetEdgeContainerAppResponseBodyApp) GetGatewayType() *string {
	return s.GatewayType
}

func (s *GetEdgeContainerAppResponseBodyApp) GetHealthCheck() *GetEdgeContainerAppResponseBodyAppHealthCheck {
	return s.HealthCheck
}

func (s *GetEdgeContainerAppResponseBodyApp) GetName() *string {
	return s.Name
}

func (s *GetEdgeContainerAppResponseBodyApp) GetQuicCid() *string {
	return s.QuicCid
}

func (s *GetEdgeContainerAppResponseBodyApp) GetRemarks() *string {
	return s.Remarks
}

func (s *GetEdgeContainerAppResponseBodyApp) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *GetEdgeContainerAppResponseBodyApp) GetStatus() *string {
	return s.Status
}

func (s *GetEdgeContainerAppResponseBodyApp) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *GetEdgeContainerAppResponseBodyApp) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetEdgeContainerAppResponseBodyApp) GetVersionCount() *int32 {
	return s.VersionCount
}

func (s *GetEdgeContainerAppResponseBodyApp) SetAppId(v string) *GetEdgeContainerAppResponseBodyApp {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) SetCreateTime(v string) *GetEdgeContainerAppResponseBodyApp {
	s.CreateTime = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) SetDomainName(v string) *GetEdgeContainerAppResponseBodyApp {
	s.DomainName = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) SetGatewayType(v string) *GetEdgeContainerAppResponseBodyApp {
	s.GatewayType = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) SetHealthCheck(v *GetEdgeContainerAppResponseBodyAppHealthCheck) *GetEdgeContainerAppResponseBodyApp {
	s.HealthCheck = v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) SetName(v string) *GetEdgeContainerAppResponseBodyApp {
	s.Name = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) SetQuicCid(v string) *GetEdgeContainerAppResponseBodyApp {
	s.QuicCid = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) SetRemarks(v string) *GetEdgeContainerAppResponseBodyApp {
	s.Remarks = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) SetServicePort(v int32) *GetEdgeContainerAppResponseBodyApp {
	s.ServicePort = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) SetStatus(v string) *GetEdgeContainerAppResponseBodyApp {
	s.Status = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) SetTargetPort(v int32) *GetEdgeContainerAppResponseBodyApp {
	s.TargetPort = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) SetUpdateTime(v string) *GetEdgeContainerAppResponseBodyApp {
	s.UpdateTime = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) SetVersionCount(v int32) *GetEdgeContainerAppResponseBodyApp {
	s.VersionCount = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyApp) Validate() error {
	if s.HealthCheck != nil {
		if err := s.HealthCheck.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEdgeContainerAppResponseBodyAppHealthCheck struct {
	// The number of consecutive failed health checks required for an application to be considered as unhealthy.
	//
	// example:
	//
	// 5
	FailTimes *int32 `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	// The domain name that is used for health checks.
	//
	// example:
	//
	// test.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The range of health check status codes that indicate successful health checks.
	//
	// example:
	//
	// http_2xx
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The interval between health checks. Unit: seconds.
	//
	// example:
	//
	// 5
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The HTTP method that the health check request uses.
	//
	// example:
	//
	// HEAD
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The health check port.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The number of consecutive successful health checks required for an application to be considered as healthy.
	//
	// example:
	//
	// 3
	SuccTimes *int32 `json:"SuccTimes,omitempty" xml:"SuccTimes,omitempty"`
	// The timeout period of the health check. Unit: seconds.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The health check type. Valid values:
	//
	// 	- l7
	//
	// 	- l4
	//
	// example:
	//
	// l7
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The health check URL.
	//
	// example:
	//
	// /health_check
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s GetEdgeContainerAppResponseBodyAppHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResponseBodyAppHealthCheck) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) GetFailTimes() *int32 {
	return s.FailTimes
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) GetHost() *string {
	return s.Host
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) GetHttpCode() *string {
	return s.HttpCode
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) GetInterval() *int32 {
	return s.Interval
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) GetMethod() *string {
	return s.Method
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) GetPort() *int32 {
	return s.Port
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) GetSuccTimes() *int32 {
	return s.SuccTimes
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) GetType() *string {
	return s.Type
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) GetUri() *string {
	return s.Uri
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) SetFailTimes(v int32) *GetEdgeContainerAppResponseBodyAppHealthCheck {
	s.FailTimes = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) SetHost(v string) *GetEdgeContainerAppResponseBodyAppHealthCheck {
	s.Host = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) SetHttpCode(v string) *GetEdgeContainerAppResponseBodyAppHealthCheck {
	s.HttpCode = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) SetInterval(v int32) *GetEdgeContainerAppResponseBodyAppHealthCheck {
	s.Interval = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) SetMethod(v string) *GetEdgeContainerAppResponseBodyAppHealthCheck {
	s.Method = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) SetPort(v int32) *GetEdgeContainerAppResponseBodyAppHealthCheck {
	s.Port = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) SetSuccTimes(v int32) *GetEdgeContainerAppResponseBodyAppHealthCheck {
	s.SuccTimes = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) SetTimeout(v int32) *GetEdgeContainerAppResponseBodyAppHealthCheck {
	s.Timeout = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) SetType(v string) *GetEdgeContainerAppResponseBodyAppHealthCheck {
	s.Type = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) SetUri(v string) *GetEdgeContainerAppResponseBodyAppHealthCheck {
	s.Uri = &v
	return s
}

func (s *GetEdgeContainerAppResponseBodyAppHealthCheck) Validate() error {
	return dara.Validate(s)
}
