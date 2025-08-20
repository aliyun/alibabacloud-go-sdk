// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventCenterRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventType(v string) *ListEventCenterRecordRequest
	GetEventType() *string
	SetInstanceId(v string) *ListEventCenterRecordRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListEventCenterRecordRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListEventCenterRecordRequest
	GetPageSize() *int32
	SetRepoName(v string) *ListEventCenterRecordRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *ListEventCenterRecordRequest
	GetRepoNamespaceName() *string
	SetRuleId(v string) *ListEventCenterRecordRequest
	GetRuleId() *string
}

type ListEventCenterRecordRequest struct {
	// The type of the event. Valid values:
	//
	// 	- `cr:Artifact:DeliveryChainCompleted`: The delivery chain is processed.
	//
	// 	- `cr:Artifact:SynchronizationCompleted`: The image is replicated.
	//
	// 	- `cr:Artifact:BuildCompleted`: The image is built.
	//
	// 	- `cr:Artifact:ScanCompleted`: The image is scanned.
	//
	// 	- `cr:Artifact:SigningCompleted`: The image is signed.
	//
	// example:
	//
	// cr:Artifact:DeliveryChainCompleted
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The ID of the event notification rule.
	//
	// example:
	//
	// crecr-n6pbhgjxtla***
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListEventCenterRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventCenterRecordRequest) GoString() string {
	return s.String()
}

func (s *ListEventCenterRecordRequest) GetEventType() *string {
	return s.EventType
}

func (s *ListEventCenterRecordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventCenterRecordRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListEventCenterRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEventCenterRecordRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *ListEventCenterRecordRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListEventCenterRecordRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *ListEventCenterRecordRequest) SetEventType(v string) *ListEventCenterRecordRequest {
	s.EventType = &v
	return s
}

func (s *ListEventCenterRecordRequest) SetInstanceId(v string) *ListEventCenterRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *ListEventCenterRecordRequest) SetPageNo(v int32) *ListEventCenterRecordRequest {
	s.PageNo = &v
	return s
}

func (s *ListEventCenterRecordRequest) SetPageSize(v int32) *ListEventCenterRecordRequest {
	s.PageSize = &v
	return s
}

func (s *ListEventCenterRecordRequest) SetRepoName(v string) *ListEventCenterRecordRequest {
	s.RepoName = &v
	return s
}

func (s *ListEventCenterRecordRequest) SetRepoNamespaceName(v string) *ListEventCenterRecordRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListEventCenterRecordRequest) SetRuleId(v string) *ListEventCenterRecordRequest {
	s.RuleId = &v
	return s
}

func (s *ListEventCenterRecordRequest) Validate() error {
	return dara.Validate(s)
}
