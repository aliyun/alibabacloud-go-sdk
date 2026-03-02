// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyOfficeSiteAcceleratorResponseBody
	GetRequestId() *string
}

type ModifyOfficeSiteAcceleratorResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOfficeSiteAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOfficeSiteAcceleratorResponseBody) SetRequestId(v string) *ModifyOfficeSiteAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOfficeSiteAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
