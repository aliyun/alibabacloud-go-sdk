// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationMaintainConfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyActiveOperationMaintainConfResponseBody
	GetRequestId() *string
}

type ModifyActiveOperationMaintainConfResponseBody struct {
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyActiveOperationMaintainConfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationMaintainConfResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintainConfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyActiveOperationMaintainConfResponseBody) SetRequestId(v string) *ModifyActiveOperationMaintainConfResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfResponseBody) Validate() error {
	return dara.Validate(s)
}
