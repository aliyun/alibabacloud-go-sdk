// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentChunksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDocumentChunksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDocumentChunksResponse
	GetStatusCode() *int32
	SetBody(v *ListDocumentChunksResponseBody) *ListDocumentChunksResponse
	GetBody() *ListDocumentChunksResponseBody
}

type ListDocumentChunksResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDocumentChunksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDocumentChunksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentChunksResponse) GoString() string {
	return s.String()
}

func (s *ListDocumentChunksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDocumentChunksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDocumentChunksResponse) GetBody() *ListDocumentChunksResponseBody {
	return s.Body
}

func (s *ListDocumentChunksResponse) SetHeaders(v map[string]*string) *ListDocumentChunksResponse {
	s.Headers = v
	return s
}

func (s *ListDocumentChunksResponse) SetStatusCode(v int32) *ListDocumentChunksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDocumentChunksResponse) SetBody(v *ListDocumentChunksResponseBody) *ListDocumentChunksResponse {
	s.Body = v
	return s
}

func (s *ListDocumentChunksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
