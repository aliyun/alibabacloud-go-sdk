// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPurgeQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPurgeQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPurgeQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetPurgeQuotaResponseBody) *GetPurgeQuotaResponse
	GetBody() *GetPurgeQuotaResponseBody
}

type GetPurgeQuotaResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPurgeQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPurgeQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPurgeQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetPurgeQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPurgeQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPurgeQuotaResponse) GetBody() *GetPurgeQuotaResponseBody {
	return s.Body
}

func (s *GetPurgeQuotaResponse) SetHeaders(v map[string]*string) *GetPurgeQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetPurgeQuotaResponse) SetStatusCode(v int32) *GetPurgeQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPurgeQuotaResponse) SetBody(v *GetPurgeQuotaResponseBody) *GetPurgeQuotaResponse {
	s.Body = v
	return s
}

func (s *GetPurgeQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
