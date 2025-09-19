// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPublishBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchName(v string) *AddPublishBatchRequest
	GetBatchName() *string
	SetInterval(v int32) *AddPublishBatchRequest
	GetInterval() *int32
	SetOperationBase(v int32) *AddPublishBatchRequest
	GetOperationBase() *int32
	SetUpgradeVersion(v string) *AddPublishBatchRequest
	GetUpgradeVersion() *string
}

type AddPublishBatchRequest struct {
	// The name of the release batch.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BatchName *string `json:"BatchName,omitempty" xml:"BatchName,omitempty"`
	// The interval between two release batches.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The dimension based on which the asset is selected. Valid values:
	//
	// 	- **0**: selects the asset by instance.
	//
	// 	- **1**: selects the asset by machine group.
	//
	// 	- **2**: selects the asset by the ID of the instance that is deployed in the virtual private cloud (VPC).
	//
	// example:
	//
	// 0
	OperationBase *int32 `json:"OperationBase,omitempty" xml:"OperationBase,omitempty"`
	// The version to which you want to upgrade the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.9
	UpgradeVersion *string `json:"UpgradeVersion,omitempty" xml:"UpgradeVersion,omitempty"`
}

func (s AddPublishBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPublishBatchRequest) GoString() string {
	return s.String()
}

func (s *AddPublishBatchRequest) GetBatchName() *string {
	return s.BatchName
}

func (s *AddPublishBatchRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *AddPublishBatchRequest) GetOperationBase() *int32 {
	return s.OperationBase
}

func (s *AddPublishBatchRequest) GetUpgradeVersion() *string {
	return s.UpgradeVersion
}

func (s *AddPublishBatchRequest) SetBatchName(v string) *AddPublishBatchRequest {
	s.BatchName = &v
	return s
}

func (s *AddPublishBatchRequest) SetInterval(v int32) *AddPublishBatchRequest {
	s.Interval = &v
	return s
}

func (s *AddPublishBatchRequest) SetOperationBase(v int32) *AddPublishBatchRequest {
	s.OperationBase = &v
	return s
}

func (s *AddPublishBatchRequest) SetUpgradeVersion(v string) *AddPublishBatchRequest {
	s.UpgradeVersion = &v
	return s
}

func (s *AddPublishBatchRequest) Validate() error {
	return dara.Validate(s)
}
