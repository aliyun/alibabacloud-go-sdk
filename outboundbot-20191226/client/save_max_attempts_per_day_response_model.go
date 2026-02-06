// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMaxAttemptsPerDayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveMaxAttemptsPerDayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveMaxAttemptsPerDayResponse
	GetStatusCode() *int32
	SetBody(v *SaveMaxAttemptsPerDayResponseBody) *SaveMaxAttemptsPerDayResponse
	GetBody() *SaveMaxAttemptsPerDayResponseBody
}

type SaveMaxAttemptsPerDayResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveMaxAttemptsPerDayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveMaxAttemptsPerDayResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveMaxAttemptsPerDayResponse) GoString() string {
	return s.String()
}

func (s *SaveMaxAttemptsPerDayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveMaxAttemptsPerDayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveMaxAttemptsPerDayResponse) GetBody() *SaveMaxAttemptsPerDayResponseBody {
	return s.Body
}

func (s *SaveMaxAttemptsPerDayResponse) SetHeaders(v map[string]*string) *SaveMaxAttemptsPerDayResponse {
	s.Headers = v
	return s
}

func (s *SaveMaxAttemptsPerDayResponse) SetStatusCode(v int32) *SaveMaxAttemptsPerDayResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveMaxAttemptsPerDayResponse) SetBody(v *SaveMaxAttemptsPerDayResponseBody) *SaveMaxAttemptsPerDayResponse {
	s.Body = v
	return s
}

func (s *SaveMaxAttemptsPerDayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
