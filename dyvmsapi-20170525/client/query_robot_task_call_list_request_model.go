// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotTaskCallListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallResult(v string) *QueryRobotTaskCallListRequest
	GetCallResult() *string
	SetCalled(v string) *QueryRobotTaskCallListRequest
	GetCalled() *string
	SetDialogCountFrom(v string) *QueryRobotTaskCallListRequest
	GetDialogCountFrom() *string
	SetDialogCountTo(v string) *QueryRobotTaskCallListRequest
	GetDialogCountTo() *string
	SetDurationFrom(v string) *QueryRobotTaskCallListRequest
	GetDurationFrom() *string
	SetDurationTo(v string) *QueryRobotTaskCallListRequest
	GetDurationTo() *string
	SetHangupDirection(v string) *QueryRobotTaskCallListRequest
	GetHangupDirection() *string
	SetOwnerId(v int64) *QueryRobotTaskCallListRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *QueryRobotTaskCallListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryRobotTaskCallListRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *QueryRobotTaskCallListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryRobotTaskCallListRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *QueryRobotTaskCallListRequest
	GetTaskId() *string
}

type QueryRobotTaskCallListRequest struct {
	// The call result. Valid values:
	//
	// 	- **200002**: The line is busy.
	//
	// 	- **200005**: The called party cannot be connected.
	//
	// 	- **200010**: The phone of the called party is powered off.
	//
	// 	- **200011**: The called party is out of service.
	//
	// 	- **200012**: The call is lost.
	//
	// example:
	//
	// 200002
	CallResult *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	// The called number.
	//
	// example:
	//
	// 1300****0000
	Called *string `json:"Called,omitempty" xml:"Called,omitempty"`
	// The minimum number of conversation rounds in the call.
	//
	// example:
	//
	// 0
	DialogCountFrom *string `json:"DialogCountFrom,omitempty" xml:"DialogCountFrom,omitempty"`
	// The maximum number of conversation rounds in the call.
	//
	// example:
	//
	// 5
	DialogCountTo *string `json:"DialogCountTo,omitempty" xml:"DialogCountTo,omitempty"`
	// The minimum call duration.
	//
	// example:
	//
	// 0
	DurationFrom *string `json:"DurationFrom,omitempty" xml:"DurationFrom,omitempty"`
	// The maximum call duration.
	//
	// example:
	//
	// 60
	DurationTo *string `json:"DurationTo,omitempty" xml:"DurationTo,omitempty"`
	// The party who hangs up. Valid values:
	//
	// 	- **0**: the called party.
	//
	// 	- **1**: the robot.
	//
	// example:
	//
	// 1
	HangupDirection *string `json:"HangupDirection,omitempty" xml:"HangupDirection,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](https://help.aliyun.com/document_detail/393531.html) operation to obtain the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1045001****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryRobotTaskCallListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotTaskCallListRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallListRequest) GetCallResult() *string {
	return s.CallResult
}

func (s *QueryRobotTaskCallListRequest) GetCalled() *string {
	return s.Called
}

func (s *QueryRobotTaskCallListRequest) GetDialogCountFrom() *string {
	return s.DialogCountFrom
}

func (s *QueryRobotTaskCallListRequest) GetDialogCountTo() *string {
	return s.DialogCountTo
}

func (s *QueryRobotTaskCallListRequest) GetDurationFrom() *string {
	return s.DurationFrom
}

func (s *QueryRobotTaskCallListRequest) GetDurationTo() *string {
	return s.DurationTo
}

func (s *QueryRobotTaskCallListRequest) GetHangupDirection() *string {
	return s.HangupDirection
}

func (s *QueryRobotTaskCallListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryRobotTaskCallListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryRobotTaskCallListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRobotTaskCallListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryRobotTaskCallListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryRobotTaskCallListRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryRobotTaskCallListRequest) SetCallResult(v string) *QueryRobotTaskCallListRequest {
	s.CallResult = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetCalled(v string) *QueryRobotTaskCallListRequest {
	s.Called = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetDialogCountFrom(v string) *QueryRobotTaskCallListRequest {
	s.DialogCountFrom = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetDialogCountTo(v string) *QueryRobotTaskCallListRequest {
	s.DialogCountTo = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetDurationFrom(v string) *QueryRobotTaskCallListRequest {
	s.DurationFrom = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetDurationTo(v string) *QueryRobotTaskCallListRequest {
	s.DurationTo = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetHangupDirection(v string) *QueryRobotTaskCallListRequest {
	s.HangupDirection = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetOwnerId(v int64) *QueryRobotTaskCallListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetPageNo(v int32) *QueryRobotTaskCallListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetPageSize(v int32) *QueryRobotTaskCallListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetResourceOwnerAccount(v string) *QueryRobotTaskCallListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetResourceOwnerId(v int64) *QueryRobotTaskCallListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetTaskId(v string) *QueryRobotTaskCallListRequest {
	s.TaskId = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) Validate() error {
	return dara.Validate(s)
}
