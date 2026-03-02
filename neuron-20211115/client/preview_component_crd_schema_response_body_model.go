// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewComponentCrdSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PreviewComponentCrdSchemaResponseBody
	GetRequestId() *string
	SetSchema(v string) *PreviewComponentCrdSchemaResponseBody
	GetSchema() *string
}

type PreviewComponentCrdSchemaResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Schema    *string `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s PreviewComponentCrdSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewComponentCrdSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewComponentCrdSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewComponentCrdSchemaResponseBody) GetSchema() *string {
	return s.Schema
}

func (s *PreviewComponentCrdSchemaResponseBody) SetRequestId(v string) *PreviewComponentCrdSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewComponentCrdSchemaResponseBody) SetSchema(v string) *PreviewComponentCrdSchemaResponseBody {
	s.Schema = &v
	return s
}

func (s *PreviewComponentCrdSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}
