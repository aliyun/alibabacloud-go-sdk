// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReplicationVaultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateReplicationVaultResponseBody
	GetCode() *string
	SetMessage(v string) *CreateReplicationVaultResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateReplicationVaultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateReplicationVaultResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateReplicationVaultResponseBody
	GetTaskId() *string
	SetVaultId(v string) *CreateReplicationVaultResponseBody
	GetVaultId() *string
}

type CreateReplicationVaultResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
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
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the job that is used to initialize the backup vault. You can call the DescribeTask operation to query the job status.
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

func (s CreateReplicationVaultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationVaultResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReplicationVaultResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateReplicationVaultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateReplicationVaultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateReplicationVaultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateReplicationVaultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateReplicationVaultResponseBody) GetVaultId() *string {
	return s.VaultId
}

func (s *CreateReplicationVaultResponseBody) SetCode(v string) *CreateReplicationVaultResponseBody {
	s.Code = &v
	return s
}

func (s *CreateReplicationVaultResponseBody) SetMessage(v string) *CreateReplicationVaultResponseBody {
	s.Message = &v
	return s
}

func (s *CreateReplicationVaultResponseBody) SetRequestId(v string) *CreateReplicationVaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReplicationVaultResponseBody) SetSuccess(v bool) *CreateReplicationVaultResponseBody {
	s.Success = &v
	return s
}

func (s *CreateReplicationVaultResponseBody) SetTaskId(v string) *CreateReplicationVaultResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateReplicationVaultResponseBody) SetVaultId(v string) *CreateReplicationVaultResponseBody {
	s.VaultId = &v
	return s
}

func (s *CreateReplicationVaultResponseBody) Validate() error {
	return dara.Validate(s)
}
