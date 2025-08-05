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
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the initialization task used to initialize the backup vault. You can call the DescribeTask operation to query the status of an initialization task.
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
