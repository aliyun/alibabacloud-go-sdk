// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentId(v string) *DeleteDocumentRequest
	GetDocumentId() *string
	SetInstanceId(v string) *DeleteDocumentRequest
	GetInstanceId() *string
	SetRequestId(v string) *DeleteDocumentRequest
	GetRequestId() *string
	SetSchemaId(v string) *DeleteDocumentRequest
	GetSchemaId() *string
}

type DeleteDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d278629c-c687-4aa3-b044-4fe9b012****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// DE803553-8AA9-4B9D-9E4E-A82BC69EDCEE
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

func (s DeleteDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *DeleteDocumentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDocumentRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocumentRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *DeleteDocumentRequest) SetDocumentId(v string) *DeleteDocumentRequest {
	s.DocumentId = &v
	return s
}

func (s *DeleteDocumentRequest) SetInstanceId(v string) *DeleteDocumentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDocumentRequest) SetRequestId(v string) *DeleteDocumentRequest {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentRequest) SetSchemaId(v string) *DeleteDocumentRequest {
	s.SchemaId = &v
	return s
}

func (s *DeleteDocumentRequest) Validate() error {
	return dara.Validate(s)
}
