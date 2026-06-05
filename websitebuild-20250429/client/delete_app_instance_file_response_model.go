// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInstanceFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppInstanceFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppInstanceFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppInstanceFileResponseBody) *DeleteAppInstanceFileResponse
	GetBody() *DeleteAppInstanceFileResponseBody
}

type DeleteAppInstanceFileResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppInstanceFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppInstanceFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstanceFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppInstanceFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppInstanceFileResponse) GetBody() *DeleteAppInstanceFileResponseBody {
	return s.Body
}

func (s *DeleteAppInstanceFileResponse) SetHeaders(v map[string]*string) *DeleteAppInstanceFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppInstanceFileResponse) SetStatusCode(v int32) *DeleteAppInstanceFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppInstanceFileResponse) SetBody(v *DeleteAppInstanceFileResponseBody) *DeleteAppInstanceFileResponse {
	s.Body = v
	return s
}

func (s *DeleteAppInstanceFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
