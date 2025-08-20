// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeSetId(v string) *ListChangeSetsRequest
	GetChangeSetId() *string
	SetChangeSetName(v []*string) *ListChangeSetsRequest
	GetChangeSetName() []*string
	SetExecutionStatus(v []*string) *ListChangeSetsRequest
	GetExecutionStatus() []*string
	SetPageNumber(v int64) *ListChangeSetsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListChangeSetsRequest
	GetPageSize() *int64
	SetRegionId(v string) *ListChangeSetsRequest
	GetRegionId() *string
	SetStackId(v string) *ListChangeSetsRequest
	GetStackId() *string
	SetStatus(v []*string) *ListChangeSetsRequest
	GetStatus() []*string
}

type ListChangeSetsRequest struct {
	// The ID of the change set. If detailed information about the change set is not required, you can use this parameter to replace the GetChangeSet operation.
	//
	// example:
	//
	// 1f6521a4-05af-4975-afe9-bc4b45ad****
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The name of change set N. Maximum value of N: 5. You can use an asterisk (\\*) as a wildcard for fuzzy search.
	//
	// example:
	//
	// MyChangeSet
	ChangeSetName []*string `json:"ChangeSetName,omitempty" xml:"ChangeSetName,omitempty" type:"Repeated"`
	// The execution status of change set N. Maximum value of N: 5. Valid values:
	//
	// 	- UNAVAILABLE
	//
	// 	- AVAILABLE
	//
	// 	- EXECUTE_IN_PROGRESS
	//
	// 	- EXECUTE_COMPLETE
	//
	// 	- EXECUTE_FAILED
	//
	// 	- OBSOLETE
	//
	// example:
	//
	// AVAILABLE
	ExecutionStatus []*string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty" type:"Repeated"`
	// The page number.\\
	//
	// Pages start from page 1.\\
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Valid values: 1 to 50.\\
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the change set. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The status of change set N. Maximum value of N: 5. Valid values:
	//
	// 	- CREATE_PENDING
	//
	// 	- CREATE_IN_PROGRESS
	//
	// 	- CREATE_COMPLETE
	//
	// 	- CREATE_FAILED
	//
	// 	- DELETE_FAILED
	//
	// 	- DELETE_COMPLETE
	//
	// example:
	//
	// CREATE_COMPLETE
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s ListChangeSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChangeSetsRequest) GoString() string {
	return s.String()
}

func (s *ListChangeSetsRequest) GetChangeSetId() *string {
	return s.ChangeSetId
}

func (s *ListChangeSetsRequest) GetChangeSetName() []*string {
	return s.ChangeSetName
}

func (s *ListChangeSetsRequest) GetExecutionStatus() []*string {
	return s.ExecutionStatus
}

func (s *ListChangeSetsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListChangeSetsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListChangeSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListChangeSetsRequest) GetStackId() *string {
	return s.StackId
}

func (s *ListChangeSetsRequest) GetStatus() []*string {
	return s.Status
}

func (s *ListChangeSetsRequest) SetChangeSetId(v string) *ListChangeSetsRequest {
	s.ChangeSetId = &v
	return s
}

func (s *ListChangeSetsRequest) SetChangeSetName(v []*string) *ListChangeSetsRequest {
	s.ChangeSetName = v
	return s
}

func (s *ListChangeSetsRequest) SetExecutionStatus(v []*string) *ListChangeSetsRequest {
	s.ExecutionStatus = v
	return s
}

func (s *ListChangeSetsRequest) SetPageNumber(v int64) *ListChangeSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListChangeSetsRequest) SetPageSize(v int64) *ListChangeSetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListChangeSetsRequest) SetRegionId(v string) *ListChangeSetsRequest {
	s.RegionId = &v
	return s
}

func (s *ListChangeSetsRequest) SetStackId(v string) *ListChangeSetsRequest {
	s.StackId = &v
	return s
}

func (s *ListChangeSetsRequest) SetStatus(v []*string) *ListChangeSetsRequest {
	s.Status = v
	return s
}

func (s *ListChangeSetsRequest) Validate() error {
	return dara.Validate(s)
}
