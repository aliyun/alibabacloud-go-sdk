// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoSyncTaskByRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRepoSyncTaskByRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateRepoSyncTaskByRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateRepoSyncTaskByRuleResponseBody
	GetRequestId() *string
	SetSyncTaskId(v string) *CreateRepoSyncTaskByRuleResponseBody
	GetSyncTaskId() *string
}

type CreateRepoSyncTaskByRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 17A4C658-AE8F-4A08-821F-EDCB5FC74EE8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// rst-biu4u4pm4it5****
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
}

func (s CreateRepoSyncTaskByRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoSyncTaskByRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncTaskByRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRepoSyncTaskByRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateRepoSyncTaskByRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRepoSyncTaskByRuleResponseBody) GetSyncTaskId() *string {
	return s.SyncTaskId
}

func (s *CreateRepoSyncTaskByRuleResponseBody) SetCode(v string) *CreateRepoSyncTaskByRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleResponseBody) SetIsSuccess(v bool) *CreateRepoSyncTaskByRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleResponseBody) SetRequestId(v string) *CreateRepoSyncTaskByRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleResponseBody) SetSyncTaskId(v string) *CreateRepoSyncTaskByRuleResponseBody {
	s.SyncTaskId = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
