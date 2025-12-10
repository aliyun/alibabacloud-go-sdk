// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitResourceDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitResourceDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitResourceDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *InitResourceDirectoryResponseBody) *InitResourceDirectoryResponse
	GetBody() *InitResourceDirectoryResponseBody
}

type InitResourceDirectoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitResourceDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s InitResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *InitResourceDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitResourceDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitResourceDirectoryResponse) GetBody() *InitResourceDirectoryResponseBody {
	return s.Body
}

func (s *InitResourceDirectoryResponse) SetHeaders(v map[string]*string) *InitResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *InitResourceDirectoryResponse) SetStatusCode(v int32) *InitResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *InitResourceDirectoryResponse) SetBody(v *InitResourceDirectoryResponseBody) *InitResourceDirectoryResponse {
	s.Body = v
	return s
}

func (s *InitResourceDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
