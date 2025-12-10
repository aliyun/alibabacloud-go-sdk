// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoGroupingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAutoGroupingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAutoGroupingRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAutoGroupingRuleResponseBody) *DeleteAutoGroupingRuleResponse
	GetBody() *DeleteAutoGroupingRuleResponseBody
}

type DeleteAutoGroupingRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutoGroupingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutoGroupingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoGroupingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoGroupingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAutoGroupingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAutoGroupingRuleResponse) GetBody() *DeleteAutoGroupingRuleResponseBody {
	return s.Body
}

func (s *DeleteAutoGroupingRuleResponse) SetHeaders(v map[string]*string) *DeleteAutoGroupingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoGroupingRuleResponse) SetStatusCode(v int32) *DeleteAutoGroupingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoGroupingRuleResponse) SetBody(v *DeleteAutoGroupingRuleResponseBody) *DeleteAutoGroupingRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteAutoGroupingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
