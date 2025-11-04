// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlarmResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlarmResponseBody) *CreateAlarmResponse
	GetBody() *CreateAlarmResponseBody
}

type CreateAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmResponse) GoString() string {
	return s.String()
}

func (s *CreateAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlarmResponse) GetBody() *CreateAlarmResponseBody {
	return s.Body
}

func (s *CreateAlarmResponse) SetHeaders(v map[string]*string) *CreateAlarmResponse {
	s.Headers = v
	return s
}

func (s *CreateAlarmResponse) SetStatusCode(v int32) *CreateAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlarmResponse) SetBody(v *CreateAlarmResponseBody) *CreateAlarmResponse {
	s.Body = v
	return s
}

func (s *CreateAlarmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
