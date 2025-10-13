// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebCustomDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWebCustomDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWebCustomDomainResponse
	GetStatusCode() *int32
	SetBody(v *WebCustomDomainBody) *CreateWebCustomDomainResponse
	GetBody() *WebCustomDomainBody
}

type CreateWebCustomDomainResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebCustomDomainBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebCustomDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWebCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateWebCustomDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWebCustomDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWebCustomDomainResponse) GetBody() *WebCustomDomainBody {
	return s.Body
}

func (s *CreateWebCustomDomainResponse) SetHeaders(v map[string]*string) *CreateWebCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateWebCustomDomainResponse) SetStatusCode(v int32) *CreateWebCustomDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebCustomDomainResponse) SetBody(v *WebCustomDomainBody) *CreateWebCustomDomainResponse {
	s.Body = v
	return s
}

func (s *CreateWebCustomDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
