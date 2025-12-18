// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRemediationTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRemediationTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRemediationTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetRemediationTemplateResponseBody) *GetRemediationTemplateResponse
	GetBody() *GetRemediationTemplateResponseBody
}

type GetRemediationTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRemediationTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRemediationTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRemediationTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetRemediationTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRemediationTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRemediationTemplateResponse) GetBody() *GetRemediationTemplateResponseBody {
	return s.Body
}

func (s *GetRemediationTemplateResponse) SetHeaders(v map[string]*string) *GetRemediationTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetRemediationTemplateResponse) SetStatusCode(v int32) *GetRemediationTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRemediationTemplateResponse) SetBody(v *GetRemediationTemplateResponseBody) *GetRemediationTemplateResponse {
	s.Body = v
	return s
}

func (s *GetRemediationTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
