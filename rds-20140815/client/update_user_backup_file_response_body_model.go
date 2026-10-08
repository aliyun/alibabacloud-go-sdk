// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserBackupFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *UpdateUserBackupFileResponseBody
	GetBackupId() *string
	SetRequestId(v string) *UpdateUserBackupFileResponseBody
	GetRequestId() *string
}

type UpdateUserBackupFileResponseBody struct {
	// The user backup ID.
	//
	// example:
	//
	// b-lvn2365ev9f1****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 29EBB093-DBD8-5EEB-841D-E611B88CDE4B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserBackupFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserBackupFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserBackupFileResponseBody) GetBackupId() *string {
	return s.BackupId
}

func (s *UpdateUserBackupFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserBackupFileResponseBody) SetBackupId(v string) *UpdateUserBackupFileResponseBody {
	s.BackupId = &v
	return s
}

func (s *UpdateUserBackupFileResponseBody) SetRequestId(v string) *UpdateUserBackupFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserBackupFileResponseBody) Validate() error {
	return dara.Validate(s)
}
