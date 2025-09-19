// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCustomizeReportConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyCustomizeReportConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyCustomizeReportConfigResponse
	GetStatusCode() *int32
	SetBody(v *CopyCustomizeReportConfigResponseBody) *CopyCustomizeReportConfigResponse
	GetBody() *CopyCustomizeReportConfigResponseBody
}

type CopyCustomizeReportConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyCustomizeReportConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyCustomizeReportConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyCustomizeReportConfigResponse) GoString() string {
	return s.String()
}

func (s *CopyCustomizeReportConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyCustomizeReportConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyCustomizeReportConfigResponse) GetBody() *CopyCustomizeReportConfigResponseBody {
	return s.Body
}

func (s *CopyCustomizeReportConfigResponse) SetHeaders(v map[string]*string) *CopyCustomizeReportConfigResponse {
	s.Headers = v
	return s
}

func (s *CopyCustomizeReportConfigResponse) SetStatusCode(v int32) *CopyCustomizeReportConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyCustomizeReportConfigResponse) SetBody(v *CopyCustomizeReportConfigResponseBody) *CopyCustomizeReportConfigResponse {
	s.Body = v
	return s
}

func (s *CopyCustomizeReportConfigResponse) Validate() error {
	return dara.Validate(s)
}
