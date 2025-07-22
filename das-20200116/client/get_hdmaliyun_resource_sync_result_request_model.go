// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHDMAliyunResourceSyncResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetHDMAliyunResourceSyncResultRequest
	GetTaskId() *string
	SetUid(v string) *GetHDMAliyunResourceSyncResultRequest
	GetUid() *string
	SetUserId(v string) *GetHDMAliyunResourceSyncResultRequest
	GetUserId() *string
	SetContext(v string) *GetHDMAliyunResourceSyncResultRequest
	GetContext() *string
	SetAccessKey(v string) *GetHDMAliyunResourceSyncResultRequest
	GetAccessKey() *string
	SetSignature(v string) *GetHDMAliyunResourceSyncResultRequest
	GetSignature() *string
	SetSkipAuth(v string) *GetHDMAliyunResourceSyncResultRequest
	GetSkipAuth() *string
	SetTimestamp(v string) *GetHDMAliyunResourceSyncResultRequest
	GetTimestamp() *string
}

type GetHDMAliyunResourceSyncResultRequest struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Context   *string `json:"__context,omitempty" xml:"__context,omitempty"`
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	SkipAuth  *string `json:"skipAuth,omitempty" xml:"skipAuth,omitempty"`
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetHDMAliyunResourceSyncResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHDMAliyunResourceSyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetHDMAliyunResourceSyncResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetHDMAliyunResourceSyncResultRequest) GetUid() *string {
	return s.Uid
}

func (s *GetHDMAliyunResourceSyncResultRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetHDMAliyunResourceSyncResultRequest) GetContext() *string {
	return s.Context
}

func (s *GetHDMAliyunResourceSyncResultRequest) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetHDMAliyunResourceSyncResultRequest) GetSignature() *string {
	return s.Signature
}

func (s *GetHDMAliyunResourceSyncResultRequest) GetSkipAuth() *string {
	return s.SkipAuth
}

func (s *GetHDMAliyunResourceSyncResultRequest) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetTaskId(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetUid(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.Uid = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetUserId(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.UserId = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetContext(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.Context = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetAccessKey(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.AccessKey = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetSignature(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.Signature = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetSkipAuth(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.SkipAuth = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetTimestamp(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.Timestamp = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) Validate() error {
	return dara.Validate(s)
}
