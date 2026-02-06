// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaxAttemptsPerDayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMaxAttemptsPerDayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMaxAttemptsPerDayResponse
	GetStatusCode() *int32
	SetBody(v *GetMaxAttemptsPerDayResponseBody) *GetMaxAttemptsPerDayResponse
	GetBody() *GetMaxAttemptsPerDayResponseBody
}

type GetMaxAttemptsPerDayResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMaxAttemptsPerDayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMaxAttemptsPerDayResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMaxAttemptsPerDayResponse) GoString() string {
	return s.String()
}

func (s *GetMaxAttemptsPerDayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMaxAttemptsPerDayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMaxAttemptsPerDayResponse) GetBody() *GetMaxAttemptsPerDayResponseBody {
	return s.Body
}

func (s *GetMaxAttemptsPerDayResponse) SetHeaders(v map[string]*string) *GetMaxAttemptsPerDayResponse {
	s.Headers = v
	return s
}

func (s *GetMaxAttemptsPerDayResponse) SetStatusCode(v int32) *GetMaxAttemptsPerDayResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMaxAttemptsPerDayResponse) SetBody(v *GetMaxAttemptsPerDayResponseBody) *GetMaxAttemptsPerDayResponse {
	s.Body = v
	return s
}

func (s *GetMaxAttemptsPerDayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
