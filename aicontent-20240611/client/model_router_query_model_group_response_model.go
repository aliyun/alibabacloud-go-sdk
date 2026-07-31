// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryModelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryModelGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryModelGroupResponseBody) *ModelRouterQueryModelGroupResponse
	GetBody() *ModelRouterQueryModelGroupResponseBody
}

type ModelRouterQueryModelGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryModelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryModelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryModelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryModelGroupResponse) GetBody() *ModelRouterQueryModelGroupResponseBody {
	return s.Body
}

func (s *ModelRouterQueryModelGroupResponse) SetHeaders(v map[string]*string) *ModelRouterQueryModelGroupResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryModelGroupResponse) SetStatusCode(v int32) *ModelRouterQueryModelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupResponse) SetBody(v *ModelRouterQueryModelGroupResponseBody) *ModelRouterQueryModelGroupResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryModelGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
