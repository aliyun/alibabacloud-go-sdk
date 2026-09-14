// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossProjectDeploymentCandidatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListCrossProjectDeploymentCandidatesResponseBodyData) *ListCrossProjectDeploymentCandidatesResponseBody
	GetData() *ListCrossProjectDeploymentCandidatesResponseBodyData
	SetRequestId(v string) *ListCrossProjectDeploymentCandidatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCrossProjectDeploymentCandidatesResponseBody
	GetSuccess() *bool
}

type ListCrossProjectDeploymentCandidatesResponseBody struct {
	// The business response.
	//
	// example:
	//
	// {"RequestId":"735894D1-D5E5-50B8-8A6D-041C90A98B23","PageNumber":1,"PageSize":10,"TotalCount":1,"DeploymentCandidates":[{"ObjectId":"1","ObjectType":"ODPS_SQL","ObjectName":"object-1","ObjectVersion":"7","ChangeType":"ADD","CommitUser":"operator","CommitTime":1788739200000}]}
	Data *ListCrossProjectDeploymentCandidatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID, which is used to locate and troubleshoot this API call.
	//
	// example:
	//
	// 735894D1-D5E5-50B8-8A6D-041C90A98B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCrossProjectDeploymentCandidatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectDeploymentCandidatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrossProjectDeploymentCandidatesResponseBody) GetData() *ListCrossProjectDeploymentCandidatesResponseBodyData {
	return s.Data
}

func (s *ListCrossProjectDeploymentCandidatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrossProjectDeploymentCandidatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCrossProjectDeploymentCandidatesResponseBody) SetData(v *ListCrossProjectDeploymentCandidatesResponseBodyData) *ListCrossProjectDeploymentCandidatesResponseBody {
	s.Data = v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBody) SetRequestId(v string) *ListCrossProjectDeploymentCandidatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBody) SetSuccess(v bool) *ListCrossProjectDeploymentCandidatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCrossProjectDeploymentCandidatesResponseBodyData struct {
	// The list of candidate objects from the source workspace that are available for cross-workspace deployment.
	//
	// example:
	//
	// [{"ObjectId":"1","ObjectType":"ODPS_SQL","ObjectName":"object-1","ObjectVersion":"7","ChangeType":"ADD","CommitUser":"operator","CommitTime":1788739200000}]
	DeploymentCandidates []*ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates `json:"DeploymentCandidates,omitempty" xml:"DeploymentCandidates,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 735894D1-D5E5-50B8-8A6D-041C90A98B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCrossProjectDeploymentCandidatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectDeploymentCandidatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyData) GetDeploymentCandidates() []*ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates {
	return s.DeploymentCandidates
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyData) SetDeploymentCandidates(v []*ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) *ListCrossProjectDeploymentCandidatesResponseBodyData {
	s.DeploymentCandidates = v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyData) SetPageNumber(v int32) *ListCrossProjectDeploymentCandidatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyData) SetPageSize(v int32) *ListCrossProjectDeploymentCandidatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyData) SetRequestId(v string) *ListCrossProjectDeploymentCandidatesResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyData) SetTotalCount(v int32) *ListCrossProjectDeploymentCandidatesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyData) Validate() error {
	if s.DeploymentCandidates != nil {
		for _, item := range s.DeploymentCandidates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates struct {
	// The change type.
	//
	// example:
	//
	// ADD
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The commit time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1788739200000
	CommitTime *int64 `json:"CommitTime,omitempty" xml:"CommitTime,omitempty"`
	// The committer.
	//
	// example:
	//
	// operator
	CommitUser *string `json:"CommitUser,omitempty" xml:"CommitUser,omitempty"`
	// The candidate object ID.
	//
	// example:
	//
	// 1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The candidate object name.
	//
	// example:
	//
	// object-1
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The candidate object type.
	//
	// example:
	//
	// ODPS_SQL
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The candidate object version.
	//
	// example:
	//
	// 7
	ObjectVersion *string `json:"ObjectVersion,omitempty" xml:"ObjectVersion,omitempty"`
}

func (s ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) GoString() string {
	return s.String()
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) GetChangeType() *string {
	return s.ChangeType
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) GetCommitTime() *int64 {
	return s.CommitTime
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) GetCommitUser() *string {
	return s.CommitUser
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) GetObjectVersion() *string {
	return s.ObjectVersion
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) SetChangeType(v string) *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates {
	s.ChangeType = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) SetCommitTime(v int64) *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates {
	s.CommitTime = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) SetCommitUser(v string) *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates {
	s.CommitUser = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) SetObjectId(v string) *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates {
	s.ObjectId = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) SetObjectName(v string) *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates {
	s.ObjectName = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) SetObjectType(v string) *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates {
	s.ObjectType = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) SetObjectVersion(v string) *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates {
	s.ObjectVersion = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponseBodyDataDeploymentCandidates) Validate() error {
	return dara.Validate(s)
}
