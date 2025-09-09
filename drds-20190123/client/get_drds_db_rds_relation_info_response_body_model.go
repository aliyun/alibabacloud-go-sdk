// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDrdsDbRdsRelationInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetDrdsDbRdsRelationInfoResponseBodyData) *GetDrdsDbRdsRelationInfoResponseBody
	GetData() []*GetDrdsDbRdsRelationInfoResponseBodyData
	SetRequestId(v string) *GetDrdsDbRdsRelationInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDrdsDbRdsRelationInfoResponseBody
	GetSuccess() *bool
}

type GetDrdsDbRdsRelationInfoResponseBody struct {
	// The structure information about the storage instances of the DRDS database. Each entry corresponds to a primary storage instance.
	Data []*GetDrdsDbRdsRelationInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 82FD0D9B-9A65-40D3-B1D9-8851B1D4AF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDrdsDbRdsRelationInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDrdsDbRdsRelationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) GetData() []*GetDrdsDbRdsRelationInfoResponseBodyData {
	return s.Data
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) SetData(v []*GetDrdsDbRdsRelationInfoResponseBodyData) *GetDrdsDbRdsRelationInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) SetRequestId(v string) *GetDrdsDbRdsRelationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) SetSuccess(v bool) *GetDrdsDbRdsRelationInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDrdsDbRdsRelationInfoResponseBodyData struct {
	// The ID of the storage instance.
	//
	// example:
	//
	// rm-bp16ad920ndxxxx02
	RdsInstanceId *string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	// The IDs of the read-only storage instances.
	ReadOnlyInstanceInfo []*string `json:"ReadOnlyInstanceInfo,omitempty" xml:"ReadOnlyInstanceInfo,omitempty" type:"Repeated"`
	// The ID of the storage instance that is in use. If the specified instance in the request is a primary DRDS instance, the value of this parameter is the ID of the primary storage instance. If the specified instance in the request is a read-only DRDS instance, the value of this parameter is the ID of the secondary storage instance.
	//
	// example:
	//
	// rm-bp1l8xi1dd9xxxxbj
	UsedInstanceId *string `json:"UsedInstanceId,omitempty" xml:"UsedInstanceId,omitempty"`
	// The type of the storage instance that is in use.
	//
	// example:
	//
	// RDS
	UsedInstanceType *string `json:"UsedInstanceType,omitempty" xml:"UsedInstanceType,omitempty"`
}

func (s GetDrdsDbRdsRelationInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDrdsDbRdsRelationInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) GetRdsInstanceId() *string {
	return s.RdsInstanceId
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) GetReadOnlyInstanceInfo() []*string {
	return s.ReadOnlyInstanceInfo
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) GetUsedInstanceId() *string {
	return s.UsedInstanceId
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) GetUsedInstanceType() *string {
	return s.UsedInstanceType
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) SetRdsInstanceId(v string) *GetDrdsDbRdsRelationInfoResponseBodyData {
	s.RdsInstanceId = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) SetReadOnlyInstanceInfo(v []*string) *GetDrdsDbRdsRelationInfoResponseBodyData {
	s.ReadOnlyInstanceInfo = v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) SetUsedInstanceId(v string) *GetDrdsDbRdsRelationInfoResponseBodyData {
	s.UsedInstanceId = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) SetUsedInstanceType(v string) *GetDrdsDbRdsRelationInfoResponseBodyData {
	s.UsedInstanceType = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
