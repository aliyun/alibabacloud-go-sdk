// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteDistributionStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouteDistributionStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouteDistributionStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouteDistributionStrategyResponseBody) *DeleteRouteDistributionStrategyResponse
	GetBody() *DeleteRouteDistributionStrategyResponseBody
}

type DeleteRouteDistributionStrategyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteDistributionStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteDistributionStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteDistributionStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteDistributionStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouteDistributionStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouteDistributionStrategyResponse) GetBody() *DeleteRouteDistributionStrategyResponseBody {
	return s.Body
}

func (s *DeleteRouteDistributionStrategyResponse) SetHeaders(v map[string]*string) *DeleteRouteDistributionStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteDistributionStrategyResponse) SetStatusCode(v int32) *DeleteRouteDistributionStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteDistributionStrategyResponse) SetBody(v *DeleteRouteDistributionStrategyResponseBody) *DeleteRouteDistributionStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteRouteDistributionStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
