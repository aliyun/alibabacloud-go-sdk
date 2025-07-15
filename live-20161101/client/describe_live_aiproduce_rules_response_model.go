// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAIProduceRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveAIProduceRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveAIProduceRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveAIProduceRulesResponseBody) *DescribeLiveAIProduceRulesResponse
	GetBody() *DescribeLiveAIProduceRulesResponseBody
}

type DescribeLiveAIProduceRulesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveAIProduceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveAIProduceRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAIProduceRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveAIProduceRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveAIProduceRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveAIProduceRulesResponse) GetBody() *DescribeLiveAIProduceRulesResponseBody {
	return s.Body
}

func (s *DescribeLiveAIProduceRulesResponse) SetHeaders(v map[string]*string) *DescribeLiveAIProduceRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveAIProduceRulesResponse) SetStatusCode(v int32) *DescribeLiveAIProduceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveAIProduceRulesResponse) SetBody(v *DescribeLiveAIProduceRulesResponseBody) *DescribeLiveAIProduceRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveAIProduceRulesResponse) Validate() error {
	return dara.Validate(s)
}
