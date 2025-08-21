// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebPreciseAccessSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebPreciseAccessSwitchResponseBody
	GetRequestId() *string
}

type ModifyWebPreciseAccessSwitchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebPreciseAccessSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebPreciseAccessSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebPreciseAccessSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebPreciseAccessSwitchResponseBody) SetRequestId(v string) *ModifyWebPreciseAccessSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebPreciseAccessSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
