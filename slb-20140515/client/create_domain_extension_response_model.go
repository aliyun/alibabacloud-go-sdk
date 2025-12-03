// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainExtensionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDomainExtensionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDomainExtensionResponse
	GetStatusCode() *int32
	SetBody(v *CreateDomainExtensionResponseBody) *CreateDomainExtensionResponse
	GetBody() *CreateDomainExtensionResponseBody
}

type CreateDomainExtensionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDomainExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDomainExtensionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainExtensionResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainExtensionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDomainExtensionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDomainExtensionResponse) GetBody() *CreateDomainExtensionResponseBody {
	return s.Body
}

func (s *CreateDomainExtensionResponse) SetHeaders(v map[string]*string) *CreateDomainExtensionResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainExtensionResponse) SetStatusCode(v int32) *CreateDomainExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainExtensionResponse) SetBody(v *CreateDomainExtensionResponseBody) *CreateDomainExtensionResponse {
	s.Body = v
	return s
}

func (s *CreateDomainExtensionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
