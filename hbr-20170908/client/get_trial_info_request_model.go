// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrialInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *GetTrialInfoRequest
	GetBucket() *string
	SetCreateTime(v int64) *GetTrialInfoRequest
	GetCreateTime() *int64
	SetFileSystemId(v string) *GetTrialInfoRequest
	GetFileSystemId() *string
	SetSourceType(v string) *GetTrialInfoRequest
	GetSourceType() *string
}

type GetTrialInfoRequest struct {
	// This parameter is required only when **SourceType*	- is set to **OSS**. The name of the OSS bucket.
	//
	// example:
	//
	// hbr-backup-oss
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **NAS**. The time when the file system was created. UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1607436917
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **NAS**. The ID of the file system.
	//
	// example:
	//
	// 005494
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The data source type. Only free trial information of OSS backup and NAS backup can be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s GetTrialInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTrialInfoRequest) GoString() string {
	return s.String()
}

func (s *GetTrialInfoRequest) GetBucket() *string {
	return s.Bucket
}

func (s *GetTrialInfoRequest) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetTrialInfoRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetTrialInfoRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *GetTrialInfoRequest) SetBucket(v string) *GetTrialInfoRequest {
	s.Bucket = &v
	return s
}

func (s *GetTrialInfoRequest) SetCreateTime(v int64) *GetTrialInfoRequest {
	s.CreateTime = &v
	return s
}

func (s *GetTrialInfoRequest) SetFileSystemId(v string) *GetTrialInfoRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetTrialInfoRequest) SetSourceType(v string) *GetTrialInfoRequest {
	s.SourceType = &v
	return s
}

func (s *GetTrialInfoRequest) Validate() error {
	return dara.Validate(s)
}
