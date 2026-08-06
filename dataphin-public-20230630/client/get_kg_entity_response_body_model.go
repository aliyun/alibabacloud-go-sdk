// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKgEntityResponseBody
	GetCode() *string
	SetEntityInfo(v *GetKgEntityResponseBodyEntityInfo) *GetKgEntityResponseBody
	GetEntityInfo() *GetKgEntityResponseBodyEntityInfo
	SetHttpStatusCode(v int32) *GetKgEntityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetKgEntityResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKgEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKgEntityResponseBody
	GetSuccess() *bool
}

type GetKgEntityResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The entity record details.
	EntityInfo *GetKgEntityResponseBodyEntityInfo `json:"EntityInfo,omitempty" xml:"EntityInfo,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
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
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKgEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKgEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetKgEntityResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKgEntityResponseBody) GetEntityInfo() *GetKgEntityResponseBodyEntityInfo {
	return s.EntityInfo
}

func (s *GetKgEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetKgEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKgEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKgEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKgEntityResponseBody) SetCode(v string) *GetKgEntityResponseBody {
	s.Code = &v
	return s
}

func (s *GetKgEntityResponseBody) SetEntityInfo(v *GetKgEntityResponseBodyEntityInfo) *GetKgEntityResponseBody {
	s.EntityInfo = v
	return s
}

func (s *GetKgEntityResponseBody) SetHttpStatusCode(v int32) *GetKgEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetKgEntityResponseBody) SetMessage(v string) *GetKgEntityResponseBody {
	s.Message = &v
	return s
}

func (s *GetKgEntityResponseBody) SetRequestId(v string) *GetKgEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKgEntityResponseBody) SetSuccess(v bool) *GetKgEntityResponseBody {
	s.Success = &v
	return s
}

func (s *GetKgEntityResponseBody) Validate() error {
	if s.EntityInfo != nil {
		if err := s.EntityInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKgEntityResponseBodyEntityInfo struct {
	// The entity record ID.
	//
	// example:
	//
	// abc-xxx
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity type code.
	//
	// example:
	//
	// Company
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The list of entity record properties.
	PropertyList []*GetKgEntityResponseBodyEntityInfoPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
}

func (s GetKgEntityResponseBodyEntityInfo) String() string {
	return dara.Prettify(s)
}

func (s GetKgEntityResponseBodyEntityInfo) GoString() string {
	return s.String()
}

func (s *GetKgEntityResponseBodyEntityInfo) GetEntityId() *string {
	return s.EntityId
}

func (s *GetKgEntityResponseBodyEntityInfo) GetEntityType() *string {
	return s.EntityType
}

func (s *GetKgEntityResponseBodyEntityInfo) GetPropertyList() []*GetKgEntityResponseBodyEntityInfoPropertyList {
	return s.PropertyList
}

func (s *GetKgEntityResponseBodyEntityInfo) SetEntityId(v string) *GetKgEntityResponseBodyEntityInfo {
	s.EntityId = &v
	return s
}

func (s *GetKgEntityResponseBodyEntityInfo) SetEntityType(v string) *GetKgEntityResponseBodyEntityInfo {
	s.EntityType = &v
	return s
}

func (s *GetKgEntityResponseBodyEntityInfo) SetPropertyList(v []*GetKgEntityResponseBodyEntityInfoPropertyList) *GetKgEntityResponseBodyEntityInfo {
	s.PropertyList = v
	return s
}

func (s *GetKgEntityResponseBodyEntityInfo) Validate() error {
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

type GetKgEntityResponseBodyEntityInfoPropertyList struct {
	// The property code.
	//
	// example:
	//
	// company_name
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property data type. Valid values:
	//
	// - STRING: string
	//
	// - INTEGER: integer
	//
	// - FLOAT: floating-point number
	//
	// - BOOLEAN: Boolean
	//
	// - DATE: date
	//
	// - LIST: list
	//
	// example:
	//
	// STRING
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The property value.
	//
	// example:
	//
	// Alibaba
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetKgEntityResponseBodyEntityInfoPropertyList) String() string {
	return dara.Prettify(s)
}

func (s GetKgEntityResponseBodyEntityInfoPropertyList) GoString() string {
	return s.String()
}

func (s *GetKgEntityResponseBodyEntityInfoPropertyList) GetCode() *string {
	return s.Code
}

func (s *GetKgEntityResponseBodyEntityInfoPropertyList) GetDataType() *string {
	return s.DataType
}

func (s *GetKgEntityResponseBodyEntityInfoPropertyList) GetValue() *string {
	return s.Value
}

func (s *GetKgEntityResponseBodyEntityInfoPropertyList) SetCode(v string) *GetKgEntityResponseBodyEntityInfoPropertyList {
	s.Code = &v
	return s
}

func (s *GetKgEntityResponseBodyEntityInfoPropertyList) SetDataType(v string) *GetKgEntityResponseBodyEntityInfoPropertyList {
	s.DataType = &v
	return s
}

func (s *GetKgEntityResponseBodyEntityInfoPropertyList) SetValue(v string) *GetKgEntityResponseBodyEntityInfoPropertyList {
	s.Value = &v
	return s
}

func (s *GetKgEntityResponseBodyEntityInfoPropertyList) Validate() error {
	return dara.Validate(s)
}
