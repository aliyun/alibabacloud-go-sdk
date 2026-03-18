// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryModelResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryModelResponseBody) *ModelRouterQueryModelResponse
	GetBody() *ModelRouterQueryModelResponseBody
}

type ModelRouterQueryModelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryModelResponse) GetBody() *ModelRouterQueryModelResponseBody {
	return s.Body
}

func (s *ModelRouterQueryModelResponse) SetHeaders(v map[string]*string) *ModelRouterQueryModelResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryModelResponse) SetStatusCode(v int32) *ModelRouterQueryModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryModelResponse) SetBody(v *ModelRouterQueryModelResponseBody) *ModelRouterQueryModelResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
