// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberRiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribePhoneNumberRiskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribePhoneNumberRiskResponseBody
	GetCode() *string
	SetData(v *DescribePhoneNumberRiskResponseBodyData) *DescribePhoneNumberRiskResponseBody
	GetData() *DescribePhoneNumberRiskResponseBodyData
	SetMessage(v string) *DescribePhoneNumberRiskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePhoneNumberRiskResponseBody
	GetRequestId() *string
}

type DescribePhoneNumberRiskResponseBody struct {
	AccessDeniedDetail *string                                  `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *DescribePhoneNumberRiskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberRiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberRiskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberRiskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribePhoneNumberRiskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneNumberRiskResponseBody) GetData() *DescribePhoneNumberRiskResponseBodyData {
	return s.Data
}

func (s *DescribePhoneNumberRiskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePhoneNumberRiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhoneNumberRiskResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberRiskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePhoneNumberRiskResponseBody) SetCode(v string) *DescribePhoneNumberRiskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberRiskResponseBody) SetData(v *DescribePhoneNumberRiskResponseBodyData) *DescribePhoneNumberRiskResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberRiskResponseBody) SetMessage(v string) *DescribePhoneNumberRiskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberRiskResponseBody) SetRequestId(v string) *DescribePhoneNumberRiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberRiskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhoneNumberRiskResponseBodyData struct {
	// example:
	//
	// 示例值
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s DescribePhoneNumberRiskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberRiskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberRiskResponseBodyData) GetVerifyResult() *string {
	return s.VerifyResult
}

func (s *DescribePhoneNumberRiskResponseBodyData) SetVerifyResult(v string) *DescribePhoneNumberRiskResponseBodyData {
	s.VerifyResult = &v
	return s
}

func (s *DescribePhoneNumberRiskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
