// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeProjectMessagesResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeProjectMessagesResponseBodyResult) *DescribeProjectMessagesResponseBody
	GetResult() []*DescribeProjectMessagesResponseBodyResult
	SetSuccess(v bool) *DescribeProjectMessagesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeProjectMessagesResponseBody
	GetTotalCount() *int64
}

type DescribeProjectMessagesResponseBody struct {
	// example:
	//
	// 00eb4de1-6cff-4f56-833e-7b1e070e398d
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeProjectMessagesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 28
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProjectMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProjectMessagesResponseBody) GetResult() []*DescribeProjectMessagesResponseBodyResult {
	return s.Result
}

func (s *DescribeProjectMessagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeProjectMessagesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeProjectMessagesResponseBody) SetRequestId(v string) *DescribeProjectMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectMessagesResponseBody) SetResult(v []*DescribeProjectMessagesResponseBodyResult) *DescribeProjectMessagesResponseBody {
	s.Result = v
	return s
}

func (s *DescribeProjectMessagesResponseBody) SetSuccess(v bool) *DescribeProjectMessagesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProjectMessagesResponseBody) SetTotalCount(v int64) *DescribeProjectMessagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeProjectMessagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProjectMessagesResponseBodyResult struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1589015560000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 452611111****
	Operator     *int64  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// example:
	//
	// Provider
	OperatorRole *string `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
}

func (s DescribeProjectMessagesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectMessagesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeProjectMessagesResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *DescribeProjectMessagesResponseBodyResult) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeProjectMessagesResponseBodyResult) GetOperator() *int64 {
	return s.Operator
}

func (s *DescribeProjectMessagesResponseBodyResult) GetOperatorName() *string {
	return s.OperatorName
}

func (s *DescribeProjectMessagesResponseBodyResult) GetOperatorRole() *string {
	return s.OperatorRole
}

func (s *DescribeProjectMessagesResponseBodyResult) SetContent(v string) *DescribeProjectMessagesResponseBodyResult {
	s.Content = &v
	return s
}

func (s *DescribeProjectMessagesResponseBodyResult) SetGmtCreate(v int64) *DescribeProjectMessagesResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DescribeProjectMessagesResponseBodyResult) SetOperator(v int64) *DescribeProjectMessagesResponseBodyResult {
	s.Operator = &v
	return s
}

func (s *DescribeProjectMessagesResponseBodyResult) SetOperatorName(v string) *DescribeProjectMessagesResponseBodyResult {
	s.OperatorName = &v
	return s
}

func (s *DescribeProjectMessagesResponseBodyResult) SetOperatorRole(v string) *DescribeProjectMessagesResponseBodyResult {
	s.OperatorRole = &v
	return s
}

func (s *DescribeProjectMessagesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
