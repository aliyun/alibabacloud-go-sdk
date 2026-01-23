// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBelongAssetMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBelongAssetMappingResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetBelongAssetMappingResponseBody
	GetHttpStatusCode() *int32
	SetMappingRelationList(v []*GetBelongAssetMappingResponseBodyMappingRelationList) *GetBelongAssetMappingResponseBody
	GetMappingRelationList() []*GetBelongAssetMappingResponseBodyMappingRelationList
	SetMessage(v string) *GetBelongAssetMappingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBelongAssetMappingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBelongAssetMappingResponseBody
	GetSuccess() *bool
}

type GetBelongAssetMappingResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode      *int32                                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MappingRelationList []*GetBelongAssetMappingResponseBodyMappingRelationList `json:"MappingRelationList,omitempty" xml:"MappingRelationList,omitempty" type:"Repeated"`
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

func (s GetBelongAssetMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBelongAssetMappingResponseBody) GoString() string {
	return s.String()
}

func (s *GetBelongAssetMappingResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBelongAssetMappingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBelongAssetMappingResponseBody) GetMappingRelationList() []*GetBelongAssetMappingResponseBodyMappingRelationList {
	return s.MappingRelationList
}

func (s *GetBelongAssetMappingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBelongAssetMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBelongAssetMappingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBelongAssetMappingResponseBody) SetCode(v string) *GetBelongAssetMappingResponseBody {
	s.Code = &v
	return s
}

func (s *GetBelongAssetMappingResponseBody) SetHttpStatusCode(v int32) *GetBelongAssetMappingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBelongAssetMappingResponseBody) SetMappingRelationList(v []*GetBelongAssetMappingResponseBodyMappingRelationList) *GetBelongAssetMappingResponseBody {
	s.MappingRelationList = v
	return s
}

func (s *GetBelongAssetMappingResponseBody) SetMessage(v string) *GetBelongAssetMappingResponseBody {
	s.Message = &v
	return s
}

func (s *GetBelongAssetMappingResponseBody) SetRequestId(v string) *GetBelongAssetMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBelongAssetMappingResponseBody) SetSuccess(v bool) *GetBelongAssetMappingResponseBody {
	s.Success = &v
	return s
}

func (s *GetBelongAssetMappingResponseBody) Validate() error {
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

type GetBelongAssetMappingResponseBodyMappingRelationList struct {
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

func (s GetBelongAssetMappingResponseBodyMappingRelationList) String() string {
	return dara.Prettify(s)
}

func (s GetBelongAssetMappingResponseBodyMappingRelationList) GoString() string {
	return s.String()
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) GetAssetType() *string {
	return s.AssetType
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) GetGuid() *string {
	return s.Guid
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) GetName() *string {
	return s.Name
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) GetStandardCode() *string {
	return s.StandardCode
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) GetStandardId() *int64 {
	return s.StandardId
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) GetStandardName() *string {
	return s.StandardName
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) GetStandardSetDirectory() *string {
	return s.StandardSetDirectory
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) GetStandardSetId() *int64 {
	return s.StandardSetId
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) GetStandardSetName() *string {
	return s.StandardSetName
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) GetStandardStage() *string {
	return s.StandardStage
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) SetAssetType(v string) *GetBelongAssetMappingResponseBodyMappingRelationList {
	s.AssetType = &v
	return s
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) SetGuid(v string) *GetBelongAssetMappingResponseBodyMappingRelationList {
	s.Guid = &v
	return s
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) SetModifyTime(v string) *GetBelongAssetMappingResponseBodyMappingRelationList {
	s.ModifyTime = &v
	return s
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) SetName(v string) *GetBelongAssetMappingResponseBodyMappingRelationList {
	s.Name = &v
	return s
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) SetStandardCode(v string) *GetBelongAssetMappingResponseBodyMappingRelationList {
	s.StandardCode = &v
	return s
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) SetStandardId(v int64) *GetBelongAssetMappingResponseBodyMappingRelationList {
	s.StandardId = &v
	return s
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) SetStandardName(v string) *GetBelongAssetMappingResponseBodyMappingRelationList {
	s.StandardName = &v
	return s
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) SetStandardSetDirectory(v string) *GetBelongAssetMappingResponseBodyMappingRelationList {
	s.StandardSetDirectory = &v
	return s
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) SetStandardSetId(v int64) *GetBelongAssetMappingResponseBodyMappingRelationList {
	s.StandardSetId = &v
	return s
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) SetStandardSetName(v string) *GetBelongAssetMappingResponseBodyMappingRelationList {
	s.StandardSetName = &v
	return s
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) SetStandardStage(v string) *GetBelongAssetMappingResponseBodyMappingRelationList {
	s.StandardStage = &v
	return s
}

func (s *GetBelongAssetMappingResponseBodyMappingRelationList) Validate() error {
	return dara.Validate(s)
}
