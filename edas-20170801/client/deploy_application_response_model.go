// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployApplicationResponse
	GetStatusCode() *int32
	SetBody(v *DeployApplicationResponseBody) *DeployApplicationResponse
	GetBody() *DeployApplicationResponseBody
}

type DeployApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeployApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployApplicationResponse) GetBody() *DeployApplicationResponseBody {
	return s.Body
}

func (s *DeployApplicationResponse) SetHeaders(v map[string]*string) *DeployApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeployApplicationResponse) SetStatusCode(v int32) *DeployApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployApplicationResponse) SetBody(v *DeployApplicationResponseBody) *DeployApplicationResponse {
	s.Body = v
	return s
}

func (s *DeployApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
