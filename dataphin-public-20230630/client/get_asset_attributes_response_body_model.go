// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAssetAttributesResponseBody
	GetCode() *string
	SetData(v *GetAssetAttributesResponseBodyData) *GetAssetAttributesResponseBody
	GetData() *GetAssetAttributesResponseBodyData
	SetHttpStatusCode(v int32) *GetAssetAttributesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAssetAttributesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAssetAttributesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAssetAttributesResponseBody
	GetSuccess() *bool
}

type GetAssetAttributesResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *GetAssetAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
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

func (s GetAssetAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAssetAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAssetAttributesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAssetAttributesResponseBody) GetData() *GetAssetAttributesResponseBodyData {
	return s.Data
}

func (s *GetAssetAttributesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAssetAttributesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAssetAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAssetAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAssetAttributesResponseBody) SetCode(v string) *GetAssetAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *GetAssetAttributesResponseBody) SetData(v *GetAssetAttributesResponseBodyData) *GetAssetAttributesResponseBody {
	s.Data = v
	return s
}

func (s *GetAssetAttributesResponseBody) SetHttpStatusCode(v int32) *GetAssetAttributesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAssetAttributesResponseBody) SetMessage(v string) *GetAssetAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *GetAssetAttributesResponseBody) SetRequestId(v string) *GetAssetAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAssetAttributesResponseBody) SetSuccess(v bool) *GetAssetAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *GetAssetAttributesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAssetAttributesResponseBodyData struct {
	// The list of asset properties.
	AssetAttributeList []*GetAssetAttributesResponseBodyDataAssetAttributeList `json:"AssetAttributeList,omitempty" xml:"AssetAttributeList,omitempty" type:"Repeated"`
}

func (s GetAssetAttributesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAssetAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAssetAttributesResponseBodyData) GetAssetAttributeList() []*GetAssetAttributesResponseBodyDataAssetAttributeList {
	return s.AssetAttributeList
}

func (s *GetAssetAttributesResponseBodyData) SetAssetAttributeList(v []*GetAssetAttributesResponseBodyDataAssetAttributeList) *GetAssetAttributesResponseBodyData {
	s.AssetAttributeList = v
	return s
}

func (s *GetAssetAttributesResponseBodyData) Validate() error {
	if s.AssetAttributeList != nil {
		for _, item := range s.AssetAttributeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAssetAttributesResponseBodyDataAssetAttributeList struct {
	// The asset name.
	//
	// example:
	//
	// table_orders
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The asset type.
	//
	// example:
	//
	// TABLE
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The list of property values.
	AttributeList []*GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList `json:"AttributeList,omitempty" xml:"AttributeList,omitempty" type:"Repeated"`
	// The unique identifier of the asset.
	//
	// example:
	//
	// odps.project_a.table_orders
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// The last modified time, in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2026-06-01 10:30:00
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
}

func (s GetAssetAttributesResponseBodyDataAssetAttributeList) String() string {
	return dara.Prettify(s)
}

func (s GetAssetAttributesResponseBodyDataAssetAttributeList) GoString() string {
	return s.String()
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeList) GetAssetName() *string {
	return s.AssetName
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeList) GetAssetType() *string {
	return s.AssetType
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeList) GetAttributeList() []*GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList {
	return s.AttributeList
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeList) GetGuid() *string {
	return s.Guid
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeList) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeList) SetAssetName(v string) *GetAssetAttributesResponseBodyDataAssetAttributeList {
	s.AssetName = &v
	return s
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeList) SetAssetType(v string) *GetAssetAttributesResponseBodyDataAssetAttributeList {
	s.AssetType = &v
	return s
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeList) SetAttributeList(v []*GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList) *GetAssetAttributesResponseBodyDataAssetAttributeList {
	s.AttributeList = v
	return s
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeList) SetGuid(v string) *GetAssetAttributesResponseBodyDataAssetAttributeList {
	s.Guid = &v
	return s
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeList) SetLastModifiedTime(v string) *GetAssetAttributesResponseBodyDataAssetAttributeList {
	s.LastModifiedTime = &v
	return s
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeList) Validate() error {
	if s.AttributeList != nil {
		for _, item := range s.AttributeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList struct {
	// The property code.
	//
	// example:
	//
	// data_level
	AttributeCode *string `json:"AttributeCode,omitempty" xml:"AttributeCode,omitempty"`
	// The display name of the property.
	//
	// example:
	//
	// Data Level
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// Indicates whether the property is required.
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// The list of property values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList) String() string {
	return dara.Prettify(s)
}

func (s GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList) GoString() string {
	return s.String()
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList) GetAttributeCode() *string {
	return s.AttributeCode
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList) GetAttributeName() *string {
	return s.AttributeName
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList) GetRequired() *bool {
	return s.Required
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList) GetValues() []*string {
	return s.Values
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList) SetAttributeCode(v string) *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList {
	s.AttributeCode = &v
	return s
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList) SetAttributeName(v string) *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList {
	s.AttributeName = &v
	return s
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList) SetRequired(v bool) *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList {
	s.Required = &v
	return s
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList) SetValues(v []*string) *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList {
	s.Values = v
	return s
}

func (s *GetAssetAttributesResponseBodyDataAssetAttributeListAttributeList) Validate() error {
	return dara.Validate(s)
}
