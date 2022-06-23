// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetProductComponentDetailRequest struct {
	Uid                                *string `json:"uid,omitempty" xml:"uid,omitempty"`
	VersionUID                         *string `json:"versionUID,omitempty" xml:"versionUID,omitempty"`
	ProductComponentVersionRelationUID *string `json:"productComponentVersionRelationUID,omitempty" xml:"productComponentVersionRelationUID,omitempty"`
	TestField                          *string `json:"testField,omitempty" xml:"testField,omitempty"`
	TestField4                         *string `json:"TestField4,omitempty" xml:"TestField4,omitempty"`
	TestField5                         *string `json:"TestField5,omitempty" xml:"TestField5,omitempty"`
	TestField51                        *string `json:"TestField51,omitempty" xml:"TestField51,omitempty"`
	TestField59                        *string `json:"TestField59,omitempty" xml:"TestField59,omitempty"`
	TestField6                         *string `json:"TestField6,omitempty" xml:"TestField6,omitempty"`
	TestField2                         *string `json:"testField2,omitempty" xml:"testField2,omitempty"`
	TestField3                         *string `json:"testField3,omitempty" xml:"testField3,omitempty"`
}

func (s GetProductComponentDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductComponentDetailRequest) GoString() string {
	return s.String()
}

func (s *GetProductComponentDetailRequest) SetUid(v string) *GetProductComponentDetailRequest {
	s.Uid = &v
	return s
}

func (s *GetProductComponentDetailRequest) SetVersionUID(v string) *GetProductComponentDetailRequest {
	s.VersionUID = &v
	return s
}

func (s *GetProductComponentDetailRequest) SetProductComponentVersionRelationUID(v string) *GetProductComponentDetailRequest {
	s.ProductComponentVersionRelationUID = &v
	return s
}

func (s *GetProductComponentDetailRequest) SetTestField(v string) *GetProductComponentDetailRequest {
	s.TestField = &v
	return s
}

func (s *GetProductComponentDetailRequest) SetTestField4(v string) *GetProductComponentDetailRequest {
	s.TestField4 = &v
	return s
}

func (s *GetProductComponentDetailRequest) SetTestField5(v string) *GetProductComponentDetailRequest {
	s.TestField5 = &v
	return s
}

func (s *GetProductComponentDetailRequest) SetTestField51(v string) *GetProductComponentDetailRequest {
	s.TestField51 = &v
	return s
}

func (s *GetProductComponentDetailRequest) SetTestField59(v string) *GetProductComponentDetailRequest {
	s.TestField59 = &v
	return s
}

func (s *GetProductComponentDetailRequest) SetTestField6(v string) *GetProductComponentDetailRequest {
	s.TestField6 = &v
	return s
}

func (s *GetProductComponentDetailRequest) SetTestField2(v string) *GetProductComponentDetailRequest {
	s.TestField2 = &v
	return s
}

func (s *GetProductComponentDetailRequest) SetTestField3(v string) *GetProductComponentDetailRequest {
	s.TestField3 = &v
	return s
}

type GetProductComponentDetailResponseBody struct {
	Data    *GetProductComponentDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                    `json:"errCode,omitempty" xml:"errCode,omitempty"`
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProductComponentDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductComponentDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductComponentDetailResponseBody) SetData(v *GetProductComponentDetailResponseBodyData) *GetProductComponentDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetProductComponentDetailResponseBody) SetErrCode(v string) *GetProductComponentDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetProductComponentDetailResponseBody) SetSuccess(v bool) *GetProductComponentDetailResponseBody {
	s.Success = &v
	return s
}

