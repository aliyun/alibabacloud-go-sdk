// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReservedNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateReservedNodePoolResponseBody
	GetRequestId() *string
	SetReservedNodePoolId(v string) *CreateReservedNodePoolResponseBody
	GetReservedNodePoolId() *string
}

type CreateReservedNodePoolResponseBody struct {
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rnp-cdx****
	ReservedNodePoolId *string `json:"ReservedNodePoolId,omitempty" xml:"ReservedNodePoolId,omitempty"`
}

func (s CreateReservedNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReservedNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReservedNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateReservedNodePoolResponseBody) GetReservedNodePoolId() *string {
	return s.ReservedNodePoolId
}

func (s *CreateReservedNodePoolResponseBody) SetRequestId(v string) *CreateReservedNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReservedNodePoolResponseBody) SetReservedNodePoolId(v string) *CreateReservedNodePoolResponseBody {
	s.ReservedNodePoolId = &v
	return s
}

func (s *CreateReservedNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
