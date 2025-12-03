// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIpWhitelistResponseBody
	GetRequestId() *string
}

type ModifyIpWhitelistResponseBody struct {
	// example:
	//
	// 101CFA8A-FB88-5014-A10C-3A0DA9AD8B0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIpWhitelistResponseBody) SetRequestId(v string) *ModifyIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIpWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
