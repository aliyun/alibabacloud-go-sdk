// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossProjectDeploymentEnvironmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListCrossProjectDeploymentEnvironmentsResponseBodyData) *ListCrossProjectDeploymentEnvironmentsResponseBody
	GetData() *ListCrossProjectDeploymentEnvironmentsResponseBodyData
	SetRequestId(v string) *ListCrossProjectDeploymentEnvironmentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCrossProjectDeploymentEnvironmentsResponseBody
	GetSuccess() *bool
}

type ListCrossProjectDeploymentEnvironmentsResponseBody struct {
	// The business response.
	//
	// example:
	//
	// {"RequestId":"735894D1-D5E5-50B8-8A6D-041C90A98B23","PageNumber":1,"PageSize":10,"TotalCount":1,"DeploymentEnvironments":[{"DeploymentEnvironmentId":101,"Name":"environment-101","SourceProjectId":10,"TargetProjectId":20,"TargetProjectName":"target","Status":"Enabled"}]}
	Data *ListCrossProjectDeploymentEnvironmentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListCrossProjectDeploymentEnvironmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectDeploymentEnvironmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBody) GetData() *ListCrossProjectDeploymentEnvironmentsResponseBodyData {
	return s.Data
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBody) SetData(v *ListCrossProjectDeploymentEnvironmentsResponseBodyData) *ListCrossProjectDeploymentEnvironmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBody) SetRequestId(v string) *ListCrossProjectDeploymentEnvironmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBody) SetSuccess(v bool) *ListCrossProjectDeploymentEnvironmentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCrossProjectDeploymentEnvironmentsResponseBodyData struct {
	// The list of enabled cross-workspace deployment environments in the source project.
	//
	// example:
	//
	// [{"DeploymentEnvironmentId":101,"Name":"environment-101","SourceProjectId":10,"TargetProjectId":20,"TargetProjectName":"target","Status":"Enabled"}]
	DeploymentEnvironments []*ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments `json:"DeploymentEnvironments,omitempty" xml:"DeploymentEnvironments,omitempty" type:"Repeated"`
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

func (s ListCrossProjectDeploymentEnvironmentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectDeploymentEnvironmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyData) GetDeploymentEnvironments() []*ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments {
	return s.DeploymentEnvironments
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyData) SetDeploymentEnvironments(v []*ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) *ListCrossProjectDeploymentEnvironmentsResponseBodyData {
	s.DeploymentEnvironments = v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyData) SetPageNumber(v int32) *ListCrossProjectDeploymentEnvironmentsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyData) SetPageSize(v int32) *ListCrossProjectDeploymentEnvironmentsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyData) SetRequestId(v string) *ListCrossProjectDeploymentEnvironmentsResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyData) SetTotalCount(v int32) *ListCrossProjectDeploymentEnvironmentsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyData) Validate() error {
	if s.DeploymentEnvironments != nil {
		for _, item := range s.DeploymentEnvironments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments struct {
	// The cross-workspace deployment environment ID.
	//
	// example:
	//
	// 101
	DeploymentEnvironmentId *int64 `json:"DeploymentEnvironmentId,omitempty" xml:"DeploymentEnvironmentId,omitempty"`
	// The environment name.
	//
	// example:
	//
	// environment-101
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source project workspace ID.
	//
	// example:
	//
	// 10
	SourceProjectId *int64 `json:"SourceProjectId,omitempty" xml:"SourceProjectId,omitempty"`
	// The environment status.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The target project workspace ID.
	//
	// example:
	//
	// 20
	TargetProjectId *int64 `json:"TargetProjectId,omitempty" xml:"TargetProjectId,omitempty"`
	// The target project workspace name.
	//
	// example:
	//
	// target
	TargetProjectName *string `json:"TargetProjectName,omitempty" xml:"TargetProjectName,omitempty"`
}

func (s ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) GoString() string {
	return s.String()
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) GetDeploymentEnvironmentId() *int64 {
	return s.DeploymentEnvironmentId
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) GetName() *string {
	return s.Name
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) GetSourceProjectId() *int64 {
	return s.SourceProjectId
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) GetStatus() *string {
	return s.Status
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) GetTargetProjectId() *int64 {
	return s.TargetProjectId
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) GetTargetProjectName() *string {
	return s.TargetProjectName
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) SetDeploymentEnvironmentId(v int64) *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments {
	s.DeploymentEnvironmentId = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) SetName(v string) *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments {
	s.Name = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) SetSourceProjectId(v int64) *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments {
	s.SourceProjectId = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) SetStatus(v string) *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments {
	s.Status = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) SetTargetProjectId(v int64) *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments {
	s.TargetProjectId = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) SetTargetProjectName(v string) *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments {
	s.TargetProjectName = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponseBodyDataDeploymentEnvironments) Validate() error {
	return dara.Validate(s)
}
