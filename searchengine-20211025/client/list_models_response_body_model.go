// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListModelsResponseBody
	GetRequestId() *string
	SetResult(v []*ListModelsResponseBodyResult) *ListModelsResponseBody
	GetResult() []*ListModelsResponseBodyResult
	SetTotalCount(v int32) *ListModelsResponseBody
	GetTotalCount() *int32
}

type ListModelsResponseBody struct {
	// example:
	//
	// 38b079f1-7846-4226-8c90-3e2644b5c52b
	RequestId *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListModelsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 14
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelsResponseBody) GetResult() []*ListModelsResponseBodyResult {
	return s.Result
}

func (s *ListModelsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListModelsResponseBody) SetRequestId(v string) *ListModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelsResponseBody) SetResult(v []*ListModelsResponseBodyResult) *ListModelsResponseBody {
	s.Result = v
	return s
}

func (s *ListModelsResponseBody) SetTotalCount(v int32) *ListModelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListModelsResponseBodyResult struct {
	Content *ListModelsResponseBodyResultContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// example:
	//
	// 2024-05-21 16:05:26
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 128
	Dimension *int32 `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ok
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// text_embedding
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 2024-05-21 16:05:26
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// http://***.platform-cn-shanghai.opensearch.aliyuncs.com/v3/openapi/workspaces/default/text-embedding/ops-text-embedding-001
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListModelsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyResult) GetContent() *ListModelsResponseBodyResultContent {
	return s.Content
}

func (s *ListModelsResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListModelsResponseBodyResult) GetDimension() *int32 {
	return s.Dimension
}

func (s *ListModelsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListModelsResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListModelsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListModelsResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListModelsResponseBodyResult) GetUrl() *string {
	return s.Url
}

func (s *ListModelsResponseBodyResult) SetContent(v *ListModelsResponseBodyResultContent) *ListModelsResponseBodyResult {
	s.Content = v
	return s
}

func (s *ListModelsResponseBodyResult) SetCreateTime(v string) *ListModelsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListModelsResponseBodyResult) SetDimension(v int32) *ListModelsResponseBodyResult {
	s.Dimension = &v
	return s
}

func (s *ListModelsResponseBodyResult) SetName(v string) *ListModelsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListModelsResponseBodyResult) SetStatus(v string) *ListModelsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListModelsResponseBodyResult) SetType(v string) *ListModelsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListModelsResponseBodyResult) SetUpdateTime(v string) *ListModelsResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *ListModelsResponseBodyResult) SetUrl(v string) *ListModelsResponseBodyResult {
	s.Url = &v
	return s
}

func (s *ListModelsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListModelsResponseBodyResultContent struct {
	// example:
	//
	// POST
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// example:
	//
	// test
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// example:
	//
	// text_embedding
	ModelType *string                                      `json:"modelType,omitempty" xml:"modelType,omitempty"`
	Request   *ListModelsResponseBodyResultContentRequest  `json:"request,omitempty" xml:"request,omitempty" type:"Struct"`
	Response  *ListModelsResponseBodyResultContentResponse `json:"response,omitempty" xml:"response,omitempty" type:"Struct"`
	// example:
	//
	// http://***.platform-cn-shanghai.opensearch.aliyuncs.com/v3/openapi/workspaces/default/text-embedding/ops-text-embedding-001
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListModelsResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyResultContent) GetMethod() *string {
	return s.Method
}

func (s *ListModelsResponseBodyResultContent) GetModelName() *string {
	return s.ModelName
}

func (s *ListModelsResponseBodyResultContent) GetModelType() *string {
	return s.ModelType
}

func (s *ListModelsResponseBodyResultContent) GetRequest() *ListModelsResponseBodyResultContentRequest {
	return s.Request
}

func (s *ListModelsResponseBodyResultContent) GetResponse() *ListModelsResponseBodyResultContentResponse {
	return s.Response
}

func (s *ListModelsResponseBodyResultContent) GetUrl() *string {
	return s.Url
}

func (s *ListModelsResponseBodyResultContent) SetMethod(v string) *ListModelsResponseBodyResultContent {
	s.Method = &v
	return s
}

func (s *ListModelsResponseBodyResultContent) SetModelName(v string) *ListModelsResponseBodyResultContent {
	s.ModelName = &v
	return s
}

func (s *ListModelsResponseBodyResultContent) SetModelType(v string) *ListModelsResponseBodyResultContent {
	s.ModelType = &v
	return s
}

func (s *ListModelsResponseBodyResultContent) SetRequest(v *ListModelsResponseBodyResultContentRequest) *ListModelsResponseBodyResultContent {
	s.Request = v
	return s
}

func (s *ListModelsResponseBodyResultContent) SetResponse(v *ListModelsResponseBodyResultContentResponse) *ListModelsResponseBodyResultContent {
	s.Response = v
	return s
}

func (s *ListModelsResponseBodyResultContent) SetUrl(v string) *ListModelsResponseBodyResultContent {
	s.Url = &v
	return s
}

func (s *ListModelsResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type ListModelsResponseBodyResultContentRequest struct {
	Header     *ListModelsResponseBodyResultContentRequestHeader     `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Parameters *ListModelsResponseBodyResultContentRequestParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	// example:
	//
	// {\\"input\\": [\\"%{input}\\"], \\"input_type\\": \\"%{input_type}\\"}
	RequestBody *string                                              `json:"requestBody,omitempty" xml:"requestBody,omitempty"`
	UrlParams   *ListModelsResponseBodyResultContentRequestUrlParams `json:"urlParams,omitempty" xml:"urlParams,omitempty" type:"Struct"`
}

func (s ListModelsResponseBodyResultContentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyResultContentRequest) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyResultContentRequest) GetHeader() *ListModelsResponseBodyResultContentRequestHeader {
	return s.Header
}

func (s *ListModelsResponseBodyResultContentRequest) GetParameters() *ListModelsResponseBodyResultContentRequestParameters {
	return s.Parameters
}

func (s *ListModelsResponseBodyResultContentRequest) GetRequestBody() *string {
	return s.RequestBody
}

func (s *ListModelsResponseBodyResultContentRequest) GetUrlParams() *ListModelsResponseBodyResultContentRequestUrlParams {
	return s.UrlParams
}

func (s *ListModelsResponseBodyResultContentRequest) SetHeader(v *ListModelsResponseBodyResultContentRequestHeader) *ListModelsResponseBodyResultContentRequest {
	s.Header = v
	return s
}

func (s *ListModelsResponseBodyResultContentRequest) SetParameters(v *ListModelsResponseBodyResultContentRequestParameters) *ListModelsResponseBodyResultContentRequest {
	s.Parameters = v
	return s
}

func (s *ListModelsResponseBodyResultContentRequest) SetRequestBody(v string) *ListModelsResponseBodyResultContentRequest {
	s.RequestBody = &v
	return s
}

func (s *ListModelsResponseBodyResultContentRequest) SetUrlParams(v *ListModelsResponseBodyResultContentRequestUrlParams) *ListModelsResponseBodyResultContentRequest {
	s.UrlParams = v
	return s
}

func (s *ListModelsResponseBodyResultContentRequest) Validate() error {
	return dara.Validate(s)
}

type ListModelsResponseBodyResultContentRequestHeader struct {
	// example:
	//
	// Bearer OS-v0********6vvs
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	// example:
	//
	// application/json
	ContentType *string `json:"Content-Type,omitempty" xml:"Content-Type,omitempty"`
}

func (s ListModelsResponseBodyResultContentRequestHeader) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyResultContentRequestHeader) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyResultContentRequestHeader) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListModelsResponseBodyResultContentRequestHeader) GetContentType() *string {
	return s.ContentType
}

func (s *ListModelsResponseBodyResultContentRequestHeader) SetAuthorization(v string) *ListModelsResponseBodyResultContentRequestHeader {
	s.Authorization = &v
	return s
}

func (s *ListModelsResponseBodyResultContentRequestHeader) SetContentType(v string) *ListModelsResponseBodyResultContentRequestHeader {
	s.ContentType = &v
	return s
}

func (s *ListModelsResponseBodyResultContentRequestHeader) Validate() error {
	return dara.Validate(s)
}

type ListModelsResponseBodyResultContentRequestParameters struct {
	Build  *ListModelsResponseBodyResultContentRequestParametersBuild  `json:"build,omitempty" xml:"build,omitempty" type:"Struct"`
	Search *ListModelsResponseBodyResultContentRequestParametersSearch `json:"search,omitempty" xml:"search,omitempty" type:"Struct"`
}

func (s ListModelsResponseBodyResultContentRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyResultContentRequestParameters) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyResultContentRequestParameters) GetBuild() *ListModelsResponseBodyResultContentRequestParametersBuild {
	return s.Build
}

func (s *ListModelsResponseBodyResultContentRequestParameters) GetSearch() *ListModelsResponseBodyResultContentRequestParametersSearch {
	return s.Search
}

func (s *ListModelsResponseBodyResultContentRequestParameters) SetBuild(v *ListModelsResponseBodyResultContentRequestParametersBuild) *ListModelsResponseBodyResultContentRequestParameters {
	s.Build = v
	return s
}

func (s *ListModelsResponseBodyResultContentRequestParameters) SetSearch(v *ListModelsResponseBodyResultContentRequestParametersSearch) *ListModelsResponseBodyResultContentRequestParameters {
	s.Search = v
	return s
}

func (s *ListModelsResponseBodyResultContentRequestParameters) Validate() error {
	return dara.Validate(s)
}

type ListModelsResponseBodyResultContentRequestParametersBuild struct {
	// example:
	//
	// query
	InputType *string `json:"input_type,omitempty" xml:"input_type,omitempty"`
}

func (s ListModelsResponseBodyResultContentRequestParametersBuild) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyResultContentRequestParametersBuild) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyResultContentRequestParametersBuild) GetInputType() *string {
	return s.InputType
}

func (s *ListModelsResponseBodyResultContentRequestParametersBuild) SetInputType(v string) *ListModelsResponseBodyResultContentRequestParametersBuild {
	s.InputType = &v
	return s
}

func (s *ListModelsResponseBodyResultContentRequestParametersBuild) Validate() error {
	return dara.Validate(s)
}

type ListModelsResponseBodyResultContentRequestParametersSearch struct {
	// example:
	//
	// document
	InputType *string `json:"input_type,omitempty" xml:"input_type,omitempty"`
}

func (s ListModelsResponseBodyResultContentRequestParametersSearch) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyResultContentRequestParametersSearch) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyResultContentRequestParametersSearch) GetInputType() *string {
	return s.InputType
}

func (s *ListModelsResponseBodyResultContentRequestParametersSearch) SetInputType(v string) *ListModelsResponseBodyResultContentRequestParametersSearch {
	s.InputType = &v
	return s
}

func (s *ListModelsResponseBodyResultContentRequestParametersSearch) Validate() error {
	return dara.Validate(s)
}

type ListModelsResponseBodyResultContentRequestUrlParams struct {
	// example:
	//
	// key: value
	Build map[string]interface{} `json:"build,omitempty" xml:"build,omitempty"`
	// example:
	//
	// key: value
	Search map[string]interface{} `json:"search,omitempty" xml:"search,omitempty"`
}

func (s ListModelsResponseBodyResultContentRequestUrlParams) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyResultContentRequestUrlParams) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyResultContentRequestUrlParams) GetBuild() map[string]interface{} {
	return s.Build
}

func (s *ListModelsResponseBodyResultContentRequestUrlParams) GetSearch() map[string]interface{} {
	return s.Search
}

func (s *ListModelsResponseBodyResultContentRequestUrlParams) SetBuild(v map[string]interface{}) *ListModelsResponseBodyResultContentRequestUrlParams {
	s.Build = v
	return s
}

func (s *ListModelsResponseBodyResultContentRequestUrlParams) SetSearch(v map[string]interface{}) *ListModelsResponseBodyResultContentRequestUrlParams {
	s.Search = v
	return s
}

func (s *ListModelsResponseBodyResultContentRequestUrlParams) Validate() error {
	return dara.Validate(s)
}

type ListModelsResponseBodyResultContentResponse struct {
	// example:
	//
	// $.result.embeddings[*].embedding
	Embeddings *string `json:"embeddings,omitempty" xml:"embeddings,omitempty"`
}

func (s ListModelsResponseBodyResultContentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyResultContentResponse) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyResultContentResponse) GetEmbeddings() *string {
	return s.Embeddings
}

func (s *ListModelsResponseBodyResultContentResponse) SetEmbeddings(v string) *ListModelsResponseBodyResultContentResponse {
	s.Embeddings = &v
	return s
}

func (s *ListModelsResponseBodyResultContentResponse) Validate() error {
	return dara.Validate(s)
}
