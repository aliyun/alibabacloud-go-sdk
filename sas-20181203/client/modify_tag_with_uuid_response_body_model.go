// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTagWithUuidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTagWithUuidResponseBody
	GetRequestId() *string
}

type ModifyTagWithUuidResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 70C4B40D-D55E-4B7B-9992-8535E396B2F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTagWithUuidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTagWithUuidResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTagWithUuidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTagWithUuidResponseBody) SetRequestId(v string) *ModifyTagWithUuidResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTagWithUuidResponseBody) Validate() error {
	return dara.Validate(s)
}
