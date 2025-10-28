// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeployGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeployGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeployGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListDeployGroupResponseBody) *ListDeployGroupResponse
	GetBody() *ListDeployGroupResponseBody
}

type ListDeployGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeployGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeployGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeployGroupResponse) GoString() string {
	return s.String()
}

func (s *ListDeployGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeployGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeployGroupResponse) GetBody() *ListDeployGroupResponseBody {
	return s.Body
}

func (s *ListDeployGroupResponse) SetHeaders(v map[string]*string) *ListDeployGroupResponse {
	s.Headers = v
	return s
}

func (s *ListDeployGroupResponse) SetStatusCode(v int32) *ListDeployGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeployGroupResponse) SetBody(v *ListDeployGroupResponseBody) *ListDeployGroupResponse {
	s.Body = v
	return s
}

func (s *ListDeployGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
