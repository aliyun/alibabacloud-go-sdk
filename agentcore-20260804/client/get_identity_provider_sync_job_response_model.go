// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderSyncJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentityProviderSyncJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentityProviderSyncJobResponse
	GetStatusCode() *int32
	SetBody(v *GetIdentityProviderSyncJobResponseBody) *GetIdentityProviderSyncJobResponse
	GetBody() *GetIdentityProviderSyncJobResponseBody
}

type GetIdentityProviderSyncJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdentityProviderSyncJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentityProviderSyncJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncJobResponse) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentityProviderSyncJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentityProviderSyncJobResponse) GetBody() *GetIdentityProviderSyncJobResponseBody {
	return s.Body
}

func (s *GetIdentityProviderSyncJobResponse) SetHeaders(v map[string]*string) *GetIdentityProviderSyncJobResponse {
	s.Headers = v
	return s
}

func (s *GetIdentityProviderSyncJobResponse) SetStatusCode(v int32) *GetIdentityProviderSyncJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentityProviderSyncJobResponse) SetBody(v *GetIdentityProviderSyncJobResponseBody) *GetIdentityProviderSyncJobResponse {
	s.Body = v
	return s
}

func (s *GetIdentityProviderSyncJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
