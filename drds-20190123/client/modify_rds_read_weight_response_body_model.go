// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRdsReadWeightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRdsReadWeightResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyRdsReadWeightResponseBody
	GetSuccess() *bool
}

type ModifyRdsReadWeightResponseBody struct {
	// Indicates the ID of the request.
	//
	// example:
	//
	// B12FC174-D5CE-4A6E-83C1-0F8F86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyRdsReadWeightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRdsReadWeightResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRdsReadWeightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRdsReadWeightResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyRdsReadWeightResponseBody) SetRequestId(v string) *ModifyRdsReadWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRdsReadWeightResponseBody) SetSuccess(v bool) *ModifyRdsReadWeightResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyRdsReadWeightResponseBody) Validate() error {
	return dara.Validate(s)
}
