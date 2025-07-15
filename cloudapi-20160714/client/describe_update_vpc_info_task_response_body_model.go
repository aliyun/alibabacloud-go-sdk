// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpdateVpcInfoTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiUpdateVpcInfoResults(v *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults) *DescribeUpdateVpcInfoTaskResponseBody
	GetApiUpdateVpcInfoResults() *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults
	SetRequestId(v string) *DescribeUpdateVpcInfoTaskResponseBody
	GetRequestId() *string
}

type DescribeUpdateVpcInfoTaskResponseBody struct {
	ApiUpdateVpcInfoResults *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults `json:"ApiUpdateVpcInfoResults,omitempty" xml:"ApiUpdateVpcInfoResults,omitempty" type:"Struct"`
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUpdateVpcInfoTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpdateVpcInfoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpdateVpcInfoTaskResponseBody) GetApiUpdateVpcInfoResults() *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults {
	return s.ApiUpdateVpcInfoResults
}

func (s *DescribeUpdateVpcInfoTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUpdateVpcInfoTaskResponseBody) SetApiUpdateVpcInfoResults(v *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults) *DescribeUpdateVpcInfoTaskResponseBody {
	s.ApiUpdateVpcInfoResults = v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBody) SetRequestId(v string) *DescribeUpdateVpcInfoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults struct {
	ApiUpdateVpcInfoResult []*DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult `json:"ApiUpdateVpcInfoResult,omitempty" xml:"ApiUpdateVpcInfoResult,omitempty" type:"Repeated"`
}

func (s DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults) GoString() string {
	return s.String()
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults) GetApiUpdateVpcInfoResult() []*DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	return s.ApiUpdateVpcInfoResult
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults) SetApiUpdateVpcInfoResult(v []*DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults {
	s.ApiUpdateVpcInfoResult = v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults) Validate() error {
	return dara.Validate(s)
}

type DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult struct {
	// example:
	//
	// api_test2
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// example:
	//
	// 86364e7c166c47ba819b3f8f95ac0913
	ApiUid *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	// example:
	//
	// Success. Request Success.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// e8da6f6346184da9a30d0dc1888b1f3b
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// fe952b95072747e2a8dfd336bcff8d7f
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// example:
	//
	// success
	UpdateStatus *string `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
}

func (s DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) GoString() string {
	return s.String()
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) GetApiUid() *string {
	return s.ApiUid
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) GetStageId() *string {
	return s.StageId
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) GetStageName() *string {
	return s.StageName
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) GetUpdateStatus() *string {
	return s.UpdateStatus
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetApiName(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.ApiName = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetApiUid(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.ApiUid = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetErrorMsg(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetGroupId(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.GroupId = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetGroupName(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.GroupName = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetStageId(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.StageId = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetStageName(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.StageName = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetUpdateStatus(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.UpdateStatus = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) Validate() error {
	return dara.Validate(s)
}
