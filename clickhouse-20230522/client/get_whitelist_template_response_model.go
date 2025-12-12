// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhitelistTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWhitelistTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWhitelistTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetWhitelistTemplateResponseBody) *GetWhitelistTemplateResponse
	GetBody() *GetWhitelistTemplateResponseBody
}

type GetWhitelistTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWhitelistTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWhitelistTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWhitelistTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetWhitelistTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWhitelistTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWhitelistTemplateResponse) GetBody() *GetWhitelistTemplateResponseBody {
	return s.Body
}

func (s *GetWhitelistTemplateResponse) SetHeaders(v map[string]*string) *GetWhitelistTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetWhitelistTemplateResponse) SetStatusCode(v int32) *GetWhitelistTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWhitelistTemplateResponse) SetBody(v *GetWhitelistTemplateResponseBody) *GetWhitelistTemplateResponse {
	s.Body = v
	return s
}

func (s *GetWhitelistTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
