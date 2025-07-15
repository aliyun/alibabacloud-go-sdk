// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVSwitchResponseBody
	GetRequestId() *string
}

type DeleteVSwitchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVSwitchResponseBody) SetRequestId(v string) *DeleteVSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
