// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEipInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *CreateEipInstanceResponseBody
	GetAllocationId() *string
	SetRequestId(v string) *CreateEipInstanceResponseBody
	GetRequestId() *string
}

type CreateEipInstanceResponseBody struct {
	// The ID of the EIP.
	//
	// example:
	//
	// eip-25877c70gddh****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9DB6123F-470D-510A-A9EB-EBA799340452
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEipInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEipInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEipInstanceResponseBody) GetAllocationId() *string {
	return s.AllocationId
}

func (s *CreateEipInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEipInstanceResponseBody) SetAllocationId(v string) *CreateEipInstanceResponseBody {
	s.AllocationId = &v
	return s
}

func (s *CreateEipInstanceResponseBody) SetRequestId(v string) *CreateEipInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEipInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
