// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebCustomDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWebCustomDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWebCustomDomainResponse
	GetStatusCode() *int32
	SetBody(v *WebCustomDomainBody) *DeleteWebCustomDomainResponse
	GetBody() *WebCustomDomainBody
}

type DeleteWebCustomDomainResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebCustomDomainBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWebCustomDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebCustomDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWebCustomDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWebCustomDomainResponse) GetBody() *WebCustomDomainBody {
	return s.Body
}

func (s *DeleteWebCustomDomainResponse) SetHeaders(v map[string]*string) *DeleteWebCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebCustomDomainResponse) SetStatusCode(v int32) *DeleteWebCustomDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebCustomDomainResponse) SetBody(v *WebCustomDomainBody) *DeleteWebCustomDomainResponse {
	s.Body = v
	return s
}

func (s *DeleteWebCustomDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
