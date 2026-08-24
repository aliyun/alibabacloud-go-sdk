// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulScanGlobalConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVulScanGlobalConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVulScanGlobalConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetVulScanGlobalConfigResponseBody) *GetVulScanGlobalConfigResponse
	GetBody() *GetVulScanGlobalConfigResponseBody
}

type GetVulScanGlobalConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulScanGlobalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulScanGlobalConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVulScanGlobalConfigResponse) GoString() string {
	return s.String()
}

func (s *GetVulScanGlobalConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVulScanGlobalConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVulScanGlobalConfigResponse) GetBody() *GetVulScanGlobalConfigResponseBody {
	return s.Body
}

func (s *GetVulScanGlobalConfigResponse) SetHeaders(v map[string]*string) *GetVulScanGlobalConfigResponse {
	s.Headers = v
	return s
}

func (s *GetVulScanGlobalConfigResponse) SetStatusCode(v int32) *GetVulScanGlobalConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulScanGlobalConfigResponse) SetBody(v *GetVulScanGlobalConfigResponseBody) *GetVulScanGlobalConfigResponse {
	s.Body = v
	return s
}

func (s *GetVulScanGlobalConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
