// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecUserOperationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecUserOperationsResponseBodyData) *DescribeApisecUserOperationsResponseBody
	GetData() []*DescribeApisecUserOperationsResponseBodyData
	SetRequestId(v string) *DescribeApisecUserOperationsResponseBody
	GetRequestId() *string
}

type DescribeApisecUserOperationsResponseBody struct {
	// The user operation records for API security.
	Data []*DescribeApisecUserOperationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C9825654-327B-5156-A570-847054B4CF10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApisecUserOperationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecUserOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecUserOperationsResponseBody) GetData() []*DescribeApisecUserOperationsResponseBodyData {
	return s.Data
}

func (s *DescribeApisecUserOperationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecUserOperationsResponseBody) SetData(v []*DescribeApisecUserOperationsResponseBodyData) *DescribeApisecUserOperationsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecUserOperationsResponseBody) SetRequestId(v string) *DescribeApisecUserOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecUserOperationsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisecUserOperationsResponseBodyData struct {
	// The status of the threat detection or security event before the operation was performed.
	//
	// Valid values for threat detection:
	//
	// - **toBeConfirmed**: to be confirmed.
	//
	// - **confirmed**: confirmed.
	//
	// - **toBeFixed**: to be fixed.
	//
	// - **fixed**: fixed.
	//
	// - **ignored**: ignored.
	//
	// Valid values for a security event:
	//
	// - **toBeConfirmed**: to be confirmed.
	//
	// - **confirmed**: confirmed.
	//
	// - **ignored**: ignored.
	//
	// example:
	//
	// ignored
	FromStatus *string `json:"FromStatus,omitempty" xml:"FromStatus,omitempty"`
	// The remarks that the user added to the operation record.
	//
	// example:
	//
	// Procesed
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the threat detection or security event associated with the operation record.
	//
	// example:
	//
	// 24d997acc48a67a01e09b9c5ad861287
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The source of the operation. Valid values:
	//
	// - **system**: the operation was automatically performed by the system.
	//
	// - **custom**: the operation was manually performed by a user.
	//
	// example:
	//
	// custom
	OperationSource *string `json:"OperationSource,omitempty" xml:"OperationSource,omitempty"`
	// The time when the operation was performed. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1685072214
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The status of the threat detection or security event after the operation was performed.
	//
	// Valid values for threat detection:
	//
	// - **toBeConfirmed**: to be confirmed.
	//
	// - **confirmed**: confirmed.
	//
	// - **toBeFixed**: to be fixed.
	//
	// - **fixed**: fixed.
	//
	// - **ignored**: ignored.
	//
	// Valid values for a security event:
	//
	// - **toBeConfirmed**: to be confirmed.
	//
	// - **confirmed**: confirmed.
	//
	// - **ignored**: ignored.
	//
	// example:
	//
	// Confirmed
	ToStatus *string `json:"ToStatus,omitempty" xml:"ToStatus,omitempty"`
	// The type of the operation record. Valid values:
	//
	// - **abnormal**: threat detection.
	//
	// - **event**: security event.
	//
	// example:
	//
	// abnormal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the user who performed the operation.
	//
	// example:
	//
	// 1610954****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeApisecUserOperationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecUserOperationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecUserOperationsResponseBodyData) GetFromStatus() *string {
	return s.FromStatus
}

func (s *DescribeApisecUserOperationsResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *DescribeApisecUserOperationsResponseBodyData) GetObjectId() *string {
	return s.ObjectId
}

func (s *DescribeApisecUserOperationsResponseBodyData) GetOperationSource() *string {
	return s.OperationSource
}

func (s *DescribeApisecUserOperationsResponseBodyData) GetTime() *int64 {
	return s.Time
}

func (s *DescribeApisecUserOperationsResponseBodyData) GetToStatus() *string {
	return s.ToStatus
}

func (s *DescribeApisecUserOperationsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DescribeApisecUserOperationsResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *DescribeApisecUserOperationsResponseBodyData) SetFromStatus(v string) *DescribeApisecUserOperationsResponseBodyData {
	s.FromStatus = &v
	return s
}

func (s *DescribeApisecUserOperationsResponseBodyData) SetNote(v string) *DescribeApisecUserOperationsResponseBodyData {
	s.Note = &v
	return s
}

func (s *DescribeApisecUserOperationsResponseBodyData) SetObjectId(v string) *DescribeApisecUserOperationsResponseBodyData {
	s.ObjectId = &v
	return s
}

func (s *DescribeApisecUserOperationsResponseBodyData) SetOperationSource(v string) *DescribeApisecUserOperationsResponseBodyData {
	s.OperationSource = &v
	return s
}

func (s *DescribeApisecUserOperationsResponseBodyData) SetTime(v int64) *DescribeApisecUserOperationsResponseBodyData {
	s.Time = &v
	return s
}

func (s *DescribeApisecUserOperationsResponseBodyData) SetToStatus(v string) *DescribeApisecUserOperationsResponseBodyData {
	s.ToStatus = &v
	return s
}

func (s *DescribeApisecUserOperationsResponseBodyData) SetType(v string) *DescribeApisecUserOperationsResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeApisecUserOperationsResponseBodyData) SetUserId(v string) *DescribeApisecUserOperationsResponseBodyData {
	s.UserId = &v
	return s
}

func (s *DescribeApisecUserOperationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
