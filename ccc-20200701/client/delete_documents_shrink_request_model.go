// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentIdsShrink(v string) *DeleteDocumentsShrinkRequest
	GetDocumentIdsShrink() *string
	SetInstanceId(v string) *DeleteDocumentsShrinkRequest
	GetInstanceId() *string
	SetRequestId(v string) *DeleteDocumentsShrinkRequest
	GetRequestId() *string
	SetSchemaId(v string) *DeleteDocumentsShrinkRequest
	GetSchemaId() *string
}

type DeleteDocumentsShrinkRequest struct {
	DocumentIdsShrink *string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty"`
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

func (s DeleteDocumentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentsShrinkRequest) GetDocumentIdsShrink() *string {
	return s.DocumentIdsShrink
}

func (s *DeleteDocumentsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDocumentsShrinkRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocumentsShrinkRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *DeleteDocumentsShrinkRequest) SetDocumentIdsShrink(v string) *DeleteDocumentsShrinkRequest {
	s.DocumentIdsShrink = &v
	return s
}

func (s *DeleteDocumentsShrinkRequest) SetInstanceId(v string) *DeleteDocumentsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDocumentsShrinkRequest) SetRequestId(v string) *DeleteDocumentsShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentsShrinkRequest) SetSchemaId(v string) *DeleteDocumentsShrinkRequest {
	s.SchemaId = &v
	return s
}

func (s *DeleteDocumentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
