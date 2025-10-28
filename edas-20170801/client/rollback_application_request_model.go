// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RollbackApplicationRequest
	GetAppId() *string
	SetBatch(v int32) *RollbackApplicationRequest
	GetBatch() *int32
	SetBatchWaitTime(v int32) *RollbackApplicationRequest
	GetBatchWaitTime() *int32
	SetGroupId(v string) *RollbackApplicationRequest
	GetGroupId() *string
	SetHistoryVersion(v string) *RollbackApplicationRequest
	GetHistoryVersion() *string
}

type RollbackApplicationRequest struct {
	// The application ID. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/423162.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-4413-************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The number of batches for the rollback. Default value: 1. Valid values: 1 to 5.
	//
	// example:
	//
	// 1
	Batch *int32 `json:"Batch,omitempty" xml:"Batch,omitempty"`
	// The wait time between batches. Default value: 0. The default value indicates no wait time. Valid values: 0 to 5. Unit: minutes.
	//
	// example:
	//
	// 0
	BatchWaitTime *int32 `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	// The application group ID. You can call the ListDeployGroup operation to query the application group ID. For more information, see [ListDeployGroup](https://help.aliyun.com/document_detail/423184.html).
	//
	// If you need to roll back the application in all application groups, set this parameter to `all`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8123db90-880f-48***************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The historical version to which you want to roll back the application. Call the ListHistoryDeployVersion operation to query the historical versions of the application. Then, set this parameter based on the returned value of `PackageVersion`. For more information, see [ListHistoryDeployVersion](https://help.aliyun.com/document_detail/423163.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-11-13 14:22:22
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
}

func (s RollbackApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackApplicationRequest) GoString() string {
	return s.String()
}

func (s *RollbackApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *RollbackApplicationRequest) GetBatch() *int32 {
	return s.Batch
}

func (s *RollbackApplicationRequest) GetBatchWaitTime() *int32 {
	return s.BatchWaitTime
}

func (s *RollbackApplicationRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RollbackApplicationRequest) GetHistoryVersion() *string {
	return s.HistoryVersion
}

func (s *RollbackApplicationRequest) SetAppId(v string) *RollbackApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RollbackApplicationRequest) SetBatch(v int32) *RollbackApplicationRequest {
	s.Batch = &v
	return s
}

func (s *RollbackApplicationRequest) SetBatchWaitTime(v int32) *RollbackApplicationRequest {
	s.BatchWaitTime = &v
	return s
}

func (s *RollbackApplicationRequest) SetGroupId(v string) *RollbackApplicationRequest {
	s.GroupId = &v
	return s
}

func (s *RollbackApplicationRequest) SetHistoryVersion(v string) *RollbackApplicationRequest {
	s.HistoryVersion = &v
	return s
}

func (s *RollbackApplicationRequest) Validate() error {
	return dara.Validate(s)
}
