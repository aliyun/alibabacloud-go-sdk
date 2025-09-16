// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetModelResponseBody
	GetRequestId() *string
	SetResult(v *GetModelResponseBodyResult) *GetModelResponseBody
	GetResult() *GetModelResponseBodyResult
}

type GetModelResponseBody struct {
	// example:
	//
	// 38b079f1-7846-4226-8c90-3e2644b5c52b
	RequestId *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetModelResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelResponseBody) GetResult() *GetModelResponseBodyResult {
	return s.Result
}

func (s *GetModelResponseBody) SetRequestId(v string) *GetModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelResponseBody) SetResult(v *GetModelResponseBodyResult) *GetModelResponseBody {
	s.Result = v
	return s
}

func (s *GetModelResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetModelResponseBodyResult struct {
	Content *GetModelResponseBodyResultContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
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

func (s GetModelResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetModelResponseBodyResult) GetContent() *GetModelResponseBodyResultContent {
	return s.Content
}

func (s *GetModelResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetModelResponseBodyResult) GetDimension() *int32 {
	return s.Dimension
}

func (s *GetModelResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetModelResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetModelResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetModelResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetModelResponseBodyResult) GetUrl() *string {
	return s.Url
}

func (s *GetModelResponseBodyResult) SetContent(v *GetModelResponseBodyResultContent) *GetModelResponseBodyResult {
	s.Content = v
	return s
}

func (s *GetModelResponseBodyResult) SetCreateTime(v string) *GetModelResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetModelResponseBodyResult) SetDimension(v int32) *GetModelResponseBodyResult {
	s.Dimension = &v
	return s
}

func (s *GetModelResponseBodyResult) SetName(v string) *GetModelResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetModelResponseBodyResult) SetStatus(v string) *GetModelResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetModelResponseBodyResult) SetType(v string) *GetModelResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetModelResponseBodyResult) SetUpdateTime(v string) *GetModelResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetModelResponseBodyResult) SetUrl(v string) *GetModelResponseBodyResult {
	s.Url = &v
	return s
}

