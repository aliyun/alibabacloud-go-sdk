// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataQualityTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataQualityTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetDataQualityTemplateResponseBody) *GetDataQualityTemplateResponse
	GetBody() *GetDataQualityTemplateResponseBody
}

type GetDataQualityTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataQualityTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataQualityTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetDataQualityTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataQualityTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataQualityTemplateResponse) GetBody() *GetDataQualityTemplateResponseBody {
	return s.Body
}

func (s *GetDataQualityTemplateResponse) SetHeaders(v map[string]*string) *GetDataQualityTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetDataQualityTemplateResponse) SetStatusCode(v int32) *GetDataQualityTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataQualityTemplateResponse) SetBody(v *GetDataQualityTemplateResponseBody) *GetDataQualityTemplateResponse {
	s.Body = v
	return s
}

func (s *GetDataQualityTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
