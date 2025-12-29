// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPhoneTwiceTelVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryPhoneTwiceTelVerifyResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryPhoneTwiceTelVerifyResponseBody
	GetCode() *string
	SetData(v *QueryPhoneTwiceTelVerifyResponseBodyData) *QueryPhoneTwiceTelVerifyResponseBody
	GetData() *QueryPhoneTwiceTelVerifyResponseBodyData
	SetMessage(v string) *QueryPhoneTwiceTelVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryPhoneTwiceTelVerifyResponseBody
	GetRequestId() *string
}

type QueryPhoneTwiceTelVerifyResponseBody struct {
	AccessDeniedDetail *string                                   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *QueryPhoneTwiceTelVerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPhoneTwiceTelVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneTwiceTelVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) GetData() *QueryPhoneTwiceTelVerifyResponseBodyData {
	return s.Data
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) SetAccessDeniedDetail(v string) *QueryPhoneTwiceTelVerifyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) SetCode(v string) *QueryPhoneTwiceTelVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) SetData(v *QueryPhoneTwiceTelVerifyResponseBodyData) *QueryPhoneTwiceTelVerifyResponseBody {
	s.Data = v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) SetMessage(v string) *QueryPhoneTwiceTelVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) SetRequestId(v string) *QueryPhoneTwiceTelVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPhoneTwiceTelVerifyResponseBodyData struct {
	// example:
	//
	// 示例值示例值
	CarrierCode *string `json:"CarrierCode,omitempty" xml:"CarrierCode,omitempty"`
	// example:
	//
	// 73
	VerifyResult *int64 `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s QueryPhoneTwiceTelVerifyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneTwiceTelVerifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPhoneTwiceTelVerifyResponseBodyData) GetCarrierCode() *string {
	return s.CarrierCode
}

func (s *QueryPhoneTwiceTelVerifyResponseBodyData) GetVerifyResult() *int64 {
	return s.VerifyResult
}

func (s *QueryPhoneTwiceTelVerifyResponseBodyData) SetCarrierCode(v string) *QueryPhoneTwiceTelVerifyResponseBodyData {
	s.CarrierCode = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponseBodyData) SetVerifyResult(v int64) *QueryPhoneTwiceTelVerifyResponseBodyData {
	s.VerifyResult = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
