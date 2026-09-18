// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGraphSchemasResponseBody
	GetCode() *string
	SetItems(v []*ListGraphSchemasResponseBodyItems) *ListGraphSchemasResponseBody
	GetItems() []*ListGraphSchemasResponseBodyItems
	SetMessage(v string) *ListGraphSchemasResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGraphSchemasResponseBody
	GetRequestId() *string
}

type ListGraphSchemasResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The location clustering.
	Items []*ListGraphSchemasResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The prompt message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGraphSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGraphSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListGraphSchemasResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGraphSchemasResponseBody) GetItems() []*ListGraphSchemasResponseBodyItems {
	return s.Items
}

func (s *ListGraphSchemasResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGraphSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGraphSchemasResponseBody) SetCode(v string) *ListGraphSchemasResponseBody {
	s.Code = &v
	return s
}

func (s *ListGraphSchemasResponseBody) SetItems(v []*ListGraphSchemasResponseBodyItems) *ListGraphSchemasResponseBody {
	s.Items = v
	return s
}

func (s *ListGraphSchemasResponseBody) SetMessage(v string) *ListGraphSchemasResponseBody {
	s.Message = &v
	return s
}

func (s *ListGraphSchemasResponseBody) SetRequestId(v string) *ListGraphSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGraphSchemasResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGraphSchemasResponseBodyItems struct {
	// The currently active schema version number. The value is 0.0.0 for a quick-created placeholder graph.
	//
	// example:
	//
	// 0.0.0
	ActiveVersion *string `json:"activeVersion,omitempty" xml:"activeVersion,omitempty"`
	// The business description of the graph. An empty string is returned if this parameter is not configured.
	//
	// example:
	//
	// Customer domain semantic graph
	BusinessProfile *string `json:"businessProfile,omitempty" xml:"businessProfile,omitempty"`
	// The display name of the tool.
	//
	// example:
	//
	// CRM Graph
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The graph name.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The status of the semantic graph.
	//
	// This parameter is required.
	//
	// example:
	//
	// PUBLISHED
	GraphStatus *string `json:"graphStatus,omitempty" xml:"graphStatus,omitempty"`
	// Indicates whether the graph contains a draft.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	HasDraft *bool `json:"hasDraft,omitempty" xml:"hasDraft,omitempty"`
	// Indicates whether this is the default group.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// The number of object types. The value falls back to 0 if parsing fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	ObjectTypeCount *int64 `json:"objectTypeCount,omitempty" xml:"objectTypeCount,omitempty"`
	// The number of relations. The value falls back to 0 if parsing fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	RelationCount *int64 `json:"relationCount,omitempty" xml:"relationCount,omitempty"`
	// The list of semantic tags. An empty array [] is returned if this parameter is not configured.
	//
	// This parameter is required.
	SemanticTags []*string `json:"semanticTags,omitempty" xml:"semanticTags,omitempty" type:"Repeated"`
}

func (s ListGraphSchemasResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListGraphSchemasResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListGraphSchemasResponseBodyItems) GetActiveVersion() *string {
	return s.ActiveVersion
}

func (s *ListGraphSchemasResponseBodyItems) GetBusinessProfile() *string {
	return s.BusinessProfile
}

func (s *ListGraphSchemasResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListGraphSchemasResponseBodyItems) GetGraphName() *string {
	return s.GraphName
}

func (s *ListGraphSchemasResponseBodyItems) GetGraphStatus() *string {
	return s.GraphStatus
}

func (s *ListGraphSchemasResponseBodyItems) GetHasDraft() *bool {
	return s.HasDraft
}

func (s *ListGraphSchemasResponseBodyItems) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListGraphSchemasResponseBodyItems) GetObjectTypeCount() *int64 {
	return s.ObjectTypeCount
}

func (s *ListGraphSchemasResponseBodyItems) GetRelationCount() *int64 {
	return s.RelationCount
}

func (s *ListGraphSchemasResponseBodyItems) GetSemanticTags() []*string {
	return s.SemanticTags
}

func (s *ListGraphSchemasResponseBodyItems) SetActiveVersion(v string) *ListGraphSchemasResponseBodyItems {
	s.ActiveVersion = &v
	return s
}

func (s *ListGraphSchemasResponseBodyItems) SetBusinessProfile(v string) *ListGraphSchemasResponseBodyItems {
	s.BusinessProfile = &v
	return s
}

func (s *ListGraphSchemasResponseBodyItems) SetDisplayName(v string) *ListGraphSchemasResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *ListGraphSchemasResponseBodyItems) SetGraphName(v string) *ListGraphSchemasResponseBodyItems {
	s.GraphName = &v
	return s
}

func (s *ListGraphSchemasResponseBodyItems) SetGraphStatus(v string) *ListGraphSchemasResponseBodyItems {
	s.GraphStatus = &v
	return s
}

func (s *ListGraphSchemasResponseBodyItems) SetHasDraft(v bool) *ListGraphSchemasResponseBodyItems {
	s.HasDraft = &v
	return s
}

func (s *ListGraphSchemasResponseBodyItems) SetIsDefault(v bool) *ListGraphSchemasResponseBodyItems {
	s.IsDefault = &v
	return s
}

func (s *ListGraphSchemasResponseBodyItems) SetObjectTypeCount(v int64) *ListGraphSchemasResponseBodyItems {
	s.ObjectTypeCount = &v
	return s
}

func (s *ListGraphSchemasResponseBodyItems) SetRelationCount(v int64) *ListGraphSchemasResponseBodyItems {
	s.RelationCount = &v
	return s
}

func (s *ListGraphSchemasResponseBodyItems) SetSemanticTags(v []*string) *ListGraphSchemasResponseBodyItems {
	s.SemanticTags = v
	return s
}

func (s *ListGraphSchemasResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
