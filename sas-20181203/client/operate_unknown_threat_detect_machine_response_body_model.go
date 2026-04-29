// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateUnknownThreatDetectMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateUnknownThreatDetectMachineResponseBody
	GetRequestId() *string
}

type OperateUnknownThreatDetectMachineResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 20456DD5-5CBF-5015-9173-12CA4246B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateUnknownThreatDetectMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateUnknownThreatDetectMachineResponseBody) GoString() string {
	return s.String()
}

func (s *OperateUnknownThreatDetectMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateUnknownThreatDetectMachineResponseBody) SetRequestId(v string) *OperateUnknownThreatDetectMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateUnknownThreatDetectMachineResponseBody) Validate() error {
	return dara.Validate(s)
}
