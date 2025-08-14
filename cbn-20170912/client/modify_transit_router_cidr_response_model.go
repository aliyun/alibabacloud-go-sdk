// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouterCidrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTransitRouterCidrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTransitRouterCidrResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTransitRouterCidrResponseBody) *ModifyTransitRouterCidrResponse
	GetBody() *ModifyTransitRouterCidrResponseBody
}

type ModifyTransitRouterCidrResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTransitRouterCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTransitRouterCidrResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterCidrResponse) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterCidrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTransitRouterCidrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTransitRouterCidrResponse) GetBody() *ModifyTransitRouterCidrResponseBody {
	return s.Body
}

func (s *ModifyTransitRouterCidrResponse) SetHeaders(v map[string]*string) *ModifyTransitRouterCidrResponse {
	s.Headers = v
	return s
}

func (s *ModifyTransitRouterCidrResponse) SetStatusCode(v int32) *ModifyTransitRouterCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTransitRouterCidrResponse) SetBody(v *ModifyTransitRouterCidrResponseBody) *ModifyTransitRouterCidrResponse {
	s.Body = v
	return s
}

func (s *ModifyTransitRouterCidrResponse) Validate() error {
	return dara.Validate(s)
}
