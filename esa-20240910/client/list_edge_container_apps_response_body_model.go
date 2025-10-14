// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApps(v []*ListEdgeContainerAppsResponseBodyApps) *ListEdgeContainerAppsResponseBody
	GetApps() []*ListEdgeContainerAppsResponseBodyApps
	SetPageNumber(v int32) *ListEdgeContainerAppsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeContainerAppsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListEdgeContainerAppsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEdgeContainerAppsResponseBody
	GetTotalCount() *int32
}

type ListEdgeContainerAppsResponseBody struct {
	// The queried applications.
	Apps []*ListEdgeContainerAppsResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// The page number. Default value: **1**. Valid values: 1 to 65535.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 500.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2430E05E-1340-5773-B5E1-B743929F46F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEdgeContainerAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppsResponseBody) GetApps() []*ListEdgeContainerAppsResponseBodyApps {
	return s.Apps
}

func (s *ListEdgeContainerAppsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeContainerAppsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeContainerAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEdgeContainerAppsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEdgeContainerAppsResponseBody) SetApps(v []*ListEdgeContainerAppsResponseBodyApps) *ListEdgeContainerAppsResponseBody {
	s.Apps = v
	return s
}

func (s *ListEdgeContainerAppsResponseBody) SetPageNumber(v int32) *ListEdgeContainerAppsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBody) SetPageSize(v int32) *ListEdgeContainerAppsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBody) SetRequestId(v string) *ListEdgeContainerAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBody) SetTotalCount(v int32) *ListEdgeContainerAppsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBody) Validate() error {
	if s.Apps != nil {
		for _, item := range s.Apps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEdgeContainerAppsResponseBodyApps struct {
	// The application ID.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the version was created.
	//
	// example:
	//
	// 2022-12-01T16:16:27.418298794+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The domain name associated with the application.
	//
	// example:
	//
	// test.com
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
	HealthCheck *ListEdgeContainerAppsResponseBodyAppsHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
	// The application name.
	//
	// example:
	//
	// app-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The progress of the application creation task in percentage.
	//
	// example:
	//
	// 100
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// Indicates whether QUIC is enabled.
	//
	// example:
	//
	// false
	QuicCid *string `json:"QuicCid,omitempty" xml:"QuicCid,omitempty"`
	// The remarks.
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
	// 2023-11-15T12:11:02Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The number of versions of the application.
	//
	// example:
	//
	// 10
	VersionCount *int32 `json:"VersionCount,omitempty" xml:"VersionCount,omitempty"`
}

func (s ListEdgeContainerAppsResponseBodyApps) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppsResponseBodyApps) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetAppId() *string {
	return s.AppId
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetDomainName() *string {
	return s.DomainName
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetHealthCheck() *ListEdgeContainerAppsResponseBodyAppsHealthCheck {
	return s.HealthCheck
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetName() *string {
	return s.Name
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetPercentage() *int32 {
	return s.Percentage
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetQuicCid() *string {
	return s.QuicCid
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetRemarks() *string {
	return s.Remarks
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetStatus() *string {
	return s.Status
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListEdgeContainerAppsResponseBodyApps) GetVersionCount() *int32 {
	return s.VersionCount
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetAppId(v string) *ListEdgeContainerAppsResponseBodyApps {
	s.AppId = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetCreateTime(v string) *ListEdgeContainerAppsResponseBodyApps {
	s.CreateTime = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetDomainName(v string) *ListEdgeContainerAppsResponseBodyApps {
	s.DomainName = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetGatewayType(v string) *ListEdgeContainerAppsResponseBodyApps {
	s.GatewayType = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetHealthCheck(v *ListEdgeContainerAppsResponseBodyAppsHealthCheck) *ListEdgeContainerAppsResponseBodyApps {
	s.HealthCheck = v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetName(v string) *ListEdgeContainerAppsResponseBodyApps {
	s.Name = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetPercentage(v int32) *ListEdgeContainerAppsResponseBodyApps {
	s.Percentage = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetQuicCid(v string) *ListEdgeContainerAppsResponseBodyApps {
	s.QuicCid = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetRemarks(v string) *ListEdgeContainerAppsResponseBodyApps {
	s.Remarks = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetServicePort(v int32) *ListEdgeContainerAppsResponseBodyApps {
	s.ServicePort = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetStatus(v string) *ListEdgeContainerAppsResponseBodyApps {
	s.Status = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetTargetPort(v int32) *ListEdgeContainerAppsResponseBodyApps {
	s.TargetPort = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetUpdateTime(v string) *ListEdgeContainerAppsResponseBodyApps {
	s.UpdateTime = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) SetVersionCount(v int32) *ListEdgeContainerAppsResponseBodyApps {
	s.VersionCount = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyApps) Validate() error {
	if s.HealthCheck != nil {
		if err := s.HealthCheck.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEdgeContainerAppsResponseBodyAppsHealthCheck struct {
	// The number of consecutive failed health checks for an application to be considered unhealthy.
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
	// http_2XX
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The interval between health checks. Unit: seconds.
	//
	// example:
	//
	// 5
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The HTTP method used for health checks.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The health check port.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The number of consecutive successful health checks for an application to be considered healthy.
	//
	// example:
	//
	// 1
	SuccTimes *int32 `json:"SuccTimes,omitempty" xml:"SuccTimes,omitempty"`
	// The timeout period of health checks. Unit: seconds.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The type of health checks. Valid values:
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

func (s ListEdgeContainerAppsResponseBodyAppsHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppsResponseBodyAppsHealthCheck) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) GetFailTimes() *int32 {
	return s.FailTimes
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) GetHost() *string {
	return s.Host
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) GetInterval() *int32 {
	return s.Interval
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) GetMethod() *string {
	return s.Method
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) GetPort() *int32 {
	return s.Port
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) GetSuccTimes() *int32 {
	return s.SuccTimes
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) GetType() *string {
	return s.Type
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) GetUri() *string {
	return s.Uri
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) SetFailTimes(v int32) *ListEdgeContainerAppsResponseBodyAppsHealthCheck {
	s.FailTimes = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) SetHost(v string) *ListEdgeContainerAppsResponseBodyAppsHealthCheck {
	s.Host = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) SetHttpCode(v string) *ListEdgeContainerAppsResponseBodyAppsHealthCheck {
	s.HttpCode = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) SetInterval(v int32) *ListEdgeContainerAppsResponseBodyAppsHealthCheck {
	s.Interval = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) SetMethod(v string) *ListEdgeContainerAppsResponseBodyAppsHealthCheck {
	s.Method = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) SetPort(v int32) *ListEdgeContainerAppsResponseBodyAppsHealthCheck {
	s.Port = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) SetSuccTimes(v int32) *ListEdgeContainerAppsResponseBodyAppsHealthCheck {
	s.SuccTimes = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) SetTimeout(v int32) *ListEdgeContainerAppsResponseBodyAppsHealthCheck {
	s.Timeout = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) SetType(v string) *ListEdgeContainerAppsResponseBodyAppsHealthCheck {
	s.Type = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) SetUri(v string) *ListEdgeContainerAppsResponseBodyAppsHealthCheck {
	s.Uri = &v
	return s
}

func (s *ListEdgeContainerAppsResponseBodyAppsHealthCheck) Validate() error {
	return dara.Validate(s)
}
