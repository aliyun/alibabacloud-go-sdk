// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBackupsSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *RemoveBackupsSetRequest
	GetBackupId() *string
	SetDrdsInstanceId(v string) *RemoveBackupsSetRequest
	GetDrdsInstanceId() *string
}

type RemoveBackupsSetRequest struct {
	// The ID of the backup set. You can call the [DescribeBackupSets](https://help.aliyun.com/document_detail/139331.html) interface to query the ID of a backup set.
	//
	// This parameter is required.
	//
	// example:
	//
	// ba30d5c4-a6dc-11ea-bd40-************
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveBackupsSetRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackupsSetRequest) GoString() string {
	return s.String()
}

func (s *RemoveBackupsSetRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *RemoveBackupsSetRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *RemoveBackupsSetRequest) SetBackupId(v string) *RemoveBackupsSetRequest {
	s.BackupId = &v
	return s
}

func (s *RemoveBackupsSetRequest) SetDrdsInstanceId(v string) *RemoveBackupsSetRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RemoveBackupsSetRequest) Validate() error {
	return dara.Validate(s)
}
