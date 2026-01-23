// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetMappingRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAssetMappingRelationsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetAssetMappingRelationsResponseBody
	GetHttpStatusCode() *int32
	SetMappingRelationList(v []*GetAssetMappingRelationsResponseBodyMappingRelationList) *GetAssetMappingRelationsResponseBody
	GetMappingRelationList() []*GetAssetMappingRelationsResponseBodyMappingRelationList
	SetMessage(v string) *GetAssetMappingRelationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAssetMappingRelationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAssetMappingRelationsResponseBody
	GetSuccess() *bool
}

type GetAssetMappingRelationsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode      *int32                                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MappingRelationList []*GetAssetMappingRelationsResponseBodyMappingRelationList `json:"MappingRelationList,omitempty" xml:"MappingRelationList,omitempty" type:"Repeated"`
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

func (s GetAssetMappingRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAssetMappingRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAssetMappingRelationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAssetMappingRelationsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAssetMappingRelationsResponseBody) GetMappingRelationList() []*GetAssetMappingRelationsResponseBodyMappingRelationList {
	return s.MappingRelationList
}

func (s *GetAssetMappingRelationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAssetMappingRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAssetMappingRelationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAssetMappingRelationsResponseBody) SetCode(v string) *GetAssetMappingRelationsResponseBody {
	s.Code = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBody) SetHttpStatusCode(v int32) *GetAssetMappingRelationsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBody) SetMappingRelationList(v []*GetAssetMappingRelationsResponseBodyMappingRelationList) *GetAssetMappingRelationsResponseBody {
	s.MappingRelationList = v
	return s
}

func (s *GetAssetMappingRelationsResponseBody) SetMessage(v string) *GetAssetMappingRelationsResponseBody {
	s.Message = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBody) SetRequestId(v string) *GetAssetMappingRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBody) SetSuccess(v bool) *GetAssetMappingRelationsResponseBody {
	s.Success = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBody) Validate() error {
	if s.MappingRelationList != nil {
		for _, item := range s.MappingRelationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAssetMappingRelationsResponseBodyMappingRelationList struct {
	// example:
	//
	// COLUMN
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// example:
	//
	// 1122
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cc
	StandardCode *string `json:"StandardCode,omitempty" xml:"StandardCode,omitempty"`
	// example:
	//
	// 1
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	// example:
	//
	// test
	StandardName *string `json:"StandardName,omitempty" xml:"StandardName,omitempty"`
	// example:
	//
	// /dir1
	StandardSetDirectory *string `json:"StandardSetDirectory,omitempty" xml:"StandardSetDirectory,omitempty"`
	// example:
	//
	// 2
	StandardSetId *int64 `json:"StandardSetId,omitempty" xml:"StandardSetId,omitempty"`
	// example:
	//
	// test
	StandardSetName *string `json:"StandardSetName,omitempty" xml:"StandardSetName,omitempty"`
	// example:
	//
	// PROD
	StandardStage *string `json:"StandardStage,omitempty" xml:"StandardStage,omitempty"`
}

func (s GetAssetMappingRelationsResponseBodyMappingRelationList) String() string {
	return dara.Prettify(s)
}

func (s GetAssetMappingRelationsResponseBodyMappingRelationList) GoString() string {
	return s.String()
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) GetAssetType() *string {
	return s.AssetType
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) GetGuid() *string {
	return s.Guid
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) GetName() *string {
	return s.Name
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) GetStandardCode() *string {
	return s.StandardCode
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) GetStandardId() *int64 {
	return s.StandardId
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) GetStandardName() *string {
	return s.StandardName
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) GetStandardSetDirectory() *string {
	return s.StandardSetDirectory
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) GetStandardSetId() *int64 {
	return s.StandardSetId
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) GetStandardSetName() *string {
	return s.StandardSetName
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) GetStandardStage() *string {
	return s.StandardStage
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) SetAssetType(v string) *GetAssetMappingRelationsResponseBodyMappingRelationList {
	s.AssetType = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) SetGuid(v string) *GetAssetMappingRelationsResponseBodyMappingRelationList {
	s.Guid = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) SetModifyTime(v string) *GetAssetMappingRelationsResponseBodyMappingRelationList {
	s.ModifyTime = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) SetName(v string) *GetAssetMappingRelationsResponseBodyMappingRelationList {
	s.Name = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) SetStandardCode(v string) *GetAssetMappingRelationsResponseBodyMappingRelationList {
	s.StandardCode = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) SetStandardId(v int64) *GetAssetMappingRelationsResponseBodyMappingRelationList {
	s.StandardId = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) SetStandardName(v string) *GetAssetMappingRelationsResponseBodyMappingRelationList {
	s.StandardName = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) SetStandardSetDirectory(v string) *GetAssetMappingRelationsResponseBodyMappingRelationList {
	s.StandardSetDirectory = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) SetStandardSetId(v int64) *GetAssetMappingRelationsResponseBodyMappingRelationList {
	s.StandardSetId = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) SetStandardSetName(v string) *GetAssetMappingRelationsResponseBodyMappingRelationList {
	s.StandardSetName = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) SetStandardStage(v string) *GetAssetMappingRelationsResponseBodyMappingRelationList {
	s.StandardStage = &v
	return s
}

func (s *GetAssetMappingRelationsResponseBodyMappingRelationList) Validate() error {
	return dara.Validate(s)
}
