// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse
	GetStatusCode() *int32
	SetBody(v *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse
	GetBody() *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody
}

type SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse struct {
	Headers    map[string]*string                                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse) GetBody() *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody {
	return s.Body
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse) SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse) SetStatusCode(v int32) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse) SetBody(v *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponseBody) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse {
	s.Body = v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
