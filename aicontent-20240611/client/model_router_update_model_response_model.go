// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterUpdateModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterUpdateModelResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterUpdateModelResponseBody) *ModelRouterUpdateModelResponse
	GetBody() *ModelRouterUpdateModelResponseBody
}

type ModelRouterUpdateModelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterUpdateModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterUpdateModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateModelResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterUpdateModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterUpdateModelResponse) GetBody() *ModelRouterUpdateModelResponseBody {
	return s.Body
}

func (s *ModelRouterUpdateModelResponse) SetHeaders(v map[string]*string) *ModelRouterUpdateModelResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterUpdateModelResponse) SetStatusCode(v int32) *ModelRouterUpdateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterUpdateModelResponse) SetBody(v *ModelRouterUpdateModelResponseBody) *ModelRouterUpdateModelResponse {
	s.Body = v
	return s
}

func (s *ModelRouterUpdateModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
