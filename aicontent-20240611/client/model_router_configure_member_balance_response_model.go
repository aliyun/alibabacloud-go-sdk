// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterConfigureMemberBalanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterConfigureMemberBalanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterConfigureMemberBalanceResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterConfigureMemberBalanceResponseBody) *ModelRouterConfigureMemberBalanceResponse
	GetBody() *ModelRouterConfigureMemberBalanceResponseBody
}

type ModelRouterConfigureMemberBalanceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterConfigureMemberBalanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterConfigureMemberBalanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterConfigureMemberBalanceResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterConfigureMemberBalanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterConfigureMemberBalanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterConfigureMemberBalanceResponse) GetBody() *ModelRouterConfigureMemberBalanceResponseBody {
	return s.Body
}

func (s *ModelRouterConfigureMemberBalanceResponse) SetHeaders(v map[string]*string) *ModelRouterConfigureMemberBalanceResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterConfigureMemberBalanceResponse) SetStatusCode(v int32) *ModelRouterConfigureMemberBalanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterConfigureMemberBalanceResponse) SetBody(v *ModelRouterConfigureMemberBalanceResponseBody) *ModelRouterConfigureMemberBalanceResponse {
	s.Body = v
	return s
}

func (s *ModelRouterConfigureMemberBalanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
