// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimeMs(v int64) *ListInstancesRequest
	GetEndTimeMs() *int64
	SetInstanceIds(v []*string) *ListInstancesRequest
	GetInstanceIds() []*string
	SetInstanceStatus(v []*string) *ListInstancesRequest
	GetInstanceStatus() []*string
	SetLimit(v string) *ListInstancesRequest
	GetLimit() *string
	SetQualifier(v string) *ListInstancesRequest
	GetQualifier() *string
	SetStartKey(v string) *ListInstancesRequest
	GetStartKey() *string
	SetStartTimeMs(v int64) *ListInstancesRequest
	GetStartTimeMs() *int64
	SetWithAllActive(v bool) *ListInstancesRequest
	GetWithAllActive() *bool
}

type ListInstancesRequest struct {
	EndTimeMs      *int64    `json:"endTimeMs,omitempty" xml:"endTimeMs,omitempty"`
	InstanceIds    []*string `json:"instanceIds" xml:"instanceIds" type:"Repeated"`
	InstanceStatus []*string `json:"instanceStatus" xml:"instanceStatus" type:"Repeated"`
	Limit          *string   `json:"limit,omitempty" xml:"limit,omitempty"`
	// The function version or alias.
	//
	// example:
	//
	// LATEST
	Qualifier   *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	StartKey    *string `json:"startKey,omitempty" xml:"startKey,omitempty"`
	StartTimeMs *int64  `json:"startTimeMs,omitempty" xml:"startTimeMs,omitempty"`
	// Specifies whether to list all instances. Valid values: true and false.
	//
	// example:
	//
	// true
	WithAllActive *bool `json:"withAllActive,omitempty" xml:"withAllActive,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetEndTimeMs() *int64 {
	return s.EndTimeMs
}

func (s *ListInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListInstancesRequest) GetInstanceStatus() []*string {
	return s.InstanceStatus
}

func (s *ListInstancesRequest) GetLimit() *string {
	return s.Limit
}

func (s *ListInstancesRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *ListInstancesRequest) GetStartKey() *string {
	return s.StartKey
}

func (s *ListInstancesRequest) GetStartTimeMs() *int64 {
	return s.StartTimeMs
}

func (s *ListInstancesRequest) GetWithAllActive() *bool {
	return s.WithAllActive
}

func (s *ListInstancesRequest) SetEndTimeMs(v int64) *ListInstancesRequest {
	s.EndTimeMs = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceIds(v []*string) *ListInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *ListInstancesRequest) SetInstanceStatus(v []*string) *ListInstancesRequest {
	s.InstanceStatus = v
	return s
}

func (s *ListInstancesRequest) SetLimit(v string) *ListInstancesRequest {
	s.Limit = &v
	return s
}

func (s *ListInstancesRequest) SetQualifier(v string) *ListInstancesRequest {
	s.Qualifier = &v
	return s
}

func (s *ListInstancesRequest) SetStartKey(v string) *ListInstancesRequest {
	s.StartKey = &v
	return s
}

func (s *ListInstancesRequest) SetStartTimeMs(v int64) *ListInstancesRequest {
	s.StartTimeMs = &v
	return s
}

func (s *ListInstancesRequest) SetWithAllActive(v bool) *ListInstancesRequest {
	s.WithAllActive = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
