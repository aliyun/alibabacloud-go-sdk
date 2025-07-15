// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetNASDefaultMountTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetNASDefaultMountTargetResponseBody
	GetRequestId() *string
}

type ResetNASDefaultMountTargetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetNASDefaultMountTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetNASDefaultMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ResetNASDefaultMountTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetNASDefaultMountTargetResponseBody) SetRequestId(v string) *ResetNASDefaultMountTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetNASDefaultMountTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
