// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteModelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterDeleteModelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterDeleteModelGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterDeleteModelGroupResponseBody) *ModelRouterDeleteModelGroupResponse
	GetBody() *ModelRouterDeleteModelGroupResponseBody
}

type ModelRouterDeleteModelGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterDeleteModelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterDeleteModelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteModelGroupResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteModelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterDeleteModelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterDeleteModelGroupResponse) GetBody() *ModelRouterDeleteModelGroupResponseBody {
	return s.Body
}

func (s *ModelRouterDeleteModelGroupResponse) SetHeaders(v map[string]*string) *ModelRouterDeleteModelGroupResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterDeleteModelGroupResponse) SetStatusCode(v int32) *ModelRouterDeleteModelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterDeleteModelGroupResponse) SetBody(v *ModelRouterDeleteModelGroupResponseBody) *ModelRouterDeleteModelGroupResponse {
	s.Body = v
	return s
}

func (s *ModelRouterDeleteModelGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
