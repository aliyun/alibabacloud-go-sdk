// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentChunksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDocumentChunksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDocumentChunksResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDocumentChunksResponseBody) *DeleteDocumentChunksResponse
	GetBody() *DeleteDocumentChunksResponseBody
}

type DeleteDocumentChunksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDocumentChunksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDocumentChunksResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentChunksResponse) GoString() string {
	return s.String()
}

func (s *DeleteDocumentChunksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDocumentChunksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDocumentChunksResponse) GetBody() *DeleteDocumentChunksResponseBody {
	return s.Body
}

func (s *DeleteDocumentChunksResponse) SetHeaders(v map[string]*string) *DeleteDocumentChunksResponse {
	s.Headers = v
	return s
}

func (s *DeleteDocumentChunksResponse) SetStatusCode(v int32) *DeleteDocumentChunksResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDocumentChunksResponse) SetBody(v *DeleteDocumentChunksResponseBody) *DeleteDocumentChunksResponse {
	s.Body = v
	return s
}

func (s *DeleteDocumentChunksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
