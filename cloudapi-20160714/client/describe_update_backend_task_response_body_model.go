// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpdateBackendTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiUpdateBackendResults(v *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResults) *DescribeUpdateBackendTaskResponseBody
	GetApiUpdateBackendResults() *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResults
	SetRequestId(v string) *DescribeUpdateBackendTaskResponseBody
	GetRequestId() *string
}

type DescribeUpdateBackendTaskResponseBody struct {
	ApiUpdateBackendResults *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResults `json:"ApiUpdateBackendResults,omitempty" xml:"ApiUpdateBackendResults,omitempty" type:"Struct"`
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ016
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUpdateBackendTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpdateBackendTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpdateBackendTaskResponseBody) GetApiUpdateBackendResults() *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResults {
	return s.ApiUpdateBackendResults
}

func (s *DescribeUpdateBackendTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUpdateBackendTaskResponseBody) SetApiUpdateBackendResults(v *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResults) *DescribeUpdateBackendTaskResponseBody {
	s.ApiUpdateBackendResults = v
	return s
}

func (s *DescribeUpdateBackendTaskResponseBody) SetRequestId(v string) *DescribeUpdateBackendTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUpdateBackendTaskResponseBody) Validate() error {
	if s.ApiUpdateBackendResults != nil {
		if err := s.ApiUpdateBackendResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResults struct {
	ApiUpdateBackendResult []*DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult `json:"ApiUpdateBackendResult,omitempty" xml:"ApiUpdateBackendResult,omitempty" type:"Repeated"`
}

func (s DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResults) GoString() string {
	return s.String()
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResults) GetApiUpdateBackendResult() []*DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult {
	return s.ApiUpdateBackendResult
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResults) SetApiUpdateBackendResult(v []*DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResults {
	s.ApiUpdateBackendResult = v
	return s
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResults) Validate() error {
	if s.ApiUpdateBackendResult != nil {
		for _, item := range s.ApiUpdateBackendResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult struct {
	// example:
	//
	// checkin_linechart_today
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// example:
	//
	// 14faa7ba0572445685866ddb6a6f19da
	ApiUid *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	// example:
	//
	// c09b078bcb8f4ade9677bd8b18cdf43f
	BackendId *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	// example:
	//
	// Failed
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 3013a55c0c44483f984d26df27120513
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// imotob1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// d8f2f54f3309458b8aaceb36c01c2dd9
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// example:
	//
	// OVER
	UpdateStatus *string `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
}

func (s DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) GoString() string {
	return s.String()
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) GetApiUid() *string {
	return s.ApiUid
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) GetBackendId() *string {
	return s.BackendId
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) GetStageId() *string {
	return s.StageId
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) GetStageName() *string {
	return s.StageName
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) GetUpdateStatus() *string {
	return s.UpdateStatus
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) SetApiName(v string) *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult {
	s.ApiName = &v
	return s
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) SetApiUid(v string) *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult {
	s.ApiUid = &v
	return s
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) SetBackendId(v string) *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult {
	s.BackendId = &v
	return s
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) SetErrorMsg(v string) *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) SetGroupId(v string) *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult {
	s.GroupId = &v
	return s
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) SetGroupName(v string) *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult {
	s.GroupName = &v
	return s
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) SetStageId(v string) *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult {
	s.StageId = &v
	return s
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) SetStageName(v string) *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult {
	s.StageName = &v
	return s
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) SetUpdateStatus(v string) *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult {
	s.UpdateStatus = &v
	return s
}

func (s *DescribeUpdateBackendTaskResponseBodyApiUpdateBackendResultsApiUpdateBackendResult) Validate() error {
	return dara.Validate(s)
}
