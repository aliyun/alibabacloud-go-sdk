// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirusScanConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVirusScanConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVirusScanConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetVirusScanConfigResponseBody) *GetVirusScanConfigResponse
	GetBody() *GetVirusScanConfigResponseBody
}

type GetVirusScanConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVirusScanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVirusScanConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanConfigResponse) GoString() string {
	return s.String()
}

func (s *GetVirusScanConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVirusScanConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVirusScanConfigResponse) GetBody() *GetVirusScanConfigResponseBody {
	return s.Body
}

func (s *GetVirusScanConfigResponse) SetHeaders(v map[string]*string) *GetVirusScanConfigResponse {
	s.Headers = v
	return s
}

func (s *GetVirusScanConfigResponse) SetStatusCode(v int32) *GetVirusScanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVirusScanConfigResponse) SetBody(v *GetVirusScanConfigResponseBody) *GetVirusScanConfigResponse {
	s.Body = v
	return s
}

func (s *GetVirusScanConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
