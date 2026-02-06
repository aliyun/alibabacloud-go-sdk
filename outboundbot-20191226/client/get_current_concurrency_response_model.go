// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentConcurrencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCurrentConcurrencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCurrentConcurrencyResponse
	GetStatusCode() *int32
	SetBody(v *GetCurrentConcurrencyResponseBody) *GetCurrentConcurrencyResponse
	GetBody() *GetCurrentConcurrencyResponseBody
}

type GetCurrentConcurrencyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCurrentConcurrencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCurrentConcurrencyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentConcurrencyResponse) GoString() string {
	return s.String()
}

func (s *GetCurrentConcurrencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCurrentConcurrencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCurrentConcurrencyResponse) GetBody() *GetCurrentConcurrencyResponseBody {
	return s.Body
}

func (s *GetCurrentConcurrencyResponse) SetHeaders(v map[string]*string) *GetCurrentConcurrencyResponse {
	s.Headers = v
	return s
}

func (s *GetCurrentConcurrencyResponse) SetStatusCode(v int32) *GetCurrentConcurrencyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCurrentConcurrencyResponse) SetBody(v *GetCurrentConcurrencyResponseBody) *GetCurrentConcurrencyResponse {
	s.Body = v
	return s
}

func (s *GetCurrentConcurrencyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
