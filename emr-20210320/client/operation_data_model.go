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
	SetFailedRefundInstanceIds(v []*string) *OperationData
	GetFailedRefundInstanceIds() []*string
	SetToBeDeliveredAmounts(v int32) *OperationData
	GetToBeDeliveredAmounts() *int32
}

type OperationData struct {
	// The number of units actually delivered.
	ActualDeliveredAmounts *int32 `json:"actualDeliveredAmounts,omitempty" xml:"actualDeliveredAmounts,omitempty"`
	// The IDs of instances that could not be refunded.
	FailedRefundInstanceIds []*string `json:"failedRefundInstanceIds,omitempty" xml:"failedRefundInstanceIds,omitempty" type:"Repeated"`
	// The number of units requested.
	ToBeDeliveredAmounts *int32 `json:"toBeDeliveredAmounts,omitempty" xml:"toBeDeliveredAmounts,omitempty"`
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

func (s *OperationData) GetFailedRefundInstanceIds() []*string {
	return s.FailedRefundInstanceIds
}

func (s *OperationData) GetToBeDeliveredAmounts() *int32 {
	return s.ToBeDeliveredAmounts
}

func (s *OperationData) SetActualDeliveredAmounts(v int32) *OperationData {
	s.ActualDeliveredAmounts = &v
	return s
}

func (s *OperationData) SetFailedRefundInstanceIds(v []*string) *OperationData {
	s.FailedRefundInstanceIds = v
	return s
}

func (s *OperationData) SetToBeDeliveredAmounts(v int32) *OperationData {
	s.ToBeDeliveredAmounts = &v
	return s
}

func (s *OperationData) Validate() error {
	return dara.Validate(s)
}
