// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPipelineGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPipelineGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetPipelineGroupResponseBody) *GetPipelineGroupResponse
	GetBody() *GetPipelineGroupResponseBody
}

type GetPipelineGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPipelineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineGroupResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPipelineGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPipelineGroupResponse) GetBody() *GetPipelineGroupResponseBody {
	return s.Body
}

func (s *GetPipelineGroupResponse) SetHeaders(v map[string]*string) *GetPipelineGroupResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineGroupResponse) SetStatusCode(v int32) *GetPipelineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineGroupResponse) SetBody(v *GetPipelineGroupResponseBody) *GetPipelineGroupResponse {
	s.Body = v
	return s
}

func (s *GetPipelineGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
