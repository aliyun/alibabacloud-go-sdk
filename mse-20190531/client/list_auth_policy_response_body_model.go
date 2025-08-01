// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAuthPolicyResponseBody
	GetCode() *int32
	SetData(v *ListAuthPolicyResponseBodyData) *ListAuthPolicyResponseBody
	GetData() *ListAuthPolicyResponseBodyData
	SetHttpStatusCode(v int32) *ListAuthPolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAuthPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAuthPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAuthPolicyResponseBody
	GetSuccess() *bool
}

type ListAuthPolicyResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	Data *ListAuthPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 05A5A150-4A5F-5A8C-97D6-710776CC8408
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAuthPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthPolicyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAuthPolicyResponseBody) GetData() *ListAuthPolicyResponseBodyData {
	return s.Data
}

func (s *ListAuthPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAuthPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAuthPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAuthPolicyResponseBody) SetCode(v int32) *ListAuthPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *ListAuthPolicyResponseBody) SetData(v *ListAuthPolicyResponseBodyData) *ListAuthPolicyResponseBody {
	s.Data = v
	return s
}

func (s *ListAuthPolicyResponseBody) SetHttpStatusCode(v int32) *ListAuthPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAuthPolicyResponseBody) SetMessage(v string) *ListAuthPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *ListAuthPolicyResponseBody) SetRequestId(v string) *ListAuthPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthPolicyResponseBody) SetSuccess(v bool) *ListAuthPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *ListAuthPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAuthPolicyResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data returned.
	Result []*ListAuthPolicyResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 11
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListAuthPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAuthPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAuthPolicyResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthPolicyResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthPolicyResponseBodyData) GetResult() []*ListAuthPolicyResponseBodyDataResult {
	return s.Result
}

func (s *ListAuthPolicyResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListAuthPolicyResponseBodyData) SetPageNumber(v int32) *ListAuthPolicyResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAuthPolicyResponseBodyData) SetPageSize(v int32) *ListAuthPolicyResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAuthPolicyResponseBodyData) SetResult(v []*ListAuthPolicyResponseBodyDataResult) *ListAuthPolicyResponseBodyData {
	s.Result = v
	return s
}

func (s *ListAuthPolicyResponseBodyData) SetTotalSize(v int32) *ListAuthPolicyResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListAuthPolicyResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAuthPolicyResponseBodyDataResult struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 19039813784***
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The application ID.
	//
	// example:
	//
	// hkhon1po62@5f1b08becb*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The content of the service authentication rule.
	AuthRule []*ListAuthPolicyResponseBodyDataResultAuthRule `json:"AuthRule,omitempty" xml:"AuthRule,omitempty" type:"Repeated"`
	// The rule type. Valid values:
	//
	// 	- 0: by application
	//
	// 	- 1: by namespace
	//
	// example:
	//
	// 0
	AuthType *int32 `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// Indicates whether the rule was enabled or disabled. Valid values:
	//
	// 	- `true`: enabled
	//
	// 	- `false`: disabled
	//
	// example:
	//
	// false
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 204
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The namespace.
	//
	// example:
	//
	// c19c6c500e1ff4d7abc7bed9b8236***
	K8sNamespace *string `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	// The name of the authentication rule.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// ced54a95-4e33-4bda-be7e-37e95868***
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **SPRING_CLOUD**
	//
	// 	- **DUBBO**
	//
	// 	- **istio**
	//
	// example:
	//
	// SPRING_CLOUD
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the application.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAuthPolicyResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListAuthPolicyResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListAuthPolicyResponseBodyDataResult) GetAccountId() *string {
	return s.AccountId
}

func (s *ListAuthPolicyResponseBodyDataResult) GetAppId() *string {
	return s.AppId
}

func (s *ListAuthPolicyResponseBodyDataResult) GetAppName() *string {
	return s.AppName
}

func (s *ListAuthPolicyResponseBodyDataResult) GetAuthRule() []*ListAuthPolicyResponseBodyDataResultAuthRule {
	return s.AuthRule
}

func (s *ListAuthPolicyResponseBodyDataResult) GetAuthType() *int32 {
	return s.AuthType
}

func (s *ListAuthPolicyResponseBodyDataResult) GetEnable() *bool {
	return s.Enable
}

func (s *ListAuthPolicyResponseBodyDataResult) GetId() *int32 {
	return s.Id
}

func (s *ListAuthPolicyResponseBodyDataResult) GetK8sNamespace() *string {
	return s.K8sNamespace
}

func (s *ListAuthPolicyResponseBodyDataResult) GetName() *string {
	return s.Name
}

func (s *ListAuthPolicyResponseBodyDataResult) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListAuthPolicyResponseBodyDataResult) GetProtocol() *string {
	return s.Protocol
}

func (s *ListAuthPolicyResponseBodyDataResult) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAuthPolicyResponseBodyDataResult) GetSource() *string {
	return s.Source
}

func (s *ListAuthPolicyResponseBodyDataResult) GetStatus() *int32 {
	return s.Status
}

