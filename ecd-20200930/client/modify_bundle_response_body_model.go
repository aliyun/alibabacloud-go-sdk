// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBundleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBundleResponseBody
	GetRequestId() *string
}

type ModifyBundleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBundleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBundleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBundleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBundleResponseBody) SetRequestId(v string) *ModifyBundleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBundleResponseBody) Validate() error {
	return dara.Validate(s)
}
