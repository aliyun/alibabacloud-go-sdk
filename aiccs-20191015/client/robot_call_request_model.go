// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRobotCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumber(v string) *RobotCallRequest
	GetCalledNumber() *string
	SetCalledShowNumber(v string) *RobotCallRequest
	GetCalledShowNumber() *string
	SetEarlyMediaAsr(v bool) *RobotCallRequest
	GetEarlyMediaAsr() *bool
	SetOutId(v string) *RobotCallRequest
	GetOutId() *string
	SetOwnerId(v int64) *RobotCallRequest
	GetOwnerId() *int64
	SetParams(v string) *RobotCallRequest
	GetParams() *string
	SetRecordFlag(v bool) *RobotCallRequest
	GetRecordFlag() *bool
	SetResourceOwnerAccount(v string) *RobotCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RobotCallRequest
	GetResourceOwnerId() *int64
	SetRobotId(v int64) *RobotCallRequest
	GetRobotId() *int64
}

type RobotCallRequest struct {
	// The called number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 131****2204
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The caller number for the outbound call. You can view this number on the [Real Number Management](https://dyvmsnext.console.aliyun.com/number/list/normal) page of the voice messaging service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 132****0849
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// Specifies whether to enable the number status identification feature. Valid values:
	//
	// - **false*	- (default): Disabled.
	//
	// - **true**: Enabled.
	//
	// >After this feature is enabled, the reason why a call is not answered is recorded.
	//
	// example:
	//
	// false
	EarlyMediaAsr *bool `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	// The ID reserved for the caller. This ID is returned to the caller through a receipt message. The value is 1 to 15 bytes in length.
	//
	// example:
	//
	// 222356****
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parameter list of the robot in JSON format. The total length cannot exceed 512 bytes. You can view the parameters on the [Script Management](https://aiccs.console.aliyun.com/patter/list) > **Details*	- > **Input/Output Parameters*	- page or by calling the [ListRobotParams](https://help.aliyun.com/document_detail/2717999.html) operation.
	//
	// example:
	//
	// {
	//
	//   "name": "Zhang San",
	//
	//   "age": "25"
	//
	// }
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// Specifies whether to record the call. Valid values:
	//
	// - **false*	- (default): Do not record.
	//
	// - **true**: Record.
	//
	// example:
	//
	// true
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The robot ID, which is the script ID. You can view approved script IDs on the [Script Management](https://aiccs.console.aliyun.com/patter/list) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000035******
	RobotId *int64 `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
}

func (s RobotCallRequest) String() string {
	return dara.Prettify(s)
}

func (s RobotCallRequest) GoString() string {
	return s.String()
}

func (s *RobotCallRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *RobotCallRequest) GetCalledShowNumber() *string {
	return s.CalledShowNumber
}

func (s *RobotCallRequest) GetEarlyMediaAsr() *bool {
	return s.EarlyMediaAsr
}

func (s *RobotCallRequest) GetOutId() *string {
	return s.OutId
}

func (s *RobotCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RobotCallRequest) GetParams() *string {
	return s.Params
}

func (s *RobotCallRequest) GetRecordFlag() *bool {
	return s.RecordFlag
}

func (s *RobotCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RobotCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RobotCallRequest) GetRobotId() *int64 {
	return s.RobotId
}

func (s *RobotCallRequest) SetCalledNumber(v string) *RobotCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *RobotCallRequest) SetCalledShowNumber(v string) *RobotCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *RobotCallRequest) SetEarlyMediaAsr(v bool) *RobotCallRequest {
	s.EarlyMediaAsr = &v
	return s
}

func (s *RobotCallRequest) SetOutId(v string) *RobotCallRequest {
	s.OutId = &v
	return s
}

func (s *RobotCallRequest) SetOwnerId(v int64) *RobotCallRequest {
	s.OwnerId = &v
	return s
}

func (s *RobotCallRequest) SetParams(v string) *RobotCallRequest {
	s.Params = &v
	return s
}

func (s *RobotCallRequest) SetRecordFlag(v bool) *RobotCallRequest {
	s.RecordFlag = &v
	return s
}

func (s *RobotCallRequest) SetResourceOwnerAccount(v string) *RobotCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RobotCallRequest) SetResourceOwnerId(v int64) *RobotCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RobotCallRequest) SetRobotId(v int64) *RobotCallRequest {
	s.RobotId = &v
	return s
}

func (s *RobotCallRequest) Validate() error {
	return dara.Validate(s)
}
