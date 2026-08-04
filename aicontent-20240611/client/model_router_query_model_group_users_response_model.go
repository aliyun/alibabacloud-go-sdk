// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryModelGroupUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryModelGroupUsersResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryModelGroupUsersResponseBody) *ModelRouterQueryModelGroupUsersResponse
	GetBody() *ModelRouterQueryModelGroupUsersResponseBody
}

type ModelRouterQueryModelGroupUsersResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryModelGroupUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryModelGroupUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupUsersResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryModelGroupUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryModelGroupUsersResponse) GetBody() *ModelRouterQueryModelGroupUsersResponseBody {
	return s.Body
}

func (s *ModelRouterQueryModelGroupUsersResponse) SetHeaders(v map[string]*string) *ModelRouterQueryModelGroupUsersResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponse) SetStatusCode(v int32) *ModelRouterQueryModelGroupUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponse) SetBody(v *ModelRouterQueryModelGroupUsersResponseBody) *ModelRouterQueryModelGroupUsersResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
