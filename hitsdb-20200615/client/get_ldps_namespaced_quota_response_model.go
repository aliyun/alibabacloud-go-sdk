// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLdpsNamespacedQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLdpsNamespacedQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLdpsNamespacedQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetLdpsNamespacedQuotaResponseBody) *GetLdpsNamespacedQuotaResponse
	GetBody() *GetLdpsNamespacedQuotaResponseBody
}

type GetLdpsNamespacedQuotaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLdpsNamespacedQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLdpsNamespacedQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLdpsNamespacedQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetLdpsNamespacedQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLdpsNamespacedQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLdpsNamespacedQuotaResponse) GetBody() *GetLdpsNamespacedQuotaResponseBody {
	return s.Body
}

func (s *GetLdpsNamespacedQuotaResponse) SetHeaders(v map[string]*string) *GetLdpsNamespacedQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetLdpsNamespacedQuotaResponse) SetStatusCode(v int32) *GetLdpsNamespacedQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLdpsNamespacedQuotaResponse) SetBody(v *GetLdpsNamespacedQuotaResponseBody) *GetLdpsNamespacedQuotaResponse {
	s.Body = v
	return s
}

func (s *GetLdpsNamespacedQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
