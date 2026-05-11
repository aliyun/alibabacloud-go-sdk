// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealTimeConcurrencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRealTimeConcurrencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRealTimeConcurrencyResponse
	GetStatusCode() *int32
	SetBody(v *GetRealTimeConcurrencyResponseBody) *GetRealTimeConcurrencyResponse
	GetBody() *GetRealTimeConcurrencyResponseBody
}

type GetRealTimeConcurrencyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealTimeConcurrencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealTimeConcurrencyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRealTimeConcurrencyResponse) GoString() string {
	return s.String()
}

func (s *GetRealTimeConcurrencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRealTimeConcurrencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRealTimeConcurrencyResponse) GetBody() *GetRealTimeConcurrencyResponseBody {
	return s.Body
}

func (s *GetRealTimeConcurrencyResponse) SetHeaders(v map[string]*string) *GetRealTimeConcurrencyResponse {
	s.Headers = v
	return s
}

func (s *GetRealTimeConcurrencyResponse) SetStatusCode(v int32) *GetRealTimeConcurrencyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealTimeConcurrencyResponse) SetBody(v *GetRealTimeConcurrencyResponseBody) *GetRealTimeConcurrencyResponse {
	s.Body = v
	return s
}

func (s *GetRealTimeConcurrencyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
