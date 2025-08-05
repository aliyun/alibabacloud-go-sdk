// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTunnelMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeList(v []*int32) *QueryTunnelMetricRequest
	GetCodeList() []*int32
	SetGroupList(v []*string) *QueryTunnelMetricRequest
	GetGroupList() []*string
	SetOperationList(v []*string) *QueryTunnelMetricRequest
	GetOperationList() []*string
	SetProject(v string) *QueryTunnelMetricRequest
	GetProject() *string
	SetQuotaNickname(v string) *QueryTunnelMetricRequest
	GetQuotaNickname() *string
	SetTableList(v []*string) *QueryTunnelMetricRequest
	GetTableList() []*string
	SetTopN(v int32) *QueryTunnelMetricRequest
	GetTopN() *int32
	SetEndTime(v int64) *QueryTunnelMetricRequest
	GetEndTime() *int64
	SetStartTime(v int64) *QueryTunnelMetricRequest
	GetStartTime() *int64
	SetStrategy(v string) *QueryTunnelMetricRequest
	GetStrategy() *string
}

type QueryTunnelMetricRequest struct {
	CodeList      []*int32  `json:"codeList,omitempty" xml:"codeList,omitempty" type:"Repeated"`
	GroupList     []*string `json:"groupList,omitempty" xml:"groupList,omitempty" type:"Repeated"`
	OperationList []*string `json:"operationList,omitempty" xml:"operationList,omitempty" type:"Repeated"`
	// example:
	//
	// project_a
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// default
	QuotaNickname *string   `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	TableList     []*string `json:"tableList,omitempty" xml:"tableList,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TopN *int32 `json:"topN,omitempty" xml:"topN,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1735536322
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1735534322
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// max
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
}

func (s QueryTunnelMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTunnelMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryTunnelMetricRequest) GetCodeList() []*int32 {
	return s.CodeList
}

func (s *QueryTunnelMetricRequest) GetGroupList() []*string {
	return s.GroupList
}

func (s *QueryTunnelMetricRequest) GetOperationList() []*string {
	return s.OperationList
}

func (s *QueryTunnelMetricRequest) GetProject() *string {
	return s.Project
}

func (s *QueryTunnelMetricRequest) GetQuotaNickname() *string {
	return s.QuotaNickname
}

func (s *QueryTunnelMetricRequest) GetTableList() []*string {
	return s.TableList
}

func (s *QueryTunnelMetricRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *QueryTunnelMetricRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryTunnelMetricRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryTunnelMetricRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *QueryTunnelMetricRequest) SetCodeList(v []*int32) *QueryTunnelMetricRequest {
	s.CodeList = v
	return s
}

func (s *QueryTunnelMetricRequest) SetGroupList(v []*string) *QueryTunnelMetricRequest {
	s.GroupList = v
	return s
}

func (s *QueryTunnelMetricRequest) SetOperationList(v []*string) *QueryTunnelMetricRequest {
	s.OperationList = v
	return s
}

func (s *QueryTunnelMetricRequest) SetProject(v string) *QueryTunnelMetricRequest {
	s.Project = &v
	return s
}

func (s *QueryTunnelMetricRequest) SetQuotaNickname(v string) *QueryTunnelMetricRequest {
	s.QuotaNickname = &v
	return s
}

func (s *QueryTunnelMetricRequest) SetTableList(v []*string) *QueryTunnelMetricRequest {
	s.TableList = v
	return s
}

func (s *QueryTunnelMetricRequest) SetTopN(v int32) *QueryTunnelMetricRequest {
	s.TopN = &v
	return s
}

func (s *QueryTunnelMetricRequest) SetEndTime(v int64) *QueryTunnelMetricRequest {
	s.EndTime = &v
	return s
}

func (s *QueryTunnelMetricRequest) SetStartTime(v int64) *QueryTunnelMetricRequest {
	s.StartTime = &v
	return s
}

func (s *QueryTunnelMetricRequest) SetStrategy(v string) *QueryTunnelMetricRequest {
	s.Strategy = &v
	return s
}

func (s *QueryTunnelMetricRequest) Validate() error {
	return dara.Validate(s)
}
