// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaVerifyIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Mobile3MetaVerifyIntlResponseBody
	GetCode() *string
	SetMessage(v string) *Mobile3MetaVerifyIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *Mobile3MetaVerifyIntlResponseBody
	GetRequestId() *string
	SetResult(v *Mobile3MetaVerifyIntlResponseBodyResult) *Mobile3MetaVerifyIntlResponseBody
	GetResult() *Mobile3MetaVerifyIntlResponseBodyResult
}

type Mobile3MetaVerifyIntlResponseBody struct {
	// Return code
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// D241532C-4EE9-5A2A-A5A5-C1FD98CE2EDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result
	Result *Mobile3MetaVerifyIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s Mobile3MetaVerifyIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile3MetaVerifyIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *Mobile3MetaVerifyIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Mobile3MetaVerifyIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Mobile3MetaVerifyIntlResponseBody) GetResult() *Mobile3MetaVerifyIntlResponseBodyResult {
	return s.Result
}

func (s *Mobile3MetaVerifyIntlResponseBody) SetCode(v string) *Mobile3MetaVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBody) SetMessage(v string) *Mobile3MetaVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBody) SetRequestId(v string) *Mobile3MetaVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBody) SetResult(v *Mobile3MetaVerifyIntlResponseBodyResult) *Mobile3MetaVerifyIntlResponseBody {
	s.Result = v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Mobile3MetaVerifyIntlResponseBodyResult struct {
	// Verification result code.
	//
	// - 1: Verification consistent
	//
	// - 2: Verification inconsistent
	//
	// - 3: No record found
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// ISP name
	//
	// - CMCC: China Mobile
	//
	// - CUCC: China Unicom
	//
	// - CTCC: China Telecom
	//
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// Detailed verification results
	//
	// - 101: Verification passed
	//
	// - 201: Mobile number and name do not match, mobile number and ID number do not match
	//
	// - 202: Mobile number and name match, but mobile number and ID number do not match
	//
	// - 203: Mobile number and ID number match, but mobile number and name do not match
	//
	// - 204: Other inconsistencies
	//
	// - 301: No record found
	//
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s Mobile3MetaVerifyIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaVerifyIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) GetBizCode() *string {
	return s.BizCode
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) GetIspName() *string {
	return s.IspName
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) SetBizCode(v string) *Mobile3MetaVerifyIntlResponseBodyResult {
	s.BizCode = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) SetIspName(v string) *Mobile3MetaVerifyIntlResponseBodyResult {
	s.IspName = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) SetSubCode(v string) *Mobile3MetaVerifyIntlResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
