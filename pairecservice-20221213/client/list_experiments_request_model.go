// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExperimentGroupId(v string) *ListExperimentsRequest
	GetExperimentGroupId() *string
	SetInstanceId(v string) *ListExperimentsRequest
	GetInstanceId() *string
	SetQuery(v string) *ListExperimentsRequest
	GetQuery() *string
	SetStatus(v string) *ListExperimentsRequest
	GetStatus() *string
}

type ListExperimentsRequest struct {
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// experiment_test1
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListExperimentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentsRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentsRequest) GetExperimentGroupId() *string {
	return s.ExperimentGroupId
}

func (s *ListExperimentsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListExperimentsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListExperimentsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListExperimentsRequest) SetExperimentGroupId(v string) *ListExperimentsRequest {
	s.ExperimentGroupId = &v
	return s
}

func (s *ListExperimentsRequest) SetInstanceId(v string) *ListExperimentsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExperimentsRequest) SetQuery(v string) *ListExperimentsRequest {
	s.Query = &v
	return s
}

func (s *ListExperimentsRequest) SetStatus(v string) *ListExperimentsRequest {
	s.Status = &v
	return s
}

func (s *ListExperimentsRequest) Validate() error {
	return dara.Validate(s)
}
