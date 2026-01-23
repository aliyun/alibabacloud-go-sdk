// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStandardSetResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetStandardSetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetStandardSetResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStandardSetResponseBody
	GetRequestId() *string
	SetStandardSetInfo(v *GetStandardSetResponseBodyStandardSetInfo) *GetStandardSetResponseBody
	GetStandardSetInfo() *GetStandardSetResponseBodyStandardSetInfo
	SetSuccess(v bool) *GetStandardSetResponseBody
	GetSuccess() *bool
}

type GetStandardSetResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId       *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StandardSetInfo *GetStandardSetResponseBodyStandardSetInfo `json:"StandardSetInfo,omitempty" xml:"StandardSetInfo,omitempty" type:"Struct"`
	Success         *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStandardSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponseBody) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStandardSetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetStandardSetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStandardSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStandardSetResponseBody) GetStandardSetInfo() *GetStandardSetResponseBodyStandardSetInfo {
	return s.StandardSetInfo
}

func (s *GetStandardSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStandardSetResponseBody) SetCode(v string) *GetStandardSetResponseBody {
	s.Code = &v
	return s
}

func (s *GetStandardSetResponseBody) SetHttpStatusCode(v int32) *GetStandardSetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetStandardSetResponseBody) SetMessage(v string) *GetStandardSetResponseBody {
	s.Message = &v
	return s
}

func (s *GetStandardSetResponseBody) SetRequestId(v string) *GetStandardSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStandardSetResponseBody) SetStandardSetInfo(v *GetStandardSetResponseBodyStandardSetInfo) *GetStandardSetResponseBody {
	s.StandardSetInfo = v
	return s
}

func (s *GetStandardSetResponseBody) SetSuccess(v bool) *GetStandardSetResponseBody {
	s.Success = &v
	return s
}

