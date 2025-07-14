// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddFeishuUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchAddFeishuUsersResponseBody
	GetRequestId() *string
	SetResult(v *BatchAddFeishuUsersResponseBodyResult) *BatchAddFeishuUsersResponseBody
	GetResult() *BatchAddFeishuUsersResponseBodyResult
	SetSuccess(v bool) *BatchAddFeishuUsersResponseBody
	GetSuccess() *bool
}

type BatchAddFeishuUsersResponseBody struct {
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Result *BatchAddFeishuUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchAddFeishuUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFeishuUsersResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddFeishuUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchAddFeishuUsersResponseBody) GetResult() *BatchAddFeishuUsersResponseBodyResult {
	return s.Result
}

func (s *BatchAddFeishuUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchAddFeishuUsersResponseBody) SetRequestId(v string) *BatchAddFeishuUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAddFeishuUsersResponseBody) SetResult(v *BatchAddFeishuUsersResponseBodyResult) *BatchAddFeishuUsersResponseBody {
	s.Result = v
	return s
}

func (s *BatchAddFeishuUsersResponseBody) SetSuccess(v bool) *BatchAddFeishuUsersResponseBody {
	s.Success = &v
	return s
}

func (s *BatchAddFeishuUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchAddFeishuUsersResponseBodyResult struct {
	// example:
	//
	// 10
	FailCount   *int32                                              `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	FailResults []*BatchAddFeishuUsersResponseBodyResultFailResults `json:"FailResults,omitempty" xml:"FailResults,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	OkCount *int32 `json:"OkCount,omitempty" xml:"OkCount,omitempty"`
}

func (s BatchAddFeishuUsersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFeishuUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchAddFeishuUsersResponseBodyResult) GetFailCount() *int32 {
	return s.FailCount
}

func (s *BatchAddFeishuUsersResponseBodyResult) GetFailResults() []*BatchAddFeishuUsersResponseBodyResultFailResults {
	return s.FailResults
}

func (s *BatchAddFeishuUsersResponseBodyResult) GetOkCount() *int32 {
	return s.OkCount
}

func (s *BatchAddFeishuUsersResponseBodyResult) SetFailCount(v int32) *BatchAddFeishuUsersResponseBodyResult {
	s.FailCount = &v
	return s
}

func (s *BatchAddFeishuUsersResponseBodyResult) SetFailResults(v []*BatchAddFeishuUsersResponseBodyResultFailResults) *BatchAddFeishuUsersResponseBodyResult {
	s.FailResults = v
	return s
}

func (s *BatchAddFeishuUsersResponseBodyResult) SetOkCount(v int32) *BatchAddFeishuUsersResponseBodyResult {
	s.OkCount = &v
	return s
}

func (s *BatchAddFeishuUsersResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type BatchAddFeishuUsersResponseBodyResultFailResults struct {
	FailInfos []*BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos `json:"FailInfos,omitempty" xml:"FailInfos,omitempty" type:"Repeated"`
}

func (s BatchAddFeishuUsersResponseBodyResultFailResults) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFeishuUsersResponseBodyResultFailResults) GoString() string {
	return s.String()
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResults) GetFailInfos() []*BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos {
	return s.FailInfos
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResults) SetFailInfos(v []*BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) *BatchAddFeishuUsersResponseBodyResultFailResults {
	s.FailInfos = v
	return s
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResults) Validate() error {
	return dara.Validate(s)
}

type BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos struct {
	// example:
	//
	// ACCOUNT_EXIST
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CodeDesc *string `json:"CodeDesc,omitempty" xml:"CodeDesc,omitempty"`
	// example:
	//
	// 20
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
}

func (s BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) GoString() string {
	return s.String()
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) GetCode() *string {
	return s.Code
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) GetCodeDesc() *string {
	return s.CodeDesc
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) GetInput() *string {
	return s.Input
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) SetCode(v string) *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos {
	s.Code = &v
	return s
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) SetCodeDesc(v string) *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos {
	s.CodeDesc = &v
	return s
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) SetInput(v string) *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos {
	s.Input = &v
	return s
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) Validate() error {
	return dara.Validate(s)
}
