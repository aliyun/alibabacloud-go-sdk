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
	App       *GetEdgeContainerAppResponseBodyApp `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AppId        *string                                        `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime   *string                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName   *string                                        `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GatewayType  *string                                        `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	HealthCheck  *GetEdgeContainerAppResponseBodyAppHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
	Name         *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	QuicCid      *string                                        `json:"QuicCid,omitempty" xml:"QuicCid,omitempty"`
	Remarks      *string                                        `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	ServicePort  *int32                                         `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	Status       *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetPort   *int32                                         `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	UpdateTime   *string                                        `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VersionCount *int32                                         `json:"VersionCount,omitempty" xml:"VersionCount,omitempty"`
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
	FailTimes *int32  `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	HttpCode  *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Method    *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Port      *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	SuccTimes *int32  `json:"SuccTimes,omitempty" xml:"SuccTimes,omitempty"`
	Timeout   *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri       *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
