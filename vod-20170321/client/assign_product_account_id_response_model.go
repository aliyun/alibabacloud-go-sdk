// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignProductAccountIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignProductAccountIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignProductAccountIdResponse
	GetStatusCode() *int32
	SetBody(v *AssignProductAccountIdResponseBody) *AssignProductAccountIdResponse
	GetBody() *AssignProductAccountIdResponseBody
}

type AssignProductAccountIdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignProductAccountIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignProductAccountIdResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignProductAccountIdResponse) GoString() string {
	return s.String()
}

func (s *AssignProductAccountIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignProductAccountIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignProductAccountIdResponse) GetBody() *AssignProductAccountIdResponseBody {
	return s.Body
}

func (s *AssignProductAccountIdResponse) SetHeaders(v map[string]*string) *AssignProductAccountIdResponse {
	s.Headers = v
	return s
}

func (s *AssignProductAccountIdResponse) SetStatusCode(v int32) *AssignProductAccountIdResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignProductAccountIdResponse) SetBody(v *AssignProductAccountIdResponseBody) *AssignProductAccountIdResponse {
	s.Body = v
	return s
}

func (s *AssignProductAccountIdResponse) Validate() error {
	return dara.Validate(s)
}
