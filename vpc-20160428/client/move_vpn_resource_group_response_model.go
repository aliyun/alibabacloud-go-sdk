// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveVpnResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveVpnResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveVpnResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *MoveVpnResourceGroupResponseBody) *MoveVpnResourceGroupResponse
	GetBody() *MoveVpnResourceGroupResponseBody
}

type MoveVpnResourceGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveVpnResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveVpnResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveVpnResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveVpnResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveVpnResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveVpnResourceGroupResponse) GetBody() *MoveVpnResourceGroupResponseBody {
	return s.Body
}

func (s *MoveVpnResourceGroupResponse) SetHeaders(v map[string]*string) *MoveVpnResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveVpnResourceGroupResponse) SetStatusCode(v int32) *MoveVpnResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveVpnResourceGroupResponse) SetBody(v *MoveVpnResourceGroupResponseBody) *MoveVpnResourceGroupResponse {
	s.Body = v
	return s
}

func (s *MoveVpnResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
