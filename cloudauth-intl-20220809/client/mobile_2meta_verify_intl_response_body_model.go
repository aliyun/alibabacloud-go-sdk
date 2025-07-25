// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile2MetaVerifyIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Mobile2MetaVerifyIntlResponseBody
	GetCode() *string
	SetMessage(v string) *Mobile2MetaVerifyIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *Mobile2MetaVerifyIntlResponseBody
	GetRequestId() *string
	SetResult(v *Mobile2MetaVerifyIntlResponseBodyResult) *Mobile2MetaVerifyIntlResponseBody
	GetResult() *Mobile2MetaVerifyIntlResponseBodyResult
}

type Mobile2MetaVerifyIntlResponseBody struct {
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
	// 4EB35****87EBA1
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *Mobile2MetaVerifyIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s Mobile2MetaVerifyIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Mobile2MetaVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile2MetaVerifyIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *Mobile2MetaVerifyIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Mobile2MetaVerifyIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Mobile2MetaVerifyIntlResponseBody) GetResult() *Mobile2MetaVerifyIntlResponseBodyResult {
	return s.Result
}

func (s *Mobile2MetaVerifyIntlResponseBody) SetCode(v string) *Mobile2MetaVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile2MetaVerifyIntlResponseBody) SetMessage(v string) *Mobile2MetaVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile2MetaVerifyIntlResponseBody) SetRequestId(v string) *Mobile2MetaVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile2MetaVerifyIntlResponseBody) SetResult(v *Mobile2MetaVerifyIntlResponseBodyResult) *Mobile2MetaVerifyIntlResponseBody {
	s.Result = v
	return s
}

func (s *Mobile2MetaVerifyIntlResponseBody) Validate() error {
	return dara.Validate(s)
}

type Mobile2MetaVerifyIntlResponseBodyResult struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s Mobile2MetaVerifyIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s Mobile2MetaVerifyIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *Mobile2MetaVerifyIntlResponseBodyResult) GetBizCode() *string {
	return s.BizCode
}

func (s *Mobile2MetaVerifyIntlResponseBodyResult) GetIspName() *string {
	return s.IspName
}

func (s *Mobile2MetaVerifyIntlResponseBodyResult) SetBizCode(v string) *Mobile2MetaVerifyIntlResponseBodyResult {
	s.BizCode = &v
	return s
}

func (s *Mobile2MetaVerifyIntlResponseBodyResult) SetIspName(v string) *Mobile2MetaVerifyIntlResponseBodyResult {
	s.IspName = &v
	return s
}

func (s *Mobile2MetaVerifyIntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
