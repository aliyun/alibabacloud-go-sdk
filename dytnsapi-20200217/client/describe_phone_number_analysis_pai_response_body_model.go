// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAnalysisPaiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribePhoneNumberAnalysisPaiResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribePhoneNumberAnalysisPaiResponseBody
	GetCode() *string
	SetData(v []*string) *DescribePhoneNumberAnalysisPaiResponseBody
	GetData() []*string
	SetMessage(v string) *DescribePhoneNumberAnalysisPaiResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePhoneNumberAnalysisPaiResponseBody
	GetRequestId() *string
}

type DescribePhoneNumberAnalysisPaiResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAnalysisPaiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisPaiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) GetData() []*string {
	return s.Data
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberAnalysisPaiResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) SetCode(v string) *DescribePhoneNumberAnalysisPaiResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) SetData(v []*string) *DescribePhoneNumberAnalysisPaiResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) SetMessage(v string) *DescribePhoneNumberAnalysisPaiResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) SetRequestId(v string) *DescribePhoneNumberAnalysisPaiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) Validate() error {
	return dara.Validate(s)
}
