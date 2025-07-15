// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DocOcrResponseBody
	GetCode() *string
	SetMessage(v string) *DocOcrResponseBody
	GetMessage() *string
	SetRequestId(v string) *DocOcrResponseBody
	GetRequestId() *string
	SetResult(v *DocOcrResponseBodyResult) *DocOcrResponseBody
	GetResult() *DocOcrResponseBodyResult
}

type DocOcrResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 86C40EC3-5940-5F47-995C-BFE90B70E540
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DocOcrResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DocOcrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DocOcrResponseBody) GoString() string {
	return s.String()
}

func (s *DocOcrResponseBody) GetCode() *string {
	return s.Code
}

func (s *DocOcrResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DocOcrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DocOcrResponseBody) GetResult() *DocOcrResponseBodyResult {
	return s.Result
}

func (s *DocOcrResponseBody) SetCode(v string) *DocOcrResponseBody {
	s.Code = &v
	return s
}

func (s *DocOcrResponseBody) SetMessage(v string) *DocOcrResponseBody {
	s.Message = &v
	return s
}

func (s *DocOcrResponseBody) SetRequestId(v string) *DocOcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocOcrResponseBody) SetResult(v *DocOcrResponseBodyResult) *DocOcrResponseBody {
	s.Result = v
	return s
}

func (s *DocOcrResponseBody) Validate() error {
	return dara.Validate(s)
}

type DocOcrResponseBodyResult struct {
	ExtIdInfo *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s DocOcrResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DocOcrResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DocOcrResponseBodyResult) GetExtIdInfo() *string {
	return s.ExtIdInfo
}

func (s *DocOcrResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *DocOcrResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *DocOcrResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *DocOcrResponseBodyResult) SetExtIdInfo(v string) *DocOcrResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *DocOcrResponseBodyResult) SetPassed(v string) *DocOcrResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *DocOcrResponseBodyResult) SetSubCode(v string) *DocOcrResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *DocOcrResponseBodyResult) SetTransactionId(v string) *DocOcrResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *DocOcrResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
