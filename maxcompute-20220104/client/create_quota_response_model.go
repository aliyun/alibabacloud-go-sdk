// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQuotaResponse
	GetStatusCode() *int32
	SetBody(v *CreateQuotaResponseBody) *CreateQuotaResponse
	GetBody() *CreateQuotaResponseBody
}

type CreateQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQuotaResponse) GetBody() *CreateQuotaResponseBody {
	return s.Body
}

func (s *CreateQuotaResponse) SetHeaders(v map[string]*string) *CreateQuotaResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaResponse) SetStatusCode(v int32) *CreateQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQuotaResponse) SetBody(v *CreateQuotaResponseBody) *CreateQuotaResponse {
	s.Body = v
	return s
}

func (s *CreateQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
