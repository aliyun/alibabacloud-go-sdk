// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemplateParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemplateParamsResponse
	GetStatusCode() *int32
	SetBody(v *GetTemplateParamsResponseBody) *GetTemplateParamsResponse
	GetBody() *GetTemplateParamsResponseBody
}

type GetTemplateParamsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParamsResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemplateParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemplateParamsResponse) GetBody() *GetTemplateParamsResponseBody {
	return s.Body
}

func (s *GetTemplateParamsResponse) SetHeaders(v map[string]*string) *GetTemplateParamsResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateParamsResponse) SetStatusCode(v int32) *GetTemplateParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateParamsResponse) SetBody(v *GetTemplateParamsResponseBody) *GetTemplateParamsResponse {
	s.Body = v
	return s
}

func (s *GetTemplateParamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
