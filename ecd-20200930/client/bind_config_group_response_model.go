// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindConfigGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindConfigGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindConfigGroupResponse
	GetStatusCode() *int32
	SetBody(v *BindConfigGroupResponseBody) *BindConfigGroupResponse
	GetBody() *BindConfigGroupResponseBody
}

type BindConfigGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindConfigGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindConfigGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s BindConfigGroupResponse) GoString() string {
	return s.String()
}

func (s *BindConfigGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindConfigGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindConfigGroupResponse) GetBody() *BindConfigGroupResponseBody {
	return s.Body
}

func (s *BindConfigGroupResponse) SetHeaders(v map[string]*string) *BindConfigGroupResponse {
	s.Headers = v
	return s
}

func (s *BindConfigGroupResponse) SetStatusCode(v int32) *BindConfigGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *BindConfigGroupResponse) SetBody(v *BindConfigGroupResponseBody) *BindConfigGroupResponse {
	s.Body = v
	return s
}

func (s *BindConfigGroupResponse) Validate() error {
	return dara.Validate(s)
}
