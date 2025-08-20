// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemplateSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemplateSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetTemplateSummaryResponseBody) *GetTemplateSummaryResponse
	GetBody() *GetTemplateSummaryResponseBody
}

type GetTemplateSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemplateSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemplateSummaryResponse) GetBody() *GetTemplateSummaryResponseBody {
	return s.Body
}

func (s *GetTemplateSummaryResponse) SetHeaders(v map[string]*string) *GetTemplateSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateSummaryResponse) SetStatusCode(v int32) *GetTemplateSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateSummaryResponse) SetBody(v *GetTemplateSummaryResponseBody) *GetTemplateSummaryResponse {
	s.Body = v
	return s
}

func (s *GetTemplateSummaryResponse) Validate() error {
	return dara.Validate(s)
}
