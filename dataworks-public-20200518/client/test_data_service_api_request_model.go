// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestDataServiceApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *TestDataServiceApiRequest
	GetApiId() *int64
	SetBodyContent(v string) *TestDataServiceApiRequest
	GetBodyContent() *string
	SetBodyParams(v []*TestDataServiceApiRequestBodyParams) *TestDataServiceApiRequest
	GetBodyParams() []*TestDataServiceApiRequestBodyParams
	SetHeadParams(v []*TestDataServiceApiRequestHeadParams) *TestDataServiceApiRequest
	GetHeadParams() []*TestDataServiceApiRequestHeadParams
	SetPathParams(v []*TestDataServiceApiRequestPathParams) *TestDataServiceApiRequest
	GetPathParams() []*TestDataServiceApiRequestPathParams
	SetQueryParam(v []*TestDataServiceApiRequestQueryParam) *TestDataServiceApiRequest
	GetQueryParam() []*TestDataServiceApiRequestQueryParam
}

type TestDataServiceApiRequest struct {
	// The ID of the DataService Studio API on which the test is performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12343
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The data of the request body.
	//
	// example:
	//
	// {"name":"test"}
	BodyContent *string `json:"BodyContent,omitempty" xml:"BodyContent,omitempty"`
	// The request parameters that are contained in the request body.
	BodyParams []*TestDataServiceApiRequestBodyParams `json:"BodyParams,omitempty" xml:"BodyParams,omitempty" type:"Repeated"`
	// The request parameters that are contained in the request header.
	HeadParams []*TestDataServiceApiRequestHeadParams `json:"HeadParams,omitempty" xml:"HeadParams,omitempty" type:"Repeated"`
	// The request parameters that are contained in the request path.
	PathParams []*TestDataServiceApiRequestPathParams `json:"PathParams,omitempty" xml:"PathParams,omitempty" type:"Repeated"`
	// The request parameters that are contained in the query.
	QueryParam []*TestDataServiceApiRequestQueryParam `json:"QueryParam,omitempty" xml:"QueryParam,omitempty" type:"Repeated"`
}

func (s TestDataServiceApiRequest) String() string {
	return dara.Prettify(s)
}

func (s TestDataServiceApiRequest) GoString() string {
	return s.String()
}

func (s *TestDataServiceApiRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *TestDataServiceApiRequest) GetBodyContent() *string {
	return s.BodyContent
}

func (s *TestDataServiceApiRequest) GetBodyParams() []*TestDataServiceApiRequestBodyParams {
	return s.BodyParams
}

func (s *TestDataServiceApiRequest) GetHeadParams() []*TestDataServiceApiRequestHeadParams {
	return s.HeadParams
}

func (s *TestDataServiceApiRequest) GetPathParams() []*TestDataServiceApiRequestPathParams {
	return s.PathParams
}

func (s *TestDataServiceApiRequest) GetQueryParam() []*TestDataServiceApiRequestQueryParam {
	return s.QueryParam
}

func (s *TestDataServiceApiRequest) SetApiId(v int64) *TestDataServiceApiRequest {
	s.ApiId = &v
	return s
}

func (s *TestDataServiceApiRequest) SetBodyContent(v string) *TestDataServiceApiRequest {
	s.BodyContent = &v
	return s
}

func (s *TestDataServiceApiRequest) SetBodyParams(v []*TestDataServiceApiRequestBodyParams) *TestDataServiceApiRequest {
	s.BodyParams = v
	return s
}

func (s *TestDataServiceApiRequest) SetHeadParams(v []*TestDataServiceApiRequestHeadParams) *TestDataServiceApiRequest {
	s.HeadParams = v
	return s
}

func (s *TestDataServiceApiRequest) SetPathParams(v []*TestDataServiceApiRequestPathParams) *TestDataServiceApiRequest {
	s.PathParams = v
	return s
}

func (s *TestDataServiceApiRequest) SetQueryParam(v []*TestDataServiceApiRequestQueryParam) *TestDataServiceApiRequest {
	s.QueryParam = v
	return s
}

func (s *TestDataServiceApiRequest) Validate() error {
	return dara.Validate(s)
}

