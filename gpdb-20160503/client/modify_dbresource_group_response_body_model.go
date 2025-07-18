// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBResourceGroupResponseBody
	GetRequestId() *string
}

type ModifyDBResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBResourceGroupResponseBody) SetRequestId(v string) *ModifyDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
