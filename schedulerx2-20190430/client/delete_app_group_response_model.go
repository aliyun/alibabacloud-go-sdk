// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppGroupResponseBody) *DeleteAppGroupResponse
	GetBody() *DeleteAppGroupResponseBody
}

type DeleteAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppGroupResponse) GetBody() *DeleteAppGroupResponseBody {
	return s.Body
}

func (s *DeleteAppGroupResponse) SetHeaders(v map[string]*string) *DeleteAppGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppGroupResponse) SetStatusCode(v int32) *DeleteAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppGroupResponse) SetBody(v *DeleteAppGroupResponseBody) *DeleteAppGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteAppGroupResponse) Validate() error {
	return dara.Validate(s)
}
