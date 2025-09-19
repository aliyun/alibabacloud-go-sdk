// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageVulListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeImageVulListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeImageVulListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeImageVulListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeImageVulListResponseBody
	GetTotalCount() *int32
	SetVulRecords(v []*DescribeImageVulListResponseBodyVulRecords) *DescribeImageVulListResponseBody
	GetVulRecords() []*DescribeImageVulListResponseBodyVulRecords
}

type DescribeImageVulListResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D6B20156-49B0-5CF0-B14D-7ECA4B50DAAB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of vulnerabilities returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// An array that consists of the vulnerabilities.
	VulRecords []*DescribeImageVulListResponseBodyVulRecords `json:"VulRecords,omitempty" xml:"VulRecords,omitempty" type:"Repeated"`
}

func (s DescribeImageVulListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageVulListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageVulListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageVulListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageVulListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageVulListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageVulListResponseBody) GetVulRecords() []*DescribeImageVulListResponseBodyVulRecords {
	return s.VulRecords
}

func (s *DescribeImageVulListResponseBody) SetCurrentPage(v int32) *DescribeImageVulListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageVulListResponseBody) SetPageSize(v int32) *DescribeImageVulListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeImageVulListResponseBody) SetRequestId(v string) *DescribeImageVulListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageVulListResponseBody) SetTotalCount(v int32) *DescribeImageVulListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageVulListResponseBody) SetVulRecords(v []*DescribeImageVulListResponseBodyVulRecords) *DescribeImageVulListResponseBody {
	s.VulRecords = v
	return s
}

