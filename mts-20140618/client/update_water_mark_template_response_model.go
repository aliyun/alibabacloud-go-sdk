// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaterMarkTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWaterMarkTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWaterMarkTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWaterMarkTemplateResponseBody) *UpdateWaterMarkTemplateResponse
	GetBody() *UpdateWaterMarkTemplateResponseBody
}

type UpdateWaterMarkTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWaterMarkTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWaterMarkTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaterMarkTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateWaterMarkTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWaterMarkTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWaterMarkTemplateResponse) GetBody() *UpdateWaterMarkTemplateResponseBody {
	return s.Body
}

func (s *UpdateWaterMarkTemplateResponse) SetHeaders(v map[string]*string) *UpdateWaterMarkTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateWaterMarkTemplateResponse) SetStatusCode(v int32) *UpdateWaterMarkTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponse) SetBody(v *UpdateWaterMarkTemplateResponseBody) *UpdateWaterMarkTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateWaterMarkTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
