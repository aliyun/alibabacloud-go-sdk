// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentIds(v []*string) *DeleteDocumentsRequest
	GetDocumentIds() []*string
	SetInstanceId(v string) *DeleteDocumentsRequest
	GetInstanceId() *string
	SetRequestId(v string) *DeleteDocumentsRequest
	GetRequestId() *string
	SetSchemaId(v string) *DeleteDocumentsRequest
	GetSchemaId() *string
}

type DeleteDocumentsRequest struct {
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 9cfad875-6260-4a53-ab6e-b13e3fb31f7d
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0630E5DF-CEB0-445B-8626-D5C7481181C3
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

func (s DeleteDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentsRequest) GetDocumentIds() []*string {
	return s.DocumentIds
}

func (s *DeleteDocumentsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDocumentsRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocumentsRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *DeleteDocumentsRequest) SetDocumentIds(v []*string) *DeleteDocumentsRequest {
	s.DocumentIds = v
	return s
}

func (s *DeleteDocumentsRequest) SetInstanceId(v string) *DeleteDocumentsRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDocumentsRequest) SetRequestId(v string) *DeleteDocumentsRequest {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentsRequest) SetSchemaId(v string) *DeleteDocumentsRequest {
	s.SchemaId = &v
	return s
}

func (s *DeleteDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
