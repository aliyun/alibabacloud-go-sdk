// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderSyncJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIdentityProviderSyncJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIdentityProviderSyncJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateIdentityProviderSyncJobResponseBody) *CreateIdentityProviderSyncJobResponse
	GetBody() *CreateIdentityProviderSyncJobResponseBody
}

type CreateIdentityProviderSyncJobResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIdentityProviderSyncJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIdentityProviderSyncJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderSyncJobResponse) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderSyncJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIdentityProviderSyncJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIdentityProviderSyncJobResponse) GetBody() *CreateIdentityProviderSyncJobResponseBody {
	return s.Body
}

func (s *CreateIdentityProviderSyncJobResponse) SetHeaders(v map[string]*string) *CreateIdentityProviderSyncJobResponse {
	s.Headers = v
	return s
}

func (s *CreateIdentityProviderSyncJobResponse) SetStatusCode(v int32) *CreateIdentityProviderSyncJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIdentityProviderSyncJobResponse) SetBody(v *CreateIdentityProviderSyncJobResponseBody) *CreateIdentityProviderSyncJobResponse {
	s.Body = v
	return s
}

func (s *CreateIdentityProviderSyncJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
