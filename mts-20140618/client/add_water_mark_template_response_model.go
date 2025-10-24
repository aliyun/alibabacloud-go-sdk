// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWaterMarkTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWaterMarkTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWaterMarkTemplateResponse
	GetStatusCode() *int32
	SetBody(v *AddWaterMarkTemplateResponseBody) *AddWaterMarkTemplateResponse
	GetBody() *AddWaterMarkTemplateResponseBody
}

type AddWaterMarkTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWaterMarkTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWaterMarkTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWaterMarkTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddWaterMarkTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWaterMarkTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWaterMarkTemplateResponse) GetBody() *AddWaterMarkTemplateResponseBody {
	return s.Body
}

func (s *AddWaterMarkTemplateResponse) SetHeaders(v map[string]*string) *AddWaterMarkTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddWaterMarkTemplateResponse) SetStatusCode(v int32) *AddWaterMarkTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWaterMarkTemplateResponse) SetBody(v *AddWaterMarkTemplateResponseBody) *AddWaterMarkTemplateResponse {
	s.Body = v
	return s
}

func (s *AddWaterMarkTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
