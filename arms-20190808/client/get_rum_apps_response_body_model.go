// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppList(v []*GetRumAppsResponseBodyAppList) *GetRumAppsResponseBody
	GetAppList() []*GetRumAppsResponseBodyAppList
	SetCode(v int32) *GetRumAppsResponseBody
	GetCode() *int32
	SetHttpStatusCode(v int32) *GetRumAppsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRumAppsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRumAppsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRumAppsResponseBody
	GetSuccess() *bool
}

type GetRumAppsResponseBody struct {
	// The queried applications.
	AppList []*GetRumAppsResponseBodyAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Internal error, please contact customer service.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 70675725-8F11-4817-8106-CFE0AD71****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRumAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRumAppsResponseBody) GetAppList() []*GetRumAppsResponseBodyAppList {
	return s.AppList
}

func (s *GetRumAppsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetRumAppsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRumAppsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRumAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRumAppsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRumAppsResponseBody) SetAppList(v []*GetRumAppsResponseBodyAppList) *GetRumAppsResponseBody {
	s.AppList = v
	return s
}

func (s *GetRumAppsResponseBody) SetCode(v int32) *GetRumAppsResponseBody {
	s.Code = &v
	return s
}

func (s *GetRumAppsResponseBody) SetHttpStatusCode(v int32) *GetRumAppsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRumAppsResponseBody) SetMessage(v string) *GetRumAppsResponseBody {
	s.Message = &v
	return s
}

func (s *GetRumAppsResponseBody) SetRequestId(v string) *GetRumAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRumAppsResponseBody) SetSuccess(v bool) *GetRumAppsResponseBody {
	s.Success = &v
	return s
}

