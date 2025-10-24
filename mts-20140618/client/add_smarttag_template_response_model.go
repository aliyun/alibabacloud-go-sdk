// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSmarttagTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSmarttagTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSmarttagTemplateResponse
	GetStatusCode() *int32
	SetBody(v *AddSmarttagTemplateResponseBody) *AddSmarttagTemplateResponse
	GetBody() *AddSmarttagTemplateResponseBody
}

type AddSmarttagTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSmarttagTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSmarttagTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSmarttagTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddSmarttagTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSmarttagTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSmarttagTemplateResponse) GetBody() *AddSmarttagTemplateResponseBody {
	return s.Body
}

func (s *AddSmarttagTemplateResponse) SetHeaders(v map[string]*string) *AddSmarttagTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddSmarttagTemplateResponse) SetStatusCode(v int32) *AddSmarttagTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSmarttagTemplateResponse) SetBody(v *AddSmarttagTemplateResponseBody) *AddSmarttagTemplateResponse {
	s.Body = v
	return s
}

func (s *AddSmarttagTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
