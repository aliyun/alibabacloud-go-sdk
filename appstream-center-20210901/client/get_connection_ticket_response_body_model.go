// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *GetConnectionTicketResponseBody
	GetAppInstanceGroupId() *string
	SetAppInstanceId(v string) *GetConnectionTicketResponseBody
	GetAppInstanceId() *string
	SetAppInstancePersistentId(v string) *GetConnectionTicketResponseBody
	GetAppInstancePersistentId() *string
	SetAvatarId(v string) *GetConnectionTicketResponseBody
	GetAvatarId() *string
	SetBizRegionId(v string) *GetConnectionTicketResponseBody
	GetBizRegionId() *string
	SetOsType(v string) *GetConnectionTicketResponseBody
	GetOsType() *string
	SetRequestId(v string) *GetConnectionTicketResponseBody
	GetRequestId() *string
	SetTaskId(v string) *GetConnectionTicketResponseBody
	GetTaskId() *string
	SetTaskStatus(v string) *GetConnectionTicketResponseBody
	GetTaskStatus() *string
	SetTenantId(v int64) *GetConnectionTicketResponseBody
	GetTenantId() *int64
	SetTicket(v string) *GetConnectionTicketResponseBody
	GetTicket() *string
}

type GetConnectionTicketResponseBody struct {
	// The delivery group ID.
	//
	// example:
	//
	// aig-53fvrq1oan****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The application instance ID.
	//
	// example:
	//
	// ai-7ybdeiyoeh5e****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The persistent session ID.
	//
	// example:
	//
	// p-0bxls9m3cl7s****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// The avatar ID.
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The operating system type.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// f3d1b31c-605e-4d04-a896****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task status.
	//
	// example:
	//
	// Running
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The tenant ID (Alibaba Cloud account UID).
	//
	// example:
	//
	// 148871678899****
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The credentials for connecting to the cloud application.
	//
	// > This parameter is returned only on non-initial calls.
	//
	// example:
	//
	// DQpbRGVza3RvcF0NCkZvcmNlVGxzVHlwZT0xDQpHV1Rva2VuUGFydDE9MDAva09ROW1FUTU3dU****
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetConnectionTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBody) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *GetConnectionTicketResponseBody) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *GetConnectionTicketResponseBody) GetAppInstancePersistentId() *string {
	return s.AppInstancePersistentId
}

func (s *GetConnectionTicketResponseBody) GetAvatarId() *string {
	return s.AvatarId
}

func (s *GetConnectionTicketResponseBody) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *GetConnectionTicketResponseBody) GetOsType() *string {
	return s.OsType
}

func (s *GetConnectionTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConnectionTicketResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetConnectionTicketResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetConnectionTicketResponseBody) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetConnectionTicketResponseBody) GetTicket() *string {
	return s.Ticket
}

func (s *GetConnectionTicketResponseBody) SetAppInstanceGroupId(v string) *GetConnectionTicketResponseBody {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAppInstanceId(v string) *GetConnectionTicketResponseBody {
	s.AppInstanceId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAppInstancePersistentId(v string) *GetConnectionTicketResponseBody {
	s.AppInstancePersistentId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAvatarId(v string) *GetConnectionTicketResponseBody {
	s.AvatarId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetBizRegionId(v string) *GetConnectionTicketResponseBody {
	s.BizRegionId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetOsType(v string) *GetConnectionTicketResponseBody {
	s.OsType = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRequestId(v string) *GetConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskId(v string) *GetConnectionTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskStatus(v string) *GetConnectionTicketResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTenantId(v int64) *GetConnectionTicketResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTicket(v string) *GetConnectionTicketResponseBody {
	s.Ticket = &v
	return s
}

func (s *GetConnectionTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
