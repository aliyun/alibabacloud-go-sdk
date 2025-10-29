// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataQualityTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataQualityTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataQualityTemplateResponseBody) *CreateDataQualityTemplateResponse
	GetBody() *CreateDataQualityTemplateResponseBody
}

type CreateDataQualityTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataQualityTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataQualityTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateDataQualityTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataQualityTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataQualityTemplateResponse) GetBody() *CreateDataQualityTemplateResponseBody {
	return s.Body
}

func (s *CreateDataQualityTemplateResponse) SetHeaders(v map[string]*string) *CreateDataQualityTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateDataQualityTemplateResponse) SetStatusCode(v int32) *CreateDataQualityTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataQualityTemplateResponse) SetBody(v *CreateDataQualityTemplateResponseBody) *CreateDataQualityTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateDataQualityTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
