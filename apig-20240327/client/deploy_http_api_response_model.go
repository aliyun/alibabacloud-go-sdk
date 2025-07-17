// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployHttpApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployHttpApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployHttpApiResponse
	GetStatusCode() *int32
	SetBody(v *DeployHttpApiResponseBody) *DeployHttpApiResponse
	GetBody() *DeployHttpApiResponseBody
}

type DeployHttpApiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployHttpApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployHttpApiResponse) GoString() string {
	return s.String()
}

func (s *DeployHttpApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployHttpApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployHttpApiResponse) GetBody() *DeployHttpApiResponseBody {
	return s.Body
}

func (s *DeployHttpApiResponse) SetHeaders(v map[string]*string) *DeployHttpApiResponse {
	s.Headers = v
	return s
}

func (s *DeployHttpApiResponse) SetStatusCode(v int32) *DeployHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployHttpApiResponse) SetBody(v *DeployHttpApiResponseBody) *DeployHttpApiResponse {
	s.Body = v
	return s
}

func (s *DeployHttpApiResponse) Validate() error {
	return dara.Validate(s)
}
