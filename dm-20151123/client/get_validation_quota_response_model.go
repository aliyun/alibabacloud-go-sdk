// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidationQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetValidationQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetValidationQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetValidationQuotaResponseBody) *GetValidationQuotaResponse
	GetBody() *GetValidationQuotaResponseBody
}

type GetValidationQuotaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetValidationQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetValidationQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetValidationQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetValidationQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetValidationQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetValidationQuotaResponse) GetBody() *GetValidationQuotaResponseBody {
	return s.Body
}

func (s *GetValidationQuotaResponse) SetHeaders(v map[string]*string) *GetValidationQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetValidationQuotaResponse) SetStatusCode(v int32) *GetValidationQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetValidationQuotaResponse) SetBody(v *GetValidationQuotaResponseBody) *GetValidationQuotaResponse {
	s.Body = v
	return s
}

func (s *GetValidationQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
