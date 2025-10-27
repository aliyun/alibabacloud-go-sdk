// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRouteStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRouteStrategyResponse
	GetStatusCode() *int32
	SetBody(v *CreateRouteStrategyResponseBody) *CreateRouteStrategyResponse
	GetBody() *CreateRouteStrategyResponseBody
}

type CreateRouteStrategyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRouteStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRouteStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateRouteStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRouteStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRouteStrategyResponse) GetBody() *CreateRouteStrategyResponseBody {
	return s.Body
}

func (s *CreateRouteStrategyResponse) SetHeaders(v map[string]*string) *CreateRouteStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateRouteStrategyResponse) SetStatusCode(v int32) *CreateRouteStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouteStrategyResponse) SetBody(v *CreateRouteStrategyResponseBody) *CreateRouteStrategyResponse {
	s.Body = v
	return s
}

func (s *CreateRouteStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