func (s *DescribeImageVulListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageVulListResponseBodyVulRecords struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// CVE-2018-25010:libwebp up to 1.0.0 ApplyFilter out-of-bounds read
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// Indicates whether the vulnerability can be fixed in the Security Center console. Valid values:
	//
	// 	- **yes**: yes
	//
	// 	- **no**: no
	//
	// example:
	//
	// yes
	CanFix *string `json:"CanFix,omitempty" xml:"CanFix,omitempty"`
	// Indicates whether the package of the software that has the vulnerability can be upgraded by using Security Center. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	CanUpdate *bool `json:"CanUpdate,omitempty" xml:"CanUpdate,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// c08d5fc1a329a4b88950a253d082f1****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// docker-law
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the container.
	//
	// example:
	//
	// 04d20e98c8e2c93b7b864372084320a15a58c8671e53c972ce3a71d9c163****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The extended information about the vulnerability.
	ExtendContentJson *DescribeImageVulListResponseBodyVulRecordsExtendContentJson `json:"ExtendContentJson,omitempty" xml:"ExtendContentJson,omitempty" type:"Struct"`
	// The timestamp when the first scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1620752053000
	FirstTs *int64 `json:"FirstTs,omitempty" xml:"FirstTs,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// registry.cn-wulanchabu.aliyuncs.com/sas_test/huxin-test-001:nuxeo6-conta****
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The digest of the image.
	//
	// example:
	//
	// 8f0fbdb41d3d1ade4ffdf21558443f4c03342010563bb8c43ccc09594d507012
	ImageDigest *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// testInstance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 1.2.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.19.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The timestamp when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1631779996000
	LastTs *int64 `json:"LastTs,omitempty" xml:"LastTs,omitempty"`
	// The image layers.
	Layers []*string `json:"Layers,omitempty" xml:"Layers,omitempty" type:"Repeated"`
	// The source of the malicious file. Valid values:
	//
	// 	- **agentless**: agentless detection
	//
	// 	- **image**: image
	//
	// 	- **container**: container
	//
	// example:
	//
	// agentless
	MaliciousSource *string `json:"MaliciousSource,omitempty" xml:"MaliciousSource,omitempty"`
	// The timestamp when the information about the vulnerability was updated. Unit: milliseconds.
	//
	// example:
	//
	// 1580808765000
	ModifyTs *int64 `json:"ModifyTs,omitempty" xml:"ModifyTs,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// debian:10:CVE-2019-9893
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// test-002
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
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
	// The pod.
	//
	// example:
	//
	// 22222-7xsqq
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The ID of the vulnerability.
	//
	// example:
	//
	// 782661
	PrimaryId *int64 `json:"PrimaryId,omitempty" xml:"PrimaryId,omitempty"`
	// The Common Vulnerabilities and Exposures (CVE) ID of the associated vulnerability.
	//
	// example:
	//
	// CVE-2019-9893
	Related *string `json:"Related,omitempty" xml:"Related,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// varnish
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// default
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The tag of this vulnerability. Valid values:
	//
	// 	- **AI**: AI-related components.
	//
	// example:
	//
	// AI
	RuleTag *string `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
	// The time at which the scan was performed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1649814050000
	ScanTime *int64 `json:"ScanTime,omitempty" xml:"ScanTime,omitempty"`
	// The status of the vulnerability. Valid values:
	//
	// 	- **1**: unfixed
	//
	// 	- **7**: fixed
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag that is added to the image vulnerability.
	//
	// example:
	//
	// oval
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ID of the asset on which the vulnerability is detected.
	//
	// example:
	//
	// m-bp17m0pc0xprzbwo****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the asset on which the vulnerability is detected.
	//
	// example:
	//
	// source-test-obj-XM0Ma
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the asset on which the vulnerability is detected. Valid values:
	//
	// 	- **ECS_SNAPSHOT**: snapshot
	//
	// 	- **ECS_IMAGE**: image
	//
	// example:
	//
	// ECS_IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the vulnerability. The value is fixed as CVE, which indicates image vulnerabilities.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 0004a32a0305a7f6ab5ff9600d47****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeImageVulListResponseBodyVulRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageVulListResponseBodyVulRecords) GoString() string {
	return s.String()
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetCanFix() *string {
	return s.CanFix
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetCanUpdate() *bool {
	return s.CanUpdate
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetExtendContentJson() *DescribeImageVulListResponseBodyVulRecordsExtendContentJson {
	return s.ExtendContentJson
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetFirstTs() *int64 {
	return s.FirstTs
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetImage() *string {
	return s.Image
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetLastTs() *int64 {
	return s.LastTs
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetLayers() []*string {
	return s.Layers
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetMaliciousSource() *string {
	return s.MaliciousSource
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetModifyTs() *int64 {
	return s.ModifyTs
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetName() *string {
	return s.Name
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetPod() *string {
	return s.Pod
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetRelated() *string {
	return s.Related
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetRuleTag() *string {
	return s.RuleTag
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetScanTime() *int64 {
	return s.ScanTime
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetTag() *string {
	return s.Tag
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetTargetId() *string {
	return s.TargetId
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetTargetName() *string {
	return s.TargetName
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetType() *string {
	return s.Type
}

func (s *DescribeImageVulListResponseBodyVulRecords) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetAliasName(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.AliasName = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetCanFix(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.CanFix = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetCanUpdate(v bool) *DescribeImageVulListResponseBodyVulRecords {
	s.CanUpdate = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetClusterId(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.ClusterId = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetClusterName(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.ClusterName = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetContainerId(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.ContainerId = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetExtendContentJson(v *DescribeImageVulListResponseBodyVulRecordsExtendContentJson) *DescribeImageVulListResponseBodyVulRecords {
	s.ExtendContentJson = v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetFirstTs(v int64) *DescribeImageVulListResponseBodyVulRecords {
	s.FirstTs = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetImage(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Image = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetImageDigest(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.ImageDigest = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetInstanceName(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.InstanceName = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetInternetIp(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.InternetIp = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetIntranetIp(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.IntranetIp = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetLastTs(v int64) *DescribeImageVulListResponseBodyVulRecords {
	s.LastTs = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetLayers(v []*string) *DescribeImageVulListResponseBodyVulRecords {
	s.Layers = v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetMaliciousSource(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.MaliciousSource = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetModifyTs(v int64) *DescribeImageVulListResponseBodyVulRecords {
	s.ModifyTs = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetName(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Name = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetNamespace(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Namespace = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetNecessity(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Necessity = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetPod(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Pod = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetPrimaryId(v int64) *DescribeImageVulListResponseBodyVulRecords {
	s.PrimaryId = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetRelated(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Related = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetRepoName(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.RepoName = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetRepoNamespace(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetRuleTag(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.RuleTag = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetScanTime(v int64) *DescribeImageVulListResponseBodyVulRecords {
	s.ScanTime = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetStatus(v int32) *DescribeImageVulListResponseBodyVulRecords {
	s.Status = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetTag(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Tag = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetTargetId(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.TargetId = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetTargetName(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.TargetName = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetTargetType(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.TargetType = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetType(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Type = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetUuid(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Uuid = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) Validate() error {
	return dara.Validate(s)
}

type DescribeImageVulListResponseBodyVulRecordsExtendContentJson struct {
	// The name of the operating system.
	//
	// example:
	//
	// debian
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The version of the operating system in the image.
	//
	// example:
	//
	// 10.9
	OsRelease *string `json:"OsRelease,omitempty" xml:"OsRelease,omitempty"`
	// The details of the package of the software that has the vulnerability.
	RpmEntityList []*DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList `json:"RpmEntityList,omitempty" xml:"RpmEntityList,omitempty" type:"Repeated"`
}

func (s DescribeImageVulListResponseBodyVulRecordsExtendContentJson) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageVulListResponseBodyVulRecordsExtendContentJson) GoString() string {
	return s.String()
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJson) GetOs() *string {
	return s.Os
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJson) GetOsRelease() *string {
	return s.OsRelease
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJson) GetRpmEntityList() []*DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	return s.RpmEntityList
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJson) SetOs(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJson {
	s.Os = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJson) SetOsRelease(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJson {
	s.OsRelease = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJson) SetRpmEntityList(v []*DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) *DescribeImageVulListResponseBodyVulRecordsExtendContentJson {
	s.RpmEntityList = v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJson) Validate() error {
	return dara.Validate(s)
}

type DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList struct {
	// The complete version number of the package.
	//
	// example:
	//
	// 2.3.3-4
	FullVersion *string `json:"FullVersion,omitempty" xml:"FullVersion,omitempty"`
	// The SHA-256 value of the digest of the image layer.
	//
	// example:
	//
	// b1f5b9420803ad0657cf21566e3e20acc08581e7f22991249ef3aa80b8b1c587
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The reason why the vulnerability is detected.
	//
	// example:
	//
	// libseccomp2 version less than equals 2.3.3-4
	MatchDetail *string `json:"MatchDetail,omitempty" xml:"MatchDetail,omitempty"`
	// The details of the rule that is used to detect the vulnerability.
	//
	// example:
	//
	// ["libseccomp2 version less than equals 2.3.3-4"]
	MatchList []*string `json:"MatchList,omitempty" xml:"MatchList,omitempty" type:"Repeated"`
	// The name of the software package.
	//
	// example:
	//
	// libseccomp2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The path of the software that has the vulnerability.
	//
	// example:
	//
	// /usr/lib64/libssh2.so.1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The command that is used to fix the vulnerability.
	//
	// example:
	//
	// apt-get update && apt-get install libseccomp2  --only-upgrade
	UpdateCmd *string `json:"UpdateCmd,omitempty" xml:"UpdateCmd,omitempty"`
	// The version number of the package.
	//
	// example:
	//
	// 2.3.3-4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GoString() string {
	return s.String()
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetFullVersion() *string {
	return s.FullVersion
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetLayer() *string {
	return s.Layer
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetMatchDetail() *string {
	return s.MatchDetail
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetMatchList() []*string {
	return s.MatchList
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetName() *string {
	return s.Name
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetPath() *string {
	return s.Path
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetUpdateCmd() *string {
	return s.UpdateCmd
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetVersion() *string {
	return s.Version
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetFullVersion(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.FullVersion = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetLayer(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Layer = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetMatchDetail(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.MatchDetail = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetMatchList(v []*string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.MatchList = v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetName(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Name = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetPath(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Path = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetUpdateCmd(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.UpdateCmd = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetVersion(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Version = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) Validate() error {
	return dara.Validate(s)
}
