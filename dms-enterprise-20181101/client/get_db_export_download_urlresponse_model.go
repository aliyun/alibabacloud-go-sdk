// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDbExportDownloadURLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDbExportDownloadURLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDbExportDownloadURLResponse
	GetStatusCode() *int32
	SetBody(v *GetDbExportDownloadURLResponseBody) *GetDbExportDownloadURLResponse
	GetBody() *GetDbExportDownloadURLResponseBody
}

type GetDbExportDownloadURLResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDbExportDownloadURLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDbExportDownloadURLResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDbExportDownloadURLResponse) GoString() string {
	return s.String()
}

func (s *GetDbExportDownloadURLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDbExportDownloadURLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDbExportDownloadURLResponse) GetBody() *GetDbExportDownloadURLResponseBody {
	return s.Body
}

func (s *GetDbExportDownloadURLResponse) SetHeaders(v map[string]*string) *GetDbExportDownloadURLResponse {
	s.Headers = v
	return s
}

func (s *GetDbExportDownloadURLResponse) SetStatusCode(v int32) *GetDbExportDownloadURLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDbExportDownloadURLResponse) SetBody(v *GetDbExportDownloadURLResponseBody) *GetDbExportDownloadURLResponse {
	s.Body = v
	return s
}

func (s *GetDbExportDownloadURLResponse) Validate() error {
	return dara.Validate(s)
}
