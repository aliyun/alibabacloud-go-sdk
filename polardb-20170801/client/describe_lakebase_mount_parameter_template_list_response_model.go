// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseMountParameterTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLakebaseMountParameterTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLakebaseMountParameterTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLakebaseMountParameterTemplateListResponseBody) *DescribeLakebaseMountParameterTemplateListResponse
	GetBody() *DescribeLakebaseMountParameterTemplateListResponseBody
}

type DescribeLakebaseMountParameterTemplateListResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLakebaseMountParameterTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLakebaseMountParameterTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterTemplateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLakebaseMountParameterTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLakebaseMountParameterTemplateListResponse) GetBody() *DescribeLakebaseMountParameterTemplateListResponseBody {
	return s.Body
}

func (s *DescribeLakebaseMountParameterTemplateListResponse) SetHeaders(v map[string]*string) *DescribeLakebaseMountParameterTemplateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListResponse) SetStatusCode(v int32) *DescribeLakebaseMountParameterTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListResponse) SetBody(v *DescribeLakebaseMountParameterTemplateListResponseBody) *DescribeLakebaseMountParameterTemplateListResponse {
	s.Body = v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