type GetProductComponentDetailResponseBodyData struct {
	AppVersion                   *string                                                                  `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	Category                     *string                                                                  `json:"category,omitempty" xml:"category,omitempty"`
	ChildrenComponentVersionList []*GetProductComponentDetailResponseBodyDataChildrenComponentVersionList `json:"childrenComponentVersionList,omitempty" xml:"childrenComponentVersionList,omitempty" type:"Repeated"`
	Class                        *string                                                                  `json:"class,omitempty" xml:"class,omitempty"`
	ComponentName                *string                                                                  `json:"componentName,omitempty" xml:"componentName,omitempty"`
	ComponentUID                 *string                                                                  `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	Description                  *string                                                                  `json:"description,omitempty" xml:"description,omitempty"`
	Documents                    []*string                                                                `json:"documents,omitempty" xml:"documents,omitempty" type:"Repeated"`
	Enable                       *bool                                                                    `json:"enable,omitempty" xml:"enable,omitempty"`
	HasDependency                *bool                                                                    `json:"hasDependency,omitempty" xml:"hasDependency,omitempty"`
	ImagesMapping                *string                                                                  `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	Namespace                    *string                                                                  `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OrchestrationValues          *string                                                                  `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	PackageURL                   *string                                                                  `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	ParentComponent              *bool                                                                    `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	Priority                     *int32                                                                   `json:"priority,omitempty" xml:"priority,omitempty"`
	ProductComponentVersionUID   *string                                                                  `json:"productComponentVersionUID,omitempty" xml:"productComponentVersionUID,omitempty"`
	Provider                     *string                                                                  `json:"provider,omitempty" xml:"provider,omitempty"`
	Readme                       *string                                                                  `json:"readme,omitempty" xml:"readme,omitempty"`
	Resources                    *string                                                                  `json:"resources,omitempty" xml:"resources,omitempty"`
	Singleton                    *bool                                                                    `json:"singleton,omitempty" xml:"singleton,omitempty"`
	Uid                          *string                                                                  `json:"uid,omitempty" xml:"uid,omitempty"`
	Version                      *string                                                                  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetProductComponentDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductComponentDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductComponentDetailResponseBodyData) SetAppVersion(v string) *GetProductComponentDetailResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetCategory(v string) *GetProductComponentDetailResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetChildrenComponentVersionList(v []*GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) *GetProductComponentDetailResponseBodyData {
	s.ChildrenComponentVersionList = v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetClass(v string) *GetProductComponentDetailResponseBodyData {
	s.Class = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetComponentName(v string) *GetProductComponentDetailResponseBodyData {
	s.ComponentName = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetComponentUID(v string) *GetProductComponentDetailResponseBodyData {
	s.ComponentUID = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetDescription(v string) *GetProductComponentDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetDocuments(v []*string) *GetProductComponentDetailResponseBodyData {
	s.Documents = v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetEnable(v bool) *GetProductComponentDetailResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetHasDependency(v bool) *GetProductComponentDetailResponseBodyData {
	s.HasDependency = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetImagesMapping(v string) *GetProductComponentDetailResponseBodyData {
	s.ImagesMapping = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetNamespace(v string) *GetProductComponentDetailResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetOrchestrationValues(v string) *GetProductComponentDetailResponseBodyData {
	s.OrchestrationValues = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetPackageURL(v string) *GetProductComponentDetailResponseBodyData {
	s.PackageURL = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetParentComponent(v bool) *GetProductComponentDetailResponseBodyData {
	s.ParentComponent = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetPriority(v int32) *GetProductComponentDetailResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetProductComponentVersionUID(v string) *GetProductComponentDetailResponseBodyData {
	s.ProductComponentVersionUID = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetProvider(v string) *GetProductComponentDetailResponseBodyData {
	s.Provider = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetReadme(v string) *GetProductComponentDetailResponseBodyData {
	s.Readme = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetResources(v string) *GetProductComponentDetailResponseBodyData {
	s.Resources = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetSingleton(v bool) *GetProductComponentDetailResponseBodyData {
	s.Singleton = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetUid(v string) *GetProductComponentDetailResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetVersion(v string) *GetProductComponentDetailResponseBodyData {
	s.Version = &v
	return s
}

type GetProductComponentDetailResponseBodyDataChildrenComponentVersionList struct {
	AppVersion                 *string   `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	Category                   *string   `json:"category,omitempty" xml:"category,omitempty"`
	Class                      *string   `json:"class,omitempty" xml:"class,omitempty"`
	ComponentName              *string   `json:"componentName,omitempty" xml:"componentName,omitempty"`
	ComponentUID               *string   `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	Description                *string   `json:"description,omitempty" xml:"description,omitempty"`
	Documents                  []*string `json:"documents,omitempty" xml:"documents,omitempty" type:"Repeated"`
	Enable                     *bool     `json:"enable,omitempty" xml:"enable,omitempty"`
	ImagesMapping              *string   `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	Namespace                  *string   `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OrchestrationValues        *string   `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	PackageURL                 *string   `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	ParentComponent            *bool     `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	Priority                   *int32    `json:"priority,omitempty" xml:"priority,omitempty"`
	ProductComponentVersionUID *string   `json:"productComponentVersionUID,omitempty" xml:"productComponentVersionUID,omitempty"`
	Provider                   *string   `json:"provider,omitempty" xml:"provider,omitempty"`
	Readme                     *string   `json:"readme,omitempty" xml:"readme,omitempty"`
	Resources                  *string   `json:"resources,omitempty" xml:"resources,omitempty"`
	Singleton                  *bool     `json:"singleton,omitempty" xml:"singleton,omitempty"`
	Uid                        *string   `json:"uid,omitempty" xml:"uid,omitempty"`
	Version                    *string   `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) String() string {
	return tea.Prettify(s)
}

func (s GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) GoString() string {
	return s.String()
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetAppVersion(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.AppVersion = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetCategory(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Category = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetClass(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Class = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetComponentName(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.ComponentName = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetComponentUID(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.ComponentUID = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetDescription(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Description = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetDocuments(v []*string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Documents = v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetEnable(v bool) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Enable = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetImagesMapping(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.ImagesMapping = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetNamespace(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Namespace = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetOrchestrationValues(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.OrchestrationValues = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetPackageURL(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.PackageURL = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetParentComponent(v bool) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.ParentComponent = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetPriority(v int32) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Priority = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetProductComponentVersionUID(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.ProductComponentVersionUID = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetProvider(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Provider = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetReadme(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Readme = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetResources(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Resources = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetSingleton(v bool) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Singleton = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetUid(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Uid = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetVersion(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Version = &v
	return s
}

type GetProductComponentDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProductComponentDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductComponentDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductComponentDetailResponse) GoString() string {
	return s.String()
}

func (s *GetProductComponentDetailResponse) SetHeaders(v map[string]*string) *GetProductComponentDetailResponse {
	s.Headers = v
	return s
}

func (s *GetProductComponentDetailResponse) SetStatusCode(v int32) *GetProductComponentDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductComponentDetailResponse) SetBody(v *GetProductComponentDetailResponseBody) *GetProductComponentDetailResponse {
	s.Body = v
	return s
}

type QkTestResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QkTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QkTestResponseBody) GoString() string {
	return s.String()
}

func (s *QkTestResponseBody) SetRequestId(v string) *QkTestResponseBody {
	s.RequestId = &v
	return s
}

type QkTestResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QkTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QkTestResponse) String() string {
	return tea.Prettify(s)
}

func (s QkTestResponse) GoString() string {
	return s.String()
}

func (s *QkTestResponse) SetHeaders(v map[string]*string) *QkTestResponse {
	s.Headers = v
	return s
}

func (s *QkTestResponse) SetStatusCode(v int32) *QkTestResponse {
	s.StatusCode = &v
	return s
}

func (s *QkTestResponse) SetBody(v *QkTestResponseBody) *QkTestResponse {
	s.Body = v
	return s
}

type TestProRequest struct {
	Uid                                *string `json:"uid,omitempty" xml:"uid,omitempty"`
	VersionUID                         *string `json:"versionUID,omitempty" xml:"versionUID,omitempty"`
	ProductComponentVersionRelationUID *string `json:"productComponentVersionRelationUID,omitempty" xml:"productComponentVersionRelationUID,omitempty"`
	TestField                          *string `json:"testField,omitempty" xml:"testField,omitempty"`
	TestField2                         *string `json:"testField2,omitempty" xml:"testField2,omitempty"`
	TestFieldTwo                       *string `json:"testFieldTwo,omitempty" xml:"testFieldTwo,omitempty"`
	TestField4                         *string `json:"testField4,omitempty" xml:"testField4,omitempty"`
	TestField5                         *string `json:"testField5,omitempty" xml:"testField5,omitempty"`
	TestField6                         *string `json:"testField6,omitempty" xml:"testField6,omitempty"`
	TestField8                         *string `json:"testField8,omitempty" xml:"testField8,omitempty"`
	TestField3                         *string `json:"testField3,omitempty" xml:"testField3,omitempty"`
	TestField7                         *string `json:"testField7,omitempty" xml:"testField7,omitempty"`
}

func (s TestProRequest) String() string {
	return tea.Prettify(s)
}

func (s TestProRequest) GoString() string {
	return s.String()
}

func (s *TestProRequest) SetUid(v string) *TestProRequest {
	s.Uid = &v
	return s
}

func (s *TestProRequest) SetVersionUID(v string) *TestProRequest {
	s.VersionUID = &v
	return s
}

func (s *TestProRequest) SetProductComponentVersionRelationUID(v string) *TestProRequest {
	s.ProductComponentVersionRelationUID = &v
	return s
}

func (s *TestProRequest) SetTestField(v string) *TestProRequest {
	s.TestField = &v
	return s
}

func (s *TestProRequest) SetTestField2(v string) *TestProRequest {
	s.TestField2 = &v
	return s
}

func (s *TestProRequest) SetTestFieldTwo(v string) *TestProRequest {
	s.TestFieldTwo = &v
	return s
}

func (s *TestProRequest) SetTestField4(v string) *TestProRequest {
	s.TestField4 = &v
	return s
}

func (s *TestProRequest) SetTestField5(v string) *TestProRequest {
	s.TestField5 = &v
	return s
}

func (s *TestProRequest) SetTestField6(v string) *TestProRequest {
	s.TestField6 = &v
	return s
}

func (s *TestProRequest) SetTestField8(v string) *TestProRequest {
	s.TestField8 = &v
	return s
}

func (s *TestProRequest) SetTestField3(v string) *TestProRequest {
	s.TestField3 = &v
	return s
}

func (s *TestProRequest) SetTestField7(v string) *TestProRequest {
	s.TestField7 = &v
	return s
}

type TestProResponseBody struct {
	Data    *TestProResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                  `json:"errCode,omitempty" xml:"errCode,omitempty"`
	Success *bool                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TestProResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TestProResponseBody) GoString() string {
	return s.String()
}

func (s *TestProResponseBody) SetData(v *TestProResponseBodyData) *TestProResponseBody {
	s.Data = v
	return s
}

func (s *TestProResponseBody) SetErrCode(v string) *TestProResponseBody {
	s.ErrCode = &v
	return s
}

func (s *TestProResponseBody) SetSuccess(v bool) *TestProResponseBody {
	s.Success = &v
	return s
}

type TestProResponseBodyData struct {
	AppVersion                   *string                                                `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	Category                     *string                                                `json:"category,omitempty" xml:"category,omitempty"`
	ChildrenComponentVersionList []*TestProResponseBodyDataChildrenComponentVersionList `json:"childrenComponentVersionList,omitempty" xml:"childrenComponentVersionList,omitempty" type:"Repeated"`
	Class                        *string                                                `json:"class,omitempty" xml:"class,omitempty"`
	ComponentName                *string                                                `json:"componentName,omitempty" xml:"componentName,omitempty"`
	ComponentUID                 *string                                                `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	Description                  *string                                                `json:"description,omitempty" xml:"description,omitempty"`
	Documents                    []*string                                              `json:"documents,omitempty" xml:"documents,omitempty" type:"Repeated"`
	Enable                       *bool                                                  `json:"enable,omitempty" xml:"enable,omitempty"`
	HasDependency                *bool                                                  `json:"hasDependency,omitempty" xml:"hasDependency,omitempty"`
	ImagesMapping                *string                                                `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	Namespace                    *string                                                `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OrchestrationValues          *string                                                `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	PackageURL                   *string                                                `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	ParentComponent              *bool                                                  `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	Priority                     *int32                                                 `json:"priority,omitempty" xml:"priority,omitempty"`
	ProductComponentVersionUID   *string                                                `json:"productComponentVersionUID,omitempty" xml:"productComponentVersionUID,omitempty"`
	Provider                     *string                                                `json:"provider,omitempty" xml:"provider,omitempty"`
	Readme                       *string                                                `json:"readme,omitempty" xml:"readme,omitempty"`
	Resources                    *string                                                `json:"resources,omitempty" xml:"resources,omitempty"`
	Singleton                    *bool                                                  `json:"singleton,omitempty" xml:"singleton,omitempty"`
	Uid                          *string                                                `json:"uid,omitempty" xml:"uid,omitempty"`
	Version                      *string                                                `json:"version,omitempty" xml:"version,omitempty"`
}

