// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReportTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteReportTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteReportTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteReportTemplateResponseBody) *DeleteReportTemplateResponse
	GetBody() *DeleteReportTemplateResponseBody
}

type DeleteReportTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReportTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReportTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteReportTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteReportTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteReportTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteReportTemplateResponse) GetBody() *DeleteReportTemplateResponseBody {
	return s.Body
}

func (s *DeleteReportTemplateResponse) SetHeaders(v map[string]*string) *DeleteReportTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteReportTemplateResponse) SetStatusCode(v int32) *DeleteReportTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReportTemplateResponse) SetBody(v *DeleteReportTemplateResponseBody) *DeleteReportTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteReportTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
