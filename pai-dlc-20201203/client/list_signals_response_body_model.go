// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSignalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *ListSignalsResponseBody
	GetJobId() *string
	SetPageNumber(v int64) *ListSignalsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSignalsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListSignalsResponseBody
	GetRequestId() *string
	SetSignals(v []*ListSignalsResponseBodySignals) *ListSignalsResponseBody
	GetSignals() []*ListSignalsResponseBodySignals
	SetTotalCount(v int64) *ListSignalsResponseBody
	GetTotalCount() *int64
}

type ListSignalsResponseBody struct {
	// example:
	//
	// dlc-...
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-xxxxxxx
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signals   []*ListSignalsResponseBodySignals `json:"Signals,omitempty" xml:"Signals,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSignalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSignalsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSignalsResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *ListSignalsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSignalsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSignalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSignalsResponseBody) GetSignals() []*ListSignalsResponseBodySignals {
	return s.Signals
}

func (s *ListSignalsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSignalsResponseBody) SetJobId(v string) *ListSignalsResponseBody {
	s.JobId = &v
	return s
}

func (s *ListSignalsResponseBody) SetPageNumber(v int64) *ListSignalsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSignalsResponseBody) SetPageSize(v int64) *ListSignalsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSignalsResponseBody) SetRequestId(v string) *ListSignalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSignalsResponseBody) SetSignals(v []*ListSignalsResponseBodySignals) *ListSignalsResponseBody {
	s.Signals = v
	return s
}

func (s *ListSignalsResponseBody) SetTotalCount(v int64) *ListSignalsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSignalsResponseBody) Validate() error {
	if s.Signals != nil {
		for _, item := range s.Signals {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSignalsResponseBodySignals struct {
	// example:
	//
	// 2025-12-30T14:07:38+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2025-12-30T14:07:38+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// signal delivered to 1 pods
	Message  *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	PodNames []*string `json:"PodNames,omitempty" xml:"PodNames,omitempty" type:"Repeated"`
	// example:
	//
	// Completed
	Reason *string   `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Roles  []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// pods
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// SIGUSR1
	Signal *string `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// example:
	//
	// oper*********
	SignalId *string `json:"SignalId,omitempty" xml:"SignalId,omitempty"`
	// example:
	//
	// Stopped
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSignalsResponseBodySignals) String() string {
	return dara.Prettify(s)
}

func (s ListSignalsResponseBodySignals) GoString() string {
	return s.String()
}

func (s *ListSignalsResponseBodySignals) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListSignalsResponseBodySignals) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListSignalsResponseBodySignals) GetMessage() *string {
	return s.Message
}

func (s *ListSignalsResponseBodySignals) GetPodNames() []*string {
	return s.PodNames
}

func (s *ListSignalsResponseBodySignals) GetReason() *string {
	return s.Reason
}

func (s *ListSignalsResponseBodySignals) GetRoles() []*string {
	return s.Roles
}

func (s *ListSignalsResponseBodySignals) GetScope() *string {
	return s.Scope
}

func (s *ListSignalsResponseBodySignals) GetSignal() *string {
	return s.Signal
}

func (s *ListSignalsResponseBodySignals) GetSignalId() *string {
	return s.SignalId
}

func (s *ListSignalsResponseBodySignals) GetStatus() *string {
	return s.Status
}

func (s *ListSignalsResponseBodySignals) SetGmtCreated(v string) *ListSignalsResponseBodySignals {
	s.GmtCreated = &v
	return s
}

func (s *ListSignalsResponseBodySignals) SetGmtModified(v string) *ListSignalsResponseBodySignals {
	s.GmtModified = &v
	return s
}

func (s *ListSignalsResponseBodySignals) SetMessage(v string) *ListSignalsResponseBodySignals {
	s.Message = &v
	return s
}

func (s *ListSignalsResponseBodySignals) SetPodNames(v []*string) *ListSignalsResponseBodySignals {
	s.PodNames = v
	return s
}

func (s *ListSignalsResponseBodySignals) SetReason(v string) *ListSignalsResponseBodySignals {
	s.Reason = &v
	return s
}

func (s *ListSignalsResponseBodySignals) SetRoles(v []*string) *ListSignalsResponseBodySignals {
	s.Roles = v
	return s
}

func (s *ListSignalsResponseBodySignals) SetScope(v string) *ListSignalsResponseBodySignals {
	s.Scope = &v
	return s
}

func (s *ListSignalsResponseBodySignals) SetSignal(v string) *ListSignalsResponseBodySignals {
	s.Signal = &v
	return s
}

func (s *ListSignalsResponseBodySignals) SetSignalId(v string) *ListSignalsResponseBodySignals {
	s.SignalId = &v
	return s
}

func (s *ListSignalsResponseBodySignals) SetStatus(v string) *ListSignalsResponseBodySignals {
	s.Status = &v
	return s
}

func (s *ListSignalsResponseBodySignals) Validate() error {
	return dara.Validate(s)
}
