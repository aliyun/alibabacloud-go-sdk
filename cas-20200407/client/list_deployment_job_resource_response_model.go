// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentJobResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeploymentJobResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeploymentJobResourceResponse
	GetStatusCode() *int32
	SetBody(v *ListDeploymentJobResourceResponseBody) *ListDeploymentJobResourceResponse
	GetBody() *ListDeploymentJobResourceResponseBody
}

type ListDeploymentJobResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentJobResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentJobResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobResourceResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeploymentJobResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeploymentJobResourceResponse) GetBody() *ListDeploymentJobResourceResponseBody {
	return s.Body
}

func (s *ListDeploymentJobResourceResponse) SetHeaders(v map[string]*string) *ListDeploymentJobResourceResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentJobResourceResponse) SetStatusCode(v int32) *ListDeploymentJobResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentJobResourceResponse) SetBody(v *ListDeploymentJobResourceResponseBody) *ListDeploymentJobResourceResponse {
	s.Body = v
	return s
}

func (s *ListDeploymentJobResourceResponse) Validate() error {
	return dara.Validate(s)
}
