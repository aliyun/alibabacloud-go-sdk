// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineManagementConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePipelineManagementConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePipelineManagementConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePipelineManagementConfigResponseBody) *UpdatePipelineManagementConfigResponse
	GetBody() *UpdatePipelineManagementConfigResponseBody
}

type UpdatePipelineManagementConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePipelineManagementConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePipelineManagementConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineManagementConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineManagementConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePipelineManagementConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePipelineManagementConfigResponse) GetBody() *UpdatePipelineManagementConfigResponseBody {
	return s.Body
}

func (s *UpdatePipelineManagementConfigResponse) SetHeaders(v map[string]*string) *UpdatePipelineManagementConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineManagementConfigResponse) SetStatusCode(v int32) *UpdatePipelineManagementConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelineManagementConfigResponse) SetBody(v *UpdatePipelineManagementConfigResponseBody) *UpdatePipelineManagementConfigResponse {
	s.Body = v
	return s
}

func (s *UpdatePipelineManagementConfigResponse) Validate() error {
	return dara.Validate(s)
}
