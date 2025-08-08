// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectOperateLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeProjectOperateLogsResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeProjectOperateLogsResponseBodyResult) *DescribeProjectOperateLogsResponseBody
	GetResult() []*DescribeProjectOperateLogsResponseBodyResult
	SetSuccess(v bool) *DescribeProjectOperateLogsResponseBody
	GetSuccess() *bool
}

type DescribeProjectOperateLogsResponseBody struct {
	// example:
	//
	// e6037e86-657f-4194-bb6a-7b26973aec8d
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeProjectOperateLogsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProjectOperateLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectOperateLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectOperateLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProjectOperateLogsResponseBody) GetResult() []*DescribeProjectOperateLogsResponseBodyResult {
	return s.Result
}

func (s *DescribeProjectOperateLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeProjectOperateLogsResponseBody) SetRequestId(v string) *DescribeProjectOperateLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectOperateLogsResponseBody) SetResult(v []*DescribeProjectOperateLogsResponseBodyResult) *DescribeProjectOperateLogsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeProjectOperateLogsResponseBody) SetSuccess(v bool) *DescribeProjectOperateLogsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProjectOperateLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProjectOperateLogsResponseBodyResult struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1587624497000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 0
	Operator     *int64  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// example:
	//
	// System
	OperatorRole *string `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
}

func (s DescribeProjectOperateLogsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectOperateLogsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeProjectOperateLogsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeProjectOperateLogsResponseBodyResult) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeProjectOperateLogsResponseBodyResult) GetOperator() *int64 {
	return s.Operator
}

func (s *DescribeProjectOperateLogsResponseBodyResult) GetOperatorName() *string {
	return s.OperatorName
}

func (s *DescribeProjectOperateLogsResponseBodyResult) GetOperatorRole() *string {
	return s.OperatorRole
}

func (s *DescribeProjectOperateLogsResponseBodyResult) SetDescription(v string) *DescribeProjectOperateLogsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeProjectOperateLogsResponseBodyResult) SetGmtCreate(v int64) *DescribeProjectOperateLogsResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DescribeProjectOperateLogsResponseBodyResult) SetOperator(v int64) *DescribeProjectOperateLogsResponseBodyResult {
	s.Operator = &v
	return s
}

func (s *DescribeProjectOperateLogsResponseBodyResult) SetOperatorName(v string) *DescribeProjectOperateLogsResponseBodyResult {
	s.OperatorName = &v
	return s
}

func (s *DescribeProjectOperateLogsResponseBodyResult) SetOperatorRole(v string) *DescribeProjectOperateLogsResponseBodyResult {
	s.OperatorRole = &v
	return s
}

func (s *DescribeProjectOperateLogsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
