// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateConfigRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAggregateConfigRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAggregateConfigRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAggregateConfigRulesResponseBody) *DeleteAggregateConfigRulesResponse
	GetBody() *DeleteAggregateConfigRulesResponseBody
}

type DeleteAggregateConfigRulesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAggregateConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAggregateConfigRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAggregateConfigRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAggregateConfigRulesResponse) GetBody() *DeleteAggregateConfigRulesResponseBody {
	return s.Body
}

func (s *DeleteAggregateConfigRulesResponse) SetHeaders(v map[string]*string) *DeleteAggregateConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAggregateConfigRulesResponse) SetStatusCode(v int32) *DeleteAggregateConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAggregateConfigRulesResponse) SetBody(v *DeleteAggregateConfigRulesResponseBody) *DeleteAggregateConfigRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteAggregateConfigRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
