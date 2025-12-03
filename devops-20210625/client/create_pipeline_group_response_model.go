// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePipelineGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePipelineGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreatePipelineGroupResponseBody) *CreatePipelineGroupResponse
	GetBody() *CreatePipelineGroupResponseBody
}

type CreatePipelineGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePipelineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineGroupResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePipelineGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePipelineGroupResponse) GetBody() *CreatePipelineGroupResponseBody {
	return s.Body
}

func (s *CreatePipelineGroupResponse) SetHeaders(v map[string]*string) *CreatePipelineGroupResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineGroupResponse) SetStatusCode(v int32) *CreatePipelineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineGroupResponse) SetBody(v *CreatePipelineGroupResponseBody) *CreatePipelineGroupResponse {
	s.Body = v
	return s
}

func (s *CreatePipelineGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
