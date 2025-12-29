// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertNoThreeElementVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CertNoThreeElementVerificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CertNoThreeElementVerificationResponseBody
	GetCode() *string
	SetData(v *CertNoThreeElementVerificationResponseBodyData) *CertNoThreeElementVerificationResponseBody
	GetData() *CertNoThreeElementVerificationResponseBodyData
	SetMessage(v string) *CertNoThreeElementVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CertNoThreeElementVerificationResponseBody
	GetRequestId() *string
}

type CertNoThreeElementVerificationResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CertNoThreeElementVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD******177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CertNoThreeElementVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CertNoThreeElementVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CertNoThreeElementVerificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CertNoThreeElementVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CertNoThreeElementVerificationResponseBody) GetData() *CertNoThreeElementVerificationResponseBodyData {
	return s.Data
}

func (s *CertNoThreeElementVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CertNoThreeElementVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CertNoThreeElementVerificationResponseBody) SetAccessDeniedDetail(v string) *CertNoThreeElementVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CertNoThreeElementVerificationResponseBody) SetCode(v string) *CertNoThreeElementVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *CertNoThreeElementVerificationResponseBody) SetData(v *CertNoThreeElementVerificationResponseBodyData) *CertNoThreeElementVerificationResponseBody {
	s.Data = v
	return s
}

func (s *CertNoThreeElementVerificationResponseBody) SetMessage(v string) *CertNoThreeElementVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *CertNoThreeElementVerificationResponseBody) SetRequestId(v string) *CertNoThreeElementVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertNoThreeElementVerificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CertNoThreeElementVerificationResponseBodyData struct {
	// example:
	//
	// 1
	IsConsistent *string `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
}

func (s CertNoThreeElementVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CertNoThreeElementVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CertNoThreeElementVerificationResponseBodyData) GetIsConsistent() *string {
	return s.IsConsistent
}

func (s *CertNoThreeElementVerificationResponseBodyData) SetIsConsistent(v string) *CertNoThreeElementVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

func (s *CertNoThreeElementVerificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
