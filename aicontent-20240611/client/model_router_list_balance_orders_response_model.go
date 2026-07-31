// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListBalanceOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterListBalanceOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterListBalanceOrdersResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterListBalanceOrdersResponseBody) *ModelRouterListBalanceOrdersResponse
	GetBody() *ModelRouterListBalanceOrdersResponseBody
}

type ModelRouterListBalanceOrdersResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterListBalanceOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterListBalanceOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListBalanceOrdersResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterListBalanceOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterListBalanceOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterListBalanceOrdersResponse) GetBody() *ModelRouterListBalanceOrdersResponseBody {
	return s.Body
}

func (s *ModelRouterListBalanceOrdersResponse) SetHeaders(v map[string]*string) *ModelRouterListBalanceOrdersResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterListBalanceOrdersResponse) SetStatusCode(v int32) *ModelRouterListBalanceOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterListBalanceOrdersResponse) SetBody(v *ModelRouterListBalanceOrdersResponseBody) *ModelRouterListBalanceOrdersResponse {
	s.Body = v
	return s
}

func (s *ModelRouterListBalanceOrdersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
