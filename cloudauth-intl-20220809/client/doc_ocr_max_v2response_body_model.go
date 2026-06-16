// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrMaxV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DocOcrMaxV2ResponseBody
	GetCode() *string
	SetMessage(v string) *DocOcrMaxV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *DocOcrMaxV2ResponseBody
	GetRequestId() *string
	SetResult(v *DocOcrMaxV2ResponseBodyResult) *DocOcrMaxV2ResponseBody
	GetResult() *DocOcrMaxV2ResponseBodyResult
}

type DocOcrMaxV2ResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5E63B760-0ECB-5C07-8503-A65C27876968
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DocOcrMaxV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DocOcrMaxV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DocOcrMaxV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DocOcrMaxV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *DocOcrMaxV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DocOcrMaxV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DocOcrMaxV2ResponseBody) GetResult() *DocOcrMaxV2ResponseBodyResult {
	return s.Result
}

func (s *DocOcrMaxV2ResponseBody) SetCode(v string) *DocOcrMaxV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DocOcrMaxV2ResponseBody) SetMessage(v string) *DocOcrMaxV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DocOcrMaxV2ResponseBody) SetRequestId(v string) *DocOcrMaxV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocOcrMaxV2ResponseBody) SetResult(v *DocOcrMaxV2ResponseBodyResult) *DocOcrMaxV2ResponseBody {
	s.Result = v
	return s
}

func (s *DocOcrMaxV2ResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DocOcrMaxV2ResponseBodyResult struct {
	// example:
	//
	// {
	//
	//   "ocrIdInfo": {
	//
	//     "id_number": "*****719******",
	//
	//     "address": "xxxxxx,
	//
	//     "ethnicity": "汉",
	//
	//     "date_of_birth": "1990年06月02日",
	//
	//     "sex": "女",
	//
	//     "name": "何**"
	//
	//   },
	//
	//   "ocrStandardData": {
	//
	//     "given_name_s": "**",
	//
	//     "surname_s": "HE",
	//
	//     "date_of_birth_s": "1990-06-02",
	//
	//     "sex_s": "F"
	//
	//   }
	//
	// }
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

func (s DocOcrMaxV2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DocOcrMaxV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DocOcrMaxV2ResponseBodyResult) GetExtIdInfo() *string {
	return s.ExtIdInfo
}

func (s *DocOcrMaxV2ResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *DocOcrMaxV2ResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *DocOcrMaxV2ResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *DocOcrMaxV2ResponseBodyResult) SetExtIdInfo(v string) *DocOcrMaxV2ResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *DocOcrMaxV2ResponseBodyResult) SetPassed(v string) *DocOcrMaxV2ResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *DocOcrMaxV2ResponseBodyResult) SetSubCode(v string) *DocOcrMaxV2ResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *DocOcrMaxV2ResponseBodyResult) SetTransactionId(v string) *DocOcrMaxV2ResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *DocOcrMaxV2ResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
