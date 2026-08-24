// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVulScanGlobalConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVulScanGlobalConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVulScanGlobalConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVulScanGlobalConfigResponseBody) *UpdateVulScanGlobalConfigResponse
	GetBody() *UpdateVulScanGlobalConfigResponseBody
}

type UpdateVulScanGlobalConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVulScanGlobalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVulScanGlobalConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVulScanGlobalConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateVulScanGlobalConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVulScanGlobalConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVulScanGlobalConfigResponse) GetBody() *UpdateVulScanGlobalConfigResponseBody {
	return s.Body
}

func (s *UpdateVulScanGlobalConfigResponse) SetHeaders(v map[string]*string) *UpdateVulScanGlobalConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateVulScanGlobalConfigResponse) SetStatusCode(v int32) *UpdateVulScanGlobalConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVulScanGlobalConfigResponse) SetBody(v *UpdateVulScanGlobalConfigResponseBody) *UpdateVulScanGlobalConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateVulScanGlobalConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
