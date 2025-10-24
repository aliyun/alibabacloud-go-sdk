// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackupFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *BackupFileResponseBody
	GetCount() *int64
	SetData(v []*BackupFileResponseBodyData) *BackupFileResponseBody
	GetData() []*BackupFileResponseBodyData
	SetRequestId(v string) *BackupFileResponseBody
	GetRequestId() *string
	SetTaskId(v string) *BackupFileResponseBody
	GetTaskId() *string
}

type BackupFileResponseBody struct {
	// The number of instances that are backed up.
	//
	// example:
	//
	// 100
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The object that is returned.
	//
	// example:
	//
	// 6C8439B9-7DBF-57F4-92AE-55A9B9D3****
	Data []*BackupFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6C8439B9-7DBF-57F4-92AE-55A9B9D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the batch task.
	//
	// example:
	//
	// t-22ex666a5mco5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BackupFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BackupFileResponseBody) GoString() string {
	return s.String()
}

func (s *BackupFileResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *BackupFileResponseBody) GetData() []*BackupFileResponseBodyData {
	return s.Data
}

func (s *BackupFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BackupFileResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *BackupFileResponseBody) SetCount(v int64) *BackupFileResponseBody {
	s.Count = &v
	return s
}

func (s *BackupFileResponseBody) SetData(v []*BackupFileResponseBodyData) *BackupFileResponseBody {
	s.Data = v
	return s
}

func (s *BackupFileResponseBody) SetRequestId(v string) *BackupFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *BackupFileResponseBody) SetTaskId(v string) *BackupFileResponseBody {
	s.TaskId = &v
	return s
}

func (s *BackupFileResponseBody) Validate() error {
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

type BackupFileResponseBodyData struct {
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The ID of the backup file.
	//
	// example:
	//
	// bf-b0qbg3pbpjkn7****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// The name of the backup file.
	//
	// example:
	//
	// a-58ftsoo90p0qa****.ab
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-22ex666a5mco5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BackupFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BackupFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *BackupFileResponseBodyData) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *BackupFileResponseBodyData) GetBackupFileId() *string {
	return s.BackupFileId
}

func (s *BackupFileResponseBodyData) GetBackupFileName() *string {
	return s.BackupFileName
}

func (s *BackupFileResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *BackupFileResponseBodyData) SetAndroidInstanceId(v string) *BackupFileResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *BackupFileResponseBodyData) SetBackupFileId(v string) *BackupFileResponseBodyData {
	s.BackupFileId = &v
	return s
}

func (s *BackupFileResponseBodyData) SetBackupFileName(v string) *BackupFileResponseBodyData {
	s.BackupFileName = &v
	return s
}

func (s *BackupFileResponseBodyData) SetTaskId(v string) *BackupFileResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *BackupFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
