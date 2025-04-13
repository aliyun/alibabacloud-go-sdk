// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetMqttConnectRequest struct {
	Request map[string]interface{} `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s GetMqttConnectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMqttConnectRequest) GoString() string {
	return s.String()
}

func (s *GetMqttConnectRequest) SetRequest(v map[string]interface{}) *GetMqttConnectRequest {
	s.Request = v
	return s
}

type GetMqttConnectResponseBody struct {
	AccessDeniedDetail *GetMqttConnectResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *GetMqttConnectResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode     *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMqttConnectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMqttConnectResponseBody) GoString() string {
	return s.String()
}

func (s *GetMqttConnectResponseBody) SetAccessDeniedDetail(v *GetMqttConnectResponseBodyAccessDeniedDetail) *GetMqttConnectResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetMqttConnectResponseBody) SetCode(v string) *GetMqttConnectResponseBody {
	s.Code = &v
	return s
}

func (s *GetMqttConnectResponseBody) SetData(v *GetMqttConnectResponseBodyData) *GetMqttConnectResponseBody {
	s.Data = v
	return s
}

func (s *GetMqttConnectResponseBody) SetHttpStatusCode(v int32) *GetMqttConnectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMqttConnectResponseBody) SetMessage(v string) *GetMqttConnectResponseBody {
	s.Message = &v
	return s
}

func (s *GetMqttConnectResponseBody) SetRequestId(v string) *GetMqttConnectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMqttConnectResponseBody) SetSuccess(v bool) *GetMqttConnectResponseBody {
	s.Success = &v
	return s
}

type GetMqttConnectResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetMqttConnectResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s GetMqttConnectResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetMqttConnectResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetMqttConnectResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetMqttConnectResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetMqttConnectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetMqttConnectResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetMqttConnectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetMqttConnectResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetMqttConnectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetMqttConnectResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetMqttConnectResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetMqttConnectResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetMqttConnectResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetMqttConnectResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetMqttConnectResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type GetMqttConnectResponseBodyData struct {
	AccessKey  *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QosLevel   *int32  `json:"QosLevel,omitempty" xml:"QosLevel,omitempty"`
	ServerUri  *string `json:"ServerUri,omitempty" xml:"ServerUri,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetMqttConnectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMqttConnectResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMqttConnectResponseBodyData) SetAccessKey(v string) *GetMqttConnectResponseBodyData {
	s.AccessKey = &v
	return s
}

func (s *GetMqttConnectResponseBodyData) SetClientId(v string) *GetMqttConnectResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *GetMqttConnectResponseBodyData) SetExpireTime(v int64) *GetMqttConnectResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetMqttConnectResponseBodyData) SetInstanceId(v string) *GetMqttConnectResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetMqttConnectResponseBodyData) SetQosLevel(v int32) *GetMqttConnectResponseBodyData {
	s.QosLevel = &v
	return s
}

func (s *GetMqttConnectResponseBodyData) SetServerUri(v string) *GetMqttConnectResponseBodyData {
	s.ServerUri = &v
	return s
}

func (s *GetMqttConnectResponseBodyData) SetToken(v string) *GetMqttConnectResponseBodyData {
	s.Token = &v
	return s
}

type GetMqttConnectResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMqttConnectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMqttConnectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMqttConnectResponse) GoString() string {
	return s.String()
}

func (s *GetMqttConnectResponse) SetHeaders(v map[string]*string) *GetMqttConnectResponse {
	s.Headers = v
	return s
}

func (s *GetMqttConnectResponse) SetStatusCode(v int32) *GetMqttConnectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMqttConnectResponse) SetBody(v *GetMqttConnectResponseBody) *GetMqttConnectResponse {
	s.Body = v
	return s
}

type GetNonceRequest struct {
	Request map[string]interface{} `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s GetNonceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNonceRequest) GoString() string {
	return s.String()
}

func (s *GetNonceRequest) SetRequest(v map[string]interface{}) *GetNonceRequest {
	s.Request = v
	return s
}

type GetNonceResponseBody struct {
	AccessDeniedDetail *GetNonceResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *GetNonceResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode     *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNonceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNonceResponseBody) GoString() string {
	return s.String()
}

func (s *GetNonceResponseBody) SetAccessDeniedDetail(v *GetNonceResponseBodyAccessDeniedDetail) *GetNonceResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetNonceResponseBody) SetCode(v string) *GetNonceResponseBody {
	s.Code = &v
	return s
}

func (s *GetNonceResponseBody) SetData(v *GetNonceResponseBodyData) *GetNonceResponseBody {
	s.Data = v
	return s
}

func (s *GetNonceResponseBody) SetHttpStatusCode(v int32) *GetNonceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNonceResponseBody) SetMessage(v string) *GetNonceResponseBody {
	s.Message = &v
	return s
}

func (s *GetNonceResponseBody) SetRequestId(v string) *GetNonceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNonceResponseBody) SetSuccess(v bool) *GetNonceResponseBody {
	s.Success = &v
	return s
}

type GetNonceResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetNonceResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s GetNonceResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetNonceResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetNonceResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetNonceResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetNonceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetNonceResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetNonceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetNonceResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetNonceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetNonceResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetNonceResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetNonceResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetNonceResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetNonceResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetNonceResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type GetNonceResponseBodyData struct {
	ExpiresIn *int64  `json:"ExpiresIn,omitempty" xml:"ExpiresIn,omitempty"`
	Nonce     *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
}

func (s GetNonceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNonceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNonceResponseBodyData) SetExpiresIn(v int64) *GetNonceResponseBodyData {
	s.ExpiresIn = &v
	return s
}

func (s *GetNonceResponseBodyData) SetNonce(v string) *GetNonceResponseBodyData {
	s.Nonce = &v
	return s
}

type GetNonceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNonceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNonceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNonceResponse) GoString() string {
	return s.String()
}

func (s *GetNonceResponse) SetHeaders(v map[string]*string) *GetNonceResponse {
	s.Headers = v
	return s
}

func (s *GetNonceResponse) SetStatusCode(v int32) *GetNonceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNonceResponse) SetBody(v *GetNonceResponseBody) *GetNonceResponse {
	s.Body = v
	return s
}

type ListMeasurePointListByNodeCodePageRequest struct {
	Request map[string]interface{} `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s ListMeasurePointListByNodeCodePageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMeasurePointListByNodeCodePageRequest) GoString() string {
	return s.String()
}

func (s *ListMeasurePointListByNodeCodePageRequest) SetRequest(v map[string]interface{}) *ListMeasurePointListByNodeCodePageRequest {
	s.Request = v
	return s
}

