// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterDeleteModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterDeleteModelResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterDeleteModelResponseBody) *ModelRouterDeleteModelResponse
	GetBody() *ModelRouterDeleteModelResponseBody
}

type ModelRouterDeleteModelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterDeleteModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterDeleteModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteModelResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterDeleteModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterDeleteModelResponse) GetBody() *ModelRouterDeleteModelResponseBody {
	return s.Body
}

func (s *ModelRouterDeleteModelResponse) SetHeaders(v map[string]*string) *ModelRouterDeleteModelResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterDeleteModelResponse) SetStatusCode(v int32) *ModelRouterDeleteModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterDeleteModelResponse) SetBody(v *ModelRouterDeleteModelResponseBody) *ModelRouterDeleteModelResponse {
	s.Body = v
	return s
}

func (s *ModelRouterDeleteModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
