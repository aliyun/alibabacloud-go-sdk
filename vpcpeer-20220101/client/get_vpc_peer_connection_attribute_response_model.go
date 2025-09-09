// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcPeerConnectionAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpcPeerConnectionAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpcPeerConnectionAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetVpcPeerConnectionAttributeResponseBody) *GetVpcPeerConnectionAttributeResponse
	GetBody() *GetVpcPeerConnectionAttributeResponseBody
}

type GetVpcPeerConnectionAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpcPeerConnectionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpcPeerConnectionAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPeerConnectionAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetVpcPeerConnectionAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpcPeerConnectionAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpcPeerConnectionAttributeResponse) GetBody() *GetVpcPeerConnectionAttributeResponseBody {
	return s.Body
}

func (s *GetVpcPeerConnectionAttributeResponse) SetHeaders(v map[string]*string) *GetVpcPeerConnectionAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponse) SetStatusCode(v int32) *GetVpcPeerConnectionAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponse) SetBody(v *GetVpcPeerConnectionAttributeResponseBody) *GetVpcPeerConnectionAttributeResponse {
	s.Body = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponse) Validate() error {
	return dara.Validate(s)
}
