// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataFlowSubTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeDataFlowSubTasksRequest
	GetFileSystemId() *string
	SetFilters(v []*DescribeDataFlowSubTasksRequestFilters) *DescribeDataFlowSubTasksRequest
	GetFilters() []*DescribeDataFlowSubTasksRequestFilters
	SetMaxResults(v int64) *DescribeDataFlowSubTasksRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeDataFlowSubTasksRequest
	GetNextToken() *string
}

type DescribeDataFlowSubTasksRequest struct {
	// The ID of the file system. Only CPFS for Lingjun file systems are supported. The ID must start with `bmcpfs-`.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-370lx1ev9ss27o0****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The filter keys used to query streaming tasks of a data flow.
	//
	// if can be null:
	// false
	Filters []*DescribeDataFlowSubTasksRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 20 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results when the returned results are truncated. You can use NextToken to initiate another request to retrieve the results after the current truncation point.
	//
	// example:
	//
	// iWk0AQAAAAAvY2FzZS8=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDataFlowSubTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowSubTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowSubTasksRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeDataFlowSubTasksRequest) GetFilters() []*DescribeDataFlowSubTasksRequestFilters {
	return s.Filters
}

func (s *DescribeDataFlowSubTasksRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeDataFlowSubTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDataFlowSubTasksRequest) SetFileSystemId(v string) *DescribeDataFlowSubTasksRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDataFlowSubTasksRequest) SetFilters(v []*DescribeDataFlowSubTasksRequestFilters) *DescribeDataFlowSubTasksRequest {
	s.Filters = v
	return s
}

func (s *DescribeDataFlowSubTasksRequest) SetMaxResults(v int64) *DescribeDataFlowSubTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDataFlowSubTasksRequest) SetNextToken(v string) *DescribeDataFlowSubTasksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDataFlowSubTasksRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataFlowSubTasksRequestFilters struct {
	// The name of the filter key.
	//
	// Valid values:
	//
	// - DataFlowIds: filters by data flow ID.
	//
	// - DataFlowTaskIds: filters by data flow task ID.
	//
	// - DataFlowSubTaskIds: filters by data flow streaming task ID.
	//
	// - Status: filters by data flow status.
	//
	// - SrcFilePath: filters by source file path.
	//
	// - DstFilePath: filters by destination file path.
	//
	// example:
	//
	// DataFlowSubTaskIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter key. This parameter does not support wildcards.
	//
	// - If Key is set to DataFlowIds, Value is set to a data flow ID or a part of a data flow ID. You can specify one or more data flow IDs. A maximum of 10 data flow IDs can be specified. Example: `df-194433a5be31****` or `df-194433a512a2****,df-234533a5be31****`.
	//
	// - If Key is set to DataFlowTaskIds, Value is set to a data flow task ID or a part of a data flow task ID. You can specify one or more data flow task IDs. A maximum of 10 data flow task IDs can be specified. Example: `task-29ee8e890f45****` or `task-29ee8e890f45****,task-38ae8e890f45****`.
	//
	// - If Key is set to DataFlowSubTaskIds, Value is set to a data flow streaming task ID or a part of a data flow streaming task ID. You can specify one or more data flow streaming task IDs. A maximum of 10 data flow streaming task IDs can be specified. Example: `subTaskId-370kyfmyknxcyzw****` or `subTaskId-247kyfmyknxcyzw****,subTaskId-256kyfmyknxcyzw****`.
	//
	// - If Key is set to Status, Value is set to the status of the data flow task. Valid values: EXPIRED, CREATED, RUNNING, COMPLETE, CANCELING, FAILED, and CANCELED. Combined queries are supported.
	//
	// - If Key is set to SrcFilePath, Value is set to the source file path. The path can be up to 1,023 characters in length.
	//
	// - If Key is set to DstFilePath, Value is set to the destination file path. The path can be up to 1,023 characters in length.
	//
	// example:
	//
	// subTaskId-370kyfmyknxcyzw****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataFlowSubTasksRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowSubTasksRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowSubTasksRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeDataFlowSubTasksRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeDataFlowSubTasksRequestFilters) SetKey(v string) *DescribeDataFlowSubTasksRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeDataFlowSubTasksRequestFilters) SetValue(v string) *DescribeDataFlowSubTasksRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeDataFlowSubTasksRequestFilters) Validate() error {
	return dara.Validate(s)
}
