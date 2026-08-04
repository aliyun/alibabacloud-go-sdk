// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetMemberApiKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterGetMemberApiKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterGetMemberApiKeysResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterGetMemberApiKeysResponseBody) *ModelRouterGetMemberApiKeysResponse
	GetBody() *ModelRouterGetMemberApiKeysResponseBody
}

type ModelRouterGetMemberApiKeysResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterGetMemberApiKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterGetMemberApiKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetMemberApiKeysResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterGetMemberApiKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterGetMemberApiKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterGetMemberApiKeysResponse) GetBody() *ModelRouterGetMemberApiKeysResponseBody {
	return s.Body
}

func (s *ModelRouterGetMemberApiKeysResponse) SetHeaders(v map[string]*string) *ModelRouterGetMemberApiKeysResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponse) SetStatusCode(v int32) *ModelRouterGetMemberApiKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponse) SetBody(v *ModelRouterGetMemberApiKeysResponseBody) *ModelRouterGetMemberApiKeysResponse {
	s.Body = v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
