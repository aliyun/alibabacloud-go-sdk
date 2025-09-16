// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *CreateModelRequestContent) *CreateModelRequest
	GetContent() *CreateModelRequestContent
	SetName(v string) *CreateModelRequest
	GetName() *string
	SetDryRun(v string) *CreateModelRequest
	GetDryRun() *string
}

type CreateModelRequest struct {
	Content *CreateModelRequestContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	DryRun *string `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequest) GoString() string {
	return s.String()
}

func (s *CreateModelRequest) GetContent() *CreateModelRequestContent {
	return s.Content
}

func (s *CreateModelRequest) GetName() *string {
	return s.Name
}

func (s *CreateModelRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *CreateModelRequest) SetContent(v *CreateModelRequestContent) *CreateModelRequest {
	s.Content = v
	return s
}

func (s *CreateModelRequest) SetName(v string) *CreateModelRequest {
	s.Name = &v
	return s
}

func (s *CreateModelRequest) SetDryRun(v string) *CreateModelRequest {
	s.DryRun = &v
	return s
}

func (s *CreateModelRequest) Validate() error {
	return dara.Validate(s)
}

type CreateModelRequestContent struct {
	// example:
	//
	// 128
	Dimension *int32 `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// example:
	//
	// POST
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// example:
	//
	// text_embedding
	ModelType *string                            `json:"modelType,omitempty" xml:"modelType,omitempty"`
	Request   *CreateModelRequestContentRequest  `json:"request,omitempty" xml:"request,omitempty" type:"Struct"`
	Response  *CreateModelRequestContentResponse `json:"response,omitempty" xml:"response,omitempty" type:"Struct"`
	// example:
	//
	// http://***.platform-cn-shanghai.opensearch.aliyuncs.com/v3/openapi/workspaces/default/text-embedding/ops-text-embedding-001
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateModelRequestContent) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequestContent) GoString() string {
	return s.String()
}

func (s *CreateModelRequestContent) GetDimension() *int32 {
	return s.Dimension
}

func (s *CreateModelRequestContent) GetMethod() *string {
	return s.Method
}

func (s *CreateModelRequestContent) GetModelType() *string {
	return s.ModelType
}

func (s *CreateModelRequestContent) GetRequest() *CreateModelRequestContentRequest {
	return s.Request
}

func (s *CreateModelRequestContent) GetResponse() *CreateModelRequestContentResponse {
	return s.Response
}

func (s *CreateModelRequestContent) GetUrl() *string {
	return s.Url
}

func (s *CreateModelRequestContent) SetDimension(v int32) *CreateModelRequestContent {
	s.Dimension = &v
	return s
}

func (s *CreateModelRequestContent) SetMethod(v string) *CreateModelRequestContent {
	s.Method = &v
	return s
}

func (s *CreateModelRequestContent) SetModelType(v string) *CreateModelRequestContent {
	s.ModelType = &v
	return s
}

func (s *CreateModelRequestContent) SetRequest(v *CreateModelRequestContentRequest) *CreateModelRequestContent {
	s.Request = v
	return s
}

func (s *CreateModelRequestContent) SetResponse(v *CreateModelRequestContentResponse) *CreateModelRequestContent {
	s.Response = v
	return s
}

func (s *CreateModelRequestContent) SetUrl(v string) *CreateModelRequestContent {
	s.Url = &v
	return s
}

func (s *CreateModelRequestContent) Validate() error {
	return dara.Validate(s)
}

