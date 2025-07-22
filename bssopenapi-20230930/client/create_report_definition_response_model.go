// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReportDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReportDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *CreateReportDefinitionResponseBody) *CreateReportDefinitionResponse
	GetBody() *CreateReportDefinitionResponseBody
}

type CreateReportDefinitionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReportDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReportDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReportDefinitionResponse) GoString() string {
	return s.String()
}

func (s *CreateReportDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReportDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReportDefinitionResponse) GetBody() *CreateReportDefinitionResponseBody {
	return s.Body
}

func (s *CreateReportDefinitionResponse) SetHeaders(v map[string]*string) *CreateReportDefinitionResponse {
	s.Headers = v
	return s
}

func (s *CreateReportDefinitionResponse) SetStatusCode(v int32) *CreateReportDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReportDefinitionResponse) SetBody(v *CreateReportDefinitionResponseBody) *CreateReportDefinitionResponse {
	s.Body = v
	return s
}

func (s *CreateReportDefinitionResponse) Validate() error {
	return dara.Validate(s)
}
