// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgNeighborResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKgNeighborResponseBody
	GetCode() *string
	SetData(v *GetKgNeighborResponseBodyData) *GetKgNeighborResponseBody
	GetData() *GetKgNeighborResponseBodyData
	SetHttpStatusCode(v int32) *GetKgNeighborResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetKgNeighborResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKgNeighborResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKgNeighborResponseBody
	GetSuccess() *bool
}

type GetKgNeighborResponseBody struct {
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetKgNeighborResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKgNeighborResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKgNeighborResponseBody) GoString() string {
	return s.String()
}

func (s *GetKgNeighborResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKgNeighborResponseBody) GetData() *GetKgNeighborResponseBodyData {
	return s.Data
}

func (s *GetKgNeighborResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetKgNeighborResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKgNeighborResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKgNeighborResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKgNeighborResponseBody) SetCode(v string) *GetKgNeighborResponseBody {
	s.Code = &v
	return s
}

func (s *GetKgNeighborResponseBody) SetData(v *GetKgNeighborResponseBodyData) *GetKgNeighborResponseBody {
	s.Data = v
	return s
}

func (s *GetKgNeighborResponseBody) SetHttpStatusCode(v int32) *GetKgNeighborResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetKgNeighborResponseBody) SetMessage(v string) *GetKgNeighborResponseBody {
	s.Message = &v
	return s
}

func (s *GetKgNeighborResponseBody) SetRequestId(v string) *GetKgNeighborResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKgNeighborResponseBody) SetSuccess(v bool) *GetKgNeighborResponseBody {
	s.Success = &v
	return s
}

