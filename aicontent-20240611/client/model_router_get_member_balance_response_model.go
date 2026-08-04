// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetMemberBalanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterGetMemberBalanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterGetMemberBalanceResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterGetMemberBalanceResponseBody) *ModelRouterGetMemberBalanceResponse
	GetBody() *ModelRouterGetMemberBalanceResponseBody
}

type ModelRouterGetMemberBalanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterGetMemberBalanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterGetMemberBalanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetMemberBalanceResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterGetMemberBalanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterGetMemberBalanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterGetMemberBalanceResponse) GetBody() *ModelRouterGetMemberBalanceResponseBody {
	return s.Body
}

func (s *ModelRouterGetMemberBalanceResponse) SetHeaders(v map[string]*string) *ModelRouterGetMemberBalanceResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterGetMemberBalanceResponse) SetStatusCode(v int32) *ModelRouterGetMemberBalanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterGetMemberBalanceResponse) SetBody(v *ModelRouterGetMemberBalanceResponseBody) *ModelRouterGetMemberBalanceResponse {
	s.Body = v
	return s
}

func (s *ModelRouterGetMemberBalanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
