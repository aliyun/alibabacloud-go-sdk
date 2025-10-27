// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStrategyTemplateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStrategyTemplateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStrategyTemplateDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetStrategyTemplateDetailResponseBody) *GetStrategyTemplateDetailResponse
	GetBody() *GetStrategyTemplateDetailResponseBody
}

type GetStrategyTemplateDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStrategyTemplateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStrategyTemplateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStrategyTemplateDetailResponse) GoString() string {
	return s.String()
}

func (s *GetStrategyTemplateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStrategyTemplateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStrategyTemplateDetailResponse) GetBody() *GetStrategyTemplateDetailResponseBody {
	return s.Body
}

func (s *GetStrategyTemplateDetailResponse) SetHeaders(v map[string]*string) *GetStrategyTemplateDetailResponse {
	s.Headers = v
	return s
}

func (s *GetStrategyTemplateDetailResponse) SetStatusCode(v int32) *GetStrategyTemplateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStrategyTemplateDetailResponse) SetBody(v *GetStrategyTemplateDetailResponseBody) *GetStrategyTemplateDetailResponse {
	s.Body = v
	return s
}

func (s *GetStrategyTemplateDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
