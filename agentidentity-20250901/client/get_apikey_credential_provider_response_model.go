// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAPIKeyCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAPIKeyCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAPIKeyCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetAPIKeyCredentialProviderResponseBody) *GetAPIKeyCredentialProviderResponse
	GetBody() *GetAPIKeyCredentialProviderResponseBody
}

type GetAPIKeyCredentialProviderResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAPIKeyCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAPIKeyCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAPIKeyCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *GetAPIKeyCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAPIKeyCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAPIKeyCredentialProviderResponse) GetBody() *GetAPIKeyCredentialProviderResponseBody {
	return s.Body
}

func (s *GetAPIKeyCredentialProviderResponse) SetHeaders(v map[string]*string) *GetAPIKeyCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *GetAPIKeyCredentialProviderResponse) SetStatusCode(v int32) *GetAPIKeyCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAPIKeyCredentialProviderResponse) SetBody(v *GetAPIKeyCredentialProviderResponseBody) *GetAPIKeyCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *GetAPIKeyCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
