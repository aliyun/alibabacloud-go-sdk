// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecommendTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecommendTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecommendTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *RecommendTemplatesResponseBody) *RecommendTemplatesResponse
	GetBody() *RecommendTemplatesResponseBody
}

type RecommendTemplatesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecommendTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecommendTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s RecommendTemplatesResponse) GoString() string {
	return s.String()
}

func (s *RecommendTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecommendTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecommendTemplatesResponse) GetBody() *RecommendTemplatesResponseBody {
	return s.Body
}

func (s *RecommendTemplatesResponse) SetHeaders(v map[string]*string) *RecommendTemplatesResponse {
	s.Headers = v
	return s
}

func (s *RecommendTemplatesResponse) SetStatusCode(v int32) *RecommendTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *RecommendTemplatesResponse) SetBody(v *RecommendTemplatesResponseBody) *RecommendTemplatesResponse {
	s.Body = v
	return s
}

func (s *RecommendTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
