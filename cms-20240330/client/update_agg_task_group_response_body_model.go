// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggTaskGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggTaskGroupConfigHash(v string) *UpdateAggTaskGroupResponseBody
	GetAggTaskGroupConfigHash() *string
	SetAggTaskGroupId(v string) *UpdateAggTaskGroupResponseBody
	GetAggTaskGroupId() *string
	SetAggTaskGroupName(v string) *UpdateAggTaskGroupResponseBody
	GetAggTaskGroupName() *string
	SetRequestId(v string) *UpdateAggTaskGroupResponseBody
	GetRequestId() *string
	SetSourcePrometheusId(v string) *UpdateAggTaskGroupResponseBody
	GetSourcePrometheusId() *string
	SetStatus(v string) *UpdateAggTaskGroupResponseBody
	GetStatus() *string
}

type UpdateAggTaskGroupResponseBody struct {
	// Summary of the aggregation task group configuration.
	//
	// example:
	//
	// a54136014dc386a92c83a6ef1e97ff22
	AggTaskGroupConfigHash *string `json:"aggTaskGroupConfigHash,omitempty" xml:"aggTaskGroupConfigHash,omitempty"`
	// Aggregation task group ID
	//
	// example:
	//
	// aggTaskGroup-5fb2c3ade63a4709bcb059d13493b7b8
	AggTaskGroupId *string `json:"aggTaskGroupId,omitempty" xml:"aggTaskGroupId,omitempty"`
	// Aggregation task group name
	//
	// example:
	//
	// pipeline-aggtask-group
	AggTaskGroupName *string `json:"aggTaskGroupName,omitempty" xml:"aggTaskGroupName,omitempty"`
	// Request ID
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Source Prometheus instance ID of the aggregation task group
	//
	// example:
	//
	// rw-083e17834e279f8c627fe91a2d72
	SourcePrometheusId *string `json:"sourcePrometheusId,omitempty" xml:"sourcePrometheusId,omitempty"`
	// Current status of the aggregation task group
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateAggTaskGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggTaskGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAggTaskGroupResponseBody) GetAggTaskGroupConfigHash() *string {
	return s.AggTaskGroupConfigHash
}

func (s *UpdateAggTaskGroupResponseBody) GetAggTaskGroupId() *string {
	return s.AggTaskGroupId
}

func (s *UpdateAggTaskGroupResponseBody) GetAggTaskGroupName() *string {
	return s.AggTaskGroupName
}

func (s *UpdateAggTaskGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAggTaskGroupResponseBody) GetSourcePrometheusId() *string {
	return s.SourcePrometheusId
}

func (s *UpdateAggTaskGroupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateAggTaskGroupResponseBody) SetAggTaskGroupConfigHash(v string) *UpdateAggTaskGroupResponseBody {
	s.AggTaskGroupConfigHash = &v
	return s
}

func (s *UpdateAggTaskGroupResponseBody) SetAggTaskGroupId(v string) *UpdateAggTaskGroupResponseBody {
	s.AggTaskGroupId = &v
	return s
}

func (s *UpdateAggTaskGroupResponseBody) SetAggTaskGroupName(v string) *UpdateAggTaskGroupResponseBody {
	s.AggTaskGroupName = &v
	return s
}

func (s *UpdateAggTaskGroupResponseBody) SetRequestId(v string) *UpdateAggTaskGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAggTaskGroupResponseBody) SetSourcePrometheusId(v string) *UpdateAggTaskGroupResponseBody {
	s.SourcePrometheusId = &v
	return s
}

func (s *UpdateAggTaskGroupResponseBody) SetStatus(v string) *UpdateAggTaskGroupResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateAggTaskGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
