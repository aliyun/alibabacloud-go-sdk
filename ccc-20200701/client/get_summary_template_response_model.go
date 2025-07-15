// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSummaryTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSummaryTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSummaryTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetSummaryTemplateResponseBody) *GetSummaryTemplateResponse
	GetBody() *GetSummaryTemplateResponseBody
}

type GetSummaryTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSummaryTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSummaryTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetSummaryTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSummaryTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSummaryTemplateResponse) GetBody() *GetSummaryTemplateResponseBody {
	return s.Body
}

func (s *GetSummaryTemplateResponse) SetHeaders(v map[string]*string) *GetSummaryTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetSummaryTemplateResponse) SetStatusCode(v int32) *GetSummaryTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSummaryTemplateResponse) SetBody(v *GetSummaryTemplateResponseBody) *GetSummaryTemplateResponse {
	s.Body = v
	return s
}

func (s *GetSummaryTemplateResponse) Validate() error {
	return dara.Validate(s)
}
