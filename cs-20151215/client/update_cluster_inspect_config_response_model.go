// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterInspectConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClusterInspectConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClusterInspectConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClusterInspectConfigResponseBody) *UpdateClusterInspectConfigResponse
	GetBody() *UpdateClusterInspectConfigResponseBody
}

type UpdateClusterInspectConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClusterInspectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClusterInspectConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterInspectConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterInspectConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClusterInspectConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClusterInspectConfigResponse) GetBody() *UpdateClusterInspectConfigResponseBody {
	return s.Body
}

func (s *UpdateClusterInspectConfigResponse) SetHeaders(v map[string]*string) *UpdateClusterInspectConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterInspectConfigResponse) SetStatusCode(v int32) *UpdateClusterInspectConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterInspectConfigResponse) SetBody(v *UpdateClusterInspectConfigResponseBody) *UpdateClusterInspectConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateClusterInspectConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
