// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppInstanceResponseBody) *DeleteAppInstanceResponse
	GetBody() *DeleteAppInstanceResponseBody
}

type DeleteAppInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppInstanceResponse) GetBody() *DeleteAppInstanceResponseBody {
	return s.Body
}

func (s *DeleteAppInstanceResponse) SetHeaders(v map[string]*string) *DeleteAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppInstanceResponse) SetStatusCode(v int32) *DeleteAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppInstanceResponse) SetBody(v *DeleteAppInstanceResponseBody) *DeleteAppInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteAppInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
