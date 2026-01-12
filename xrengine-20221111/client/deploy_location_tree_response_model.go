// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployLocationTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployLocationTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployLocationTreeResponse
	GetStatusCode() *int32
	SetBody(v *DeployLocationTreeResponseBody) *DeployLocationTreeResponse
	GetBody() *DeployLocationTreeResponseBody
}

type DeployLocationTreeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployLocationTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployLocationTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployLocationTreeResponse) GoString() string {
	return s.String()
}

func (s *DeployLocationTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployLocationTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployLocationTreeResponse) GetBody() *DeployLocationTreeResponseBody {
	return s.Body
}

func (s *DeployLocationTreeResponse) SetHeaders(v map[string]*string) *DeployLocationTreeResponse {
	s.Headers = v
	return s
}

func (s *DeployLocationTreeResponse) SetStatusCode(v int32) *DeployLocationTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployLocationTreeResponse) SetBody(v *DeployLocationTreeResponseBody) *DeployLocationTreeResponse {
	s.Body = v
	return s
}

func (s *DeployLocationTreeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
