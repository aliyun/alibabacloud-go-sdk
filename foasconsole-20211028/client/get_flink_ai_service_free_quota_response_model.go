// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlinkAiServiceFreeQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFlinkAiServiceFreeQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFlinkAiServiceFreeQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetFlinkAiServiceFreeQuotaResponseBody) *GetFlinkAiServiceFreeQuotaResponse
	GetBody() *GetFlinkAiServiceFreeQuotaResponseBody
}

type GetFlinkAiServiceFreeQuotaResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFlinkAiServiceFreeQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlinkAiServiceFreeQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFlinkAiServiceFreeQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetFlinkAiServiceFreeQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFlinkAiServiceFreeQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFlinkAiServiceFreeQuotaResponse) GetBody() *GetFlinkAiServiceFreeQuotaResponseBody {
	return s.Body
}

func (s *GetFlinkAiServiceFreeQuotaResponse) SetHeaders(v map[string]*string) *GetFlinkAiServiceFreeQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetFlinkAiServiceFreeQuotaResponse) SetStatusCode(v int32) *GetFlinkAiServiceFreeQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlinkAiServiceFreeQuotaResponse) SetBody(v *GetFlinkAiServiceFreeQuotaResponseBody) *GetFlinkAiServiceFreeQuotaResponse {
	s.Body = v
	return s
}

func (s *GetFlinkAiServiceFreeQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