type CreateModelRequestContentRequest struct {
	Header     *CreateModelRequestContentRequestHeader     `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Parameters *CreateModelRequestContentRequestParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	// example:
	//
	// {\\"input\\": [\\"%{input}\\"], \\"input_type\\": \\"%{input_type}\\"}
	RequestBody *string                                    `json:"requestBody,omitempty" xml:"requestBody,omitempty"`
	UrlParams   *CreateModelRequestContentRequestUrlParams `json:"urlParams,omitempty" xml:"urlParams,omitempty" type:"Struct"`
}

func (s CreateModelRequestContentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequestContentRequest) GoString() string {
	return s.String()
}

func (s *CreateModelRequestContentRequest) GetHeader() *CreateModelRequestContentRequestHeader {
	return s.Header
}

func (s *CreateModelRequestContentRequest) GetParameters() *CreateModelRequestContentRequestParameters {
	return s.Parameters
}

func (s *CreateModelRequestContentRequest) GetRequestBody() *string {
	return s.RequestBody
}

func (s *CreateModelRequestContentRequest) GetUrlParams() *CreateModelRequestContentRequestUrlParams {
	return s.UrlParams
}

func (s *CreateModelRequestContentRequest) SetHeader(v *CreateModelRequestContentRequestHeader) *CreateModelRequestContentRequest {
	s.Header = v
	return s
}

func (s *CreateModelRequestContentRequest) SetParameters(v *CreateModelRequestContentRequestParameters) *CreateModelRequestContentRequest {
	s.Parameters = v
	return s
}

func (s *CreateModelRequestContentRequest) SetRequestBody(v string) *CreateModelRequestContentRequest {
	s.RequestBody = &v
	return s
}

func (s *CreateModelRequestContentRequest) SetUrlParams(v *CreateModelRequestContentRequestUrlParams) *CreateModelRequestContentRequest {
	s.UrlParams = v
	return s
}

func (s *CreateModelRequestContentRequest) Validate() error {
	return dara.Validate(s)
}

type CreateModelRequestContentRequestHeader struct {
	// example:
	//
	// Bearer OS-v0********6vvs
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	// example:
	//
	// application/json
	ContentType *string `json:"Content-Type,omitempty" xml:"Content-Type,omitempty"`
}

func (s CreateModelRequestContentRequestHeader) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequestContentRequestHeader) GoString() string {
	return s.String()
}

func (s *CreateModelRequestContentRequestHeader) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateModelRequestContentRequestHeader) GetContentType() *string {
	return s.ContentType
}

func (s *CreateModelRequestContentRequestHeader) SetAuthorization(v string) *CreateModelRequestContentRequestHeader {
	s.Authorization = &v
	return s
}

func (s *CreateModelRequestContentRequestHeader) SetContentType(v string) *CreateModelRequestContentRequestHeader {
	s.ContentType = &v
	return s
}

func (s *CreateModelRequestContentRequestHeader) Validate() error {
	return dara.Validate(s)
}

type CreateModelRequestContentRequestParameters struct {
	Build  *CreateModelRequestContentRequestParametersBuild  `json:"build,omitempty" xml:"build,omitempty" type:"Struct"`
	Search *CreateModelRequestContentRequestParametersSearch `json:"search,omitempty" xml:"search,omitempty" type:"Struct"`
}

func (s CreateModelRequestContentRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequestContentRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateModelRequestContentRequestParameters) GetBuild() *CreateModelRequestContentRequestParametersBuild {
	return s.Build
}

func (s *CreateModelRequestContentRequestParameters) GetSearch() *CreateModelRequestContentRequestParametersSearch {
	return s.Search
}

func (s *CreateModelRequestContentRequestParameters) SetBuild(v *CreateModelRequestContentRequestParametersBuild) *CreateModelRequestContentRequestParameters {
	s.Build = v
	return s
}

func (s *CreateModelRequestContentRequestParameters) SetSearch(v *CreateModelRequestContentRequestParametersSearch) *CreateModelRequestContentRequestParameters {
	s.Search = v
	return s
}

func (s *CreateModelRequestContentRequestParameters) Validate() error {
	return dara.Validate(s)
}

type CreateModelRequestContentRequestParametersBuild struct {
	// example:
	//
	// query
	InputType *string `json:"input_type,omitempty" xml:"input_type,omitempty"`
}

func (s CreateModelRequestContentRequestParametersBuild) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequestContentRequestParametersBuild) GoString() string {
	return s.String()
}

func (s *CreateModelRequestContentRequestParametersBuild) GetInputType() *string {
	return s.InputType
}

func (s *CreateModelRequestContentRequestParametersBuild) SetInputType(v string) *CreateModelRequestContentRequestParametersBuild {
	s.InputType = &v
	return s
}

func (s *CreateModelRequestContentRequestParametersBuild) Validate() error {
	return dara.Validate(s)
}

type CreateModelRequestContentRequestParametersSearch struct {
	// example:
	//
	// document
	InputType *string `json:"input_type,omitempty" xml:"input_type,omitempty"`
}

func (s CreateModelRequestContentRequestParametersSearch) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequestContentRequestParametersSearch) GoString() string {
	return s.String()
}

func (s *CreateModelRequestContentRequestParametersSearch) GetInputType() *string {
	return s.InputType
}

func (s *CreateModelRequestContentRequestParametersSearch) SetInputType(v string) *CreateModelRequestContentRequestParametersSearch {
	s.InputType = &v
	return s
}

func (s *CreateModelRequestContentRequestParametersSearch) Validate() error {
	return dara.Validate(s)
}

type CreateModelRequestContentRequestUrlParams struct {
	// example:
	//
	// key: value
	Build map[string]interface{} `json:"build,omitempty" xml:"build,omitempty"`
	// example:
	//
	// key: value
	Search map[string]interface{} `json:"search,omitempty" xml:"search,omitempty"`
}

func (s CreateModelRequestContentRequestUrlParams) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequestContentRequestUrlParams) GoString() string {
	return s.String()
}

func (s *CreateModelRequestContentRequestUrlParams) GetBuild() map[string]interface{} {
	return s.Build
}

func (s *CreateModelRequestContentRequestUrlParams) GetSearch() map[string]interface{} {
	return s.Search
}

func (s *CreateModelRequestContentRequestUrlParams) SetBuild(v map[string]interface{}) *CreateModelRequestContentRequestUrlParams {
	s.Build = v
	return s
}

func (s *CreateModelRequestContentRequestUrlParams) SetSearch(v map[string]interface{}) *CreateModelRequestContentRequestUrlParams {
	s.Search = v
	return s
}

func (s *CreateModelRequestContentRequestUrlParams) Validate() error {
	return dara.Validate(s)
}

type CreateModelRequestContentResponse struct {
	// example:
	//
	// $.result.embeddings[*].embedding
	Embeddings *string `json:"embeddings,omitempty" xml:"embeddings,omitempty"`
}

func (s CreateModelRequestContentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequestContentResponse) GoString() string {
	return s.String()
}

func (s *CreateModelRequestContentResponse) GetEmbeddings() *string {
	return s.Embeddings
}

func (s *CreateModelRequestContentResponse) SetEmbeddings(v string) *CreateModelRequestContentResponse {
	s.Embeddings = &v
	return s
}

func (s *CreateModelRequestContentResponse) Validate() error {
	return dara.Validate(s)
}
