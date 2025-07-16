// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportTemplateByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReportTemplateByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReportTemplateByNameResponse
	GetStatusCode() *int32
	SetBody(v *GetReportTemplateByNameResponseBody) *GetReportTemplateByNameResponse
	GetBody() *GetReportTemplateByNameResponseBody
}

type GetReportTemplateByNameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReportTemplateByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReportTemplateByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateByNameResponse) GoString() string {
	return s.String()
}

func (s *GetReportTemplateByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReportTemplateByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReportTemplateByNameResponse) GetBody() *GetReportTemplateByNameResponseBody {
	return s.Body
}

func (s *GetReportTemplateByNameResponse) SetHeaders(v map[string]*string) *GetReportTemplateByNameResponse {
	s.Headers = v
	return s
}

func (s *GetReportTemplateByNameResponse) SetStatusCode(v int32) *GetReportTemplateByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReportTemplateByNameResponse) SetBody(v *GetReportTemplateByNameResponseBody) *GetReportTemplateByNameResponse {
	s.Body = v
	return s
}

func (s *GetReportTemplateByNameResponse) Validate() error {
	return dara.Validate(s)
}
