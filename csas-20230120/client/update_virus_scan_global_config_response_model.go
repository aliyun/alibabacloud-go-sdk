// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirusScanGlobalConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVirusScanGlobalConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVirusScanGlobalConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVirusScanGlobalConfigResponseBody) *UpdateVirusScanGlobalConfigResponse
	GetBody() *UpdateVirusScanGlobalConfigResponseBody
}

type UpdateVirusScanGlobalConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVirusScanGlobalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVirusScanGlobalConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirusScanGlobalConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateVirusScanGlobalConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVirusScanGlobalConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVirusScanGlobalConfigResponse) GetBody() *UpdateVirusScanGlobalConfigResponseBody {
	return s.Body
}

func (s *UpdateVirusScanGlobalConfigResponse) SetHeaders(v map[string]*string) *UpdateVirusScanGlobalConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateVirusScanGlobalConfigResponse) SetStatusCode(v int32) *UpdateVirusScanGlobalConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVirusScanGlobalConfigResponse) SetBody(v *UpdateVirusScanGlobalConfigResponseBody) *UpdateVirusScanGlobalConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateVirusScanGlobalConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
