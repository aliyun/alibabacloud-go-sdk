// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryExplainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeQueryExplainResponseBody
	GetCode() *string
	SetData(v []*DescribeQueryExplainResponseBodyData) *DescribeQueryExplainResponseBody
	GetData() []*DescribeQueryExplainResponseBodyData
	SetMessage(v string) *DescribeQueryExplainResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeQueryExplainResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeQueryExplainResponseBody
	GetSuccess() *string
}

type DescribeQueryExplainResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List<ExplainedSqlDO>
	Data []*DescribeQueryExplainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeQueryExplainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryExplainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQueryExplainResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeQueryExplainResponseBody) GetData() []*DescribeQueryExplainResponseBodyData {
	return s.Data
}

func (s *DescribeQueryExplainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeQueryExplainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQueryExplainResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeQueryExplainResponseBody) SetCode(v string) *DescribeQueryExplainResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeQueryExplainResponseBody) SetData(v []*DescribeQueryExplainResponseBodyData) *DescribeQueryExplainResponseBody {
	s.Data = v
	return s
}

func (s *DescribeQueryExplainResponseBody) SetMessage(v string) *DescribeQueryExplainResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeQueryExplainResponseBody) SetRequestId(v string) *DescribeQueryExplainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQueryExplainResponseBody) SetSuccess(v string) *DescribeQueryExplainResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeQueryExplainResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeQueryExplainResponseBodyData struct {
	Argument           *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	AvgRowSize         *string `json:"AvgRowSize,omitempty" xml:"AvgRowSize,omitempty"`
	DefinedValues      *string `json:"DefinedValues,omitempty" xml:"DefinedValues,omitempty"`
	EstimateCPU        *string `json:"EstimateCPU,omitempty" xml:"EstimateCPU,omitempty"`
	EstimateExecutions *string `json:"EstimateExecutions,omitempty" xml:"EstimateExecutions,omitempty"`
	EstimateIO         *string `json:"EstimateIO,omitempty" xml:"EstimateIO,omitempty"`
	EstimateRows       *string `json:"EstimateRows,omitempty" xml:"EstimateRows,omitempty"`
	Extra              *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// 1
	Id        *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	IndexList []*string `json:"IndexList,omitempty" xml:"IndexList,omitempty" type:"Repeated"`
	// example:
	//
	// PRIMARY
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 3
	KeyLen          *string   `json:"KeyLen,omitempty" xml:"KeyLen,omitempty"`
	LogicalOp       *string   `json:"LogicalOp,omitempty" xml:"LogicalOp,omitempty"`
	LogicalPlanList []*string `json:"LogicalPlanList,omitempty" xml:"LogicalPlanList,omitempty" type:"Repeated"`
	NodeId          *string   `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OutputList      *string   `json:"OutputList,omitempty" xml:"OutputList,omitempty"`
	Parallel        *string   `json:"Parallel,omitempty" xml:"Parallel,omitempty"`
	Parent          *string   `json:"Parent,omitempty" xml:"Parent,omitempty"`
	PhysicalOp      *string   `json:"PhysicalOp,omitempty" xml:"PhysicalOp,omitempty"`
	// example:
	//
	// test_idx
	PossibleKeys *string `json:"PossibleKeys,omitempty" xml:"PossibleKeys,omitempty"`
	QueryPlan    *string `json:"QueryPlan,omitempty" xml:"QueryPlan,omitempty"`
	// example:
	//
	// test_column
	Ref *string `json:"Ref,omitempty" xml:"Ref,omitempty"`
	// example:
	//
	// 1000
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// example:
	//
	// SIMPLE
	SelectType *string `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	StmtId     *string `json:"StmtId,omitempty" xml:"StmtId,omitempty"`
	StmtText   *string `json:"StmtText,omitempty" xml:"StmtText,omitempty"`
	// example:
	//
	// test
	Table            *string   `json:"Table,omitempty" xml:"Table,omitempty"`
	TableList        []*string `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Repeated"`
	TotalSubtreeCost *string   `json:"TotalSubtreeCost,omitempty" xml:"TotalSubtreeCost,omitempty"`
	// example:
	//
	// eq_ref
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Warnings *string `json:"Warnings,omitempty" xml:"Warnings,omitempty"`
}

func (s DescribeQueryExplainResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryExplainResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeQueryExplainResponseBodyData) GetArgument() *string {
	return s.Argument
}

func (s *DescribeQueryExplainResponseBodyData) GetAvgRowSize() *string {
	return s.AvgRowSize
}

func (s *DescribeQueryExplainResponseBodyData) GetDefinedValues() *string {
	return s.DefinedValues
}

func (s *DescribeQueryExplainResponseBodyData) GetEstimateCPU() *string {
	return s.EstimateCPU
}

func (s *DescribeQueryExplainResponseBodyData) GetEstimateExecutions() *string {
	return s.EstimateExecutions
}

func (s *DescribeQueryExplainResponseBodyData) GetEstimateIO() *string {
	return s.EstimateIO
}

func (s *DescribeQueryExplainResponseBodyData) GetEstimateRows() *string {
	return s.EstimateRows
}

func (s *DescribeQueryExplainResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *DescribeQueryExplainResponseBodyData) GetId() *string {
	return s.Id
}

