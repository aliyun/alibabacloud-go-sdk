// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrossBorderOptimizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *GetCrossBorderOptimizationResponseBody
	GetEnable() *string
	SetRequestId(v string) *GetCrossBorderOptimizationResponseBody
	GetRequestId() *string
}

type GetCrossBorderOptimizationResponseBody struct {
	// Whether to enable Chinese mainland network access optimization. By default, it is disabled. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CF521A24-633F-5350-A6A5-42AD503D0D20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCrossBorderOptimizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCrossBorderOptimizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrossBorderOptimizationResponseBody) GetEnable() *string {
	return s.Enable
}

func (s *GetCrossBorderOptimizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCrossBorderOptimizationResponseBody) SetEnable(v string) *GetCrossBorderOptimizationResponseBody {
	s.Enable = &v
	return s
}

func (s *GetCrossBorderOptimizationResponseBody) SetRequestId(v string) *GetCrossBorderOptimizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCrossBorderOptimizationResponseBody) Validate() error {
	return dara.Validate(s)
}
