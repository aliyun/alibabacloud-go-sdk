// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCompactionServiceSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCompactionServiceSwitchResponseBody
	GetRequestId() *string
}

type ModifyCompactionServiceSwitchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 21ABF219-10E0-571B-94B8-9C9AE5022BF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCompactionServiceSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCompactionServiceSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCompactionServiceSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCompactionServiceSwitchResponseBody) SetRequestId(v string) *ModifyCompactionServiceSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCompactionServiceSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
