// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDocumentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDocumentResponseBody) *UpdateDocumentResponse
	GetBody() *UpdateDocumentResponseBody
}

type UpdateDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocumentResponse) GoString() string {
	return s.String()
}

func (s *UpdateDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDocumentResponse) GetBody() *UpdateDocumentResponseBody {
	return s.Body
}

func (s *UpdateDocumentResponse) SetHeaders(v map[string]*string) *UpdateDocumentResponse {
	s.Headers = v
	return s
}

func (s *UpdateDocumentResponse) SetStatusCode(v int32) *UpdateDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDocumentResponse) SetBody(v *UpdateDocumentResponseBody) *UpdateDocumentResponse {
	s.Body = v
	return s
}

func (s *UpdateDocumentResponse) Validate() error {
	return dara.Validate(s)
}
