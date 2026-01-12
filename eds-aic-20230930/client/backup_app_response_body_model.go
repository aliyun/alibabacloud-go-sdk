// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackupAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *BackupAppResponseBody
	GetCount() *int64
	SetData(v []*BackupAppResponseBodyData) *BackupAppResponseBody
	GetData() []*BackupAppResponseBodyData
	SetRequestId(v string) *BackupAppResponseBody
	GetRequestId() *string
	SetTaskId(v string) *BackupAppResponseBody
	GetTaskId() *string
}

type BackupAppResponseBody struct {
	// example:
	//
	// 1
	Count *int64                       `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []*BackupAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// t-14xwibw7pyrjd****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BackupAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BackupAppResponseBody) GoString() string {
	return s.String()
}

func (s *BackupAppResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *BackupAppResponseBody) GetData() []*BackupAppResponseBodyData {
	return s.Data
}

func (s *BackupAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BackupAppResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *BackupAppResponseBody) SetCount(v int64) *BackupAppResponseBody {
	s.Count = &v
	return s
}

func (s *BackupAppResponseBody) SetData(v []*BackupAppResponseBodyData) *BackupAppResponseBody {
	s.Data = v
	return s
}

func (s *BackupAppResponseBody) SetRequestId(v string) *BackupAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *BackupAppResponseBody) SetTaskId(v string) *BackupAppResponseBody {
	s.TaskId = &v
	return s
}

func (s *BackupAppResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BackupAppResponseBodyData struct {
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// example:
	//
	// bf-b0qbg3pbpjkn7****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// example:
	//
	// MyBackup
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// example:
	//
	// t-4ks224ujixw****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BackupAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BackupAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *BackupAppResponseBodyData) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *BackupAppResponseBodyData) GetBackupFileId() *string {
	return s.BackupFileId
}

func (s *BackupAppResponseBodyData) GetBackupFileName() *string {
	return s.BackupFileName
}

func (s *BackupAppResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *BackupAppResponseBodyData) SetAndroidInstanceId(v string) *BackupAppResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *BackupAppResponseBodyData) SetBackupFileId(v string) *BackupAppResponseBodyData {
	s.BackupFileId = &v
	return s
}

func (s *BackupAppResponseBodyData) SetBackupFileName(v string) *BackupAppResponseBodyData {
	s.BackupFileName = &v
	return s
}

func (s *BackupAppResponseBodyData) SetTaskId(v string) *BackupAppResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *BackupAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
