// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDispatchConsoleAPIForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DispatchConsoleAPIForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DispatchConsoleAPIForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *DispatchConsoleAPIForPartnerResponseBody) *DispatchConsoleAPIForPartnerResponse
	GetBody() *DispatchConsoleAPIForPartnerResponseBody
}

type DispatchConsoleAPIForPartnerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DispatchConsoleAPIForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DispatchConsoleAPIForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s DispatchConsoleAPIForPartnerResponse) GoString() string {
	return s.String()
}

func (s *DispatchConsoleAPIForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DispatchConsoleAPIForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DispatchConsoleAPIForPartnerResponse) GetBody() *DispatchConsoleAPIForPartnerResponseBody {
	return s.Body
}

func (s *DispatchConsoleAPIForPartnerResponse) SetHeaders(v map[string]*string) *DispatchConsoleAPIForPartnerResponse {
	s.Headers = v
	return s
}

func (s *DispatchConsoleAPIForPartnerResponse) SetStatusCode(v int32) *DispatchConsoleAPIForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *DispatchConsoleAPIForPartnerResponse) SetBody(v *DispatchConsoleAPIForPartnerResponseBody) *DispatchConsoleAPIForPartnerResponse {
	s.Body = v
	return s
}

func (s *DispatchConsoleAPIForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
