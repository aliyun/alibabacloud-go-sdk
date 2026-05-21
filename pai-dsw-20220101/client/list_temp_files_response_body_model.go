// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTempFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTempFilesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListTempFilesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTempFilesResponseBody
	GetMessage() *string
	SetQuota(v *ListTempFilesResponseBodyQuota) *ListTempFilesResponseBody
	GetQuota() *ListTempFilesResponseBodyQuota
	SetRequestId(v string) *ListTempFilesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTempFilesResponseBody
	GetSuccess() *bool
	SetTempFiles(v []*ListTempFilesResponseBodyTempFiles) *ListTempFilesResponseBody
	GetTempFiles() []*ListTempFilesResponseBodyTempFiles
	SetTotalCount(v int64) *ListTempFilesResponseBody
	GetTotalCount() *int64
}

type ListTempFilesResponseBody struct {
	// example:
	//
	// "200"
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// "XXX"
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Quota   *ListTempFilesResponseBodyQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TempFiles []*ListTempFilesResponseBodyTempFiles `json:"TempFiles,omitempty" xml:"TempFiles,omitempty" type:"Repeated"`
	// example:
	//
	// 35
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTempFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTempFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTempFilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTempFilesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTempFilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTempFilesResponseBody) GetQuota() *ListTempFilesResponseBodyQuota {
	return s.Quota
}

func (s *ListTempFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTempFilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTempFilesResponseBody) GetTempFiles() []*ListTempFilesResponseBodyTempFiles {
	return s.TempFiles
}

func (s *ListTempFilesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTempFilesResponseBody) SetCode(v string) *ListTempFilesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTempFilesResponseBody) SetHttpStatusCode(v int32) *ListTempFilesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTempFilesResponseBody) SetMessage(v string) *ListTempFilesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTempFilesResponseBody) SetQuota(v *ListTempFilesResponseBodyQuota) *ListTempFilesResponseBody {
	s.Quota = v
	return s
}

func (s *ListTempFilesResponseBody) SetRequestId(v string) *ListTempFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTempFilesResponseBody) SetSuccess(v bool) *ListTempFilesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTempFilesResponseBody) SetTempFiles(v []*ListTempFilesResponseBodyTempFiles) *ListTempFilesResponseBody {
	s.TempFiles = v
	return s
}

