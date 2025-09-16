// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneSqlInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloneSqlInstanceResponseBody
	GetRequestId() *string
	SetResult(v *CloneSqlInstanceResponseBodyResult) *CloneSqlInstanceResponseBody
	GetResult() *CloneSqlInstanceResponseBodyResult
}

type CloneSqlInstanceResponseBody struct {
	// id of request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// NodeVO
	Result *CloneSqlInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CloneSqlInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneSqlInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CloneSqlInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneSqlInstanceResponseBody) GetResult() *CloneSqlInstanceResponseBodyResult {
	return s.Result
}

func (s *CloneSqlInstanceResponseBody) SetRequestId(v string) *CloneSqlInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneSqlInstanceResponseBody) SetResult(v *CloneSqlInstanceResponseBodyResult) *CloneSqlInstanceResponseBody {
	s.Result = v
	return s
}

func (s *CloneSqlInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CloneSqlInstanceResponseBodyResult struct {
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
	// -cn-pl32rf0****
	InstanceId *int64 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// True
	IsDir *int32 `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
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

func (s CloneSqlInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CloneSqlInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CloneSqlInstanceResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CloneSqlInstanceResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CloneSqlInstanceResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CloneSqlInstanceResponseBodyResult) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *CloneSqlInstanceResponseBodyResult) GetIsDir() *int32 {
	return s.IsDir
}

func (s *CloneSqlInstanceResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CloneSqlInstanceResponseBodyResult) GetParent() *int64 {
	return s.Parent
}

func (s *CloneSqlInstanceResponseBodyResult) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CloneSqlInstanceResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CloneSqlInstanceResponseBodyResult) SetGmtCreate(v string) *CloneSqlInstanceResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *CloneSqlInstanceResponseBodyResult) SetGmtModified(v string) *CloneSqlInstanceResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *CloneSqlInstanceResponseBodyResult) SetId(v int64) *CloneSqlInstanceResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CloneSqlInstanceResponseBodyResult) SetInstanceId(v int64) *CloneSqlInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *CloneSqlInstanceResponseBodyResult) SetIsDir(v int32) *CloneSqlInstanceResponseBodyResult {
	s.IsDir = &v
	return s
}

func (s *CloneSqlInstanceResponseBodyResult) SetName(v string) *CloneSqlInstanceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CloneSqlInstanceResponseBodyResult) SetParent(v int64) *CloneSqlInstanceResponseBodyResult {
	s.Parent = &v
	return s
}

func (s *CloneSqlInstanceResponseBodyResult) SetTemplateId(v int64) *CloneSqlInstanceResponseBodyResult {
	s.TemplateId = &v
	return s
}

func (s *CloneSqlInstanceResponseBodyResult) SetType(v string) *CloneSqlInstanceResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CloneSqlInstanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
