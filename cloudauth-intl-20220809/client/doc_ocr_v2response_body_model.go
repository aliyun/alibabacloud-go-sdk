// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DocOcrV2ResponseBody
	GetCode() *string
	SetMessage(v string) *DocOcrV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *DocOcrV2ResponseBody
	GetRequestId() *string
	SetResult(v *DocOcrV2ResponseBodyResult) *DocOcrV2ResponseBody
	GetResult() *DocOcrV2ResponseBodyResult
}

type DocOcrV2ResponseBody struct {
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
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DocOcrV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DocOcrV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DocOcrV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DocOcrV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *DocOcrV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DocOcrV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DocOcrV2ResponseBody) GetResult() *DocOcrV2ResponseBodyResult {
	return s.Result
}

func (s *DocOcrV2ResponseBody) SetCode(v string) *DocOcrV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DocOcrV2ResponseBody) SetMessage(v string) *DocOcrV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DocOcrV2ResponseBody) SetRequestId(v string) *DocOcrV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocOcrV2ResponseBody) SetResult(v *DocOcrV2ResponseBodyResult) *DocOcrV2ResponseBody {
	s.Result = v
	return s
}

func (s *DocOcrV2ResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DocOcrV2ResponseBodyResult struct {
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

func (s DocOcrV2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DocOcrV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DocOcrV2ResponseBodyResult) GetExtIdInfo() *string {
	return s.ExtIdInfo
}

func (s *DocOcrV2ResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *DocOcrV2ResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *DocOcrV2ResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *DocOcrV2ResponseBodyResult) SetExtIdInfo(v string) *DocOcrV2ResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *DocOcrV2ResponseBodyResult) SetPassed(v string) *DocOcrV2ResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *DocOcrV2ResponseBodyResult) SetSubCode(v string) *DocOcrV2ResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *DocOcrV2ResponseBodyResult) SetTransactionId(v string) *DocOcrV2ResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *DocOcrV2ResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
