// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityRuleListResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityRuleListResponseBody) *GetQualityRuleListResponse
	GetBody() *GetQualityRuleListResponseBody
}

type GetQualityRuleListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleListResponse) GoString() string {
	return s.String()
}

func (s *GetQualityRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityRuleListResponse) GetBody() *GetQualityRuleListResponseBody {
	return s.Body
}

func (s *GetQualityRuleListResponse) SetHeaders(v map[string]*string) *GetQualityRuleListResponse {
	s.Headers = v
	return s
}

func (s *GetQualityRuleListResponse) SetStatusCode(v int32) *GetQualityRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityRuleListResponse) SetBody(v *GetQualityRuleListResponseBody) *GetQualityRuleListResponse {
	s.Body = v
	return s
}

func (s *GetQualityRuleListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
