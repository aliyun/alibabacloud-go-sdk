// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteBridgeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyOfficeSiteBridgeInfoResponseBody
	GetRequestId() *string
}

type ModifyOfficeSiteBridgeInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOfficeSiteBridgeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteBridgeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteBridgeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOfficeSiteBridgeInfoResponseBody) SetRequestId(v string) *ModifyOfficeSiteBridgeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOfficeSiteBridgeInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
