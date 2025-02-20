// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type InstanceConfigDto struct {
	// example:
	//
	// storage_root_path
	ConfigKey *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
	// example:
	//
	// BE
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// example:
	//
	// value1
	ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty"`
	// example:
	//
	// ng-e24924dxxxxx
	NodeGroupId *string `json:"nodeGroupId,omitempty" xml:"nodeGroupId,omitempty"`
}

func (s InstanceConfigDto) String() string {
	return tea.Prettify(s)
}

func (s InstanceConfigDto) GoString() string {
	return s.String()
}

func (s *InstanceConfigDto) SetConfigKey(v string) *InstanceConfigDto {
	s.ConfigKey = &v
	return s
}

func (s *InstanceConfigDto) SetConfigType(v string) *InstanceConfigDto {
	s.ConfigType = &v
	return s
}

func (s *InstanceConfigDto) SetConfigValue(v string) *InstanceConfigDto {
	s.ConfigValue = &v
	return s
}

func (s *InstanceConfigDto) SetNodeGroupId(v string) *InstanceConfigDto {
	s.NodeGroupId = &v
	return s
}

type ResourceSpec struct {
	Cu                       *int32  `json:"cu,omitempty" xml:"cu,omitempty"`
	DiskNumber               *int32  `json:"diskNumber,omitempty" xml:"diskNumber,omitempty"`
	LocalStorageInstanceType *string `json:"localStorageInstanceType,omitempty" xml:"localStorageInstanceType,omitempty"`
	NodeNumber               *int32  `json:"nodeNumber,omitempty" xml:"nodeNumber,omitempty"`
	SpecType                 *string `json:"specType,omitempty" xml:"specType,omitempty"`
	StoragePerformanceLevel  *string `json:"storagePerformanceLevel,omitempty" xml:"storagePerformanceLevel,omitempty"`
	StorageSize              *int32  `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
}

func (s ResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s ResourceSpec) GoString() string {
	return s.String()
}

func (s *ResourceSpec) SetCu(v int32) *ResourceSpec {
	s.Cu = &v
	return s
}

func (s *ResourceSpec) SetDiskNumber(v int32) *ResourceSpec {
	s.DiskNumber = &v
	return s
}

func (s *ResourceSpec) SetLocalStorageInstanceType(v string) *ResourceSpec {
	s.LocalStorageInstanceType = &v
	return s
}

func (s *ResourceSpec) SetNodeNumber(v int32) *ResourceSpec {
	s.NodeNumber = &v
	return s
}

func (s *ResourceSpec) SetSpecType(v string) *ResourceSpec {
	s.SpecType = &v
	return s
}

func (s *ResourceSpec) SetStoragePerformanceLevel(v string) *ResourceSpec {
	s.StoragePerformanceLevel = &v
	return s
}

func (s *ResourceSpec) SetStorageSize(v int32) *ResourceSpec {
	s.StorageSize = &v
	return s
}

type ModifyCuRequest struct {
	FastMode *bool `json:"FastMode,omitempty" xml:"FastMode,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The warehouse ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The number of CUs to which you want to change.
	//
	// Valid values:
	//
	// 	- 4
	//
	// 	- 8
	//
	// 	- 16
	//
	// 	- 32
	//
	// 	- 64
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyCuRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCuRequest) GoString() string {
	return s.String()
}

func (s *ModifyCuRequest) SetFastMode(v bool) *ModifyCuRequest {
	s.FastMode = &v
	return s
}

func (s *ModifyCuRequest) SetInstanceId(v string) *ModifyCuRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCuRequest) SetNodeGroupId(v string) *ModifyCuRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyCuRequest) SetTarget(v int32) *ModifyCuRequest {
	s.Target = &v
	return s
}

type ModifyCuResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 24151320976****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyCuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCuResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCuResponseBody) SetAccessDeniedDetail(v string) *ModifyCuResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyCuResponseBody) SetData(v int64) *ModifyCuResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyCuResponseBody) SetErrCode(v string) *ModifyCuResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyCuResponseBody) SetErrMessage(v string) *ModifyCuResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyCuResponseBody) SetHttpStatusCode(v int32) *ModifyCuResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyCuResponseBody) SetRequestId(v string) *ModifyCuResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCuResponseBody) SetSuccess(v bool) *ModifyCuResponseBody {
	s.Success = &v
	return s
}

type ModifyCuResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCuResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCuResponse) GoString() string {
	return s.String()
}

func (s *ModifyCuResponse) SetHeaders(v map[string]*string) *ModifyCuResponse {
	s.Headers = v
	return s
}

func (s *ModifyCuResponse) SetStatusCode(v int32) *ModifyCuResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCuResponse) SetBody(v *ModifyCuResponseBody) *ModifyCuResponse {
	s.Body = v
	return s
}

type ModifyCuPreCheckRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The warehouse ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-d332aa8bca48****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The number of CUs to which you want to change.
	//
	// Valid values:
	//
	// 	- 2
	//
	// 	- 4
	//
	// 	- 8
	//
	// 	- 16
	//
	// 	- 32
	//
	// 	- 64
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyCuPreCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCuPreCheckRequest) GoString() string {
	return s.String()
}

