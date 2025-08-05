// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserIPSWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUserIPSWhitelistResponseBody
	GetRequestId() *string
}

type ModifyUserIPSWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserIPSWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserIPSWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserIPSWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserIPSWhitelistResponseBody) SetRequestId(v string) *ModifyUserIPSWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserIPSWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
