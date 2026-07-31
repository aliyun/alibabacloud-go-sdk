// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchBindModelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterBatchBindModelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterBatchBindModelGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterBatchBindModelGroupResponseBody) *ModelRouterBatchBindModelGroupResponse
	GetBody() *ModelRouterBatchBindModelGroupResponseBody
}

type ModelRouterBatchBindModelGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterBatchBindModelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterBatchBindModelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchBindModelGroupResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchBindModelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterBatchBindModelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterBatchBindModelGroupResponse) GetBody() *ModelRouterBatchBindModelGroupResponseBody {
	return s.Body
}

func (s *ModelRouterBatchBindModelGroupResponse) SetHeaders(v map[string]*string) *ModelRouterBatchBindModelGroupResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterBatchBindModelGroupResponse) SetStatusCode(v int32) *ModelRouterBatchBindModelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterBatchBindModelGroupResponse) SetBody(v *ModelRouterBatchBindModelGroupResponseBody) *ModelRouterBatchBindModelGroupResponse {
	s.Body = v
	return s
}

func (s *ModelRouterBatchBindModelGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
