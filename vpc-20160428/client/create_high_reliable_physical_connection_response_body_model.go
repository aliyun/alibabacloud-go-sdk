// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHighReliablePhysicalConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorInfoList(v *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoList) *CreateHighReliablePhysicalConnectionResponseBody
	GetErrorInfoList() *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoList
	SetPhysicalConnectionList(v *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionList) *CreateHighReliablePhysicalConnectionResponseBody
	GetPhysicalConnectionList() *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionList
	SetRequestId(v string) *CreateHighReliablePhysicalConnectionResponseBody
	GetRequestId() *string
}

type CreateHighReliablePhysicalConnectionResponseBody struct {
	// If the request fails the dry run, the following error codes and error messages may be returned:
	//
	// - pconn.high.reliable.dryrun.error.disable.outbound.data.transfer.billing. Billing for outbound data transfer is not enabled.
	//
	// - pconn.high.reliable.dryrun.error.incompatable.device.capacity. No device in the access point supports advanced features.
	//
	// - pconn.high.reliable.dryrun.error.quota.exceeded. The quota is insufficient.
	//
	// - pconn.high.reliable.dryrun.error.not.enough.resource. The access point resources are insufficient.
	ErrorInfoList *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoList `json:"ErrorInfoList,omitempty" xml:"ErrorInfoList,omitempty" type:"Struct"`
	// The Express Connect circuits.
	PhysicalConnectionList *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionList `json:"PhysicalConnectionList,omitempty" xml:"PhysicalConnectionList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHighReliablePhysicalConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHighReliablePhysicalConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHighReliablePhysicalConnectionResponseBody) GetErrorInfoList() *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoList {
	return s.ErrorInfoList
}

func (s *CreateHighReliablePhysicalConnectionResponseBody) GetPhysicalConnectionList() *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionList {
	return s.PhysicalConnectionList
}

func (s *CreateHighReliablePhysicalConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHighReliablePhysicalConnectionResponseBody) SetErrorInfoList(v *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoList) *CreateHighReliablePhysicalConnectionResponseBody {
	s.ErrorInfoList = v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponseBody) SetPhysicalConnectionList(v *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionList) *CreateHighReliablePhysicalConnectionResponseBody {
	s.PhysicalConnectionList = v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponseBody) SetRequestId(v string) *CreateHighReliablePhysicalConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateHighReliablePhysicalConnectionResponseBodyErrorInfoList struct {
	ErrorInfoList []*CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList `json:"errorInfoList,omitempty" xml:"errorInfoList,omitempty" type:"Repeated"`
}

func (s CreateHighReliablePhysicalConnectionResponseBodyErrorInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreateHighReliablePhysicalConnectionResponseBodyErrorInfoList) GoString() string {
	return s.String()
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoList) GetErrorInfoList() []*CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList {
	return s.ErrorInfoList
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoList) SetErrorInfoList(v []*CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList) *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoList {
	s.ErrorInfoList = v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoList) Validate() error {
	return dara.Validate(s)
}

type CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList struct {
	// Error codes.
	//
	// example:
	//
	// pconn.high.reliable.dryrun.error.disable.outbound.data.transfer.billing
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// pconn.high.reliable.dryrun.error.disable.outbound.data.transfer.billing
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the Express Connect circuit.
	//
	// example:
	//
	// pc-j5e5qqo616p81ncspbll1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList) GoString() string {
	return s.String()
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList) SetErrorCode(v string) *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList {
	s.ErrorCode = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList) SetErrorMessage(v string) *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList {
	s.ErrorMessage = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList) SetInstanceId(v string) *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList {
	s.InstanceId = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyErrorInfoListErrorInfoList) Validate() error {
	return dara.Validate(s)
}

type CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionList struct {
	PhysicalConnectionList []*CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList `json:"physicalConnectionList,omitempty" xml:"physicalConnectionList,omitempty" type:"Repeated"`
}

func (s CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionList) String() string {
	return dara.Prettify(s)
}

func (s CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionList) GoString() string {
	return s.String()
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionList) GetPhysicalConnectionList() []*CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList {
	return s.PhysicalConnectionList
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionList) SetPhysicalConnectionList(v []*CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList) *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionList {
	s.PhysicalConnectionList = v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionList) Validate() error {
	return dara.Validate(s)
}

type CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList struct {
	// The ID of the Express Connect circuit.
	//
	// example:
	//
	// pc-j5e5qqo616p81ncspbll1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList) String() string {
	return dara.Prettify(s)
}

func (s CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList) GoString() string {
	return s.String()
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList) SetInstanceId(v string) *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList {
	s.InstanceId = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList) SetRegionNo(v string) *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList {
	s.RegionNo = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponseBodyPhysicalConnectionListPhysicalConnectionList) Validate() error {
	return dara.Validate(s)
}
