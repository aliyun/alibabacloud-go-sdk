// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeploymentJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeploymentJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeploymentJobResponseBody) *DeleteDeploymentJobResponse
	GetBody() *DeleteDeploymentJobResponseBody
}

type DeleteDeploymentJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeploymentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeploymentJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeploymentJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeploymentJobResponse) GetBody() *DeleteDeploymentJobResponseBody {
	return s.Body
}

func (s *DeleteDeploymentJobResponse) SetHeaders(v map[string]*string) *DeleteDeploymentJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeploymentJobResponse) SetStatusCode(v int32) *DeleteDeploymentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeploymentJobResponse) SetBody(v *DeleteDeploymentJobResponseBody) *DeleteDeploymentJobResponse {
	s.Body = v
	return s
}

func (s *DeleteDeploymentJobResponse) Validate() error {
	return dara.Validate(s)
}
