// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterVpcResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterVpcResponseBody) *ModifyDBClusterVpcResponse
	GetBody() *ModifyDBClusterVpcResponseBody
}

type ModifyDBClusterVpcResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterVpcResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterVpcResponse) GetBody() *ModifyDBClusterVpcResponseBody {
	return s.Body
}

func (s *ModifyDBClusterVpcResponse) SetHeaders(v map[string]*string) *ModifyDBClusterVpcResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterVpcResponse) SetStatusCode(v int32) *ModifyDBClusterVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterVpcResponse) SetBody(v *ModifyDBClusterVpcResponseBody) *ModifyDBClusterVpcResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterVpcResponse) Validate() error {
	return dara.Validate(s)
}