type ListMeasurePointListByNodeCodePageResponseBody struct {
	AccessDeniedDetail *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *string                                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *ListMeasurePointListByNodeCodePageResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode     *int32                                                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMeasurePointListByNodeCodePageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMeasurePointListByNodeCodePageResponseBody) GoString() string {
	return s.String()
}

func (s *ListMeasurePointListByNodeCodePageResponseBody) SetAccessDeniedDetail(v *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail) *ListMeasurePointListByNodeCodePageResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBody) SetCode(v string) *ListMeasurePointListByNodeCodePageResponseBody {
	s.Code = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBody) SetData(v *ListMeasurePointListByNodeCodePageResponseBodyData) *ListMeasurePointListByNodeCodePageResponseBody {
	s.Data = v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBody) SetHttpStatusCode(v int32) *ListMeasurePointListByNodeCodePageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBody) SetMessage(v string) *ListMeasurePointListByNodeCodePageResponseBody {
	s.Message = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBody) SetRequestId(v string) *ListMeasurePointListByNodeCodePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBody) SetSuccess(v bool) *ListMeasurePointListByNodeCodePageResponseBody {
	s.Success = &v
	return s
}

type ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ListMeasurePointListByNodeCodePageResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type ListMeasurePointListByNodeCodePageResponseBodyData struct {
	Count       *int32                                                        `json:"Count,omitempty" xml:"Count,omitempty"`
	CurrentPage *int32                                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DataList    []*ListMeasurePointListByNodeCodePageResponseBodyDataDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	TotalPage   *int32                                                        `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListMeasurePointListByNodeCodePageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMeasurePointListByNodeCodePageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyData) SetCount(v int32) *ListMeasurePointListByNodeCodePageResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyData) SetCurrentPage(v int32) *ListMeasurePointListByNodeCodePageResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyData) SetDataList(v []*ListMeasurePointListByNodeCodePageResponseBodyDataDataList) *ListMeasurePointListByNodeCodePageResponseBodyData {
	s.DataList = v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyData) SetTotalPage(v int32) *ListMeasurePointListByNodeCodePageResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListMeasurePointListByNodeCodePageResponseBodyDataDataList struct {
	AcqFreq           *string `json:"AcqFreq,omitempty" xml:"AcqFreq,omitempty"`
	AlarmAttribute    *string `json:"AlarmAttribute,omitempty" xml:"AlarmAttribute,omitempty"`
	AlarmLevel        *string `json:"AlarmLevel,omitempty" xml:"AlarmLevel,omitempty"`
	CalcAttribute     *string `json:"CalcAttribute,omitempty" xml:"CalcAttribute,omitempty"`
	CalcFormula       *string `json:"CalcFormula,omitempty" xml:"CalcFormula,omitempty"`
	CalcType          *string `json:"CalcType,omitempty" xml:"CalcType,omitempty"`
	Code              *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateUser        *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	DataSourceId      *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	DataType          *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrMsgDetail      *string `json:"ErrMsgDetail,omitempty" xml:"ErrMsgDetail,omitempty"`
	FullName          *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	GmtCreate         *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModify         *int64  `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	Id                *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	InitValue         *string `json:"InitValue,omitempty" xml:"InitValue,omitempty"`
	IsEnable          *string `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	LastUploadTime    *string `json:"LastUploadTime,omitempty" xml:"LastUploadTime,omitempty"`
	LatestValue       *string `json:"LatestValue,omitempty" xml:"LatestValue,omitempty"`
	LowerLimit        *string `json:"LowerLimit,omitempty" xml:"LowerLimit,omitempty"`
	MeasurePointState *string `json:"MeasurePointState,omitempty" xml:"MeasurePointState,omitempty"`
	ModifyUser        *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeCode          *string `json:"NodeCode,omitempty" xml:"NodeCode,omitempty"`
	NodeId            *int32  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName          *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	ParentFullName    *string `json:"ParentFullName,omitempty" xml:"ParentFullName,omitempty"`
	PayLoad           *string `json:"PayLoad,omitempty" xml:"PayLoad,omitempty"`
	ProtocolConfig    *string `json:"ProtocolConfig,omitempty" xml:"ProtocolConfig,omitempty"`
	SourcePoint       *string `json:"SourcePoint,omitempty" xml:"SourcePoint,omitempty"`
	TenantCode        *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	Time              *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Unit              *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	UpperLimit        *string `json:"UpperLimit,omitempty" xml:"UpperLimit,omitempty"`
}

func (s ListMeasurePointListByNodeCodePageResponseBodyDataDataList) String() string {
	return tea.Prettify(s)
}

func (s ListMeasurePointListByNodeCodePageResponseBodyDataDataList) GoString() string {
	return s.String()
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetAcqFreq(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.AcqFreq = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetAlarmAttribute(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.AlarmAttribute = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetAlarmLevel(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.AlarmLevel = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetCalcAttribute(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.CalcAttribute = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetCalcFormula(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.CalcFormula = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetCalcType(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.CalcType = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetCode(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.Code = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetCreateUser(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.CreateUser = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetDataSourceId(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.DataSourceId = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetDataType(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.DataType = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetDescription(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.Description = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetErrMsgDetail(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.ErrMsgDetail = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetFullName(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.FullName = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetGmtCreate(v int64) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.GmtCreate = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetGmtModify(v int64) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.GmtModify = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetId(v int32) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.Id = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetInitValue(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.InitValue = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetIsEnable(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.IsEnable = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetLastUploadTime(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.LastUploadTime = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetLatestValue(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.LatestValue = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetLowerLimit(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.LowerLimit = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetMeasurePointState(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.MeasurePointState = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetModifyUser(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.ModifyUser = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetName(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.Name = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetNodeCode(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.NodeCode = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetNodeId(v int32) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.NodeId = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetNodeName(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.NodeName = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetParentFullName(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.ParentFullName = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetPayLoad(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.PayLoad = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetProtocolConfig(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.ProtocolConfig = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetSourcePoint(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.SourcePoint = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetTenantCode(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.TenantCode = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetTime(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.Time = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetType(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.Type = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetUnit(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.Unit = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponseBodyDataDataList) SetUpperLimit(v string) *ListMeasurePointListByNodeCodePageResponseBodyDataDataList {
	s.UpperLimit = &v
	return s
}

type ListMeasurePointListByNodeCodePageResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMeasurePointListByNodeCodePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMeasurePointListByNodeCodePageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMeasurePointListByNodeCodePageResponse) GoString() string {
	return s.String()
}

func (s *ListMeasurePointListByNodeCodePageResponse) SetHeaders(v map[string]*string) *ListMeasurePointListByNodeCodePageResponse {
	s.Headers = v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponse) SetStatusCode(v int32) *ListMeasurePointListByNodeCodePageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMeasurePointListByNodeCodePageResponse) SetBody(v *ListMeasurePointListByNodeCodePageResponseBody) *ListMeasurePointListByNodeCodePageResponse {
	s.Body = v
	return s
}

type MultiFieldBatchUploadRequest struct {
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultiFieldBatchUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s MultiFieldBatchUploadRequest) GoString() string {
	return s.String()
}

func (s *MultiFieldBatchUploadRequest) SetBody(v map[string]interface{}) *MultiFieldBatchUploadRequest {
	s.Body = v
	return s
}

type MultiFieldBatchUploadResponseBody struct {
	AccessDeniedDetail *MultiFieldBatchUploadResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *MultiFieldBatchUploadResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode     *int32                                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MultiFieldBatchUploadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MultiFieldBatchUploadResponseBody) GoString() string {
	return s.String()
}

func (s *MultiFieldBatchUploadResponseBody) SetAccessDeniedDetail(v *MultiFieldBatchUploadResponseBodyAccessDeniedDetail) *MultiFieldBatchUploadResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *MultiFieldBatchUploadResponseBody) SetCode(v string) *MultiFieldBatchUploadResponseBody {
	s.Code = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBody) SetData(v *MultiFieldBatchUploadResponseBodyData) *MultiFieldBatchUploadResponseBody {
	s.Data = v
	return s
}

func (s *MultiFieldBatchUploadResponseBody) SetHttpStatusCode(v int32) *MultiFieldBatchUploadResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBody) SetMessage(v string) *MultiFieldBatchUploadResponseBody {
	s.Message = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBody) SetRequestId(v string) *MultiFieldBatchUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBody) SetSuccess(v bool) *MultiFieldBatchUploadResponseBody {
	s.Success = &v
	return s
}

type MultiFieldBatchUploadResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s MultiFieldBatchUploadResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s MultiFieldBatchUploadResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *MultiFieldBatchUploadResponseBodyAccessDeniedDetail) SetAuthAction(v string) *MultiFieldBatchUploadResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *MultiFieldBatchUploadResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *MultiFieldBatchUploadResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *MultiFieldBatchUploadResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *MultiFieldBatchUploadResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *MultiFieldBatchUploadResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBodyAccessDeniedDetail) SetPolicyType(v string) *MultiFieldBatchUploadResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type MultiFieldBatchUploadResponseBodyData struct {
	Count   *int64                                          `json:"Count,omitempty" xml:"Count,omitempty"`
	Message []*MultiFieldBatchUploadResponseBodyDataMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s MultiFieldBatchUploadResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MultiFieldBatchUploadResponseBodyData) GoString() string {
	return s.String()
}

func (s *MultiFieldBatchUploadResponseBodyData) SetCount(v int64) *MultiFieldBatchUploadResponseBodyData {
	s.Count = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBodyData) SetMessage(v []*MultiFieldBatchUploadResponseBodyDataMessage) *MultiFieldBatchUploadResponseBodyData {
	s.Message = v
	return s
}

type MultiFieldBatchUploadResponseBodyDataMessage struct {
	ErrorMsg     *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	MeasurePoint *string `json:"MeasurePoint,omitempty" xml:"MeasurePoint,omitempty"`
	Node         *string `json:"Node,omitempty" xml:"Node,omitempty"`
}

func (s MultiFieldBatchUploadResponseBodyDataMessage) String() string {
	return tea.Prettify(s)
}

func (s MultiFieldBatchUploadResponseBodyDataMessage) GoString() string {
	return s.String()
}

func (s *MultiFieldBatchUploadResponseBodyDataMessage) SetErrorMsg(v string) *MultiFieldBatchUploadResponseBodyDataMessage {
	s.ErrorMsg = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBodyDataMessage) SetMeasurePoint(v string) *MultiFieldBatchUploadResponseBodyDataMessage {
	s.MeasurePoint = &v
	return s
}

func (s *MultiFieldBatchUploadResponseBodyDataMessage) SetNode(v string) *MultiFieldBatchUploadResponseBodyDataMessage {
	s.Node = &v
	return s
}

type MultiFieldBatchUploadResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MultiFieldBatchUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultiFieldBatchUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s MultiFieldBatchUploadResponse) GoString() string {
	return s.String()
}

func (s *MultiFieldBatchUploadResponse) SetHeaders(v map[string]*string) *MultiFieldBatchUploadResponse {
	s.Headers = v
	return s
}

func (s *MultiFieldBatchUploadResponse) SetStatusCode(v int32) *MultiFieldBatchUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *MultiFieldBatchUploadResponse) SetBody(v *MultiFieldBatchUploadResponseBody) *MultiFieldBatchUploadResponse {
	s.Body = v
	return s
}

type MultiSourcePointBatchUploadRequest struct {
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultiSourcePointBatchUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s MultiSourcePointBatchUploadRequest) GoString() string {
	return s.String()
}

func (s *MultiSourcePointBatchUploadRequest) SetBody(v map[string]interface{}) *MultiSourcePointBatchUploadRequest {
	s.Body = v
	return s
}

type MultiSourcePointBatchUploadResponseBody struct {
	AccessDeniedDetail *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *MultiSourcePointBatchUploadResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode     *int32                                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MultiSourcePointBatchUploadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MultiSourcePointBatchUploadResponseBody) GoString() string {
	return s.String()
}

func (s *MultiSourcePointBatchUploadResponseBody) SetAccessDeniedDetail(v *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail) *MultiSourcePointBatchUploadResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBody) SetCode(v string) *MultiSourcePointBatchUploadResponseBody {
	s.Code = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBody) SetData(v *MultiSourcePointBatchUploadResponseBodyData) *MultiSourcePointBatchUploadResponseBody {
	s.Data = v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBody) SetHttpStatusCode(v int32) *MultiSourcePointBatchUploadResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBody) SetMessage(v string) *MultiSourcePointBatchUploadResponseBody {
	s.Message = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBody) SetRequestId(v string) *MultiSourcePointBatchUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBody) SetSuccess(v bool) *MultiSourcePointBatchUploadResponseBody {
	s.Success = &v
	return s
}

type MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail) SetAuthAction(v string) *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail) SetPolicyType(v string) *MultiSourcePointBatchUploadResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type MultiSourcePointBatchUploadResponseBodyData struct {
	Count   *int64                                                `json:"Count,omitempty" xml:"Count,omitempty"`
	Message []*MultiSourcePointBatchUploadResponseBodyDataMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s MultiSourcePointBatchUploadResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MultiSourcePointBatchUploadResponseBodyData) GoString() string {
	return s.String()
}

func (s *MultiSourcePointBatchUploadResponseBodyData) SetCount(v int64) *MultiSourcePointBatchUploadResponseBodyData {
	s.Count = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBodyData) SetMessage(v []*MultiSourcePointBatchUploadResponseBodyDataMessage) *MultiSourcePointBatchUploadResponseBodyData {
	s.Message = v
	return s
}

type MultiSourcePointBatchUploadResponseBodyDataMessage struct {
	ErrorMsg    *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Node        *string `json:"Node,omitempty" xml:"Node,omitempty"`
	SourcePoint *string `json:"SourcePoint,omitempty" xml:"SourcePoint,omitempty"`
}

func (s MultiSourcePointBatchUploadResponseBodyDataMessage) String() string {
	return tea.Prettify(s)
}

func (s MultiSourcePointBatchUploadResponseBodyDataMessage) GoString() string {
	return s.String()
}

func (s *MultiSourcePointBatchUploadResponseBodyDataMessage) SetErrorMsg(v string) *MultiSourcePointBatchUploadResponseBodyDataMessage {
	s.ErrorMsg = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBodyDataMessage) SetNode(v string) *MultiSourcePointBatchUploadResponseBodyDataMessage {
	s.Node = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponseBodyDataMessage) SetSourcePoint(v string) *MultiSourcePointBatchUploadResponseBodyDataMessage {
	s.SourcePoint = &v
	return s
}

type MultiSourcePointBatchUploadResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MultiSourcePointBatchUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultiSourcePointBatchUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s MultiSourcePointBatchUploadResponse) GoString() string {
	return s.String()
}

func (s *MultiSourcePointBatchUploadResponse) SetHeaders(v map[string]*string) *MultiSourcePointBatchUploadResponse {
	s.Headers = v
	return s
}

func (s *MultiSourcePointBatchUploadResponse) SetStatusCode(v int32) *MultiSourcePointBatchUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *MultiSourcePointBatchUploadResponse) SetBody(v *MultiSourcePointBatchUploadResponseBody) *MultiSourcePointBatchUploadResponse {
	s.Body = v
	return s
}

type QueryFieldLatestBySourcePointRequest struct {
	Request map[string]interface{} `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s QueryFieldLatestBySourcePointRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFieldLatestBySourcePointRequest) GoString() string {
	return s.String()
}

