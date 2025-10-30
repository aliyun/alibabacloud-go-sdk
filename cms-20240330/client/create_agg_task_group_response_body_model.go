// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggTaskGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggTaskGroupConfigHash(v string) *CreateAggTaskGroupResponseBody
	GetAggTaskGroupConfigHash() *string
	SetAggTaskGroupId(v string) *CreateAggTaskGroupResponseBody
	GetAggTaskGroupId() *string
	SetAggTaskGroupName(v string) *CreateAggTaskGroupResponseBody
	GetAggTaskGroupName() *string
	SetRequestId(v string) *CreateAggTaskGroupResponseBody
	GetRequestId() *string
	SetSourcePrometheusId(v string) *CreateAggTaskGroupResponseBody
	GetSourcePrometheusId() *string
	SetStatus(v string) *CreateAggTaskGroupResponseBody
	GetStatus() *string
}

type CreateAggTaskGroupResponseBody struct {
	// Summary of the aggregation task group configuration.
	//
	// example:
	//
	// a54136014dc386a92c83a6ef1e97ff22
	AggTaskGroupConfigHash *string `json:"aggTaskGroupConfigHash,omitempty" xml:"aggTaskGroupConfigHash,omitempty"`
	// Aggregation task group ID.
	//
	// example:
	//
	// aggTaskGroup-f4b8e50525cf41c894488c0c71ec483f
	AggTaskGroupId *string `json:"aggTaskGroupId,omitempty" xml:"aggTaskGroupId,omitempty"`
	// Aggregation task group name.
	//
	// example:
	//
	// pipeline-aggtask-group
	AggTaskGroupName *string `json:"aggTaskGroupName,omitempty" xml:"aggTaskGroupName,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 16C0A6D6-C3E7-511D-A60B-A87FD85F5BA7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Source Prometheus instance ID of the aggregation task group.
	//
	// example:
	//
	// rw-ecc04af14729b1a16e40a0d10068
	SourcePrometheusId *string `json:"sourcePrometheusId,omitempty" xml:"sourcePrometheusId,omitempty"`
	// Current status of the aggregation task group.
	//
	// example:
	//
	// Pending2Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateAggTaskGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAggTaskGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAggTaskGroupResponseBody) GetAggTaskGroupConfigHash() *string {
	return s.AggTaskGroupConfigHash
}

func (s *CreateAggTaskGroupResponseBody) GetAggTaskGroupId() *string {
	return s.AggTaskGroupId
}

func (s *CreateAggTaskGroupResponseBody) GetAggTaskGroupName() *string {
	return s.AggTaskGroupName
}

func (s *CreateAggTaskGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAggTaskGroupResponseBody) GetSourcePrometheusId() *string {
	return s.SourcePrometheusId
}

func (s *CreateAggTaskGroupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateAggTaskGroupResponseBody) SetAggTaskGroupConfigHash(v string) *CreateAggTaskGroupResponseBody {
	s.AggTaskGroupConfigHash = &v
	return s
}

func (s *CreateAggTaskGroupResponseBody) SetAggTaskGroupId(v string) *CreateAggTaskGroupResponseBody {
	s.AggTaskGroupId = &v
	return s
}

func (s *CreateAggTaskGroupResponseBody) SetAggTaskGroupName(v string) *CreateAggTaskGroupResponseBody {
	s.AggTaskGroupName = &v
	return s
}

func (s *CreateAggTaskGroupResponseBody) SetRequestId(v string) *CreateAggTaskGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAggTaskGroupResponseBody) SetSourcePrometheusId(v string) *CreateAggTaskGroupResponseBody {
	s.SourcePrometheusId = &v
	return s
}

func (s *CreateAggTaskGroupResponseBody) SetStatus(v string) *CreateAggTaskGroupResponseBody {
	s.Status = &v
	return s
}

func (s *CreateAggTaskGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
