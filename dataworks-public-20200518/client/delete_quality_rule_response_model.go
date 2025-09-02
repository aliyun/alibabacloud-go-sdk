// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQualityRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQualityRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQualityRuleResponseBody) *DeleteQualityRuleResponse
	GetBody() *DeleteQualityRuleResponseBody
}

type DeleteQualityRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQualityRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualityRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQualityRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQualityRuleResponse) GetBody() *DeleteQualityRuleResponseBody {
	return s.Body
}

func (s *DeleteQualityRuleResponse) SetHeaders(v map[string]*string) *DeleteQualityRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualityRuleResponse) SetStatusCode(v int32) *DeleteQualityRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQualityRuleResponse) SetBody(v *DeleteQualityRuleResponseBody) *DeleteQualityRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteQualityRuleResponse) Validate() error {
	return dara.Validate(s)
}
