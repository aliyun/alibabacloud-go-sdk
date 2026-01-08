// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAppsResponseBody
	GetRequestId() *string
	SetResult(v []*ListAppsResponseBodyResult) *ListAppsResponseBody
	GetResult() []*ListAppsResponseBodyResult
	SetTotalCount(v int32) *ListAppsResponseBody
	GetTotalCount() *int32
}

type ListAppsResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListAppsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppsResponseBody) GetResult() []*ListAppsResponseBodyResult {
	return s.Result
}

func (s *ListAppsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsResponseBody) SetResult(v []*ListAppsResponseBodyResult) *ListAppsResponseBody {
	s.Result = v
	return s
}

func (s *ListAppsResponseBody) SetTotalCount(v int32) *ListAppsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAppsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppsResponseBodyResult struct {
	// example:
	//
	// test-abc
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// es-severless-test-app
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// 2022-12-27T07:09:11.000Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 应用备注
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 2022-12-27T07:09:11.000Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// OwnerID账号ID
	//
	// example:
	//
	// *********7595
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// 代表region的资源属性字段
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 代表资源状态的资源属性字段
	//
	// example:
	//
	// ACTIVE
	Status *string                           `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*ListAppsResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 7.10
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListAppsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAppsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyResult) GetAppId() *string {
	return s.AppId
}

func (s *ListAppsResponseBodyResult) GetAppName() *string {
	return s.AppName
}

func (s *ListAppsResponseBodyResult) GetAppType() *string {
	return s.AppType
}

func (s *ListAppsResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAppsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListAppsResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAppsResponseBodyResult) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListAppsResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAppsResponseBodyResult) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAppsResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListAppsResponseBodyResult) GetTags() []*ListAppsResponseBodyResultTags {
	return s.Tags
}

func (s *ListAppsResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *ListAppsResponseBodyResult) SetAppId(v string) *ListAppsResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetAppName(v string) *ListAppsResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetAppType(v string) *ListAppsResponseBodyResult {
	s.AppType = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetCreateTime(v string) *ListAppsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetDescription(v string) *ListAppsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetInstanceId(v string) *ListAppsResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetModifiedTime(v string) *ListAppsResponseBodyResult {
	s.ModifiedTime = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetOwnerId(v string) *ListAppsResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetRegionId(v string) *ListAppsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetStatus(v string) *ListAppsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetTags(v []*ListAppsResponseBodyResultTags) *ListAppsResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListAppsResponseBodyResult) SetVersion(v string) *ListAppsResponseBodyResult {
	s.Version = &v
	return s
}

func (s *ListAppsResponseBodyResult) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppsResponseBodyResultTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAppsResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s ListAppsResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyResultTags) GetKey() *string {
	return s.Key
}

func (s *ListAppsResponseBodyResultTags) GetValue() *string {
	return s.Value
}

func (s *ListAppsResponseBodyResultTags) SetKey(v string) *ListAppsResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *ListAppsResponseBodyResultTags) SetValue(v string) *ListAppsResponseBodyResultTags {
	s.Value = &v
	return s
}

func (s *ListAppsResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}
