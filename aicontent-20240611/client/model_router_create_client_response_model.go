// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCreateClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCreateClientResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCreateClientResponseBody) *ModelRouterCreateClientResponse
	GetBody() *ModelRouterCreateClientResponseBody
}

type ModelRouterCreateClientResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCreateClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCreateClientResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateClientResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCreateClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCreateClientResponse) GetBody() *ModelRouterCreateClientResponseBody {
	return s.Body
}

func (s *ModelRouterCreateClientResponse) SetHeaders(v map[string]*string) *ModelRouterCreateClientResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCreateClientResponse) SetStatusCode(v int32) *ModelRouterCreateClientResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCreateClientResponse) SetBody(v *ModelRouterCreateClientResponseBody) *ModelRouterCreateClientResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCreateClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
