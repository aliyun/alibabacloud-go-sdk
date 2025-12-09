// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBotRuleLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBotRuleLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBotRuleLabelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBotRuleLabelsResponseBody) *DescribeBotRuleLabelsResponse
	GetBody() *DescribeBotRuleLabelsResponseBody
}

type DescribeBotRuleLabelsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBotRuleLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBotRuleLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBotRuleLabelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBotRuleLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBotRuleLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBotRuleLabelsResponse) GetBody() *DescribeBotRuleLabelsResponseBody {
	return s.Body
}

func (s *DescribeBotRuleLabelsResponse) SetHeaders(v map[string]*string) *DescribeBotRuleLabelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBotRuleLabelsResponse) SetStatusCode(v int32) *DescribeBotRuleLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBotRuleLabelsResponse) SetBody(v *DescribeBotRuleLabelsResponseBody) *DescribeBotRuleLabelsResponse {
	s.Body = v
	return s
}

func (s *DescribeBotRuleLabelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
