// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggTaskGroupStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggTaskGroupConfigHash(v string) *UpdateAggTaskGroupStatusResponseBody
	GetAggTaskGroupConfigHash() *string
	SetAggTaskGroupId(v string) *UpdateAggTaskGroupStatusResponseBody
	GetAggTaskGroupId() *string
	SetAggTaskGroupName(v string) *UpdateAggTaskGroupStatusResponseBody
	GetAggTaskGroupName() *string
	SetRequestId(v string) *UpdateAggTaskGroupStatusResponseBody
	GetRequestId() *string
	SetSourcePrometheusId(v string) *UpdateAggTaskGroupStatusResponseBody
	GetSourcePrometheusId() *string
	SetStatus(v string) *UpdateAggTaskGroupStatusResponseBody
	GetStatus() *string
}

type UpdateAggTaskGroupStatusResponseBody struct {
	// example:
	//
	// a54136014dc386a92c83a6ef1e97ff22
	AggTaskGroupConfigHash *string `json:"aggTaskGroupConfigHash,omitempty" xml:"aggTaskGroupConfigHash,omitempty"`
	// example:
	//
	// aggTaskGroup-xxx
	AggTaskGroupId *string `json:"aggTaskGroupId,omitempty" xml:"aggTaskGroupId,omitempty"`
	// example:
	//
	// pipeline-aggtask-group
	AggTaskGroupName *string `json:"aggTaskGroupName,omitempty" xml:"aggTaskGroupName,omitempty"`
	// example:
	//
	// 0CEC5375-C554-562B-A65F-***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// rw-xxx
	SourcePrometheusId *string `json:"sourcePrometheusId,omitempty" xml:"sourcePrometheusId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateAggTaskGroupStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggTaskGroupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAggTaskGroupStatusResponseBody) GetAggTaskGroupConfigHash() *string {
	return s.AggTaskGroupConfigHash
}

func (s *UpdateAggTaskGroupStatusResponseBody) GetAggTaskGroupId() *string {
	return s.AggTaskGroupId
}

func (s *UpdateAggTaskGroupStatusResponseBody) GetAggTaskGroupName() *string {
	return s.AggTaskGroupName
}

func (s *UpdateAggTaskGroupStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAggTaskGroupStatusResponseBody) GetSourcePrometheusId() *string {
	return s.SourcePrometheusId
}

func (s *UpdateAggTaskGroupStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateAggTaskGroupStatusResponseBody) SetAggTaskGroupConfigHash(v string) *UpdateAggTaskGroupStatusResponseBody {
	s.AggTaskGroupConfigHash = &v
	return s
}

func (s *UpdateAggTaskGroupStatusResponseBody) SetAggTaskGroupId(v string) *UpdateAggTaskGroupStatusResponseBody {
	s.AggTaskGroupId = &v
	return s
}

func (s *UpdateAggTaskGroupStatusResponseBody) SetAggTaskGroupName(v string) *UpdateAggTaskGroupStatusResponseBody {
	s.AggTaskGroupName = &v
	return s
}

func (s *UpdateAggTaskGroupStatusResponseBody) SetRequestId(v string) *UpdateAggTaskGroupStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAggTaskGroupStatusResponseBody) SetSourcePrometheusId(v string) *UpdateAggTaskGroupStatusResponseBody {
	s.SourcePrometheusId = &v
	return s
}

func (s *UpdateAggTaskGroupStatusResponseBody) SetStatus(v string) *UpdateAggTaskGroupStatusResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateAggTaskGroupStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
