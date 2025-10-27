// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCanFixVulListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCanFixVulListResponseBody
	GetRequestId() *string
	SetVulRecords(v []*DescribeCanFixVulListResponseBodyVulRecords) *DescribeCanFixVulListResponseBody
	GetVulRecords() []*DescribeCanFixVulListResponseBodyVulRecords
}

type DescribeCanFixVulListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1408FDB3-46F4-513C-9918-FE7D356DF048
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the vulnerability.
	VulRecords []*DescribeCanFixVulListResponseBodyVulRecords `json:"VulRecords,omitempty" xml:"VulRecords,omitempty" type:"Repeated"`
}

func (s DescribeCanFixVulListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCanFixVulListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCanFixVulListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCanFixVulListResponseBody) GetVulRecords() []*DescribeCanFixVulListResponseBodyVulRecords {
	return s.VulRecords
}

func (s *DescribeCanFixVulListResponseBody) SetRequestId(v string) *DescribeCanFixVulListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCanFixVulListResponseBody) SetVulRecords(v []*DescribeCanFixVulListResponseBodyVulRecords) *DescribeCanFixVulListResponseBody {
	s.VulRecords = v
	return s
}

func (s *DescribeCanFixVulListResponseBody) Validate() error {
	if s.VulRecords != nil {
		for _, item := range s.VulRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCanFixVulListResponseBodyVulRecords struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// CVE-2018-25010:libwebp up to 1.0.0 ApplyFilter out-of-bounds read
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// Indicates whether the vulnerability can be fixed in the Security Center console. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// yes
	CanFix *string `json:"CanFix,omitempty" xml:"CanFix,omitempty"`
	// Indicates whether the packages of the software that has the vulnerability can be upgraded by using Security Center. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CanUpdate *bool `json:"CanUpdate,omitempty" xml:"CanUpdate,omitempty"`
	// The cluster ID.
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
	// The container ID.
	//
	// example:
	//
	// 04d20e98c8e2c93b7b864372084320a15a58c8671e53c972ce3a71d9c163****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The extended information about the vulnerability.
	ExtendContentJson *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson `json:"ExtendContentJson,omitempty" xml:"ExtendContentJson,omitempty" type:"Struct"`
	// The timestamp generated when the vulnerability was first detected. Unit: milliseconds.
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
	// The unique identifier of the image.
	//
	// example:
	//
	// 8f0fbdb41d3d1ade4ffdf21558443f4c03342010563bb8c43ccc09594d50****
	ImageDigest *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	// The name of the instance.
	//
	// The name must be 3 to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// testInstance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the asset.
	//
	// example:
	//
	// 1.2.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the asset.
	//
	// example:
	//
	// 172.19.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The timestamp generated when the vulnerability was last detected. Unit: milliseconds.
	//
	// example:
	//
	// 1620404763000
	LastTs *int64 `json:"LastTs,omitempty" xml:"LastTs,omitempty"`
	// The image layers.
	Layers []*string `json:"Layers,omitempty" xml:"Layers,omitempty" type:"Repeated"`
	// The source of the malicious file. Valid values:
	//
	// 	- **agentless**
	//
	// 	- **image**
	//
	// 	- **container**
	//
	// example:
	//
	// agentless
	MaliciousSource *string `json:"MaliciousSource,omitempty" xml:"MaliciousSource,omitempty"`
	// The timestamp generated when the vulnerability status was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1620404763000
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
	// 	- **asap**: high
	//
	// 	- **later**: medium
	//
	// 	- **nntf**: low
	//
	// >  We recommend that you fix high-level vulnerabilities as soon as possible.
	//
	// example:
	//
	// asap,later,nntf
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// The name of the container group.
	//
	// example:
	//
	// 22222-7xsqq
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The vulnerability ID.
	//
	// example:
	//
	// 782661
	PrimaryId *int64 `json:"PrimaryId,omitempty" xml:"PrimaryId,omitempty"`
	// The CVE IDs related to the vulnerability. Multiple CVE IDs are separated by commas (,).
	//
	// example:
	//
	// CVE-2017-7518,CVE-2017-12188
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
	// 3rdparty
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The timestamp generated when the scan task was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1649814050000
	ScanTime *int64 `json:"ScanTime,omitempty" xml:"ScanTime,omitempty"`
	// The status of the vulnerability. Valid values:
	//
	// 	- **1**: The vulnerability is unfixed.
	//
	// 	- **4**: The vulnerability is being fixed.
	//
	// 	- **7**: The vulnerability is fixed.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag that is added to the image.
	//
	// example:
	//
	// latest
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ID of the asset that is scanned.
	//
	// example:
	//
	// 300269
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the asset that is scanned.
	//
	// example:
	//
	// source-test-obj-XM0Ma
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the asset that is scanned. Valid values:
	//
	// 	- **IMAGE**
	//
	// 	- **ECS_IMAGE**
	//
	// 	- **ECS_SNAPSHOT**
	//
	// example:
	//
	// ECS_IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: system vulnerability
	//
	// 	- **sca**: application vulnerability
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the container image.
	//
	// example:
	//
	// 0004a32a0305a7f6ab5ff9600d47****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeCanFixVulListResponseBodyVulRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeCanFixVulListResponseBodyVulRecords) GoString() string {
	return s.String()
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetCanFix() *string {
	return s.CanFix
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetCanUpdate() *bool {
	return s.CanUpdate
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetExtendContentJson() *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson {
	return s.ExtendContentJson
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetFirstTs() *int64 {
	return s.FirstTs
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetImage() *string {
	return s.Image
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetLastTs() *int64 {
	return s.LastTs
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetLayers() []*string {
	return s.Layers
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetMaliciousSource() *string {
	return s.MaliciousSource
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetModifyTs() *int64 {
	return s.ModifyTs
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetName() *string {
	return s.Name
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetPod() *string {
	return s.Pod
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetRelated() *string {
	return s.Related
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetScanTime() *int64 {
	return s.ScanTime
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetTag() *string {
	return s.Tag
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetTargetId() *string {
	return s.TargetId
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetTargetName() *string {
	return s.TargetName
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetType() *string {
	return s.Type
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetAliasName(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.AliasName = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetCanFix(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.CanFix = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetCanUpdate(v bool) *DescribeCanFixVulListResponseBodyVulRecords {
	s.CanUpdate = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetClusterId(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.ClusterId = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetClusterName(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.ClusterName = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetContainerId(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.ContainerId = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetExtendContentJson(v *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson) *DescribeCanFixVulListResponseBodyVulRecords {
	s.ExtendContentJson = v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetFirstTs(v int64) *DescribeCanFixVulListResponseBodyVulRecords {
	s.FirstTs = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetImage(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.Image = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetImageDigest(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.ImageDigest = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetInstanceName(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.InstanceName = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetInternetIp(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.InternetIp = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetIntranetIp(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.IntranetIp = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetLastTs(v int64) *DescribeCanFixVulListResponseBodyVulRecords {
	s.LastTs = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetLayers(v []*string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.Layers = v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetMaliciousSource(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.MaliciousSource = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetModifyTs(v int64) *DescribeCanFixVulListResponseBodyVulRecords {
	s.ModifyTs = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetName(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.Name = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetNamespace(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.Namespace = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetNecessity(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.Necessity = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetPod(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.Pod = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetPrimaryId(v int64) *DescribeCanFixVulListResponseBodyVulRecords {
	s.PrimaryId = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetRelated(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.Related = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetRepoName(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.RepoName = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetRepoNamespace(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetScanTime(v int64) *DescribeCanFixVulListResponseBodyVulRecords {
	s.ScanTime = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetStatus(v int32) *DescribeCanFixVulListResponseBodyVulRecords {
	s.Status = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetTag(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.Tag = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetTargetId(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.TargetId = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetTargetName(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.TargetName = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetTargetType(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.TargetType = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetType(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.Type = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) SetUuid(v string) *DescribeCanFixVulListResponseBodyVulRecords {
	s.Uuid = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecords) Validate() error {
	if s.ExtendContentJson != nil {
		if err := s.ExtendContentJson.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson struct {
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
	// The RPM packages.
	RpmEntityList []*DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList `json:"RpmEntityList,omitempty" xml:"RpmEntityList,omitempty" type:"Repeated"`
}

func (s DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson) String() string {
	return dara.Prettify(s)
}

func (s DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson) GoString() string {
	return s.String()
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson) GetOs() *string {
	return s.Os
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson) GetOsRelease() *string {
	return s.OsRelease
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson) GetRpmEntityList() []*DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	return s.RpmEntityList
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson) SetOs(v string) *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson {
	s.Os = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson) SetOsRelease(v string) *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson {
	s.OsRelease = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson) SetRpmEntityList(v []*DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson {
	s.RpmEntityList = v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJson) Validate() error {
	if s.RpmEntityList != nil {
		for _, item := range s.RpmEntityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList struct {
	// The complete version number of the software package.
	//
	// example:
	//
	// 3.10.0-693.2.2.el7
	FullVersion *string `json:"FullVersion,omitempty" xml:"FullVersion,omitempty"`
	// The SHA-256 value of the digest of the image layer.
	//
	// example:
	//
	// b1f5b9420803ad0657cf21566e3e20acc08581e7f22991249ef3aa80b8b1****
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The information about the detected vulnerability.
	//
	// example:
	//
	// python-perf version less than 0:3.10.0-693.21.1.el7
	MatchDetail *string `json:"MatchDetail,omitempty" xml:"MatchDetail,omitempty"`
	// The rule that is used to detect the vulnerability.
	MatchList []*string `json:"MatchList,omitempty" xml:"MatchList,omitempty" type:"Repeated"`
	// The name of the software package.
	//
	// example:
	//
	// python-perf
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The path of the software that has the vulnerability.
	//
	// example:
	//
	// /usr/lib64/python2.7/site-packages
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The command that is used to fix the vulnerability.
	//
	// example:
	//
	// apt-get update && apt-get install libseccomp2  --only-upgrade
	UpdateCmd *string `json:"UpdateCmd,omitempty" xml:"UpdateCmd,omitempty"`
	// The version number of the software package.
	//
	// example:
	//
	// 3.10.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GoString() string {
	return s.String()
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetFullVersion() *string {
	return s.FullVersion
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetLayer() *string {
	return s.Layer
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetMatchDetail() *string {
	return s.MatchDetail
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetMatchList() []*string {
	return s.MatchList
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetName() *string {
	return s.Name
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetPath() *string {
	return s.Path
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetUpdateCmd() *string {
	return s.UpdateCmd
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetVersion() *string {
	return s.Version
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetFullVersion(v string) *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.FullVersion = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetLayer(v string) *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Layer = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetMatchDetail(v string) *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.MatchDetail = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetMatchList(v []*string) *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.MatchList = v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetName(v string) *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Name = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetPath(v string) *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Path = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetUpdateCmd(v string) *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.UpdateCmd = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetVersion(v string) *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Version = &v
	return s
}

func (s *DescribeCanFixVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) Validate() error {
	return dara.Validate(s)
}
