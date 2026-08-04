// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateMemberApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCreateMemberApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCreateMemberApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCreateMemberApiKeyResponseBody) *ModelRouterCreateMemberApiKeyResponse
	GetBody() *ModelRouterCreateMemberApiKeyResponseBody
}

type ModelRouterCreateMemberApiKeyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCreateMemberApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCreateMemberApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateMemberApiKeyResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateMemberApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCreateMemberApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCreateMemberApiKeyResponse) GetBody() *ModelRouterCreateMemberApiKeyResponseBody {
	return s.Body
}

func (s *ModelRouterCreateMemberApiKeyResponse) SetHeaders(v map[string]*string) *ModelRouterCreateMemberApiKeyResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCreateMemberApiKeyResponse) SetStatusCode(v int32) *ModelRouterCreateMemberApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCreateMemberApiKeyResponse) SetBody(v *ModelRouterCreateMemberApiKeyResponseBody) *ModelRouterCreateMemberApiKeyResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCreateMemberApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
