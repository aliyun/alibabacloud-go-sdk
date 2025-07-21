// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExportDeviceInfoOssUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExportDeviceInfoOssUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExportDeviceInfoOssUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetExportDeviceInfoOssUrlResponseBody) *GetExportDeviceInfoOssUrlResponse
	GetBody() *GetExportDeviceInfoOssUrlResponseBody
}

type GetExportDeviceInfoOssUrlResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExportDeviceInfoOssUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExportDeviceInfoOssUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExportDeviceInfoOssUrlResponse) GoString() string {
	return s.String()
}

func (s *GetExportDeviceInfoOssUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExportDeviceInfoOssUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExportDeviceInfoOssUrlResponse) GetBody() *GetExportDeviceInfoOssUrlResponseBody {
	return s.Body
}

func (s *GetExportDeviceInfoOssUrlResponse) SetHeaders(v map[string]*string) *GetExportDeviceInfoOssUrlResponse {
	s.Headers = v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponse) SetStatusCode(v int32) *GetExportDeviceInfoOssUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponse) SetBody(v *GetExportDeviceInfoOssUrlResponseBody) *GetExportDeviceInfoOssUrlResponse {
	s.Body = v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponse) Validate() error {
	return dara.Validate(s)
}
