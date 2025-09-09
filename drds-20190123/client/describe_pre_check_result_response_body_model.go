// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPreCheckResult(v *DescribePreCheckResultResponseBodyPreCheckResult) *DescribePreCheckResultResponseBody
	GetPreCheckResult() *DescribePreCheckResultResponseBodyPreCheckResult
	SetRequestId(v string) *DescribePreCheckResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribePreCheckResultResponseBody
	GetSuccess() *bool
}

type DescribePreCheckResultResponseBody struct {
	// Indicates the result of the precheck task.
	PreCheckResult *DescribePreCheckResultResponseBodyPreCheckResult `json:"PreCheckResult,omitempty" xml:"PreCheckResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// EED1A59A-CFEA-5CF8-BB4A-090E75B3D05E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePreCheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponseBody) GetPreCheckResult() *DescribePreCheckResultResponseBodyPreCheckResult {
	return s.PreCheckResult
}

func (s *DescribePreCheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePreCheckResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePreCheckResultResponseBody) SetPreCheckResult(v *DescribePreCheckResultResponseBodyPreCheckResult) *DescribePreCheckResultResponseBody {
	s.PreCheckResult = v
	return s
}

func (s *DescribePreCheckResultResponseBody) SetRequestId(v string) *DescribePreCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePreCheckResultResponseBody) SetSuccess(v bool) *DescribePreCheckResultResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePreCheckResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckResultResponseBodyPreCheckResult struct {
	// Indicates the name of the precheck task.
	PreCheckName *string `json:"PreCheckName,omitempty" xml:"PreCheckName,omitempty"`
	// Indicates the state of the precheck task.
	//
	// example:
	//
	// FAIL
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Indicates the details about the subtasks of the precheck task.
	SubCheckItems []*DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems `json:"SubCheckItems,omitempty" xml:"SubCheckItems,omitempty" type:"Repeated"`
}

func (s DescribePreCheckResultResponseBodyPreCheckResult) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckResultResponseBodyPreCheckResult) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) GetPreCheckName() *string {
	return s.PreCheckName
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) GetState() *string {
	return s.State
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) GetSubCheckItems() []*DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	return s.SubCheckItems
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) SetPreCheckName(v string) *DescribePreCheckResultResponseBodyPreCheckResult {
	s.PreCheckName = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) SetState(v string) *DescribePreCheckResultResponseBodyPreCheckResult {
	s.State = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) SetSubCheckItems(v []*DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) *DescribePreCheckResultResponseBodyPreCheckResult {
	s.SubCheckItems = v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResult) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems struct {
	// Indicates the error code that is returned by a subtask.
	//
	// example:
	//
	// 1004
	ErrorMsgCode *string `json:"ErrorMsgCode,omitempty" xml:"ErrorMsgCode,omitempty"`
	// Indicates an error message.
	ErrorMsgParams []*string `json:"ErrorMsgParams,omitempty" xml:"ErrorMsgParams,omitempty" type:"Repeated"`
	// Indicates the name of the subtask.
	PreCheckItemName *string `json:"PreCheckItemName,omitempty" xml:"PreCheckItemName,omitempty"`
	// Indicates the state of the subtask.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) GetErrorMsgCode() *string {
	return s.ErrorMsgCode
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) GetErrorMsgParams() []*string {
	return s.ErrorMsgParams
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) GetPreCheckItemName() *string {
	return s.PreCheckItemName
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) GetState() *string {
	return s.State
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetErrorMsgCode(v string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.ErrorMsgCode = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetErrorMsgParams(v []*string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.ErrorMsgParams = v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetPreCheckItemName(v string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.PreCheckItemName = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) SetState(v string) *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems {
	s.State = &v
	return s
}

func (s *DescribePreCheckResultResponseBodyPreCheckResultSubCheckItems) Validate() error {
	return dara.Validate(s)
}