func (s TestProResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TestProResponseBodyData) GoString() string {
	return s.String()
}

func (s *TestProResponseBodyData) SetAppVersion(v string) *TestProResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *TestProResponseBodyData) SetCategory(v string) *TestProResponseBodyData {
	s.Category = &v
	return s
}

func (s *TestProResponseBodyData) SetChildrenComponentVersionList(v []*TestProResponseBodyDataChildrenComponentVersionList) *TestProResponseBodyData {
	s.ChildrenComponentVersionList = v
	return s
}

func (s *TestProResponseBodyData) SetClass(v string) *TestProResponseBodyData {
	s.Class = &v
	return s
}

func (s *TestProResponseBodyData) SetComponentName(v string) *TestProResponseBodyData {
	s.ComponentName = &v
	return s
}

func (s *TestProResponseBodyData) SetComponentUID(v string) *TestProResponseBodyData {
	s.ComponentUID = &v
	return s
}

func (s *TestProResponseBodyData) SetDescription(v string) *TestProResponseBodyData {
	s.Description = &v
	return s
}

func (s *TestProResponseBodyData) SetDocuments(v []*string) *TestProResponseBodyData {
	s.Documents = v
	return s
}

func (s *TestProResponseBodyData) SetEnable(v bool) *TestProResponseBodyData {
	s.Enable = &v
	return s
}

