// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetStorageListRequest
	GetAppId() *string
	SetStatus(v string) *GetStorageListRequest
	GetStatus() *string
	SetStorageType(v string) *GetStorageListRequest
	GetStorageType() *string
}

type GetStorageListRequest struct {
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The OSS storage status.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage type.
	//
	// example:
	//
	// vod_oss_bucket
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s GetStorageListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStorageListRequest) GoString() string {
	return s.String()
}

func (s *GetStorageListRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetStorageListRequest) GetStatus() *string {
	return s.Status
}

func (s *GetStorageListRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *GetStorageListRequest) SetAppId(v string) *GetStorageListRequest {
	s.AppId = &v
	return s
}

func (s *GetStorageListRequest) SetStatus(v string) *GetStorageListRequest {
	s.Status = &v
	return s
}

func (s *GetStorageListRequest) SetStorageType(v string) *GetStorageListRequest {
	s.StorageType = &v
	return s
}

func (s *GetStorageListRequest) Validate() error {
	return dara.Validate(s)
}
