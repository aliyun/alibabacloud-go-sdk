// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserBackupFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DeleteUserBackupFileResponseBody
	GetBackupId() *string
	SetRequestId(v string) *DeleteUserBackupFileResponseBody
	GetRequestId() *string
}

type DeleteUserBackupFileResponseBody struct {
	// The ID of the deleted full backup file.
	//
	// example:
	//
	// b-w1haya7e4i25********
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F28AE40B-203B-4CFE-B81F-FD981CD97B17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserBackupFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserBackupFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserBackupFileResponseBody) GetBackupId() *string {
	return s.BackupId
}

func (s *DeleteUserBackupFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserBackupFileResponseBody) SetBackupId(v string) *DeleteUserBackupFileResponseBody {
	s.BackupId = &v
	return s
}

func (s *DeleteUserBackupFileResponseBody) SetRequestId(v string) *DeleteUserBackupFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserBackupFileResponseBody) Validate() error {
	return dara.Validate(s)
}
