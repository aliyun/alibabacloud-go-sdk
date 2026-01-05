// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSmsTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSmsTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *GetSmsTemplateListResponseBody) *GetSmsTemplateListResponse
	GetBody() *GetSmsTemplateListResponseBody
}

type GetSmsTemplateListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmsTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmsTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateListResponse) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSmsTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSmsTemplateListResponse) GetBody() *GetSmsTemplateListResponseBody {
	return s.Body
}

func (s *GetSmsTemplateListResponse) SetHeaders(v map[string]*string) *GetSmsTemplateListResponse {
	s.Headers = v
	return s
}

func (s *GetSmsTemplateListResponse) SetStatusCode(v int32) *GetSmsTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmsTemplateListResponse) SetBody(v *GetSmsTemplateListResponseBody) *GetSmsTemplateListResponse {
	s.Body = v
	return s
}

func (s *GetSmsTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
