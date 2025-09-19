// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCustomizeReportConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveCustomizeReportConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveCustomizeReportConfigResponse
	GetStatusCode() *int32
	SetBody(v *SaveCustomizeReportConfigResponseBody) *SaveCustomizeReportConfigResponse
	GetBody() *SaveCustomizeReportConfigResponseBody
}

type SaveCustomizeReportConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveCustomizeReportConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveCustomizeReportConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveCustomizeReportConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveCustomizeReportConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveCustomizeReportConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveCustomizeReportConfigResponse) GetBody() *SaveCustomizeReportConfigResponseBody {
	return s.Body
}

func (s *SaveCustomizeReportConfigResponse) SetHeaders(v map[string]*string) *SaveCustomizeReportConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveCustomizeReportConfigResponse) SetStatusCode(v int32) *SaveCustomizeReportConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveCustomizeReportConfigResponse) SetBody(v *SaveCustomizeReportConfigResponseBody) *SaveCustomizeReportConfigResponse {
	s.Body = v
	return s
}

func (s *SaveCustomizeReportConfigResponse) Validate() error {
	return dara.Validate(s)
}
