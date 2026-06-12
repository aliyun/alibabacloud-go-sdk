// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceTestCasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListServiceTestCasesResponseBodyData) *ListServiceTestCasesResponseBody
	GetData() []*ListServiceTestCasesResponseBodyData
	SetMaxResults(v int32) *ListServiceTestCasesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceTestCasesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceTestCasesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListServiceTestCasesResponseBody
	GetTotalCount() *int32
}

type ListServiceTestCasesResponseBody struct {
	// The data returned.
	Data []*ListServiceTestCasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The number of entries returned on each page. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA3AE512-6D30-549A-B52D-B9042CA8D515
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that meet the filter criteria.
	//
	// example:
	//
	// 18
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceTestCasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestCasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceTestCasesResponseBody) GetData() []*ListServiceTestCasesResponseBodyData {
	return s.Data
}

func (s *ListServiceTestCasesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceTestCasesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceTestCasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceTestCasesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServiceTestCasesResponseBody) SetData(v []*ListServiceTestCasesResponseBodyData) *ListServiceTestCasesResponseBody {
	s.Data = v
	return s
}

func (s *ListServiceTestCasesResponseBody) SetMaxResults(v int32) *ListServiceTestCasesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTestCasesResponseBody) SetNextToken(v string) *ListServiceTestCasesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceTestCasesResponseBody) SetRequestId(v string) *ListServiceTestCasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceTestCasesResponseBody) SetTotalCount(v int32) *ListServiceTestCasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceTestCasesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceTestCasesResponseBodyData struct {
	// The template name.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The ID of the test case.
	//
	// example:
	//
	// stc-83fcee1383354e35b151
	TestCaseId *string `json:"TestCaseId,omitempty" xml:"TestCaseId,omitempty"`
	// The name of the test case.
	//
	// example:
	//
	// case1
	TestCaseName *string `json:"TestCaseName,omitempty" xml:"TestCaseName,omitempty"`
	// The user-configured YAML file.
	//
	// example:
	//
	// ---
	//
	// parameters:
	//
	//   PayType: "PostPaid"
	//
	//   EcsInstanceType: "$[iact3-auto]"
	//
	//   InstancePassword: "$[iact3-auto]"
	TestConfig *string `json:"TestConfig,omitempty" xml:"TestConfig,omitempty"`
}

func (s ListServiceTestCasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestCasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServiceTestCasesResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListServiceTestCasesResponseBodyData) GetTestCaseId() *string {
	return s.TestCaseId
}

func (s *ListServiceTestCasesResponseBodyData) GetTestCaseName() *string {
	return s.TestCaseName
}

func (s *ListServiceTestCasesResponseBodyData) GetTestConfig() *string {
	return s.TestConfig
}

func (s *ListServiceTestCasesResponseBodyData) SetTemplateName(v string) *ListServiceTestCasesResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *ListServiceTestCasesResponseBodyData) SetTestCaseId(v string) *ListServiceTestCasesResponseBodyData {
	s.TestCaseId = &v
	return s
}

func (s *ListServiceTestCasesResponseBodyData) SetTestCaseName(v string) *ListServiceTestCasesResponseBodyData {
	s.TestCaseName = &v
	return s
}

func (s *ListServiceTestCasesResponseBodyData) SetTestConfig(v string) *ListServiceTestCasesResponseBodyData {
	s.TestConfig = &v
	return s
}

func (s *ListServiceTestCasesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