func (s *ModifyCuPreCheckRequest) SetInstanceId(v string) *ModifyCuPreCheckRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCuPreCheckRequest) SetNodeGroupId(v string) *ModifyCuPreCheckRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyCuPreCheckRequest) SetTarget(v int32) *ModifyCuPreCheckRequest {
	s.Target = &v
	return s
}

type ModifyCuPreCheckResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	Data *ModifyCuPreCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyCuPreCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCuPreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCuPreCheckResponseBody) SetAccessDeniedDetail(v string) *ModifyCuPreCheckResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyCuPreCheckResponseBody) SetData(v *ModifyCuPreCheckResponseBodyData) *ModifyCuPreCheckResponseBody {
	s.Data = v
	return s
}

func (s *ModifyCuPreCheckResponseBody) SetErrCode(v string) *ModifyCuPreCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyCuPreCheckResponseBody) SetErrMessage(v string) *ModifyCuPreCheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyCuPreCheckResponseBody) SetHttpStatusCode(v int32) *ModifyCuPreCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyCuPreCheckResponseBody) SetRequestId(v string) *ModifyCuPreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCuPreCheckResponseBody) SetSuccess(v bool) *ModifyCuPreCheckResponseBody {
	s.Success = &v
	return s
}

type ModifyCuPreCheckResponseBodyData struct {
	// Indicates whether the number of CUs can be modified.
	//
	// example:
	//
	// false
	Allow *bool `json:"Allow,omitempty" xml:"Allow,omitempty"`
	// The reason why the number of CUs cannot be modified.
	//
	// example:
	//
	// Failed to find node group[ng-3d5ce6454354****].
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ModifyCuPreCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyCuPreCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyCuPreCheckResponseBodyData) SetAllow(v bool) *ModifyCuPreCheckResponseBodyData {
	s.Allow = &v
	return s
}

func (s *ModifyCuPreCheckResponseBodyData) SetReason(v string) *ModifyCuPreCheckResponseBodyData {
	s.Reason = &v
	return s
}

type ModifyCuPreCheckResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCuPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCuPreCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCuPreCheckResponse) GoString() string {
	return s.String()
}

func (s *ModifyCuPreCheckResponse) SetHeaders(v map[string]*string) *ModifyCuPreCheckResponse {
	s.Headers = v
	return s
}

func (s *ModifyCuPreCheckResponse) SetStatusCode(v int32) *ModifyCuPreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCuPreCheckResponse) SetBody(v *ModifyCuPreCheckResponseBody) *ModifyCuPreCheckResponse {
	s.Body = v
	return s
}

