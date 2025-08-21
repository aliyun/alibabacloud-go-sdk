// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAIProtectSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebAIProtectSwitchResponseBody
	GetRequestId() *string
}

type ModifyWebAIProtectSwitchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebAIProtectSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAIProtectSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebAIProtectSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebAIProtectSwitchResponseBody) SetRequestId(v string) *ModifyWebAIProtectSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebAIProtectSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
