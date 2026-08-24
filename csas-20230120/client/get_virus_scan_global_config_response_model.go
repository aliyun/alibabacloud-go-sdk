// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirusScanGlobalConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVirusScanGlobalConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVirusScanGlobalConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetVirusScanGlobalConfigResponseBody) *GetVirusScanGlobalConfigResponse
	GetBody() *GetVirusScanGlobalConfigResponseBody
}

type GetVirusScanGlobalConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVirusScanGlobalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVirusScanGlobalConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanGlobalConfigResponse) GoString() string {
	return s.String()
}

func (s *GetVirusScanGlobalConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVirusScanGlobalConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVirusScanGlobalConfigResponse) GetBody() *GetVirusScanGlobalConfigResponseBody {
	return s.Body
}

func (s *GetVirusScanGlobalConfigResponse) SetHeaders(v map[string]*string) *GetVirusScanGlobalConfigResponse {
	s.Headers = v
	return s
}

func (s *GetVirusScanGlobalConfigResponse) SetStatusCode(v int32) *GetVirusScanGlobalConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVirusScanGlobalConfigResponse) SetBody(v *GetVirusScanGlobalConfigResponseBody) *GetVirusScanGlobalConfigResponse {
	s.Body = v
	return s
}

func (s *GetVirusScanGlobalConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