func (s *TestProResponseBodyData) SetHasDependency(v bool) *TestProResponseBodyData {
	s.HasDependency = &v
	return s
}

func (s *TestProResponseBodyData) SetImagesMapping(v string) *TestProResponseBodyData {
	s.ImagesMapping = &v
	return s
}

func (s *TestProResponseBodyData) SetNamespace(v string) *TestProResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *TestProResponseBodyData) SetOrchestrationValues(v string) *TestProResponseBodyData {
	s.OrchestrationValues = &v
	return s
}

func (s *TestProResponseBodyData) SetPackageURL(v string) *TestProResponseBodyData {
	s.PackageURL = &v
	return s
}

func (s *TestProResponseBodyData) SetParentComponent(v bool) *TestProResponseBodyData {
	s.ParentComponent = &v
	return s
}

func (s *TestProResponseBodyData) SetPriority(v int32) *TestProResponseBodyData {
	s.Priority = &v
	return s
}

func (s *TestProResponseBodyData) SetProductComponentVersionUID(v string) *TestProResponseBodyData {
	s.ProductComponentVersionUID = &v
	return s
}

func (s *TestProResponseBodyData) SetProvider(v string) *TestProResponseBodyData {
	s.Provider = &v
	return s
}

func (s *TestProResponseBodyData) SetReadme(v string) *TestProResponseBodyData {
	s.Readme = &v
	return s
}

