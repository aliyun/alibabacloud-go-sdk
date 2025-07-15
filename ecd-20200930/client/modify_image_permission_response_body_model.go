// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImagePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyImagePermissionResponseBody
	GetRequestId() *string
}

type ModifyImagePermissionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImagePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyImagePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImagePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyImagePermissionResponseBody) SetRequestId(v string) *ModifyImagePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyImagePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
