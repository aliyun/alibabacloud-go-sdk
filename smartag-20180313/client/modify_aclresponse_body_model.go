// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyACLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyACLResponseBody
	GetRequestId() *string
}

type ModifyACLResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 076FD0BE-67D5-4338-A2A1-C54DE7D78B16
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyACLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyACLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyACLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyACLResponseBody) SetRequestId(v string) *ModifyACLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyACLResponseBody) Validate() error {
	return dara.Validate(s)
}
