// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceUpgradeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceUpgradeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceUpgradeStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceUpgradeStatusResponseBody) *GetDeviceUpgradeStatusResponse
	GetBody() *GetDeviceUpgradeStatusResponseBody
}

type GetDeviceUpgradeStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceUpgradeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceUpgradeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceUpgradeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceUpgradeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceUpgradeStatusResponse) GetBody() *GetDeviceUpgradeStatusResponseBody {
	return s.Body
}

func (s *GetDeviceUpgradeStatusResponse) SetHeaders(v map[string]*string) *GetDeviceUpgradeStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceUpgradeStatusResponse) SetStatusCode(v int32) *GetDeviceUpgradeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceUpgradeStatusResponse) SetBody(v *GetDeviceUpgradeStatusResponseBody) *GetDeviceUpgradeStatusResponse {
	s.Body = v
	return s
}

func (s *GetDeviceUpgradeStatusResponse) Validate() error {
	return dara.Validate(s)
}
