// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaterMarkTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWaterMarkTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWaterMarkTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWaterMarkTemplateResponseBody) *DeleteWaterMarkTemplateResponse
	GetBody() *DeleteWaterMarkTemplateResponseBody
}

type DeleteWaterMarkTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWaterMarkTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWaterMarkTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaterMarkTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteWaterMarkTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWaterMarkTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWaterMarkTemplateResponse) GetBody() *DeleteWaterMarkTemplateResponseBody {
	return s.Body
}

func (s *DeleteWaterMarkTemplateResponse) SetHeaders(v map[string]*string) *DeleteWaterMarkTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteWaterMarkTemplateResponse) SetStatusCode(v int32) *DeleteWaterMarkTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWaterMarkTemplateResponse) SetBody(v *DeleteWaterMarkTemplateResponseBody) *DeleteWaterMarkTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteWaterMarkTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
