// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainNameCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainNameCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveTaskForSubmittingDomainNameCredentialResponse
	GetStatusCode() *int32
	SetBody(v *SaveTaskForSubmittingDomainNameCredentialResponseBody) *SaveTaskForSubmittingDomainNameCredentialResponse
	GetBody() *SaveTaskForSubmittingDomainNameCredentialResponseBody
}

type SaveTaskForSubmittingDomainNameCredentialResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTaskForSubmittingDomainNameCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTaskForSubmittingDomainNameCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainNameCredentialResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainNameCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveTaskForSubmittingDomainNameCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveTaskForSubmittingDomainNameCredentialResponse) GetBody() *SaveTaskForSubmittingDomainNameCredentialResponseBody {
	return s.Body
}

func (s *SaveTaskForSubmittingDomainNameCredentialResponse) SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainNameCredentialResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialResponse) SetStatusCode(v int32) *SaveTaskForSubmittingDomainNameCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialResponse) SetBody(v *SaveTaskForSubmittingDomainNameCredentialResponseBody) *SaveTaskForSubmittingDomainNameCredentialResponse {
	s.Body = v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
