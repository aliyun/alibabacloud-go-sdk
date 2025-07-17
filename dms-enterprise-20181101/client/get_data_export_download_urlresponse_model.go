// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataExportDownloadURLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataExportDownloadURLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataExportDownloadURLResponse
	GetStatusCode() *int32
	SetBody(v *GetDataExportDownloadURLResponseBody) *GetDataExportDownloadURLResponse
	GetBody() *GetDataExportDownloadURLResponseBody
}

type GetDataExportDownloadURLResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataExportDownloadURLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataExportDownloadURLResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportDownloadURLResponse) GoString() string {
	return s.String()
}

func (s *GetDataExportDownloadURLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataExportDownloadURLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataExportDownloadURLResponse) GetBody() *GetDataExportDownloadURLResponseBody {
	return s.Body
}

func (s *GetDataExportDownloadURLResponse) SetHeaders(v map[string]*string) *GetDataExportDownloadURLResponse {
	s.Headers = v
	return s
}

func (s *GetDataExportDownloadURLResponse) SetStatusCode(v int32) *GetDataExportDownloadURLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataExportDownloadURLResponse) SetBody(v *GetDataExportDownloadURLResponseBody) *GetDataExportDownloadURLResponse {
	s.Body = v
	return s
}

func (s *GetDataExportDownloadURLResponse) Validate() error {
	return dara.Validate(s)
}
