// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribePhoneNumberAnalysisResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribePhoneNumberAnalysisResponseBody
	GetCode() *string
	SetData(v *DescribePhoneNumberAnalysisResponseBodyData) *DescribePhoneNumberAnalysisResponseBody
	GetData() *DescribePhoneNumberAnalysisResponseBodyData
	SetMessage(v string) *DescribePhoneNumberAnalysisResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePhoneNumberAnalysisResponseBody
	GetRequestId() *string
}

type DescribePhoneNumberAnalysisResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribePhoneNumberAnalysisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribePhoneNumberAnalysisResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneNumberAnalysisResponseBody) GetData() *DescribePhoneNumberAnalysisResponseBodyData {
	return s.Data
}

func (s *DescribePhoneNumberAnalysisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePhoneNumberAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetCode(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetData(v *DescribePhoneNumberAnalysisResponseBodyData) *DescribePhoneNumberAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetMessage(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetRequestId(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhoneNumberAnalysisResponseBodyData struct {
	List []*DescribePhoneNumberAnalysisResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribePhoneNumberAnalysisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponseBodyData) GetList() []*DescribePhoneNumberAnalysisResponseBodyDataList {
	return s.List
}

func (s *DescribePhoneNumberAnalysisResponseBodyData) SetList(v []*DescribePhoneNumberAnalysisResponseBodyDataList) *DescribePhoneNumberAnalysisResponseBodyData {
	s.List = v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePhoneNumberAnalysisResponseBodyDataList struct {
	// example:
	//
	// NO
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 18752785620
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DescribePhoneNumberAnalysisResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponseBodyDataList) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneNumberAnalysisResponseBodyDataList) GetNumber() *string {
	return s.Number
}

func (s *DescribePhoneNumberAnalysisResponseBodyDataList) SetCode(v string) *DescribePhoneNumberAnalysisResponseBodyDataList {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBodyDataList) SetNumber(v string) *DescribePhoneNumberAnalysisResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
