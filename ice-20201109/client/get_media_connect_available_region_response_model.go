// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectAvailableRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaConnectAvailableRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaConnectAvailableRegionResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaConnectAvailableRegionResponseBody) *GetMediaConnectAvailableRegionResponse
	GetBody() *GetMediaConnectAvailableRegionResponseBody
}

type GetMediaConnectAvailableRegionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaConnectAvailableRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaConnectAvailableRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectAvailableRegionResponse) GoString() string {
	return s.String()
}

func (s *GetMediaConnectAvailableRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaConnectAvailableRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaConnectAvailableRegionResponse) GetBody() *GetMediaConnectAvailableRegionResponseBody {
	return s.Body
}

func (s *GetMediaConnectAvailableRegionResponse) SetHeaders(v map[string]*string) *GetMediaConnectAvailableRegionResponse {
	s.Headers = v
	return s
}

func (s *GetMediaConnectAvailableRegionResponse) SetStatusCode(v int32) *GetMediaConnectAvailableRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaConnectAvailableRegionResponse) SetBody(v *GetMediaConnectAvailableRegionResponseBody) *GetMediaConnectAvailableRegionResponse {
	s.Body = v
	return s
}

func (s *GetMediaConnectAvailableRegionResponse) Validate() error {
	return dara.Validate(s)
}
