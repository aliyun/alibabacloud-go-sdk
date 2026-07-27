// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKgRelationResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetKgRelationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetKgRelationResponseBody
	GetMessage() *string
	SetRelationInfo(v *GetKgRelationResponseBodyRelationInfo) *GetKgRelationResponseBody
	GetRelationInfo() *GetKgRelationResponseBodyRelationInfo
	SetRequestId(v string) *GetKgRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKgRelationResponseBody
	GetSuccess() *bool
}

type GetKgRelationResponseBody struct {
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
	Message      *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RelationInfo *GetKgRelationResponseBodyRelationInfo `json:"RelationInfo,omitempty" xml:"RelationInfo,omitempty" type:"Struct"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKgRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKgRelationResponseBody) GoString() string {
	return s.String()
}

func (s *GetKgRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKgRelationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetKgRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKgRelationResponseBody) GetRelationInfo() *GetKgRelationResponseBodyRelationInfo {
	return s.RelationInfo
}

func (s *GetKgRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKgRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKgRelationResponseBody) SetCode(v string) *GetKgRelationResponseBody {
	s.Code = &v
	return s
}

func (s *GetKgRelationResponseBody) SetHttpStatusCode(v int32) *GetKgRelationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetKgRelationResponseBody) SetMessage(v string) *GetKgRelationResponseBody {
	s.Message = &v
	return s
}

func (s *GetKgRelationResponseBody) SetRelationInfo(v *GetKgRelationResponseBodyRelationInfo) *GetKgRelationResponseBody {
	s.RelationInfo = v
	return s
}

func (s *GetKgRelationResponseBody) SetRequestId(v string) *GetKgRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKgRelationResponseBody) SetSuccess(v bool) *GetKgRelationResponseBody {
	s.Success = &v
	return s
}

func (s *GetKgRelationResponseBody) Validate() error {
	if s.RelationInfo != nil {
		if err := s.RelationInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKgRelationResponseBodyRelationInfo struct {
	PropertyList []*GetKgRelationResponseBodyRelationInfoPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	// example:
	//
	// abc-xxx
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// example:
	//
	// BELONG_TO
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// example:
	//
	// abc-xxx
	SourceEntityId *string `json:"SourceEntityId,omitempty" xml:"SourceEntityId,omitempty"`
	// example:
	//
	// abd-xxx
	TargetEntityId *string `json:"TargetEntityId,omitempty" xml:"TargetEntityId,omitempty"`
}

func (s GetKgRelationResponseBodyRelationInfo) String() string {
	return dara.Prettify(s)
}

func (s GetKgRelationResponseBodyRelationInfo) GoString() string {
	return s.String()
}

func (s *GetKgRelationResponseBodyRelationInfo) GetPropertyList() []*GetKgRelationResponseBodyRelationInfoPropertyList {
	return s.PropertyList
}

func (s *GetKgRelationResponseBodyRelationInfo) GetRelationId() *string {
	return s.RelationId
}

func (s *GetKgRelationResponseBodyRelationInfo) GetRelationType() *string {
	return s.RelationType
}

func (s *GetKgRelationResponseBodyRelationInfo) GetSourceEntityId() *string {
	return s.SourceEntityId
}

func (s *GetKgRelationResponseBodyRelationInfo) GetTargetEntityId() *string {
	return s.TargetEntityId
}

func (s *GetKgRelationResponseBodyRelationInfo) SetPropertyList(v []*GetKgRelationResponseBodyRelationInfoPropertyList) *GetKgRelationResponseBodyRelationInfo {
	s.PropertyList = v
	return s
}

func (s *GetKgRelationResponseBodyRelationInfo) SetRelationId(v string) *GetKgRelationResponseBodyRelationInfo {
	s.RelationId = &v
	return s
}

func (s *GetKgRelationResponseBodyRelationInfo) SetRelationType(v string) *GetKgRelationResponseBodyRelationInfo {
	s.RelationType = &v
	return s
}

func (s *GetKgRelationResponseBodyRelationInfo) SetSourceEntityId(v string) *GetKgRelationResponseBodyRelationInfo {
	s.SourceEntityId = &v
	return s
}

func (s *GetKgRelationResponseBodyRelationInfo) SetTargetEntityId(v string) *GetKgRelationResponseBodyRelationInfo {
	s.TargetEntityId = &v
	return s
}

func (s *GetKgRelationResponseBodyRelationInfo) Validate() error {
	if s.PropertyList != nil {
		for _, item := range s.PropertyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKgRelationResponseBodyRelationInfoPropertyList struct {
	// example:
	//
	// company_name
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// STRING
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// Alibaba
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetKgRelationResponseBodyRelationInfoPropertyList) String() string {
	return dara.Prettify(s)
}

func (s GetKgRelationResponseBodyRelationInfoPropertyList) GoString() string {
	return s.String()
}

func (s *GetKgRelationResponseBodyRelationInfoPropertyList) GetCode() *string {
	return s.Code
}

func (s *GetKgRelationResponseBodyRelationInfoPropertyList) GetDataType() *string {
	return s.DataType
}

func (s *GetKgRelationResponseBodyRelationInfoPropertyList) GetValue() *string {
	return s.Value
}

func (s *GetKgRelationResponseBodyRelationInfoPropertyList) SetCode(v string) *GetKgRelationResponseBodyRelationInfoPropertyList {
	s.Code = &v
	return s
}

func (s *GetKgRelationResponseBodyRelationInfoPropertyList) SetDataType(v string) *GetKgRelationResponseBodyRelationInfoPropertyList {
	s.DataType = &v
	return s
}

func (s *GetKgRelationResponseBodyRelationInfoPropertyList) SetValue(v string) *GetKgRelationResponseBodyRelationInfoPropertyList {
	s.Value = &v
	return s
}

func (s *GetKgRelationResponseBodyRelationInfoPropertyList) Validate() error {
	return dara.Validate(s)
}
