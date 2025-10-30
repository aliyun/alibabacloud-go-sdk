// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDocumentCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDocumentCollectionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDocumentCollectionResponseBody) *DeleteDocumentCollectionResponse
	GetBody() *DeleteDocumentCollectionResponseBody
}

type DeleteDocumentCollectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDocumentCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDocumentCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentCollectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteDocumentCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDocumentCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDocumentCollectionResponse) GetBody() *DeleteDocumentCollectionResponseBody {
	return s.Body
}

func (s *DeleteDocumentCollectionResponse) SetHeaders(v map[string]*string) *DeleteDocumentCollectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteDocumentCollectionResponse) SetStatusCode(v int32) *DeleteDocumentCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDocumentCollectionResponse) SetBody(v *DeleteDocumentCollectionResponseBody) *DeleteDocumentCollectionResponse {
	s.Body = v
	return s
}

func (s *DeleteDocumentCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
