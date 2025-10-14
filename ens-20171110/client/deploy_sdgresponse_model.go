// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploySDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeploySDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeploySDGResponse
	GetStatusCode() *int32
	SetBody(v *DeploySDGResponseBody) *DeploySDGResponse
	GetBody() *DeploySDGResponseBody
}

type DeploySDGResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeploySDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeploySDGResponse) String() string {
	return dara.Prettify(s)
}

func (s DeploySDGResponse) GoString() string {
	return s.String()
}

func (s *DeploySDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeploySDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeploySDGResponse) GetBody() *DeploySDGResponseBody {
	return s.Body
}

func (s *DeploySDGResponse) SetHeaders(v map[string]*string) *DeploySDGResponse {
	s.Headers = v
	return s
}

func (s *DeploySDGResponse) SetStatusCode(v int32) *DeploySDGResponse {
	s.StatusCode = &v
	return s
}

func (s *DeploySDGResponse) SetBody(v *DeploySDGResponseBody) *DeploySDGResponse {
	s.Body = v
	return s
}

func (s *DeploySDGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
