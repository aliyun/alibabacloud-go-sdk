// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOncallScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOncallScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOncallScheduleResponse
	GetStatusCode() *int32
	SetBody(v *GetOncallScheduleResponseBody) *GetOncallScheduleResponse
	GetBody() *GetOncallScheduleResponseBody
}

type GetOncallScheduleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOncallScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOncallScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOncallScheduleResponse) GoString() string {
	return s.String()
}

func (s *GetOncallScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOncallScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOncallScheduleResponse) GetBody() *GetOncallScheduleResponseBody {
	return s.Body
}

func (s *GetOncallScheduleResponse) SetHeaders(v map[string]*string) *GetOncallScheduleResponse {
	s.Headers = v
	return s
}

func (s *GetOncallScheduleResponse) SetStatusCode(v int32) *GetOncallScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOncallScheduleResponse) SetBody(v *GetOncallScheduleResponseBody) *GetOncallScheduleResponse {
	s.Body = v
	return s
}

func (s *GetOncallScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
