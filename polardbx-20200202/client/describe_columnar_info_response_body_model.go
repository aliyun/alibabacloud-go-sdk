// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnarInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeColumnarInfoResponseBodyData) *DescribeColumnarInfoResponseBody
	GetData() *DescribeColumnarInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeColumnarInfoResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeColumnarInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeColumnarInfoResponseBody
	GetSuccess() *bool
}

type DescribeColumnarInfoResponseBody struct {
	Data *DescribeColumnarInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 14036EBE-****-44DB-ACE9-AC6157D3A6FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeColumnarInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColumnarInfoResponseBody) GetData() *DescribeColumnarInfoResponseBodyData {
	return s.Data
}

func (s *DescribeColumnarInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeColumnarInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeColumnarInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeColumnarInfoResponseBody) SetData(v *DescribeColumnarInfoResponseBodyData) *DescribeColumnarInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeColumnarInfoResponseBody) SetHttpStatusCode(v int32) *DescribeColumnarInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeColumnarInfoResponseBody) SetRequestId(v string) *DescribeColumnarInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColumnarInfoResponseBody) SetSuccess(v bool) *DescribeColumnarInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeColumnarInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeColumnarInfoResponseBodyData struct {
	// example:
	//
	// 30
	BinlogPersistTime *int32 `json:"BinlogPersistTime,omitempty" xml:"BinlogPersistTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 524288000
	BinlogSize *int32 `json:"BinlogSize,omitempty" xml:"BinlogSize,omitempty"`
	// example:
	//
	// ON
	CheckSumSwitch *string `json:"CheckSumSwitch,omitempty" xml:"CheckSumSwitch,omitempty"`
	// example:
	//
	// polarx.n8.medium.col
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// polarx-col-kernel-5.4.20-20250819_17555906
	ColumnarNewVersion *string `json:"ColumnarNewVersion,omitempty" xml:"ColumnarNewVersion,omitempty"`
	// example:
	//
	// polarx-col-kernel-5.4.20-20250819_17555906
	ColumnarVersion      *string                                                     `json:"ColumnarVersion,omitempty" xml:"ColumnarVersion,omitempty"`
	InstanceTopologyList []*DescribeColumnarInfoResponseBodyDataInstanceTopologyList `json:"InstanceTopologyList,omitempty" xml:"InstanceTopologyList,omitempty" type:"Repeated"`
	// server id
	//
	// This parameter is required.
	//
	// example:
	//
	// 53755321
	ServerId *int32 `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
}

func (s DescribeColumnarInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeColumnarInfoResponseBodyData) GetBinlogPersistTime() *int32 {
	return s.BinlogPersistTime
}

func (s *DescribeColumnarInfoResponseBodyData) GetBinlogSize() *int32 {
	return s.BinlogSize
}

func (s *DescribeColumnarInfoResponseBodyData) GetCheckSumSwitch() *string {
	return s.CheckSumSwitch
}

func (s *DescribeColumnarInfoResponseBodyData) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeColumnarInfoResponseBodyData) GetColumnarNewVersion() *string {
	return s.ColumnarNewVersion
}

func (s *DescribeColumnarInfoResponseBodyData) GetColumnarVersion() *string {
	return s.ColumnarVersion
}

func (s *DescribeColumnarInfoResponseBodyData) GetInstanceTopologyList() []*DescribeColumnarInfoResponseBodyDataInstanceTopologyList {
	return s.InstanceTopologyList
}

func (s *DescribeColumnarInfoResponseBodyData) GetServerId() *int32 {
	return s.ServerId
}

func (s *DescribeColumnarInfoResponseBodyData) SetBinlogPersistTime(v int32) *DescribeColumnarInfoResponseBodyData {
	s.BinlogPersistTime = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyData) SetBinlogSize(v int32) *DescribeColumnarInfoResponseBodyData {
	s.BinlogSize = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyData) SetCheckSumSwitch(v string) *DescribeColumnarInfoResponseBodyData {
	s.CheckSumSwitch = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyData) SetClassCode(v string) *DescribeColumnarInfoResponseBodyData {
	s.ClassCode = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyData) SetColumnarNewVersion(v string) *DescribeColumnarInfoResponseBodyData {
	s.ColumnarNewVersion = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyData) SetColumnarVersion(v string) *DescribeColumnarInfoResponseBodyData {
	s.ColumnarVersion = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyData) SetInstanceTopologyList(v []*DescribeColumnarInfoResponseBodyDataInstanceTopologyList) *DescribeColumnarInfoResponseBodyData {
	s.InstanceTopologyList = v
	return s
}

func (s *DescribeColumnarInfoResponseBodyData) SetServerId(v int32) *DescribeColumnarInfoResponseBodyData {
	s.ServerId = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyData) Validate() error {
	if s.InstanceTopologyList != nil {
		for _, item := range s.InstanceTopologyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeColumnarInfoResponseBodyDataInstanceTopologyList struct {
	// example:
	//
	// ***
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// pxc-***
	InstanceName  *string                                                                  `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PhysicalNodes []*DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes `json:"PhysicalNodes,omitempty" xml:"PhysicalNodes,omitempty" type:"Repeated"`
}

func (s DescribeColumnarInfoResponseBodyDataInstanceTopologyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarInfoResponseBodyDataInstanceTopologyList) GoString() string {
	return s.String()
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyList) GetComment() *string {
	return s.Comment
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyList) GetPhysicalNodes() []*DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	return s.PhysicalNodes
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyList) SetComment(v string) *DescribeColumnarInfoResponseBodyDataInstanceTopologyList {
	s.Comment = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyList) SetInstanceName(v string) *DescribeColumnarInfoResponseBodyDataInstanceTopologyList {
	s.InstanceName = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyList) SetPhysicalNodes(v []*DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) *DescribeColumnarInfoResponseBodyDataInstanceTopologyList {
	s.PhysicalNodes = v
	return s
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyList) Validate() error {
	if s.PhysicalNodes != nil {
		for _, item := range s.PhysicalNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes struct {
	// example:
	//
	// cn-hangzhou-h
	AZone *string `json:"AZone,omitempty" xml:"AZone,omitempty"`
	// example:
	//
	// 204800
	Disk *int32 `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// example:
	//
	// polarx.n8.medium.col
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// example:
	//
	// ***
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// polarx-col-kernel-5.4.20-20250819_17555906
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GoString() string {
	return s.String()
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetAZone() *string {
	return s.AZone
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetDisk() *int32 {
	return s.Disk
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetStatus() *string {
	return s.Status
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetVersion() *string {
	return s.Version
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetAZone(v string) *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.AZone = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetDisk(v int32) *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.Disk = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetNodeClass(v string) *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.NodeClass = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetNodeId(v string) *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.NodeId = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetStatus(v string) *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.Status = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetVersion(v string) *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.Version = &v
	return s
}

func (s *DescribeColumnarInfoResponseBodyDataInstanceTopologyListPhysicalNodes) Validate() error {
	return dara.Validate(s)
}
