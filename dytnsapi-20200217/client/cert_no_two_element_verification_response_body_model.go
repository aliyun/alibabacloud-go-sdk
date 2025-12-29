// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertNoTwoElementVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CertNoTwoElementVerificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CertNoTwoElementVerificationResponseBody
	GetCode() *string
	SetData(v *CertNoTwoElementVerificationResponseBodyData) *CertNoTwoElementVerificationResponseBody
	GetData() *CertNoTwoElementVerificationResponseBodyData
	SetMessage(v string) *CertNoTwoElementVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CertNoTwoElementVerificationResponseBody
	GetRequestId() *string
}

type CertNoTwoElementVerificationResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CertNoTwoElementVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CertNoTwoElementVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CertNoTwoElementVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CertNoTwoElementVerificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CertNoTwoElementVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CertNoTwoElementVerificationResponseBody) GetData() *CertNoTwoElementVerificationResponseBodyData {
	return s.Data
}

func (s *CertNoTwoElementVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CertNoTwoElementVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CertNoTwoElementVerificationResponseBody) SetAccessDeniedDetail(v string) *CertNoTwoElementVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CertNoTwoElementVerificationResponseBody) SetCode(v string) *CertNoTwoElementVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *CertNoTwoElementVerificationResponseBody) SetData(v *CertNoTwoElementVerificationResponseBodyData) *CertNoTwoElementVerificationResponseBody {
	s.Data = v
	return s
}

func (s *CertNoTwoElementVerificationResponseBody) SetMessage(v string) *CertNoTwoElementVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *CertNoTwoElementVerificationResponseBody) SetRequestId(v string) *CertNoTwoElementVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertNoTwoElementVerificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CertNoTwoElementVerificationResponseBodyData struct {
	// example:
	//
	// 1
	IsConsistent *string `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
}

func (s CertNoTwoElementVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CertNoTwoElementVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CertNoTwoElementVerificationResponseBodyData) GetIsConsistent() *string {
	return s.IsConsistent
}

func (s *CertNoTwoElementVerificationResponseBodyData) SetIsConsistent(v string) *CertNoTwoElementVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

func (s *CertNoTwoElementVerificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
