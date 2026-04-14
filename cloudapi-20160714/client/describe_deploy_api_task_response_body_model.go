// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeployApiTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeployedResults(v *DescribeDeployApiTaskResponseBodyDeployedResults) *DescribeDeployApiTaskResponseBody
	GetDeployedResults() *DescribeDeployApiTaskResponseBodyDeployedResults
	SetRequestId(v string) *DescribeDeployApiTaskResponseBody
	GetRequestId() *string
}

type DescribeDeployApiTaskResponseBody struct {
	DeployedResults *DescribeDeployApiTaskResponseBodyDeployedResults `json:"DeployedResults,omitempty" xml:"DeployedResults,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CA4B3261-F14A-5E33-8608-F75A1DF27AD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeployApiTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployApiTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeployApiTaskResponseBody) GetDeployedResults() *DescribeDeployApiTaskResponseBodyDeployedResults {
	return s.DeployedResults
}

func (s *DescribeDeployApiTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeployApiTaskResponseBody) SetDeployedResults(v *DescribeDeployApiTaskResponseBodyDeployedResults) *DescribeDeployApiTaskResponseBody {
	s.DeployedResults = v
	return s
}

func (s *DescribeDeployApiTaskResponseBody) SetRequestId(v string) *DescribeDeployApiTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeployApiTaskResponseBody) Validate() error {
	if s.DeployedResults != nil {
		if err := s.DeployedResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDeployApiTaskResponseBodyDeployedResults struct {
	DeployedResult []*DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult `json:"DeployedResult,omitempty" xml:"DeployedResult,omitempty" type:"Repeated"`
}

func (s DescribeDeployApiTaskResponseBodyDeployedResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployApiTaskResponseBodyDeployedResults) GoString() string {
	return s.String()
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResults) GetDeployedResult() []*DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult {
	return s.DeployedResult
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResults) SetDeployedResult(v []*DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) *DescribeDeployApiTaskResponseBodyDeployedResults {
	s.DeployedResult = v
	return s
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResults) Validate() error {
	if s.DeployedResult != nil {
		for _, item := range s.DeployedResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult struct {
	ApiUid         *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	DeployedStatus *string `json:"DeployedStatus,omitempty" xml:"DeployedStatus,omitempty"`
	ErrorMsg       *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	StageName      *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) GoString() string {
	return s.String()
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) GetApiUid() *string {
	return s.ApiUid
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) GetDeployedStatus() *string {
	return s.DeployedStatus
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) GetStageName() *string {
	return s.StageName
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) SetApiUid(v string) *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult {
	s.ApiUid = &v
	return s
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) SetDeployedStatus(v string) *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult {
	s.DeployedStatus = &v
	return s
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) SetErrorMsg(v string) *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) SetGroupId(v string) *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult {
	s.GroupId = &v
	return s
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) SetStageName(v string) *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult {
	s.StageName = &v
	return s
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) Validate() error {
	return dara.Validate(s)
}
