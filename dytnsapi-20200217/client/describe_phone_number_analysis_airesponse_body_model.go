// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAnalysisAIResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribePhoneNumberAnalysisAIResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribePhoneNumberAnalysisAIResponseBody
	GetCode() *string
	SetData(v *DescribePhoneNumberAnalysisAIResponseBodyData) *DescribePhoneNumberAnalysisAIResponseBody
	GetData() *DescribePhoneNumberAnalysisAIResponseBodyData
	SetMessage(v string) *DescribePhoneNumberAnalysisAIResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePhoneNumberAnalysisAIResponseBody
	GetRequestId() *string
}

type DescribePhoneNumberAnalysisAIResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code. Valid values:
	//
	// - OK: success
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *DescribePhoneNumberAnalysisAIResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the returned status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD******177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAnalysisAIResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisAIResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) GetData() *DescribePhoneNumberAnalysisAIResponseBodyData {
	return s.Data
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberAnalysisAIResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) SetCode(v string) *DescribePhoneNumberAnalysisAIResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) SetData(v *DescribePhoneNumberAnalysisAIResponseBodyData) *DescribePhoneNumberAnalysisAIResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) SetMessage(v string) *DescribePhoneNumberAnalysisAIResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) SetRequestId(v string) *DescribePhoneNumberAnalysisAIResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhoneNumberAnalysisAIResponseBodyData struct {
	// The returned result code.
	//
	// - YES: valid
	//
	// - NO: invalid
	//
	// - UNKNOWN: unknown
	//
	// example:
	//
	// YES
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The phone number that was passed in.
	//
	// example:
	//
	// 187****5620
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DescribePhoneNumberAnalysisAIResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisAIResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisAIResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneNumberAnalysisAIResponseBodyData) GetNumber() *string {
	return s.Number
}

func (s *DescribePhoneNumberAnalysisAIResponseBodyData) SetCode(v string) *DescribePhoneNumberAnalysisAIResponseBodyData {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponseBodyData) SetNumber(v string) *DescribePhoneNumberAnalysisAIResponseBodyData {
	s.Number = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponseBodyData) Validate() error {
	return dara.Validate(s)
}
