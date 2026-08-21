// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOnlineHeatmapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceOnlineHeatmapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceOnlineHeatmapResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceOnlineHeatmapResponseBody) *GetDeviceOnlineHeatmapResponse
	GetBody() *GetDeviceOnlineHeatmapResponseBody
}

type GetDeviceOnlineHeatmapResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceOnlineHeatmapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceOnlineHeatmapResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOnlineHeatmapResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceOnlineHeatmapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceOnlineHeatmapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceOnlineHeatmapResponse) GetBody() *GetDeviceOnlineHeatmapResponseBody {
	return s.Body
}

func (s *GetDeviceOnlineHeatmapResponse) SetHeaders(v map[string]*string) *GetDeviceOnlineHeatmapResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceOnlineHeatmapResponse) SetStatusCode(v int32) *GetDeviceOnlineHeatmapResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceOnlineHeatmapResponse) SetBody(v *GetDeviceOnlineHeatmapResponseBody) *GetDeviceOnlineHeatmapResponse {
	s.Body = v
	return s
}

func (s *GetDeviceOnlineHeatmapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
