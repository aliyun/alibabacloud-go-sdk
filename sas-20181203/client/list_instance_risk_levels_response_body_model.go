// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRiskLevelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRiskLevels(v []*ListInstanceRiskLevelsResponseBodyInstanceRiskLevels) *ListInstanceRiskLevelsResponseBody
	GetInstanceRiskLevels() []*ListInstanceRiskLevelsResponseBodyInstanceRiskLevels
	SetRequestId(v string) *ListInstanceRiskLevelsResponseBody
	GetRequestId() *string
}

type ListInstanceRiskLevelsResponseBody struct {
	// The risk levels of instances.
	InstanceRiskLevels []*ListInstanceRiskLevelsResponseBodyInstanceRiskLevels `json:"InstanceRiskLevels,omitempty" xml:"InstanceRiskLevels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F5CF78A7-30AA-59DB-847F-13EE3AE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceRiskLevelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRiskLevelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceRiskLevelsResponseBody) GetInstanceRiskLevels() []*ListInstanceRiskLevelsResponseBodyInstanceRiskLevels {
	return s.InstanceRiskLevels
}

func (s *ListInstanceRiskLevelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceRiskLevelsResponseBody) SetInstanceRiskLevels(v []*ListInstanceRiskLevelsResponseBodyInstanceRiskLevels) *ListInstanceRiskLevelsResponseBody {
	s.InstanceRiskLevels = v
	return s
}

func (s *ListInstanceRiskLevelsResponseBody) SetRequestId(v string) *ListInstanceRiskLevelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceRiskLevelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceRiskLevelsResponseBodyInstanceRiskLevels struct {
	// The ID of the server.
	//
	// example:
	//
	// i-m5efigezp50l2cmb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// 	- **none**
	//
	// example:
	//
	// high
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The UUID of the server for which you want to modify the defense rule. You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to query the UUIDs of servers.
	//
	// example:
	//
	// f2d6e901-1004-4ca8-9dae-53ec04a92765
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListInstanceRiskLevelsResponseBodyInstanceRiskLevels) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRiskLevelsResponseBodyInstanceRiskLevels) GoString() string {
	return s.String()
}

func (s *ListInstanceRiskLevelsResponseBodyInstanceRiskLevels) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceRiskLevelsResponseBodyInstanceRiskLevels) GetLevel() *string {
	return s.Level
}

func (s *ListInstanceRiskLevelsResponseBodyInstanceRiskLevels) GetUuid() *string {
	return s.Uuid
}

func (s *ListInstanceRiskLevelsResponseBodyInstanceRiskLevels) SetInstanceId(v string) *ListInstanceRiskLevelsResponseBodyInstanceRiskLevels {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceRiskLevelsResponseBodyInstanceRiskLevels) SetLevel(v string) *ListInstanceRiskLevelsResponseBodyInstanceRiskLevels {
	s.Level = &v
	return s
}

func (s *ListInstanceRiskLevelsResponseBodyInstanceRiskLevels) SetUuid(v string) *ListInstanceRiskLevelsResponseBodyInstanceRiskLevels {
	s.Uuid = &v
	return s
}

func (s *ListInstanceRiskLevelsResponseBodyInstanceRiskLevels) Validate() error {
	return dara.Validate(s)
}
