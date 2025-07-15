// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNASDefaultMountTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNASDefaultMountTargetResponseBody
	GetRequestId() *string
}

type ModifyNASDefaultMountTargetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNASDefaultMountTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNASDefaultMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNASDefaultMountTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNASDefaultMountTargetResponseBody) SetRequestId(v string) *ModifyNASDefaultMountTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNASDefaultMountTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
