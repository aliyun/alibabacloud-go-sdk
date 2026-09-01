// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceDirectoryAccountId(v int64) *GetCheckProcessRequest
	GetResourceDirectoryAccountId() *int64
	SetTaskId(v string) *GetCheckProcessRequest
	GetTaskId() *string
}

type GetCheckProcessRequest struct {
	// The ID of the member accounts in the resource folder (Alibaba Cloud account).
	//
	// > You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The ID of the cloud service configuration check task that you want to query.
	//
	// > You can call the [SubmitCheck](~~SubmitCheck~~) operation to obtain this parameter.
	//
	// example:
	//
	// 5347c7b6-c85c-4070-846a-3029e08e****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetCheckProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckProcessRequest) GoString() string {
	return s.String()
}

func (s *GetCheckProcessRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *GetCheckProcessRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCheckProcessRequest) SetResourceDirectoryAccountId(v int64) *GetCheckProcessRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *GetCheckProcessRequest) SetTaskId(v string) *GetCheckProcessRequest {
	s.TaskId = &v
	return s
}

func (s *GetCheckProcessRequest) Validate() error {
	return dara.Validate(s)
}
