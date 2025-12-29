// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPhoneNumberOnlineTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryPhoneNumberOnlineTimeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryPhoneNumberOnlineTimeResponseBody
	GetCode() *string
	SetData(v *QueryPhoneNumberOnlineTimeResponseBodyData) *QueryPhoneNumberOnlineTimeResponseBody
	GetData() *QueryPhoneNumberOnlineTimeResponseBodyData
	SetMessage(v string) *QueryPhoneNumberOnlineTimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryPhoneNumberOnlineTimeResponseBody
	GetRequestId() *string
}

type QueryPhoneNumberOnlineTimeResponseBody struct {
	AccessDeniedDetail *string                                     `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *QueryPhoneNumberOnlineTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPhoneNumberOnlineTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneNumberOnlineTimeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) GetData() *QueryPhoneNumberOnlineTimeResponseBodyData {
	return s.Data
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) SetAccessDeniedDetail(v string) *QueryPhoneNumberOnlineTimeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) SetCode(v string) *QueryPhoneNumberOnlineTimeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) SetData(v *QueryPhoneNumberOnlineTimeResponseBodyData) *QueryPhoneNumberOnlineTimeResponseBody {
	s.Data = v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) SetMessage(v string) *QueryPhoneNumberOnlineTimeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) SetRequestId(v string) *QueryPhoneNumberOnlineTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPhoneNumberOnlineTimeResponseBodyData struct {
	// example:
	//
	// 示例值示例值示例值
	CarrierCode *string `json:"CarrierCode,omitempty" xml:"CarrierCode,omitempty"`
	// example:
	//
	// 29
	VerifyResult *int64 `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s QueryPhoneNumberOnlineTimeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneNumberOnlineTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPhoneNumberOnlineTimeResponseBodyData) GetCarrierCode() *string {
	return s.CarrierCode
}

func (s *QueryPhoneNumberOnlineTimeResponseBodyData) GetVerifyResult() *int64 {
	return s.VerifyResult
}

func (s *QueryPhoneNumberOnlineTimeResponseBodyData) SetCarrierCode(v string) *QueryPhoneNumberOnlineTimeResponseBodyData {
	s.CarrierCode = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponseBodyData) SetVerifyResult(v int64) *QueryPhoneNumberOnlineTimeResponseBodyData {
	s.VerifyResult = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