func (s *QueryFieldLatestBySourcePointRequest) SetRequest(v map[string]interface{}) *QueryFieldLatestBySourcePointRequest {
	s.Request = v
	return s
}

type QueryFieldLatestBySourcePointResponseBody struct {
	AccessDeniedDetail *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               []*QueryFieldLatestBySourcePointResponseBodyData             `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode     *int32                                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFieldLatestBySourcePointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFieldLatestBySourcePointResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFieldLatestBySourcePointResponseBody) SetAccessDeniedDetail(v *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail) *QueryFieldLatestBySourcePointResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBody) SetCode(v string) *QueryFieldLatestBySourcePointResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBody) SetData(v []*QueryFieldLatestBySourcePointResponseBodyData) *QueryFieldLatestBySourcePointResponseBody {
	s.Data = v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBody) SetHttpStatusCode(v int32) *QueryFieldLatestBySourcePointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBody) SetMessage(v string) *QueryFieldLatestBySourcePointResponseBody {
	s.Message = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBody) SetRequestId(v string) *QueryFieldLatestBySourcePointResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBody) SetSuccess(v bool) *QueryFieldLatestBySourcePointResponseBody {
	s.Success = &v
	return s
}

type QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail) SetAuthAction(v string) *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail) SetPolicyType(v string) *QueryFieldLatestBySourcePointResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type QueryFieldLatestBySourcePointResponseBodyData struct {
	MeasurePoint *string                                                `json:"MeasurePoint,omitempty" xml:"MeasurePoint,omitempty"`
	Node         *string                                                `json:"Node,omitempty" xml:"Node,omitempty"`
	SourcePoint  *string                                                `json:"SourcePoint,omitempty" xml:"SourcePoint,omitempty"`
	ValueType    *string                                                `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	Values       []*QueryFieldLatestBySourcePointResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryFieldLatestBySourcePointResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFieldLatestBySourcePointResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFieldLatestBySourcePointResponseBodyData) SetMeasurePoint(v string) *QueryFieldLatestBySourcePointResponseBodyData {
	s.MeasurePoint = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyData) SetNode(v string) *QueryFieldLatestBySourcePointResponseBodyData {
	s.Node = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyData) SetSourcePoint(v string) *QueryFieldLatestBySourcePointResponseBodyData {
	s.SourcePoint = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyData) SetValueType(v string) *QueryFieldLatestBySourcePointResponseBodyData {
	s.ValueType = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyData) SetValues(v []*QueryFieldLatestBySourcePointResponseBodyDataValues) *QueryFieldLatestBySourcePointResponseBodyData {
	s.Values = v
	return s
}

