// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVaultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVaultResponseBody
	GetCode() *string
	SetMessage(v string) *CreateVaultResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVaultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVaultResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateVaultResponseBody
	GetTaskId() *string
	SetVaultId(v string) *CreateVaultResponseBody
	GetVaultId() *string
}

type CreateVaultResponseBody struct {
	// The return code. A value of 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, \\`successful\\` is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the backup vault initialization task. Use the DescribeTask operation to query the task status.
	//
	// example:
	//
	// t-*********************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-*********************
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateVaultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVaultResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVaultResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVaultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVaultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVaultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVaultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVaultResponseBody) GetVaultId() *string {
	return s.VaultId
}

func (s *CreateVaultResponseBody) SetCode(v string) *CreateVaultResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVaultResponseBody) SetMessage(v string) *CreateVaultResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVaultResponseBody) SetRequestId(v string) *CreateVaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVaultResponseBody) SetSuccess(v bool) *CreateVaultResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVaultResponseBody) SetTaskId(v string) *CreateVaultResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateVaultResponseBody) SetVaultId(v string) *CreateVaultResponseBody {
	s.VaultId = &v
	return s
}

func (s *CreateVaultResponseBody) Validate() error {
	return dara.Validate(s)
}
