// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlInstanceContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSqlInstanceContentResponseBody
	GetRequestId() *string
	SetResult(v *UpdateSqlInstanceContentResponseBodyResult) *UpdateSqlInstanceContentResponseBody
	GetResult() *UpdateSqlInstanceContentResponseBodyResult
}

type UpdateSqlInstanceContentResponseBody struct {
	// id of request
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// InstanceVersionVO
	Result *UpdateSqlInstanceContentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateSqlInstanceContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlInstanceContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSqlInstanceContentResponseBody) GetResult() *UpdateSqlInstanceContentResponseBodyResult {
	return s.Result
}

func (s *UpdateSqlInstanceContentResponseBody) SetRequestId(v string) *UpdateSqlInstanceContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBody) SetResult(v *UpdateSqlInstanceContentResponseBodyResult) *UpdateSqlInstanceContentResponseBody {
	s.Result = v
	return s
}

func (s *UpdateSqlInstanceContentResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateSqlInstanceContentResponseBodyResult struct {
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
	// 1719221186114
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

func (s UpdateSqlInstanceContentResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlInstanceContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceContentResponseBodyResult) GetCombineParams() *string {
	return s.CombineParams
}

func (s *UpdateSqlInstanceContentResponseBodyResult) GetComment() *string {
	return s.Comment
}

func (s *UpdateSqlInstanceContentResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *UpdateSqlInstanceContentResponseBodyResult) GetDynamicParams() *string {
	return s.DynamicParams
}

func (s *UpdateSqlInstanceContentResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateSqlInstanceContentResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateSqlInstanceContentResponseBodyResult) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *UpdateSqlInstanceContentResponseBodyResult) GetKvpairs() *string {
	return s.Kvpairs
}

func (s *UpdateSqlInstanceContentResponseBodyResult) GetRelatedTemplateId() *int64 {
	return s.RelatedTemplateId
}

func (s *UpdateSqlInstanceContentResponseBodyResult) GetStaticParams() *string {
	return s.StaticParams
}

func (s *UpdateSqlInstanceContentResponseBodyResult) GetTemplateParams() *string {
	return s.TemplateParams
}

func (s *UpdateSqlInstanceContentResponseBodyResult) GetVersion() *int64 {
	return s.Version
}

func (s *UpdateSqlInstanceContentResponseBodyResult) SetCombineParams(v string) *UpdateSqlInstanceContentResponseBodyResult {
	s.CombineParams = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBodyResult) SetComment(v string) *UpdateSqlInstanceContentResponseBodyResult {
	s.Comment = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBodyResult) SetContent(v string) *UpdateSqlInstanceContentResponseBodyResult {
	s.Content = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBodyResult) SetDynamicParams(v string) *UpdateSqlInstanceContentResponseBodyResult {
	s.DynamicParams = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBodyResult) SetGmtCreate(v string) *UpdateSqlInstanceContentResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBodyResult) SetGmtModified(v string) *UpdateSqlInstanceContentResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBodyResult) SetInstanceId(v int64) *UpdateSqlInstanceContentResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBodyResult) SetKvpairs(v string) *UpdateSqlInstanceContentResponseBodyResult {
	s.Kvpairs = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBodyResult) SetRelatedTemplateId(v int64) *UpdateSqlInstanceContentResponseBodyResult {
	s.RelatedTemplateId = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBodyResult) SetStaticParams(v string) *UpdateSqlInstanceContentResponseBodyResult {
	s.StaticParams = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBodyResult) SetTemplateParams(v string) *UpdateSqlInstanceContentResponseBodyResult {
	s.TemplateParams = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBodyResult) SetVersion(v int64) *UpdateSqlInstanceContentResponseBodyResult {
	s.Version = &v
	return s
}

func (s *UpdateSqlInstanceContentResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
