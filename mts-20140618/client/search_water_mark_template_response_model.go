// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchWaterMarkTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchWaterMarkTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchWaterMarkTemplateResponse
	GetStatusCode() *int32
	SetBody(v *SearchWaterMarkTemplateResponseBody) *SearchWaterMarkTemplateResponse
	GetBody() *SearchWaterMarkTemplateResponseBody
}

type SearchWaterMarkTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchWaterMarkTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchWaterMarkTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchWaterMarkTemplateResponse) GoString() string {
	return s.String()
}

func (s *SearchWaterMarkTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchWaterMarkTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchWaterMarkTemplateResponse) GetBody() *SearchWaterMarkTemplateResponseBody {
	return s.Body
}

func (s *SearchWaterMarkTemplateResponse) SetHeaders(v map[string]*string) *SearchWaterMarkTemplateResponse {
	s.Headers = v
	return s
}

func (s *SearchWaterMarkTemplateResponse) SetStatusCode(v int32) *SearchWaterMarkTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchWaterMarkTemplateResponse) SetBody(v *SearchWaterMarkTemplateResponseBody) *SearchWaterMarkTemplateResponse {
	s.Body = v
	return s
}

func (s *SearchWaterMarkTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
