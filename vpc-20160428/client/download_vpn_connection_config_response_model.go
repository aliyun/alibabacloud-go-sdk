// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadVpnConnectionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadVpnConnectionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadVpnConnectionConfigResponse
	GetStatusCode() *int32
	SetBody(v *DownloadVpnConnectionConfigResponseBody) *DownloadVpnConnectionConfigResponse
	GetBody() *DownloadVpnConnectionConfigResponseBody
}

type DownloadVpnConnectionConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadVpnConnectionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadVpnConnectionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadVpnConnectionConfigResponse) GoString() string {
	return s.String()
}

func (s *DownloadVpnConnectionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadVpnConnectionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadVpnConnectionConfigResponse) GetBody() *DownloadVpnConnectionConfigResponseBody {
	return s.Body
}

func (s *DownloadVpnConnectionConfigResponse) SetHeaders(v map[string]*string) *DownloadVpnConnectionConfigResponse {
	s.Headers = v
	return s
}

func (s *DownloadVpnConnectionConfigResponse) SetStatusCode(v int32) *DownloadVpnConnectionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadVpnConnectionConfigResponse) SetBody(v *DownloadVpnConnectionConfigResponseBody) *DownloadVpnConnectionConfigResponse {
	s.Body = v
	return s
}

func (s *DownloadVpnConnectionConfigResponse) Validate() error {
	return dara.Validate(s)
}
