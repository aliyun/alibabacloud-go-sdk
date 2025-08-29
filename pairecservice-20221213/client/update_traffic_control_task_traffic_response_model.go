// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficControlTaskTrafficResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTrafficControlTaskTrafficResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTrafficControlTaskTrafficResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTrafficControlTaskTrafficResponseBody) *UpdateTrafficControlTaskTrafficResponse
	GetBody() *UpdateTrafficControlTaskTrafficResponseBody
}

type UpdateTrafficControlTaskTrafficResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrafficControlTaskTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrafficControlTaskTrafficResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficControlTaskTrafficResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskTrafficResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTrafficControlTaskTrafficResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTrafficControlTaskTrafficResponse) GetBody() *UpdateTrafficControlTaskTrafficResponseBody {
	return s.Body
}

func (s *UpdateTrafficControlTaskTrafficResponse) SetHeaders(v map[string]*string) *UpdateTrafficControlTaskTrafficResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrafficControlTaskTrafficResponse) SetStatusCode(v int32) *UpdateTrafficControlTaskTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficResponse) SetBody(v *UpdateTrafficControlTaskTrafficResponseBody) *UpdateTrafficControlTaskTrafficResponse {
	s.Body = v
	return s
}

func (s *UpdateTrafficControlTaskTrafficResponse) Validate() error {
	return dara.Validate(s)
}
