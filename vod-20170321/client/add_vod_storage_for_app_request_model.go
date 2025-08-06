// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVodStorageForAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AddVodStorageForAppRequest
	GetAppId() *string
	SetStorageLocation(v string) *AddVodStorageForAppRequest
	GetStorageLocation() *string
	SetStorageType(v string) *AddVodStorageForAppRequest
	GetStorageType() *string
}

type AddVodStorageForAppRequest struct {
	// The IDs of applications. You can obtain the application ID from the `AppId` parameter in the response to the [CreateAppInfo](~~CreateAppInfo~~) or [ListAppInfo](~~ListAppInfo~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The address of an Object Storage Service (OSS) bucket. This parameter does not take effect. You can call this operation to add only VOD buckets.
	//
	// example:
	//
	// out-****.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The storage class. Default value: **vod_oss_bucket**.
	//
	// example:
	//
	// vod_oss_bucket
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s AddVodStorageForAppRequest) String() string {
	return dara.Prettify(s)
}

func (s AddVodStorageForAppRequest) GoString() string {
	return s.String()
}

func (s *AddVodStorageForAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddVodStorageForAppRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *AddVodStorageForAppRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *AddVodStorageForAppRequest) SetAppId(v string) *AddVodStorageForAppRequest {
	s.AppId = &v
	return s
}

func (s *AddVodStorageForAppRequest) SetStorageLocation(v string) *AddVodStorageForAppRequest {
	s.StorageLocation = &v
	return s
}

func (s *AddVodStorageForAppRequest) SetStorageType(v string) *AddVodStorageForAppRequest {
	s.StorageType = &v
	return s
}

func (s *AddVodStorageForAppRequest) Validate() error {
	return dara.Validate(s)
}
