// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentId(v string) *GetDocumentRequest
	GetDocumentId() *string
	SetInstanceId(v string) *GetDocumentRequest
	GetInstanceId() *string
	SetRequestId(v string) *GetDocumentRequest
	GetRequestId() *string
	SetSchemaId(v string) *GetDocumentRequest
	GetSchemaId() *string
}

type GetDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 51e155ce-3747-4f21-b402-13c69597b920
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// FCEFE806-E67C-577E-865B-4ED398F2F648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// schema id
	//
	// This parameter is required.
	//
	// example:
	//
	// profile
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *GetDocumentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDocumentRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetDocumentRequest) SetDocumentId(v string) *GetDocumentRequest {
	s.DocumentId = &v
	return s
}

func (s *GetDocumentRequest) SetInstanceId(v string) *GetDocumentRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDocumentRequest) SetRequestId(v string) *GetDocumentRequest {
	s.RequestId = &v
	return s
}

func (s *GetDocumentRequest) SetSchemaId(v string) *GetDocumentRequest {
	s.SchemaId = &v
	return s
}

func (s *GetDocumentRequest) Validate() error {
	return dara.Validate(s)
}
