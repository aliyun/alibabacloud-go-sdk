// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterDeleteClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterDeleteClientResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterDeleteClientResponseBody) *ModelRouterDeleteClientResponse
	GetBody() *ModelRouterDeleteClientResponseBody
}

type ModelRouterDeleteClientResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterDeleteClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterDeleteClientResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteClientResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterDeleteClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterDeleteClientResponse) GetBody() *ModelRouterDeleteClientResponseBody {
	return s.Body
}

func (s *ModelRouterDeleteClientResponse) SetHeaders(v map[string]*string) *ModelRouterDeleteClientResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterDeleteClientResponse) SetStatusCode(v int32) *ModelRouterDeleteClientResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterDeleteClientResponse) SetBody(v *ModelRouterDeleteClientResponseBody) *ModelRouterDeleteClientResponse {
	s.Body = v
	return s
}

func (s *ModelRouterDeleteClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
