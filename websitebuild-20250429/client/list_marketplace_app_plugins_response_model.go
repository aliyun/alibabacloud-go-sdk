// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMarketplaceAppPluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMarketplaceAppPluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMarketplaceAppPluginsResponse
	GetStatusCode() *int32
	SetBody(v *ListMarketplaceAppPluginsResponseBody) *ListMarketplaceAppPluginsResponse
	GetBody() *ListMarketplaceAppPluginsResponseBody
}

type ListMarketplaceAppPluginsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMarketplaceAppPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMarketplaceAppPluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMarketplaceAppPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListMarketplaceAppPluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMarketplaceAppPluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMarketplaceAppPluginsResponse) GetBody() *ListMarketplaceAppPluginsResponseBody {
	return s.Body
}

func (s *ListMarketplaceAppPluginsResponse) SetHeaders(v map[string]*string) *ListMarketplaceAppPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListMarketplaceAppPluginsResponse) SetStatusCode(v int32) *ListMarketplaceAppPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponse) SetBody(v *ListMarketplaceAppPluginsResponseBody) *ListMarketplaceAppPluginsResponse {
	s.Body = v
	return s
}

func (s *ListMarketplaceAppPluginsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
