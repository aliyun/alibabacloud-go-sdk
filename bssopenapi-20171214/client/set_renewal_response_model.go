// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRenewalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetRenewalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetRenewalResponse
	GetStatusCode() *int32
	SetBody(v *SetRenewalResponseBody) *SetRenewalResponse
	GetBody() *SetRenewalResponseBody
}

type SetRenewalResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRenewalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRenewalResponse) String() string {
	return dara.Prettify(s)
}

func (s SetRenewalResponse) GoString() string {
	return s.String()
}

func (s *SetRenewalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetRenewalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetRenewalResponse) GetBody() *SetRenewalResponseBody {
	return s.Body
}

func (s *SetRenewalResponse) SetHeaders(v map[string]*string) *SetRenewalResponse {
	s.Headers = v
	return s
}

func (s *SetRenewalResponse) SetStatusCode(v int32) *SetRenewalResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRenewalResponse) SetBody(v *SetRenewalResponseBody) *SetRenewalResponse {
	s.Body = v
	return s
}

func (s *SetRenewalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
