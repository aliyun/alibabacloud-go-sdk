// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMmsTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMmsTimerResponse
	GetStatusCode() *int32
	SetBody(v *GetMmsTimerResponseBody) *GetMmsTimerResponse
	GetBody() *GetMmsTimerResponseBody
}

type GetMmsTimerResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTimerResponse) GoString() string {
	return s.String()
}

func (s *GetMmsTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMmsTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMmsTimerResponse) GetBody() *GetMmsTimerResponseBody {
	return s.Body
}

func (s *GetMmsTimerResponse) SetHeaders(v map[string]*string) *GetMmsTimerResponse {
	s.Headers = v
	return s
}

func (s *GetMmsTimerResponse) SetStatusCode(v int32) *GetMmsTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsTimerResponse) SetBody(v *GetMmsTimerResponseBody) *GetMmsTimerResponse {
	s.Body = v
	return s
}

func (s *GetMmsTimerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
