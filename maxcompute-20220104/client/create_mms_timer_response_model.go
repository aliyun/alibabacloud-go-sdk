// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMmsTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMmsTimerResponse
	GetStatusCode() *int32
	SetBody(v *CreateMmsTimerResponseBody) *CreateMmsTimerResponse
	GetBody() *CreateMmsTimerResponseBody
}

type CreateMmsTimerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMmsTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMmsTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsTimerResponse) GoString() string {
	return s.String()
}

func (s *CreateMmsTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMmsTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMmsTimerResponse) GetBody() *CreateMmsTimerResponseBody {
	return s.Body
}

func (s *CreateMmsTimerResponse) SetHeaders(v map[string]*string) *CreateMmsTimerResponse {
	s.Headers = v
	return s
}

func (s *CreateMmsTimerResponse) SetStatusCode(v int32) *CreateMmsTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMmsTimerResponse) SetBody(v *CreateMmsTimerResponseBody) *CreateMmsTimerResponse {
	s.Body = v
	return s
}

func (s *CreateMmsTimerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
