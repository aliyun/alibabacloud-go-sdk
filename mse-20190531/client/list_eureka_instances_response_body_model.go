// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEurekaInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListEurekaInstancesResponseBodyData) *ListEurekaInstancesResponseBody
	GetData() []*ListEurekaInstancesResponseBodyData
	SetErrorCode(v string) *ListEurekaInstancesResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListEurekaInstancesResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListEurekaInstancesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListEurekaInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEurekaInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListEurekaInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEurekaInstancesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListEurekaInstancesResponseBody
	GetTotalCount() *int32
}

type ListEurekaInstancesResponseBody struct {
	// The details of the data.
	Data []*ListEurekaInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// The request is successfully processed.
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
	// 316F5F64-F73D-42DC-8632-01E308B6****
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

func (s ListEurekaInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEurekaInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEurekaInstancesResponseBody) GetData() []*ListEurekaInstancesResponseBodyData {
	return s.Data
}

func (s *ListEurekaInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListEurekaInstancesResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListEurekaInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEurekaInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEurekaInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEurekaInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEurekaInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEurekaInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEurekaInstancesResponseBody) SetData(v []*ListEurekaInstancesResponseBodyData) *ListEurekaInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetErrorCode(v string) *ListEurekaInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetHttpCode(v string) *ListEurekaInstancesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetMessage(v string) *ListEurekaInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetPageNumber(v int32) *ListEurekaInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetPageSize(v int32) *ListEurekaInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetRequestId(v string) *ListEurekaInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetSuccess(v bool) *ListEurekaInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetTotalCount(v int32) *ListEurekaInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) Validate() error {
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

type ListEurekaInstancesResponseBodyData struct {
	// The name of the application.
	//
	// example:
	//
	// CONTACTINFO
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The timeout period of the instance.\\
	//
	// After the specified timeout period expires, the service is unavailable by default and is deleted.
	//
	// example:
	//
	// 90
	DurationInSecs *int32 `json:"DurationInSecs,omitempty" xml:"DurationInSecs,omitempty"`
	// The URL of the homepage.
	//
	// example:
	//
	// http://30.5.XX.XX:8091/
	HomePageUrl *string `json:"HomePageUrl,omitempty" xml:"HomePageUrl,omitempty"`
	// The hostname.
	//
	// example:
	//
	// 30.5.XX.XX
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// L-PC1A6A28-****.hz.ali.com:contactinfo:8091
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 30.5.XX.XX
	IpAddr *string `json:"IpAddr,omitempty" xml:"IpAddr,omitempty"`
	// The time when the instance was last modified.
	//
	// example:
	//
	// 20201009115543
	LastDirtyTimestamp *int64 `json:"LastDirtyTimestamp,omitempty" xml:"LastDirtyTimestamp,omitempty"`
	// The time when the instance heartbeat was last checked.
	//
	// example:
	//
	// 20201010071203
	LastUpdatedTimestamp *int64 `json:"LastUpdatedTimestamp,omitempty" xml:"LastUpdatedTimestamp,omitempty"`
	// The metadata.
	//
	// example:
	//
	// [string]
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The service port number.
	//
	// example:
	//
	// 8091
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The maximum interval between two heartbeat checks after a heartbeat check times out.\\
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	RenewalIntervalInSecs *int32 `json:"RenewalIntervalInSecs,omitempty" xml:"RenewalIntervalInSecs,omitempty"`
	// The security port.
	//
	// example:
	//
	// 443
	SecurePort *int32 `json:"SecurePort,omitempty" xml:"SecurePort,omitempty"`
	// The number of service providers. The value is in the following format: Number of healthy instances/Total number of instances.
	//
	// example:
	//
	// 1/1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The virtual IP address (VIP).
	//
	// example:
	//
	// contactinfo
	VipAddress *string `json:"VipAddress,omitempty" xml:"VipAddress,omitempty"`
}

func (s ListEurekaInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEurekaInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEurekaInstancesResponseBodyData) GetApp() *string {
	return s.App
}

func (s *ListEurekaInstancesResponseBodyData) GetDurationInSecs() *int32 {
	return s.DurationInSecs
}

func (s *ListEurekaInstancesResponseBodyData) GetHomePageUrl() *string {
	return s.HomePageUrl
}

func (s *ListEurekaInstancesResponseBodyData) GetHostName() *string {
	return s.HostName
}

func (s *ListEurekaInstancesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEurekaInstancesResponseBodyData) GetIpAddr() *string {
	return s.IpAddr
}

func (s *ListEurekaInstancesResponseBodyData) GetLastDirtyTimestamp() *int64 {
	return s.LastDirtyTimestamp
}

func (s *ListEurekaInstancesResponseBodyData) GetLastUpdatedTimestamp() *int64 {
	return s.LastUpdatedTimestamp
}

func (s *ListEurekaInstancesResponseBodyData) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *ListEurekaInstancesResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *ListEurekaInstancesResponseBodyData) GetRenewalIntervalInSecs() *int32 {
	return s.RenewalIntervalInSecs
}

func (s *ListEurekaInstancesResponseBodyData) GetSecurePort() *int32 {
	return s.SecurePort
}

func (s *ListEurekaInstancesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListEurekaInstancesResponseBodyData) GetVipAddress() *string {
	return s.VipAddress
}

func (s *ListEurekaInstancesResponseBodyData) SetApp(v string) *ListEurekaInstancesResponseBodyData {
	s.App = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetDurationInSecs(v int32) *ListEurekaInstancesResponseBodyData {
	s.DurationInSecs = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetHomePageUrl(v string) *ListEurekaInstancesResponseBodyData {
	s.HomePageUrl = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetHostName(v string) *ListEurekaInstancesResponseBodyData {
	s.HostName = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetInstanceId(v string) *ListEurekaInstancesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetIpAddr(v string) *ListEurekaInstancesResponseBodyData {
	s.IpAddr = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetLastDirtyTimestamp(v int64) *ListEurekaInstancesResponseBodyData {
	s.LastDirtyTimestamp = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetLastUpdatedTimestamp(v int64) *ListEurekaInstancesResponseBodyData {
	s.LastUpdatedTimestamp = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetMetadata(v map[string]interface{}) *ListEurekaInstancesResponseBodyData {
	s.Metadata = v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetPort(v int32) *ListEurekaInstancesResponseBodyData {
	s.Port = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetRenewalIntervalInSecs(v int32) *ListEurekaInstancesResponseBodyData {
	s.RenewalIntervalInSecs = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetSecurePort(v int32) *ListEurekaInstancesResponseBodyData {
	s.SecurePort = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetStatus(v string) *ListEurekaInstancesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetVipAddress(v string) *ListEurekaInstancesResponseBodyData {
	s.VipAddress = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
