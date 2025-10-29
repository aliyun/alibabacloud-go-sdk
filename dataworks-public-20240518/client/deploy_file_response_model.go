// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployFileResponse
	GetStatusCode() *int32
	SetBody(v *DeployFileResponseBody) *DeployFileResponse
	GetBody() *DeployFileResponseBody
}

type DeployFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployFileResponse) GoString() string {
	return s.String()
}

func (s *DeployFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployFileResponse) GetBody() *DeployFileResponseBody {
	return s.Body
}

func (s *DeployFileResponse) SetHeaders(v map[string]*string) *DeployFileResponse {
	s.Headers = v
	return s
}

func (s *DeployFileResponse) SetStatusCode(v int32) *DeployFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployFileResponse) SetBody(v *DeployFileResponseBody) *DeployFileResponse {
	s.Body = v
	return s
}

func (s *DeployFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
