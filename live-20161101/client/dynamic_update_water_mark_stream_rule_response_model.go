// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDynamicUpdateWaterMarkStreamRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DynamicUpdateWaterMarkStreamRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DynamicUpdateWaterMarkStreamRuleResponse
	GetStatusCode() *int32
	SetBody(v *DynamicUpdateWaterMarkStreamRuleResponseBody) *DynamicUpdateWaterMarkStreamRuleResponse
	GetBody() *DynamicUpdateWaterMarkStreamRuleResponseBody
}

type DynamicUpdateWaterMarkStreamRuleResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DynamicUpdateWaterMarkStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DynamicUpdateWaterMarkStreamRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DynamicUpdateWaterMarkStreamRuleResponse) GoString() string {
	return s.String()
}

func (s *DynamicUpdateWaterMarkStreamRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DynamicUpdateWaterMarkStreamRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DynamicUpdateWaterMarkStreamRuleResponse) GetBody() *DynamicUpdateWaterMarkStreamRuleResponseBody {
	return s.Body
}

func (s *DynamicUpdateWaterMarkStreamRuleResponse) SetHeaders(v map[string]*string) *DynamicUpdateWaterMarkStreamRuleResponse {
	s.Headers = v
	return s
}

func (s *DynamicUpdateWaterMarkStreamRuleResponse) SetStatusCode(v int32) *DynamicUpdateWaterMarkStreamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DynamicUpdateWaterMarkStreamRuleResponse) SetBody(v *DynamicUpdateWaterMarkStreamRuleResponseBody) *DynamicUpdateWaterMarkStreamRuleResponse {
	s.Body = v
	return s
}

func (s *DynamicUpdateWaterMarkStreamRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
