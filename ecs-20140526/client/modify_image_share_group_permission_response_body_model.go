// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageShareGroupPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyImageShareGroupPermissionResponseBody
	GetRequestId() *string
}

type ModifyImageShareGroupPermissionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 123e4567-e89b-42d3-a456-426614174000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageShareGroupPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageShareGroupPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageShareGroupPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyImageShareGroupPermissionResponseBody) SetRequestId(v string) *ModifyImageShareGroupPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyImageShareGroupPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
