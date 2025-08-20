// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRepoSyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelRepoSyncTaskResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CancelRepoSyncTaskResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CancelRepoSyncTaskResponseBody
	GetRequestId() *string
}

type CancelRepoSyncTaskResponseBody struct {
	// The response code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the CancelRepoSyncTask request is successful. Valid values:
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EB9C5722-51E2-4497-A573-575B0CA5CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelRepoSyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelRepoSyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRepoSyncTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelRepoSyncTaskResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CancelRepoSyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelRepoSyncTaskResponseBody) SetCode(v string) *CancelRepoSyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelRepoSyncTaskResponseBody) SetIsSuccess(v bool) *CancelRepoSyncTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CancelRepoSyncTaskResponseBody) SetRequestId(v string) *CancelRepoSyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelRepoSyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
