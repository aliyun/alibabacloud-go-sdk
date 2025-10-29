// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataQualityRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataQualityRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataQualityRuleResponseBody) *DeleteDataQualityRuleResponse
	GetBody() *DeleteDataQualityRuleResponseBody
}

type DeleteDataQualityRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataQualityRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataQualityRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataQualityRuleResponse) GetBody() *DeleteDataQualityRuleResponseBody {
	return s.Body
}

func (s *DeleteDataQualityRuleResponse) SetHeaders(v map[string]*string) *DeleteDataQualityRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataQualityRuleResponse) SetStatusCode(v int32) *DeleteDataQualityRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataQualityRuleResponse) SetBody(v *DeleteDataQualityRuleResponseBody) *DeleteDataQualityRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteDataQualityRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
