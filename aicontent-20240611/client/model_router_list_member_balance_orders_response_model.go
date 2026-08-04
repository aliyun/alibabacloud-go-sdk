// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListMemberBalanceOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterListMemberBalanceOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterListMemberBalanceOrdersResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterListMemberBalanceOrdersResponseBody) *ModelRouterListMemberBalanceOrdersResponse
	GetBody() *ModelRouterListMemberBalanceOrdersResponseBody
}

type ModelRouterListMemberBalanceOrdersResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterListMemberBalanceOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterListMemberBalanceOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListMemberBalanceOrdersResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterListMemberBalanceOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterListMemberBalanceOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterListMemberBalanceOrdersResponse) GetBody() *ModelRouterListMemberBalanceOrdersResponseBody {
	return s.Body
}

func (s *ModelRouterListMemberBalanceOrdersResponse) SetHeaders(v map[string]*string) *ModelRouterListMemberBalanceOrdersResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponse) SetStatusCode(v int32) *ModelRouterListMemberBalanceOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponse) SetBody(v *ModelRouterListMemberBalanceOrdersResponseBody) *ModelRouterListMemberBalanceOrdersResponse {
	s.Body = v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
