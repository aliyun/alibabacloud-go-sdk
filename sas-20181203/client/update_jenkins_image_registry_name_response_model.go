// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJenkinsImageRegistryNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateJenkinsImageRegistryNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateJenkinsImageRegistryNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateJenkinsImageRegistryNameResponseBody) *UpdateJenkinsImageRegistryNameResponse
	GetBody() *UpdateJenkinsImageRegistryNameResponseBody
}

type UpdateJenkinsImageRegistryNameResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJenkinsImageRegistryNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJenkinsImageRegistryNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateJenkinsImageRegistryNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateJenkinsImageRegistryNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateJenkinsImageRegistryNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateJenkinsImageRegistryNameResponse) GetBody() *UpdateJenkinsImageRegistryNameResponseBody {
	return s.Body
}

func (s *UpdateJenkinsImageRegistryNameResponse) SetHeaders(v map[string]*string) *UpdateJenkinsImageRegistryNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateJenkinsImageRegistryNameResponse) SetStatusCode(v int32) *UpdateJenkinsImageRegistryNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJenkinsImageRegistryNameResponse) SetBody(v *UpdateJenkinsImageRegistryNameResponseBody) *UpdateJenkinsImageRegistryNameResponse {
	s.Body = v
	return s
}

func (s *UpdateJenkinsImageRegistryNameResponse) Validate() error {
	return dara.Validate(s)
}
