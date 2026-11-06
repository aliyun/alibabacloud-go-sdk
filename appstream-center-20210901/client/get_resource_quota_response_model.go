// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceQuotaResponseBody) *GetResourceQuotaResponse
	GetBody() *GetResourceQuotaResponseBody
}

type GetResourceQuotaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetResourceQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceQuotaResponse) GetBody() *GetResourceQuotaResponseBody {
	return s.Body
}

func (s *GetResourceQuotaResponse) SetHeaders(v map[string]*string) *GetResourceQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetResourceQuotaResponse) SetStatusCode(v int32) *GetResourceQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceQuotaResponse) SetBody(v *GetResourceQuotaResponseBody) *GetResourceQuotaResponse {
	s.Body = v
	return s
}

func (s *GetResourceQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
