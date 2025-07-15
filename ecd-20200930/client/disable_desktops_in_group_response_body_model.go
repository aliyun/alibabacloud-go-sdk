// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDesktopsInGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableDesktopsInGroupResponseBody
	GetRequestId() *string
}

type DisableDesktopsInGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 34FB4D97-C0D9-5534-ABC6-90C7EBD5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDesktopsInGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDesktopsInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDesktopsInGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDesktopsInGroupResponseBody) SetRequestId(v string) *DisableDesktopsInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDesktopsInGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
