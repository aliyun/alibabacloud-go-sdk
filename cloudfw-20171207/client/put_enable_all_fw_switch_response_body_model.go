// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEnableAllFwSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutEnableAllFwSwitchResponseBody
	GetRequestId() *string
}

type PutEnableAllFwSwitchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B2841452-CB8D-4F7D-B247-38E1CF7334F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutEnableAllFwSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutEnableAllFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutEnableAllFwSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutEnableAllFwSwitchResponseBody) SetRequestId(v string) *PutEnableAllFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutEnableAllFwSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
