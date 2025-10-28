// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeployGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeployGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeployGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeployGroupResponseBody) *DeleteDeployGroupResponse
	GetBody() *DeleteDeployGroupResponseBody
}

type DeleteDeployGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeployGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeployGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeployGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeployGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeployGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeployGroupResponse) GetBody() *DeleteDeployGroupResponseBody {
	return s.Body
}

func (s *DeleteDeployGroupResponse) SetHeaders(v map[string]*string) *DeleteDeployGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeployGroupResponse) SetStatusCode(v int32) *DeleteDeployGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeployGroupResponse) SetBody(v *DeleteDeployGroupResponseBody) *DeleteDeployGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteDeployGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
