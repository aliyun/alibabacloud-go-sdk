// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePipelineGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePipelineGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePipelineGroupResponseBody) *UpdatePipelineGroupResponse
	GetBody() *UpdatePipelineGroupResponseBody
}

type UpdatePipelineGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePipelineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePipelineGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePipelineGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePipelineGroupResponse) GetBody() *UpdatePipelineGroupResponseBody {
	return s.Body
}

func (s *UpdatePipelineGroupResponse) SetHeaders(v map[string]*string) *UpdatePipelineGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineGroupResponse) SetStatusCode(v int32) *UpdatePipelineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelineGroupResponse) SetBody(v *UpdatePipelineGroupResponseBody) *UpdatePipelineGroupResponse {
	s.Body = v
	return s
}

func (s *UpdatePipelineGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
