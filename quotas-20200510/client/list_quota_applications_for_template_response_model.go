// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaApplicationsForTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotaApplicationsForTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotaApplicationsForTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotaApplicationsForTemplateResponseBody) *ListQuotaApplicationsForTemplateResponse
	GetBody() *ListQuotaApplicationsForTemplateResponseBody
}

type ListQuotaApplicationsForTemplateResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaApplicationsForTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaApplicationsForTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsForTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsForTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotaApplicationsForTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotaApplicationsForTemplateResponse) GetBody() *ListQuotaApplicationsForTemplateResponseBody {
	return s.Body
}

func (s *ListQuotaApplicationsForTemplateResponse) SetHeaders(v map[string]*string) *ListQuotaApplicationsForTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponse) SetStatusCode(v int32) *ListQuotaApplicationsForTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponse) SetBody(v *ListQuotaApplicationsForTemplateResponseBody) *ListQuotaApplicationsForTemplateResponse {
	s.Body = v
	return s
}

func (s *ListQuotaApplicationsForTemplateResponse) Validate() error {
	return dara.Validate(s)
}
