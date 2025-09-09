// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarDbReadWeightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPolarDbReadWeightResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyPolarDbReadWeightResponseBody
	GetSuccess() *bool
}

type ModifyPolarDbReadWeightResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B12FC174-D5CE-4A6E-83C1-0F8F86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the database creation failure records were removed from the PolarDB-X instance.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPolarDbReadWeightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarDbReadWeightResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolarDbReadWeightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPolarDbReadWeightResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyPolarDbReadWeightResponseBody) SetRequestId(v string) *ModifyPolarDbReadWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPolarDbReadWeightResponseBody) SetSuccess(v bool) *ModifyPolarDbReadWeightResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyPolarDbReadWeightResponseBody) Validate() error {
	return dara.Validate(s)
}
