// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryModelGroupModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryModelGroupModelsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryModelGroupModelsResponseBody) *ModelRouterQueryModelGroupModelsResponse
	GetBody() *ModelRouterQueryModelGroupModelsResponseBody
}

type ModelRouterQueryModelGroupModelsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryModelGroupModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryModelGroupModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupModelsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryModelGroupModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryModelGroupModelsResponse) GetBody() *ModelRouterQueryModelGroupModelsResponseBody {
	return s.Body
}

func (s *ModelRouterQueryModelGroupModelsResponse) SetHeaders(v map[string]*string) *ModelRouterQueryModelGroupModelsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponse) SetStatusCode(v int32) *ModelRouterQueryModelGroupModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponse) SetBody(v *ModelRouterQueryModelGroupModelsResponseBody) *ModelRouterQueryModelGroupModelsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
