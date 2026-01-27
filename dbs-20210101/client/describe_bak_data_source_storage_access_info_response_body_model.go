// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBakDataSourceStorageAccessInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeBakDataSourceStorageAccessInfoResponseBody
	GetCode() *string
	SetData(v *DescribeBakDataSourceStorageAccessInfoResponseBodyData) *DescribeBakDataSourceStorageAccessInfoResponseBody
	GetData() *DescribeBakDataSourceStorageAccessInfoResponseBodyData
	SetErrCode(v string) *DescribeBakDataSourceStorageAccessInfoResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeBakDataSourceStorageAccessInfoResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeBakDataSourceStorageAccessInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeBakDataSourceStorageAccessInfoResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeBakDataSourceStorageAccessInfoResponseBody
	GetSuccess() *string
}

type DescribeBakDataSourceStorageAccessInfoResponseBody struct {
	// example:
	//
	// Success
	Code *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeBakDataSourceStorageAccessInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Request.Forbidden
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Have no Permissions
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D13761C3-9971-5C02-B789-3F3CD159****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBakDataSourceStorageAccessInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBakDataSourceStorageAccessInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) GetData() *DescribeBakDataSourceStorageAccessInfoResponseBodyData {
	return s.Data
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) SetCode(v string) *DescribeBakDataSourceStorageAccessInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) SetData(v *DescribeBakDataSourceStorageAccessInfoResponseBodyData) *DescribeBakDataSourceStorageAccessInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) SetErrCode(v string) *DescribeBakDataSourceStorageAccessInfoResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) SetErrMessage(v string) *DescribeBakDataSourceStorageAccessInfoResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) SetMessage(v string) *DescribeBakDataSourceStorageAccessInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) SetRequestId(v string) *DescribeBakDataSourceStorageAccessInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) SetSuccess(v string) *DescribeBakDataSourceStorageAccessInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBakDataSourceStorageAccessInfoResponseBodyData struct {
	OssAccessInfo *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo `json:"OssAccessInfo,omitempty" xml:"OssAccessInfo,omitempty" type:"Struct"`
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeBakDataSourceStorageAccessInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBakDataSourceStorageAccessInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyData) GetOssAccessInfo() *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo {
	return s.OssAccessInfo
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyData) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyData) SetOssAccessInfo(v *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) *DescribeBakDataSourceStorageAccessInfoResponseBodyData {
	s.OssAccessInfo = v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyData) SetStorageType(v string) *DescribeBakDataSourceStorageAccessInfoResponseBodyData {
	s.StorageType = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyData) Validate() error {
	if s.OssAccessInfo != nil {
		if err := s.OssAccessInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo struct {
	// example:
	//
	// TMP.3Kk6fNt7hhNmHrGYFkLe5WAo8qL18Hk2rKSZkDkZvrey1ksVAFgJ8Vty8isrBHaH5KUNYAwtcW8HUPzjjsNLo*******
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// 6XzFspxx5wKjuJ2QwATkxnV7fcFxR*****
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// example:
	//
	// dbs-bakdata-test-cn-beijing-backup-oss-0-*****
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// 2024-07-11T05:14:44Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// VN/2023/11/13/activity/20080101/activity-*****.tar
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// example:
	//
	// oss-cn-beijing.****.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// example:
	//
	// oss-cn-beijing-****.aliyuncs.com
	OssInternalEndpoint *string `json:"OssInternalEndpoint,omitempty" xml:"OssInternalEndpoint,omitempty"`
	// example:
	//
	// cn-beijing
	OssRegion *string `json:"OssRegion,omitempty" xml:"OssRegion,omitempty"`
	// example:
	//
	// 554c0098-9858-4871-9c4c-33d9d*****
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) GoString() string {
	return s.String()
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) GetOssInternalEndpoint() *string {
	return s.OssInternalEndpoint
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) GetOssRegion() *string {
	return s.OssRegion
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) SetAccessKeyId(v string) *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) SetAccessKeySecret(v string) *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo {
	s.AccessKeySecret = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) SetBucketName(v string) *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo {
	s.BucketName = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) SetExpiredTime(v string) *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) SetObjectKey(v string) *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo {
	s.ObjectKey = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) SetOssEndpoint(v string) *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) SetOssInternalEndpoint(v string) *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo {
	s.OssInternalEndpoint = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) SetOssRegion(v string) *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo {
	s.OssRegion = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) SetSecurityToken(v string) *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo {
	s.SecurityToken = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoResponseBodyDataOssAccessInfo) Validate() error {
	return dara.Validate(s)
}
