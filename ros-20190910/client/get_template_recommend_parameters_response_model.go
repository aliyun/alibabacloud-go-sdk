// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateRecommendParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemplateRecommendParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemplateRecommendParametersResponse
	GetStatusCode() *int32
	SetBody(v *GetTemplateRecommendParametersResponseBody) *GetTemplateRecommendParametersResponse
	GetBody() *GetTemplateRecommendParametersResponseBody
}

type GetTemplateRecommendParametersResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateRecommendParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateRecommendParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRecommendParametersResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateRecommendParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemplateRecommendParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemplateRecommendParametersResponse) GetBody() *GetTemplateRecommendParametersResponseBody {
	return s.Body
}

func (s *GetTemplateRecommendParametersResponse) SetHeaders(v map[string]*string) *GetTemplateRecommendParametersResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateRecommendParametersResponse) SetStatusCode(v int32) *GetTemplateRecommendParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateRecommendParametersResponse) SetBody(v *GetTemplateRecommendParametersResponseBody) *GetTemplateRecommendParametersResponse {
	s.Body = v
	return s
}

func (s *GetTemplateRecommendParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
