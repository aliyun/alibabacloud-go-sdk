// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDynamicTagRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDynamicTagRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDynamicTagRuleListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDynamicTagRuleListResponseBody) *DescribeDynamicTagRuleListResponse
	GetBody() *DescribeDynamicTagRuleListResponseBody
}

type DescribeDynamicTagRuleListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDynamicTagRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDynamicTagRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDynamicTagRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDynamicTagRuleListResponse) GetBody() *DescribeDynamicTagRuleListResponseBody {
	return s.Body
}

func (s *DescribeDynamicTagRuleListResponse) SetHeaders(v map[string]*string) *DescribeDynamicTagRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDynamicTagRuleListResponse) SetStatusCode(v int32) *DescribeDynamicTagRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponse) SetBody(v *DescribeDynamicTagRuleListResponseBody) *DescribeDynamicTagRuleListResponse {
	s.Body = v
	return s
}

func (s *DescribeDynamicTagRuleListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
