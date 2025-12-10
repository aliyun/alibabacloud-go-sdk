// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoGroupingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoGroupingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoGroupingRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoGroupingRuleResponseBody) *GetAutoGroupingRuleResponse
	GetBody() *GetAutoGroupingRuleResponseBody
}

type GetAutoGroupingRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoGroupingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoGroupingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoGroupingRuleResponse) GoString() string {
	return s.String()
}

func (s *GetAutoGroupingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoGroupingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoGroupingRuleResponse) GetBody() *GetAutoGroupingRuleResponseBody {
	return s.Body
}

func (s *GetAutoGroupingRuleResponse) SetHeaders(v map[string]*string) *GetAutoGroupingRuleResponse {
	s.Headers = v
	return s
}

func (s *GetAutoGroupingRuleResponse) SetStatusCode(v int32) *GetAutoGroupingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoGroupingRuleResponse) SetBody(v *GetAutoGroupingRuleResponseBody) *GetAutoGroupingRuleResponse {
	s.Body = v
	return s
}

func (s *GetAutoGroupingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