type QueryFieldLatestBySourcePointResponseBodyDataValues struct {
	EventTime   *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	ProcessTime *string `json:"ProcessTime,omitempty" xml:"ProcessTime,omitempty"`
	Quality     *int32  `json:"Quality,omitempty" xml:"Quality,omitempty"`
	SampleType  *string `json:"SampleType,omitempty" xml:"SampleType,omitempty"`
	Time        *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	Timestamp   *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryFieldLatestBySourcePointResponseBodyDataValues) String() string {
	return tea.Prettify(s)
}

func (s QueryFieldLatestBySourcePointResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *QueryFieldLatestBySourcePointResponseBodyDataValues) SetEventTime(v string) *QueryFieldLatestBySourcePointResponseBodyDataValues {
	s.EventTime = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyDataValues) SetProcessTime(v string) *QueryFieldLatestBySourcePointResponseBodyDataValues {
	s.ProcessTime = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyDataValues) SetQuality(v int32) *QueryFieldLatestBySourcePointResponseBodyDataValues {
	s.Quality = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyDataValues) SetSampleType(v string) *QueryFieldLatestBySourcePointResponseBodyDataValues {
	s.SampleType = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyDataValues) SetTime(v int64) *QueryFieldLatestBySourcePointResponseBodyDataValues {
	s.Time = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyDataValues) SetTimestamp(v int64) *QueryFieldLatestBySourcePointResponseBodyDataValues {
	s.Timestamp = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponseBodyDataValues) SetValue(v string) *QueryFieldLatestBySourcePointResponseBodyDataValues {
	s.Value = &v
	return s
}

type QueryFieldLatestBySourcePointResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFieldLatestBySourcePointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFieldLatestBySourcePointResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFieldLatestBySourcePointResponse) GoString() string {
	return s.String()
}

func (s *QueryFieldLatestBySourcePointResponse) SetHeaders(v map[string]*string) *QueryFieldLatestBySourcePointResponse {
	s.Headers = v
	return s
}

func (s *QueryFieldLatestBySourcePointResponse) SetStatusCode(v int32) *QueryFieldLatestBySourcePointResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFieldLatestBySourcePointResponse) SetBody(v *QueryFieldLatestBySourcePointResponseBody) *QueryFieldLatestBySourcePointResponse {
	s.Body = v
	return s
}

