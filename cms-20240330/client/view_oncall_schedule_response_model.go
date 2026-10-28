// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewOncallScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ViewOncallScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ViewOncallScheduleResponse
	GetStatusCode() *int32
	SetBody(v *ViewOncallScheduleResponseBody) *ViewOncallScheduleResponse
	GetBody() *ViewOncallScheduleResponseBody
}

type ViewOncallScheduleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ViewOncallScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ViewOncallScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s ViewOncallScheduleResponse) GoString() string {
	return s.String()
}

func (s *ViewOncallScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ViewOncallScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ViewOncallScheduleResponse) GetBody() *ViewOncallScheduleResponseBody {
	return s.Body
}

func (s *ViewOncallScheduleResponse) SetHeaders(v map[string]*string) *ViewOncallScheduleResponse {
	s.Headers = v
	return s
}

func (s *ViewOncallScheduleResponse) SetStatusCode(v int32) *ViewOncallScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *ViewOncallScheduleResponse) SetBody(v *ViewOncallScheduleResponseBody) *ViewOncallScheduleResponse {
	s.Body = v
	return s
}

func (s *ViewOncallScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
