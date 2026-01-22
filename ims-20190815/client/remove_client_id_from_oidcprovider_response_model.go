// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClientIdFromOIDCProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveClientIdFromOIDCProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveClientIdFromOIDCProviderResponse
	GetStatusCode() *int32
	SetBody(v *RemoveClientIdFromOIDCProviderResponseBody) *RemoveClientIdFromOIDCProviderResponse
	GetBody() *RemoveClientIdFromOIDCProviderResponseBody
}

type RemoveClientIdFromOIDCProviderResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveClientIdFromOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveClientIdFromOIDCProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveClientIdFromOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *RemoveClientIdFromOIDCProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveClientIdFromOIDCProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveClientIdFromOIDCProviderResponse) GetBody() *RemoveClientIdFromOIDCProviderResponseBody {
	return s.Body
}

func (s *RemoveClientIdFromOIDCProviderResponse) SetHeaders(v map[string]*string) *RemoveClientIdFromOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponse) SetStatusCode(v int32) *RemoveClientIdFromOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponse) SetBody(v *RemoveClientIdFromOIDCProviderResponseBody) *RemoveClientIdFromOIDCProviderResponse {
	s.Body = v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
