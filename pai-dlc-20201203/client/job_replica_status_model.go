// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobReplicaStatus interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int32) *JobReplicaStatus
	GetActive() *int32
	SetType(v string) *JobReplicaStatus
	GetType() *string
}

type JobReplicaStatus struct {
	Active *int32  `json:"Active,omitempty" xml:"Active,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s JobReplicaStatus) String() string {
	return dara.Prettify(s)
}

func (s JobReplicaStatus) GoString() string {
	return s.String()
}

func (s *JobReplicaStatus) GetActive() *int32 {
	return s.Active
}

func (s *JobReplicaStatus) GetType() *string {
	return s.Type
}

func (s *JobReplicaStatus) SetActive(v int32) *JobReplicaStatus {
	s.Active = &v
	return s
}

func (s *JobReplicaStatus) SetType(v string) *JobReplicaStatus {
	s.Type = &v
	return s
}

func (s *JobReplicaStatus) Validate() error {
	return dara.Validate(s)
}
