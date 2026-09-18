// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGraphDraftResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveGraphDraftResourceResponseBody
	GetCode() *string
	SetDraftChangeId(v int64) *SaveGraphDraftResourceResponseBody
	GetDraftChangeId() *int64
	SetDraftContentHash(v string) *SaveGraphDraftResourceResponseBody
	GetDraftContentHash() *string
	SetElementType(v string) *SaveGraphDraftResourceResponseBody
	GetElementType() *string
	SetGmtModified(v string) *SaveGraphDraftResourceResponseBody
	GetGmtModified() *string
	SetGraphName(v string) *SaveGraphDraftResourceResponseBody
	GetGraphName() *string
	SetMessage(v string) *SaveGraphDraftResourceResponseBody
	GetMessage() *string
	SetOperationType(v string) *SaveGraphDraftResourceResponseBody
	GetOperationType() *string
	SetRequestId(v string) *SaveGraphDraftResourceResponseBody
	GetRequestId() *string
	SetResourceName(v string) *SaveGraphDraftResourceResponseBody
	GetResourceName() *string
	SetResourceType(v string) *SaveGraphDraftResourceResponseBody
	GetResourceType() *string
}

type SaveGraphDraftResourceResponseBody struct {
	// The business status code. A value of 200 indicates success. A failure returns a backend error code (ERR.	- / InvalidParameter.*).
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The draft change ID (the draftChangeId returned by listGraphDraftResources).
	//
	// example:
	//
	// 401001
	DraftChangeId *int64 `json:"draftChangeId,omitempty" xml:"draftChangeId,omitempty"`
	// The online content hash on which the draft save is based (draft starting point), a 64-character SHA-256 hexadecimal string.
	//
	// example:
	//
	// sha256:abc123
	DraftContentHash *string `json:"draftContentHash,omitempty" xml:"draftContentHash,omitempty"`
	// The element type. Currently, only text is supported.
	//
	// example:
	//
	// object_type
	ElementType *string `json:"elementType,omitempty" xml:"elementType,omitempty"`
	// The last modified time in ISO 8601 format.
	//
	// example:
	//
	// 2026-09-08T10:30:00+00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The graph name.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The status code description.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The operation type.
	//
	// example:
	//
	// UPDATE
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// customer
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// object
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s SaveGraphDraftResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveGraphDraftResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveGraphDraftResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveGraphDraftResourceResponseBody) GetDraftChangeId() *int64 {
	return s.DraftChangeId
}

func (s *SaveGraphDraftResourceResponseBody) GetDraftContentHash() *string {
	return s.DraftContentHash
}

func (s *SaveGraphDraftResourceResponseBody) GetElementType() *string {
	return s.ElementType
}

func (s *SaveGraphDraftResourceResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *SaveGraphDraftResourceResponseBody) GetGraphName() *string {
	return s.GraphName
}

func (s *SaveGraphDraftResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveGraphDraftResourceResponseBody) GetOperationType() *string {
	return s.OperationType
}

func (s *SaveGraphDraftResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveGraphDraftResourceResponseBody) GetResourceName() *string {
	return s.ResourceName
}

func (s *SaveGraphDraftResourceResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *SaveGraphDraftResourceResponseBody) SetCode(v string) *SaveGraphDraftResourceResponseBody {
	s.Code = &v
	return s
}

func (s *SaveGraphDraftResourceResponseBody) SetDraftChangeId(v int64) *SaveGraphDraftResourceResponseBody {
	s.DraftChangeId = &v
	return s
}

func (s *SaveGraphDraftResourceResponseBody) SetDraftContentHash(v string) *SaveGraphDraftResourceResponseBody {
	s.DraftContentHash = &v
	return s
}

func (s *SaveGraphDraftResourceResponseBody) SetElementType(v string) *SaveGraphDraftResourceResponseBody {
	s.ElementType = &v
	return s
}

func (s *SaveGraphDraftResourceResponseBody) SetGmtModified(v string) *SaveGraphDraftResourceResponseBody {
	s.GmtModified = &v
	return s
}

func (s *SaveGraphDraftResourceResponseBody) SetGraphName(v string) *SaveGraphDraftResourceResponseBody {
	s.GraphName = &v
	return s
}

func (s *SaveGraphDraftResourceResponseBody) SetMessage(v string) *SaveGraphDraftResourceResponseBody {
	s.Message = &v
	return s
}

func (s *SaveGraphDraftResourceResponseBody) SetOperationType(v string) *SaveGraphDraftResourceResponseBody {
	s.OperationType = &v
	return s
}

func (s *SaveGraphDraftResourceResponseBody) SetRequestId(v string) *SaveGraphDraftResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveGraphDraftResourceResponseBody) SetResourceName(v string) *SaveGraphDraftResourceResponseBody {
	s.ResourceName = &v
	return s
}

func (s *SaveGraphDraftResourceResponseBody) SetResourceType(v string) *SaveGraphDraftResourceResponseBody {
	s.ResourceType = &v
	return s
}

func (s *SaveGraphDraftResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
