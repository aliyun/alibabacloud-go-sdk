// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCreateModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCreateModelResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCreateModelResponseBody) *ModelRouterCreateModelResponse
	GetBody() *ModelRouterCreateModelResponseBody
}

type ModelRouterCreateModelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCreateModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCreateModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateModelResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCreateModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCreateModelResponse) GetBody() *ModelRouterCreateModelResponseBody {
	return s.Body
}

func (s *ModelRouterCreateModelResponse) SetHeaders(v map[string]*string) *ModelRouterCreateModelResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCreateModelResponse) SetStatusCode(v int32) *ModelRouterCreateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCreateModelResponse) SetBody(v *ModelRouterCreateModelResponseBody) *ModelRouterCreateModelResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCreateModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
