// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlInstanceNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSqlInstanceNameResponseBody
	GetRequestId() *string
	SetResult(v *UpdateSqlInstanceNameResponseBodyResult) *UpdateSqlInstanceNameResponseBody
	GetResult() *UpdateSqlInstanceNameResponseBodyResult
}

type UpdateSqlInstanceNameResponseBody struct {
	// id of request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// NodeVO
	Result *UpdateSqlInstanceNameResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateSqlInstanceNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSqlInstanceNameResponseBody) GetResult() *UpdateSqlInstanceNameResponseBodyResult {
	return s.Result
}

func (s *UpdateSqlInstanceNameResponseBody) SetRequestId(v string) *UpdateSqlInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSqlInstanceNameResponseBody) SetResult(v *UpdateSqlInstanceNameResponseBodyResult) *UpdateSqlInstanceNameResponseBody {
	s.Result = v
	return s
}

func (s *UpdateSqlInstanceNameResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateSqlInstanceNameResponseBodyResult struct {
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
	// 22
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
	InstanceId *int64 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// true
	IsDir *int32 `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// -1
	Parent *int64 `json:"parent,omitempty" xml:"parent,omitempty"`
	// example:
	//
	// 1
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// table, instance, template, function
	//
	// example:
	//
	// instance
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateSqlInstanceNameResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlInstanceNameResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceNameResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateSqlInstanceNameResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateSqlInstanceNameResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *UpdateSqlInstanceNameResponseBodyResult) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *UpdateSqlInstanceNameResponseBodyResult) GetIsDir() *int32 {
	return s.IsDir
}

func (s *UpdateSqlInstanceNameResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateSqlInstanceNameResponseBodyResult) GetParent() *int64 {
	return s.Parent
}

func (s *UpdateSqlInstanceNameResponseBodyResult) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *UpdateSqlInstanceNameResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *UpdateSqlInstanceNameResponseBodyResult) SetGmtCreate(v string) *UpdateSqlInstanceNameResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *UpdateSqlInstanceNameResponseBodyResult) SetGmtModified(v string) *UpdateSqlInstanceNameResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *UpdateSqlInstanceNameResponseBodyResult) SetId(v int64) *UpdateSqlInstanceNameResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateSqlInstanceNameResponseBodyResult) SetInstanceId(v int64) *UpdateSqlInstanceNameResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *UpdateSqlInstanceNameResponseBodyResult) SetIsDir(v int32) *UpdateSqlInstanceNameResponseBodyResult {
	s.IsDir = &v
	return s
}

func (s *UpdateSqlInstanceNameResponseBodyResult) SetName(v string) *UpdateSqlInstanceNameResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateSqlInstanceNameResponseBodyResult) SetParent(v int64) *UpdateSqlInstanceNameResponseBodyResult {
	s.Parent = &v
	return s
}

func (s *UpdateSqlInstanceNameResponseBodyResult) SetTemplateId(v int64) *UpdateSqlInstanceNameResponseBodyResult {
	s.TemplateId = &v
	return s
}

func (s *UpdateSqlInstanceNameResponseBodyResult) SetType(v string) *UpdateSqlInstanceNameResponseBodyResult {
	s.Type = &v
	return s
}

func (s *UpdateSqlInstanceNameResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
