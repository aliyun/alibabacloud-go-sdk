// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterConnectionStringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterConnectionStringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterConnectionStringResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterConnectionStringResponseBody) *ModifyClusterConnectionStringResponse
	GetBody() *ModifyClusterConnectionStringResponseBody
}

type ModifyClusterConnectionStringResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterConnectionStringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterConnectionStringResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterConnectionStringResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterConnectionStringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterConnectionStringResponse) GetBody() *ModifyClusterConnectionStringResponseBody {
	return s.Body
}

func (s *ModifyClusterConnectionStringResponse) SetHeaders(v map[string]*string) *ModifyClusterConnectionStringResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterConnectionStringResponse) SetStatusCode(v int32) *ModifyClusterConnectionStringResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterConnectionStringResponse) SetBody(v *ModifyClusterConnectionStringResponseBody) *ModifyClusterConnectionStringResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterConnectionStringResponse) Validate() error {
	return dara.Validate(s)
}
