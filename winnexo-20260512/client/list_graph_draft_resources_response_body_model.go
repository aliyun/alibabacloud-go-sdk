// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphDraftResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGraphDraftResourcesResponseBody
	GetCode() *string
	SetItems(v []*ListGraphDraftResourcesResponseBodyItems) *ListGraphDraftResourcesResponseBody
	GetItems() []*ListGraphDraftResourcesResponseBodyItems
	SetMessage(v string) *ListGraphDraftResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGraphDraftResourcesResponseBody
	GetRequestId() *string
}

type ListGraphDraftResourcesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The MCP card list.
	Items []*ListGraphDraftResourcesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The prompt message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGraphDraftResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGraphDraftResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGraphDraftResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGraphDraftResourcesResponseBody) GetItems() []*ListGraphDraftResourcesResponseBodyItems {
	return s.Items
}

func (s *ListGraphDraftResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGraphDraftResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGraphDraftResourcesResponseBody) SetCode(v string) *ListGraphDraftResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBody) SetItems(v []*ListGraphDraftResourcesResponseBodyItems) *ListGraphDraftResourcesResponseBody {
	s.Items = v
	return s
}

func (s *ListGraphDraftResourcesResponseBody) SetMessage(v string) *ListGraphDraftResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBody) SetRequestId(v string) *ListGraphDraftResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBody) Validate() error {
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

type ListGraphDraftResourcesResponseBodyItems struct {
	// The hash of the draft content itself. The value is a 64-character SHA-256 hexadecimal string.
	//
	// example:
	//
	// e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
	BaseContentHash *string `json:"baseContentHash,omitempty" xml:"baseContentHash,omitempty"`
	// The active schema version number on which the draft is based.
	//
	// example:
	//
	// 1.2.0
	BaseSchemaVersion *string `json:"baseSchemaVersion,omitempty" xml:"baseSchemaVersion,omitempty"`
	// The unique ID of the draft change. This ID is referenced when you revoke a draft or publish changes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 401001
	DraftChangeId *int64 `json:"draftChangeId,omitempty" xml:"draftChangeId,omitempty"`
	// The hash of the online content on which the draft was based when it was saved (draft starting point). The value is a 64-character SHA-256 hexadecimal string.
	//
	// example:
	//
	// a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e
	DraftContentHash *string `json:"draftContentHash,omitempty" xml:"draftContentHash,omitempty"`
	// The edit mode. In the current implementation, the value is always YAML, which corresponds to sourceType.
	//
	// example:
	//
	// ADVANCED
	EditMode *string `json:"editMode,omitempty" xml:"editMode,omitempty"`
	// The actual publish effect relative to the current online state. After a draft is saved, the online graph may have changed, and the operation intent is adjusted based on the current online state.
	//
	// example:
	//
	// UPDATE
	EffectiveOperation *string `json:"effectiveOperation,omitempty" xml:"effectiveOperation,omitempty"`
	// The element type. Currently, only text is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// indicator
	ElementType *string `json:"elementType,omitempty" xml:"elementType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-09-07T09:00:00+00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time in ISO 8601 format.
	//
	// example:
	//
	// 2026-09-08T10:30:00+00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Indicates whether the draft baseline has expired. The value is true if the hash of the online content at the time the draft was saved is inconsistent with the hash of the current active content. The ONLINE_CHANGED risk is prompted during publishing.
	//
	// This parameter is required.
	//
	// example:
	//
	// true/false
	HasOnlineChanged *bool `json:"hasOnlineChanged,omitempty" xml:"hasOnlineChanged,omitempty"`
	// The operation type.
	//
	// example:
	//
	// UPDATE
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	// The resource name of the agent at runtime.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer_contract_amount
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// element
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The online risk aggregation JSON text (risk_code / risk_message). The value is null if no risk exists.
	//
	// example:
	//
	// {"risk_code": "ONLINE_CHANGED", "risk_message": "The online content has changed."}
	Risk *string `json:"risk,omitempty" xml:"risk,omitempty"`
	// The skill source type.
	//
	// example:
	//
	// YAML
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListGraphDraftResourcesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListGraphDraftResourcesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetBaseContentHash() *string {
	return s.BaseContentHash
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetBaseSchemaVersion() *string {
	return s.BaseSchemaVersion
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetDraftChangeId() *int64 {
	return s.DraftChangeId
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetDraftContentHash() *string {
	return s.DraftContentHash
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetEditMode() *string {
	return s.EditMode
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetEffectiveOperation() *string {
	return s.EffectiveOperation
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetElementType() *string {
	return s.ElementType
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetHasOnlineChanged() *bool {
	return s.HasOnlineChanged
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetOperationType() *string {
	return s.OperationType
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetRisk() *string {
	return s.Risk
}

func (s *ListGraphDraftResourcesResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetBaseContentHash(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.BaseContentHash = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetBaseSchemaVersion(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.BaseSchemaVersion = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetDraftChangeId(v int64) *ListGraphDraftResourcesResponseBodyItems {
	s.DraftChangeId = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetDraftContentHash(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.DraftContentHash = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetEditMode(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.EditMode = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetEffectiveOperation(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.EffectiveOperation = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetElementType(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.ElementType = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetGmtCreate(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetGmtModified(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetHasOnlineChanged(v bool) *ListGraphDraftResourcesResponseBodyItems {
	s.HasOnlineChanged = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetOperationType(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.OperationType = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetResourceName(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.ResourceName = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetResourceType(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.ResourceType = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetRisk(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.Risk = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) SetSourceType(v string) *ListGraphDraftResourcesResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListGraphDraftResourcesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
