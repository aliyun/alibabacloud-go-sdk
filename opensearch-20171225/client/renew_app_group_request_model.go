// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PrepayOrderInfo) *RenewAppGroupRequest
	GetBody() *PrepayOrderInfo
	SetClientToken(v string) *RenewAppGroupRequest
	GetClientToken() *string
}

type RenewAppGroupRequest struct {
	// The renewal request body.
	Body *PrepayOrderInfo `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 74db41d8cd3c784209093aa76afbe89e
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s RenewAppGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewAppGroupRequest) GoString() string {
	return s.String()
}

func (s *RenewAppGroupRequest) GetBody() *PrepayOrderInfo {
	return s.Body
}

func (s *RenewAppGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewAppGroupRequest) SetBody(v *PrepayOrderInfo) *RenewAppGroupRequest {
	s.Body = v
	return s
}

func (s *RenewAppGroupRequest) SetClientToken(v string) *RenewAppGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewAppGroupRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
