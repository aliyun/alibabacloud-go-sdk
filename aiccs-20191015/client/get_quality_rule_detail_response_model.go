// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityRuleDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityRuleDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityRuleDetailResponseBody) *GetQualityRuleDetailResponse
	GetBody() *GetQualityRuleDetailResponseBody
}

type GetQualityRuleDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityRuleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityRuleDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleDetailResponse) GoString() string {
	return s.String()
}

func (s *GetQualityRuleDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityRuleDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityRuleDetailResponse) GetBody() *GetQualityRuleDetailResponseBody {
	return s.Body
}

func (s *GetQualityRuleDetailResponse) SetHeaders(v map[string]*string) *GetQualityRuleDetailResponse {
	s.Headers = v
	return s
}

func (s *GetQualityRuleDetailResponse) SetStatusCode(v int32) *GetQualityRuleDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityRuleDetailResponse) SetBody(v *GetQualityRuleDetailResponseBody) *GetQualityRuleDetailResponse {
	s.Body = v
	return s
}

func (s *GetQualityRuleDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
