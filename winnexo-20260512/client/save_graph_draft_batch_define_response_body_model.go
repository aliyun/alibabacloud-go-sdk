// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGraphDraftBatchDefineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveGraphDraftBatchDefineResponseBody
	GetCode() *string
	SetGraphName(v string) *SaveGraphDraftBatchDefineResponseBody
	GetGraphName() *string
	SetItems(v []*SaveGraphDraftBatchDefineResponseBodyItems) *SaveGraphDraftBatchDefineResponseBody
	GetItems() []*SaveGraphDraftBatchDefineResponseBodyItems
	SetMessage(v string) *SaveGraphDraftBatchDefineResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveGraphDraftBatchDefineResponseBody
	GetRequestId() *string
	SetSaveMode(v string) *SaveGraphDraftBatchDefineResponseBody
	GetSaveMode() *string
	SetSavedCount(v int32) *SaveGraphDraftBatchDefineResponseBody
	GetSavedCount() *int32
}

type SaveGraphDraftBatchDefineResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The graph name.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The list of MCP cards.
	Items []*SaveGraphDraftBatchDefineResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The status code description.
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
	// The save mode.
	//
	// example:
	//
	// FULL_YAML
	SaveMode *string `json:"saveMode,omitempty" xml:"saveMode,omitempty"`
	// The number of saved items.
	//
	// example:
	//
	// 2
	SavedCount *int32 `json:"savedCount,omitempty" xml:"savedCount,omitempty"`
}

func (s SaveGraphDraftBatchDefineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveGraphDraftBatchDefineResponseBody) GoString() string {
	return s.String()
}

func (s *SaveGraphDraftBatchDefineResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveGraphDraftBatchDefineResponseBody) GetGraphName() *string {
	return s.GraphName
}

func (s *SaveGraphDraftBatchDefineResponseBody) GetItems() []*SaveGraphDraftBatchDefineResponseBodyItems {
	return s.Items
}

func (s *SaveGraphDraftBatchDefineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveGraphDraftBatchDefineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveGraphDraftBatchDefineResponseBody) GetSaveMode() *string {
	return s.SaveMode
}

func (s *SaveGraphDraftBatchDefineResponseBody) GetSavedCount() *int32 {
	return s.SavedCount
}

func (s *SaveGraphDraftBatchDefineResponseBody) SetCode(v string) *SaveGraphDraftBatchDefineResponseBody {
	s.Code = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBody) SetGraphName(v string) *SaveGraphDraftBatchDefineResponseBody {
	s.GraphName = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBody) SetItems(v []*SaveGraphDraftBatchDefineResponseBodyItems) *SaveGraphDraftBatchDefineResponseBody {
	s.Items = v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBody) SetMessage(v string) *SaveGraphDraftBatchDefineResponseBody {
	s.Message = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBody) SetRequestId(v string) *SaveGraphDraftBatchDefineResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBody) SetSaveMode(v string) *SaveGraphDraftBatchDefineResponseBody {
	s.SaveMode = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBody) SetSavedCount(v int32) *SaveGraphDraftBatchDefineResponseBody {
	s.SavedCount = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBody) Validate() error {
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

type SaveGraphDraftBatchDefineResponseBodyItems struct {
	// The hash of the draft content itself, a 64-character SHA-256 hexadecimal string.
	//
	// example:
	//
	// a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e
	BaseContentHash *string `json:"baseContentHash,omitempty" xml:"baseContentHash,omitempty"`
	// The active schema version number on which the draft is based.
	//
	// example:
	//
	// v1.0.3
	BaseSchemaVersion *string `json:"baseSchemaVersion,omitempty" xml:"baseSchemaVersion,omitempty"`
	// The unique draft change ID, referenced when revoking drafts or publishing.
	//
	// example:
	//
	// 401001
	DraftChangeId *int64 `json:"draftChangeId,omitempty" xml:"draftChangeId,omitempty"`
	// The online content hash on which the draft save is based (draft starting point), a 64-character SHA-256 hexadecimal string.
	//
	// example:
	//
	// e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
	DraftContentHash *string `json:"draftContentHash,omitempty" xml:"draftContentHash,omitempty"`
	// The element type. Currently, only text is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// object_type
	ElementType *string `json:"elementType,omitempty" xml:"elementType,omitempty"`
	// The update time in ISO 8601 format.
	//
	// example:
	//
	// 2026-09-11T10:30:00+00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The operation type.
	//
	// example:
	//
	// UPDATE
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	// The resource name of the agent runtime.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// object
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The source type.
	//
	// example:
	//
	// BATCH_DEFINE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s SaveGraphDraftBatchDefineResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s SaveGraphDraftBatchDefineResponseBodyItems) GoString() string {
	return s.String()
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) GetBaseContentHash() *string {
	return s.BaseContentHash
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) GetBaseSchemaVersion() *string {
	return s.BaseSchemaVersion
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) GetDraftChangeId() *int64 {
	return s.DraftChangeId
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) GetDraftContentHash() *string {
	return s.DraftContentHash
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) GetElementType() *string {
	return s.ElementType
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) GetOperationType() *string {
	return s.OperationType
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) GetResourceName() *string {
	return s.ResourceName
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) SetBaseContentHash(v string) *SaveGraphDraftBatchDefineResponseBodyItems {
	s.BaseContentHash = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) SetBaseSchemaVersion(v string) *SaveGraphDraftBatchDefineResponseBodyItems {
	s.BaseSchemaVersion = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) SetDraftChangeId(v int64) *SaveGraphDraftBatchDefineResponseBodyItems {
	s.DraftChangeId = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) SetDraftContentHash(v string) *SaveGraphDraftBatchDefineResponseBodyItems {
	s.DraftContentHash = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) SetElementType(v string) *SaveGraphDraftBatchDefineResponseBodyItems {
	s.ElementType = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) SetGmtModified(v string) *SaveGraphDraftBatchDefineResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) SetOperationType(v string) *SaveGraphDraftBatchDefineResponseBodyItems {
	s.OperationType = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) SetResourceName(v string) *SaveGraphDraftBatchDefineResponseBodyItems {
	s.ResourceName = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) SetResourceType(v string) *SaveGraphDraftBatchDefineResponseBodyItems {
	s.ResourceType = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) SetSourceType(v string) *SaveGraphDraftBatchDefineResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
