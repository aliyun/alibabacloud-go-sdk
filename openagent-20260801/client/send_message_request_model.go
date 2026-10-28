// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SendMessageRequest
	GetContent() *string
	SetResume(v []*SendMessageRequestResume) *SendMessageRequest
	GetResume() []*SendMessageRequestResume
	SetRunId(v string) *SendMessageRequest
	GetRunId() *string
	SetType(v string) *SendMessageRequest
	GetType() *string
}

type SendMessageRequest struct {
	// The message text in MESSAGE mode. This parameter cannot be empty. The type can be omitted in this mode.
	//
	// example:
	//
	// Query the status of the sample resource without making any changes
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The list of interrupt replies. This parameter is required when `type` is set to `RESUME`. You must reply to all pending interrupts at once, and each `interruptId` can appear only once.
	Resume []*SendMessageRequestResume `json:"resume,omitempty" xml:"resume,omitempty" type:"Repeated"`
	// The ID of the active run to resubscribe to. This parameter is required in RECONNECT mode.
	//
	// example:
	//
	// run_example_001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The request type. Valid values: `MESSAGE` (sends a message), `RESUME` (resumes a run), and `RECONNECT` (reconnects to a run). Default value: `MESSAGE`.
	//
	// example:
	//
	// MESSAGE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SendMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) GetContent() *string {
	return s.Content
}

func (s *SendMessageRequest) GetResume() []*SendMessageRequestResume {
	return s.Resume
}

func (s *SendMessageRequest) GetRunId() *string {
	return s.RunId
}

func (s *SendMessageRequest) GetType() *string {
	return s.Type
}

func (s *SendMessageRequest) SetContent(v string) *SendMessageRequest {
	s.Content = &v
	return s
}

func (s *SendMessageRequest) SetResume(v []*SendMessageRequestResume) *SendMessageRequest {
	s.Resume = v
	return s
}

func (s *SendMessageRequest) SetRunId(v string) *SendMessageRequest {
	s.RunId = &v
	return s
}

func (s *SendMessageRequest) SetType(v string) *SendMessageRequest {
	s.Type = &v
	return s
}

func (s *SendMessageRequest) Validate() error {
	if s.Resume != nil {
		for _, item := range s.Resume {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendMessageRequestResume struct {
	// The ID of the interrupt to be replied to. Obtain this value from `data.outcome.interrupts[].id` in the SSE `RUN_FINISHED` event or `data.outcome.interrupts[].id` in the `GetRun` response.
	//
	// example:
	//
	// interrupt_example_001
	InterruptId *string `json:"interruptId,omitempty" xml:"interruptId,omitempty"`
	// The interrupt handling result. This parameter is required when `status` is set to `resolved` and must conform to the `responseSchema` of the corresponding interrupt. Do not pass this parameter when `status` is set to `cancelled`.
	//
	// example:
	//
	// {"region":"cn-hangzhou"}
	Payload interface{} `json:"payload,omitempty" xml:"payload,omitempty"`
	// The reply method for the current interrupt. `resolved` indicates that the answer, confirmation, or tool result required by the interrupt is submitted, and `payload` must be provided. `cancelled` indicates that the interrupt is canceled, and `payload` is not passed.
	//
	// example:
	//
	// resolved
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SendMessageRequestResume) String() string {
	return dara.Prettify(s)
}

func (s SendMessageRequestResume) GoString() string {
	return s.String()
}

func (s *SendMessageRequestResume) GetInterruptId() *string {
	return s.InterruptId
}

func (s *SendMessageRequestResume) GetPayload() interface{} {
	return s.Payload
}

func (s *SendMessageRequestResume) GetStatus() *string {
	return s.Status
}

func (s *SendMessageRequestResume) SetInterruptId(v string) *SendMessageRequestResume {
	s.InterruptId = &v
	return s
}

func (s *SendMessageRequestResume) SetPayload(v interface{}) *SendMessageRequestResume {
	s.Payload = v
	return s
}

func (s *SendMessageRequestResume) SetStatus(v string) *SendMessageRequestResume {
	s.Status = &v
	return s
}

func (s *SendMessageRequestResume) Validate() error {
	return dara.Validate(s)
}
