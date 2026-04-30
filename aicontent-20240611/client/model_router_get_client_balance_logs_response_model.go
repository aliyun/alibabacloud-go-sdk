// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetClientBalanceLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterGetClientBalanceLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterGetClientBalanceLogsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterGetClientBalanceLogsResponseBody) *ModelRouterGetClientBalanceLogsResponse
	GetBody() *ModelRouterGetClientBalanceLogsResponseBody
}

type ModelRouterGetClientBalanceLogsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterGetClientBalanceLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterGetClientBalanceLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetClientBalanceLogsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterGetClientBalanceLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterGetClientBalanceLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterGetClientBalanceLogsResponse) GetBody() *ModelRouterGetClientBalanceLogsResponseBody {
	return s.Body
}

func (s *ModelRouterGetClientBalanceLogsResponse) SetHeaders(v map[string]*string) *ModelRouterGetClientBalanceLogsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponse) SetStatusCode(v int32) *ModelRouterGetClientBalanceLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponse) SetBody(v *ModelRouterGetClientBalanceLogsResponseBody) *ModelRouterGetClientBalanceLogsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
