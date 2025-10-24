// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWaterMarkTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryWaterMarkTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryWaterMarkTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *QueryWaterMarkTemplateListResponseBody) *QueryWaterMarkTemplateListResponse
	GetBody() *QueryWaterMarkTemplateListResponseBody
}

type QueryWaterMarkTemplateListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWaterMarkTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWaterMarkTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryWaterMarkTemplateListResponse) GoString() string {
	return s.String()
}

func (s *QueryWaterMarkTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryWaterMarkTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryWaterMarkTemplateListResponse) GetBody() *QueryWaterMarkTemplateListResponseBody {
	return s.Body
}

func (s *QueryWaterMarkTemplateListResponse) SetHeaders(v map[string]*string) *QueryWaterMarkTemplateListResponse {
	s.Headers = v
	return s
}

func (s *QueryWaterMarkTemplateListResponse) SetStatusCode(v int32) *QueryWaterMarkTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponse) SetBody(v *QueryWaterMarkTemplateListResponseBody) *QueryWaterMarkTemplateListResponse {
	s.Body = v
	return s
}

func (s *QueryWaterMarkTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