type ModifyDiskNumberRequest struct {
	FastMode *bool `json:"FastMode,omitempty" xml:"FastMode,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The warehouse ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The number of disks to which you want to change to.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyDiskNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskNumberRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskNumberRequest) SetFastMode(v bool) *ModifyDiskNumberRequest {
	s.FastMode = &v
	return s
}

func (s *ModifyDiskNumberRequest) SetInstanceId(v string) *ModifyDiskNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDiskNumberRequest) SetNodeGroupId(v string) *ModifyDiskNumberRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyDiskNumberRequest) SetTarget(v int32) *ModifyDiskNumberRequest {
	s.Target = &v
	return s
}

type ModifyDiskNumberResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 24151320976****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDiskNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskNumberResponseBody) SetAccessDeniedDetail(v string) *ModifyDiskNumberResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) SetData(v int64) *ModifyDiskNumberResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) SetErrCode(v string) *ModifyDiskNumberResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) SetErrMessage(v string) *ModifyDiskNumberResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) SetHttpStatusCode(v int32) *ModifyDiskNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) SetRequestId(v string) *ModifyDiskNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) SetSuccess(v bool) *ModifyDiskNumberResponseBody {
	s.Success = &v
	return s
}

type ModifyDiskNumberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskNumberResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskNumberResponse) SetHeaders(v map[string]*string) *ModifyDiskNumberResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskNumberResponse) SetStatusCode(v int32) *ModifyDiskNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskNumberResponse) SetBody(v *ModifyDiskNumberResponseBody) *ModifyDiskNumberResponse {
	s.Body = v
	return s
}

type ModifyDiskPerformanceLevelRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The warehouse ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The disk performance level to which you want to change.
	//
	// Valid values:
	//
	// 	- pl0
	//
	// 	- pl1
	//
	// 	- pl2
	//
	// 	- pl3
	//
	// This parameter is required.
	//
	// example:
	//
	// pl2
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyDiskPerformanceLevelRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskPerformanceLevelRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskPerformanceLevelRequest) SetInstanceId(v string) *ModifyDiskPerformanceLevelRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDiskPerformanceLevelRequest) SetNodeGroupId(v string) *ModifyDiskPerformanceLevelRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyDiskPerformanceLevelRequest) SetTarget(v string) *ModifyDiskPerformanceLevelRequest {
	s.Target = &v
	return s
}

type ModifyDiskPerformanceLevelResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 24151320976****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDiskPerformanceLevelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskPerformanceLevelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetAccessDeniedDetail(v string) *ModifyDiskPerformanceLevelResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetData(v int64) *ModifyDiskPerformanceLevelResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetErrCode(v string) *ModifyDiskPerformanceLevelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetErrMessage(v string) *ModifyDiskPerformanceLevelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetHttpStatusCode(v int32) *ModifyDiskPerformanceLevelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetRequestId(v string) *ModifyDiskPerformanceLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetSuccess(v bool) *ModifyDiskPerformanceLevelResponseBody {
	s.Success = &v
	return s
}

type ModifyDiskPerformanceLevelResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskPerformanceLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskPerformanceLevelResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskPerformanceLevelResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskPerformanceLevelResponse) SetHeaders(v map[string]*string) *ModifyDiskPerformanceLevelResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskPerformanceLevelResponse) SetStatusCode(v int32) *ModifyDiskPerformanceLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponse) SetBody(v *ModifyDiskPerformanceLevelResponseBody) *ModifyDiskPerformanceLevelResponse {
	s.Body = v
	return s
}

type ModifyDiskSizeRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The warehouse ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The disk size to which you want to change to. Unit: GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 500
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyDiskSizeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskSizeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskSizeRequest) SetInstanceId(v string) *ModifyDiskSizeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDiskSizeRequest) SetNodeGroupId(v string) *ModifyDiskSizeRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyDiskSizeRequest) SetTarget(v int32) *ModifyDiskSizeRequest {
	s.Target = &v
	return s
}

type ModifyDiskSizeResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 24151320976****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDiskSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskSizeResponseBody) SetAccessDeniedDetail(v string) *ModifyDiskSizeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) SetData(v int64) *ModifyDiskSizeResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) SetErrCode(v string) *ModifyDiskSizeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) SetErrMessage(v string) *ModifyDiskSizeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) SetHttpStatusCode(v int32) *ModifyDiskSizeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) SetRequestId(v string) *ModifyDiskSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) SetSuccess(v bool) *ModifyDiskSizeResponseBody {
	s.Success = &v
	return s
}

type ModifyDiskSizeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskSizeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskSizeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskSizeResponse) SetHeaders(v map[string]*string) *ModifyDiskSizeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskSizeResponse) SetStatusCode(v int32) *ModifyDiskSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskSizeResponse) SetBody(v *ModifyDiskSizeResponseBody) *ModifyDiskSizeResponse {
	s.Body = v
	return s
}

type ModifyNodeNumberRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The warehouse ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The number of nodes to which you want to change to.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyNodeNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeNumberRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberRequest) SetInstanceId(v string) *ModifyNodeNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyNodeNumberRequest) SetNodeGroupId(v string) *ModifyNodeNumberRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyNodeNumberRequest) SetTarget(v int32) *ModifyNodeNumberRequest {
	s.Target = &v
	return s
}

type ModifyNodeNumberResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 24151320976****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyNodeNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberResponseBody) SetAccessDeniedDetail(v string) *ModifyNodeNumberResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) SetData(v int64) *ModifyNodeNumberResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) SetErrCode(v string) *ModifyNodeNumberResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) SetErrMessage(v string) *ModifyNodeNumberResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) SetHttpStatusCode(v int32) *ModifyNodeNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) SetRequestId(v string) *ModifyNodeNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) SetSuccess(v bool) *ModifyNodeNumberResponseBody {
	s.Success = &v
	return s
}

type ModifyNodeNumberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodeNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodeNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeNumberResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberResponse) SetHeaders(v map[string]*string) *ModifyNodeNumberResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodeNumberResponse) SetStatusCode(v int32) *ModifyNodeNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodeNumberResponse) SetBody(v *ModifyNodeNumberResponseBody) *ModifyNodeNumberResponse {
	s.Body = v
	return s
}

type ModifyNodeNumberPreCheckRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The warehouse ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The number of nodes to which you want to change to.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyNodeNumberPreCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeNumberPreCheckRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberPreCheckRequest) SetInstanceId(v string) *ModifyNodeNumberPreCheckRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyNodeNumberPreCheckRequest) SetNodeGroupId(v string) *ModifyNodeNumberPreCheckRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyNodeNumberPreCheckRequest) SetTarget(v int32) *ModifyNodeNumberPreCheckRequest {
	s.Target = &v
	return s
}

type ModifyNodeNumberPreCheckResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	Data *ModifyNodeNumberPreCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyNodeNumberPreCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeNumberPreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetAccessDeniedDetail(v string) *ModifyNodeNumberPreCheckResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetData(v *ModifyNodeNumberPreCheckResponseBodyData) *ModifyNodeNumberPreCheckResponseBody {
	s.Data = v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetErrCode(v string) *ModifyNodeNumberPreCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetErrMessage(v string) *ModifyNodeNumberPreCheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetHttpStatusCode(v int32) *ModifyNodeNumberPreCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetRequestId(v string) *ModifyNodeNumberPreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetSuccess(v bool) *ModifyNodeNumberPreCheckResponseBody {
	s.Success = &v
	return s
}

type ModifyNodeNumberPreCheckResponseBodyData struct {
	// Indicates whether the number of nodes can be modified.
	//
	// example:
	//
	// true
	Allow *bool `json:"Allow,omitempty" xml:"Allow,omitempty"`
	// The reason why the number of nodes cannot be modified.
	//
	// example:
	//
	// Failed to find node group[ng-3d5ce6454354****].
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ModifyNodeNumberPreCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeNumberPreCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberPreCheckResponseBodyData) SetAllow(v bool) *ModifyNodeNumberPreCheckResponseBodyData {
	s.Allow = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBodyData) SetReason(v string) *ModifyNodeNumberPreCheckResponseBodyData {
	s.Reason = &v
	return s
}

type ModifyNodeNumberPreCheckResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodeNumberPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodeNumberPreCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeNumberPreCheckResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberPreCheckResponse) SetHeaders(v map[string]*string) *ModifyNodeNumberPreCheckResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodeNumberPreCheckResponse) SetStatusCode(v int32) *ModifyNodeNumberPreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponse) SetBody(v *ModifyNodeNumberPreCheckResponseBody) *ModifyNodeNumberPreCheckResponse {
	s.Body = v
	return s
}

type QueryUpgradableVersionsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to query the minor versions that you can upgrade to. Default value: true. Valid values:
	//
	// 	- true: The minor versions that you can upgrade to.
	//
	// 	- false: The major versions that you can upgrade to.
	//
	// example:
	//
	// true
	Minor *bool `json:"Minor,omitempty" xml:"Minor,omitempty"`
}

func (s QueryUpgradableVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUpgradableVersionsRequest) GoString() string {
	return s.String()
}

func (s *QueryUpgradableVersionsRequest) SetInstanceId(v string) *QueryUpgradableVersionsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryUpgradableVersionsRequest) SetMinor(v bool) *QueryUpgradableVersionsRequest {
	s.Minor = &v
	return s
}

type QueryUpgradableVersionsResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The versions that you can upgrade to.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUpgradableVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUpgradableVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUpgradableVersionsResponseBody) SetAccessDeniedDetail(v string) *QueryUpgradableVersionsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) SetData(v []*string) *QueryUpgradableVersionsResponseBody {
	s.Data = v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) SetErrCode(v string) *QueryUpgradableVersionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) SetErrMessage(v string) *QueryUpgradableVersionsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) SetHttpStatusCode(v int32) *QueryUpgradableVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) SetRequestId(v string) *QueryUpgradableVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) SetSuccess(v bool) *QueryUpgradableVersionsResponseBody {
	s.Success = &v
	return s
}

type QueryUpgradableVersionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUpgradableVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUpgradableVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUpgradableVersionsResponse) GoString() string {
	return s.String()
}

func (s *QueryUpgradableVersionsResponse) SetHeaders(v map[string]*string) *QueryUpgradableVersionsResponse {
	s.Headers = v
	return s
}

func (s *QueryUpgradableVersionsResponse) SetStatusCode(v int32) *QueryUpgradableVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUpgradableVersionsResponse) SetBody(v *QueryUpgradableVersionsResponseBody) *QueryUpgradableVersionsResponse {
	s.Body = v
	return s
}

type ReleaseInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleaseInstanceResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBody) SetAccessDeniedDetail(v string) *ReleaseInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetData(v bool) *ReleaseInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetErrCode(v string) *ReleaseInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetErrMessage(v string) *ReleaseInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetHttpStatusCode(v int32) *ReleaseInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetRequestId(v string) *ReleaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetSuccess(v bool) *ReleaseInstanceResponseBody {
	s.Success = &v
	return s
}

type ReleaseInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponse) SetHeaders(v map[string]*string) *ReleaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstanceResponse) SetStatusCode(v int32) *ReleaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstanceResponse) SetBody(v *ReleaseInstanceResponseBody) *ReleaseInstanceResponse {
	s.Body = v
	return s
}

type RestartInstanceRequest struct {
	// Specifies whether to restart compute nodes in quick restart mode. Default value: false. Valid values:
	//
	// 	- true: Compute nodes are restarted in quick restart mode in multiple batches. The batches are executed in parallel, and the nodes in each batch are restarted at the same time.
	//
	// 	- false: Compute nodes are restarted in rolling restart mode.
	//
	// example:
	//
	// true
	FastMode *bool `json:"FastMode,omitempty" xml:"FastMode,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) SetFastMode(v bool) *RestartInstanceRequest {
	s.FastMode = &v
	return s
}

