// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDocumentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDocumentsResponse
	GetStatusCode() *int32
	SetBody(v *ListDocumentsResponseBody) *ListDocumentsResponse
	GetBody() *ListDocumentsResponseBody
}

type ListDocumentsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDocumentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsResponse) GoString() string {
	return s.String()
}

func (s *ListDocumentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDocumentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDocumentsResponse) GetBody() *ListDocumentsResponseBody {
	return s.Body
}

func (s *ListDocumentsResponse) SetHeaders(v map[string]*string) *ListDocumentsResponse {
	s.Headers = v
	return s
}

func (s *ListDocumentsResponse) SetStatusCode(v int32) *ListDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDocumentsResponse) SetBody(v *ListDocumentsResponseBody) *ListDocumentsResponse {
	s.Body = v
	return s
}

func (s *ListDocumentsResponse) Validate() error {
	return dara.Validate(s)
}
