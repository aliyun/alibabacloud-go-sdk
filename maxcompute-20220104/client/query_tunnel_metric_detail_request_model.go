// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTunnelMetricDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAscOrder(v bool) *QueryTunnelMetricDetailRequest
	GetAscOrder() *bool
	SetCodeList(v []*int64) *QueryTunnelMetricDetailRequest
	GetCodeList() []*int64
	SetGroupList(v []*string) *QueryTunnelMetricDetailRequest
	GetGroupList() []*string
	SetLimit(v int64) *QueryTunnelMetricDetailRequest
	GetLimit() *int64
	SetOperationList(v []*string) *QueryTunnelMetricDetailRequest
	GetOperationList() []*string
	SetOrderColumn(v string) *QueryTunnelMetricDetailRequest
	GetOrderColumn() *string
	SetProject(v string) *QueryTunnelMetricDetailRequest
	GetProject() *string
	SetQuotaNickname(v string) *QueryTunnelMetricDetailRequest
	GetQuotaNickname() *string
	SetTableList(v []*string) *QueryTunnelMetricDetailRequest
	GetTableList() []*string
	SetEndTime(v int64) *QueryTunnelMetricDetailRequest
	GetEndTime() *int64
	SetStartTime(v int64) *QueryTunnelMetricDetailRequest
	GetStartTime() *int64
}

type QueryTunnelMetricDetailRequest struct {
	// Specifies whether to sort the results in ascending order. If you set this parameter to `true`, the results are sorted in ascending order. If you set this to `false` or leave it unspecified, the results are sorted in descending order.
	//
	// example:
	//
	// false
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// A list of error codes to filter the results.
	CodeList []*int64 `json:"codeList,omitempty" xml:"codeList,omitempty" type:"Repeated"`
	// The list of grouping dimensions.
	GroupList []*string `json:"groupList,omitempty" xml:"groupList,omitempty" type:"Repeated"`
	// The maximum number of entries to return. Default: 10. Maximum: 100.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// A list of operation types.
	OperationList []*string `json:"operationList,omitempty" xml:"operationList,omitempty" type:"Repeated"`
	// The sort column.
	//
	// example:
	//
	// maxValue
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// project_a
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The nickname of the level-2 tunnel quota.
	//
	// The nickname for a shared quota is `default`.
	//
	// For an exclusive quota, the nickname is in the `quotaNickname#subQuotaNickname` format.
	//
	// example:
	//
	// quota_A
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	// A list of table names.
	//
	// The `project` parameter is required if you specify a list of tables. All tables must belong to the specified project.
	TableList []*string `json:"tableList,omitempty" xml:"tableList,omitempty" type:"Repeated"`
	// The end of the query time range. This is a Unix timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1735536322
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start of the query time range. This is a Unix timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1735534322
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s QueryTunnelMetricDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTunnelMetricDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryTunnelMetricDetailRequest) GetAscOrder() *bool {
	return s.AscOrder
}

func (s *QueryTunnelMetricDetailRequest) GetCodeList() []*int64 {
	return s.CodeList
}

func (s *QueryTunnelMetricDetailRequest) GetGroupList() []*string {
	return s.GroupList
}

func (s *QueryTunnelMetricDetailRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *QueryTunnelMetricDetailRequest) GetOperationList() []*string {
	return s.OperationList
}

func (s *QueryTunnelMetricDetailRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QueryTunnelMetricDetailRequest) GetProject() *string {
	return s.Project
}

func (s *QueryTunnelMetricDetailRequest) GetQuotaNickname() *string {
	return s.QuotaNickname
}

func (s *QueryTunnelMetricDetailRequest) GetTableList() []*string {
	return s.TableList
}

func (s *QueryTunnelMetricDetailRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryTunnelMetricDetailRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryTunnelMetricDetailRequest) SetAscOrder(v bool) *QueryTunnelMetricDetailRequest {
	s.AscOrder = &v
	return s
}

func (s *QueryTunnelMetricDetailRequest) SetCodeList(v []*int64) *QueryTunnelMetricDetailRequest {
	s.CodeList = v
	return s
}

func (s *QueryTunnelMetricDetailRequest) SetGroupList(v []*string) *QueryTunnelMetricDetailRequest {
	s.GroupList = v
	return s
}

func (s *QueryTunnelMetricDetailRequest) SetLimit(v int64) *QueryTunnelMetricDetailRequest {
	s.Limit = &v
	return s
}

func (s *QueryTunnelMetricDetailRequest) SetOperationList(v []*string) *QueryTunnelMetricDetailRequest {
	s.OperationList = v
	return s
}

func (s *QueryTunnelMetricDetailRequest) SetOrderColumn(v string) *QueryTunnelMetricDetailRequest {
	s.OrderColumn = &v
	return s
}

func (s *QueryTunnelMetricDetailRequest) SetProject(v string) *QueryTunnelMetricDetailRequest {
	s.Project = &v
	return s
}

func (s *QueryTunnelMetricDetailRequest) SetQuotaNickname(v string) *QueryTunnelMetricDetailRequest {
	s.QuotaNickname = &v
	return s
}

func (s *QueryTunnelMetricDetailRequest) SetTableList(v []*string) *QueryTunnelMetricDetailRequest {
	s.TableList = v
	return s
}

func (s *QueryTunnelMetricDetailRequest) SetEndTime(v int64) *QueryTunnelMetricDetailRequest {
	s.EndTime = &v
	return s
}

func (s *QueryTunnelMetricDetailRequest) SetStartTime(v int64) *QueryTunnelMetricDetailRequest {
	s.StartTime = &v
	return s
}

func (s *QueryTunnelMetricDetailRequest) Validate() error {
	return dara.Validate(s)
}
