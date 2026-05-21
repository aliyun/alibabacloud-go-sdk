// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTempFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *GetTempFileResponseBody
	GetBandwidth() *int32
	SetCapacity(v int64) *GetTempFileResponseBody
	GetCapacity() *int64
	SetCode(v string) *GetTempFileResponseBody
	GetCode() *string
	SetDescription(v string) *GetTempFileResponseBody
	GetDescription() *string
	SetDownloadUrl(v string) *GetTempFileResponseBody
	GetDownloadUrl() *string
	SetGmtCreateTime(v string) *GetTempFileResponseBody
	GetGmtCreateTime() *string
	SetGmtExpiredTime(v string) *GetTempFileResponseBody
	GetGmtExpiredTime() *string
	SetGmtModifiedTime(v string) *GetTempFileResponseBody
	GetGmtModifiedTime() *string
	SetHttpStatusCode(v int32) *GetTempFileResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *GetTempFileResponseBody
	GetInstanceId() *string
	SetMessage(v string) *GetTempFileResponseBody
	GetMessage() *string
	SetName(v string) *GetTempFileResponseBody
	GetName() *string
	SetOwnerId(v string) *GetTempFileResponseBody
	GetOwnerId() *string
	SetPrefix(v string) *GetTempFileResponseBody
	GetPrefix() *string
	SetRequestId(v string) *GetTempFileResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetTempFileResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetTempFileResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *GetTempFileResponseBody
	GetTaskId() *string
	SetType(v string) *GetTempFileResponseBody
	GetType() *string
	SetUploadUrl(v string) *GetTempFileResponseBody
	GetUploadUrl() *string
	SetUserId(v string) *GetTempFileResponseBody
	GetUserId() *string
	SetUuid(v string) *GetTempFileResponseBody
	GetUuid() *string
}

type GetTempFileResponseBody struct {
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
	// "200"
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s GetTempFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTempFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetTempFileResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *GetTempFileResponseBody) GetCapacity() *int64 {
	return s.Capacity
}

func (s *GetTempFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTempFileResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetTempFileResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetTempFileResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetTempFileResponseBody) GetGmtExpiredTime() *string {
	return s.GmtExpiredTime
}

func (s *GetTempFileResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetTempFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTempFileResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTempFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTempFileResponseBody) GetName() *string {
	return s.Name
}

func (s *GetTempFileResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetTempFileResponseBody) GetPrefix() *string {
	return s.Prefix
}

func (s *GetTempFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTempFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetTempFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTempFileResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTempFileResponseBody) GetType() *string {
	return s.Type
}

func (s *GetTempFileResponseBody) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *GetTempFileResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetTempFileResponseBody) GetUuid() *string {
	return s.Uuid
}

func (s *GetTempFileResponseBody) SetBandwidth(v int32) *GetTempFileResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *GetTempFileResponseBody) SetCapacity(v int64) *GetTempFileResponseBody {
	s.Capacity = &v
	return s
}

func (s *GetTempFileResponseBody) SetCode(v string) *GetTempFileResponseBody {
	s.Code = &v
	return s
}

func (s *GetTempFileResponseBody) SetDescription(v string) *GetTempFileResponseBody {
	s.Description = &v
	return s
}

func (s *GetTempFileResponseBody) SetDownloadUrl(v string) *GetTempFileResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *GetTempFileResponseBody) SetGmtCreateTime(v string) *GetTempFileResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTempFileResponseBody) SetGmtExpiredTime(v string) *GetTempFileResponseBody {
	s.GmtExpiredTime = &v
	return s
}

func (s *GetTempFileResponseBody) SetGmtModifiedTime(v string) *GetTempFileResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetTempFileResponseBody) SetHttpStatusCode(v int32) *GetTempFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTempFileResponseBody) SetInstanceId(v string) *GetTempFileResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetTempFileResponseBody) SetMessage(v string) *GetTempFileResponseBody {
	s.Message = &v
	return s
}

func (s *GetTempFileResponseBody) SetName(v string) *GetTempFileResponseBody {
	s.Name = &v
	return s
}

func (s *GetTempFileResponseBody) SetOwnerId(v string) *GetTempFileResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetTempFileResponseBody) SetPrefix(v string) *GetTempFileResponseBody {
	s.Prefix = &v
	return s
}

func (s *GetTempFileResponseBody) SetRequestId(v string) *GetTempFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTempFileResponseBody) SetStatus(v string) *GetTempFileResponseBody {
	s.Status = &v
	return s
}

func (s *GetTempFileResponseBody) SetSuccess(v bool) *GetTempFileResponseBody {
	s.Success = &v
	return s
}

func (s *GetTempFileResponseBody) SetTaskId(v string) *GetTempFileResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetTempFileResponseBody) SetType(v string) *GetTempFileResponseBody {
	s.Type = &v
	return s
}

func (s *GetTempFileResponseBody) SetUploadUrl(v string) *GetTempFileResponseBody {
	s.UploadUrl = &v
	return s
}

func (s *GetTempFileResponseBody) SetUserId(v string) *GetTempFileResponseBody {
	s.UserId = &v
	return s
}

func (s *GetTempFileResponseBody) SetUuid(v string) *GetTempFileResponseBody {
	s.Uuid = &v
	return s
}

func (s *GetTempFileResponseBody) Validate() error {
	return dara.Validate(s)
}
