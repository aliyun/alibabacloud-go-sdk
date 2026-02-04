// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceQuotaResponseBody) *GetInstanceQuotaResponse
	GetBody() *GetInstanceQuotaResponseBody
}

type GetInstanceQuotaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceQuotaResponse) GetBody() *GetInstanceQuotaResponseBody {
	return s.Body
}

func (s *GetInstanceQuotaResponse) SetHeaders(v map[string]*string) *GetInstanceQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceQuotaResponse) SetStatusCode(v int32) *GetInstanceQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceQuotaResponse) SetBody(v *GetInstanceQuotaResponseBody) *GetInstanceQuotaResponse {
	s.Body = v
	return s
}

func (s *GetInstanceQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
