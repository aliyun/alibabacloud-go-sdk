// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAnalysisTransparentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribePhoneNumberAnalysisTransparentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribePhoneNumberAnalysisTransparentResponseBody
	GetCode() *string
	SetData(v *DescribePhoneNumberAnalysisTransparentResponseBodyData) *DescribePhoneNumberAnalysisTransparentResponseBody
	GetData() *DescribePhoneNumberAnalysisTransparentResponseBodyData
	SetMessage(v string) *DescribePhoneNumberAnalysisTransparentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePhoneNumberAnalysisTransparentResponseBody
	GetRequestId() *string
}

type DescribePhoneNumberAnalysisTransparentResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribePhoneNumberAnalysisTransparentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAnalysisTransparentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisTransparentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) GetData() *DescribePhoneNumberAnalysisTransparentResponseBodyData {
	return s.Data
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberAnalysisTransparentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) SetCode(v string) *DescribePhoneNumberAnalysisTransparentResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) SetData(v *DescribePhoneNumberAnalysisTransparentResponseBodyData) *DescribePhoneNumberAnalysisTransparentResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) SetMessage(v string) *DescribePhoneNumberAnalysisTransparentResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) SetRequestId(v string) *DescribePhoneNumberAnalysisTransparentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhoneNumberAnalysisTransparentResponseBodyData struct {
	// example:
	//
	// 示例值示例值示例值
	DeviceRisk *string `json:"Device_risk,omitempty" xml:"Device_risk,omitempty"`
	// example:
	//
	// 示例值示例值
	IpRisk *string `json:"Ip_risk,omitempty" xml:"Ip_risk,omitempty"`
	// example:
	//
	// 0.6
	Score1 *string `json:"Score1,omitempty" xml:"Score1,omitempty"`
	// example:
	//
	// 0.2
	Score2 *string `json:"Score2,omitempty" xml:"Score2,omitempty"`
	// example:
	//
	// 0.8
	Score3 *string `json:"Score3,omitempty" xml:"Score3,omitempty"`
}

func (s DescribePhoneNumberAnalysisTransparentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisTransparentResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) GetDeviceRisk() *string {
	return s.DeviceRisk
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) GetIpRisk() *string {
	return s.IpRisk
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) GetScore1() *string {
	return s.Score1
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) GetScore2() *string {
	return s.Score2
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) GetScore3() *string {
	return s.Score3
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) SetDeviceRisk(v string) *DescribePhoneNumberAnalysisTransparentResponseBodyData {
	s.DeviceRisk = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) SetIpRisk(v string) *DescribePhoneNumberAnalysisTransparentResponseBodyData {
	s.IpRisk = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) SetScore1(v string) *DescribePhoneNumberAnalysisTransparentResponseBodyData {
	s.Score1 = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) SetScore2(v string) *DescribePhoneNumberAnalysisTransparentResponseBodyData {
	s.Score2 = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) SetScore3(v string) *DescribePhoneNumberAnalysisTransparentResponseBodyData {
	s.Score3 = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
