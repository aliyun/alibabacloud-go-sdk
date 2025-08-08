// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcubeVhostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryVhostResult(v *QueryMcubeVhostResponseBodyQueryVhostResult) *QueryMcubeVhostResponseBody
	GetQueryVhostResult() *QueryMcubeVhostResponseBodyQueryVhostResult
	SetRequestId(v string) *QueryMcubeVhostResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryMcubeVhostResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *QueryMcubeVhostResponseBody
	GetResultMessage() *string
}

type QueryMcubeVhostResponseBody struct {
	QueryVhostResult *QueryMcubeVhostResponseBodyQueryVhostResult `json:"QueryVhostResult,omitempty" xml:"QueryVhostResult,omitempty" type:"Struct"`
	RequestId        *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode       *string                                      `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage    *string                                      `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMcubeVhostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeVhostResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMcubeVhostResponseBody) GetQueryVhostResult() *QueryMcubeVhostResponseBodyQueryVhostResult {
	return s.QueryVhostResult
}

func (s *QueryMcubeVhostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMcubeVhostResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryMcubeVhostResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryMcubeVhostResponseBody) SetQueryVhostResult(v *QueryMcubeVhostResponseBodyQueryVhostResult) *QueryMcubeVhostResponseBody {
	s.QueryVhostResult = v
	return s
}

func (s *QueryMcubeVhostResponseBody) SetRequestId(v string) *QueryMcubeVhostResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMcubeVhostResponseBody) SetResultCode(v string) *QueryMcubeVhostResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMcubeVhostResponseBody) SetResultMessage(v string) *QueryMcubeVhostResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMcubeVhostResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMcubeVhostResponseBodyQueryVhostResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMcubeVhostResponseBodyQueryVhostResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeVhostResponseBodyQueryVhostResult) GoString() string {
	return s.String()
}

func (s *QueryMcubeVhostResponseBodyQueryVhostResult) GetData() *string {
	return s.Data
}

func (s *QueryMcubeVhostResponseBodyQueryVhostResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryMcubeVhostResponseBodyQueryVhostResult) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMcubeVhostResponseBodyQueryVhostResult) SetData(v string) *QueryMcubeVhostResponseBodyQueryVhostResult {
	s.Data = &v
	return s
}

func (s *QueryMcubeVhostResponseBodyQueryVhostResult) SetResultMsg(v string) *QueryMcubeVhostResponseBodyQueryVhostResult {
	s.ResultMsg = &v
	return s
}

func (s *QueryMcubeVhostResponseBodyQueryVhostResult) SetSuccess(v bool) *QueryMcubeVhostResponseBodyQueryVhostResult {
	s.Success = &v
	return s
}

func (s *QueryMcubeVhostResponseBodyQueryVhostResult) Validate() error {
	return dara.Validate(s)
}
