// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVaultReplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVaultReplicationResponseBody
	GetCode() *string
	SetMessage(v string) *CreateVaultReplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVaultReplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVaultReplicationResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateVaultReplicationResponseBody
	GetTaskId() *string
}

type CreateVaultReplicationResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 280DD872-EE25-52E8-9CB4-491067173DD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// t-*********************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVaultReplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVaultReplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVaultReplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVaultReplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVaultReplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVaultReplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVaultReplicationResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVaultReplicationResponseBody) SetCode(v string) *CreateVaultReplicationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVaultReplicationResponseBody) SetMessage(v string) *CreateVaultReplicationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVaultReplicationResponseBody) SetRequestId(v string) *CreateVaultReplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVaultReplicationResponseBody) SetSuccess(v bool) *CreateVaultReplicationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVaultReplicationResponseBody) SetTaskId(v string) *CreateVaultReplicationResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateVaultReplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
