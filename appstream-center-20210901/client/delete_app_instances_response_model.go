// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppInstancesResponseBody) *DeleteAppInstancesResponse
	GetBody() *DeleteAppInstancesResponseBody
}

type DeleteAppInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppInstancesResponse) GetBody() *DeleteAppInstancesResponseBody {
	return s.Body
}

func (s *DeleteAppInstancesResponse) SetHeaders(v map[string]*string) *DeleteAppInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppInstancesResponse) SetStatusCode(v int32) *DeleteAppInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppInstancesResponse) SetBody(v *DeleteAppInstancesResponseBody) *DeleteAppInstancesResponse {
	s.Body = v
	return s
}

func (s *DeleteAppInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
