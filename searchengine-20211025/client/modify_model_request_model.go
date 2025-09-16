// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ModifyModelRequestContent) *ModifyModelRequest
	GetContent() *ModifyModelRequestContent
	SetStatus(v string) *ModifyModelRequest
	GetStatus() *string
	SetDryRun(v string) *ModifyModelRequest
	GetDryRun() *string
}

type ModifyModelRequest struct {
	Content *ModifyModelRequestContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	DryRun *string `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelRequest) GoString() string {
	return s.String()
}

func (s *ModifyModelRequest) GetContent() *ModifyModelRequestContent {
	return s.Content
}

func (s *ModifyModelRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyModelRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *ModifyModelRequest) SetContent(v *ModifyModelRequestContent) *ModifyModelRequest {
	s.Content = v
	return s
}

func (s *ModifyModelRequest) SetStatus(v string) *ModifyModelRequest {
	s.Status = &v
	return s
}

func (s *ModifyModelRequest) SetDryRun(v string) *ModifyModelRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyModelRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyModelRequestContent struct {
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
	Request   *ModifyModelRequestContentRequest  `json:"request,omitempty" xml:"request,omitempty" type:"Struct"`
	Response  *ModifyModelRequestContentResponse `json:"response,omitempty" xml:"response,omitempty" type:"Struct"`
	// example:
	//
	// http://***.platform-cn-shanghai.opensearch.aliyuncs.com/v3/openapi/workspaces/default/text-embedding/ops-text-embedding-001
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ModifyModelRequestContent) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelRequestContent) GoString() string {
	return s.String()
}

func (s *ModifyModelRequestContent) GetDimension() *int32 {
	return s.Dimension
}

func (s *ModifyModelRequestContent) GetMethod() *string {
	return s.Method
}

func (s *ModifyModelRequestContent) GetModelType() *string {
	return s.ModelType
}

func (s *ModifyModelRequestContent) GetRequest() *ModifyModelRequestContentRequest {
	return s.Request
}

func (s *ModifyModelRequestContent) GetResponse() *ModifyModelRequestContentResponse {
	return s.Response
}

func (s *ModifyModelRequestContent) GetUrl() *string {
	return s.Url
}

func (s *ModifyModelRequestContent) SetDimension(v int32) *ModifyModelRequestContent {
	s.Dimension = &v
	return s
}

func (s *ModifyModelRequestContent) SetMethod(v string) *ModifyModelRequestContent {
	s.Method = &v
	return s
}

func (s *ModifyModelRequestContent) SetModelType(v string) *ModifyModelRequestContent {
	s.ModelType = &v
	return s
}

func (s *ModifyModelRequestContent) SetRequest(v *ModifyModelRequestContentRequest) *ModifyModelRequestContent {
	s.Request = v
	return s
}

func (s *ModifyModelRequestContent) SetResponse(v *ModifyModelRequestContentResponse) *ModifyModelRequestContent {
	s.Response = v
	return s
}

func (s *ModifyModelRequestContent) SetUrl(v string) *ModifyModelRequestContent {
	s.Url = &v
	return s
}

func (s *ModifyModelRequestContent) Validate() error {
	return dara.Validate(s)
}

type ModifyModelRequestContentRequest struct {
	Header     *ModifyModelRequestContentRequestHeader     `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Parameters *ModifyModelRequestContentRequestParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	// example:
	//
	// {\\"input\\": [\\"%{input}\\"], \\"input_type\\": \\"%{input_type}\\"}
	RequestBody *string                                    `json:"requestBody,omitempty" xml:"requestBody,omitempty"`
	UrlParams   *ModifyModelRequestContentRequestUrlParams `json:"urlParams,omitempty" xml:"urlParams,omitempty" type:"Struct"`
}

func (s ModifyModelRequestContentRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelRequestContentRequest) GoString() string {
	return s.String()
}

func (s *ModifyModelRequestContentRequest) GetHeader() *ModifyModelRequestContentRequestHeader {
	return s.Header
}

func (s *ModifyModelRequestContentRequest) GetParameters() *ModifyModelRequestContentRequestParameters {
	return s.Parameters
}

func (s *ModifyModelRequestContentRequest) GetRequestBody() *string {
	return s.RequestBody
}

func (s *ModifyModelRequestContentRequest) GetUrlParams() *ModifyModelRequestContentRequestUrlParams {
	return s.UrlParams
}

func (s *ModifyModelRequestContentRequest) SetHeader(v *ModifyModelRequestContentRequestHeader) *ModifyModelRequestContentRequest {
	s.Header = v
	return s
}

func (s *ModifyModelRequestContentRequest) SetParameters(v *ModifyModelRequestContentRequestParameters) *ModifyModelRequestContentRequest {
	s.Parameters = v
	return s
}

func (s *ModifyModelRequestContentRequest) SetRequestBody(v string) *ModifyModelRequestContentRequest {
	s.RequestBody = &v
	return s
}

func (s *ModifyModelRequestContentRequest) SetUrlParams(v *ModifyModelRequestContentRequestUrlParams) *ModifyModelRequestContentRequest {
	s.UrlParams = v
	return s
}

func (s *ModifyModelRequestContentRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyModelRequestContentRequestHeader struct {
	// example:
	//
	// Bearer OS-v0********6vvs
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	// example:
	//
	// application/json
	ContentType *string `json:"Content-Type,omitempty" xml:"Content-Type,omitempty"`
}

func (s ModifyModelRequestContentRequestHeader) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelRequestContentRequestHeader) GoString() string {
	return s.String()
}