func (s *DescribeQueryExplainResponseBodyData) GetIndexList() []*string {
	return s.IndexList
}

func (s *DescribeQueryExplainResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *DescribeQueryExplainResponseBodyData) GetKeyLen() *string {
	return s.KeyLen
}

func (s *DescribeQueryExplainResponseBodyData) GetLogicalOp() *string {
	return s.LogicalOp
}

func (s *DescribeQueryExplainResponseBodyData) GetLogicalPlanList() []*string {
	return s.LogicalPlanList
}

func (s *DescribeQueryExplainResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeQueryExplainResponseBodyData) GetOutputList() *string {
	return s.OutputList
}

func (s *DescribeQueryExplainResponseBodyData) GetParallel() *string {
	return s.Parallel
}

func (s *DescribeQueryExplainResponseBodyData) GetParent() *string {
	return s.Parent
}

func (s *DescribeQueryExplainResponseBodyData) GetPhysicalOp() *string {
	return s.PhysicalOp
}

func (s *DescribeQueryExplainResponseBodyData) GetPossibleKeys() *string {
	return s.PossibleKeys
}

func (s *DescribeQueryExplainResponseBodyData) GetQueryPlan() *string {
	return s.QueryPlan
}

func (s *DescribeQueryExplainResponseBodyData) GetRef() *string {
	return s.Ref
}

func (s *DescribeQueryExplainResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *DescribeQueryExplainResponseBodyData) GetSelectType() *string {
	return s.SelectType
}

func (s *DescribeQueryExplainResponseBodyData) GetStmtId() *string {
	return s.StmtId
}

func (s *DescribeQueryExplainResponseBodyData) GetStmtText() *string {
	return s.StmtText
}

func (s *DescribeQueryExplainResponseBodyData) GetTable() *string {
	return s.Table
}

func (s *DescribeQueryExplainResponseBodyData) GetTableList() []*string {
	return s.TableList
}

func (s *DescribeQueryExplainResponseBodyData) GetTotalSubtreeCost() *string {
	return s.TotalSubtreeCost
}

func (s *DescribeQueryExplainResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DescribeQueryExplainResponseBodyData) GetWarnings() *string {
	return s.Warnings
}

func (s *DescribeQueryExplainResponseBodyData) SetArgument(v string) *DescribeQueryExplainResponseBodyData {
	s.Argument = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetAvgRowSize(v string) *DescribeQueryExplainResponseBodyData {
	s.AvgRowSize = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetDefinedValues(v string) *DescribeQueryExplainResponseBodyData {
	s.DefinedValues = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetEstimateCPU(v string) *DescribeQueryExplainResponseBodyData {
	s.EstimateCPU = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetEstimateExecutions(v string) *DescribeQueryExplainResponseBodyData {
	s.EstimateExecutions = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetEstimateIO(v string) *DescribeQueryExplainResponseBodyData {
	s.EstimateIO = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetEstimateRows(v string) *DescribeQueryExplainResponseBodyData {
	s.EstimateRows = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetExtra(v string) *DescribeQueryExplainResponseBodyData {
	s.Extra = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetId(v string) *DescribeQueryExplainResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetIndexList(v []*string) *DescribeQueryExplainResponseBodyData {
	s.IndexList = v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetKey(v string) *DescribeQueryExplainResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetKeyLen(v string) *DescribeQueryExplainResponseBodyData {
	s.KeyLen = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetLogicalOp(v string) *DescribeQueryExplainResponseBodyData {
	s.LogicalOp = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetLogicalPlanList(v []*string) *DescribeQueryExplainResponseBodyData {
	s.LogicalPlanList = v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetNodeId(v string) *DescribeQueryExplainResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetOutputList(v string) *DescribeQueryExplainResponseBodyData {
	s.OutputList = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetParallel(v string) *DescribeQueryExplainResponseBodyData {
	s.Parallel = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetParent(v string) *DescribeQueryExplainResponseBodyData {
	s.Parent = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetPhysicalOp(v string) *DescribeQueryExplainResponseBodyData {
	s.PhysicalOp = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetPossibleKeys(v string) *DescribeQueryExplainResponseBodyData {
	s.PossibleKeys = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetQueryPlan(v string) *DescribeQueryExplainResponseBodyData {
	s.QueryPlan = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetRef(v string) *DescribeQueryExplainResponseBodyData {
	s.Ref = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetRows(v string) *DescribeQueryExplainResponseBodyData {
	s.Rows = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetSelectType(v string) *DescribeQueryExplainResponseBodyData {
	s.SelectType = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetStmtId(v string) *DescribeQueryExplainResponseBodyData {
	s.StmtId = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetStmtText(v string) *DescribeQueryExplainResponseBodyData {
	s.StmtText = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetTable(v string) *DescribeQueryExplainResponseBodyData {
	s.Table = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetTableList(v []*string) *DescribeQueryExplainResponseBodyData {
	s.TableList = v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetTotalSubtreeCost(v string) *DescribeQueryExplainResponseBodyData {
	s.TotalSubtreeCost = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetType(v string) *DescribeQueryExplainResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) SetWarnings(v string) *DescribeQueryExplainResponseBodyData {
	s.Warnings = &v
	return s
}

func (s *DescribeQueryExplainResponseBodyData) Validate() error {
	return dara.Validate(s)
}
