// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootMachineResponseBody
	GetRequestId() *string
}

type RebootMachineResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 151F6EB6-D5F3-417A-AF7B-4D84975DB586
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootMachineResponseBody) GoString() string {
	return s.String()
}

func (s *RebootMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootMachineResponseBody) SetRequestId(v string) *RebootMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootMachineResponseBody) Validate() error {
	return dara.Validate(s)
}
