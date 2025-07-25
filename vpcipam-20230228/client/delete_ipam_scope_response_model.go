// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpamScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpamScopeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpamScopeResponseBody) *DeleteIpamScopeResponse
	GetBody() *DeleteIpamScopeResponseBody
}

type DeleteIpamScopeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpamScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpamScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamScopeResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpamScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpamScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpamScopeResponse) GetBody() *DeleteIpamScopeResponseBody {
	return s.Body
}

func (s *DeleteIpamScopeResponse) SetHeaders(v map[string]*string) *DeleteIpamScopeResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpamScopeResponse) SetStatusCode(v int32) *DeleteIpamScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpamScopeResponse) SetBody(v *DeleteIpamScopeResponseBody) *DeleteIpamScopeResponse {
	s.Body = v
	return s
}

func (s *DeleteIpamScopeResponse) Validate() error {
	return dara.Validate(s)
}
