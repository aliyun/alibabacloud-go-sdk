// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppLayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppLayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppLayoutResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppLayoutResponseBody) *DeleteAppLayoutResponse
	GetBody() *DeleteAppLayoutResponseBody
}

type DeleteAppLayoutResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppLayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppLayoutResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppLayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppLayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppLayoutResponse) GetBody() *DeleteAppLayoutResponseBody {
	return s.Body
}

func (s *DeleteAppLayoutResponse) SetHeaders(v map[string]*string) *DeleteAppLayoutResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppLayoutResponse) SetStatusCode(v int32) *DeleteAppLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppLayoutResponse) SetBody(v *DeleteAppLayoutResponseBody) *DeleteAppLayoutResponse {
	s.Body = v
	return s
}

func (s *DeleteAppLayoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
