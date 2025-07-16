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
	// 4035409E-7297-4115-ABC9-C1C3942BC069
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
