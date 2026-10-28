// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOncallScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOncallScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOncallScheduleResponse
	GetStatusCode() *int32
	SetBody(v *CreateOncallScheduleResponseBody) *CreateOncallScheduleResponse
	GetBody() *CreateOncallScheduleResponseBody
}

type CreateOncallScheduleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOncallScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOncallScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOncallScheduleResponse) GoString() string {
	return s.String()
}

func (s *CreateOncallScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOncallScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOncallScheduleResponse) GetBody() *CreateOncallScheduleResponseBody {
	return s.Body
}

func (s *CreateOncallScheduleResponse) SetHeaders(v map[string]*string) *CreateOncallScheduleResponse {
	s.Headers = v
	return s
}

func (s *CreateOncallScheduleResponse) SetStatusCode(v int32) *CreateOncallScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOncallScheduleResponse) SetBody(v *CreateOncallScheduleResponseBody) *CreateOncallScheduleResponse {
	s.Body = v
	return s
}

func (s *CreateOncallScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
