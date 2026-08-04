// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetMemberBalanceLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterGetMemberBalanceLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterGetMemberBalanceLogsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterGetMemberBalanceLogsResponseBody) *ModelRouterGetMemberBalanceLogsResponse
	GetBody() *ModelRouterGetMemberBalanceLogsResponseBody
}

type ModelRouterGetMemberBalanceLogsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterGetMemberBalanceLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterGetMemberBalanceLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetMemberBalanceLogsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterGetMemberBalanceLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterGetMemberBalanceLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterGetMemberBalanceLogsResponse) GetBody() *ModelRouterGetMemberBalanceLogsResponseBody {
	return s.Body
}

func (s *ModelRouterGetMemberBalanceLogsResponse) SetHeaders(v map[string]*string) *ModelRouterGetMemberBalanceLogsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponse) SetStatusCode(v int32) *ModelRouterGetMemberBalanceLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponse) SetBody(v *ModelRouterGetMemberBalanceLogsResponseBody) *ModelRouterGetMemberBalanceLogsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