func (s *GetModelResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetModelResponseBodyResultContent struct {
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
	ModelType *string                                    `json:"modelType,omitempty" xml:"modelType,omitempty"`
	Request   *GetModelResponseBodyResultContentRequest  `json:"request,omitempty" xml:"request,omitempty" type:"Struct"`
	Response  *GetModelResponseBodyResultContentResponse `json:"response,omitempty" xml:"response,omitempty" type:"Struct"`
	// example:
	//
	// http://***.platform-cn-shanghai.opensearch.aliyuncs.com/v3/openapi/workspaces/default/text-embedding/ops-text-embedding-001
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetModelResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *GetModelResponseBodyResultContent) GetMethod() *string {
	return s.Method
}

func (s *GetModelResponseBodyResultContent) GetModelName() *string {
	return s.ModelName
}

func (s *GetModelResponseBodyResultContent) GetModelType() *string {
	return s.ModelType
}

func (s *GetModelResponseBodyResultContent) GetRequest() *GetModelResponseBodyResultContentRequest {
	return s.Request
}

func (s *GetModelResponseBodyResultContent) GetResponse() *GetModelResponseBodyResultContentResponse {
	return s.Response
}

func (s *GetModelResponseBodyResultContent) GetUrl() *string {
	return s.Url
}

func (s *GetModelResponseBodyResultContent) SetMethod(v string) *GetModelResponseBodyResultContent {
	s.Method = &v
	return s
}

func (s *GetModelResponseBodyResultContent) SetModelName(v string) *GetModelResponseBodyResultContent {
	s.ModelName = &v
	return s
}

func (s *GetModelResponseBodyResultContent) SetModelType(v string) *GetModelResponseBodyResultContent {
	s.ModelType = &v
	return s
}

func (s *GetModelResponseBodyResultContent) SetRequest(v *GetModelResponseBodyResultContentRequest) *GetModelResponseBodyResultContent {
	s.Request = v
	return s
}

func (s *GetModelResponseBodyResultContent) SetResponse(v *GetModelResponseBodyResultContentResponse) *GetModelResponseBodyResultContent {
	s.Response = v
	return s
}

func (s *GetModelResponseBodyResultContent) SetUrl(v string) *GetModelResponseBodyResultContent {
	s.Url = &v
	return s
}

func (s *GetModelResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type GetModelResponseBodyResultContentRequest struct {
	Header     *GetModelResponseBodyResultContentRequestHeader     `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Parameters *GetModelResponseBodyResultContentRequestParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	// example:
	//
	// {\\"input\\": [\\"%{input}\\"], \\"input_type\\": \\"%{input_type}\\"}
	RequestBody *string                                            `json:"requestBody,omitempty" xml:"requestBody,omitempty"`
	UrlParams   *GetModelResponseBodyResultContentRequestUrlParams `json:"urlParams,omitempty" xml:"urlParams,omitempty" type:"Struct"`
}

func (s GetModelResponseBodyResultContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBodyResultContentRequest) GoString() string {
	return s.String()
}

func (s *GetModelResponseBodyResultContentRequest) GetHeader() *GetModelResponseBodyResultContentRequestHeader {
	return s.Header
}

func (s *GetModelResponseBodyResultContentRequest) GetParameters() *GetModelResponseBodyResultContentRequestParameters {
	return s.Parameters
}

func (s *GetModelResponseBodyResultContentRequest) GetRequestBody() *string {
	return s.RequestBody
}

func (s *GetModelResponseBodyResultContentRequest) GetUrlParams() *GetModelResponseBodyResultContentRequestUrlParams {
	return s.UrlParams
}

func (s *GetModelResponseBodyResultContentRequest) SetHeader(v *GetModelResponseBodyResultContentRequestHeader) *GetModelResponseBodyResultContentRequest {
	s.Header = v
	return s
}

func (s *GetModelResponseBodyResultContentRequest) SetParameters(v *GetModelResponseBodyResultContentRequestParameters) *GetModelResponseBodyResultContentRequest {
	s.Parameters = v
	return s
}

func (s *GetModelResponseBodyResultContentRequest) SetRequestBody(v string) *GetModelResponseBodyResultContentRequest {
	s.RequestBody = &v
	return s
}

func (s *GetModelResponseBodyResultContentRequest) SetUrlParams(v *GetModelResponseBodyResultContentRequestUrlParams) *GetModelResponseBodyResultContentRequest {
	s.UrlParams = v
	return s
}

func (s *GetModelResponseBodyResultContentRequest) Validate() error {
	return dara.Validate(s)
}

type GetModelResponseBodyResultContentRequestHeader struct {
	// example:
	//
	// Bearer OS-v0********6vvs
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	// example:
	//
	// application/json
	ContentType *string `json:"Content-Type,omitempty" xml:"Content-Type,omitempty"`
}

func (s GetModelResponseBodyResultContentRequestHeader) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBodyResultContentRequestHeader) GoString() string {
	return s.String()
}

func (s *GetModelResponseBodyResultContentRequestHeader) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetModelResponseBodyResultContentRequestHeader) GetContentType() *string {
	return s.ContentType
}

func (s *GetModelResponseBodyResultContentRequestHeader) SetAuthorization(v string) *GetModelResponseBodyResultContentRequestHeader {
	s.Authorization = &v
	return s
}

func (s *GetModelResponseBodyResultContentRequestHeader) SetContentType(v string) *GetModelResponseBodyResultContentRequestHeader {
	s.ContentType = &v
	return s
}

func (s *GetModelResponseBodyResultContentRequestHeader) Validate() error {
	return dara.Validate(s)
}

type GetModelResponseBodyResultContentRequestParameters struct {
	Build  *GetModelResponseBodyResultContentRequestParametersBuild  `json:"build,omitempty" xml:"build,omitempty" type:"Struct"`
	Search *GetModelResponseBodyResultContentRequestParametersSearch `json:"search,omitempty" xml:"search,omitempty" type:"Struct"`
}

func (s GetModelResponseBodyResultContentRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBodyResultContentRequestParameters) GoString() string {
	return s.String()
}

func (s *GetModelResponseBodyResultContentRequestParameters) GetBuild() *GetModelResponseBodyResultContentRequestParametersBuild {
	return s.Build
}

func (s *GetModelResponseBodyResultContentRequestParameters) GetSearch() *GetModelResponseBodyResultContentRequestParametersSearch {
	return s.Search
}

func (s *GetModelResponseBodyResultContentRequestParameters) SetBuild(v *GetModelResponseBodyResultContentRequestParametersBuild) *GetModelResponseBodyResultContentRequestParameters {
	s.Build = v
	return s
}

func (s *GetModelResponseBodyResultContentRequestParameters) SetSearch(v *GetModelResponseBodyResultContentRequestParametersSearch) *GetModelResponseBodyResultContentRequestParameters {
	s.Search = v
	return s
}

func (s *GetModelResponseBodyResultContentRequestParameters) Validate() error {
	return dara.Validate(s)
}

type GetModelResponseBodyResultContentRequestParametersBuild struct {
	// example:
	//
	// query
	InputType *string `json:"input_type,omitempty" xml:"input_type,omitempty"`
}

func (s GetModelResponseBodyResultContentRequestParametersBuild) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBodyResultContentRequestParametersBuild) GoString() string {
	return s.String()
}

func (s *GetModelResponseBodyResultContentRequestParametersBuild) GetInputType() *string {
	return s.InputType
}

func (s *GetModelResponseBodyResultContentRequestParametersBuild) SetInputType(v string) *GetModelResponseBodyResultContentRequestParametersBuild {
	s.InputType = &v
	return s
}

func (s *GetModelResponseBodyResultContentRequestParametersBuild) Validate() error {
	return dara.Validate(s)
}

type GetModelResponseBodyResultContentRequestParametersSearch struct {
	// example:
	//
	// document
	InputType *string `json:"input_type,omitempty" xml:"input_type,omitempty"`
}

func (s GetModelResponseBodyResultContentRequestParametersSearch) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBodyResultContentRequestParametersSearch) GoString() string {
	return s.String()
}

func (s *GetModelResponseBodyResultContentRequestParametersSearch) GetInputType() *string {
	return s.InputType
}

func (s *GetModelResponseBodyResultContentRequestParametersSearch) SetInputType(v string) *GetModelResponseBodyResultContentRequestParametersSearch {
	s.InputType = &v
	return s
}

func (s *GetModelResponseBodyResultContentRequestParametersSearch) Validate() error {
	return dara.Validate(s)
}

type GetModelResponseBodyResultContentRequestUrlParams struct {
	// example:
	//
	// key: value
	Build map[string]interface{} `json:"build,omitempty" xml:"build,omitempty"`
	// example:
	//
	// key: value
	Search map[string]interface{} `json:"search,omitempty" xml:"search,omitempty"`
}

func (s GetModelResponseBodyResultContentRequestUrlParams) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBodyResultContentRequestUrlParams) GoString() string {
	return s.String()
}

func (s *GetModelResponseBodyResultContentRequestUrlParams) GetBuild() map[string]interface{} {
	return s.Build
}

func (s *GetModelResponseBodyResultContentRequestUrlParams) GetSearch() map[string]interface{} {
	return s.Search
}

func (s *GetModelResponseBodyResultContentRequestUrlParams) SetBuild(v map[string]interface{}) *GetModelResponseBodyResultContentRequestUrlParams {
	s.Build = v
	return s
}

func (s *GetModelResponseBodyResultContentRequestUrlParams) SetSearch(v map[string]interface{}) *GetModelResponseBodyResultContentRequestUrlParams {
	s.Search = v
	return s
}

func (s *GetModelResponseBodyResultContentRequestUrlParams) Validate() error {
	return dara.Validate(s)
}

type GetModelResponseBodyResultContentResponse struct {
	// example:
	//
	// $.result.embeddings[*].embedding
	Embeddings *string `json:"embeddings,omitempty" xml:"embeddings,omitempty"`
}

func (s GetModelResponseBodyResultContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBodyResultContentResponse) GoString() string {
	return s.String()
}

func (s *GetModelResponseBodyResultContentResponse) GetEmbeddings() *string {
	return s.Embeddings
}

func (s *GetModelResponseBodyResultContentResponse) SetEmbeddings(v string) *GetModelResponseBodyResultContentResponse {
	s.Embeddings = &v
	return s
}

func (s *GetModelResponseBodyResultContentResponse) Validate() error {
	return dara.Validate(s)
}