type QueryIndustryDeviceDataRequest struct {
	Request map[string]interface{} `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s QueryIndustryDeviceDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceDataRequest) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceDataRequest) SetRequest(v map[string]interface{}) *QueryIndustryDeviceDataRequest {
	s.Request = v
	return s
}

type QueryIndustryDeviceDataResponseBody struct {
	AccessDeniedDetail *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               []*QueryIndustryDeviceDataResponseBodyData             `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode     *int32                                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryIndustryDeviceDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceDataResponseBody) SetAccessDeniedDetail(v *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail) *QueryIndustryDeviceDataResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *QueryIndustryDeviceDataResponseBody) SetCode(v string) *QueryIndustryDeviceDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBody) SetData(v []*QueryIndustryDeviceDataResponseBodyData) *QueryIndustryDeviceDataResponseBody {
	s.Data = v
	return s
}

func (s *QueryIndustryDeviceDataResponseBody) SetHttpStatusCode(v int32) *QueryIndustryDeviceDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBody) SetMessage(v string) *QueryIndustryDeviceDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBody) SetRequestId(v string) *QueryIndustryDeviceDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBody) SetSuccess(v bool) *QueryIndustryDeviceDataResponseBody {
	s.Success = &v
	return s
}

type QueryIndustryDeviceDataResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s QueryIndustryDeviceDataResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceDataResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail) SetAuthAction(v string) *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail) SetPolicyType(v string) *QueryIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type QueryIndustryDeviceDataResponseBodyData struct {
	MeasurePoint *string                                          `json:"MeasurePoint,omitempty" xml:"MeasurePoint,omitempty"`
	Node         *string                                          `json:"Node,omitempty" xml:"Node,omitempty"`
	ValueType    *string                                          `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	Values       []*QueryIndustryDeviceDataResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryIndustryDeviceDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceDataResponseBodyData) SetMeasurePoint(v string) *QueryIndustryDeviceDataResponseBodyData {
	s.MeasurePoint = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyData) SetNode(v string) *QueryIndustryDeviceDataResponseBodyData {
	s.Node = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyData) SetValueType(v string) *QueryIndustryDeviceDataResponseBodyData {
	s.ValueType = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyData) SetValues(v []*QueryIndustryDeviceDataResponseBodyDataValues) *QueryIndustryDeviceDataResponseBodyData {
	s.Values = v
	return s
}

type QueryIndustryDeviceDataResponseBodyDataValues struct {
	EventTime   *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	ProcessTime *string `json:"ProcessTime,omitempty" xml:"ProcessTime,omitempty"`
	Quality     *int32  `json:"Quality,omitempty" xml:"Quality,omitempty"`
	Time        *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	TimeStamp   *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryIndustryDeviceDataResponseBodyDataValues) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceDataResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceDataResponseBodyDataValues) SetEventTime(v string) *QueryIndustryDeviceDataResponseBodyDataValues {
	s.EventTime = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyDataValues) SetProcessTime(v string) *QueryIndustryDeviceDataResponseBodyDataValues {
	s.ProcessTime = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyDataValues) SetQuality(v int32) *QueryIndustryDeviceDataResponseBodyDataValues {
	s.Quality = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyDataValues) SetTime(v int64) *QueryIndustryDeviceDataResponseBodyDataValues {
	s.Time = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyDataValues) SetTimeStamp(v int64) *QueryIndustryDeviceDataResponseBodyDataValues {
	s.TimeStamp = &v
	return s
}

func (s *QueryIndustryDeviceDataResponseBodyDataValues) SetValue(v string) *QueryIndustryDeviceDataResponseBodyDataValues {
	s.Value = &v
	return s
}

type QueryIndustryDeviceDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryIndustryDeviceDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryIndustryDeviceDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceDataResponse) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceDataResponse) SetHeaders(v map[string]*string) *QueryIndustryDeviceDataResponse {
	s.Headers = v
	return s
}

func (s *QueryIndustryDeviceDataResponse) SetStatusCode(v int32) *QueryIndustryDeviceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryIndustryDeviceDataResponse) SetBody(v *QueryIndustryDeviceDataResponseBody) *QueryIndustryDeviceDataResponse {
	s.Body = v
	return s
}

type QueryIndustryDeviceLimitsDataRequest struct {
	Request map[string]interface{} `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s QueryIndustryDeviceLimitsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceLimitsDataRequest) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceLimitsDataRequest) SetRequest(v map[string]interface{}) *QueryIndustryDeviceLimitsDataRequest {
	s.Request = v
	return s
}

type QueryIndustryDeviceLimitsDataResponseBody struct {
	AccessDeniedDetail *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               []*QueryIndustryDeviceLimitsDataResponseBodyData             `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode     *string                                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryIndustryDeviceLimitsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceLimitsDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceLimitsDataResponseBody) SetAccessDeniedDetail(v *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail) *QueryIndustryDeviceLimitsDataResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBody) SetCode(v string) *QueryIndustryDeviceLimitsDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBody) SetData(v []*QueryIndustryDeviceLimitsDataResponseBodyData) *QueryIndustryDeviceLimitsDataResponseBody {
	s.Data = v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBody) SetHttpStatusCode(v string) *QueryIndustryDeviceLimitsDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBody) SetMessage(v string) *QueryIndustryDeviceLimitsDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBody) SetRequestId(v string) *QueryIndustryDeviceLimitsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBody) SetSuccess(v bool) *QueryIndustryDeviceLimitsDataResponseBody {
	s.Success = &v
	return s
}

type QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail) SetAuthAction(v string) *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail) SetPolicyType(v string) *QueryIndustryDeviceLimitsDataResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type QueryIndustryDeviceLimitsDataResponseBodyData struct {
	MeasurePoint *string                                                `json:"MeasurePoint,omitempty" xml:"MeasurePoint,omitempty"`
	Node         *string                                                `json:"Node,omitempty" xml:"Node,omitempty"`
	ValueType    *string                                                `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	Values       []*QueryIndustryDeviceLimitsDataResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryIndustryDeviceLimitsDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceLimitsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyData) SetMeasurePoint(v string) *QueryIndustryDeviceLimitsDataResponseBodyData {
	s.MeasurePoint = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyData) SetNode(v string) *QueryIndustryDeviceLimitsDataResponseBodyData {
	s.Node = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyData) SetValueType(v string) *QueryIndustryDeviceLimitsDataResponseBodyData {
	s.ValueType = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyData) SetValues(v []*QueryIndustryDeviceLimitsDataResponseBodyDataValues) *QueryIndustryDeviceLimitsDataResponseBodyData {
	s.Values = v
	return s
}

