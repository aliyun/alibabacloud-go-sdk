// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAreaBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebAreaBlockResponseBody
	GetRequestId() *string
}

type ModifyWebAreaBlockResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5AA2BD65-E289-4E91-9DD9-3E1FB2140D17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebAreaBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAreaBlockResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebAreaBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebAreaBlockResponseBody) SetRequestId(v string) *ModifyWebAreaBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebAreaBlockResponseBody) Validate() error {
	return dara.Validate(s)
}