func (s *GetKgNeighborResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKgNeighborResponseBodyData struct {
	EdgeList []*GetKgNeighborResponseBodyDataEdgeList `json:"EdgeList,omitempty" xml:"EdgeList,omitempty" type:"Repeated"`
	NodeList []*GetKgNeighborResponseBodyDataNodeList `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Repeated"`
}

func (s GetKgNeighborResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKgNeighborResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKgNeighborResponseBodyData) GetEdgeList() []*GetKgNeighborResponseBodyDataEdgeList {
	return s.EdgeList
}

func (s *GetKgNeighborResponseBodyData) GetNodeList() []*GetKgNeighborResponseBodyDataNodeList {
	return s.NodeList
}

func (s *GetKgNeighborResponseBodyData) SetEdgeList(v []*GetKgNeighborResponseBodyDataEdgeList) *GetKgNeighborResponseBodyData {
	s.EdgeList = v
	return s
}

func (s *GetKgNeighborResponseBodyData) SetNodeList(v []*GetKgNeighborResponseBodyDataNodeList) *GetKgNeighborResponseBodyData {
	s.NodeList = v
	return s
}

func (s *GetKgNeighborResponseBodyData) Validate() error {
	if s.EdgeList != nil {
		for _, item := range s.EdgeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NodeList != nil {
		for _, item := range s.NodeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKgNeighborResponseBodyDataEdgeList struct {
	PropertyList []*GetKgNeighborResponseBodyDataEdgeListPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
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

func (s GetKgNeighborResponseBodyDataEdgeList) String() string {
	return dara.Prettify(s)
}

func (s GetKgNeighborResponseBodyDataEdgeList) GoString() string {
	return s.String()
}

func (s *GetKgNeighborResponseBodyDataEdgeList) GetPropertyList() []*GetKgNeighborResponseBodyDataEdgeListPropertyList {
	return s.PropertyList
}

func (s *GetKgNeighborResponseBodyDataEdgeList) GetRelationId() *string {
	return s.RelationId
}

func (s *GetKgNeighborResponseBodyDataEdgeList) GetRelationType() *string {
	return s.RelationType
}

func (s *GetKgNeighborResponseBodyDataEdgeList) GetSourceEntityId() *string {
	return s.SourceEntityId
}

func (s *GetKgNeighborResponseBodyDataEdgeList) GetTargetEntityId() *string {
	return s.TargetEntityId
}

func (s *GetKgNeighborResponseBodyDataEdgeList) SetPropertyList(v []*GetKgNeighborResponseBodyDataEdgeListPropertyList) *GetKgNeighborResponseBodyDataEdgeList {
	s.PropertyList = v
	return s
}

func (s *GetKgNeighborResponseBodyDataEdgeList) SetRelationId(v string) *GetKgNeighborResponseBodyDataEdgeList {
	s.RelationId = &v
	return s
}

func (s *GetKgNeighborResponseBodyDataEdgeList) SetRelationType(v string) *GetKgNeighborResponseBodyDataEdgeList {
	s.RelationType = &v
	return s
}

func (s *GetKgNeighborResponseBodyDataEdgeList) SetSourceEntityId(v string) *GetKgNeighborResponseBodyDataEdgeList {
	s.SourceEntityId = &v
	return s
}

func (s *GetKgNeighborResponseBodyDataEdgeList) SetTargetEntityId(v string) *GetKgNeighborResponseBodyDataEdgeList {
	s.TargetEntityId = &v
	return s
}

func (s *GetKgNeighborResponseBodyDataEdgeList) Validate() error {
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

type GetKgNeighborResponseBodyDataEdgeListPropertyList struct {
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

func (s GetKgNeighborResponseBodyDataEdgeListPropertyList) String() string {
	return dara.Prettify(s)
}

func (s GetKgNeighborResponseBodyDataEdgeListPropertyList) GoString() string {
	return s.String()
}

func (s *GetKgNeighborResponseBodyDataEdgeListPropertyList) GetCode() *string {
	return s.Code
}

func (s *GetKgNeighborResponseBodyDataEdgeListPropertyList) GetDataType() *string {
	return s.DataType
}

func (s *GetKgNeighborResponseBodyDataEdgeListPropertyList) GetValue() *string {
	return s.Value
}

func (s *GetKgNeighborResponseBodyDataEdgeListPropertyList) SetCode(v string) *GetKgNeighborResponseBodyDataEdgeListPropertyList {
	s.Code = &v
	return s
}

func (s *GetKgNeighborResponseBodyDataEdgeListPropertyList) SetDataType(v string) *GetKgNeighborResponseBodyDataEdgeListPropertyList {
	s.DataType = &v
	return s
}

func (s *GetKgNeighborResponseBodyDataEdgeListPropertyList) SetValue(v string) *GetKgNeighborResponseBodyDataEdgeListPropertyList {
	s.Value = &v
	return s
}

func (s *GetKgNeighborResponseBodyDataEdgeListPropertyList) Validate() error {
	return dara.Validate(s)
}

type GetKgNeighborResponseBodyDataNodeList struct {
	// example:
	//
	// abc-xxx
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// Company
	EntityType   *string                                              `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	PropertyList []*GetKgNeighborResponseBodyDataNodeListPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
}

func (s GetKgNeighborResponseBodyDataNodeList) String() string {
	return dara.Prettify(s)
}

func (s GetKgNeighborResponseBodyDataNodeList) GoString() string {
	return s.String()
}

func (s *GetKgNeighborResponseBodyDataNodeList) GetEntityId() *string {
	return s.EntityId
}

func (s *GetKgNeighborResponseBodyDataNodeList) GetEntityType() *string {
	return s.EntityType
}

func (s *GetKgNeighborResponseBodyDataNodeList) GetPropertyList() []*GetKgNeighborResponseBodyDataNodeListPropertyList {
	return s.PropertyList
}

func (s *GetKgNeighborResponseBodyDataNodeList) SetEntityId(v string) *GetKgNeighborResponseBodyDataNodeList {
	s.EntityId = &v
	return s
}

func (s *GetKgNeighborResponseBodyDataNodeList) SetEntityType(v string) *GetKgNeighborResponseBodyDataNodeList {
	s.EntityType = &v
	return s
}

func (s *GetKgNeighborResponseBodyDataNodeList) SetPropertyList(v []*GetKgNeighborResponseBodyDataNodeListPropertyList) *GetKgNeighborResponseBodyDataNodeList {
	s.PropertyList = v
	return s
}

func (s *GetKgNeighborResponseBodyDataNodeList) Validate() error {
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

type GetKgNeighborResponseBodyDataNodeListPropertyList struct {
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

func (s GetKgNeighborResponseBodyDataNodeListPropertyList) String() string {
	return dara.Prettify(s)
}

func (s GetKgNeighborResponseBodyDataNodeListPropertyList) GoString() string {
	return s.String()
}

func (s *GetKgNeighborResponseBodyDataNodeListPropertyList) GetCode() *string {
	return s.Code
}

func (s *GetKgNeighborResponseBodyDataNodeListPropertyList) GetDataType() *string {
	return s.DataType
}

func (s *GetKgNeighborResponseBodyDataNodeListPropertyList) GetValue() *string {
	return s.Value
}

func (s *GetKgNeighborResponseBodyDataNodeListPropertyList) SetCode(v string) *GetKgNeighborResponseBodyDataNodeListPropertyList {
	s.Code = &v
	return s
}

func (s *GetKgNeighborResponseBodyDataNodeListPropertyList) SetDataType(v string) *GetKgNeighborResponseBodyDataNodeListPropertyList {
	s.DataType = &v
	return s
}

func (s *GetKgNeighborResponseBodyDataNodeListPropertyList) SetValue(v string) *GetKgNeighborResponseBodyDataNodeListPropertyList {
	s.Value = &v
	return s
}

func (s *GetKgNeighborResponseBodyDataNodeListPropertyList) Validate() error {
	return dara.Validate(s)
}
