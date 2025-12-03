// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskWarningLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDiskWarningLineResponseBody
	GetRequestId() *string
}

type ModifyDiskWarningLineResponseBody struct {
	// example:
	//
	// FC4A930D-3AEE-4C9D-BC70-C0F2EEEAA174
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskWarningLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskWarningLineResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskWarningLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskWarningLineResponseBody) SetRequestId(v string) *ModifyDiskWarningLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskWarningLineResponseBody) Validate() error {
	return dara.Validate(s)
}
