// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppCodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppCodeResponseBody) *DeleteAppCodeResponse
	GetBody() *DeleteAppCodeResponseBody
}

type DeleteAppCodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppCodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppCodeResponse) GetBody() *DeleteAppCodeResponseBody {
	return s.Body
}

func (s *DeleteAppCodeResponse) SetHeaders(v map[string]*string) *DeleteAppCodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppCodeResponse) SetStatusCode(v int32) *DeleteAppCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppCodeResponse) SetBody(v *DeleteAppCodeResponseBody) *DeleteAppCodeResponse {
	s.Body = v
	return s
}

func (s *DeleteAppCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