type QueryIndustryDeviceLimitsDataResponseBodyDataValues struct {
	EventTime   *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	ProcessTime *string `json:"ProcessTime,omitempty" xml:"ProcessTime,omitempty"`
	Quality     *int32  `json:"Quality,omitempty" xml:"Quality,omitempty"`
	Time        *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	TimeStamp   *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryIndustryDeviceLimitsDataResponseBodyDataValues) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceLimitsDataResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyDataValues) SetEventTime(v string) *QueryIndustryDeviceLimitsDataResponseBodyDataValues {
	s.EventTime = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyDataValues) SetProcessTime(v string) *QueryIndustryDeviceLimitsDataResponseBodyDataValues {
	s.ProcessTime = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyDataValues) SetQuality(v int32) *QueryIndustryDeviceLimitsDataResponseBodyDataValues {
	s.Quality = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyDataValues) SetTime(v int64) *QueryIndustryDeviceLimitsDataResponseBodyDataValues {
	s.Time = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyDataValues) SetTimeStamp(v int64) *QueryIndustryDeviceLimitsDataResponseBodyDataValues {
	s.TimeStamp = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponseBodyDataValues) SetValue(v string) *QueryIndustryDeviceLimitsDataResponseBodyDataValues {
	s.Value = &v
	return s
}

type QueryIndustryDeviceLimitsDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryIndustryDeviceLimitsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryIndustryDeviceLimitsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceLimitsDataResponse) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceLimitsDataResponse) SetHeaders(v map[string]*string) *QueryIndustryDeviceLimitsDataResponse {
	s.Headers = v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponse) SetStatusCode(v int32) *QueryIndustryDeviceLimitsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryIndustryDeviceLimitsDataResponse) SetBody(v *QueryIndustryDeviceLimitsDataResponseBody) *QueryIndustryDeviceLimitsDataResponse {
	s.Body = v
	return s
}

type QueryIndustryDeviceStatusDataRequest struct {
	Request map[string]interface{} `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s QueryIndustryDeviceStatusDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceStatusDataRequest) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceStatusDataRequest) SetRequest(v map[string]interface{}) *QueryIndustryDeviceStatusDataRequest {
	s.Request = v
	return s
}

type QueryIndustryDeviceStatusDataResponseBody struct {
	AccessDeniedDetail *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               []*QueryIndustryDeviceStatusDataResponseBodyData             `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode     *int32                                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryIndustryDeviceStatusDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceStatusDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceStatusDataResponseBody) SetAccessDeniedDetail(v *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail) *QueryIndustryDeviceStatusDataResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBody) SetCode(v string) *QueryIndustryDeviceStatusDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBody) SetData(v []*QueryIndustryDeviceStatusDataResponseBodyData) *QueryIndustryDeviceStatusDataResponseBody {
	s.Data = v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBody) SetHttpStatusCode(v int32) *QueryIndustryDeviceStatusDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBody) SetMessage(v string) *QueryIndustryDeviceStatusDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBody) SetRequestId(v string) *QueryIndustryDeviceStatusDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBody) SetSuccess(v bool) *QueryIndustryDeviceStatusDataResponseBody {
	s.Success = &v
	return s
}

type QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail) SetAuthAction(v string) *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail) SetPolicyType(v string) *QueryIndustryDeviceStatusDataResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type QueryIndustryDeviceStatusDataResponseBodyData struct {
	MeasurePoint *string                  `json:"MeasurePoint,omitempty" xml:"MeasurePoint,omitempty"`
	Node         *string                  `json:"Node,omitempty" xml:"Node,omitempty"`
	ValueType    *string                  `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	Values       []map[string]interface{} `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryIndustryDeviceStatusDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceStatusDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceStatusDataResponseBodyData) SetMeasurePoint(v string) *QueryIndustryDeviceStatusDataResponseBodyData {
	s.MeasurePoint = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBodyData) SetNode(v string) *QueryIndustryDeviceStatusDataResponseBodyData {
	s.Node = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBodyData) SetValueType(v string) *QueryIndustryDeviceStatusDataResponseBodyData {
	s.ValueType = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponseBodyData) SetValues(v []map[string]interface{}) *QueryIndustryDeviceStatusDataResponseBodyData {
	s.Values = v
	return s
}

type QueryIndustryDeviceStatusDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryIndustryDeviceStatusDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryIndustryDeviceStatusDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryIndustryDeviceStatusDataResponse) GoString() string {
	return s.String()
}

func (s *QueryIndustryDeviceStatusDataResponse) SetHeaders(v map[string]*string) *QueryIndustryDeviceStatusDataResponse {
	s.Headers = v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponse) SetStatusCode(v int32) *QueryIndustryDeviceStatusDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryIndustryDeviceStatusDataResponse) SetBody(v *QueryIndustryDeviceStatusDataResponseBody) *QueryIndustryDeviceStatusDataResponse {
	s.Body = v
	return s
}

type SourcePointBatchRequest struct {
	Request map[string]interface{} `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s SourcePointBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s SourcePointBatchRequest) GoString() string {
	return s.String()
}

func (s *SourcePointBatchRequest) SetRequest(v map[string]interface{}) *SourcePointBatchRequest {
	s.Request = v
	return s
}

type SourcePointBatchResponseBody struct {
	AccessDeniedDetail *SourcePointBatchResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *SourcePointBatchResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode     *int32                                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SourcePointBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SourcePointBatchResponseBody) GoString() string {
	return s.String()
}

func (s *SourcePointBatchResponseBody) SetAccessDeniedDetail(v *SourcePointBatchResponseBodyAccessDeniedDetail) *SourcePointBatchResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *SourcePointBatchResponseBody) SetCode(v string) *SourcePointBatchResponseBody {
	s.Code = &v
	return s
}

func (s *SourcePointBatchResponseBody) SetData(v *SourcePointBatchResponseBodyData) *SourcePointBatchResponseBody {
	s.Data = v
	return s
}

func (s *SourcePointBatchResponseBody) SetHttpStatusCode(v int32) *SourcePointBatchResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SourcePointBatchResponseBody) SetMessage(v string) *SourcePointBatchResponseBody {
	s.Message = &v
	return s
}

func (s *SourcePointBatchResponseBody) SetRequestId(v string) *SourcePointBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *SourcePointBatchResponseBody) SetSuccess(v bool) *SourcePointBatchResponseBody {
	s.Success = &v
	return s
}

type SourcePointBatchResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s SourcePointBatchResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s SourcePointBatchResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *SourcePointBatchResponseBodyAccessDeniedDetail) SetAuthAction(v string) *SourcePointBatchResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *SourcePointBatchResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *SourcePointBatchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *SourcePointBatchResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *SourcePointBatchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *SourcePointBatchResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *SourcePointBatchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *SourcePointBatchResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *SourcePointBatchResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *SourcePointBatchResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *SourcePointBatchResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *SourcePointBatchResponseBodyAccessDeniedDetail) SetPolicyType(v string) *SourcePointBatchResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type SourcePointBatchResponseBodyData struct {
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s SourcePointBatchResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SourcePointBatchResponseBodyData) GoString() string {
	return s.String()
}

func (s *SourcePointBatchResponseBodyData) SetCount(v int32) *SourcePointBatchResponseBodyData {
	s.Count = &v
	return s
}

type SourcePointBatchResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SourcePointBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SourcePointBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s SourcePointBatchResponse) GoString() string {
	return s.String()
}

func (s *SourcePointBatchResponse) SetHeaders(v map[string]*string) *SourcePointBatchResponse {
	s.Headers = v
	return s
}

func (s *SourcePointBatchResponse) SetStatusCode(v int32) *SourcePointBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *SourcePointBatchResponse) SetBody(v *SourcePointBatchResponseBody) *SourcePointBatchResponse {
	s.Body = v
	return s
}

type UploadIndustryDeviceDataRequest struct {
	Request map[string]interface{} `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s UploadIndustryDeviceDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadIndustryDeviceDataRequest) GoString() string {
	return s.String()
}

