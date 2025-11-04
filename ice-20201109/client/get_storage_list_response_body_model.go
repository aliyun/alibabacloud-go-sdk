// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetStorageListResponseBody
	GetRequestId() *string
	SetStorageInfoList(v []*GetStorageListResponseBodyStorageInfoList) *GetStorageListResponseBody
	GetStorageInfoList() []*GetStorageListResponseBodyStorageInfoList
}

type GetStorageListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******73-8B78-5D86-A50C-49B96C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The storage configurations.
	StorageInfoList []*GetStorageListResponseBodyStorageInfoList `json:"StorageInfoList,omitempty" xml:"StorageInfoList,omitempty" type:"Repeated"`
}

func (s GetStorageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStorageListResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStorageListResponseBody) GetStorageInfoList() []*GetStorageListResponseBodyStorageInfoList {
	return s.StorageInfoList
}

func (s *GetStorageListResponseBody) SetRequestId(v string) *GetStorageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStorageListResponseBody) SetStorageInfoList(v []*GetStorageListResponseBodyStorageInfoList) *GetStorageListResponseBody {
	s.StorageInfoList = v
	return s
}

func (s *GetStorageListResponseBody) Validate() error {
	if s.StorageInfoList != nil {
		for _, item := range s.StorageInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStorageListResponseBodyStorageInfoList struct {
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the configuration was created.
	//
	// example:
	//
	// 2024-06-06T01:55:07Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether it is the default storage location.
	//
	// example:
	//
	// true
	DefaultStorage *bool `json:"DefaultStorage,omitempty" xml:"DefaultStorage,omitempty"`
	// Indicates whether temporary files created during editing processes are stored in this location.
	//
	// example:
	//
	// false
	EditingTempFileStorage *bool `json:"EditingTempFileStorage,omitempty" xml:"EditingTempFileStorage,omitempty"`
	// The time when the configuration was last modified.
	//
	// example:
	//
	// 2024-06-06T03:07:07Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The file path.
	//
	// example:
	//
	// your-path/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The OSS storage status.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The bucket.
	//
	// example:
	//
	// your-bucket
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The storage type.
	//
	// example:
	//
	// vod_oss_bucket
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s GetStorageListResponseBodyStorageInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetStorageListResponseBodyStorageInfoList) GoString() string {
	return s.String()
}

func (s *GetStorageListResponseBodyStorageInfoList) GetAppId() *string {
	return s.AppId
}

func (s *GetStorageListResponseBodyStorageInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetStorageListResponseBodyStorageInfoList) GetDefaultStorage() *bool {
	return s.DefaultStorage
}

func (s *GetStorageListResponseBodyStorageInfoList) GetEditingTempFileStorage() *bool {
	return s.EditingTempFileStorage
}

func (s *GetStorageListResponseBodyStorageInfoList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetStorageListResponseBodyStorageInfoList) GetPath() *string {
	return s.Path
}

func (s *GetStorageListResponseBodyStorageInfoList) GetStatus() *string {
	return s.Status
}

func (s *GetStorageListResponseBodyStorageInfoList) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetStorageListResponseBodyStorageInfoList) GetStorageType() *string {
	return s.StorageType
}

func (s *GetStorageListResponseBodyStorageInfoList) SetAppId(v string) *GetStorageListResponseBodyStorageInfoList {
	s.AppId = &v
	return s
}

func (s *GetStorageListResponseBodyStorageInfoList) SetCreationTime(v string) *GetStorageListResponseBodyStorageInfoList {
	s.CreationTime = &v
	return s
}

func (s *GetStorageListResponseBodyStorageInfoList) SetDefaultStorage(v bool) *GetStorageListResponseBodyStorageInfoList {
	s.DefaultStorage = &v
	return s
}

func (s *GetStorageListResponseBodyStorageInfoList) SetEditingTempFileStorage(v bool) *GetStorageListResponseBodyStorageInfoList {
	s.EditingTempFileStorage = &v
	return s
}

func (s *GetStorageListResponseBodyStorageInfoList) SetModifiedTime(v string) *GetStorageListResponseBodyStorageInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *GetStorageListResponseBodyStorageInfoList) SetPath(v string) *GetStorageListResponseBodyStorageInfoList {
	s.Path = &v
	return s
}

func (s *GetStorageListResponseBodyStorageInfoList) SetStatus(v string) *GetStorageListResponseBodyStorageInfoList {
	s.Status = &v
	return s
}

func (s *GetStorageListResponseBodyStorageInfoList) SetStorageLocation(v string) *GetStorageListResponseBodyStorageInfoList {
	s.StorageLocation = &v
	return s
}

func (s *GetStorageListResponseBodyStorageInfoList) SetStorageType(v string) *GetStorageListResponseBodyStorageInfoList {
	s.StorageType = &v
	return s
}

func (s *GetStorageListResponseBodyStorageInfoList) Validate() error {
	return dara.Validate(s)
}
