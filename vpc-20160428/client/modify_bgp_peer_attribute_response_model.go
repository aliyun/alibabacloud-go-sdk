// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBgpPeerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBgpPeerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBgpPeerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBgpPeerAttributeResponseBody) *ModifyBgpPeerAttributeResponse
	GetBody() *ModifyBgpPeerAttributeResponseBody
}

type ModifyBgpPeerAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBgpPeerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBgpPeerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBgpPeerAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyBgpPeerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBgpPeerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBgpPeerAttributeResponse) GetBody() *ModifyBgpPeerAttributeResponseBody {
	return s.Body
}

func (s *ModifyBgpPeerAttributeResponse) SetHeaders(v map[string]*string) *ModifyBgpPeerAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyBgpPeerAttributeResponse) SetStatusCode(v int32) *ModifyBgpPeerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBgpPeerAttributeResponse) SetBody(v *ModifyBgpPeerAttributeResponseBody) *ModifyBgpPeerAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyBgpPeerAttributeResponse) Validate() error {
	return dara.Validate(s)
}
