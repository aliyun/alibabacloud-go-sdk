// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSqlInstanceResponseBody
	GetRequestId() *string
	SetResult(v *GetSqlInstanceResponseBodyResult) *GetSqlInstanceResponseBody
	GetResult() *GetSqlInstanceResponseBodyResult
}

type GetSqlInstanceResponseBody struct {
	// id of request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// InstanceVersionVO
	Result *GetSqlInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSqlInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSqlInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSqlInstanceResponseBody) GetResult() *GetSqlInstanceResponseBodyResult {
	return s.Result
}

func (s *GetSqlInstanceResponseBody) SetRequestId(v string) *GetSqlInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlInstanceResponseBody) SetResult(v *GetSqlInstanceResponseBodyResult) *GetSqlInstanceResponseBody {
	s.Result = v
	return s
}

func (s *GetSqlInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSqlInstanceResponseBodyResult struct {
	// example:
	//
	// {\\"trace\\":\\"INFO\\",\\"databaseName\\":\\"general\\",\\"iquan.plan.cache.enable\\":true,\\"iquan.plan.prepare.level\\":\\"jni.post.optimize\\",\\"urlencode_data\\":false,\\"formatType\\":\\"string\\",\\"timeout\\":1000}
	CombineParams *string `json:"combineParams,omitempty" xml:"combineParams,omitempty"`
	// example:
	//
	// init version
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
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
	// 1719220182844
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1719221186114
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// ha-cn-pl32rf0****
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

func (s GetSqlInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetSqlInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSqlInstanceResponseBodyResult) GetCombineParams() *string {
	return s.CombineParams
}

func (s *GetSqlInstanceResponseBodyResult) GetComment() *string {
	return s.Comment
}

func (s *GetSqlInstanceResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *GetSqlInstanceResponseBodyResult) GetDynamicParams() *string {
	return s.DynamicParams
}

func (s *GetSqlInstanceResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetSqlInstanceResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetSqlInstanceResponseBodyResult) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetSqlInstanceResponseBodyResult) GetKvpairs() *string {
	return s.Kvpairs
}

func (s *GetSqlInstanceResponseBodyResult) GetRelatedTemplateId() *int64 {
	return s.RelatedTemplateId
}

func (s *GetSqlInstanceResponseBodyResult) GetStaticParams() *string {
	return s.StaticParams
}

func (s *GetSqlInstanceResponseBodyResult) GetTemplateParams() *string {
	return s.TemplateParams
}

func (s *GetSqlInstanceResponseBodyResult) GetVersion() *int64 {
	return s.Version
}

func (s *GetSqlInstanceResponseBodyResult) SetCombineParams(v string) *GetSqlInstanceResponseBodyResult {
	s.CombineParams = &v
	return s
}

func (s *GetSqlInstanceResponseBodyResult) SetComment(v string) *GetSqlInstanceResponseBodyResult {
	s.Comment = &v
	return s
}

func (s *GetSqlInstanceResponseBodyResult) SetContent(v string) *GetSqlInstanceResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetSqlInstanceResponseBodyResult) SetDynamicParams(v string) *GetSqlInstanceResponseBodyResult {
	s.DynamicParams = &v
	return s
}

func (s *GetSqlInstanceResponseBodyResult) SetGmtCreate(v string) *GetSqlInstanceResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *GetSqlInstanceResponseBodyResult) SetGmtModified(v string) *GetSqlInstanceResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *GetSqlInstanceResponseBodyResult) SetInstanceId(v int64) *GetSqlInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *GetSqlInstanceResponseBodyResult) SetKvpairs(v string) *GetSqlInstanceResponseBodyResult {
	s.Kvpairs = &v
	return s
}

func (s *GetSqlInstanceResponseBodyResult) SetRelatedTemplateId(v int64) *GetSqlInstanceResponseBodyResult {
	s.RelatedTemplateId = &v
	return s
}

func (s *GetSqlInstanceResponseBodyResult) SetStaticParams(v string) *GetSqlInstanceResponseBodyResult {
	s.StaticParams = &v
	return s
}

func (s *GetSqlInstanceResponseBodyResult) SetTemplateParams(v string) *GetSqlInstanceResponseBodyResult {
	s.TemplateParams = &v
	return s
}

func (s *GetSqlInstanceResponseBodyResult) SetVersion(v int64) *GetSqlInstanceResponseBodyResult {
	s.Version = &v
	return s
}

func (s *GetSqlInstanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