func (s *UploadIndustryDeviceDataRequest) SetRequest(v map[string]interface{}) *UploadIndustryDeviceDataRequest {
	s.Request = v
	return s
}

type UploadIndustryDeviceDataResponseBody struct {
	AccessDeniedDetail *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *UploadIndustryDeviceDataResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode     *int32                                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadIndustryDeviceDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadIndustryDeviceDataResponseBody) GoString() string {
	return s.String()
}

func (s *UploadIndustryDeviceDataResponseBody) SetAccessDeniedDetail(v *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail) *UploadIndustryDeviceDataResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UploadIndustryDeviceDataResponseBody) SetCode(v string) *UploadIndustryDeviceDataResponseBody {
	s.Code = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBody) SetData(v *UploadIndustryDeviceDataResponseBodyData) *UploadIndustryDeviceDataResponseBody {
	s.Data = v
	return s
}

func (s *UploadIndustryDeviceDataResponseBody) SetHttpStatusCode(v int32) *UploadIndustryDeviceDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBody) SetMessage(v string) *UploadIndustryDeviceDataResponseBody {
	s.Message = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBody) SetRequestId(v string) *UploadIndustryDeviceDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBody) SetSuccess(v bool) *UploadIndustryDeviceDataResponseBody {
	s.Success = &v
	return s
}

type UploadIndustryDeviceDataResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s UploadIndustryDeviceDataResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s UploadIndustryDeviceDataResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UploadIndustryDeviceDataResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type UploadIndustryDeviceDataResponseBodyData struct {
	Count   *int32                                             `json:"Count,omitempty" xml:"Count,omitempty"`
	Message []*UploadIndustryDeviceDataResponseBodyDataMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UploadIndustryDeviceDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UploadIndustryDeviceDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadIndustryDeviceDataResponseBodyData) SetCount(v int32) *UploadIndustryDeviceDataResponseBodyData {
	s.Count = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBodyData) SetMessage(v []*UploadIndustryDeviceDataResponseBodyDataMessage) *UploadIndustryDeviceDataResponseBodyData {
	s.Message = v
	return s
}

type UploadIndustryDeviceDataResponseBodyDataMessage struct {
	ErrorMsg     *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	MeasurePoint *string `json:"MeasurePoint,omitempty" xml:"MeasurePoint,omitempty"`
	Node         *string `json:"Node,omitempty" xml:"Node,omitempty"`
}

func (s UploadIndustryDeviceDataResponseBodyDataMessage) String() string {
	return tea.Prettify(s)
}

func (s UploadIndustryDeviceDataResponseBodyDataMessage) GoString() string {
	return s.String()
}

func (s *UploadIndustryDeviceDataResponseBodyDataMessage) SetErrorMsg(v string) *UploadIndustryDeviceDataResponseBodyDataMessage {
	s.ErrorMsg = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBodyDataMessage) SetMeasurePoint(v string) *UploadIndustryDeviceDataResponseBodyDataMessage {
	s.MeasurePoint = &v
	return s
}

func (s *UploadIndustryDeviceDataResponseBodyDataMessage) SetNode(v string) *UploadIndustryDeviceDataResponseBodyDataMessage {
	s.Node = &v
	return s
}

type UploadIndustryDeviceDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadIndustryDeviceDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadIndustryDeviceDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadIndustryDeviceDataResponse) GoString() string {
	return s.String()
}

func (s *UploadIndustryDeviceDataResponse) SetHeaders(v map[string]*string) *UploadIndustryDeviceDataResponse {
	s.Headers = v
	return s
}

