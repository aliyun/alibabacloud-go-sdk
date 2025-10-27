// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAntiBruteForceRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAntiBruteForceRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAntiBruteForceRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAntiBruteForceRuleResponseBody) *DeleteAntiBruteForceRuleResponse
	GetBody() *DeleteAntiBruteForceRuleResponseBody
}

type DeleteAntiBruteForceRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAntiBruteForceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAntiBruteForceRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAntiBruteForceRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAntiBruteForceRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAntiBruteForceRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAntiBruteForceRuleResponse) GetBody() *DeleteAntiBruteForceRuleResponseBody {
	return s.Body
}

func (s *DeleteAntiBruteForceRuleResponse) SetHeaders(v map[string]*string) *DeleteAntiBruteForceRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAntiBruteForceRuleResponse) SetStatusCode(v int32) *DeleteAntiBruteForceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAntiBruteForceRuleResponse) SetBody(v *DeleteAntiBruteForceRuleResponseBody) *DeleteAntiBruteForceRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteAntiBruteForceRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
