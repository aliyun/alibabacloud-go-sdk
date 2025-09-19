// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageGroupedVulListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *DescribeImageGroupedVulListRequest
	GetAliasName() *string
	SetClusterId(v string) *DescribeImageGroupedVulListRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *DescribeImageGroupedVulListRequest
	GetCurrentPage() *int32
	SetCveId(v string) *DescribeImageGroupedVulListRequest
	GetCveId() *string
	SetGroupId(v string) *DescribeImageGroupedVulListRequest
	GetGroupId() *string
	SetImageDigest(v string) *DescribeImageGroupedVulListRequest
	GetImageDigest() *string
	SetImageLayer(v string) *DescribeImageGroupedVulListRequest
	GetImageLayer() *string
	SetImageTag(v string) *DescribeImageGroupedVulListRequest
	GetImageTag() *string
	SetIsLatest(v int32) *DescribeImageGroupedVulListRequest
	GetIsLatest() *int32
	SetLang(v string) *DescribeImageGroupedVulListRequest
	GetLang() *string
	SetName(v string) *DescribeImageGroupedVulListRequest
	GetName() *string
	SetNecessity(v string) *DescribeImageGroupedVulListRequest
	GetNecessity() *string
	SetPageSize(v int32) *DescribeImageGroupedVulListRequest
	GetPageSize() *int32
	SetPatchId(v int64) *DescribeImageGroupedVulListRequest
	GetPatchId() *int64
	SetRepoId(v string) *DescribeImageGroupedVulListRequest
	GetRepoId() *string
	SetRepoInstanceId(v string) *DescribeImageGroupedVulListRequest
	GetRepoInstanceId() *string
	SetRepoName(v string) *DescribeImageGroupedVulListRequest
	GetRepoName() *string
	SetRepoNamespace(v string) *DescribeImageGroupedVulListRequest
	GetRepoNamespace() *string
	SetRepoRegionId(v string) *DescribeImageGroupedVulListRequest
	GetRepoRegionId() *string
	SetRuleTag(v string) *DescribeImageGroupedVulListRequest
	GetRuleTag() *string
	SetScanRange(v []*string) *DescribeImageGroupedVulListRequest
	GetScanRange() []*string
	SetType(v string) *DescribeImageGroupedVulListRequest
	GetType() *string
	SetUuids(v string) *DescribeImageGroupedVulListRequest
	GetUuids() *string
}

type DescribeImageGroupedVulListRequest struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// High severity vulnerability that affects org.eclipse.jetty:jetty-server
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The ID of the container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the ID of the container cluster.
	//
	// example:
	//
	// c60b77fe62093480db6164a3c2fa5****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The Common Vulnerabilities and Exposures (CVE) ID of the vulnerability.
	//
	// example:
	//
	// CVE-2017-15420
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// The ID of the asset group.
	//
	// example:
	//
	// 1311
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The SHA-256 value of the image digest.
	//
	// example:
	//
	// w213412341dfsfasdfafadfasfasf
	ImageDigest *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	// The layer of the image.
	//
	// example:
	//
	// b1f5b9420803ad0657cf21566e3e20acc08581e7f22991249ef3aa80b8b1c587
	ImageLayer *string `json:"ImageLayer,omitempty" xml:"ImageLayer,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// oval
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// Specifies whether to query the vulnerabilities in the latest images. If you do not specify this parameter, the vulnerabilities in all images are queried. Valid values:
	//
	// 	- **0**: does not query the vulnerabilities in the latest images.
	//
	// 	- **1**: queries the vulnerabilities in the latest images.
	//
	// example:
	//
	// 0
	IsLatest *int32 `json:"IsLatest,omitempty" xml:"IsLatest,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// debian:10:CVE-2019-9893
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority to fix the vulnerability. Valid values:
	//
	// 	- **asap**: high. You must fix the vulnerability at the earliest opportunity.
	//
	// 	- **later**: medium. You can fix the vulnerability based on your business requirements.
	//
	// 	- **nntf**: low. You can ignore the vulnerability.
	//
	// example:
	//
	// asap
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the patch that is used to fix the vulnerability.
	//
	// example:
	//
	// 1341512412
	PatchId *int64 `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// qew****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The instance ID of the image repository.
	//
	// example:
	//
	// i-qewqrqcsadf****
	RepoInstanceId *string `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// libssh2
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// libssh2
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The region ID of the image repository.
	//
	// example:
	//
	// cn-hangzhou
	RepoRegionId *string `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	// The tag of this vulnerability. Valid values:
	//
	// 	- **AI**: AI-related components.
	//
	// example:
	//
	// AI
	RuleTag *string `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
	// An array consisting of the types of the assets that you want to scan.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
	// The type of the vulnerability that you want to query. Valid values:
	//
	// 	- **cve**: image system vulnerability
	//
	// 	- **sca**: image application vulnerability
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the asset. Separate multiple UUIDs with commas (,).
	//
	// example:
	//
	// uuid-13134124****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeImageGroupedVulListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageGroupedVulListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageGroupedVulListRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeImageGroupedVulListRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeImageGroupedVulListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageGroupedVulListRequest) GetCveId() *string {
	return s.CveId
}

func (s *DescribeImageGroupedVulListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeImageGroupedVulListRequest) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *DescribeImageGroupedVulListRequest) GetImageLayer() *string {
	return s.ImageLayer
}

func (s *DescribeImageGroupedVulListRequest) GetImageTag() *string {
	return s.ImageTag
}

func (s *DescribeImageGroupedVulListRequest) GetIsLatest() *int32 {
	return s.IsLatest
}

func (s *DescribeImageGroupedVulListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageGroupedVulListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeImageGroupedVulListRequest) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeImageGroupedVulListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageGroupedVulListRequest) GetPatchId() *int64 {
	return s.PatchId
}

func (s *DescribeImageGroupedVulListRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *DescribeImageGroupedVulListRequest) GetRepoInstanceId() *string {
	return s.RepoInstanceId
}

func (s *DescribeImageGroupedVulListRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageGroupedVulListRequest) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageGroupedVulListRequest) GetRepoRegionId() *string {
	return s.RepoRegionId
}

func (s *DescribeImageGroupedVulListRequest) GetRuleTag() *string {
	return s.RuleTag
}

func (s *DescribeImageGroupedVulListRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeImageGroupedVulListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeImageGroupedVulListRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeImageGroupedVulListRequest) SetAliasName(v string) *DescribeImageGroupedVulListRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetClusterId(v string) *DescribeImageGroupedVulListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetCurrentPage(v int32) *DescribeImageGroupedVulListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetCveId(v string) *DescribeImageGroupedVulListRequest {
	s.CveId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetGroupId(v string) *DescribeImageGroupedVulListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetImageDigest(v string) *DescribeImageGroupedVulListRequest {
	s.ImageDigest = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetImageLayer(v string) *DescribeImageGroupedVulListRequest {
	s.ImageLayer = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetImageTag(v string) *DescribeImageGroupedVulListRequest {
	s.ImageTag = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetIsLatest(v int32) *DescribeImageGroupedVulListRequest {
	s.IsLatest = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetLang(v string) *DescribeImageGroupedVulListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetName(v string) *DescribeImageGroupedVulListRequest {
	s.Name = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetNecessity(v string) *DescribeImageGroupedVulListRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetPageSize(v int32) *DescribeImageGroupedVulListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetPatchId(v int64) *DescribeImageGroupedVulListRequest {
	s.PatchId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetRepoId(v string) *DescribeImageGroupedVulListRequest {
	s.RepoId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetRepoInstanceId(v string) *DescribeImageGroupedVulListRequest {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetRepoName(v string) *DescribeImageGroupedVulListRequest {
	s.RepoName = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetRepoNamespace(v string) *DescribeImageGroupedVulListRequest {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetRepoRegionId(v string) *DescribeImageGroupedVulListRequest {
	s.RepoRegionId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetRuleTag(v string) *DescribeImageGroupedVulListRequest {
	s.RuleTag = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetScanRange(v []*string) *DescribeImageGroupedVulListRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetType(v string) *DescribeImageGroupedVulListRequest {
	s.Type = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetUuids(v string) *DescribeImageGroupedVulListRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) Validate() error {
	return dara.Validate(s)
}
