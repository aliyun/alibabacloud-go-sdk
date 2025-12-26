// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecExamplesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExamples(v []*DescribeApisecExamplesResponseBodyExamples) *DescribeApisecExamplesResponseBody
	GetExamples() []*DescribeApisecExamplesResponseBodyExamples
	SetMaxResults(v int32) *DescribeApisecExamplesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeApisecExamplesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeApisecExamplesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApisecExamplesResponseBody
	GetTotalCount() *int32
}

type DescribeApisecExamplesResponseBody struct {
	Examples []*DescribeApisecExamplesResponseBodyExamples `json:"Examples,omitempty" xml:"Examples,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisecExamplesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecExamplesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecExamplesResponseBody) GetExamples() []*DescribeApisecExamplesResponseBodyExamples {
	return s.Examples
}

func (s *DescribeApisecExamplesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeApisecExamplesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeApisecExamplesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecExamplesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApisecExamplesResponseBody) SetExamples(v []*DescribeApisecExamplesResponseBodyExamples) *DescribeApisecExamplesResponseBody {
	s.Examples = v
	return s
}

func (s *DescribeApisecExamplesResponseBody) SetMaxResults(v int32) *DescribeApisecExamplesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeApisecExamplesResponseBody) SetNextToken(v string) *DescribeApisecExamplesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeApisecExamplesResponseBody) SetRequestId(v string) *DescribeApisecExamplesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecExamplesResponseBody) SetTotalCount(v int32) *DescribeApisecExamplesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisecExamplesResponseBody) Validate() error {
	if s.Examples != nil {
		for _, item := range s.Examples {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisecExamplesResponseBodyExamples struct {
	// example:
	//
	// http://www.test.com/api/v1/hello.php?token=TkJGQw
	ApiUrl     *string `json:"ApiUrl,omitempty" xml:"ApiUrl,omitempty"`
	PocPayload *string `json:"PocPayload,omitempty" xml:"PocPayload,omitempty"`
	// example:
	//
	// http
	Protocol             *string                                                           `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Request              *string                                                           `json:"Request,omitempty" xml:"Request,omitempty"`
	RequestSensitiveData []*DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData `json:"RequestSensitiveData,omitempty" xml:"RequestSensitiveData,omitempty" type:"Repeated"`
	// example:
	//
	// {
	//
	//   "header": {
	//
	//     "Connection": "keep-alive",
	//
	//     "Content-Encoding": "gzip",
	//
	//     "Content-Type": "text/html; charset=UTF-8"
	//
	//   },
	//
	//   "body": "xxxx",
	//
	//   "status": 200
	//
	// }
	Response              *string                                                            `json:"Response,omitempty" xml:"Response,omitempty"`
	ResponseSensitiveData []*DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData `json:"ResponseSensitiveData,omitempty" xml:"ResponseSensitiveData,omitempty" type:"Repeated"`
}

func (s DescribeApisecExamplesResponseBodyExamples) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecExamplesResponseBodyExamples) GoString() string {
	return s.String()
}

func (s *DescribeApisecExamplesResponseBodyExamples) GetApiUrl() *string {
	return s.ApiUrl
}

func (s *DescribeApisecExamplesResponseBodyExamples) GetPocPayload() *string {
	return s.PocPayload
}

func (s *DescribeApisecExamplesResponseBodyExamples) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeApisecExamplesResponseBodyExamples) GetRequest() *string {
	return s.Request
}

func (s *DescribeApisecExamplesResponseBodyExamples) GetRequestSensitiveData() []*DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData {
	return s.RequestSensitiveData
}

func (s *DescribeApisecExamplesResponseBodyExamples) GetResponse() *string {
	return s.Response
}

func (s *DescribeApisecExamplesResponseBodyExamples) GetResponseSensitiveData() []*DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData {
	return s.ResponseSensitiveData
}

func (s *DescribeApisecExamplesResponseBodyExamples) SetApiUrl(v string) *DescribeApisecExamplesResponseBodyExamples {
	s.ApiUrl = &v
	return s
}

func (s *DescribeApisecExamplesResponseBodyExamples) SetPocPayload(v string) *DescribeApisecExamplesResponseBodyExamples {
	s.PocPayload = &v
	return s
}

func (s *DescribeApisecExamplesResponseBodyExamples) SetProtocol(v string) *DescribeApisecExamplesResponseBodyExamples {
	s.Protocol = &v
	return s
}

func (s *DescribeApisecExamplesResponseBodyExamples) SetRequest(v string) *DescribeApisecExamplesResponseBodyExamples {
	s.Request = &v
	return s
}

func (s *DescribeApisecExamplesResponseBodyExamples) SetRequestSensitiveData(v []*DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData) *DescribeApisecExamplesResponseBodyExamples {
	s.RequestSensitiveData = v
	return s
}

func (s *DescribeApisecExamplesResponseBodyExamples) SetResponse(v string) *DescribeApisecExamplesResponseBodyExamples {
	s.Response = &v
	return s
}

func (s *DescribeApisecExamplesResponseBodyExamples) SetResponseSensitiveData(v []*DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData) *DescribeApisecExamplesResponseBodyExamples {
	s.ResponseSensitiveData = v
	return s
}

func (s *DescribeApisecExamplesResponseBodyExamples) Validate() error {
	if s.RequestSensitiveData != nil {
		for _, item := range s.RequestSensitiveData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResponseSensitiveData != nil {
		for _, item := range s.ResponseSensitiveData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData struct {
	// example:
	//
	// 1000
	SensitiveCode     *string   `json:"SensitiveCode,omitempty" xml:"SensitiveCode,omitempty"`
	SensitiveDataList []*string `json:"SensitiveDataList,omitempty" xml:"SensitiveDataList,omitempty" type:"Repeated"`
}

func (s DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData) GoString() string {
	return s.String()
}

func (s *DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData) GetSensitiveCode() *string {
	return s.SensitiveCode
}

func (s *DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData) GetSensitiveDataList() []*string {
	return s.SensitiveDataList
}

func (s *DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData) SetSensitiveCode(v string) *DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData {
	s.SensitiveCode = &v
	return s
}

func (s *DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData) SetSensitiveDataList(v []*string) *DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData {
	s.SensitiveDataList = v
	return s
}

func (s *DescribeApisecExamplesResponseBodyExamplesRequestSensitiveData) Validate() error {
	return dara.Validate(s)
}

type DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData struct {
	// example:
	//
	// 1000
	SensitiveCode     *string   `json:"SensitiveCode,omitempty" xml:"SensitiveCode,omitempty"`
	SensitiveDataList []*string `json:"SensitiveDataList,omitempty" xml:"SensitiveDataList,omitempty" type:"Repeated"`
}

func (s DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData) GoString() string {
	return s.String()
}

func (s *DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData) GetSensitiveCode() *string {
	return s.SensitiveCode
}

func (s *DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData) GetSensitiveDataList() []*string {
	return s.SensitiveDataList
}

func (s *DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData) SetSensitiveCode(v string) *DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData {
	s.SensitiveCode = &v
	return s
}

func (s *DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData) SetSensitiveDataList(v []*string) *DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData {
	s.SensitiveDataList = v
	return s
}

func (s *DescribeApisecExamplesResponseBodyExamplesResponseSensitiveData) Validate() error {
	return dara.Validate(s)
}
