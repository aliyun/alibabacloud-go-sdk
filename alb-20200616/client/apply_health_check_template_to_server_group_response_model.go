// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyHealthCheckTemplateToServerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyHealthCheckTemplateToServerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyHealthCheckTemplateToServerGroupResponse
	GetStatusCode() *int32
	SetBody(v *ApplyHealthCheckTemplateToServerGroupResponseBody) *ApplyHealthCheckTemplateToServerGroupResponse
	GetBody() *ApplyHealthCheckTemplateToServerGroupResponseBody
}

type ApplyHealthCheckTemplateToServerGroupResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyHealthCheckTemplateToServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyHealthCheckTemplateToServerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyHealthCheckTemplateToServerGroupResponse) GoString() string {
	return s.String()
}

func (s *ApplyHealthCheckTemplateToServerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyHealthCheckTemplateToServerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyHealthCheckTemplateToServerGroupResponse) GetBody() *ApplyHealthCheckTemplateToServerGroupResponseBody {
	return s.Body
}

func (s *ApplyHealthCheckTemplateToServerGroupResponse) SetHeaders(v map[string]*string) *ApplyHealthCheckTemplateToServerGroupResponse {
	s.Headers = v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupResponse) SetStatusCode(v int32) *ApplyHealthCheckTemplateToServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupResponse) SetBody(v *ApplyHealthCheckTemplateToServerGroupResponseBody) *ApplyHealthCheckTemplateToServerGroupResponse {
	s.Body = v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
