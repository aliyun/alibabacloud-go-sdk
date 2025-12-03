// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAppsAuthToApiProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAppsAuthToApiProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAppsAuthToApiProductResponse
	GetStatusCode() *int32
	SetBody(v *SetAppsAuthToApiProductResponseBody) *SetAppsAuthToApiProductResponse
	GetBody() *SetAppsAuthToApiProductResponseBody
}

type SetAppsAuthToApiProductResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAppsAuthToApiProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAppsAuthToApiProductResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAppsAuthToApiProductResponse) GoString() string {
	return s.String()
}

func (s *SetAppsAuthToApiProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAppsAuthToApiProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAppsAuthToApiProductResponse) GetBody() *SetAppsAuthToApiProductResponseBody {
	return s.Body
}

func (s *SetAppsAuthToApiProductResponse) SetHeaders(v map[string]*string) *SetAppsAuthToApiProductResponse {
	s.Headers = v
	return s
}

func (s *SetAppsAuthToApiProductResponse) SetStatusCode(v int32) *SetAppsAuthToApiProductResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAppsAuthToApiProductResponse) SetBody(v *SetAppsAuthToApiProductResponseBody) *SetAppsAuthToApiProductResponse {
	s.Body = v
	return s
}

func (s *SetAppsAuthToApiProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
