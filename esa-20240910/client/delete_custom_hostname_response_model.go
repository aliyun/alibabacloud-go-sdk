// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomHostnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomHostnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomHostnameResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomHostnameResponseBody) *DeleteCustomHostnameResponse
	GetBody() *DeleteCustomHostnameResponseBody
}

type DeleteCustomHostnameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomHostnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomHostnameResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomHostnameResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomHostnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomHostnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomHostnameResponse) GetBody() *DeleteCustomHostnameResponseBody {
	return s.Body
}

func (s *DeleteCustomHostnameResponse) SetHeaders(v map[string]*string) *DeleteCustomHostnameResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomHostnameResponse) SetStatusCode(v int32) *DeleteCustomHostnameResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomHostnameResponse) SetBody(v *DeleteCustomHostnameResponseBody) *DeleteCustomHostnameResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomHostnameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
