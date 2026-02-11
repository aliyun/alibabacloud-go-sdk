// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployModelFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployModelFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployModelFileResponse
	GetStatusCode() *int32
	SetBody(v *DeployModelFileResponseBody) *DeployModelFileResponse
	GetBody() *DeployModelFileResponseBody
}

type DeployModelFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployModelFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployModelFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployModelFileResponse) GoString() string {
	return s.String()
}

func (s *DeployModelFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployModelFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployModelFileResponse) GetBody() *DeployModelFileResponseBody {
	return s.Body
}

func (s *DeployModelFileResponse) SetHeaders(v map[string]*string) *DeployModelFileResponse {
	s.Headers = v
	return s
}

func (s *DeployModelFileResponse) SetStatusCode(v int32) *DeployModelFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployModelFileResponse) SetBody(v *DeployModelFileResponseBody) *DeployModelFileResponse {
	s.Body = v
	return s
}

func (s *DeployModelFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
