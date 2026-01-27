// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppResponseBody) *DeleteAppResponse
	GetBody() *DeleteAppResponseBody
}

type DeleteAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppResponse) GetBody() *DeleteAppResponseBody {
	return s.Body
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetStatusCode(v int32) *DeleteAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

func (s *DeleteAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
