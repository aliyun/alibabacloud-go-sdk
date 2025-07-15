// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationGroupResponseBody) *DeleteApplicationGroupResponse
	GetBody() *DeleteApplicationGroupResponseBody
}

type DeleteApplicationGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationGroupResponse) GetBody() *DeleteApplicationGroupResponseBody {
	return s.Body
}

func (s *DeleteApplicationGroupResponse) SetHeaders(v map[string]*string) *DeleteApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationGroupResponse) SetStatusCode(v int32) *DeleteApplicationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationGroupResponse) SetBody(v *DeleteApplicationGroupResponseBody) *DeleteApplicationGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationGroupResponse) Validate() error {
	return dara.Validate(s)
}
