// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicDisposalProcessesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListDynamicDisposalProcessesRequest
	GetCurrentPage() *int64
	SetDevTag(v string) *ListDynamicDisposalProcessesRequest
	GetDevTag() *string
	SetDisposalAction(v string) *ListDynamicDisposalProcessesRequest
	GetDisposalAction() *string
	SetDisposalProcessId(v string) *ListDynamicDisposalProcessesRequest
	GetDisposalProcessId() *string
	SetEndTime(v int64) *ListDynamicDisposalProcessesRequest
	GetEndTime() *int64
	SetPageSize(v int64) *ListDynamicDisposalProcessesRequest
	GetPageSize() *int64
	SetRecoveryType(v string) *ListDynamicDisposalProcessesRequest
	GetRecoveryType() *string
	SetStartTime(v int64) *ListDynamicDisposalProcessesRequest
	GetStartTime() *int64
	SetStatus(v string) *ListDynamicDisposalProcessesRequest
	GetStatus() *string
	SetUserName(v string) *ListDynamicDisposalProcessesRequest
	GetUserName() *string
}

type ListDynamicDisposalProcessesRequest struct {
	// The page number to display in the paginated query. Range: 1~10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Terminal device ID.
	//
	// example:
	//
	// E7798391-2554-FE83-E0B7-045DDED629A8
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// Disposal action.
	//
	// - **ztna_connect**: Prohibit connection to the zero-trust intranet.
	//
	// - **nac_connect**: Prohibit connection to the office network access.
	//
	// - **none**: No disposal action.
	//
	// example:
	//
	// none
	DisposalAction *string `json:"DisposalAction,omitempty" xml:"DisposalAction,omitempty"`
	// Disposal process ID.
	//
	// example:
	//
	// dp-xxxxxxxx
	DisposalProcessId *string `json:"DisposalProcessId,omitempty" xml:"DisposalProcessId,omitempty"`
	// The end time for querying dynamic disposal processes. Format: Unix timestamp (in seconds).
	//
	// example:
	//
	// 1743143296
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of items per page in the paginated query. Range: 1~1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Recovery type.
	//
	// - **auto**: Automatic recovery.
	//
	// - **console**: Console recovery.
	//
	// - **auth**: Recovery by authentication and reporting.
	//
	// example:
	//
	// auto
	RecoveryType *string `json:"RecoveryType,omitempty" xml:"RecoveryType,omitempty"`
	// The start time for querying dynamic disposal processes. Format: Unix timestamp (in seconds).
	//
	// example:
	//
	// 1743143296
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Disposal status. Values:
	//
	// - **disposal**: In the disposal state.
	//
	// - **finished**: Already automatically recovered.
	//
	// - **recovery**: Recovered by authentication and reporting or console recovery.
	//
	// example:
	//
	// disposal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Username.
	//
	// example:
	//
	// xiaoming
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListDynamicDisposalProcessesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicDisposalProcessesRequest) GoString() string {
	return s.String()
}

func (s *ListDynamicDisposalProcessesRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListDynamicDisposalProcessesRequest) GetDevTag() *string {
	return s.DevTag
}

func (s *ListDynamicDisposalProcessesRequest) GetDisposalAction() *string {
	return s.DisposalAction
}

func (s *ListDynamicDisposalProcessesRequest) GetDisposalProcessId() *string {
	return s.DisposalProcessId
}

func (s *ListDynamicDisposalProcessesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDynamicDisposalProcessesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDynamicDisposalProcessesRequest) GetRecoveryType() *string {
	return s.RecoveryType
}

func (s *ListDynamicDisposalProcessesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDynamicDisposalProcessesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDynamicDisposalProcessesRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListDynamicDisposalProcessesRequest) SetCurrentPage(v int64) *ListDynamicDisposalProcessesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDynamicDisposalProcessesRequest) SetDevTag(v string) *ListDynamicDisposalProcessesRequest {
	s.DevTag = &v
	return s
}

func (s *ListDynamicDisposalProcessesRequest) SetDisposalAction(v string) *ListDynamicDisposalProcessesRequest {
	s.DisposalAction = &v
	return s
}

func (s *ListDynamicDisposalProcessesRequest) SetDisposalProcessId(v string) *ListDynamicDisposalProcessesRequest {
	s.DisposalProcessId = &v
	return s
}

func (s *ListDynamicDisposalProcessesRequest) SetEndTime(v int64) *ListDynamicDisposalProcessesRequest {
	s.EndTime = &v
	return s
}

func (s *ListDynamicDisposalProcessesRequest) SetPageSize(v int64) *ListDynamicDisposalProcessesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDynamicDisposalProcessesRequest) SetRecoveryType(v string) *ListDynamicDisposalProcessesRequest {
	s.RecoveryType = &v
	return s
}

func (s *ListDynamicDisposalProcessesRequest) SetStartTime(v int64) *ListDynamicDisposalProcessesRequest {
	s.StartTime = &v
	return s
}

func (s *ListDynamicDisposalProcessesRequest) SetStatus(v string) *ListDynamicDisposalProcessesRequest {
	s.Status = &v
	return s
}

func (s *ListDynamicDisposalProcessesRequest) SetUserName(v string) *ListDynamicDisposalProcessesRequest {
	s.UserName = &v
	return s
}

func (s *ListDynamicDisposalProcessesRequest) Validate() error {
	return dara.Validate(s)
}
