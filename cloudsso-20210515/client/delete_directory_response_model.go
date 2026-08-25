// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDirectoryResponseBody) *DeleteDirectoryResponse
	GetBody() *DeleteDirectoryResponseBody
}

type DeleteDirectoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDirectoryResponse) GetBody() *DeleteDirectoryResponseBody {
	return s.Body
}

func (s *DeleteDirectoryResponse) SetHeaders(v map[string]*string) *DeleteDirectoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteDirectoryResponse) SetStatusCode(v int32) *DeleteDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDirectoryResponse) SetBody(v *DeleteDirectoryResponseBody) *DeleteDirectoryResponse {
	s.Body = v
	return s
}

func (s *DeleteDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
