// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCheckTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCheckTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCheckTemplateResponseBody) *GetDataCheckTemplateResponse
	GetBody() *GetDataCheckTemplateResponseBody
}

type GetDataCheckTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCheckTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCheckTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCheckTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCheckTemplateResponse) GetBody() *GetDataCheckTemplateResponseBody {
	return s.Body
}

func (s *GetDataCheckTemplateResponse) SetHeaders(v map[string]*string) *GetDataCheckTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetDataCheckTemplateResponse) SetStatusCode(v int32) *GetDataCheckTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCheckTemplateResponse) SetBody(v *GetDataCheckTemplateResponseBody) *GetDataCheckTemplateResponse {
	s.Body = v
	return s
}

func (s *GetDataCheckTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