type TestDataServiceApiRequestBodyParams struct {
	// The name of the parameter.
	//
	// example:
	//
	// name
	ParamKey *string `json:"ParamKey,omitempty" xml:"ParamKey,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// test
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s TestDataServiceApiRequestBodyParams) String() string {
	return dara.Prettify(s)
}

func (s TestDataServiceApiRequestBodyParams) GoString() string {
	return s.String()
}

func (s *TestDataServiceApiRequestBodyParams) GetParamKey() *string {
	return s.ParamKey
}

func (s *TestDataServiceApiRequestBodyParams) GetParamValue() *string {
	return s.ParamValue
}

func (s *TestDataServiceApiRequestBodyParams) SetParamKey(v string) *TestDataServiceApiRequestBodyParams {
	s.ParamKey = &v
	return s
}

func (s *TestDataServiceApiRequestBodyParams) SetParamValue(v string) *TestDataServiceApiRequestBodyParams {
	s.ParamValue = &v
	return s
}

func (s *TestDataServiceApiRequestBodyParams) Validate() error {
	return dara.Validate(s)
}

type TestDataServiceApiRequestHeadParams struct {
	// The name of the parameter.
	//
	// example:
	//
	// requestId
	ParamKey *string `json:"ParamKey,omitempty" xml:"ParamKey,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// abcd
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s TestDataServiceApiRequestHeadParams) String() string {
	return dara.Prettify(s)
}

func (s TestDataServiceApiRequestHeadParams) GoString() string {
	return s.String()
}

func (s *TestDataServiceApiRequestHeadParams) GetParamKey() *string {
	return s.ParamKey
}

func (s *TestDataServiceApiRequestHeadParams) GetParamValue() *string {
	return s.ParamValue
}

func (s *TestDataServiceApiRequestHeadParams) SetParamKey(v string) *TestDataServiceApiRequestHeadParams {
	s.ParamKey = &v
	return s
}

func (s *TestDataServiceApiRequestHeadParams) SetParamValue(v string) *TestDataServiceApiRequestHeadParams {
	s.ParamValue = &v
	return s
}

func (s *TestDataServiceApiRequestHeadParams) Validate() error {
	return dara.Validate(s)
}

type TestDataServiceApiRequestPathParams struct {
	// The name of the parameter.
	//
	// example:
	//
	// path1
	ParamKey *string `json:"ParamKey,omitempty" xml:"ParamKey,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// api
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s TestDataServiceApiRequestPathParams) String() string {
	return dara.Prettify(s)
}

func (s TestDataServiceApiRequestPathParams) GoString() string {
	return s.String()
}

func (s *TestDataServiceApiRequestPathParams) GetParamKey() *string {
	return s.ParamKey
}

func (s *TestDataServiceApiRequestPathParams) GetParamValue() *string {
	return s.ParamValue
}

func (s *TestDataServiceApiRequestPathParams) SetParamKey(v string) *TestDataServiceApiRequestPathParams {
	s.ParamKey = &v
	return s
}

func (s *TestDataServiceApiRequestPathParams) SetParamValue(v string) *TestDataServiceApiRequestPathParams {
	s.ParamValue = &v
	return s
}

func (s *TestDataServiceApiRequestPathParams) Validate() error {
	return dara.Validate(s)
}

type TestDataServiceApiRequestQueryParam struct {
	// The name of the parameter.
	//
	// example:
	//
	// name
	ParamKey *string `json:"ParamKey,omitempty" xml:"ParamKey,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// test
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s TestDataServiceApiRequestQueryParam) String() string {
	return dara.Prettify(s)
}

func (s TestDataServiceApiRequestQueryParam) GoString() string {
	return s.String()
}

func (s *TestDataServiceApiRequestQueryParam) GetParamKey() *string {
	return s.ParamKey
}

func (s *TestDataServiceApiRequestQueryParam) GetParamValue() *string {
	return s.ParamValue
}

func (s *TestDataServiceApiRequestQueryParam) SetParamKey(v string) *TestDataServiceApiRequestQueryParam {
	s.ParamKey = &v
	return s
}

func (s *TestDataServiceApiRequestQueryParam) SetParamValue(v string) *TestDataServiceApiRequestQueryParam {
	s.ParamValue = &v
	return s
}

func (s *TestDataServiceApiRequestQueryParam) Validate() error {
	return dara.Validate(s)
}
