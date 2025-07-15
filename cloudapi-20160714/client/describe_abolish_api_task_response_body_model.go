// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAbolishApiTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiAbolishResults(v *DescribeAbolishApiTaskResponseBodyApiAbolishResults) *DescribeAbolishApiTaskResponseBody
	GetApiAbolishResults() *DescribeAbolishApiTaskResponseBodyApiAbolishResults
	SetRequestId(v string) *DescribeAbolishApiTaskResponseBody
	GetRequestId() *string
}

type DescribeAbolishApiTaskResponseBody struct {
	// The result returned.
	ApiAbolishResults *DescribeAbolishApiTaskResponseBodyApiAbolishResults `json:"ApiAbolishResults,omitempty" xml:"ApiAbolishResults,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E8515BA6-81CD-4191-A7CF-C4FCDD3C0D99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAbolishApiTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbolishApiTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAbolishApiTaskResponseBody) GetApiAbolishResults() *DescribeAbolishApiTaskResponseBodyApiAbolishResults {
	return s.ApiAbolishResults
}

func (s *DescribeAbolishApiTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAbolishApiTaskResponseBody) SetApiAbolishResults(v *DescribeAbolishApiTaskResponseBodyApiAbolishResults) *DescribeAbolishApiTaskResponseBody {
	s.ApiAbolishResults = v
	return s
}

func (s *DescribeAbolishApiTaskResponseBody) SetRequestId(v string) *DescribeAbolishApiTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAbolishApiTaskResponseBodyApiAbolishResults struct {
	ApiAbolishResult []*DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult `json:"ApiAbolishResult,omitempty" xml:"ApiAbolishResult,omitempty" type:"Repeated"`
}

func (s DescribeAbolishApiTaskResponseBodyApiAbolishResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbolishApiTaskResponseBodyApiAbolishResults) GoString() string {
	return s.String()
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResults) GetApiAbolishResult() []*DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	return s.ApiAbolishResult
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResults) SetApiAbolishResult(v []*DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) *DescribeAbolishApiTaskResponseBodyApiAbolishResults {
	s.ApiAbolishResult = v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResults) Validate() error {
	return dara.Validate(s)
}

type DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult struct {
	// The unpublishing status.
	//
	// example:
	//
	// OVER
	AbolishStatus *string `json:"AbolishStatus,omitempty" xml:"AbolishStatus,omitempty"`
	// The name of the API.
	//
	// example:
	//
	// v2_page_consent
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The ID of the API.
	//
	// example:
	//
	// 4e26cdbbb113416dba1f0285bed29979
	ApiUid *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success. Request Success.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 160cb6505e1c43a6b84346856d74eb47
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group.
	//
	// example:
	//
	// wb2022021401619286
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the runtime environment.
	//
	// example:
	//
	// 0919f2854a88484c91dc9253347c78f9
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) GoString() string {
	return s.String()
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) GetAbolishStatus() *string {
	return s.AbolishStatus
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) GetApiUid() *string {
	return s.ApiUid
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) GetStageId() *string {
	return s.StageId
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) GetStageName() *string {
	return s.StageName
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetAbolishStatus(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.AbolishStatus = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetApiName(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.ApiName = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetApiUid(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.ApiUid = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetErrorMsg(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetGroupId(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.GroupId = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetGroupName(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.GroupName = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetStageId(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.StageId = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetStageName(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.StageName = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) Validate() error {
	return dara.Validate(s)
}
