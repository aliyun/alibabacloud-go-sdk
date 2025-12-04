// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpdGrantRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpdGrantRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpdGrantRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpdGrantRuleResponseBody) *DeleteVpdGrantRuleResponse
	GetBody() *DeleteVpdGrantRuleResponseBody
}

type DeleteVpdGrantRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpdGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpdGrantRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpdGrantRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpdGrantRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpdGrantRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpdGrantRuleResponse) GetBody() *DeleteVpdGrantRuleResponseBody {
	return s.Body
}

func (s *DeleteVpdGrantRuleResponse) SetHeaders(v map[string]*string) *DeleteVpdGrantRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpdGrantRuleResponse) SetStatusCode(v int32) *DeleteVpdGrantRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpdGrantRuleResponse) SetBody(v *DeleteVpdGrantRuleResponseBody) *DeleteVpdGrantRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteVpdGrantRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
