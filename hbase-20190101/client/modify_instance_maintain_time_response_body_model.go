// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMaintainTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceMaintainTimeResponseBody
	GetRequestId() *string
}

type ModifyInstanceMaintainTimeResponseBody struct {
	// example:
	//
	// C9085433-A56A-4089-B49A-DF5A4E2B7B06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMaintainTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceMaintainTimeResponseBody) SetRequestId(v string) *ModifyInstanceMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceMaintainTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
