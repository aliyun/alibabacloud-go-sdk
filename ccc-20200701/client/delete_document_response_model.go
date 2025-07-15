// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDocumentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDocumentResponseBody) *DeleteDocumentResponse
	GetBody() *DeleteDocumentResponseBody
}

type DeleteDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentResponse) GoString() string {
	return s.String()
}

func (s *DeleteDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDocumentResponse) GetBody() *DeleteDocumentResponseBody {
	return s.Body
}

func (s *DeleteDocumentResponse) SetHeaders(v map[string]*string) *DeleteDocumentResponse {
	s.Headers = v
	return s
}

func (s *DeleteDocumentResponse) SetStatusCode(v int32) *DeleteDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDocumentResponse) SetBody(v *DeleteDocumentResponseBody) *DeleteDocumentResponse {
	s.Body = v
	return s
}

func (s *DeleteDocumentResponse) Validate() error {
	return dara.Validate(s)
}
