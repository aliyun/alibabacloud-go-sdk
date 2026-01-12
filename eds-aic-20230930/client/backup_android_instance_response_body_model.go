// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackupAndroidInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *BackupAndroidInstanceResponseBody
	GetCount() *int64
	SetData(v []*BackupAndroidInstanceResponseBodyData) *BackupAndroidInstanceResponseBody
	GetData() []*BackupAndroidInstanceResponseBodyData
	SetRequestId(v string) *BackupAndroidInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *BackupAndroidInstanceResponseBody
	GetTaskId() *string
}

type BackupAndroidInstanceResponseBody struct {
	// example:
	//
	// 1
	Count *int64                                   `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []*BackupAndroidInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 440D7342-5FC2-5E7C-B2DB-D0B4EAC2BDF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// t-14xwibw7pyrjd****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BackupAndroidInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BackupAndroidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *BackupAndroidInstanceResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *BackupAndroidInstanceResponseBody) GetData() []*BackupAndroidInstanceResponseBodyData {
	return s.Data
}

func (s *BackupAndroidInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BackupAndroidInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *BackupAndroidInstanceResponseBody) SetCount(v int64) *BackupAndroidInstanceResponseBody {
	s.Count = &v
	return s
}

func (s *BackupAndroidInstanceResponseBody) SetData(v []*BackupAndroidInstanceResponseBodyData) *BackupAndroidInstanceResponseBody {
	s.Data = v
	return s
}

func (s *BackupAndroidInstanceResponseBody) SetRequestId(v string) *BackupAndroidInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BackupAndroidInstanceResponseBody) SetTaskId(v string) *BackupAndroidInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *BackupAndroidInstanceResponseBody) Validate() error {
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

type BackupAndroidInstanceResponseBodyData struct {
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
	// a-58ftsoo90p0qa****.ab
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// example:
	//
	// t-bp67acfmxazb4p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BackupAndroidInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BackupAndroidInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *BackupAndroidInstanceResponseBodyData) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *BackupAndroidInstanceResponseBodyData) GetBackupFileId() *string {
	return s.BackupFileId
}

func (s *BackupAndroidInstanceResponseBodyData) GetBackupFileName() *string {
	return s.BackupFileName
}

func (s *BackupAndroidInstanceResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *BackupAndroidInstanceResponseBodyData) SetAndroidInstanceId(v string) *BackupAndroidInstanceResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *BackupAndroidInstanceResponseBodyData) SetBackupFileId(v string) *BackupAndroidInstanceResponseBodyData {
	s.BackupFileId = &v
	return s
}

func (s *BackupAndroidInstanceResponseBodyData) SetBackupFileName(v string) *BackupAndroidInstanceResponseBodyData {
	s.BackupFileName = &v
	return s
}

func (s *BackupAndroidInstanceResponseBodyData) SetTaskId(v string) *BackupAndroidInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *BackupAndroidInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
