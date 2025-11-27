// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegrationVersionForCSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIntegrationVersionForCSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIntegrationVersionForCSResponse
	GetStatusCode() *int32
	SetBody(v *GetIntegrationVersionForCSResponseBody) *GetIntegrationVersionForCSResponse
	GetBody() *GetIntegrationVersionForCSResponseBody
}

type GetIntegrationVersionForCSResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIntegrationVersionForCSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIntegrationVersionForCSResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationVersionForCSResponse) GoString() string {
	return s.String()
}

func (s *GetIntegrationVersionForCSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIntegrationVersionForCSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIntegrationVersionForCSResponse) GetBody() *GetIntegrationVersionForCSResponseBody {
	return s.Body
}

func (s *GetIntegrationVersionForCSResponse) SetHeaders(v map[string]*string) *GetIntegrationVersionForCSResponse {
	s.Headers = v
	return s
}

func (s *GetIntegrationVersionForCSResponse) SetStatusCode(v int32) *GetIntegrationVersionForCSResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntegrationVersionForCSResponse) SetBody(v *GetIntegrationVersionForCSResponseBody) *GetIntegrationVersionForCSResponse {
	s.Body = v
	return s
}

func (s *GetIntegrationVersionForCSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
