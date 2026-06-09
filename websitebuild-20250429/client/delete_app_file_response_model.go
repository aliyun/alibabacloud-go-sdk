// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppFileResponseBody) *DeleteAppFileResponse
	GetBody() *DeleteAppFileResponseBody
}

type DeleteAppFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppFileResponse) GetBody() *DeleteAppFileResponseBody {
	return s.Body
}

func (s *DeleteAppFileResponse) SetHeaders(v map[string]*string) *DeleteAppFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppFileResponse) SetStatusCode(v int32) *DeleteAppFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppFileResponse) SetBody(v *DeleteAppFileResponseBody) *DeleteAppFileResponse {
	s.Body = v
	return s
}

func (s *DeleteAppFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
