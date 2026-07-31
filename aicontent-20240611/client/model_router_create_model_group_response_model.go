// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateModelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCreateModelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCreateModelGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCreateModelGroupResponseBody) *ModelRouterCreateModelGroupResponse
	GetBody() *ModelRouterCreateModelGroupResponseBody
}

type ModelRouterCreateModelGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCreateModelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCreateModelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateModelGroupResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateModelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCreateModelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCreateModelGroupResponse) GetBody() *ModelRouterCreateModelGroupResponseBody {
	return s.Body
}

func (s *ModelRouterCreateModelGroupResponse) SetHeaders(v map[string]*string) *ModelRouterCreateModelGroupResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCreateModelGroupResponse) SetStatusCode(v int32) *ModelRouterCreateModelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCreateModelGroupResponse) SetBody(v *ModelRouterCreateModelGroupResponseBody) *ModelRouterCreateModelGroupResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCreateModelGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
