// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportUserBackupFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *ImportUserBackupFileResponseBody
	GetBackupId() *string
	SetRequestId(v string) *ImportUserBackupFileResponseBody
	GetRequestId() *string
	SetStatus(v bool) *ImportUserBackupFileResponseBody
	GetStatus() *bool
}

type ImportUserBackupFileResponseBody struct {
	// The user backup ID.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// b-n8tpg24c6i0v****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A140DD14-DCC9-4548-9C72-52A49A58A310
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the user backup is imported. The value **true*	- is returned if the import is successful. Otherwise, an error message is returned.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ImportUserBackupFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportUserBackupFileResponseBody) GoString() string {
	return s.String()
}

func (s *ImportUserBackupFileResponseBody) GetBackupId() *string {
	return s.BackupId
}

func (s *ImportUserBackupFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportUserBackupFileResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *ImportUserBackupFileResponseBody) SetBackupId(v string) *ImportUserBackupFileResponseBody {
	s.BackupId = &v
	return s
}

func (s *ImportUserBackupFileResponseBody) SetRequestId(v string) *ImportUserBackupFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportUserBackupFileResponseBody) SetStatus(v bool) *ImportUserBackupFileResponseBody {
	s.Status = &v
	return s
}

func (s *ImportUserBackupFileResponseBody) Validate() error {
	return dara.Validate(s)
}
