// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRobotTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaller(v string) *CreateRobotTaskRequest
	GetCaller() *string
	SetCorpName(v string) *CreateRobotTaskRequest
	GetCorpName() *string
	SetDialogId(v int64) *CreateRobotTaskRequest
	GetDialogId() *int64
	SetIsSelfLine(v bool) *CreateRobotTaskRequest
	GetIsSelfLine() *bool
	SetNumberStatusIdent(v bool) *CreateRobotTaskRequest
	GetNumberStatusIdent() *bool
	SetOwnerId(v int64) *CreateRobotTaskRequest
	GetOwnerId() *int64
	SetRecallInterval(v int32) *CreateRobotTaskRequest
	GetRecallInterval() *int32
	SetRecallStateCodes(v string) *CreateRobotTaskRequest
	GetRecallStateCodes() *string
	SetRecallTimes(v int32) *CreateRobotTaskRequest
	GetRecallTimes() *int32
	SetResourceOwnerAccount(v string) *CreateRobotTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateRobotTaskRequest
	GetResourceOwnerId() *int64
	SetRetryType(v int32) *CreateRobotTaskRequest
	GetRetryType() *int32
	SetTaskName(v string) *CreateRobotTaskRequest
	GetTaskName() *string
}

type CreateRobotTaskRequest struct {
	// The calling number.
	//
	// You must use the phone numbers that you have purchased and separate multiple numbers with commas (,). You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and choose **Real Number Service*	- > **Real Number Management*	- to view the numbers you purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0571****5678
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The company name, which must be the same as the **enterprise name*	- on the qualification management page.
	//
	// example:
	//
	// Alibaba
	CorpName *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	// The ID of the robot or communication script that is used to initiate the call.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and choose **Intelligent Voice Robot*	- > **Communication Script Management*	- to view the communication script ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000023****
	DialogId *int64 `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	// Specifies whether to call the self-managed line. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// > If you set this parameter to **true**, calling numbers are not verified.
	//
	// example:
	//
	// false
	IsSelfLine *bool `json:"IsSelfLine,omitempty" xml:"IsSelfLine,omitempty"`
	// Specifies whether to enable number status identification. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// > If you set this parameter to **true**, the reason why a call is not answered is recorded.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	NumberStatusIdent *bool  `json:"NumberStatusIdent,omitempty" xml:"NumberStatusIdent,omitempty"`
	OwnerId           *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The redial interval. Unit: minutes. The value must be greater than 1.
	//
	// > The maximum redial interval is 30 minutes.
	//
	// example:
	//
	// 5
	RecallInterval *int32 `json:"RecallInterval,omitempty" xml:"RecallInterval,omitempty"`
	// The call state in which redial is required. Separate multiple call states with commas (,). Valid values:
	//
	// 	- **200010**: The phone of the called party is powered off.
	//
	// 	- **200011**: The number of the called party is out of service.
	//
	// 	- **200002**: The line is busy.
	//
	// 	- **200012**: The call is lost.
	//
	// 	- **200005**: The called party cannot be connected.
	//
	// 	- **200003**: The called party does not respond to the call.
	//
	// example:
	//
	// 200010,200011
	RecallStateCodes *string `json:"RecallStateCodes,omitempty" xml:"RecallStateCodes,omitempty"`
	// The number of redial times.
	//
	// example:
	//
	// 1
	RecallTimes          *int32  `json:"RecallTimes,omitempty" xml:"RecallTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable auto-redial. Valid values:
	//
	// 	- **1**: enables auto-redial.
	//
	// 	- **0**: disables auto-redial.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RetryType *int32 `json:"RetryType,omitempty" xml:"RetryType,omitempty"`
	// The task name. The task name can be up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test Template
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateRobotTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRobotTaskRequest) GetCaller() *string {
	return s.Caller
}

func (s *CreateRobotTaskRequest) GetCorpName() *string {
	return s.CorpName
}

func (s *CreateRobotTaskRequest) GetDialogId() *int64 {
	return s.DialogId
}

func (s *CreateRobotTaskRequest) GetIsSelfLine() *bool {
	return s.IsSelfLine
}

func (s *CreateRobotTaskRequest) GetNumberStatusIdent() *bool {
	return s.NumberStatusIdent
}

func (s *CreateRobotTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRobotTaskRequest) GetRecallInterval() *int32 {
	return s.RecallInterval
}

func (s *CreateRobotTaskRequest) GetRecallStateCodes() *string {
	return s.RecallStateCodes
}

func (s *CreateRobotTaskRequest) GetRecallTimes() *int32 {
	return s.RecallTimes
}

func (s *CreateRobotTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateRobotTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateRobotTaskRequest) GetRetryType() *int32 {
	return s.RetryType
}

func (s *CreateRobotTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateRobotTaskRequest) SetCaller(v string) *CreateRobotTaskRequest {
	s.Caller = &v
	return s
}

func (s *CreateRobotTaskRequest) SetCorpName(v string) *CreateRobotTaskRequest {
	s.CorpName = &v
	return s
}

func (s *CreateRobotTaskRequest) SetDialogId(v int64) *CreateRobotTaskRequest {
	s.DialogId = &v
	return s
}

func (s *CreateRobotTaskRequest) SetIsSelfLine(v bool) *CreateRobotTaskRequest {
	s.IsSelfLine = &v
	return s
}

func (s *CreateRobotTaskRequest) SetNumberStatusIdent(v bool) *CreateRobotTaskRequest {
	s.NumberStatusIdent = &v
	return s
}

func (s *CreateRobotTaskRequest) SetOwnerId(v int64) *CreateRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRobotTaskRequest) SetRecallInterval(v int32) *CreateRobotTaskRequest {
	s.RecallInterval = &v
	return s
}

func (s *CreateRobotTaskRequest) SetRecallStateCodes(v string) *CreateRobotTaskRequest {
	s.RecallStateCodes = &v
	return s
}

func (s *CreateRobotTaskRequest) SetRecallTimes(v int32) *CreateRobotTaskRequest {
	s.RecallTimes = &v
	return s
}

func (s *CreateRobotTaskRequest) SetResourceOwnerAccount(v string) *CreateRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRobotTaskRequest) SetResourceOwnerId(v int64) *CreateRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRobotTaskRequest) SetRetryType(v int32) *CreateRobotTaskRequest {
	s.RetryType = &v
	return s
}

func (s *CreateRobotTaskRequest) SetTaskName(v string) *CreateRobotTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateRobotTaskRequest) Validate() error {
	return dara.Validate(s)
}
