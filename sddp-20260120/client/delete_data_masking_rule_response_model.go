// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataMaskingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataMaskingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataMaskingRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataMaskingRuleResponseBody) *DeleteDataMaskingRuleResponse
	GetBody() *DeleteDataMaskingRuleResponseBody
}

type DeleteDataMaskingRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataMaskingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataMaskingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataMaskingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataMaskingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataMaskingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataMaskingRuleResponse) GetBody() *DeleteDataMaskingRuleResponseBody {
	return s.Body
}

func (s *DeleteDataMaskingRuleResponse) SetHeaders(v map[string]*string) *DeleteDataMaskingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataMaskingRuleResponse) SetStatusCode(v int32) *DeleteDataMaskingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataMaskingRuleResponse) SetBody(v *DeleteDataMaskingRuleResponseBody) *DeleteDataMaskingRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteDataMaskingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
