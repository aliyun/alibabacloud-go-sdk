// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardLookupTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStandardLookupTableResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetStandardLookupTableResponseBody
	GetHttpStatusCode() *int32
	SetLookupTableInfo(v *GetStandardLookupTableResponseBodyLookupTableInfo) *GetStandardLookupTableResponseBody
	GetLookupTableInfo() *GetStandardLookupTableResponseBodyLookupTableInfo
	SetMessage(v string) *GetStandardLookupTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStandardLookupTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStandardLookupTableResponseBody
	GetSuccess() *bool
}

type GetStandardLookupTableResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode  *int32                                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	LookupTableInfo *GetStandardLookupTableResponseBodyLookupTableInfo `json:"LookupTableInfo,omitempty" xml:"LookupTableInfo,omitempty" type:"Struct"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStandardLookupTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStandardLookupTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetStandardLookupTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStandardLookupTableResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetStandardLookupTableResponseBody) GetLookupTableInfo() *GetStandardLookupTableResponseBodyLookupTableInfo {
	return s.LookupTableInfo
}

func (s *GetStandardLookupTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStandardLookupTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStandardLookupTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStandardLookupTableResponseBody) SetCode(v string) *GetStandardLookupTableResponseBody {
	s.Code = &v
	return s
}

func (s *GetStandardLookupTableResponseBody) SetHttpStatusCode(v int32) *GetStandardLookupTableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetStandardLookupTableResponseBody) SetLookupTableInfo(v *GetStandardLookupTableResponseBodyLookupTableInfo) *GetStandardLookupTableResponseBody {
	s.LookupTableInfo = v
	return s
}

func (s *GetStandardLookupTableResponseBody) SetMessage(v string) *GetStandardLookupTableResponseBody {
	s.Message = &v
	return s
}

func (s *GetStandardLookupTableResponseBody) SetRequestId(v string) *GetStandardLookupTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStandardLookupTableResponseBody) SetSuccess(v bool) *GetStandardLookupTableResponseBody {
	s.Success = &v
	return s
}

func (s *GetStandardLookupTableResponseBody) Validate() error {
	if s.LookupTableInfo != nil {
		if err := s.LookupTableInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardLookupTableResponseBodyLookupTableInfo struct {
	// example:
	//
	// CITY
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 30012021
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// test
	Description        *string                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectoryReference *GetStandardLookupTableResponseBodyLookupTableInfoDirectoryReference `json:"DirectoryReference,omitempty" xml:"DirectoryReference,omitempty" type:"Struct"`
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 30012021
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// test
	LastModifierName     *string                                                                  `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	LookupTableValueList []*GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList `json:"LookupTableValueList,omitempty" xml:"LookupTableValueList,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 城市码表
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30012021
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
}

func (s GetStandardLookupTableResponseBodyLookupTableInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStandardLookupTableResponseBodyLookupTableInfo) GoString() string {
	return s.String()
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetCode() *string {
	return s.Code
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetDescription() *string {
	return s.Description
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetDirectoryReference() *GetStandardLookupTableResponseBodyLookupTableInfoDirectoryReference {
	return s.DirectoryReference
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetId() *int64 {
	return s.Id
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetLookupTableValueList() []*GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList {
	return s.LookupTableValueList
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetName() *string {
	return s.Name
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetOwner() *string {
	return s.Owner
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetCode(v string) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.Code = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetCreateTime(v string) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.CreateTime = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetCreator(v string) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.Creator = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetCreatorName(v string) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.CreatorName = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetDescription(v string) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.Description = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetDirectoryReference(v *GetStandardLookupTableResponseBodyLookupTableInfoDirectoryReference) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.DirectoryReference = v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetId(v int64) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.Id = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetLastModifier(v string) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.LastModifier = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetLastModifierName(v string) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.LastModifierName = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetLookupTableValueList(v []*GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.LookupTableValueList = v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetModifyTime(v string) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetName(v string) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.Name = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetOwner(v string) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.Owner = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) SetOwnerName(v string) *GetStandardLookupTableResponseBodyLookupTableInfo {
	s.OwnerName = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfo) Validate() error {
	if s.DirectoryReference != nil {
		if err := s.DirectoryReference.Validate(); err != nil {
			return err
		}
	}
	if s.LookupTableValueList != nil {
		for _, item := range s.LookupTableValueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStandardLookupTableResponseBodyLookupTableInfoDirectoryReference struct {
	// example:
	//
	// /dir1/dir2
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
}

func (s GetStandardLookupTableResponseBodyLookupTableInfoDirectoryReference) String() string {
	return dara.Prettify(s)
}

func (s GetStandardLookupTableResponseBodyLookupTableInfoDirectoryReference) GoString() string {
	return s.String()
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfoDirectoryReference) GetDirectory() *string {
	return s.Directory
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfoDirectoryReference) SetDirectory(v string) *GetStandardLookupTableResponseBodyLookupTableInfoDirectoryReference {
	s.Directory = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfoDirectoryReference) Validate() error {
	return dara.Validate(s)
}

type GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList struct {
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// HZ
	EnglishName *string `json:"EnglishName,omitempty" xml:"EnglishName,omitempty"`
	// example:
	//
	// Hangzhou
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 杭州
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList) GoString() string {
	return s.String()
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList) GetDescription() *string {
	return s.Description
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList) GetEnglishName() *string {
	return s.EnglishName
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList) GetName() *string {
	return s.Name
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList) GetValue() *string {
	return s.Value
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList) SetDescription(v string) *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList {
	s.Description = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList) SetEnglishName(v string) *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList {
	s.EnglishName = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList) SetName(v string) *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList {
	s.Name = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList) SetValue(v string) *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList {
	s.Value = &v
	return s
}

func (s *GetStandardLookupTableResponseBodyLookupTableInfoLookupTableValueList) Validate() error {
	return dara.Validate(s)
}
