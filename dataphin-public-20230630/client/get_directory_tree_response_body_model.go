// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectoryTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDirectoryTreeResponseBody
	GetCode() *string
	SetData(v []*GetDirectoryTreeResponseBodyData) *GetDirectoryTreeResponseBody
	GetData() []*GetDirectoryTreeResponseBodyData
	SetHttpStatusCode(v int32) *GetDirectoryTreeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDirectoryTreeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDirectoryTreeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDirectoryTreeResponseBody
	GetSuccess() *bool
}

type GetDirectoryTreeResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetDirectoryTreeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDirectoryTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryTreeResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectoryTreeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDirectoryTreeResponseBody) GetData() []*GetDirectoryTreeResponseBodyData {
	return s.Data
}

func (s *GetDirectoryTreeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDirectoryTreeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDirectoryTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDirectoryTreeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDirectoryTreeResponseBody) SetCode(v string) *GetDirectoryTreeResponseBody {
	s.Code = &v
	return s
}

func (s *GetDirectoryTreeResponseBody) SetData(v []*GetDirectoryTreeResponseBodyData) *GetDirectoryTreeResponseBody {
	s.Data = v
	return s
}

func (s *GetDirectoryTreeResponseBody) SetHttpStatusCode(v int32) *GetDirectoryTreeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDirectoryTreeResponseBody) SetMessage(v string) *GetDirectoryTreeResponseBody {
	s.Message = &v
	return s
}

func (s *GetDirectoryTreeResponseBody) SetRequestId(v string) *GetDirectoryTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDirectoryTreeResponseBody) SetSuccess(v bool) *GetDirectoryTreeResponseBody {
	s.Success = &v
	return s
}

