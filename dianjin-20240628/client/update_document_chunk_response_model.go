// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDocumentChunkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDocumentChunkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDocumentChunkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDocumentChunkResponseBody) *UpdateDocumentChunkResponse
	GetBody() *UpdateDocumentChunkResponseBody
}

type UpdateDocumentChunkResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDocumentChunkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDocumentChunkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocumentChunkResponse) GoString() string {
	return s.String()
}

func (s *UpdateDocumentChunkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDocumentChunkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDocumentChunkResponse) GetBody() *UpdateDocumentChunkResponseBody {
	return s.Body
}

func (s *UpdateDocumentChunkResponse) SetHeaders(v map[string]*string) *UpdateDocumentChunkResponse {
	s.Headers = v
	return s
}

func (s *UpdateDocumentChunkResponse) SetStatusCode(v int32) *UpdateDocumentChunkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDocumentChunkResponse) SetBody(v *UpdateDocumentChunkResponseBody) *UpdateDocumentChunkResponse {
	s.Body = v
	return s
}

func (s *UpdateDocumentChunkResponse) Validate() error {
	return dara.Validate(s)
}
