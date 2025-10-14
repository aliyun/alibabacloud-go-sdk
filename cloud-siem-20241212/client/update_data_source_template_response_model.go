// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataSourceTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataSourceTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataSourceTemplateResponseBody) *UpdateDataSourceTemplateResponse
	GetBody() *UpdateDataSourceTemplateResponseBody
}

type UpdateDataSourceTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataSourceTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataSourceTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataSourceTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataSourceTemplateResponse) GetBody() *UpdateDataSourceTemplateResponseBody {
	return s.Body
}

func (s *UpdateDataSourceTemplateResponse) SetHeaders(v map[string]*string) *UpdateDataSourceTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataSourceTemplateResponse) SetStatusCode(v int32) *UpdateDataSourceTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataSourceTemplateResponse) SetBody(v *UpdateDataSourceTemplateResponseBody) *UpdateDataSourceTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateDataSourceTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