func (s *GetStandardSetResponseBody) Validate() error {
	if s.StandardSetInfo != nil {
		if err := s.StandardSetInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardSetResponseBodyStandardSetInfo struct {
	ApprovalConfig *GetStandardSetResponseBodyStandardSetInfoApprovalConfig `json:"ApprovalConfig,omitempty" xml:"ApprovalConfig,omitempty" type:"Struct"`
	// example:
	//
	// CITY
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator    *GetStandardSetResponseBodyStandardSetInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// 1001
	DefaultStandardTemplateId *int64 `json:"DefaultStandardTemplateId,omitempty" xml:"DefaultStandardTemplateId,omitempty"`
	// example:
	//
	// test
	Description        *string                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryReference *GetStandardSetResponseBodyStandardSetInfoDirectoryReference `json:"DirectoryReference,omitempty" xml:"DirectoryReference,omitempty" type:"Struct"`
	// example:
	//
	// 1234
	Id              *int64                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	LastModifier    *GetStandardSetResponseBodyStandardSetInfoLastModifier      `json:"LastModifier,omitempty" xml:"LastModifier,omitempty" type:"Struct"`
	MaintainerList  []*GetStandardSetResponseBodyStandardSetInfoMaintainerList  `json:"MaintainerList,omitempty" xml:"MaintainerList,omitempty" type:"Repeated"`
	MemberGroupList []*GetStandardSetResponseBodyStandardSetInfoMemberGroupList `json:"MemberGroupList,omitempty" xml:"MemberGroupList,omitempty" type:"Repeated"`
	MemberList      []*GetStandardSetResponseBodyStandardSetInfoMemberList      `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test
	Name                  *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	OfflineApprovalConfig *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig `json:"OfflineApprovalConfig,omitempty" xml:"OfflineApprovalConfig,omitempty" type:"Struct"`
	VisibilityConfig      *GetStandardSetResponseBodyStandardSetInfoVisibilityConfig      `json:"VisibilityConfig,omitempty" xml:"VisibilityConfig,omitempty" type:"Struct"`
}

func (s GetStandardSetResponseBodyStandardSetInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponseBodyStandardSetInfo) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetApprovalConfig() *GetStandardSetResponseBodyStandardSetInfoApprovalConfig {
	return s.ApprovalConfig
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetCode() *string {
	return s.Code
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetCreator() *GetStandardSetResponseBodyStandardSetInfoCreator {
	return s.Creator
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetDefaultStandardTemplateId() *int64 {
	return s.DefaultStandardTemplateId
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetDescription() *string {
	return s.Description
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetDirectoryReference() *GetStandardSetResponseBodyStandardSetInfoDirectoryReference {
	return s.DirectoryReference
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetId() *int64 {
	return s.Id
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetLastModifier() *GetStandardSetResponseBodyStandardSetInfoLastModifier {
	return s.LastModifier
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetMaintainerList() []*GetStandardSetResponseBodyStandardSetInfoMaintainerList {
	return s.MaintainerList
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetMemberGroupList() []*GetStandardSetResponseBodyStandardSetInfoMemberGroupList {
	return s.MemberGroupList
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetMemberList() []*GetStandardSetResponseBodyStandardSetInfoMemberList {
	return s.MemberList
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetName() *string {
	return s.Name
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetOfflineApprovalConfig() *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig {
	return s.OfflineApprovalConfig
}

func (s *GetStandardSetResponseBodyStandardSetInfo) GetVisibilityConfig() *GetStandardSetResponseBodyStandardSetInfoVisibilityConfig {
	return s.VisibilityConfig
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetApprovalConfig(v *GetStandardSetResponseBodyStandardSetInfoApprovalConfig) *GetStandardSetResponseBodyStandardSetInfo {
	s.ApprovalConfig = v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetCode(v string) *GetStandardSetResponseBodyStandardSetInfo {
	s.Code = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetCreateTime(v string) *GetStandardSetResponseBodyStandardSetInfo {
	s.CreateTime = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetCreator(v *GetStandardSetResponseBodyStandardSetInfoCreator) *GetStandardSetResponseBodyStandardSetInfo {
	s.Creator = v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetDefaultStandardTemplateId(v int64) *GetStandardSetResponseBodyStandardSetInfo {
	s.DefaultStandardTemplateId = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetDescription(v string) *GetStandardSetResponseBodyStandardSetInfo {
	s.Description = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetDirectoryReference(v *GetStandardSetResponseBodyStandardSetInfoDirectoryReference) *GetStandardSetResponseBodyStandardSetInfo {
	s.DirectoryReference = v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetId(v int64) *GetStandardSetResponseBodyStandardSetInfo {
	s.Id = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetLastModifier(v *GetStandardSetResponseBodyStandardSetInfoLastModifier) *GetStandardSetResponseBodyStandardSetInfo {
	s.LastModifier = v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetMaintainerList(v []*GetStandardSetResponseBodyStandardSetInfoMaintainerList) *GetStandardSetResponseBodyStandardSetInfo {
	s.MaintainerList = v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetMemberGroupList(v []*GetStandardSetResponseBodyStandardSetInfoMemberGroupList) *GetStandardSetResponseBodyStandardSetInfo {
	s.MemberGroupList = v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetMemberList(v []*GetStandardSetResponseBodyStandardSetInfoMemberList) *GetStandardSetResponseBodyStandardSetInfo {
	s.MemberList = v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetModifyTime(v string) *GetStandardSetResponseBodyStandardSetInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetName(v string) *GetStandardSetResponseBodyStandardSetInfo {
	s.Name = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetOfflineApprovalConfig(v *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig) *GetStandardSetResponseBodyStandardSetInfo {
	s.OfflineApprovalConfig = v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) SetVisibilityConfig(v *GetStandardSetResponseBodyStandardSetInfoVisibilityConfig) *GetStandardSetResponseBodyStandardSetInfo {
	s.VisibilityConfig = v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfo) Validate() error {
	if s.ApprovalConfig != nil {
		if err := s.ApprovalConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.DirectoryReference != nil {
		if err := s.DirectoryReference.Validate(); err != nil {
			return err
		}
	}
	if s.LastModifier != nil {
		if err := s.LastModifier.Validate(); err != nil {
			return err
		}
	}
	if s.MaintainerList != nil {
		for _, item := range s.MaintainerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MemberGroupList != nil {
		for _, item := range s.MemberGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MemberList != nil {
		for _, item := range s.MemberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OfflineApprovalConfig != nil {
		if err := s.OfflineApprovalConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VisibilityConfig != nil {
		if err := s.VisibilityConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardSetResponseBodyStandardSetInfoApprovalConfig struct {
	// example:
	//
	// BY_DEFAULT
	ApprovalType    *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	EnableApproval  *bool   `json:"EnableApproval,omitempty" xml:"EnableApproval,omitempty"`
	IsSubmitInBatch *bool   `json:"IsSubmitInBatch,omitempty" xml:"IsSubmitInBatch,omitempty"`
	// example:
	//
	// 1121
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetStandardSetResponseBodyStandardSetInfoApprovalConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponseBodyStandardSetInfoApprovalConfig) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponseBodyStandardSetInfoApprovalConfig) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *GetStandardSetResponseBodyStandardSetInfoApprovalConfig) GetEnableApproval() *bool {
	return s.EnableApproval
}

func (s *GetStandardSetResponseBodyStandardSetInfoApprovalConfig) GetIsSubmitInBatch() *bool {
	return s.IsSubmitInBatch
}

func (s *GetStandardSetResponseBodyStandardSetInfoApprovalConfig) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetStandardSetResponseBodyStandardSetInfoApprovalConfig) SetApprovalType(v string) *GetStandardSetResponseBodyStandardSetInfoApprovalConfig {
	s.ApprovalType = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoApprovalConfig) SetEnableApproval(v bool) *GetStandardSetResponseBodyStandardSetInfoApprovalConfig {
	s.EnableApproval = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoApprovalConfig) SetIsSubmitInBatch(v bool) *GetStandardSetResponseBodyStandardSetInfoApprovalConfig {
	s.IsSubmitInBatch = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoApprovalConfig) SetTemplateId(v int64) *GetStandardSetResponseBodyStandardSetInfoApprovalConfig {
	s.TemplateId = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoApprovalConfig) Validate() error {
	return dara.Validate(s)
}

type GetStandardSetResponseBodyStandardSetInfoCreator struct {
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardSetResponseBodyStandardSetInfoCreator) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponseBodyStandardSetInfoCreator) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponseBodyStandardSetInfoCreator) GetId() *string {
	return s.Id
}

func (s *GetStandardSetResponseBodyStandardSetInfoCreator) GetName() *string {
	return s.Name
}

func (s *GetStandardSetResponseBodyStandardSetInfoCreator) SetId(v string) *GetStandardSetResponseBodyStandardSetInfoCreator {
	s.Id = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoCreator) SetName(v string) *GetStandardSetResponseBodyStandardSetInfoCreator {
	s.Name = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoCreator) Validate() error {
	return dara.Validate(s)
}

type GetStandardSetResponseBodyStandardSetInfoDirectoryReference struct {
	// example:
	//
	// /dir1
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
}

func (s GetStandardSetResponseBodyStandardSetInfoDirectoryReference) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponseBodyStandardSetInfoDirectoryReference) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponseBodyStandardSetInfoDirectoryReference) GetDirectory() *string {
	return s.Directory
}

func (s *GetStandardSetResponseBodyStandardSetInfoDirectoryReference) SetDirectory(v string) *GetStandardSetResponseBodyStandardSetInfoDirectoryReference {
	s.Directory = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoDirectoryReference) Validate() error {
	return dara.Validate(s)
}

type GetStandardSetResponseBodyStandardSetInfoLastModifier struct {
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardSetResponseBodyStandardSetInfoLastModifier) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponseBodyStandardSetInfoLastModifier) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponseBodyStandardSetInfoLastModifier) GetId() *string {
	return s.Id
}

func (s *GetStandardSetResponseBodyStandardSetInfoLastModifier) GetName() *string {
	return s.Name
}

func (s *GetStandardSetResponseBodyStandardSetInfoLastModifier) SetId(v string) *GetStandardSetResponseBodyStandardSetInfoLastModifier {
	s.Id = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoLastModifier) SetName(v string) *GetStandardSetResponseBodyStandardSetInfoLastModifier {
	s.Name = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoLastModifier) Validate() error {
	return dara.Validate(s)
}

type GetStandardSetResponseBodyStandardSetInfoMaintainerList struct {
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardSetResponseBodyStandardSetInfoMaintainerList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponseBodyStandardSetInfoMaintainerList) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponseBodyStandardSetInfoMaintainerList) GetId() *string {
	return s.Id
}

func (s *GetStandardSetResponseBodyStandardSetInfoMaintainerList) GetName() *string {
	return s.Name
}

func (s *GetStandardSetResponseBodyStandardSetInfoMaintainerList) SetId(v string) *GetStandardSetResponseBodyStandardSetInfoMaintainerList {
	s.Id = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoMaintainerList) SetName(v string) *GetStandardSetResponseBodyStandardSetInfoMaintainerList {
	s.Name = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoMaintainerList) Validate() error {
	return dara.Validate(s)
}

type GetStandardSetResponseBodyStandardSetInfoMemberGroupList struct {
	// example:
	//
	// 1121
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// testGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardSetResponseBodyStandardSetInfoMemberGroupList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponseBodyStandardSetInfoMemberGroupList) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponseBodyStandardSetInfoMemberGroupList) GetId() *string {
	return s.Id
}

func (s *GetStandardSetResponseBodyStandardSetInfoMemberGroupList) GetName() *string {
	return s.Name
}

func (s *GetStandardSetResponseBodyStandardSetInfoMemberGroupList) SetId(v string) *GetStandardSetResponseBodyStandardSetInfoMemberGroupList {
	s.Id = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoMemberGroupList) SetName(v string) *GetStandardSetResponseBodyStandardSetInfoMemberGroupList {
	s.Name = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoMemberGroupList) Validate() error {
	return dara.Validate(s)
}

type GetStandardSetResponseBodyStandardSetInfoMemberList struct {
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardSetResponseBodyStandardSetInfoMemberList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponseBodyStandardSetInfoMemberList) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponseBodyStandardSetInfoMemberList) GetId() *string {
	return s.Id
}

func (s *GetStandardSetResponseBodyStandardSetInfoMemberList) GetName() *string {
	return s.Name
}

func (s *GetStandardSetResponseBodyStandardSetInfoMemberList) SetId(v string) *GetStandardSetResponseBodyStandardSetInfoMemberList {
	s.Id = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoMemberList) SetName(v string) *GetStandardSetResponseBodyStandardSetInfoMemberList {
	s.Name = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoMemberList) Validate() error {
	return dara.Validate(s)
}

type GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig struct {
	// example:
	//
	// BY_DEFAULT
	ApprovalType    *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	EnableApproval  *bool   `json:"EnableApproval,omitempty" xml:"EnableApproval,omitempty"`
	IsSubmitInBatch *bool   `json:"IsSubmitInBatch,omitempty" xml:"IsSubmitInBatch,omitempty"`
	// example:
	//
	// 1121
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig) GetEnableApproval() *bool {
	return s.EnableApproval
}

func (s *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig) GetIsSubmitInBatch() *bool {
	return s.IsSubmitInBatch
}

func (s *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig) SetApprovalType(v string) *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig {
	s.ApprovalType = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig) SetEnableApproval(v bool) *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig {
	s.EnableApproval = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig) SetIsSubmitInBatch(v bool) *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig {
	s.IsSubmitInBatch = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig) SetTemplateId(v int64) *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig {
	s.TemplateId = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoOfflineApprovalConfig) Validate() error {
	return dara.Validate(s)
}

type GetStandardSetResponseBodyStandardSetInfoVisibilityConfig struct {
	SpecifiedUserList []*GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList `json:"SpecifiedUserList,omitempty" xml:"SpecifiedUserList,omitempty" type:"Repeated"`
	// example:
	//
	// PUBLIC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStandardSetResponseBodyStandardSetInfoVisibilityConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponseBodyStandardSetInfoVisibilityConfig) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponseBodyStandardSetInfoVisibilityConfig) GetSpecifiedUserList() []*GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList {
	return s.SpecifiedUserList
}

func (s *GetStandardSetResponseBodyStandardSetInfoVisibilityConfig) GetType() *string {
	return s.Type
}

func (s *GetStandardSetResponseBodyStandardSetInfoVisibilityConfig) SetSpecifiedUserList(v []*GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList) *GetStandardSetResponseBodyStandardSetInfoVisibilityConfig {
	s.SpecifiedUserList = v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoVisibilityConfig) SetType(v string) *GetStandardSetResponseBodyStandardSetInfoVisibilityConfig {
	s.Type = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoVisibilityConfig) Validate() error {
	if s.SpecifiedUserList != nil {
		for _, item := range s.SpecifiedUserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList struct {
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList) GetId() *string {
	return s.Id
}

func (s *GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList) GetName() *string {
	return s.Name
}

func (s *GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList) SetId(v string) *GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList {
	s.Id = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList) SetName(v string) *GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList {
	s.Name = &v
	return s
}

func (s *GetStandardSetResponseBodyStandardSetInfoVisibilityConfigSpecifiedUserList) Validate() error {
	return dara.Validate(s)
}
