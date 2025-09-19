// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVulTargetResponseBody
	GetRequestId() *string
}

type ModifyVulTargetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 38597320-A990-5444-9A4C-7A1269610C2A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVulTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVulTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVulTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVulTargetResponseBody) SetRequestId(v string) *ModifyVulTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVulTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
