// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *ListOperationCheckRequest
	GetCheckId() *int64
	SetEndTime(v int64) *ListOperationCheckRequest
	GetEndTime() *int64
	SetLang(v string) *ListOperationCheckRequest
	GetLang() *string
	SetOperationTaskInstances(v []*ListOperationCheckRequestOperationTaskInstances) *ListOperationCheckRequest
	GetOperationTaskInstances() []*ListOperationCheckRequestOperationTaskInstances
	SetStartTime(v int64) *ListOperationCheckRequest
	GetStartTime() *int64
	SetType(v string) *ListOperationCheckRequest
	GetType() *string
}

type ListOperationCheckRequest struct {
	// Check item ID.
	//
	// > Obtain this parameter by calling the [ListCheckResult](~~ListCheckResult~~) interface.
	//
	// example:
	//
	// 23
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// Timestamp (in milliseconds) of the end time of the queried task.
	//
	// example:
	//
	// 1719923175001
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Language type for request and response messages, default value is zh. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Information about the operated instances.
	//
	// This parameter is required.
	OperationTaskInstances []*ListOperationCheckRequestOperationTaskInstances `json:"OperationTaskInstances,omitempty" xml:"OperationTaskInstances,omitempty" type:"Repeated"`
	// Timestamp (in milliseconds) of the start time of the queried task.
	//
	// example:
	//
	// 1719923175000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Task type corresponding to the task:
	//
	// - **REPAIR**: Repair task
	//
	// - **ROLLBACK**: Rollback task
	//
	// This parameter is required.
	//
	// example:
	//
	// REPAIR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListOperationCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationCheckRequest) GoString() string {
	return s.String()
}

func (s *ListOperationCheckRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListOperationCheckRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListOperationCheckRequest) GetLang() *string {
	return s.Lang
}

func (s *ListOperationCheckRequest) GetOperationTaskInstances() []*ListOperationCheckRequestOperationTaskInstances {
	return s.OperationTaskInstances
}

func (s *ListOperationCheckRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListOperationCheckRequest) GetType() *string {
	return s.Type
}

func (s *ListOperationCheckRequest) SetCheckId(v int64) *ListOperationCheckRequest {
	s.CheckId = &v
	return s
}

func (s *ListOperationCheckRequest) SetEndTime(v int64) *ListOperationCheckRequest {
	s.EndTime = &v
	return s
}

func (s *ListOperationCheckRequest) SetLang(v string) *ListOperationCheckRequest {
	s.Lang = &v
	return s
}

func (s *ListOperationCheckRequest) SetOperationTaskInstances(v []*ListOperationCheckRequestOperationTaskInstances) *ListOperationCheckRequest {
	s.OperationTaskInstances = v
	return s
}

func (s *ListOperationCheckRequest) SetStartTime(v int64) *ListOperationCheckRequest {
	s.StartTime = &v
	return s
}

func (s *ListOperationCheckRequest) SetType(v string) *ListOperationCheckRequest {
	s.Type = &v
	return s
}

func (s *ListOperationCheckRequest) Validate() error {
	if s.OperationTaskInstances != nil {
		for _, item := range s.OperationTaskInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperationCheckRequestOperationTaskInstances struct {
	// Cloud asset instance ID.
	//
	// example:
	//
	// r-bp1642ib4bg2bm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Asset vendor. Values:
	//
	// - **ALIYUN**: Alibaba Cloud
	//
	// example:
	//
	// ALIYUN
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListOperationCheckRequestOperationTaskInstances) String() string {
	return dara.Prettify(s)
}

func (s ListOperationCheckRequestOperationTaskInstances) GoString() string {
	return s.String()
}

func (s *ListOperationCheckRequestOperationTaskInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationCheckRequestOperationTaskInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOperationCheckRequestOperationTaskInstances) GetVendor() *string {
	return s.Vendor
}

func (s *ListOperationCheckRequestOperationTaskInstances) SetInstanceId(v string) *ListOperationCheckRequestOperationTaskInstances {
	s.InstanceId = &v
	return s
}

func (s *ListOperationCheckRequestOperationTaskInstances) SetRegionId(v string) *ListOperationCheckRequestOperationTaskInstances {
	s.RegionId = &v
	return s
}

func (s *ListOperationCheckRequestOperationTaskInstances) SetVendor(v string) *ListOperationCheckRequestOperationTaskInstances {
	s.Vendor = &v
	return s
}

func (s *ListOperationCheckRequestOperationTaskInstances) Validate() error {
	return dara.Validate(s)
}
