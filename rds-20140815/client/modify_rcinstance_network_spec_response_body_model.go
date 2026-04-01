// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceNetworkSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRCInstanceNetworkSpecResponseBody
	GetRequestId() *string
}

type ModifyRCInstanceNetworkSpecResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8F347CA3-D6AB-5045-9026-24578801F781
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCInstanceNetworkSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceNetworkSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceNetworkSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCInstanceNetworkSpecResponseBody) SetRequestId(v string) *ModifyRCInstanceNetworkSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCInstanceNetworkSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
