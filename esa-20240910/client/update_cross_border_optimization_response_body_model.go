// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrossBorderOptimizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCrossBorderOptimizationResponseBody
	GetRequestId() *string
}

type UpdateCrossBorderOptimizationResponseBody struct {
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCrossBorderOptimizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrossBorderOptimizationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCrossBorderOptimizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCrossBorderOptimizationResponseBody) SetRequestId(v string) *UpdateCrossBorderOptimizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCrossBorderOptimizationResponseBody) Validate() error {
	return dara.Validate(s)
}
