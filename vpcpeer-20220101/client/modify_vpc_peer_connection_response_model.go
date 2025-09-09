// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcPeerConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcPeerConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcPeerConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcPeerConnectionResponseBody) *ModifyVpcPeerConnectionResponse
	GetBody() *ModifyVpcPeerConnectionResponseBody
}

type ModifyVpcPeerConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcPeerConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcPeerConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcPeerConnectionResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcPeerConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcPeerConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcPeerConnectionResponse) GetBody() *ModifyVpcPeerConnectionResponseBody {
	return s.Body
}

func (s *ModifyVpcPeerConnectionResponse) SetHeaders(v map[string]*string) *ModifyVpcPeerConnectionResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcPeerConnectionResponse) SetStatusCode(v int32) *ModifyVpcPeerConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcPeerConnectionResponse) SetBody(v *ModifyVpcPeerConnectionResponseBody) *ModifyVpcPeerConnectionResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcPeerConnectionResponse) Validate() error {
	return dara.Validate(s)
}
