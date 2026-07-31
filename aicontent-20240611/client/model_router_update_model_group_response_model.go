// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateModelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterUpdateModelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterUpdateModelGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterUpdateModelGroupResponseBody) *ModelRouterUpdateModelGroupResponse
	GetBody() *ModelRouterUpdateModelGroupResponseBody
}

type ModelRouterUpdateModelGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterUpdateModelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterUpdateModelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateModelGroupResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateModelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterUpdateModelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterUpdateModelGroupResponse) GetBody() *ModelRouterUpdateModelGroupResponseBody {
	return s.Body
}

func (s *ModelRouterUpdateModelGroupResponse) SetHeaders(v map[string]*string) *ModelRouterUpdateModelGroupResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterUpdateModelGroupResponse) SetStatusCode(v int32) *ModelRouterUpdateModelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterUpdateModelGroupResponse) SetBody(v *ModelRouterUpdateModelGroupResponseBody) *ModelRouterUpdateModelGroupResponse {
	s.Body = v
	return s
}

func (s *ModelRouterUpdateModelGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
