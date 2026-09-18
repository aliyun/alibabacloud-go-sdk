// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGroupDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGroupDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *ListGroupDirectoryResponseBody) *ListGroupDirectoryResponse
	GetBody() *ListGroupDirectoryResponseBody
}

type ListGroupDirectoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGroupDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ListGroupDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGroupDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGroupDirectoryResponse) GetBody() *ListGroupDirectoryResponseBody {
	return s.Body
}

func (s *ListGroupDirectoryResponse) SetHeaders(v map[string]*string) *ListGroupDirectoryResponse {
	s.Headers = v
	return s
}

func (s *ListGroupDirectoryResponse) SetStatusCode(v int32) *ListGroupDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupDirectoryResponse) SetBody(v *ListGroupDirectoryResponseBody) *ListGroupDirectoryResponse {
	s.Body = v
	return s
}

func (s *ListGroupDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
