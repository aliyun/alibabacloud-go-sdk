// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveContactTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveContactTemplateResponse
	GetStatusCode() *int32
	SetBody(v *SaveContactTemplateResponseBody) *SaveContactTemplateResponse
	GetBody() *SaveContactTemplateResponseBody
}

type SaveContactTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveContactTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveContactTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveContactTemplateResponse) GoString() string {
	return s.String()
}

func (s *SaveContactTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveContactTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveContactTemplateResponse) GetBody() *SaveContactTemplateResponseBody {
	return s.Body
}

func (s *SaveContactTemplateResponse) SetHeaders(v map[string]*string) *SaveContactTemplateResponse {
	s.Headers = v
	return s
}

func (s *SaveContactTemplateResponse) SetStatusCode(v int32) *SaveContactTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveContactTemplateResponse) SetBody(v *SaveContactTemplateResponseBody) *SaveContactTemplateResponse {
	s.Body = v
	return s
}

func (s *SaveContactTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
