// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineRuntimeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetHotlineRuntimeInfoRequest
	GetAccountName() *string
	SetInstanceId(v string) *GetHotlineRuntimeInfoRequest
	GetInstanceId() *string
}

type GetHotlineRuntimeInfoRequest struct {
	// The agent account name. It is unique within the instance (logon name).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123@****.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The instance ID.
	//
	// You can log on to the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview) and view the instance ID in **Instance Management**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetHotlineRuntimeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineRuntimeInfoRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineRuntimeInfoRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetHotlineRuntimeInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHotlineRuntimeInfoRequest) SetAccountName(v string) *GetHotlineRuntimeInfoRequest {
	s.AccountName = &v
	return s
}

func (s *GetHotlineRuntimeInfoRequest) SetInstanceId(v string) *GetHotlineRuntimeInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineRuntimeInfoRequest) Validate() error {
	return dara.Validate(s)
}
