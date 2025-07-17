// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIndexDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIndexDocumentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIndexDocumentResponseBody) *DeleteIndexDocumentResponse
	GetBody() *DeleteIndexDocumentResponseBody
}

type DeleteIndexDocumentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIndexDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIndexDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexDocumentResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIndexDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIndexDocumentResponse) GetBody() *DeleteIndexDocumentResponseBody {
	return s.Body
}

func (s *DeleteIndexDocumentResponse) SetHeaders(v map[string]*string) *DeleteIndexDocumentResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexDocumentResponse) SetStatusCode(v int32) *DeleteIndexDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndexDocumentResponse) SetBody(v *DeleteIndexDocumentResponseBody) *DeleteIndexDocumentResponse {
	s.Body = v
	return s
}

func (s *DeleteIndexDocumentResponse) Validate() error {
	return dara.Validate(s)
}
