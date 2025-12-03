// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelineGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePipelineGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePipelineGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeletePipelineGroupResponseBody) *DeletePipelineGroupResponse
	GetBody() *DeletePipelineGroupResponseBody
}

type DeletePipelineGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePipelineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePipelineGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelineGroupResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelineGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePipelineGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePipelineGroupResponse) GetBody() *DeletePipelineGroupResponseBody {
	return s.Body
}

func (s *DeletePipelineGroupResponse) SetHeaders(v map[string]*string) *DeletePipelineGroupResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelineGroupResponse) SetStatusCode(v int32) *DeletePipelineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePipelineGroupResponse) SetBody(v *DeletePipelineGroupResponseBody) *DeletePipelineGroupResponse {
	s.Body = v
	return s
}

func (s *DeletePipelineGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
