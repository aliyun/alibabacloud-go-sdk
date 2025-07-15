// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupFileIdList(v []*string) *DeleteBackupFileRequest
	GetBackupFileIdList() []*string
}

type DeleteBackupFileRequest struct {
	// This parameter is required.
	BackupFileIdList []*string `json:"BackupFileIdList,omitempty" xml:"BackupFileIdList,omitempty" type:"Repeated"`
}

func (s DeleteBackupFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupFileRequest) GetBackupFileIdList() []*string {
	return s.BackupFileIdList
}

func (s *DeleteBackupFileRequest) SetBackupFileIdList(v []*string) *DeleteBackupFileRequest {
	s.BackupFileIdList = v
	return s
}

func (s *DeleteBackupFileRequest) Validate() error {
	return dara.Validate(s)
}
