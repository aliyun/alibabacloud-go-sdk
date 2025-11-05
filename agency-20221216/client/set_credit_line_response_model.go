// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCreditLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCreditLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCreditLineResponse
	GetStatusCode() *int32
	SetBody(v *SetCreditLineResponseBody) *SetCreditLineResponse
	GetBody() *SetCreditLineResponseBody
}

type SetCreditLineResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCreditLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCreditLineResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCreditLineResponse) GoString() string {
	return s.String()
}

func (s *SetCreditLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCreditLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCreditLineResponse) GetBody() *SetCreditLineResponseBody {
	return s.Body
}

func (s *SetCreditLineResponse) SetHeaders(v map[string]*string) *SetCreditLineResponse {
	s.Headers = v
	return s
}

func (s *SetCreditLineResponse) SetStatusCode(v int32) *SetCreditLineResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCreditLineResponse) SetBody(v *SetCreditLineResponseBody) *SetCreditLineResponse {
	s.Body = v
	return s
}

func (s *SetCreditLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
