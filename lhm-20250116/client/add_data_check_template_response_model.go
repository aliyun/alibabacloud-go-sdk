// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataCheckTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDataCheckTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDataCheckTemplateResponse
	GetStatusCode() *int32
	SetBody(v *AddDataCheckTemplateResponseBody) *AddDataCheckTemplateResponse
	GetBody() *AddDataCheckTemplateResponseBody
}

type AddDataCheckTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataCheckTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDataCheckTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddDataCheckTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDataCheckTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDataCheckTemplateResponse) GetBody() *AddDataCheckTemplateResponseBody {
	return s.Body
}

func (s *AddDataCheckTemplateResponse) SetHeaders(v map[string]*string) *AddDataCheckTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddDataCheckTemplateResponse) SetStatusCode(v int32) *AddDataCheckTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataCheckTemplateResponse) SetBody(v *AddDataCheckTemplateResponseBody) *AddDataCheckTemplateResponse {
	s.Body = v
	return s
}

func (s *AddDataCheckTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
