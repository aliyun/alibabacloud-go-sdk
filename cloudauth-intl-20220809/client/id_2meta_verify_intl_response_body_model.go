// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaVerifyIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Id2MetaVerifyIntlResponseBody
	GetCode() *string
	SetMessage(v string) *Id2MetaVerifyIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *Id2MetaVerifyIntlResponseBody
	GetRequestId() *string
	SetResult(v *Id2MetaVerifyIntlResponseBodyResult) *Id2MetaVerifyIntlResponseBody
	GetResult() *Id2MetaVerifyIntlResponseBodyResult
}

type Id2MetaVerifyIntlResponseBody struct {
	// [Status codes](https://www.alibabacloud.com/help/en/ekyc/latest/ok4bwxwmu1n94o76?spm=a2c63.p38356.0.i54#942707fca218x).
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed description of the response code.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EFA11401-C961-5E89-A2D3-BF9883E5CC3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result
	Result *Id2MetaVerifyIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s Id2MetaVerifyIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *Id2MetaVerifyIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Id2MetaVerifyIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Id2MetaVerifyIntlResponseBody) GetResult() *Id2MetaVerifyIntlResponseBodyResult {
	return s.Result
}

func (s *Id2MetaVerifyIntlResponseBody) SetCode(v string) *Id2MetaVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *Id2MetaVerifyIntlResponseBody) SetMessage(v string) *Id2MetaVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *Id2MetaVerifyIntlResponseBody) SetRequestId(v string) *Id2MetaVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id2MetaVerifyIntlResponseBody) SetResult(v *Id2MetaVerifyIntlResponseBodyResult) *Id2MetaVerifyIntlResponseBody {
	s.Result = v
	return s
}

func (s *Id2MetaVerifyIntlResponseBody) Validate() error {
	return dara.Validate(s)
}

type Id2MetaVerifyIntlResponseBodyResult struct {
	// The verification result:
	//
	// - 1: The information is consistent. This result is billable.
	//
	// - 2: The information is inconsistent. This result is billable.
	//
	// - 3: No record is found. This result is not billable.
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
}

func (s Id2MetaVerifyIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyIntlResponseBodyResult) GetBizCode() *string {
	return s.BizCode
}

func (s *Id2MetaVerifyIntlResponseBodyResult) SetBizCode(v string) *Id2MetaVerifyIntlResponseBodyResult {
	s.BizCode = &v
	return s
}

func (s *Id2MetaVerifyIntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