func (s *ModifyModelRequestContentRequestHeader) GetAuthorization() *string {
	return s.Authorization
}

func (s *ModifyModelRequestContentRequestHeader) GetContentType() *string {
	return s.ContentType
}

func (s *ModifyModelRequestContentRequestHeader) SetAuthorization(v string) *ModifyModelRequestContentRequestHeader {
	s.Authorization = &v
	return s
}

func (s *ModifyModelRequestContentRequestHeader) SetContentType(v string) *ModifyModelRequestContentRequestHeader {
	s.ContentType = &v
	return s
}

func (s *ModifyModelRequestContentRequestHeader) Validate() error {
	return dara.Validate(s)
}

type ModifyModelRequestContentRequestParameters struct {
	Build  *ModifyModelRequestContentRequestParametersBuild  `json:"build,omitempty" xml:"build,omitempty" type:"Struct"`
	Search *ModifyModelRequestContentRequestParametersSearch `json:"search,omitempty" xml:"search,omitempty" type:"Struct"`
}

func (s ModifyModelRequestContentRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelRequestContentRequestParameters) GoString() string {
	return s.String()
}

func (s *ModifyModelRequestContentRequestParameters) GetBuild() *ModifyModelRequestContentRequestParametersBuild {
	return s.Build
}

func (s *ModifyModelRequestContentRequestParameters) GetSearch() *ModifyModelRequestContentRequestParametersSearch {
	return s.Search
}

func (s *ModifyModelRequestContentRequestParameters) SetBuild(v *ModifyModelRequestContentRequestParametersBuild) *ModifyModelRequestContentRequestParameters {
	s.Build = v
	return s
}

func (s *ModifyModelRequestContentRequestParameters) SetSearch(v *ModifyModelRequestContentRequestParametersSearch) *ModifyModelRequestContentRequestParameters {
	s.Search = v
	return s
}

func (s *ModifyModelRequestContentRequestParameters) Validate() error {
	return dara.Validate(s)
}

type ModifyModelRequestContentRequestParametersBuild struct {
	// example:
	//
	// query
	InputType *string `json:"input_type,omitempty" xml:"input_type,omitempty"`
}

func (s ModifyModelRequestContentRequestParametersBuild) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelRequestContentRequestParametersBuild) GoString() string {
	return s.String()
}

func (s *ModifyModelRequestContentRequestParametersBuild) GetInputType() *string {
	return s.InputType
}

func (s *ModifyModelRequestContentRequestParametersBuild) SetInputType(v string) *ModifyModelRequestContentRequestParametersBuild {
	s.InputType = &v
	return s
}

func (s *ModifyModelRequestContentRequestParametersBuild) Validate() error {
	return dara.Validate(s)
}

type ModifyModelRequestContentRequestParametersSearch struct {
	// example:
	//
	// document
	InputType *string `json:"input_type,omitempty" xml:"input_type,omitempty"`
}

func (s ModifyModelRequestContentRequestParametersSearch) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelRequestContentRequestParametersSearch) GoString() string {
	return s.String()
}

func (s *ModifyModelRequestContentRequestParametersSearch) GetInputType() *string {
	return s.InputType
}

func (s *ModifyModelRequestContentRequestParametersSearch) SetInputType(v string) *ModifyModelRequestContentRequestParametersSearch {
	s.InputType = &v
	return s
}

func (s *ModifyModelRequestContentRequestParametersSearch) Validate() error {
	return dara.Validate(s)
}

type ModifyModelRequestContentRequestUrlParams struct {
	// example:
	//
	// key: value
	Build map[string]interface{} `json:"build,omitempty" xml:"build,omitempty"`
	// example:
	//
	// key: value
	Search map[string]interface{} `json:"search,omitempty" xml:"search,omitempty"`
}

func (s ModifyModelRequestContentRequestUrlParams) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelRequestContentRequestUrlParams) GoString() string {
	return s.String()
}

func (s *ModifyModelRequestContentRequestUrlParams) GetBuild() map[string]interface{} {
	return s.Build
}

func (s *ModifyModelRequestContentRequestUrlParams) GetSearch() map[string]interface{} {
	return s.Search
}

func (s *ModifyModelRequestContentRequestUrlParams) SetBuild(v map[string]interface{}) *ModifyModelRequestContentRequestUrlParams {
	s.Build = v
	return s
}

func (s *ModifyModelRequestContentRequestUrlParams) SetSearch(v map[string]interface{}) *ModifyModelRequestContentRequestUrlParams {
	s.Search = v
	return s
}

func (s *ModifyModelRequestContentRequestUrlParams) Validate() error {
	return dara.Validate(s)
}

type ModifyModelRequestContentResponse struct {
	// example:
	//
	// $.result.embeddings[*].embedding
	Embeddings *string `json:"embeddings,omitempty" xml:"embeddings,omitempty"`
}

func (s ModifyModelRequestContentResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelRequestContentResponse) GoString() string {
	return s.String()
}

func (s *ModifyModelRequestContentResponse) GetEmbeddings() *string {
	return s.Embeddings
}

func (s *ModifyModelRequestContentResponse) SetEmbeddings(v string) *ModifyModelRequestContentResponse {
	s.Embeddings = &v
	return s
}

func (s *ModifyModelRequestContentResponse) Validate() error {
	return dara.Validate(s)
}