func (s *TestProResponseBodyData) SetResources(v string) *TestProResponseBodyData {
	s.Resources = &v
	return s
}

func (s *TestProResponseBodyData) SetSingleton(v bool) *TestProResponseBodyData {
	s.Singleton = &v
	return s
}

func (s *TestProResponseBodyData) SetUid(v string) *TestProResponseBodyData {
	s.Uid = &v
	return s
}

func (s *TestProResponseBodyData) SetVersion(v string) *TestProResponseBodyData {
	s.Version = &v
	return s
}

type TestProResponseBodyDataChildrenComponentVersionList struct {
	AppVersion                 *string   `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	Category                   *string   `json:"category,omitempty" xml:"category,omitempty"`
	Class                      *string   `json:"class,omitempty" xml:"class,omitempty"`
	ComponentName              *string   `json:"componentName,omitempty" xml:"componentName,omitempty"`
	ComponentUID               *string   `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	Description                *string   `json:"description,omitempty" xml:"description,omitempty"`
	Documents                  []*string `json:"documents,omitempty" xml:"documents,omitempty" type:"Repeated"`
	Enable                     *bool     `json:"enable,omitempty" xml:"enable,omitempty"`
	ImagesMapping              *string   `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	Namespace                  *string   `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OrchestrationValues        *string   `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	PackageURL                 *string   `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	ParentComponent            *bool     `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	Priority                   *int32    `json:"priority,omitempty" xml:"priority,omitempty"`
	ProductComponentVersionUID *string   `json:"productComponentVersionUID,omitempty" xml:"productComponentVersionUID,omitempty"`
	Provider                   *string   `json:"provider,omitempty" xml:"provider,omitempty"`
	Readme                     *string   `json:"readme,omitempty" xml:"readme,omitempty"`
	Resources                  *string   `json:"resources,omitempty" xml:"resources,omitempty"`
	Singleton                  *bool     `json:"singleton,omitempty" xml:"singleton,omitempty"`
	Uid                        *string   `json:"uid,omitempty" xml:"uid,omitempty"`
	Version                    *string   `json:"version,omitempty" xml:"version,omitempty"`
}

func (s TestProResponseBodyDataChildrenComponentVersionList) String() string {
	return tea.Prettify(s)
}

func (s TestProResponseBodyDataChildrenComponentVersionList) GoString() string {
	return s.String()
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetAppVersion(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.AppVersion = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetCategory(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Category = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetClass(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Class = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetComponentName(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.ComponentName = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetComponentUID(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.ComponentUID = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetDescription(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Description = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetDocuments(v []*string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Documents = v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetEnable(v bool) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Enable = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetImagesMapping(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.ImagesMapping = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetNamespace(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Namespace = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetOrchestrationValues(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.OrchestrationValues = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetPackageURL(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.PackageURL = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetParentComponent(v bool) *TestProResponseBodyDataChildrenComponentVersionList {
	s.ParentComponent = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetPriority(v int32) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Priority = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetProductComponentVersionUID(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.ProductComponentVersionUID = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetProvider(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Provider = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetReadme(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Readme = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetResources(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Resources = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetSingleton(v bool) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Singleton = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetUid(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Uid = &v
	return s
}

func (s *TestProResponseBodyDataChildrenComponentVersionList) SetVersion(v string) *TestProResponseBodyDataChildrenComponentVersionList {
	s.Version = &v
	return s
}

type TestProResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TestProResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TestProResponse) String() string {
	return tea.Prettify(s)
}

func (s TestProResponse) GoString() string {
	return s.String()
}

func (s *TestProResponse) SetHeaders(v map[string]*string) *TestProResponse {
	s.Headers = v
	return s
}

func (s *TestProResponse) SetStatusCode(v int32) *TestProResponse {
	s.StatusCode = &v
	return s
}

func (s *TestProResponse) SetBody(v *TestProResponseBody) *TestProResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("lmztest"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetProductComponentDetailWithOptions(request *GetProductComponentDetailRequest, runtime *util.RuntimeOptions) (_result *GetProductComponentDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProductComponentDetail"),
		Version:     tea.String("2010-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductComponentDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductComponentDetail(request *GetProductComponentDetailRequest) (_result *GetProductComponentDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProductComponentDetailResponse{}
	_body, _err := client.GetProductComponentDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QkTestWithOptions(runtime *util.RuntimeOptions) (_result *QkTestResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QkTest"),
		Version:     tea.String("2010-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QkTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QkTest() (_result *QkTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QkTestResponse{}
	_body, _err := client.QkTestWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TestProWithOptions(request *TestProRequest, runtime *util.RuntimeOptions) (_result *TestProResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TestPro"),
		Version:     tea.String("2010-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TestProResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TestPro(request *TestProRequest) (_result *TestProResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TestProResponse{}
	_body, _err := client.TestProWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
