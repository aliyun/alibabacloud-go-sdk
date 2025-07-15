// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDocumentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDocumentsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDocumentsResponseBody) *DeleteDocumentsResponse
	GetBody() *DeleteDocumentsResponseBody
}

type DeleteDocumentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDocumentsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDocumentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDocumentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDocumentsResponse) GetBody() *DeleteDocumentsResponseBody {
	return s.Body
}

func (s *DeleteDocumentsResponse) SetHeaders(v map[string]*string) *DeleteDocumentsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDocumentsResponse) SetStatusCode(v int32) *DeleteDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDocumentsResponse) SetBody(v *DeleteDocumentsResponseBody) *DeleteDocumentsResponse {
	s.Body = v
	return s
}

func (s *DeleteDocumentsResponse) Validate() error {
	return dara.Validate(s)
}
