// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebIpSetSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebIpSetSwitchResponseBody
	GetRequestId() *string
}

type ModifyWebIpSetSwitchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebIpSetSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebIpSetSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebIpSetSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebIpSetSwitchResponseBody) SetRequestId(v string) *ModifyWebIpSetSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebIpSetSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
