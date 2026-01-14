// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDdosFromAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDdosId(v string) *DetachDdosFromAcceleratorResponseBody
	GetDdosId() *string
	SetRequestId(v string) *DetachDdosFromAcceleratorResponseBody
	GetRequestId() *string
}

type DetachDdosFromAcceleratorResponseBody struct {
	// Deprecated
	//
	// The ID of the Anti-DDoS Pro/Premium instance that was disassociated from the GA instance.
	//
	// example:
	//
	// ddoscoo-cn-zz11vq7j****
	DdosId *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachDdosFromAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachDdosFromAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *DetachDdosFromAcceleratorResponseBody) GetDdosId() *string {
	return s.DdosId
}

func (s *DetachDdosFromAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachDdosFromAcceleratorResponseBody) SetDdosId(v string) *DetachDdosFromAcceleratorResponseBody {
	s.DdosId = &v
	return s
}

func (s *DetachDdosFromAcceleratorResponseBody) SetRequestId(v string) *DetachDdosFromAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachDdosFromAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
