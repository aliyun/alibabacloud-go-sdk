// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvancedQueryTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAdvancedQueryTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAdvancedQueryTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetAdvancedQueryTemplateResponseBody) *GetAdvancedQueryTemplateResponse
	GetBody() *GetAdvancedQueryTemplateResponseBody
}

type GetAdvancedQueryTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdvancedQueryTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdvancedQueryTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAdvancedQueryTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetAdvancedQueryTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAdvancedQueryTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAdvancedQueryTemplateResponse) GetBody() *GetAdvancedQueryTemplateResponseBody {
	return s.Body
}

func (s *GetAdvancedQueryTemplateResponse) SetHeaders(v map[string]*string) *GetAdvancedQueryTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetAdvancedQueryTemplateResponse) SetStatusCode(v int32) *GetAdvancedQueryTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdvancedQueryTemplateResponse) SetBody(v *GetAdvancedQueryTemplateResponseBody) *GetAdvancedQueryTemplateResponse {
	s.Body = v
	return s
}

func (s *GetAdvancedQueryTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
