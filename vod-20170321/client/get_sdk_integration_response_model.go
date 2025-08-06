// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSdkIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSdkIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSdkIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *GetSdkIntegrationResponseBody) *GetSdkIntegrationResponse
	GetBody() *GetSdkIntegrationResponseBody
}

type GetSdkIntegrationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSdkIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSdkIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSdkIntegrationResponse) GoString() string {
	return s.String()
}

func (s *GetSdkIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSdkIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSdkIntegrationResponse) GetBody() *GetSdkIntegrationResponseBody {
	return s.Body
}

func (s *GetSdkIntegrationResponse) SetHeaders(v map[string]*string) *GetSdkIntegrationResponse {
	s.Headers = v
	return s
}

func (s *GetSdkIntegrationResponse) SetStatusCode(v int32) *GetSdkIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSdkIntegrationResponse) SetBody(v *GetSdkIntegrationResponseBody) *GetSdkIntegrationResponse {
	s.Body = v
	return s
}

func (s *GetSdkIntegrationResponse) Validate() error {
	return dara.Validate(s)
}
