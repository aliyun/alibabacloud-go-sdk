// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationData interface {
	dara.Model
	String() string
	GoString() string
	SetActualDeliveredAmounts(v int32) *OperationData
	GetActualDeliveredAmounts() *int32
	SetToBeDeliveredAmounts(v int32) *OperationData
	GetToBeDeliveredAmounts() *int32
}

type OperationData struct {
	ActualDeliveredAmounts *int32 `json:"actualDeliveredAmounts,omitempty" xml:"actualDeliveredAmounts,omitempty"`
	ToBeDeliveredAmounts   *int32 `json:"toBeDeliveredAmounts,omitempty" xml:"toBeDeliveredAmounts,omitempty"`
}

func (s OperationData) String() string {
	return dara.Prettify(s)
}

func (s OperationData) GoString() string {
	return s.String()
}

func (s *OperationData) GetActualDeliveredAmounts() *int32 {
	return s.ActualDeliveredAmounts
}

func (s *OperationData) GetToBeDeliveredAmounts() *int32 {
	return s.ToBeDeliveredAmounts
}

func (s *OperationData) SetActualDeliveredAmounts(v int32) *OperationData {
	s.ActualDeliveredAmounts = &v
	return s
}

func (s *OperationData) SetToBeDeliveredAmounts(v int32) *OperationData {
	s.ToBeDeliveredAmounts = &v
	return s
}

func (s *OperationData) Validate() error {
	return dara.Validate(s)
}
