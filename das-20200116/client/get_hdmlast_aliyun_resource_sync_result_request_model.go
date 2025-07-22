// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHDMLastAliyunResourceSyncResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUid(v string) *GetHDMLastAliyunResourceSyncResultRequest
	GetUid() *string
	SetUserId(v string) *GetHDMLastAliyunResourceSyncResultRequest
	GetUserId() *string
	SetContext(v string) *GetHDMLastAliyunResourceSyncResultRequest
	GetContext() *string
	SetAccessKey(v string) *GetHDMLastAliyunResourceSyncResultRequest
	GetAccessKey() *string
	SetSignature(v string) *GetHDMLastAliyunResourceSyncResultRequest
	GetSignature() *string
	SetSkipAuth(v string) *GetHDMLastAliyunResourceSyncResultRequest
	GetSkipAuth() *string
	SetTimestamp(v string) *GetHDMLastAliyunResourceSyncResultRequest
	GetTimestamp() *string
}

type GetHDMLastAliyunResourceSyncResultRequest struct {
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Context   *string `json:"__context,omitempty" xml:"__context,omitempty"`
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	SkipAuth  *string `json:"skipAuth,omitempty" xml:"skipAuth,omitempty"`
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetHDMLastAliyunResourceSyncResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHDMLastAliyunResourceSyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) GetUid() *string {
	return s.Uid
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) GetContext() *string {
	return s.Context
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) GetSignature() *string {
	return s.Signature
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) GetSkipAuth() *string {
	return s.SkipAuth
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetUid(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.Uid = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetUserId(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.UserId = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetContext(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.Context = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetAccessKey(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.AccessKey = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetSignature(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.Signature = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetSkipAuth(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.SkipAuth = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetTimestamp(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.Timestamp = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) Validate() error {
	return dara.Validate(s)
}
