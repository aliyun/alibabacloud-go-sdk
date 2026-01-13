// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficControlTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTrafficControlTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTrafficControlTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTrafficControlTaskResponseBody) *UpdateTrafficControlTaskResponse
	GetBody() *UpdateTrafficControlTaskResponseBody
}

type UpdateTrafficControlTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrafficControlTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTrafficControlTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTrafficControlTaskResponse) GetBody() *UpdateTrafficControlTaskResponseBody {
	return s.Body
}

func (s *UpdateTrafficControlTaskResponse) SetHeaders(v map[string]*string) *UpdateTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrafficControlTaskResponse) SetStatusCode(v int32) *UpdateTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrafficControlTaskResponse) SetBody(v *UpdateTrafficControlTaskResponseBody) *UpdateTrafficControlTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateTrafficControlTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
