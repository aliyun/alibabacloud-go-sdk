// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDeployGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeDeployGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeDeployGroupResponse
	GetStatusCode() *int32
	SetBody(v *ChangeDeployGroupResponseBody) *ChangeDeployGroupResponse
	GetBody() *ChangeDeployGroupResponseBody
}

type ChangeDeployGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeDeployGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeDeployGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeDeployGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeDeployGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeDeployGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeDeployGroupResponse) GetBody() *ChangeDeployGroupResponseBody {
	return s.Body
}

func (s *ChangeDeployGroupResponse) SetHeaders(v map[string]*string) *ChangeDeployGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeDeployGroupResponse) SetStatusCode(v int32) *ChangeDeployGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeDeployGroupResponse) SetBody(v *ChangeDeployGroupResponseBody) *ChangeDeployGroupResponse {
	s.Body = v
	return s
}

func (s *ChangeDeployGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
