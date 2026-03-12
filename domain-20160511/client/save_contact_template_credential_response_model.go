// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactTemplateCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveContactTemplateCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveContactTemplateCredentialResponse
	GetStatusCode() *int32
	SetBody(v *SaveContactTemplateCredentialResponseBody) *SaveContactTemplateCredentialResponse
	GetBody() *SaveContactTemplateCredentialResponseBody
}

type SaveContactTemplateCredentialResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveContactTemplateCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveContactTemplateCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveContactTemplateCredentialResponse) GoString() string {
	return s.String()
}

func (s *SaveContactTemplateCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveContactTemplateCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveContactTemplateCredentialResponse) GetBody() *SaveContactTemplateCredentialResponseBody {
	return s.Body
}

func (s *SaveContactTemplateCredentialResponse) SetHeaders(v map[string]*string) *SaveContactTemplateCredentialResponse {
	s.Headers = v
	return s
}

func (s *SaveContactTemplateCredentialResponse) SetStatusCode(v int32) *SaveContactTemplateCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveContactTemplateCredentialResponse) SetBody(v *SaveContactTemplateCredentialResponseBody) *SaveContactTemplateCredentialResponse {
	s.Body = v
	return s
}

func (s *SaveContactTemplateCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
