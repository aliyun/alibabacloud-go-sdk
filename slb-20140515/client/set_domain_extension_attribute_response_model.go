// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainExtensionAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDomainExtensionAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDomainExtensionAttributeResponse
	GetStatusCode() *int32
	SetBody(v *SetDomainExtensionAttributeResponseBody) *SetDomainExtensionAttributeResponse
	GetBody() *SetDomainExtensionAttributeResponseBody
}

type SetDomainExtensionAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDomainExtensionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDomainExtensionAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDomainExtensionAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetDomainExtensionAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDomainExtensionAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDomainExtensionAttributeResponse) GetBody() *SetDomainExtensionAttributeResponseBody {
	return s.Body
}

func (s *SetDomainExtensionAttributeResponse) SetHeaders(v map[string]*string) *SetDomainExtensionAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetDomainExtensionAttributeResponse) SetStatusCode(v int32) *SetDomainExtensionAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDomainExtensionAttributeResponse) SetBody(v *SetDomainExtensionAttributeResponseBody) *SetDomainExtensionAttributeResponse {
	s.Body = v
	return s
}

func (s *SetDomainExtensionAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
