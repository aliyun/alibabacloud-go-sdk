// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateApiKeyStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterUpdateApiKeyStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterUpdateApiKeyStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterUpdateApiKeyStatusResponseBody) *ModelRouterUpdateApiKeyStatusResponse
	GetBody() *ModelRouterUpdateApiKeyStatusResponseBody
}

type ModelRouterUpdateApiKeyStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterUpdateApiKeyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterUpdateApiKeyStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateApiKeyStatusResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateApiKeyStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterUpdateApiKeyStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterUpdateApiKeyStatusResponse) GetBody() *ModelRouterUpdateApiKeyStatusResponseBody {
	return s.Body
}

func (s *ModelRouterUpdateApiKeyStatusResponse) SetHeaders(v map[string]*string) *ModelRouterUpdateApiKeyStatusResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterUpdateApiKeyStatusResponse) SetStatusCode(v int32) *ModelRouterUpdateApiKeyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterUpdateApiKeyStatusResponse) SetBody(v *ModelRouterUpdateApiKeyStatusResponseBody) *ModelRouterUpdateApiKeyStatusResponse {
	s.Body = v
	return s
}

func (s *ModelRouterUpdateApiKeyStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
