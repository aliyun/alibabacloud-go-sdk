// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppListResponseBody) *DeleteAppListResponse
	GetBody() *DeleteAppListResponseBody
}

type DeleteAppListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppListResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppListResponse) GetBody() *DeleteAppListResponseBody {
	return s.Body
}

func (s *DeleteAppListResponse) SetHeaders(v map[string]*string) *DeleteAppListResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppListResponse) SetStatusCode(v int32) *DeleteAppListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppListResponse) SetBody(v *DeleteAppListResponseBody) *DeleteAppListResponse {
	s.Body = v
	return s
}

func (s *DeleteAppListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
