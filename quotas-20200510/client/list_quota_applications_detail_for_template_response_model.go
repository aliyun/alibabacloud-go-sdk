// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaApplicationsDetailForTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotaApplicationsDetailForTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotaApplicationsDetailForTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotaApplicationsDetailForTemplateResponseBody) *ListQuotaApplicationsDetailForTemplateResponse
	GetBody() *ListQuotaApplicationsDetailForTemplateResponseBody
}

type ListQuotaApplicationsDetailForTemplateResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaApplicationsDetailForTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaApplicationsDetailForTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsDetailForTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsDetailForTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotaApplicationsDetailForTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotaApplicationsDetailForTemplateResponse) GetBody() *ListQuotaApplicationsDetailForTemplateResponseBody {
	return s.Body
}

func (s *ListQuotaApplicationsDetailForTemplateResponse) SetHeaders(v map[string]*string) *ListQuotaApplicationsDetailForTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponse) SetStatusCode(v int32) *ListQuotaApplicationsDetailForTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponse) SetBody(v *ListQuotaApplicationsDetailForTemplateResponseBody) *ListQuotaApplicationsDetailForTemplateResponse {
	s.Body = v
	return s
}

func (s *ListQuotaApplicationsDetailForTemplateResponse) Validate() error {
	return dara.Validate(s)
}
