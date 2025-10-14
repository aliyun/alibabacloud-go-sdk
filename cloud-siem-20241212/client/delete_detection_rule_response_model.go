// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDetectionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDetectionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDetectionRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDetectionRuleResponseBody) *DeleteDetectionRuleResponse
	GetBody() *DeleteDetectionRuleResponseBody
}

type DeleteDetectionRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDetectionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDetectionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDetectionRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDetectionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDetectionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDetectionRuleResponse) GetBody() *DeleteDetectionRuleResponseBody {
	return s.Body
}

func (s *DeleteDetectionRuleResponse) SetHeaders(v map[string]*string) *DeleteDetectionRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDetectionRuleResponse) SetStatusCode(v int32) *DeleteDetectionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDetectionRuleResponse) SetBody(v *DeleteDetectionRuleResponseBody) *DeleteDetectionRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteDetectionRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