func (s *RestartInstanceRequest) SetInstanceId(v string) *RestartInstanceRequest {
	s.InstanceId = &v
	return s
}

type RestartInstanceResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unsupported.Non.Whitelist.Operation
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) SetAccessDeniedDetail(v string) *RestartInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RestartInstanceResponseBody) SetData(v bool) *RestartInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *RestartInstanceResponseBody) SetErrCode(v string) *RestartInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RestartInstanceResponseBody) SetErrMessage(v string) *RestartInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RestartInstanceResponseBody) SetHttpStatusCode(v int32) *RestartInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstanceResponseBody) SetSuccess(v bool) *RestartInstanceResponseBody {
	s.Success = &v
	return s
}

type RestartInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponse) SetHeaders(v map[string]*string) *RestartInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartInstanceResponse) SetStatusCode(v int32) *RestartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartInstanceResponse) SetBody(v *RestartInstanceResponseBody) *RestartInstanceResponse {
	s.Body = v
	return s
}

type UpdateInstanceNameRequest struct {
	// The new name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// new_name
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameRequest) SetClusterName(v string) *UpdateInstanceNameRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateInstanceNameRequest) SetInstanceId(v string) *UpdateInstanceNameRequest {
	s.InstanceId = &v
	return s
}

type UpdateInstanceNameResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponseBody) SetAccessDeniedDetail(v string) *UpdateInstanceNameResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetData(v bool) *UpdateInstanceNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetErrCode(v string) *UpdateInstanceNameResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetErrMessage(v string) *UpdateInstanceNameResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetRequestId(v string) *UpdateInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetSuccess(v bool) *UpdateInstanceNameResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponse) SetHeaders(v map[string]*string) *UpdateInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceNameResponse) SetStatusCode(v int32) *UpdateInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceNameResponse) SetBody(v *UpdateInstanceNameResponseBody) *UpdateInstanceNameResponse {
	s.Body = v
	return s
}

type UpgradeVersionRequest struct {
	FastMode *bool `json:"FastMode,omitempty" xml:"FastMode,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the minor version is upgraded. Default value: true. Valid values:
	//
	// 	- true: The minor version is upgraded.
	//
	// 	- false: The major version is upgraded.
	//
	// example:
	//
	// true
	Minor *bool `json:"Minor,omitempty" xml:"Minor,omitempty"`
	// The version to which you want to upgrade.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3.3.2-1.77-1.6.4
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
}

func (s UpgradeVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeVersionRequest) SetFastMode(v bool) *UpgradeVersionRequest {
	s.FastMode = &v
	return s
}

func (s *UpgradeVersionRequest) SetInstanceId(v string) *UpgradeVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeVersionRequest) SetMinor(v bool) *UpgradeVersionRequest {
	s.Minor = &v
	return s
}

func (s *UpgradeVersionRequest) SetTargetVersion(v string) *UpgradeVersionRequest {
	s.TargetVersion = &v
	return s
}

type UpgradeVersionResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid instance status: [cannot upgrade when instance is not running].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeVersionResponseBody) SetAccessDeniedDetail(v string) *UpgradeVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpgradeVersionResponseBody) SetData(v bool) *UpgradeVersionResponseBody {
	s.Data = &v
	return s
}

func (s *UpgradeVersionResponseBody) SetErrCode(v string) *UpgradeVersionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpgradeVersionResponseBody) SetErrMessage(v string) *UpgradeVersionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpgradeVersionResponseBody) SetHttpStatusCode(v int32) *UpgradeVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpgradeVersionResponseBody) SetRequestId(v string) *UpgradeVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeVersionResponseBody) SetSuccess(v bool) *UpgradeVersionResponseBody {
	s.Success = &v
	return s
}

type UpgradeVersionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeVersionResponse) SetHeaders(v map[string]*string) *UpgradeVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeVersionResponse) SetStatusCode(v int32) *UpgradeVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeVersionResponse) SetBody(v *UpgradeVersionResponseBody) *UpgradeVersionResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("starrocks"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the number of CUs for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/en/product/ecs?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.47c9281fkIZGiB#pricing) of EMR Serverless StarRocks instances.
//
// Before you call this operation, take note of the following items:
//
// 	- You can modify the number of CUs for a warehouse of only StarRocks instances of Standard Edition.
//
// 	- You can increase the number of disks only for warehouses of the standard specifications.
//
// 	- The instance must be in the Running state.
//
// After you modify the number of CUs for a warehouse, the billing of CUs has the following changes:
//
// 	- Pay-as-you-go StarRocks instances: You are charged based on the number of CUs.
//
// 	- Subscription StarRocks instances: You are charged additionally based on the price difference between the number of CUs before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyCuRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCuResponse
func (client *Client) ModifyCuWithOptions(request *ModifyCuRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyCuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FastMode)) {
		query["FastMode"] = request.FastMode
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyCu"),
		Version:     tea.String("2022-10-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/resourceChange/modifyCu"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyCuResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyCuResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the number of CUs for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/en/product/ecs?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.47c9281fkIZGiB#pricing) of EMR Serverless StarRocks instances.
//
// Before you call this operation, take note of the following items:
//
// 	- You can modify the number of CUs for a warehouse of only StarRocks instances of Standard Edition.
//
// 	- You can increase the number of disks only for warehouses of the standard specifications.
//
// 	- The instance must be in the Running state.
//
// After you modify the number of CUs for a warehouse, the billing of CUs has the following changes:
//
// 	- Pay-as-you-go StarRocks instances: You are charged based on the number of CUs.
//
// 	- Subscription StarRocks instances: You are charged additionally based on the price difference between the number of CUs before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyCuRequest
//
// @return ModifyCuResponse
func (client *Client) ModifyCu(request *ModifyCuRequest) (_result *ModifyCuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyCuResponse{}
	_body, _err := client.ModifyCuWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a precheck before you modify the number of CUs for a warehouse.
//
// @param request - ModifyCuPreCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCuPreCheckResponse
func (client *Client) ModifyCuPreCheckWithOptions(request *ModifyCuPreCheckRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyCuPreCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyCuPreCheck"),
		Version:     tea.String("2022-10-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/resourceChange/modifyCuPreCheck"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyCuPreCheckResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyCuPreCheckResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Performs a precheck before you modify the number of CUs for a warehouse.
//
// @param request - ModifyCuPreCheckRequest
//
// @return ModifyCuPreCheckResponse
func (client *Client) ModifyCuPreCheck(request *ModifyCuPreCheckRequest) (_result *ModifyCuPreCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyCuPreCheckResponse{}
	_body, _err := client.ModifyCuPreCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Increases the number of disks for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/en/product/ecs?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.47c9281fkIZGiB#pricing) of EMR Serverless StarRocks instances. Before you call this operation, take note of the following items:
//
// 	- You can increase the number of disks only for StarRocks instances of Standard Edition.
//
// 	- You can increase the number of disks only for warehouses of the standard specifications.
//
// 	- The instance must be in the Running state.
//
// After you increase the number of disks for a warehouse, the billing of disks has the following changes:
//
// 	- Pay-as-you-go StarRocks instances: You are charged for the disk based on the new disk type.
//
// 	- Subscription StarRocks instances: You are charged additionally based on the price difference between the number of disks before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyDiskNumberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskNumberResponse
func (client *Client) ModifyDiskNumberWithOptions(request *ModifyDiskNumberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyDiskNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FastMode)) {
		query["FastMode"] = request.FastMode
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDiskNumber"),
		Version:     tea.String("2022-10-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/resourceChange/modifyDiskNumber"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyDiskNumberResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyDiskNumberResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Increases the number of disks for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/en/product/ecs?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.47c9281fkIZGiB#pricing) of EMR Serverless StarRocks instances. Before you call this operation, take note of the following items:
//
// 	- You can increase the number of disks only for StarRocks instances of Standard Edition.
//
// 	- You can increase the number of disks only for warehouses of the standard specifications.
//
// 	- The instance must be in the Running state.
//
// After you increase the number of disks for a warehouse, the billing of disks has the following changes:
//
// 	- Pay-as-you-go StarRocks instances: You are charged for the disk based on the new disk type.
//
// 	- Subscription StarRocks instances: You are charged additionally based on the price difference between the number of disks before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyDiskNumberRequest
//
// @return ModifyDiskNumberResponse
func (client *Client) ModifyDiskNumber(request *ModifyDiskNumberRequest) (_result *ModifyDiskNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDiskNumberResponse{}
	_body, _err := client.ModifyDiskNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the disk performance level for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/en/product/ecs?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.47c9281fkIZGiB#pricing) of EMR Serverless StarRocks instances.
//
// Before you call this operation, take note of the following items:
//
// 	- You can modify the disk performance level only for StarRocks instances of Standard Edition.
//
// 	- You can modify the disk performance level only for warehouses of the standard specifications.
//
// 	- The instance must be in the Running state.
//
// 	- You cannot downgrade the performance level to PL0.
//
// 	- The performance level of an Enterprise SSD (ESSD) is limited by the ESSD disk size. If you cannot upgrade the performance level of an ESSD, expand the ESSD and try again. For more information, see [Enterprise SSDs](https://help.aliyun.com/document_detail/122389.html).
//
// After the disk performance level is changed, the billing of the disk has the following changes:
//
// 	- Pay-as-you-go StarRocks instances: You are charged for the disk based on the new disk type.
//
// 	- Subscription StarRocks instances: You are charged additionally based on the price difference between the disk performance level before and sfter the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyDiskPerformanceLevelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskPerformanceLevelResponse
func (client *Client) ModifyDiskPerformanceLevelWithOptions(request *ModifyDiskPerformanceLevelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyDiskPerformanceLevelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDiskPerformanceLevel"),
		Version:     tea.String("2022-10-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/resourceChange/modifyDiskPerformanceLevel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyDiskPerformanceLevelResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyDiskPerformanceLevelResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the disk performance level for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/en/product/ecs?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.47c9281fkIZGiB#pricing) of EMR Serverless StarRocks instances.
//
// Before you call this operation, take note of the following items:
//
// 	- You can modify the disk performance level only for StarRocks instances of Standard Edition.
//
// 	- You can modify the disk performance level only for warehouses of the standard specifications.
//
// 	- The instance must be in the Running state.
//
// 	- You cannot downgrade the performance level to PL0.
//
// 	- The performance level of an Enterprise SSD (ESSD) is limited by the ESSD disk size. If you cannot upgrade the performance level of an ESSD, expand the ESSD and try again. For more information, see [Enterprise SSDs](https://help.aliyun.com/document_detail/122389.html).
//
// After the disk performance level is changed, the billing of the disk has the following changes:
//
// 	- Pay-as-you-go StarRocks instances: You are charged for the disk based on the new disk type.
//
// 	- Subscription StarRocks instances: You are charged additionally based on the price difference between the disk performance level before and sfter the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyDiskPerformanceLevelRequest
//
// @return ModifyDiskPerformanceLevelResponse
func (client *Client) ModifyDiskPerformanceLevel(request *ModifyDiskPerformanceLevelRequest) (_result *ModifyDiskPerformanceLevelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDiskPerformanceLevelResponse{}
	_body, _err := client.ModifyDiskPerformanceLevelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Expands the disk size for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/en/product/ecs?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.47c9281fkIZGiB#pricing) of EMR Serverless StarRocks instances. Before you call this operation, take note of the following items:
//
// 	- You can expand the disk size only for StarRocks instances of Standard Edition.
//
// 	- You can expand the disk size only for warehouses of the standard specifications.
//
// 	- The instance must be in the Running state.
//
// After you expand the disk size, the billing of disks has the following changes:
//
// 	- Pay-as-you-go StarRocks instances: You are charged for the disk based on the new disk size.
//
// 	- Subscription StarRocks instances: You are charged additionally based on the price difference between the disk size before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyDiskSizeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskSizeResponse
func (client *Client) ModifyDiskSizeWithOptions(request *ModifyDiskSizeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyDiskSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDiskSize"),
		Version:     tea.String("2022-10-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/resourceChange/modifyDiskSize"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyDiskSizeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyDiskSizeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Expands the disk size for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/en/product/ecs?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.47c9281fkIZGiB#pricing) of EMR Serverless StarRocks instances. Before you call this operation, take note of the following items:
//
// 	- You can expand the disk size only for StarRocks instances of Standard Edition.
//
// 	- You can expand the disk size only for warehouses of the standard specifications.
//
// 	- The instance must be in the Running state.
//
// After you expand the disk size, the billing of disks has the following changes:
//
// 	- Pay-as-you-go StarRocks instances: You are charged for the disk based on the new disk size.
//
// 	- Subscription StarRocks instances: You are charged additionally based on the price difference between the disk size before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyDiskSizeRequest
//
// @return ModifyDiskSizeResponse
func (client *Client) ModifyDiskSize(request *ModifyDiskSizeRequest) (_result *ModifyDiskSizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDiskSizeResponse{}
	_body, _err := client.ModifyDiskSizeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the number of nodes in a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/en/product/ecs?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.47c9281fkIZGiB#pricing) of EMR Serverless StarRocks instances.
//
// Before you call this operation, take note of the following items:
//
// 	- You can modify the number of nodes in a warehouse of only StarRocks instances of Standard Edition.
//
// 	- The instance must be in the Running state.
//
// 	- The number of frontend nodes (FEs) cannot be an even number, and you cannot reduce the number of FE nodes.
//
// After you modify the number of nodes in a warehouse, the billing of nodes has the following changes:
//
// 	- Pay-as-you-go StarRocks instances: You are charged based on the number of nodes.
//
// 	- Subscription StarRocks instances: You are charged additionally based on the price difference between the number of nodes before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyNodeNumberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodeNumberResponse
func (client *Client) ModifyNodeNumberWithOptions(request *ModifyNodeNumberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyNodeNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNodeNumber"),
		Version:     tea.String("2022-10-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/resourceChange/modifyNodeNumber"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyNodeNumberResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyNodeNumberResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the number of nodes in a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/en/product/ecs?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.47c9281fkIZGiB#pricing) of EMR Serverless StarRocks instances.
//
// Before you call this operation, take note of the following items:
//
// 	- You can modify the number of nodes in a warehouse of only StarRocks instances of Standard Edition.
//
// 	- The instance must be in the Running state.
//
// 	- The number of frontend nodes (FEs) cannot be an even number, and you cannot reduce the number of FE nodes.
//
// After you modify the number of nodes in a warehouse, the billing of nodes has the following changes:
//
// 	- Pay-as-you-go StarRocks instances: You are charged based on the number of nodes.
//
// 	- Subscription StarRocks instances: You are charged additionally based on the price difference between the number of nodes before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyNodeNumberRequest
//
// @return ModifyNodeNumberResponse
func (client *Client) ModifyNodeNumber(request *ModifyNodeNumberRequest) (_result *ModifyNodeNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyNodeNumberResponse{}
	_body, _err := client.ModifyNodeNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a precheck before you modify the number of nodes in a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// @param request - ModifyNodeNumberPreCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodeNumberPreCheckResponse
func (client *Client) ModifyNodeNumberPreCheckWithOptions(request *ModifyNodeNumberPreCheckRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyNodeNumberPreCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNodeNumberPreCheck"),
		Version:     tea.String("2022-10-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/resourceChange/modifyNodeNumberPreCheck"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyNodeNumberPreCheckResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyNodeNumberPreCheckResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Performs a precheck before you modify the number of nodes in a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// @param request - ModifyNodeNumberPreCheckRequest
//
// @return ModifyNodeNumberPreCheckResponse
func (client *Client) ModifyNodeNumberPreCheck(request *ModifyNodeNumberPreCheckRequest) (_result *ModifyNodeNumberPreCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyNodeNumberPreCheckResponse{}
	_body, _err := client.ModifyNodeNumberPreCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the versions of an E-MapReduce (EMR) Serverless StarRocks instance that the versions that you can upgrade to. The versions of a StarRocks instance include the major version and minor version. You can view the major version and minor version of a StarRocks instance in the Version Information section of the Instance Details tab in the EMR console. You can call this operation to query the minor versions or major versions that the versions that you can upgrade to.
//
// @param request - QueryUpgradableVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUpgradableVersionsResponse
func (client *Client) QueryUpgradableVersionsWithOptions(request *QueryUpgradableVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryUpgradableVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Minor)) {
		query["Minor"] = request.Minor
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUpgradableVersions"),
		Version:     tea.String("2022-10-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/starrocks/queryUpgradableVersions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryUpgradableVersionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryUpgradableVersionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the versions of an E-MapReduce (EMR) Serverless StarRocks instance that the versions that you can upgrade to. The versions of a StarRocks instance include the major version and minor version. You can view the major version and minor version of a StarRocks instance in the Version Information section of the Instance Details tab in the EMR console. You can call this operation to query the minor versions or major versions that the versions that you can upgrade to.
//
// @param request - QueryUpgradableVersionsRequest
//
// @return QueryUpgradableVersionsResponse
func (client *Client) QueryUpgradableVersions(request *QueryUpgradableVersionsRequest) (_result *QueryUpgradableVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryUpgradableVersionsResponse{}
	_body, _err := client.QueryUpgradableVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go E-MapReduce (EMR) Serverless StarRocks instance. To unsubscribe from a subscription instance, go to the Unsubscribe page of the Expenses and Costs console.
//
// Description:
//
// *
//
// **Warning:*	- After an instance is released, all physical resources used by the instance are recycled. Relevant data is erased and cannot be restored.
//
// @param request - ReleaseInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseInstance"),
		Version:     tea.String("2022-10-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cluster/release"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReleaseInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReleaseInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Releases a pay-as-you-go E-MapReduce (EMR) Serverless StarRocks instance. To unsubscribe from a subscription instance, go to the Unsubscribe page of the Expenses and Costs console.
//
// Description:
//
// *
//
// **Warning:*	- After an instance is released, all physical resources used by the instance are recycled. Relevant data is erased and cannot be restored.
//
// @param request - ReleaseInstanceRequest
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// This operation is an asynchronous operation. After you call this operation to restart a StarRocks instance, the operation sets the status of the instance to Restarting and begins the restart process. When the status of the instance changes to Running, the instance is restarted.
//
// @param request - RestartInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstanceResponse
func (client *Client) RestartInstanceWithOptions(request *RestartInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FastMode)) {
		query["FastMode"] = request.FastMode
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartInstance"),
		Version:     tea.String("2022-10-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/starrocks/restartCluster"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RestartInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RestartInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Restarts an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// This operation is an asynchronous operation. After you call this operation to restart a StarRocks instance, the operation sets the status of the instance to Restarting and begins the restart process. When the status of the instance changes to Running, the instance is restarted.
//
// @param request - RestartInstanceRequest
//
// @return RestartInstanceResponse
func (client *Client) RestartInstance(request *RestartInstanceRequest) (_result *RestartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// @param request - UpdateInstanceNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceNameWithOptions(request *UpdateInstanceNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceName"),
		Version:     tea.String("2022-10-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cluster/update_name"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateInstanceNameResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateInstanceNameResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the name of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// @param request - UpdateInstanceNameRequest
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceName(request *UpdateInstanceNameRequest) (_result *UpdateInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.UpdateInstanceNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the version of an E-MapReduce (EMR) Serverless StarRocks instance. The versions of a StarRocks instance include the major version and minor version. You can view the major version and minor version of a StarRocks instance in the Version Information section of the Instance Details tab in the EMR console. This operation can be used to upgrade the minor version or major version of a StarRocks instance. You can call the QueryUpgradableVersions operation to query the versions that you can upgrade to.
//
// Description:
//
// The instance must be in the Running state when you call this operation.
//
// @param request - UpgradeVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeVersionResponse
func (client *Client) UpgradeVersionWithOptions(request *UpgradeVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpgradeVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FastMode)) {
		query["FastMode"] = request.FastMode
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Minor)) {
		query["Minor"] = request.Minor
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVersion)) {
		query["TargetVersion"] = request.TargetVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeVersion"),
		Version:     tea.String("2022-10-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/starrocks/upgradeVersion"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpgradeVersionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpgradeVersionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Upgrades the version of an E-MapReduce (EMR) Serverless StarRocks instance. The versions of a StarRocks instance include the major version and minor version. You can view the major version and minor version of a StarRocks instance in the Version Information section of the Instance Details tab in the EMR console. This operation can be used to upgrade the minor version or major version of a StarRocks instance. You can call the QueryUpgradableVersions operation to query the versions that you can upgrade to.
//
// Description:
//
// The instance must be in the Running state when you call this operation.
//
// @param request - UpgradeVersionRequest
//
// @return UpgradeVersionResponse
func (client *Client) UpgradeVersion(request *UpgradeVersionRequest) (_result *UpgradeVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeVersionResponse{}
	_body, _err := client.UpgradeVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
