// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebCustomDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWebCustomDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWebCustomDomainResponse
	GetStatusCode() *int32
	SetBody(v *WebCustomDomainBody) *UpdateWebCustomDomainResponse
	GetBody() *WebCustomDomainBody
}

type UpdateWebCustomDomainResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebCustomDomainBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebCustomDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebCustomDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWebCustomDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWebCustomDomainResponse) GetBody() *WebCustomDomainBody {
	return s.Body
}

func (s *UpdateWebCustomDomainResponse) SetHeaders(v map[string]*string) *UpdateWebCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebCustomDomainResponse) SetStatusCode(v int32) *UpdateWebCustomDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebCustomDomainResponse) SetBody(v *WebCustomDomainBody) *UpdateWebCustomDomainResponse {
	s.Body = v
	return s
}

func (s *UpdateWebCustomDomainResponse) Validate() error {
	return dara.Validate(s)
}
