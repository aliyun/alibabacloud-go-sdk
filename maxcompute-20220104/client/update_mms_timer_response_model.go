// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMmsTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMmsTimerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMmsTimerResponseBody) *UpdateMmsTimerResponse
	GetBody() *UpdateMmsTimerResponseBody
}

type UpdateMmsTimerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmsTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmsTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsTimerResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmsTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMmsTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMmsTimerResponse) GetBody() *UpdateMmsTimerResponseBody {
	return s.Body
}

func (s *UpdateMmsTimerResponse) SetHeaders(v map[string]*string) *UpdateMmsTimerResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmsTimerResponse) SetStatusCode(v int32) *UpdateMmsTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmsTimerResponse) SetBody(v *UpdateMmsTimerResponseBody) *UpdateMmsTimerResponse {
	s.Body = v
	return s
}

func (s *UpdateMmsTimerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
