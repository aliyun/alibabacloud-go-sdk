// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateEipSegmentAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEipSegmentInstanceId(v string) *AllocateEipSegmentAddressResponseBody
	GetEipSegmentInstanceId() *string
	SetRequestId(v string) *AllocateEipSegmentAddressResponseBody
	GetRequestId() *string
}

type AllocateEipSegmentAddressResponseBody struct {
	// The ID of the contiguous EIP group.
	//
	// example:
	//
	// eipsg-2zett8ba055tbsxme****
	EipSegmentInstanceId *string `json:"EipSegmentInstanceId,omitempty" xml:"EipSegmentInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F7A6301A-64BA-41EC-8284-8F4838C15D1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateEipSegmentAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipSegmentAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateEipSegmentAddressResponseBody) GetEipSegmentInstanceId() *string {
	return s.EipSegmentInstanceId
}

func (s *AllocateEipSegmentAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateEipSegmentAddressResponseBody) SetEipSegmentInstanceId(v string) *AllocateEipSegmentAddressResponseBody {
	s.EipSegmentInstanceId = &v
	return s
}

func (s *AllocateEipSegmentAddressResponseBody) SetRequestId(v string) *AllocateEipSegmentAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateEipSegmentAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
