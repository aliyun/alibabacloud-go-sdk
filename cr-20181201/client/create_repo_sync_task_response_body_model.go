// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoSyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRepoSyncTaskResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateRepoSyncTaskResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateRepoSyncTaskResponseBody
	GetRequestId() *string
	SetSyncTaskId(v string) *CreateRepoSyncTaskResponseBody
	GetSyncTaskId() *string
}

type CreateRepoSyncTaskResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// 8F8A0BA6-7F06-4BAE-B147-10BD6A25****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rst-gbch330f0c****
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
}

func (s CreateRepoSyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoSyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRepoSyncTaskResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateRepoSyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRepoSyncTaskResponseBody) GetSyncTaskId() *string {
	return s.SyncTaskId
}

func (s *CreateRepoSyncTaskResponseBody) SetCode(v string) *CreateRepoSyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoSyncTaskResponseBody) SetIsSuccess(v bool) *CreateRepoSyncTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoSyncTaskResponseBody) SetRequestId(v string) *CreateRepoSyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoSyncTaskResponseBody) SetSyncTaskId(v string) *CreateRepoSyncTaskResponseBody {
	s.SyncTaskId = &v
	return s
}

func (s *CreateRepoSyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
