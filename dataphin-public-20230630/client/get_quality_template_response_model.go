// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityTemplateResponseBody) *GetQualityTemplateResponse
	GetBody() *GetQualityTemplateResponseBody
}

type GetQualityTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetQualityTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityTemplateResponse) GetBody() *GetQualityTemplateResponseBody {
	return s.Body
}

func (s *GetQualityTemplateResponse) SetHeaders(v map[string]*string) *GetQualityTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetQualityTemplateResponse) SetStatusCode(v int32) *GetQualityTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityTemplateResponse) SetBody(v *GetQualityTemplateResponseBody) *GetQualityTemplateResponse {
	s.Body = v
	return s
}

func (s *GetQualityTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
