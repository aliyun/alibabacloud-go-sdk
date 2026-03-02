// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncResourcePlanOperationResult interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *AsyncResourcePlanOperationResult
	GetMessage() *string
	SetPlan(v string) *AsyncResourcePlanOperationResult
	GetPlan() *string
	SetTicketStatus(v string) *AsyncResourcePlanOperationResult
	GetTicketStatus() *string
}

type AsyncResourcePlanOperationResult struct {
	// The information about the ticket. The value of this parameter is returned when the value of the ticketStatus parameter is FAILED or EXECUTING.
	//
	// example:
	//
	// "create resource plan failed"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The resource configuration plan of the deployment in expert mode. The value of this parameter is returned when the value of the ticketStatus parameter is FINISHED.
	//
	// example:
	//
	// {\\"ssgProfiles\\":[{\\"name\\":\\"default\\",\\"cpu\\":1.13,\\"heap\\":\\"1 gb\\",\\"offHeap\\":\\"32 mb\\",\\"managed\\":{},\\"extended\\":{}}],\\"nodes\\":[{\\"id\\":1,\\"type\\":\\"StreamExecTableSourceScan\\",\\"desc\\":\\"Source: datagen_source[78]\\",\\"profile\\":{\\"group\\":\\"default\\",\\"parallelism\\":1,\\"maxParallelism\\":32768,\\"minParallelism\\":1}},{\\"id\\":2,\\"type\\":\\"StreamExecSink\\",\\"desc\\":\\"Sink: blackhole_sink[79]\\",\\"profile\\":{\\"group\\":\\"default\\",\\"parallelism\\":1,\\"maxParallelism\\":32768,\\"minParallelism\\":1}}],\\"edges\\":[{\\"source\\":1,\\"target\\":2,\\"mode\\":\\"PIPELINED\\",\\"strategy\\":\\"FORWARD\\"}],\\"vertices\\":{\\"717c7b8afebbfb7137f6f0f99beb2a94\\":[1,2]}}
	Plan *string `json:"plan,omitempty" xml:"plan,omitempty"`
	// The status of the ticket that applies for an asynchronous operation. Valid values:
	//
	// 	- EXECUTING
	//
	// 	- FINISHED
	//
	// 	- FAILED
	//
	// example:
	//
	// FINISHED
	TicketStatus *string `json:"ticketStatus,omitempty" xml:"ticketStatus,omitempty"`
}

func (s AsyncResourcePlanOperationResult) String() string {
	return dara.Prettify(s)
}

func (s AsyncResourcePlanOperationResult) GoString() string {
	return s.String()
}

func (s *AsyncResourcePlanOperationResult) GetMessage() *string {
	return s.Message
}

func (s *AsyncResourcePlanOperationResult) GetPlan() *string {
	return s.Plan
}

func (s *AsyncResourcePlanOperationResult) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *AsyncResourcePlanOperationResult) SetMessage(v string) *AsyncResourcePlanOperationResult {
	s.Message = &v
	return s
}

func (s *AsyncResourcePlanOperationResult) SetPlan(v string) *AsyncResourcePlanOperationResult {
	s.Plan = &v
	return s
}

func (s *AsyncResourcePlanOperationResult) SetTicketStatus(v string) *AsyncResourcePlanOperationResult {
	s.TicketStatus = &v
	return s
}

func (s *AsyncResourcePlanOperationResult) Validate() error {
	return dara.Validate(s)
}
