// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupExpireTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupExpireTime(v string) *ModifyBackupExpireTimeResponseBody
	GetBackupExpireTime() *string
	SetBackupId(v string) *ModifyBackupExpireTimeResponseBody
	GetBackupId() *string
	SetRequestId(v string) *ModifyBackupExpireTimeResponseBody
	GetRequestId() *string
}

type ModifyBackupExpireTimeResponseBody struct {
	// The time-to-live (TTL) of the backup. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format and is in UTC.
	//
	// > - *9999-01-01*T*00:00:00*&#x5A;*&#x20;indicates that the backup is retained permanently.*
	//
	// >
	//
	// >   **
	//
	// >
	//
	// > **
	//
	// example:
	//
	// 2025-03-29T03:47:12Z
	BackupExpireTime *string `json:"BackupExpireTime,omitempty" xml:"BackupExpireTime,omitempty"`
	// The backup ID.
	//
	// example:
	//
	// 260032xxxx
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A0181AC4-XXXX-XXXX-87CA-100C70B86729
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupExpireTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupExpireTimeResponseBody) GetBackupExpireTime() *string {
	return s.BackupExpireTime
}

func (s *ModifyBackupExpireTimeResponseBody) GetBackupId() *string {
	return s.BackupId
}

func (s *ModifyBackupExpireTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupExpireTimeResponseBody) SetBackupExpireTime(v string) *ModifyBackupExpireTimeResponseBody {
	s.BackupExpireTime = &v
	return s
}

func (s *ModifyBackupExpireTimeResponseBody) SetBackupId(v string) *ModifyBackupExpireTimeResponseBody {
	s.BackupId = &v
	return s
}

func (s *ModifyBackupExpireTimeResponseBody) SetRequestId(v string) *ModifyBackupExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupExpireTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
