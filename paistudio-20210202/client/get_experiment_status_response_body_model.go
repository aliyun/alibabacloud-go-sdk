// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodes(v []*GetExperimentStatusResponseBodyNodes) *GetExperimentStatusResponseBody
	GetNodes() []*GetExperimentStatusResponseBodyNodes
	SetRequestId(v string) *GetExperimentStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetExperimentStatusResponseBody
	GetStatus() *string
}

type GetExperimentStatusResponseBody struct {
	Nodes []*GetExperimentStatusResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetExperimentStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentStatusResponseBody) GetNodes() []*GetExperimentStatusResponseBodyNodes {
	return s.Nodes
}

func (s *GetExperimentStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExperimentStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetExperimentStatusResponseBody) SetNodes(v []*GetExperimentStatusResponseBodyNodes) *GetExperimentStatusResponseBody {
	s.Nodes = v
	return s
}

func (s *GetExperimentStatusResponseBody) SetRequestId(v string) *GetExperimentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentStatusResponseBody) SetStatus(v string) *GetExperimentStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetExperimentStatusResponseBody) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetExperimentStatusResponseBodyNodes struct {
	// example:
	//
	// 2021-01-21T18:12:35.232Z
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	// example:
	//
	// job-mewqhd72nsrqujn1px
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// id-2317-1608984201281-74996
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// flow-wayrh3k605s7i51wey
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// example:
	//
	// node-6hc0eocnmzf4pk9njc
	RunNodeId *string `json:"RunNodeId,omitempty" xml:"RunNodeId,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	StartedAt *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetExperimentStatusResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentStatusResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *GetExperimentStatusResponseBodyNodes) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetExperimentStatusResponseBodyNodes) GetJobId() *string {
	return s.JobId
}

func (s *GetExperimentStatusResponseBodyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *GetExperimentStatusResponseBodyNodes) GetRunId() *string {
	return s.RunId
}

func (s *GetExperimentStatusResponseBodyNodes) GetRunNodeId() *string {
	return s.RunNodeId
}

func (s *GetExperimentStatusResponseBodyNodes) GetStartedAt() *string {
	return s.StartedAt
}

func (s *GetExperimentStatusResponseBodyNodes) GetStatus() *string {
	return s.Status
}

func (s *GetExperimentStatusResponseBodyNodes) SetFinishedAt(v string) *GetExperimentStatusResponseBodyNodes {
	s.FinishedAt = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) SetJobId(v string) *GetExperimentStatusResponseBodyNodes {
	s.JobId = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) SetNodeId(v string) *GetExperimentStatusResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) SetRunId(v string) *GetExperimentStatusResponseBodyNodes {
	s.RunId = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) SetRunNodeId(v string) *GetExperimentStatusResponseBodyNodes {
	s.RunNodeId = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) SetStartedAt(v string) *GetExperimentStatusResponseBodyNodes {
	s.StartedAt = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) SetStatus(v string) *GetExperimentStatusResponseBodyNodes {
	s.Status = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) Validate() error {
	return dara.Validate(s)
}
