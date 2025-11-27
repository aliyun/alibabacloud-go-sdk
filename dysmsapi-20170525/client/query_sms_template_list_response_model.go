// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySmsTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySmsTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *QuerySmsTemplateListResponseBody) *QuerySmsTemplateListResponse
	GetBody() *QuerySmsTemplateListResponseBody
}

type QuerySmsTemplateListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmsTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTemplateListResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySmsTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySmsTemplateListResponse) GetBody() *QuerySmsTemplateListResponseBody {
	return s.Body
}

func (s *QuerySmsTemplateListResponse) SetHeaders(v map[string]*string) *QuerySmsTemplateListResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsTemplateListResponse) SetStatusCode(v int32) *QuerySmsTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsTemplateListResponse) SetBody(v *QuerySmsTemplateListResponseBody) *QuerySmsTemplateListResponse {
	s.Body = v
	return s
}

func (s *QuerySmsTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
