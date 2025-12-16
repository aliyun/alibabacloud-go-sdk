// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitAuthVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitAuthVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *InitAuthVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitAuthVerifyResponseBody
	GetRequestId() *string
	SetResult(v *InitAuthVerifyResponseBodyResult) *InitAuthVerifyResponseBody
	GetResult() *InitAuthVerifyResponseBodyResult
}

type InitAuthVerifyResponseBody struct {
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
	// B506328A-D84B-4750-82C7-6A207C585CF1
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *InitAuthVerifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s InitAuthVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitAuthVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *InitAuthVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitAuthVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitAuthVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitAuthVerifyResponseBody) GetResult() *InitAuthVerifyResponseBodyResult {
	return s.Result
}

func (s *InitAuthVerifyResponseBody) SetCode(v string) *InitAuthVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *InitAuthVerifyResponseBody) SetMessage(v string) *InitAuthVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *InitAuthVerifyResponseBody) SetRequestId(v string) *InitAuthVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitAuthVerifyResponseBody) SetResult(v *InitAuthVerifyResponseBodyResult) *InitAuthVerifyResponseBody {
	s.Result = v
	return s
}

func (s *InitAuthVerifyResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitAuthVerifyResponseBodyResult struct {
	// example:
	//
	// shif9d1a185b8dde7cd07bf0943a448b
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
}

func (s InitAuthVerifyResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s InitAuthVerifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InitAuthVerifyResponseBodyResult) GetCertifyId() *string {
	return s.CertifyId
}

func (s *InitAuthVerifyResponseBodyResult) SetCertifyId(v string) *InitAuthVerifyResponseBodyResult {
	s.CertifyId = &v
	return s
}

func (s *InitAuthVerifyResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