func (s *UploadIndustryDeviceDataResponse) SetStatusCode(v int32) *UploadIndustryDeviceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadIndustryDeviceDataResponse) SetBody(v *UploadIndustryDeviceDataResponseBody) *UploadIndustryDeviceDataResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou": tea.String("et-industry.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("et-industry-openapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - GetMqttConnectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMqttConnectResponse
func (client *Client) GetMqttConnectWithOptions(request *GetMqttConnectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMqttConnectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["Request"] = request.Request
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMqttConnect"),
		Version:     tea.String("2020-08-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/igate/collaboration/pop/getmqttconnect"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetMqttConnectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetMqttConnectResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetMqttConnectRequest
//
// @return GetMqttConnectResponse
func (client *Client) GetMqttConnect(request *GetMqttConnectRequest) (_result *GetMqttConnectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMqttConnectResponse{}
	_body, _err := client.GetMqttConnectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetNonceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNonceResponse
func (client *Client) GetNonceWithOptions(request *GetNonceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetNonceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["Request"] = request.Request
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNonce"),
		Version:     tea.String("2020-08-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/igate/collaboration/pop/getnonce"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetNonceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetNonceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetNonceRequest
//
// @return GetNonceResponse
func (client *Client) GetNonce(request *GetNonceRequest) (_result *GetNonceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNonceResponse{}
	_body, _err := client.GetNonceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListMeasurePointListByNodeCodePageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMeasurePointListByNodeCodePageResponse
func (client *Client) ListMeasurePointListByNodeCodePageWithOptions(request *ListMeasurePointListByNodeCodePageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMeasurePointListByNodeCodePageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["Request"] = request.Request
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMeasurePointListByNodeCodePage"),
		Version:     tea.String("2020-08-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/igate/node/pop/measurepointlistbynodecodepage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListMeasurePointListByNodeCodePageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListMeasurePointListByNodeCodePageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListMeasurePointListByNodeCodePageRequest
//
// @return ListMeasurePointListByNodeCodePageResponse
func (client *Client) ListMeasurePointListByNodeCodePage(request *ListMeasurePointListByNodeCodePageRequest) (_result *ListMeasurePointListByNodeCodePageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMeasurePointListByNodeCodePageResponse{}
	_body, _err := client.ListMeasurePointListByNodeCodePageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多值批量上报
//
// @param request - MultiFieldBatchUploadRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MultiFieldBatchUploadResponse
func (client *Client) MultiFieldBatchUploadWithOptions(request *MultiFieldBatchUploadRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MultiFieldBatchUploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MultiFieldBatchUpload"),
		Version:     tea.String("2020-08-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/igate/timeseries/upload/pop/multifieldbatchv2"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &MultiFieldBatchUploadResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &MultiFieldBatchUploadResponse{}
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
// 多值批量上报
//
// @param request - MultiFieldBatchUploadRequest
//
// @return MultiFieldBatchUploadResponse
func (client *Client) MultiFieldBatchUpload(request *MultiFieldBatchUploadRequest) (_result *MultiFieldBatchUploadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MultiFieldBatchUploadResponse{}
	_body, _err := client.MultiFieldBatchUploadWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多源点位批量上报
//
// @param request - MultiSourcePointBatchUploadRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MultiSourcePointBatchUploadResponse
func (client *Client) MultiSourcePointBatchUploadWithOptions(request *MultiSourcePointBatchUploadRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MultiSourcePointBatchUploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MultiSourcePointBatchUpload"),
		Version:     tea.String("2020-08-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/igate/timeseries/upload/pop/sourcepointbatchv2"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &MultiSourcePointBatchUploadResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &MultiSourcePointBatchUploadResponse{}
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
// 多源点位批量上报
//
// @param request - MultiSourcePointBatchUploadRequest
//
// @return MultiSourcePointBatchUploadResponse
func (client *Client) MultiSourcePointBatchUpload(request *MultiSourcePointBatchUploadRequest) (_result *MultiSourcePointBatchUploadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MultiSourcePointBatchUploadResponse{}
	_body, _err := client.MultiSourcePointBatchUploadWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryFieldLatestBySourcePointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFieldLatestBySourcePointResponse
func (client *Client) QueryFieldLatestBySourcePointWithOptions(request *QueryFieldLatestBySourcePointRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryFieldLatestBySourcePointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["Request"] = request.Request
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFieldLatestBySourcePoint"),
		Version:     tea.String("2020-08-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/igate/timeseries/query/pop/multifieldlatestbysourcepoint"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryFieldLatestBySourcePointResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryFieldLatestBySourcePointResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - QueryFieldLatestBySourcePointRequest
//
// @return QueryFieldLatestBySourcePointResponse
func (client *Client) QueryFieldLatestBySourcePoint(request *QueryFieldLatestBySourcePointRequest) (_result *QueryFieldLatestBySourcePointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryFieldLatestBySourcePointResponse{}
	_body, _err := client.QueryFieldLatestBySourcePointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryIndustryDeviceDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryIndustryDeviceDataResponse
func (client *Client) QueryIndustryDeviceDataWithOptions(request *QueryIndustryDeviceDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryIndustryDeviceDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["Request"] = request.Request
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryIndustryDeviceData"),
		Version:     tea.String("2020-08-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/igate/timeseries/query/pop/multifieldlatest"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryIndustryDeviceDataResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryIndustryDeviceDataResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - QueryIndustryDeviceDataRequest
//
// @return QueryIndustryDeviceDataResponse
func (client *Client) QueryIndustryDeviceData(request *QueryIndustryDeviceDataRequest) (_result *QueryIndustryDeviceDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryIndustryDeviceDataResponse{}
	_body, _err := client.QueryIndustryDeviceDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryIndustryDeviceLimitsDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryIndustryDeviceLimitsDataResponse
func (client *Client) QueryIndustryDeviceLimitsDataWithOptions(request *QueryIndustryDeviceLimitsDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryIndustryDeviceLimitsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["Request"] = request.Request
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryIndustryDeviceLimitsData"),
		Version:     tea.String("2020-08-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/igate/timeseries/query/pop/multifieldrange"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryIndustryDeviceLimitsDataResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryIndustryDeviceLimitsDataResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - QueryIndustryDeviceLimitsDataRequest
//
// @return QueryIndustryDeviceLimitsDataResponse
func (client *Client) QueryIndustryDeviceLimitsData(request *QueryIndustryDeviceLimitsDataRequest) (_result *QueryIndustryDeviceLimitsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryIndustryDeviceLimitsDataResponse{}
	_body, _err := client.QueryIndustryDeviceLimitsDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryIndustryDeviceStatusDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryIndustryDeviceStatusDataResponse
func (client *Client) QueryIndustryDeviceStatusDataWithOptions(request *QueryIndustryDeviceStatusDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryIndustryDeviceStatusDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["Request"] = request.Request
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryIndustryDeviceStatusData"),
		Version:     tea.String("2020-08-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/igate/timeseries/query/pop/multifieldrangestatus"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryIndustryDeviceStatusDataResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryIndustryDeviceStatusDataResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - QueryIndustryDeviceStatusDataRequest
//
// @return QueryIndustryDeviceStatusDataResponse
func (client *Client) QueryIndustryDeviceStatusData(request *QueryIndustryDeviceStatusDataRequest) (_result *QueryIndustryDeviceStatusDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryIndustryDeviceStatusDataResponse{}
	_body, _err := client.QueryIndustryDeviceStatusDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SourcePointBatchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SourcePointBatchResponse
func (client *Client) SourcePointBatchWithOptions(request *SourcePointBatchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SourcePointBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["Request"] = request.Request
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SourcePointBatch"),
		Version:     tea.String("2020-08-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/igate/timeseries/upload/pop/sourcepointbatch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SourcePointBatchResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SourcePointBatchResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - SourcePointBatchRequest
//
// @return SourcePointBatchResponse
func (client *Client) SourcePointBatch(request *SourcePointBatchRequest) (_result *SourcePointBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SourcePointBatchResponse{}
	_body, _err := client.SourcePointBatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UploadIndustryDeviceDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadIndustryDeviceDataResponse
func (client *Client) UploadIndustryDeviceDataWithOptions(request *UploadIndustryDeviceDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadIndustryDeviceDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["Request"] = request.Request
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadIndustryDeviceData"),
		Version:     tea.String("2020-08-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/igate/timeseries/upload/pop/multifieldbatch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UploadIndustryDeviceDataResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UploadIndustryDeviceDataResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - UploadIndustryDeviceDataRequest
//
// @return UploadIndustryDeviceDataResponse
func (client *Client) UploadIndustryDeviceData(request *UploadIndustryDeviceDataRequest) (_result *UploadIndustryDeviceDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadIndustryDeviceDataResponse{}
	_body, _err := client.UploadIndustryDeviceDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
