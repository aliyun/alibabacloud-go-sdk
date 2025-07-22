// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetAgentResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetAgentResponseBody
	GetStatus() *string
	SetStopTime(v string) *GetAgentResponseBody
	GetStopTime() *string
}

type GetAgentResponseBody struct {
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1751513144838
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1751513144838
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAgentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetAgentResponseBody) GetStopTime() *string {
	return s.StopTime
}

func (s *GetAgentResponseBody) SetMessage(v string) *GetAgentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentResponseBody) SetRequestId(v string) *GetAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentResponseBody) SetStartTime(v string) *GetAgentResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetAgentResponseBody) SetStatus(v string) *GetAgentResponseBody {
	s.Status = &v
	return s
}

func (s *GetAgentResponseBody) SetStopTime(v string) *GetAgentResponseBody {
	s.StopTime = &v
	return s
}

func (s *GetAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
