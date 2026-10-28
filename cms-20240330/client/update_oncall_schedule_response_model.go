// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOncallScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOncallScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOncallScheduleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOncallScheduleResponseBody) *UpdateOncallScheduleResponse
	GetBody() *UpdateOncallScheduleResponseBody
}

type UpdateOncallScheduleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOncallScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOncallScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOncallScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpdateOncallScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOncallScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOncallScheduleResponse) GetBody() *UpdateOncallScheduleResponseBody {
	return s.Body
}

func (s *UpdateOncallScheduleResponse) SetHeaders(v map[string]*string) *UpdateOncallScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpdateOncallScheduleResponse) SetStatusCode(v int32) *UpdateOncallScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOncallScheduleResponse) SetBody(v *UpdateOncallScheduleResponseBody) *UpdateOncallScheduleResponse {
	s.Body = v
	return s
}

func (s *UpdateOncallScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
