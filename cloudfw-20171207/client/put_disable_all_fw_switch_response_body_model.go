// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDisableAllFwSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutDisableAllFwSwitchResponseBody
	GetRequestId() *string
}

type PutDisableAllFwSwitchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B2841452-CB8D-4F7D-B247-38E1CF7334F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutDisableAllFwSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutDisableAllFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutDisableAllFwSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutDisableAllFwSwitchResponseBody) SetRequestId(v string) *PutDisableAllFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutDisableAllFwSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
