// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVccGrantRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVccGrantRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVccGrantRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVccGrantRuleResponseBody) *DeleteVccGrantRuleResponse
	GetBody() *DeleteVccGrantRuleResponseBody
}

type DeleteVccGrantRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVccGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVccGrantRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVccGrantRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteVccGrantRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVccGrantRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVccGrantRuleResponse) GetBody() *DeleteVccGrantRuleResponseBody {
	return s.Body
}

func (s *DeleteVccGrantRuleResponse) SetHeaders(v map[string]*string) *DeleteVccGrantRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteVccGrantRuleResponse) SetStatusCode(v int32) *DeleteVccGrantRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVccGrantRuleResponse) SetBody(v *DeleteVccGrantRuleResponseBody) *DeleteVccGrantRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteVccGrantRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
