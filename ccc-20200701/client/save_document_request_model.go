// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentId(v string) *SaveDocumentRequest
	GetDocumentId() *string
	SetDocumentJson(v string) *SaveDocumentRequest
	GetDocumentJson() *string
	SetInstanceId(v string) *SaveDocumentRequest
	GetInstanceId() *string
	SetRequestId(v string) *SaveDocumentRequest
	GetRequestId() *string
	SetSchemaId(v string) *SaveDocumentRequest
	GetSchemaId() *string
}

type SaveDocumentRequest struct {
	// example:
	//
	// xxx
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"name":"tom"}
	DocumentJson *string `json:"DocumentJson,omitempty" xml:"DocumentJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9cfad875-6260-4a53-ab6e-b13e3fb31f7d
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// EAF3C248-E123-441B-A545-B6CD02E98EED
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

func (s SaveDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveDocumentRequest) GoString() string {
	return s.String()
}

func (s *SaveDocumentRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *SaveDocumentRequest) GetDocumentJson() *string {
	return s.DocumentJson
}

func (s *SaveDocumentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveDocumentRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveDocumentRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *SaveDocumentRequest) SetDocumentId(v string) *SaveDocumentRequest {
	s.DocumentId = &v
	return s
}

func (s *SaveDocumentRequest) SetDocumentJson(v string) *SaveDocumentRequest {
	s.DocumentJson = &v
	return s
}

func (s *SaveDocumentRequest) SetInstanceId(v string) *SaveDocumentRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveDocumentRequest) SetRequestId(v string) *SaveDocumentRequest {
	s.RequestId = &v
	return s
}

func (s *SaveDocumentRequest) SetSchemaId(v string) *SaveDocumentRequest {
	s.SchemaId = &v
	return s
}

func (s *SaveDocumentRequest) Validate() error {
	return dara.Validate(s)
}
