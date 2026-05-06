// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelProviderTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelProviderTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelProviderTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetModelProviderTemplateResponseBody) *GetModelProviderTemplateResponse
	GetBody() *GetModelProviderTemplateResponseBody
}

type GetModelProviderTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelProviderTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelProviderTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelProviderTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetModelProviderTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelProviderTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelProviderTemplateResponse) GetBody() *GetModelProviderTemplateResponseBody {
	return s.Body
}

func (s *GetModelProviderTemplateResponse) SetHeaders(v map[string]*string) *GetModelProviderTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetModelProviderTemplateResponse) SetStatusCode(v int32) *GetModelProviderTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelProviderTemplateResponse) SetBody(v *GetModelProviderTemplateResponseBody) *GetModelProviderTemplateResponse {
	s.Body = v
	return s
}

func (s *GetModelProviderTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
