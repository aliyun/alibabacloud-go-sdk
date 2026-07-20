// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSignalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreated(v string) *GetSignalResponseBody
	GetGmtCreated() *string
	SetGmtModified(v string) *GetSignalResponseBody
	GetGmtModified() *string
	SetJobId(v string) *GetSignalResponseBody
	GetJobId() *string
	SetMessage(v string) *GetSignalResponseBody
	GetMessage() *string
	SetPodNames(v []*string) *GetSignalResponseBody
	GetPodNames() []*string
	SetReason(v string) *GetSignalResponseBody
	GetReason() *string
	SetRequestId(v string) *GetSignalResponseBody
	GetRequestId() *string
	SetRoles(v []*string) *GetSignalResponseBody
	GetRoles() []*string
	SetScope(v string) *GetSignalResponseBody
	GetScope() *string
	SetSignal(v string) *GetSignalResponseBody
	GetSignal() *string
	SetSignalId(v string) *GetSignalResponseBody
	GetSignalId() *string
	SetStatus(v string) *GetSignalResponseBody
	GetStatus() *string
}

type GetSignalResponseBody struct {
	// The creation time.
	//
	// example:
	//
	// 2026-03-18T10:02:04+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-03-18T10:02:04+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The job ID.
	//
	// example:
	//
	// oper***********
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The status description, which contains a summary for each pod (number of successful deliveries, names of failed or pending pods, etc.).
	//
	// example:
	//
	// signal delivered to 1 pods
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of pod names.
	PodNames []*string `json:"PodNames,omitempty" xml:"PodNames,omitempty" type:"Repeated"`
	// The status reason code, such as `Completed`, `SignalFailed`, or `StoppedByJobEnded`.
	//
	// example:
	//
	// Completed
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-xxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of role objects.
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The delivery scope.
	//
	// example:
	//
	// pods
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The signal.
	//
	// example:
	//
	// SIGUSR1
	Signal *string `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The signal ID.
	//
	// example:
	//
	// dlc***********
	SignalId *string `json:"SignalId,omitempty" xml:"SignalId,omitempty"`
	// The signal status.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSignalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSignalResponseBody) GoString() string {
	return s.String()
}

func (s *GetSignalResponseBody) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *GetSignalResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetSignalResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetSignalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSignalResponseBody) GetPodNames() []*string {
	return s.PodNames
}

func (s *GetSignalResponseBody) GetReason() *string {
	return s.Reason
}

func (s *GetSignalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSignalResponseBody) GetRoles() []*string {
	return s.Roles
}

func (s *GetSignalResponseBody) GetScope() *string {
	return s.Scope
}

func (s *GetSignalResponseBody) GetSignal() *string {
	return s.Signal
}

func (s *GetSignalResponseBody) GetSignalId() *string {
	return s.SignalId
}

func (s *GetSignalResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSignalResponseBody) SetGmtCreated(v string) *GetSignalResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *GetSignalResponseBody) SetGmtModified(v string) *GetSignalResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetSignalResponseBody) SetJobId(v string) *GetSignalResponseBody {
	s.JobId = &v
	return s
}

func (s *GetSignalResponseBody) SetMessage(v string) *GetSignalResponseBody {
	s.Message = &v
	return s
}

func (s *GetSignalResponseBody) SetPodNames(v []*string) *GetSignalResponseBody {
	s.PodNames = v
	return s
}

func (s *GetSignalResponseBody) SetReason(v string) *GetSignalResponseBody {
	s.Reason = &v
	return s
}

func (s *GetSignalResponseBody) SetRequestId(v string) *GetSignalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSignalResponseBody) SetRoles(v []*string) *GetSignalResponseBody {
	s.Roles = v
	return s
}

func (s *GetSignalResponseBody) SetScope(v string) *GetSignalResponseBody {
	s.Scope = &v
	return s
}

func (s *GetSignalResponseBody) SetSignal(v string) *GetSignalResponseBody {
	s.Signal = &v
	return s
}

func (s *GetSignalResponseBody) SetSignalId(v string) *GetSignalResponseBody {
	s.SignalId = &v
	return s
}

func (s *GetSignalResponseBody) SetStatus(v string) *GetSignalResponseBody {
	s.Status = &v
	return s
}

func (s *GetSignalResponseBody) Validate() error {
	return dara.Validate(s)
}
