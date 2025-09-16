// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateFolderResponseBody
	GetRequestId() *string
	SetResult(v *CreateFolderResponseBodyResult) *CreateFolderResponseBody
	GetResult() *CreateFolderResponseBodyResult
}

type CreateFolderResponseBody struct {
	// id of request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// NodeVO
	Result *CreateFolderResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFolderResponseBody) GetResult() *CreateFolderResponseBodyResult {
	return s.Result
}

func (s *CreateFolderResponseBody) SetRequestId(v string) *CreateFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFolderResponseBody) SetResult(v *CreateFolderResponseBodyResult) *CreateFolderResponseBody {
	s.Result = v
	return s
}

func (s *CreateFolderResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateFolderResponseBodyResult struct {
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
	// 25030
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ha-cn-pl32rf0****
	InstanceId *int64 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// True
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

func (s CreateFolderResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateFolderResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateFolderResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CreateFolderResponseBodyResult) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *CreateFolderResponseBodyResult) GetIsDir() *int32 {
	return s.IsDir
}

func (s *CreateFolderResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateFolderResponseBodyResult) GetParent() *int64 {
	return s.Parent
}

func (s *CreateFolderResponseBodyResult) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CreateFolderResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CreateFolderResponseBodyResult) SetGmtCreate(v string) *CreateFolderResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *CreateFolderResponseBodyResult) SetGmtModified(v string) *CreateFolderResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *CreateFolderResponseBodyResult) SetId(v int64) *CreateFolderResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateFolderResponseBodyResult) SetInstanceId(v int64) *CreateFolderResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *CreateFolderResponseBodyResult) SetIsDir(v int32) *CreateFolderResponseBodyResult {
	s.IsDir = &v
	return s
}

func (s *CreateFolderResponseBodyResult) SetName(v string) *CreateFolderResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateFolderResponseBodyResult) SetParent(v int64) *CreateFolderResponseBodyResult {
	s.Parent = &v
	return s
}

func (s *CreateFolderResponseBodyResult) SetTemplateId(v int64) *CreateFolderResponseBodyResult {
	s.TemplateId = &v
	return s
}

func (s *CreateFolderResponseBodyResult) SetType(v string) *CreateFolderResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateFolderResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
