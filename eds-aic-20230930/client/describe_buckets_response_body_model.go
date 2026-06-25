// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBucketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeBucketsResponseBody
	GetCode() *string
	SetData(v []*DescribeBucketsResponseBodyData) *DescribeBucketsResponseBody
	GetData() []*DescribeBucketsResponseBodyData
	SetHttpStatusCode(v int32) *DescribeBucketsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeBucketsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeBucketsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBucketsResponseBody
	GetSuccess() *bool
}

type DescribeBucketsResponseBody struct {
	// The error code.
	//
	// > This parameter is returned only when the call fails.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result object.
	Data []*DescribeBucketsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The return status code of the request.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBucketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBucketsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeBucketsResponseBody) GetData() []*DescribeBucketsResponseBodyData {
	return s.Data
}

func (s *DescribeBucketsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeBucketsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBucketsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBucketsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBucketsResponseBody) SetCode(v string) *DescribeBucketsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBucketsResponseBody) SetData(v []*DescribeBucketsResponseBodyData) *DescribeBucketsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBucketsResponseBody) SetHttpStatusCode(v int32) *DescribeBucketsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBucketsResponseBody) SetMessage(v string) *DescribeBucketsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBucketsResponseBody) SetRequestId(v string) *DescribeBucketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBucketsResponseBody) SetSuccess(v bool) *DescribeBucketsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBucketsResponseBody) Validate() error {
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

type DescribeBucketsResponseBodyData struct {
	// The name of the bucket.
	//
	// example:
	//
	// zydctest
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The public endpoint that is used to access the bucket.
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	ExtranetEndpoint *string `json:"ExtranetEndpoint,omitempty" xml:"ExtranetEndpoint,omitempty"`
	// The time when the bucket was created.
	//
	// example:
	//
	// 2024-05-15 17:33:59
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The internal endpoint that is used by a same-region ECS instance to access the bucket.
	//
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// The storage path, such as a path in Hadoop Distributed File System (HDFS) or OSS.
	//
	// example:
	//
	// center
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The information about the objects in the bucket.
	OssObjectList []*DescribeBucketsResponseBodyDataOssObjectList `json:"OssObjectList,omitempty" xml:"OssObjectList,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBucketsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBucketsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBucketsResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeBucketsResponseBodyData) GetExtranetEndpoint() *string {
	return s.ExtranetEndpoint
}

func (s *DescribeBucketsResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeBucketsResponseBodyData) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *DescribeBucketsResponseBodyData) GetLocation() *string {
	return s.Location
}

func (s *DescribeBucketsResponseBodyData) GetOssObjectList() []*DescribeBucketsResponseBodyDataOssObjectList {
	return s.OssObjectList
}

func (s *DescribeBucketsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBucketsResponseBodyData) SetBucketName(v string) *DescribeBucketsResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *DescribeBucketsResponseBodyData) SetExtranetEndpoint(v string) *DescribeBucketsResponseBodyData {
	s.ExtranetEndpoint = &v
	return s
}

func (s *DescribeBucketsResponseBodyData) SetGmtCreated(v string) *DescribeBucketsResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeBucketsResponseBodyData) SetIntranetEndpoint(v string) *DescribeBucketsResponseBodyData {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeBucketsResponseBodyData) SetLocation(v string) *DescribeBucketsResponseBodyData {
	s.Location = &v
	return s
}

func (s *DescribeBucketsResponseBodyData) SetOssObjectList(v []*DescribeBucketsResponseBodyDataOssObjectList) *DescribeBucketsResponseBodyData {
	s.OssObjectList = v
	return s
}

func (s *DescribeBucketsResponseBodyData) SetRegionId(v string) *DescribeBucketsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeBucketsResponseBodyData) Validate() error {
	if s.OssObjectList != nil {
		for _, item := range s.OssObjectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBucketsResponseBodyDataOssObjectList struct {
	// The name of the bucket.
	//
	// example:
	//
	// tf-testacceu-central-1ensbucketlifecycle44222
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The Entity Tag (ETag) is created when an object is generated. It is used to identify the content of an object.
	//
	// - For an object created by a PutObject request, the ETag value is the MD5 hash of the object content.
	//
	// - For an object created by other methods, the ETag value is a unique value generated based on specific calculation rules. It is not the MD5 hash of the object content.
	//
	// - The ETag value can be used to check if the object content has changed. Do not use the ETag value as the MD5 hash to verify the data integrity of the object.
	//
	// example:
	//
	// 5B3C1A2E053D763E1B002CC607C5A0FE1****
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// The tag key.
	//
	// example:
	//
	// con
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The time when the object was last modified.
	//
	// example:
	//
	// 2012-02-24T08:42:32.000Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The information about the bucket owner.
	Owner *DescribeBucketsResponseBodyDataOssObjectListOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// The information about database recovery if the database type is MSSQL. The value is a JSON string that contains the following parameters:
	//
	// - **name**: The name of the database.
	//
	// - **files**: The path of the database file.
	//
	// example:
	//
	// ongoing-request="true"
	RestoreInfo *string `json:"RestoreInfo,omitempty" xml:"RestoreInfo,omitempty"`
	// The size of the file. Unit: KB.
	//
	// example:
	//
	// 9
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The storage class of the bucket. Valid values:
	//
	// - Standard (default): Standard
	//
	// - IA: Infrequent Access
	//
	// - Archive: Archive Storage
	//
	// - ColdArchive: Cold Archive
	//
	// - DeepColdArchive: Deep Cold Archive
	//
	// example:
	//
	// ARCHIVE
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The type of the trigger method.
	//
	// example:
	//
	// InstanceGroup
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeBucketsResponseBodyDataOssObjectList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBucketsResponseBodyDataOssObjectList) GoString() string {
	return s.String()
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) GetETag() *string {
	return s.ETag
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) GetKey() *string {
	return s.Key
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) GetLastModified() *string {
	return s.LastModified
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) GetOwner() *DescribeBucketsResponseBodyDataOssObjectListOwner {
	return s.Owner
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) GetRestoreInfo() *string {
	return s.RestoreInfo
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) GetSize() *int64 {
	return s.Size
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) GetType() *string {
	return s.Type
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) SetBucketName(v string) *DescribeBucketsResponseBodyDataOssObjectList {
	s.BucketName = &v
	return s
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) SetETag(v string) *DescribeBucketsResponseBodyDataOssObjectList {
	s.ETag = &v
	return s
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) SetKey(v string) *DescribeBucketsResponseBodyDataOssObjectList {
	s.Key = &v
	return s
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) SetLastModified(v string) *DescribeBucketsResponseBodyDataOssObjectList {
	s.LastModified = &v
	return s
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) SetOwner(v *DescribeBucketsResponseBodyDataOssObjectListOwner) *DescribeBucketsResponseBodyDataOssObjectList {
	s.Owner = v
	return s
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) SetRestoreInfo(v string) *DescribeBucketsResponseBodyDataOssObjectList {
	s.RestoreInfo = &v
	return s
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) SetSize(v int64) *DescribeBucketsResponseBodyDataOssObjectList {
	s.Size = &v
	return s
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) SetStorageClass(v string) *DescribeBucketsResponseBodyDataOssObjectList {
	s.StorageClass = &v
	return s
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) SetType(v string) *DescribeBucketsResponseBodyDataOssObjectList {
	s.Type = &v
	return s
}

func (s *DescribeBucketsResponseBodyDataOssObjectList) Validate() error {
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBucketsResponseBodyDataOssObjectListOwner struct {
	// The name of the bucket owner.
	//
	// example:
	//
	// 佳钰
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The unique identifier of the bucket owner.
	//
	// example:
	//
	// 395
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeBucketsResponseBodyDataOssObjectListOwner) String() string {
	return dara.Prettify(s)
}

func (s DescribeBucketsResponseBodyDataOssObjectListOwner) GoString() string {
	return s.String()
}

func (s *DescribeBucketsResponseBodyDataOssObjectListOwner) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeBucketsResponseBodyDataOssObjectListOwner) GetId() *string {
	return s.Id
}

func (s *DescribeBucketsResponseBodyDataOssObjectListOwner) SetDisplayName(v string) *DescribeBucketsResponseBodyDataOssObjectListOwner {
	s.DisplayName = &v
	return s
}

func (s *DescribeBucketsResponseBodyDataOssObjectListOwner) SetId(v string) *DescribeBucketsResponseBodyDataOssObjectListOwner {
	s.Id = &v
	return s
}

func (s *DescribeBucketsResponseBodyDataOssObjectListOwner) Validate() error {
	return dara.Validate(s)
}
