// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppInstanceGroupResponseBody) *DeleteAppInstanceGroupResponse
	GetBody() *DeleteAppInstanceGroupResponseBody
}

type DeleteAppInstanceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppInstanceGroupResponse) GetBody() *DeleteAppInstanceGroupResponseBody {
	return s.Body
}

func (s *DeleteAppInstanceGroupResponse) SetHeaders(v map[string]*string) *DeleteAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppInstanceGroupResponse) SetStatusCode(v int32) *DeleteAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppInstanceGroupResponse) SetBody(v *DeleteAppInstanceGroupResponseBody) *DeleteAppInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteAppInstanceGroupResponse) Validate() error {
	return dara.Validate(s)
}
