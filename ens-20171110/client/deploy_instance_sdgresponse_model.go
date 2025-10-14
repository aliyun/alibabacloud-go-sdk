// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployInstanceSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployInstanceSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployInstanceSDGResponse
	GetStatusCode() *int32
	SetBody(v *DeployInstanceSDGResponseBody) *DeployInstanceSDGResponse
	GetBody() *DeployInstanceSDGResponseBody
}

type DeployInstanceSDGResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployInstanceSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployInstanceSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployInstanceSDGResponse) GoString() string {
	return s.String()
}

func (s *DeployInstanceSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployInstanceSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployInstanceSDGResponse) GetBody() *DeployInstanceSDGResponseBody {
	return s.Body
}

func (s *DeployInstanceSDGResponse) SetHeaders(v map[string]*string) *DeployInstanceSDGResponse {
	s.Headers = v
	return s
}

func (s *DeployInstanceSDGResponse) SetStatusCode(v int32) *DeployInstanceSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployInstanceSDGResponse) SetBody(v *DeployInstanceSDGResponseBody) *DeployInstanceSDGResponse {
	s.Body = v
	return s
}

func (s *DeployInstanceSDGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
