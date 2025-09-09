// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutStartBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutStartBackupResponseBody
	GetRequestId() *string
	SetResult(v string) *PutStartBackupResponseBody
	GetResult() *string
	SetSuccess(v bool) *PutStartBackupResponseBody
	GetSuccess() *bool
}

type PutStartBackupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5D64DE5944A1E541E0******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the backup task was submitted.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutStartBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutStartBackupResponseBody) GoString() string {
	return s.String()
}

func (s *PutStartBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutStartBackupResponseBody) GetResult() *string {
	return s.Result
}

func (s *PutStartBackupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutStartBackupResponseBody) SetRequestId(v string) *PutStartBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutStartBackupResponseBody) SetResult(v string) *PutStartBackupResponseBody {
	s.Result = &v
	return s
}

func (s *PutStartBackupResponseBody) SetSuccess(v bool) *PutStartBackupResponseBody {
	s.Success = &v
	return s
}

func (s *PutStartBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
