// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApplicationGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployApplicationGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployApplicationGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeployApplicationGroupResponseBody) *DeployApplicationGroupResponse
	GetBody() *DeployApplicationGroupResponseBody
}

type DeployApplicationGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployApplicationGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *DeployApplicationGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployApplicationGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployApplicationGroupResponse) GetBody() *DeployApplicationGroupResponseBody {
	return s.Body
}

func (s *DeployApplicationGroupResponse) SetHeaders(v map[string]*string) *DeployApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *DeployApplicationGroupResponse) SetStatusCode(v int32) *DeployApplicationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployApplicationGroupResponse) SetBody(v *DeployApplicationGroupResponseBody) *DeployApplicationGroupResponse {
	s.Body = v
	return s
}

func (s *DeployApplicationGroupResponse) Validate() error {
	return dara.Validate(s)
}