func (s *GetRumAppsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRumAppsResponseBodyAppList struct {
	// The application type. Valid values: web, miniapp, ios, and android.
	//
	// example:
	//
	// web
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The time when the application was created. The value is a timestamp.
	//
	// example:
	//
	// 1685686960872
	CreateTime interface{} `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// TEST
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The endpoint that is used to report application data.
	//
	// example:
	//
	// xxxxxxxx-default-cn.rum.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// Indicates whether the application is subscribed. Valid values: true and false.
	//
	// example:
	//
	// true
	IsSubscription *bool `json:"IsSubscription,omitempty" xml:"IsSubscription,omitempty"`
	// The application name.
	//
	// example:
	//
	// home page
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The alias of the application.
	//
	// example:
	//
	// Williamtag
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The package name of the Android application.
	//
	// example:
	//
	// com.zy.yxws
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The application ID.
	//
	// example:
	//
	// xxxxx@cc08bdxxxx20b15
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzaq3ypaqkdy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of service domain configurations. Only mobile applications are supported.
	ServiceDomainConfigs []*GetRumAppsResponseBodyAppListServiceDomainConfigs `json:"ServiceDomainConfigs,omitempty" xml:"ServiceDomainConfigs,omitempty" type:"Repeated"`
	// The name of the Simple Log Service Logstore that stores application data.
	//
	// example:
	//
	// logstore-rum
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	// The name of the Simple Log Service project that stores application data.
	//
	// example:
	//
	// proj-xtrace-xxxxxxxba6ef5466b5debf9e2f951-cn-hangzhou
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	// The status of the application. Valid values: created, running, and stopped.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*GetRumAppsResponseBodyAppListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the application. Valid value: RUM.
	//
	// example:
	//
	// RUM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRumAppsResponseBodyAppList) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppsResponseBodyAppList) GoString() string {
	return s.String()
}

func (s *GetRumAppsResponseBodyAppList) GetAppType() *string {
	return s.AppType
}

func (s *GetRumAppsResponseBodyAppList) GetCreateTime() interface{} {
	return s.CreateTime
}

func (s *GetRumAppsResponseBodyAppList) GetDescription() *string {
	return s.Description
}

func (s *GetRumAppsResponseBodyAppList) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetRumAppsResponseBodyAppList) GetIsSubscription() *bool {
	return s.IsSubscription
}

func (s *GetRumAppsResponseBodyAppList) GetName() *string {
	return s.Name
}

func (s *GetRumAppsResponseBodyAppList) GetNickName() *string {
	return s.NickName
}

func (s *GetRumAppsResponseBodyAppList) GetPackageName() *string {
	return s.PackageName
}

func (s *GetRumAppsResponseBodyAppList) GetPid() *string {
	return s.Pid
}

func (s *GetRumAppsResponseBodyAppList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRumAppsResponseBodyAppList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetRumAppsResponseBodyAppList) GetServiceDomainConfigs() []*GetRumAppsResponseBodyAppListServiceDomainConfigs {
	return s.ServiceDomainConfigs
}

func (s *GetRumAppsResponseBodyAppList) GetSlsLogstore() *string {
	return s.SlsLogstore
}

func (s *GetRumAppsResponseBodyAppList) GetSlsProject() *string {
	return s.SlsProject
}

func (s *GetRumAppsResponseBodyAppList) GetStatus() *string {
	return s.Status
}

func (s *GetRumAppsResponseBodyAppList) GetTags() []*GetRumAppsResponseBodyAppListTags {
	return s.Tags
}

func (s *GetRumAppsResponseBodyAppList) GetType() *string {
	return s.Type
}

func (s *GetRumAppsResponseBodyAppList) SetAppType(v string) *GetRumAppsResponseBodyAppList {
	s.AppType = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetCreateTime(v interface{}) *GetRumAppsResponseBodyAppList {
	s.CreateTime = v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetDescription(v string) *GetRumAppsResponseBodyAppList {
	s.Description = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetEndpoint(v string) *GetRumAppsResponseBodyAppList {
	s.Endpoint = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetIsSubscription(v bool) *GetRumAppsResponseBodyAppList {
	s.IsSubscription = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetName(v string) *GetRumAppsResponseBodyAppList {
	s.Name = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetNickName(v string) *GetRumAppsResponseBodyAppList {
	s.NickName = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetPackageName(v string) *GetRumAppsResponseBodyAppList {
	s.PackageName = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetPid(v string) *GetRumAppsResponseBodyAppList {
	s.Pid = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetRegionId(v string) *GetRumAppsResponseBodyAppList {
	s.RegionId = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetResourceGroupId(v string) *GetRumAppsResponseBodyAppList {
	s.ResourceGroupId = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetServiceDomainConfigs(v []*GetRumAppsResponseBodyAppListServiceDomainConfigs) *GetRumAppsResponseBodyAppList {
	s.ServiceDomainConfigs = v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetSlsLogstore(v string) *GetRumAppsResponseBodyAppList {
	s.SlsLogstore = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetSlsProject(v string) *GetRumAppsResponseBodyAppList {
	s.SlsProject = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetStatus(v string) *GetRumAppsResponseBodyAppList {
	s.Status = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetTags(v []*GetRumAppsResponseBodyAppListTags) *GetRumAppsResponseBodyAppList {
	s.Tags = v
	return s
}

func (s *GetRumAppsResponseBodyAppList) SetType(v string) *GetRumAppsResponseBodyAppList {
	s.Type = &v
	return s
}

func (s *GetRumAppsResponseBodyAppList) Validate() error {
	return dara.Validate(s)
}

type GetRumAppsResponseBodyAppListServiceDomainConfigs struct {
	// The description.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name or IP address.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The trace propagation protocols. This parameter is required if the tracing analysis feature is enabled.
	PropagatorTypes []*string `json:"PropagatorTypes,omitempty" xml:"PropagatorTypes,omitempty" type:"Repeated"`
	// Indicates whether the tracing analysis feature is enabled. To enable the tracing analysis feature, you must activate Managed Service for OpenTelemetry. Valid values:
	//
	// 	- `true`: enables the tracing analysis feature. If you enable the tracing analysis feature, related headers are inserted into requests for the domain name.
	//
	// 	- `false`: disables the tracing analysis feature.
	//
	// example:
	//
	// true
	Tracing *string `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
}

func (s GetRumAppsResponseBodyAppListServiceDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppsResponseBodyAppListServiceDomainConfigs) GoString() string {
	return s.String()
}

func (s *GetRumAppsResponseBodyAppListServiceDomainConfigs) GetDescription() *string {
	return s.Description
}

func (s *GetRumAppsResponseBodyAppListServiceDomainConfigs) GetDomain() *string {
	return s.Domain
}

func (s *GetRumAppsResponseBodyAppListServiceDomainConfigs) GetPropagatorTypes() []*string {
	return s.PropagatorTypes
}

func (s *GetRumAppsResponseBodyAppListServiceDomainConfigs) GetTracing() *string {
	return s.Tracing
}

func (s *GetRumAppsResponseBodyAppListServiceDomainConfigs) SetDescription(v string) *GetRumAppsResponseBodyAppListServiceDomainConfigs {
	s.Description = &v
	return s
}

func (s *GetRumAppsResponseBodyAppListServiceDomainConfigs) SetDomain(v string) *GetRumAppsResponseBodyAppListServiceDomainConfigs {
	s.Domain = &v
	return s
}

func (s *GetRumAppsResponseBodyAppListServiceDomainConfigs) SetPropagatorTypes(v []*string) *GetRumAppsResponseBodyAppListServiceDomainConfigs {
	s.PropagatorTypes = v
	return s
}

func (s *GetRumAppsResponseBodyAppListServiceDomainConfigs) SetTracing(v string) *GetRumAppsResponseBodyAppListServiceDomainConfigs {
	s.Tracing = &v
	return s
}

func (s *GetRumAppsResponseBodyAppListServiceDomainConfigs) Validate() error {
	return dara.Validate(s)
}

type GetRumAppsResponseBodyAppListTags struct {
	// The tag key. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// ok
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRumAppsResponseBodyAppListTags) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppsResponseBodyAppListTags) GoString() string {
	return s.String()
}

func (s *GetRumAppsResponseBodyAppListTags) GetKey() *string {
	return s.Key
}

func (s *GetRumAppsResponseBodyAppListTags) GetValue() *string {
	return s.Value
}

func (s *GetRumAppsResponseBodyAppListTags) SetKey(v string) *GetRumAppsResponseBodyAppListTags {
	s.Key = &v
	return s
}

func (s *GetRumAppsResponseBodyAppListTags) SetValue(v string) *GetRumAppsResponseBodyAppListTags {
	s.Value = &v
	return s
}

func (s *GetRumAppsResponseBodyAppListTags) Validate() error {
	return dara.Validate(s)
}
