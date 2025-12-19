// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitAppFailOverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitAppFailOverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitAppFailOverResponse
	GetStatusCode() *int32
	SetBody(v *InitAppFailOverResponseBody) *InitAppFailOverResponse
	GetBody() *InitAppFailOverResponseBody
}

type InitAppFailOverResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitAppFailOverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitAppFailOverResponse) String() string {
	return dara.Prettify(s)
}

func (s InitAppFailOverResponse) GoString() string {
	return s.String()
}

func (s *InitAppFailOverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitAppFailOverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitAppFailOverResponse) GetBody() *InitAppFailOverResponseBody {
	return s.Body
}

func (s *InitAppFailOverResponse) SetHeaders(v map[string]*string) *InitAppFailOverResponse {
	s.Headers = v
	return s
}

func (s *InitAppFailOverResponse) SetStatusCode(v int32) *InitAppFailOverResponse {
	s.StatusCode = &v
	return s
}

func (s *InitAppFailOverResponse) SetBody(v *InitAppFailOverResponseBody) *InitAppFailOverResponse {
	s.Body = v
	return s
}

func (s *InitAppFailOverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
