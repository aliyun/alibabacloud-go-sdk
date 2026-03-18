// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterUpdateClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterUpdateClientResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterUpdateClientResponseBody) *ModelRouterUpdateClientResponse
	GetBody() *ModelRouterUpdateClientResponseBody
}

type ModelRouterUpdateClientResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterUpdateClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterUpdateClientResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateClientResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterUpdateClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterUpdateClientResponse) GetBody() *ModelRouterUpdateClientResponseBody {
	return s.Body
}

func (s *ModelRouterUpdateClientResponse) SetHeaders(v map[string]*string) *ModelRouterUpdateClientResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterUpdateClientResponse) SetStatusCode(v int32) *ModelRouterUpdateClientResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterUpdateClientResponse) SetBody(v *ModelRouterUpdateClientResponseBody) *ModelRouterUpdateClientResponse {
	s.Body = v
	return s
}

func (s *ModelRouterUpdateClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
