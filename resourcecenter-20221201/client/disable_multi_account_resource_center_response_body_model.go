// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableMultiAccountResourceCenterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableMultiAccountResourceCenterResponseBody
	GetRequestId() *string
}

type DisableMultiAccountResourceCenterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4951F920-48DB-5731-96AA-3A7C8AE617D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableMultiAccountResourceCenterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableMultiAccountResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DisableMultiAccountResourceCenterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableMultiAccountResourceCenterResponseBody) SetRequestId(v string) *DisableMultiAccountResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableMultiAccountResourceCenterResponseBody) Validate() error {
	return dara.Validate(s)
}
