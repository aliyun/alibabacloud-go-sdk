// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceQuotaResponseBody) *GetServiceQuotaResponse
	GetBody() *GetServiceQuotaResponseBody
}

type GetServiceQuotaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetServiceQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceQuotaResponse) GetBody() *GetServiceQuotaResponseBody {
	return s.Body
}

func (s *GetServiceQuotaResponse) SetHeaders(v map[string]*string) *GetServiceQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetServiceQuotaResponse) SetStatusCode(v int32) *GetServiceQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceQuotaResponse) SetBody(v *GetServiceQuotaResponseBody) *GetServiceQuotaResponse {
	s.Body = v
	return s
}

func (s *GetServiceQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
