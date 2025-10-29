// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataQualityTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataQualityTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataQualityTemplateResponseBody) *DeleteDataQualityTemplateResponse
	GetBody() *DeleteDataQualityTemplateResponseBody
}

type DeleteDataQualityTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataQualityTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataQualityTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataQualityTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataQualityTemplateResponse) GetBody() *DeleteDataQualityTemplateResponseBody {
	return s.Body
}

func (s *DeleteDataQualityTemplateResponse) SetHeaders(v map[string]*string) *DeleteDataQualityTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataQualityTemplateResponse) SetStatusCode(v int32) *DeleteDataQualityTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataQualityTemplateResponse) SetBody(v *DeleteDataQualityTemplateResponseBody) *DeleteDataQualityTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteDataQualityTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
