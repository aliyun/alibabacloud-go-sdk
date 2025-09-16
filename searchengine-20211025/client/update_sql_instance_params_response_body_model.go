// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlInstanceParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSqlInstanceParamsResponseBody
	GetRequestId() *string
	SetResult(v *UpdateSqlInstanceParamsResponseBodyResult) *UpdateSqlInstanceParamsResponseBody
	GetResult() *UpdateSqlInstanceParamsResponseBodyResult
}

type UpdateSqlInstanceParamsResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// InstanceVersionVO
	Result *UpdateSqlInstanceParamsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateSqlInstanceParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlInstanceParamsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSqlInstanceParamsResponseBody) GetResult() *UpdateSqlInstanceParamsResponseBodyResult {
	return s.Result
}

func (s *UpdateSqlInstanceParamsResponseBody) SetRequestId(v string) *UpdateSqlInstanceParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBody) SetResult(v *UpdateSqlInstanceParamsResponseBodyResult) *UpdateSqlInstanceParamsResponseBody {
	s.Result = v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateSqlInstanceParamsResponseBodyResult struct {
	// example:
	//
	// {\\"trace\\":\\"INFO\\",\\"databaseName\\":\\"general\\",\\"iquan.plan.cache.enable\\":true,\\"iquan.plan.prepare.level\\":\\"jni.post.optimize\\",\\"urlencode_data\\":false,\\"formatType\\":\\"string\\",\\"timeout\\":1000}
	CombineParams *string `json:"combineParams,omitempty" xml:"combineParams,omitempty"`
	Comment       *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// select 	- from test
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// {\\"trace\\":\\"INFO\\",\\"databaseName\\":\\"general\\",\\"iquan.plan.cache.enable\\":true,\\"iquan.plan.prepare.level\\":\\"jni.post.optimize\\",\\"urlencode_data\\":false,\\"formatType\\":\\"string\\",\\"timeout\\":1000}
	DynamicParams *string `json:"dynamicParams,omitempty" xml:"dynamicParams,omitempty"`
	// example:
	//
	// 1719221186114
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1719220182844
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	InstanceId *int64 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// {\\"trace\\":\\"INFO\\",\\"databaseName\\":\\"general\\",\\"iquan.plan.cache.enable\\":true,\\"iquan.plan.prepare.level\\":\\"jni.post.optimize\\",\\"urlencode_data\\":false,\\"formatType\\":\\"string\\",\\"timeout\\":1000}
	Kvpairs *string `json:"kvpairs,omitempty" xml:"kvpairs,omitempty"`
	// example:
	//
	// 1
	RelatedTemplateId *int64 `json:"relatedTemplateId,omitempty" xml:"relatedTemplateId,omitempty"`
	// example:
	//
	// {\\"trace\\":\\"INFO\\",\\"databaseName\\":\\"general\\",\\"iquan.plan.cache.enable\\":true,\\"iquan.plan.prepare.level\\":\\"jni.post.optimize\\",\\"urlencode_data\\":false,\\"formatType\\":\\"string\\",\\"timeout\\":1000}
	StaticParams *string `json:"staticParams,omitempty" xml:"staticParams,omitempty"`
	// example:
	//
	// {\\"trace\\":\\"INFO\\",\\"databaseName\\":\\"general\\",\\"iquan.plan.cache.enable\\":true,\\"iquan.plan.prepare.level\\":\\"jni.post.optimize\\",\\"urlencode_data\\":false,\\"formatType\\":\\"string\\",\\"timeout\\":1000}
	TemplateParams *string `json:"templateParams,omitempty" xml:"templateParams,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateSqlInstanceParamsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlInstanceParamsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) GetCombineParams() *string {
	return s.CombineParams
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) GetComment() *string {
	return s.Comment
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) GetDynamicParams() *string {
	return s.DynamicParams
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) GetKvpairs() *string {
	return s.Kvpairs
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) GetRelatedTemplateId() *int64 {
	return s.RelatedTemplateId
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) GetStaticParams() *string {
	return s.StaticParams
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) GetTemplateParams() *string {
	return s.TemplateParams
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) GetVersion() *int64 {
	return s.Version
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) SetCombineParams(v string) *UpdateSqlInstanceParamsResponseBodyResult {
	s.CombineParams = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) SetComment(v string) *UpdateSqlInstanceParamsResponseBodyResult {
	s.Comment = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) SetContent(v string) *UpdateSqlInstanceParamsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) SetDynamicParams(v string) *UpdateSqlInstanceParamsResponseBodyResult {
	s.DynamicParams = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) SetGmtCreate(v string) *UpdateSqlInstanceParamsResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) SetGmtModified(v string) *UpdateSqlInstanceParamsResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) SetInstanceId(v int64) *UpdateSqlInstanceParamsResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) SetKvpairs(v string) *UpdateSqlInstanceParamsResponseBodyResult {
	s.Kvpairs = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) SetRelatedTemplateId(v int64) *UpdateSqlInstanceParamsResponseBodyResult {
	s.RelatedTemplateId = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) SetStaticParams(v string) *UpdateSqlInstanceParamsResponseBodyResult {
	s.StaticParams = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) SetTemplateParams(v string) *UpdateSqlInstanceParamsResponseBodyResult {
	s.TemplateParams = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) SetVersion(v int64) *UpdateSqlInstanceParamsResponseBodyResult {
	s.Version = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
