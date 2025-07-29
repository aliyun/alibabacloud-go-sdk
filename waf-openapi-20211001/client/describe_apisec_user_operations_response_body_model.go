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
	// The operation records.
	Data []*DescribeApisecUserOperationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
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
	return dara.Validate(s)
}

type DescribeApisecUserOperationsResponseBodyData struct {
	// The state before the operation.
	//
	// Valid values of the risk state:
	//
	// 	- **toBeConfirmed**
	//
	// 	- **confirmed**
	//
	// 	- **toBeFixed**
	//
	// 	- **fixed**
	//
	// 	- **ignored**
	//
	// Valid values of the event state:
	//
	// 	- **toBeConfirmed**
	//
	// 	- **confirmed**
	//
	// 	- **ignored**
	//
	// example:
	//
	// ignored
	FromStatus *string `json:"FromStatus,omitempty" xml:"FromStatus,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Handled
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The object ID of the operation record.
	//
	// example:
	//
	// 24d997acc48a67a01e09b9c5ad861287
	ObjectId        *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	OperationSource *string `json:"OperationSource,omitempty" xml:"OperationSource,omitempty"`
	// The time at which the operation was performed. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1685072214
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The state after the operation.
	//
	// Valid values of the risk state:
	//
	// 	- **toBeConfirmed**
	//
	// 	- **confirmed**
	//
	// 	- **toBeFixed**
	//
	// 	- **fixed**
	//
	// 	- **ignored**
	//
	// Valid values of the event state:
	//
	// 	- **toBeConfirmed**
	//
	// 	- **confirmed**
	//
	// 	- **ignored**
	//
	// example:
	//
	// Confirmed
	ToStatus *string `json:"ToStatus,omitempty" xml:"ToStatus,omitempty"`
	// The type of the operation record. Valid values:
	//
	// 	- **abnormal**: risk detection
	//
	// 	- **event**: security event
	//
	// example:
	//
	// abnormal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user ID.
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
