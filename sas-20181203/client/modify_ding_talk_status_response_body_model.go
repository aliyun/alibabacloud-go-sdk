// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDingTalkStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDingTalkStatusResponseBody
	GetRequestId() *string
}

type ModifyDingTalkStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 81DCBD76-196C-57A5-9C7D-F14DE8E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDingTalkStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDingTalkStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDingTalkStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDingTalkStatusResponseBody) SetRequestId(v string) *ModifyDingTalkStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDingTalkStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
