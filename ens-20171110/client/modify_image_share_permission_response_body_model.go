// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageSharePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyImageSharePermissionResponseBody
	GetRequestId() *string
}

type ModifyImageSharePermissionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 701B3BB9-9190-544D-90D1-328B0527380C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageSharePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageSharePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageSharePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyImageSharePermissionResponseBody) SetRequestId(v string) *ModifyImageSharePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyImageSharePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
