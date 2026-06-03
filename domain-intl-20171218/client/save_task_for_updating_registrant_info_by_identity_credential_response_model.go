// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse
	GetStatusCode() *int32
	SetBody(v *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse
	GetBody() *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody
}

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse struct {
	Headers    map[string]*string                                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) GetBody() *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody {
	return s.Body
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) SetHeaders(v map[string]*string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) SetStatusCode(v int32) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) SetBody(v *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse {
	s.Body = v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