func (s *GetDirectoryTreeResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDirectoryTreeResponseBodyData struct {
	Children *GetDirectoryTreeResponseBodyDataChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Struct"`
	Parent   *GetDirectoryTreeResponseBodyDataParent   `json:"Parent,omitempty" xml:"Parent,omitempty" type:"Struct"`
}

func (s GetDirectoryTreeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryTreeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDirectoryTreeResponseBodyData) GetChildren() *GetDirectoryTreeResponseBodyDataChildren {
	return s.Children
}

func (s *GetDirectoryTreeResponseBodyData) GetParent() *GetDirectoryTreeResponseBodyDataParent {
	return s.Parent
}

func (s *GetDirectoryTreeResponseBodyData) SetChildren(v *GetDirectoryTreeResponseBodyDataChildren) *GetDirectoryTreeResponseBodyData {
	s.Children = v
	return s
}

func (s *GetDirectoryTreeResponseBodyData) SetParent(v *GetDirectoryTreeResponseBodyDataParent) *GetDirectoryTreeResponseBodyData {
	s.Parent = v
	return s
}

func (s *GetDirectoryTreeResponseBodyData) Validate() error {
	if s.Children != nil {
		if err := s.Children.Validate(); err != nil {
			return err
		}
	}
	if s.Parent != nil {
		if err := s.Parent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDirectoryTreeResponseBodyDataChildren struct {
	// example:
	//
	// codeManage
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// 300000907
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// John
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// 41342
	DataCellId *int64 `json:"DataCellId,omitempty" xml:"DataCellId,omitempty"`
	// example:
	//
	// /
	DirName *string `json:"DirName,omitempty" xml:"DirName,omitempty"`
	// example:
	//
	// 2024-11-01 18:08:17
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-11-01 18:08:17
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 3
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 300000907
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// John
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// example:
	//
	// 代码管理
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 7081229106458752
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 3
	StringId *string `json:"StringId,omitempty" xml:"StringId,omitempty"`
	// example:
	//
	// codeManage
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDirectoryTreeResponseBodyDataChildren) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryTreeResponseBodyDataChildren) GoString() string {
	return s.String()
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetCategoryType() *string {
	return s.CategoryType
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetCreator() *string {
	return s.Creator
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetDataCellId() *int64 {
	return s.DataCellId
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetDirName() *string {
	return s.DirName
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetId() *int64 {
	return s.Id
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetName() *string {
	return s.Name
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetStringId() *string {
	return s.StringId
}

func (s *GetDirectoryTreeResponseBodyDataChildren) GetType() *string {
	return s.Type
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetCategoryType(v string) *GetDirectoryTreeResponseBodyDataChildren {
	s.CategoryType = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetCreator(v string) *GetDirectoryTreeResponseBodyDataChildren {
	s.Creator = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetCreatorName(v string) *GetDirectoryTreeResponseBodyDataChildren {
	s.CreatorName = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetDataCellId(v int64) *GetDirectoryTreeResponseBodyDataChildren {
	s.DataCellId = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetDirName(v string) *GetDirectoryTreeResponseBodyDataChildren {
	s.DirName = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetGmtCreate(v string) *GetDirectoryTreeResponseBodyDataChildren {
	s.GmtCreate = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetGmtModified(v string) *GetDirectoryTreeResponseBodyDataChildren {
	s.GmtModified = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetId(v int64) *GetDirectoryTreeResponseBodyDataChildren {
	s.Id = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetLastModifier(v string) *GetDirectoryTreeResponseBodyDataChildren {
	s.LastModifier = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetLastModifierName(v string) *GetDirectoryTreeResponseBodyDataChildren {
	s.LastModifierName = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetName(v string) *GetDirectoryTreeResponseBodyDataChildren {
	s.Name = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetProjectId(v int64) *GetDirectoryTreeResponseBodyDataChildren {
	s.ProjectId = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetStringId(v string) *GetDirectoryTreeResponseBodyDataChildren {
	s.StringId = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) SetType(v string) *GetDirectoryTreeResponseBodyDataChildren {
	s.Type = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataChildren) Validate() error {
	return dara.Validate(s)
}

type GetDirectoryTreeResponseBodyDataParent struct {
	// example:
	//
	// codeManage
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// 300000907
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// John
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// 3213
	DataCellId *int64 `json:"DataCellId,omitempty" xml:"DataCellId,omitempty"`
	// example:
	//
	// /
	DirName *string `json:"DirName,omitempty" xml:"DirName,omitempty"`
	// example:
	//
	// 2024-11-01 18:08:17
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-11-01 18:08:17
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 3
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 300000907
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// John
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// example:
	//
	// 代码管理
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 7081229106458752
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 3
	StringId *string `json:"StringId,omitempty" xml:"StringId,omitempty"`
	// example:
	//
	// codeManage
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDirectoryTreeResponseBodyDataParent) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryTreeResponseBodyDataParent) GoString() string {
	return s.String()
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetCategoryType() *string {
	return s.CategoryType
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetCreator() *string {
	return s.Creator
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetDataCellId() *int64 {
	return s.DataCellId
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetDirName() *string {
	return s.DirName
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetId() *int64 {
	return s.Id
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetName() *string {
	return s.Name
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetStringId() *string {
	return s.StringId
}

func (s *GetDirectoryTreeResponseBodyDataParent) GetType() *string {
	return s.Type
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetCategoryType(v string) *GetDirectoryTreeResponseBodyDataParent {
	s.CategoryType = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetCreator(v string) *GetDirectoryTreeResponseBodyDataParent {
	s.Creator = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetCreatorName(v string) *GetDirectoryTreeResponseBodyDataParent {
	s.CreatorName = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetDataCellId(v int64) *GetDirectoryTreeResponseBodyDataParent {
	s.DataCellId = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetDirName(v string) *GetDirectoryTreeResponseBodyDataParent {
	s.DirName = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetGmtCreate(v string) *GetDirectoryTreeResponseBodyDataParent {
	s.GmtCreate = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetGmtModified(v string) *GetDirectoryTreeResponseBodyDataParent {
	s.GmtModified = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetId(v int64) *GetDirectoryTreeResponseBodyDataParent {
	s.Id = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetLastModifier(v string) *GetDirectoryTreeResponseBodyDataParent {
	s.LastModifier = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetLastModifierName(v string) *GetDirectoryTreeResponseBodyDataParent {
	s.LastModifierName = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetName(v string) *GetDirectoryTreeResponseBodyDataParent {
	s.Name = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetProjectId(v int64) *GetDirectoryTreeResponseBodyDataParent {
	s.ProjectId = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetStringId(v string) *GetDirectoryTreeResponseBodyDataParent {
	s.StringId = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) SetType(v string) *GetDirectoryTreeResponseBodyDataParent {
	s.Type = &v
	return s
}

func (s *GetDirectoryTreeResponseBodyDataParent) Validate() error {
	return dara.Validate(s)
}
