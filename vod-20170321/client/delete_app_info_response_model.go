// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppInfoResponseBody) *DeleteAppInfoResponse
	GetBody() *DeleteAppInfoResponseBody
}

type DeleteAppInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppInfoResponse) GetBody() *DeleteAppInfoResponseBody {
	return s.Body
}

func (s *DeleteAppInfoResponse) SetHeaders(v map[string]*string) *DeleteAppInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppInfoResponse) SetStatusCode(v int32) *DeleteAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppInfoResponse) SetBody(v *DeleteAppInfoResponseBody) *DeleteAppInfoResponse {
	s.Body = v
	return s
}

func (s *DeleteAppInfoResponse) Validate() error {
	return dara.Validate(s)
}
