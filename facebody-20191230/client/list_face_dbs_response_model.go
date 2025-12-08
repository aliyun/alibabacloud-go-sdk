// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFaceDbsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFaceDbsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFaceDbsResponse
	GetStatusCode() *int32
	SetBody(v *ListFaceDbsResponseBody) *ListFaceDbsResponse
	GetBody() *ListFaceDbsResponseBody
}

type ListFaceDbsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFaceDbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFaceDbsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFaceDbsResponse) GoString() string {
	return s.String()
}

func (s *ListFaceDbsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFaceDbsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFaceDbsResponse) GetBody() *ListFaceDbsResponseBody {
	return s.Body
}

func (s *ListFaceDbsResponse) SetHeaders(v map[string]*string) *ListFaceDbsResponse {
	s.Headers = v
	return s
}

func (s *ListFaceDbsResponse) SetStatusCode(v int32) *ListFaceDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFaceDbsResponse) SetBody(v *ListFaceDbsResponseBody) *ListFaceDbsResponse {
	s.Body = v
	return s
}

func (s *ListFaceDbsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
