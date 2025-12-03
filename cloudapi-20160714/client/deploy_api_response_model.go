// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployApiResponse
	GetStatusCode() *int32
	SetBody(v *DeployApiResponseBody) *DeployApiResponse
	GetBody() *DeployApiResponseBody
}

type DeployApiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployApiResponse) GoString() string {
	return s.String()
}

func (s *DeployApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployApiResponse) GetBody() *DeployApiResponseBody {
	return s.Body
}

func (s *DeployApiResponse) SetHeaders(v map[string]*string) *DeployApiResponse {
	s.Headers = v
	return s
}

func (s *DeployApiResponse) SetStatusCode(v int32) *DeployApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployApiResponse) SetBody(v *DeployApiResponseBody) *DeployApiResponse {
	s.Body = v
	return s
}

func (s *DeployApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
