// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexDocumentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIndexDocumentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIndexDocumentsResponse
	GetStatusCode() *int32
	SetBody(v *ListIndexDocumentsResponseBody) *ListIndexDocumentsResponse
	GetBody() *ListIndexDocumentsResponseBody
}

type ListIndexDocumentsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIndexDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIndexDocumentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIndexDocumentsResponse) GoString() string {
	return s.String()
}

func (s *ListIndexDocumentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIndexDocumentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIndexDocumentsResponse) GetBody() *ListIndexDocumentsResponseBody {
	return s.Body
}

func (s *ListIndexDocumentsResponse) SetHeaders(v map[string]*string) *ListIndexDocumentsResponse {
	s.Headers = v
	return s
}

func (s *ListIndexDocumentsResponse) SetStatusCode(v int32) *ListIndexDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndexDocumentsResponse) SetBody(v *ListIndexDocumentsResponseBody) *ListIndexDocumentsResponse {
	s.Body = v
	return s
}

func (s *ListIndexDocumentsResponse) Validate() error {
	return dara.Validate(s)
}
