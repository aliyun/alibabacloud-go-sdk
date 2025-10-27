// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageInfoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageInfos(v []*DescribeImageInfoListResponseBodyImageInfos) *DescribeImageInfoListResponseBody
	GetImageInfos() []*DescribeImageInfoListResponseBodyImageInfos
	SetRequestId(v string) *DescribeImageInfoListResponseBody
	GetRequestId() *string
}

type DescribeImageInfoListResponseBody struct {
	// An array that consists of the information about images.
	ImageInfos []*DescribeImageInfoListResponseBodyImageInfos `json:"ImageInfos,omitempty" xml:"ImageInfos,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BC3B0DAE-CC0E-59E9-9383-6F060F22****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageInfoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageInfoListResponseBody) GetImageInfos() []*DescribeImageInfoListResponseBodyImageInfos {
	return s.ImageInfos
}

func (s *DescribeImageInfoListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageInfoListResponseBody) SetImageInfos(v []*DescribeImageInfoListResponseBodyImageInfos) *DescribeImageInfoListResponseBody {
	s.ImageInfos = v
	return s
}

func (s *DescribeImageInfoListResponseBody) SetRequestId(v string) *DescribeImageInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageInfoListResponseBody) Validate() error {
	if s.ImageInfos != nil {
		for _, item := range s.ImageInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageInfoListResponseBodyImageInfos struct {
	// The number of alerts that are generated on the current pod, application, namespace, or cluster.
	//
	// example:
	//
	// 10
	AlarmCount *int32 `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	// Indicates whether alerts are generated on the asset. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// YES
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// The digest value of the image.
	//
	// example:
	//
	// a3521b04dfdd1361a24be6263f2983cf12ba910989f4d9f7324da7e1e89f****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The endpoint of Container Registry.
	//
	// example:
	//
	// cn-hangzhou-x7
	Endpoints *string `json:"Endpoints,omitempty" xml:"Endpoints,omitempty"`
	// The time when the image was created.
	//
	// example:
	//
	// 1636962328000
	ImageCreate *int64 `json:"ImageCreate,omitempty" xml:"ImageCreate,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// bb0175afea16138815a8900adeeb0315d88a83a2376eeffa14db1d693a15****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The size of the image.
	//
	// example:
	//
	// 157408623
	ImageSize *int64 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The time when the image was updated.
	//
	// example:
	//
	// 1636974116000
	ImageUpdate *int64 `json:"ImageUpdate,omitempty" xml:"ImageUpdate,omitempty"`
	// The ID of the image instance.
	//
	// example:
	//
	// i-wz95abw6pa7y79ve****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the image instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the registration.
	//
	// example:
	//
	// harbor
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-5qk9v2rdt0s****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// opa-test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// libssh2
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The type of the image repository. Valid values:
	//
	// 	- `PUBLIC`
	//
	// 	- `PRIVATE`
	//
	// example:
	//
	// PUBLIC
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// Indicates whether the image is at risk. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// YES
	RiskStatus *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// The usage label of the image.
	//
	// example:
	//
	// PAI
	SourceBizTag *string `json:"SourceBizTag,omitempty" xml:"SourceBizTag,omitempty"`
	// The status of the image.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag that is added to the image.
	//
	// example:
	//
	// v1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The tag immutability.
	//
	// example:
	//
	// 0
	TagImmutable *int32 `json:"TagImmutable,omitempty" xml:"TagImmutable,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// ff9ca084-7faa-4ab2-8728-69024755****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The total number of vulnerabilities in your assets.
	//
	// example:
	//
	// 1
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
	// Indicates whether vulnerabilities are detected on the asset. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// YES
	VulStatus *string `json:"VulStatus,omitempty" xml:"VulStatus,omitempty"`
}

func (s DescribeImageInfoListResponseBodyImageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInfoListResponseBodyImageInfos) GoString() string {
	return s.String()
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetAlarmCount() *int32 {
	return s.AlarmCount
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetAlarmStatus() *string {
	return s.AlarmStatus
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetDigest() *string {
	return s.Digest
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetEndpoints() *string {
	return s.Endpoints
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetImageCreate() *int64 {
	return s.ImageCreate
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetImageSize() *int64 {
	return s.ImageSize
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetImageUpdate() *int64 {
	return s.ImageUpdate
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetRegistryType() *string {
	return s.RegistryType
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetRepoId() *string {
	return s.RepoId
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetRepoType() *string {
	return s.RepoType
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetRiskStatus() *string {
	return s.RiskStatus
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetSourceBizTag() *string {
	return s.SourceBizTag
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetTag() *string {
	return s.Tag
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetTagImmutable() *int32 {
	return s.TagImmutable
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetVulCount() *int32 {
	return s.VulCount
}

func (s *DescribeImageInfoListResponseBodyImageInfos) GetVulStatus() *string {
	return s.VulStatus
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetAlarmCount(v int32) *DescribeImageInfoListResponseBodyImageInfos {
	s.AlarmCount = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetAlarmStatus(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.AlarmStatus = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetDigest(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.Digest = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetEndpoints(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.Endpoints = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetImageCreate(v int64) *DescribeImageInfoListResponseBodyImageInfos {
	s.ImageCreate = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetImageId(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.ImageId = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetImageSize(v int64) *DescribeImageInfoListResponseBodyImageInfos {
	s.ImageSize = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetImageUpdate(v int64) *DescribeImageInfoListResponseBodyImageInfos {
	s.ImageUpdate = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetInstanceId(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.InstanceId = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetRegionId(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.RegionId = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetRegistryType(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.RegistryType = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetRepoId(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.RepoId = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetRepoName(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.RepoName = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetRepoNamespace(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetRepoType(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.RepoType = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetRiskStatus(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.RiskStatus = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetSourceBizTag(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.SourceBizTag = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetStatus(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.Status = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetTag(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.Tag = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetTagImmutable(v int32) *DescribeImageInfoListResponseBodyImageInfos {
	s.TagImmutable = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetUuid(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.Uuid = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetVulCount(v int32) *DescribeImageInfoListResponseBodyImageInfos {
	s.VulCount = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) SetVulStatus(v string) *DescribeImageInfoListResponseBodyImageInfos {
	s.VulStatus = &v
	return s
}

func (s *DescribeImageInfoListResponseBodyImageInfos) Validate() error {
	return dara.Validate(s)
}