func (s *ListTempFilesResponseBody) SetTotalCount(v int64) *ListTempFilesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTempFilesResponseBody) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	if s.TempFiles != nil {
		for _, item := range s.TempFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTempFilesResponseBodyQuota struct {
	// example:
	//
	// 1000
	TotalCapacity *int64 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	// example:
	//
	// 1000
	TotalFileNum *int64 `json:"TotalFileNum,omitempty" xml:"TotalFileNum,omitempty"`
	// example:
	//
	// 1000
	UsedCapacity *int64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
	// example:
	//
	// 100
	UsedFileNum *int64 `json:"UsedFileNum,omitempty" xml:"UsedFileNum,omitempty"`
}

func (s ListTempFilesResponseBodyQuota) String() string {
	return dara.Prettify(s)
}

func (s ListTempFilesResponseBodyQuota) GoString() string {
	return s.String()
}

func (s *ListTempFilesResponseBodyQuota) GetTotalCapacity() *int64 {
	return s.TotalCapacity
}

func (s *ListTempFilesResponseBodyQuota) GetTotalFileNum() *int64 {
	return s.TotalFileNum
}

func (s *ListTempFilesResponseBodyQuota) GetUsedCapacity() *int64 {
	return s.UsedCapacity
}

func (s *ListTempFilesResponseBodyQuota) GetUsedFileNum() *int64 {
	return s.UsedFileNum
}

func (s *ListTempFilesResponseBodyQuota) SetTotalCapacity(v int64) *ListTempFilesResponseBodyQuota {
	s.TotalCapacity = &v
	return s
}

func (s *ListTempFilesResponseBodyQuota) SetTotalFileNum(v int64) *ListTempFilesResponseBodyQuota {
	s.TotalFileNum = &v
	return s
}

func (s *ListTempFilesResponseBodyQuota) SetUsedCapacity(v int64) *ListTempFilesResponseBodyQuota {
	s.UsedCapacity = &v
	return s
}

func (s *ListTempFilesResponseBodyQuota) SetUsedFileNum(v int64) *ListTempFilesResponseBodyQuota {
	s.UsedFileNum = &v
	return s
}

func (s *ListTempFilesResponseBodyQuota) Validate() error {
	return dara.Validate(s)
}

type ListTempFilesResponseBodyTempFiles struct {
	// example:
	//
	// 10Mi
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// example:
	//
	// 1000
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// http://examplebucket.yourEndpoint/exampledir/exampleobject.txt?Expires=1703123270&OSSAccessKeyId&Signature=Fn7xSv2kRQraU6UqRZ3+DMzQK14=
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtExpiredTime *string `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// filename
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Owner Id
	//
	// example:
	//
	// 1612285282502324
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// /1079135428626537/209170658818096848/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// task-05cexxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// type
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// http://examplebucket.yourEndpoint/exampledir/exampleobject.txt?Expires=1703123154&OSSAccessKeyId&Signature=5ekVo7r+CeeU5oYmCpnmrzrx2IM=
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
	// example:
	//
	// 1612285282502324
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// tempfile-05cexxxxxxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListTempFilesResponseBodyTempFiles) String() string {
	return dara.Prettify(s)
}

func (s ListTempFilesResponseBodyTempFiles) GoString() string {
	return s.String()
}

func (s *ListTempFilesResponseBodyTempFiles) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListTempFilesResponseBodyTempFiles) GetCapacity() *int64 {
	return s.Capacity
}

func (s *ListTempFilesResponseBodyTempFiles) GetDescription() *string {
	return s.Description
}

func (s *ListTempFilesResponseBodyTempFiles) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListTempFilesResponseBodyTempFiles) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListTempFilesResponseBodyTempFiles) GetGmtExpiredTime() *string {
	return s.GmtExpiredTime
}

func (s *ListTempFilesResponseBodyTempFiles) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListTempFilesResponseBodyTempFiles) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTempFilesResponseBodyTempFiles) GetName() *string {
	return s.Name
}

func (s *ListTempFilesResponseBodyTempFiles) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListTempFilesResponseBodyTempFiles) GetPrefix() *string {
	return s.Prefix
}

func (s *ListTempFilesResponseBodyTempFiles) GetStatus() *string {
	return s.Status
}

func (s *ListTempFilesResponseBodyTempFiles) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTempFilesResponseBodyTempFiles) GetType() *string {
	return s.Type
}

func (s *ListTempFilesResponseBodyTempFiles) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *ListTempFilesResponseBodyTempFiles) GetUserId() *string {
	return s.UserId
}

func (s *ListTempFilesResponseBodyTempFiles) GetUuid() *string {
	return s.Uuid
}

func (s *ListTempFilesResponseBodyTempFiles) SetBandwidth(v int32) *ListTempFilesResponseBodyTempFiles {
	s.Bandwidth = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetCapacity(v int64) *ListTempFilesResponseBodyTempFiles {
	s.Capacity = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetDescription(v string) *ListTempFilesResponseBodyTempFiles {
	s.Description = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetDownloadUrl(v string) *ListTempFilesResponseBodyTempFiles {
	s.DownloadUrl = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetGmtCreateTime(v string) *ListTempFilesResponseBodyTempFiles {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetGmtExpiredTime(v string) *ListTempFilesResponseBodyTempFiles {
	s.GmtExpiredTime = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetGmtModifiedTime(v string) *ListTempFilesResponseBodyTempFiles {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetInstanceId(v string) *ListTempFilesResponseBodyTempFiles {
	s.InstanceId = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetName(v string) *ListTempFilesResponseBodyTempFiles {
	s.Name = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetOwnerId(v string) *ListTempFilesResponseBodyTempFiles {
	s.OwnerId = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetPrefix(v string) *ListTempFilesResponseBodyTempFiles {
	s.Prefix = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetStatus(v string) *ListTempFilesResponseBodyTempFiles {
	s.Status = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetTaskId(v string) *ListTempFilesResponseBodyTempFiles {
	s.TaskId = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetType(v string) *ListTempFilesResponseBodyTempFiles {
	s.Type = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetUploadUrl(v string) *ListTempFilesResponseBodyTempFiles {
	s.UploadUrl = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetUserId(v string) *ListTempFilesResponseBodyTempFiles {
	s.UserId = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) SetUuid(v string) *ListTempFilesResponseBodyTempFiles {
	s.Uuid = &v
	return s
}

func (s *ListTempFilesResponseBodyTempFiles) Validate() error {
	return dara.Validate(s)
}
