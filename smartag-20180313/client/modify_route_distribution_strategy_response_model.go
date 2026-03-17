// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouteDistributionStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRouteDistributionStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRouteDistributionStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRouteDistributionStrategyResponseBody) *ModifyRouteDistributionStrategyResponse
	GetBody() *ModifyRouteDistributionStrategyResponseBody
}

type ModifyRouteDistributionStrategyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRouteDistributionStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRouteDistributionStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouteDistributionStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyRouteDistributionStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRouteDistributionStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRouteDistributionStrategyResponse) GetBody() *ModifyRouteDistributionStrategyResponseBody {
	return s.Body
}

func (s *ModifyRouteDistributionStrategyResponse) SetHeaders(v map[string]*string) *ModifyRouteDistributionStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyRouteDistributionStrategyResponse) SetStatusCode(v int32) *ModifyRouteDistributionStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRouteDistributionStrategyResponse) SetBody(v *ModifyRouteDistributionStrategyResponseBody) *ModifyRouteDistributionStrategyResponse {
	s.Body = v
	return s
}

func (s *ModifyRouteDistributionStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
