// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenameFolderResponseBody
	GetRequestId() *string
	SetResult(v *RenameFolderResponseBodyResult) *RenameFolderResponseBody
	GetResult() *RenameFolderResponseBodyResult
}

type RenameFolderResponseBody struct {
	// id of request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// NodeVO
	Result *RenameFolderResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s RenameFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenameFolderResponseBody) GoString() string {
	return s.String()
}

func (s *RenameFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenameFolderResponseBody) GetResult() *RenameFolderResponseBodyResult {
	return s.Result
}

func (s *RenameFolderResponseBody) SetRequestId(v string) *RenameFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameFolderResponseBody) SetResult(v *RenameFolderResponseBodyResult) *RenameFolderResponseBody {
	s.Result = v
	return s
}

func (s *RenameFolderResponseBody) Validate() error {
	return dara.Validate(s)
}

type RenameFolderResponseBodyResult struct {
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
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
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
	// template
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RenameFolderResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s RenameFolderResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RenameFolderResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *RenameFolderResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *RenameFolderResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *RenameFolderResponseBodyResult) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *RenameFolderResponseBodyResult) GetIsDir() *int32 {
	return s.IsDir
}

func (s *RenameFolderResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *RenameFolderResponseBodyResult) GetParent() *int64 {
	return s.Parent
}

func (s *RenameFolderResponseBodyResult) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *RenameFolderResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *RenameFolderResponseBodyResult) SetGmtCreate(v string) *RenameFolderResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *RenameFolderResponseBodyResult) SetGmtModified(v string) *RenameFolderResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *RenameFolderResponseBodyResult) SetId(v int64) *RenameFolderResponseBodyResult {
	s.Id = &v
	return s
}

func (s *RenameFolderResponseBodyResult) SetInstanceId(v int64) *RenameFolderResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *RenameFolderResponseBodyResult) SetIsDir(v int32) *RenameFolderResponseBodyResult {
	s.IsDir = &v
	return s
}

func (s *RenameFolderResponseBodyResult) SetName(v string) *RenameFolderResponseBodyResult {
	s.Name = &v
	return s
}

func (s *RenameFolderResponseBodyResult) SetParent(v int64) *RenameFolderResponseBodyResult {
	s.Parent = &v
	return s
}

func (s *RenameFolderResponseBodyResult) SetTemplateId(v int64) *RenameFolderResponseBodyResult {
	s.TemplateId = &v
	return s
}

func (s *RenameFolderResponseBodyResult) SetType(v string) *RenameFolderResponseBodyResult {
	s.Type = &v
	return s
}

func (s *RenameFolderResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