func (s *ListAuthPolicyResponseBodyDataResult) SetAccountId(v string) *ListAuthPolicyResponseBodyDataResult {
	s.AccountId = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetAppId(v string) *ListAuthPolicyResponseBodyDataResult {
	s.AppId = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetAppName(v string) *ListAuthPolicyResponseBodyDataResult {
	s.AppName = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetAuthRule(v []*ListAuthPolicyResponseBodyDataResultAuthRule) *ListAuthPolicyResponseBodyDataResult {
	s.AuthRule = v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetAuthType(v int32) *ListAuthPolicyResponseBodyDataResult {
	s.AuthType = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetEnable(v bool) *ListAuthPolicyResponseBodyDataResult {
	s.Enable = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetId(v int32) *ListAuthPolicyResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetK8sNamespace(v string) *ListAuthPolicyResponseBodyDataResult {
	s.K8sNamespace = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetName(v string) *ListAuthPolicyResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetNamespaceId(v string) *ListAuthPolicyResponseBodyDataResult {
	s.NamespaceId = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetProtocol(v string) *ListAuthPolicyResponseBodyDataResult {
	s.Protocol = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetRegionId(v string) *ListAuthPolicyResponseBodyDataResult {
	s.RegionId = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetSource(v string) *ListAuthPolicyResponseBodyDataResult {
	s.Source = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) SetStatus(v int32) *ListAuthPolicyResponseBodyDataResult {
	s.Status = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type ListAuthPolicyResponseBodyDataResultAuthRule struct {
	// The IDs of applications.
	AppIds []*string `json:"AppIds,omitempty" xml:"AppIds,omitempty" type:"Repeated"`
	// The rule type. Valid values:
	//
	// 	- 0: by application
	//
	// 	- 1: by namespace
	//
	// example:
	//
	// 0
	AuthType *int32 `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// Indicates whether the rule is a blacklist rule.
	//
	// example:
	//
	// false
	Black *bool `json:"Black,omitempty" xml:"Black,omitempty"`
	// The queried namespaces.
	K8sNamespaces []*string `json:"K8sNamespaces,omitempty" xml:"K8sNamespaces,omitempty" type:"Repeated"`
	// The request method.
	Method *ListAuthPolicyResponseBodyDataResultAuthRuleMethod `json:"Method,omitempty" xml:"Method,omitempty" type:"Struct"`
	// The service path.
	//
	// example:
	//
	// /a
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ListAuthPolicyResponseBodyDataResultAuthRule) String() string {
	return dara.Prettify(s)
}

func (s ListAuthPolicyResponseBodyDataResultAuthRule) GoString() string {
	return s.String()
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) GetAppIds() []*string {
	return s.AppIds
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) GetAuthType() *int32 {
	return s.AuthType
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) GetBlack() *bool {
	return s.Black
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) GetK8sNamespaces() []*string {
	return s.K8sNamespaces
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) GetMethod() *ListAuthPolicyResponseBodyDataResultAuthRuleMethod {
	return s.Method
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) GetPath() *string {
	return s.Path
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) SetAppIds(v []*string) *ListAuthPolicyResponseBodyDataResultAuthRule {
	s.AppIds = v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) SetAuthType(v int32) *ListAuthPolicyResponseBodyDataResultAuthRule {
	s.AuthType = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) SetBlack(v bool) *ListAuthPolicyResponseBodyDataResultAuthRule {
	s.Black = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) SetK8sNamespaces(v []*string) *ListAuthPolicyResponseBodyDataResultAuthRule {
	s.K8sNamespaces = v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) SetMethod(v *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) *ListAuthPolicyResponseBodyDataResultAuthRule {
	s.Method = v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) SetPath(v string) *ListAuthPolicyResponseBodyDataResultAuthRule {
	s.Path = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRule) Validate() error {
	return dara.Validate(s)
}

type ListAuthPolicyResponseBodyDataResultAuthRuleMethod struct {
	// The group.
	//
	// example:
	//
	// default
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The method name.
	//
	// example:
	//
	// hello
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The types of request parameters.
	ParameterTypes []*string `json:"ParameterTypes,omitempty" xml:"ParameterTypes,omitempty" type:"Repeated"`
	// The type of the return value.
	//
	// example:
	//
	// Boolean
	ReturnType *string `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
	// The service name.
	//
	// example:
	//
	// spring-cloud-a
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The method version.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAuthPolicyResponseBodyDataResultAuthRuleMethod) String() string {
	return dara.Prettify(s)
}

func (s ListAuthPolicyResponseBodyDataResultAuthRuleMethod) GoString() string {
	return s.String()
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) GetGroup() *string {
	return s.Group
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) GetName() *string {
	return s.Name
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) GetParameterTypes() []*string {
	return s.ParameterTypes
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) GetReturnType() *string {
	return s.ReturnType
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) GetVersion() *string {
	return s.Version
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) SetGroup(v string) *ListAuthPolicyResponseBodyDataResultAuthRuleMethod {
	s.Group = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) SetName(v string) *ListAuthPolicyResponseBodyDataResultAuthRuleMethod {
	s.Name = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) SetParameterTypes(v []*string) *ListAuthPolicyResponseBodyDataResultAuthRuleMethod {
	s.ParameterTypes = v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) SetReturnType(v string) *ListAuthPolicyResponseBodyDataResultAuthRuleMethod {
	s.ReturnType = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) SetServiceName(v string) *ListAuthPolicyResponseBodyDataResultAuthRuleMethod {
	s.ServiceName = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) SetVersion(v string) *ListAuthPolicyResponseBodyDataResultAuthRuleMethod {
	s.Version = &v
	return s
}

func (s *ListAuthPolicyResponseBodyDataResultAuthRuleMethod) Validate() error {
	return dara.Validate(s)
}
