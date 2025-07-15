// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteCrossDesktopAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyOfficeSiteCrossDesktopAccessResponseBody
	GetRequestId() *string
}

type ModifyOfficeSiteCrossDesktopAccessResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOfficeSiteCrossDesktopAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteCrossDesktopAccessResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponseBody) SetRequestId(v string) *ModifyOfficeSiteCrossDesktopAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
