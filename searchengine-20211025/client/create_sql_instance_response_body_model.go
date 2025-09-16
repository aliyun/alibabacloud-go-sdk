// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSqlInstanceResponseBody
	GetRequestId() *string
	SetResult(v *CreateSqlInstanceResponseBodyResult) *CreateSqlInstanceResponseBody
	GetResult() *CreateSqlInstanceResponseBodyResult
}

type CreateSqlInstanceResponseBody struct {
	// id of request
	//
	// example:
	//
	// B43CD1BB-ABD7-59C5-B89A-6E5F6FE60A84
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// NodeVO
	Result *CreateSqlInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateSqlInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSqlInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSqlInstanceResponseBody) GetResult() *CreateSqlInstanceResponseBodyResult {
	return s.Result
}

func (s *CreateSqlInstanceResponseBody) SetRequestId(v string) *CreateSqlInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSqlInstanceResponseBody) SetResult(v *CreateSqlInstanceResponseBodyResult) *CreateSqlInstanceResponseBody {
	s.Result = v
	return s
}

func (s *CreateSqlInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateSqlInstanceResponseBodyResult struct {
	// example:
	//
	// 1719220182844
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1719220182844
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 22
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ha-cn-pl32rf0****
	InstanceId *int64 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// true
	IsDir *int32 `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// example:
	//
	// test
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

func (s CreateSqlInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateSqlInstanceResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateSqlInstanceResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateSqlInstanceResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CreateSqlInstanceResponseBodyResult) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *CreateSqlInstanceResponseBodyResult) GetIsDir() *int32 {
	return s.IsDir
}

func (s *CreateSqlInstanceResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateSqlInstanceResponseBodyResult) GetParent() *int64 {
	return s.Parent
}

func (s *CreateSqlInstanceResponseBodyResult) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CreateSqlInstanceResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CreateSqlInstanceResponseBodyResult) SetGmtCreate(v string) *CreateSqlInstanceResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *CreateSqlInstanceResponseBodyResult) SetGmtModified(v string) *CreateSqlInstanceResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *CreateSqlInstanceResponseBodyResult) SetId(v int64) *CreateSqlInstanceResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateSqlInstanceResponseBodyResult) SetInstanceId(v int64) *CreateSqlInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *CreateSqlInstanceResponseBodyResult) SetIsDir(v int32) *CreateSqlInstanceResponseBodyResult {
	s.IsDir = &v
	return s
}

func (s *CreateSqlInstanceResponseBodyResult) SetName(v string) *CreateSqlInstanceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateSqlInstanceResponseBodyResult) SetParent(v int64) *CreateSqlInstanceResponseBodyResult {
	s.Parent = &v
	return s
}

func (s *CreateSqlInstanceResponseBodyResult) SetTemplateId(v int64) *CreateSqlInstanceResponseBodyResult {
	s.TemplateId = &v
	return s
}

func (s *CreateSqlInstanceResponseBodyResult) SetType(v string) *CreateSqlInstanceResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateSqlInstanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
