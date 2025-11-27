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
	// The ID of the backup file.
	//
	// example:
	//
	// b-g14d0m772f7b********
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6A236067-4727-4B42-92CF-734E417ED69A
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
