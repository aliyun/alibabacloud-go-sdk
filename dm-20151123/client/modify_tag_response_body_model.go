// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTagResponseBody
	GetRequestId() *string
}

type ModifyTagResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 00BD30D8-2E86-523A-BFC7-63B7FF931A06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTagResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTagResponseBody) SetRequestId(v string) *ModifyTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTagResponseBody) Validate() error {
	return dara.Validate(s)
}
