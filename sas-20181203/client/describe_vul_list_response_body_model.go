// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeVulListResponseBody
	GetCurrentPage() *int32
	SetNextToken(v string) *DescribeVulListResponseBody
	GetNextToken() *string
	SetPageSize(v int32) *DescribeVulListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVulListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVulListResponseBody
	GetTotalCount() *int32
	SetVulRecords(v []*DescribeVulListResponseBodyVulRecords) *DescribeVulListResponseBody
	GetVulRecords() []*DescribeVulListResponseBodyVulRecords
}

type DescribeVulListResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The value of NextToken that is returned when the NextToken method is used.
	//
	// example:
	//
	// E17B501887A2D3AA5E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2F26AB2A-1075-488F-8472-40E5DB486ACC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of vulnerabilities returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the vulnerability.
	VulRecords []*DescribeVulListResponseBodyVulRecords `json:"VulRecords,omitempty" xml:"VulRecords,omitempty" type:"Repeated"`
}

func (s DescribeVulListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVulListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVulListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVulListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVulListResponseBody) GetVulRecords() []*DescribeVulListResponseBodyVulRecords {
	return s.VulRecords
}

func (s *DescribeVulListResponseBody) SetCurrentPage(v int32) *DescribeVulListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulListResponseBody) SetNextToken(v string) *DescribeVulListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVulListResponseBody) SetPageSize(v int32) *DescribeVulListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVulListResponseBody) SetRequestId(v string) *DescribeVulListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulListResponseBody) SetTotalCount(v int32) *DescribeVulListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulListResponseBody) SetVulRecords(v []*DescribeVulListResponseBodyVulRecords) *DescribeVulListResponseBody {
	s.VulRecords = v
	return s
}

func (s *DescribeVulListResponseBody) Validate() error {
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

type DescribeVulListResponseBodyVulRecords struct {
	// The name of the vulnerability.
	//
	// example:
	//
	// RHSA-2019:0230-Important: polkit security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The edition of Security Center that is authorized to scan the asset. Valid values:
	//
	// 	- **1**: Basic.
	//
	// 	- **6**: Anti-virus.
	//
	// 	- **5**: Advanced.
	//
	// 	- **3**: Enterprise.
	//
	// 	- **7**: Ultimate.
	//
	// 	- **10**: Value-added Plan.
	//
	// example:
	//
	// 3
	AuthVersion *string `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// Indicates whether Security Center is authorized to scan the asset. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Bind *bool `json:"Bind,omitempty" xml:"Bind,omitempty"`
	// The extended information about the vulnerability.
	ExtendContentJson *DescribeVulListResponseBodyVulRecordsExtendContentJson `json:"ExtendContentJson,omitempty" xml:"ExtendContentJson,omitempty" type:"Struct"`
	// The timestamp when the vulnerability was first detected. Unit: milliseconds.
	//
	// example:
	//
	// 1554189334000
	FirstTs *int64 `json:"FirstTs,omitempty" xml:"FirstTs,omitempty"`
	// The ID of the asset group.
	//
	// example:
	//
	// 281801
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// registry-cn-**-vpc.ack.**.com/acs/ack-node-problem-detector:v0.8.16-8ed7053-**
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The ID of the asset.
	//
	// example:
	//
	// i-bp18t***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the asset.
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
	// 1.2.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The ID of the cluster.
	//
	// >  The value of this parameter is returned only if you use the Ultimate edition of Security Center that can protect container assets.
	//
	// example:
	//
	// c863dc93bed3843de9934d4346dc4****
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The namespace.
	//
	// >  If you use the Ultimate edition of Security Center, the value of this parameter is queried from container assets. If you do not use the Ultimate edition, the value of this parameter is queried from the Security Center agent.
	//
	// example:
	//
	// default
	K8sNamespace *string `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	// The ID of the node.
	//
	// >  The value of this parameter is returned only if you use the Ultimate edition of Security Center that can protect container assets.
	//
	// example:
	//
	// i-bp1ifm6suw9mnbsr****
	K8sNodeId *string `json:"K8sNodeId,omitempty" xml:"K8sNodeId,omitempty"`
	// The name of the node.
	//
	// >  The value of this parameter is returned only if you use the Ultimate edition of Security Center that can protect container assets.
	//
	// example:
	//
	// deployment-riskai-7b67d68975-m****
	K8sNodeName *string `json:"K8sNodeName,omitempty" xml:"K8sNodeName,omitempty"`
	// The name of the pod.
	//
	// >  The value of this parameter is returned only if you use the Ultimate edition of Security Center that can protect container assets.
	//
	// example:
	//
	// deployment-riskai-7b67d68975-m****
	K8sPodName *string `json:"K8sPodName,omitempty" xml:"K8sPodName,omitempty"`
	// The timestamp when the vulnerability was last detected. Unit: milliseconds.
	//
	// example:
	//
	// 1620404763000
	LastTs *int64 `json:"LastTs,omitempty" xml:"LastTs,omitempty"`
	// The timestamp when the vulnerability status was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1620404763000
	ModifyTs *int64 `json:"ModifyTs,omitempty" xml:"ModifyTs,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// oval:com.redhat.rhsa:def:20170574
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// kube-system
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The priority to fix the vulnerability. Valid values:
	//
	// 	- **asap**: high.
	//
	// 	- **later**: medium.
	//
	// 	- **nntf**: low.
	//
	// >  We recommend that you fix **high-risk*	- vulnerabilities at the earliest opportunity.
	//
	// example:
	//
	// asap
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// Indicates whether the Security Center agent on the asset is online. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The name of the operating system for your asset.
	//
	// example:
	//
	// CentOS 7.2 64-bit
	OsName *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	// The name of the operating system for your asset.
	//
	// example:
	//
	// linux
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// The ID of the vulnerability.
	//
	// example:
	//
	// 101162078
	PrimaryId *int64 `json:"PrimaryId,omitempty" xml:"PrimaryId,omitempty"`
	// The progress of the vulnerability fixing.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Indicates whether the application protection feature is supported. Valid values:
	//
	// 	- **0**: no.
	//
	// 	- **1**: yes.
	//
	// >  If this parameter is not returned, the application protection feature is not supported.
	//
	// example:
	//
	// 1
	RaspDefend *int32 `json:"RaspDefend,omitempty" xml:"RaspDefend,omitempty"`
	// The protection mode of the application protection feature. Valid values:
	//
	// 	- **0**: unprotected.
	//
	// 	- **1**: the Monitor mode.
	//
	// 	- **2**: the Block mode.
	//
	// 	- **3**: disabled.
	//
	// example:
	//
	// 1
	RaspStatus *int32 `json:"RaspStatus,omitempty" xml:"RaspStatus,omitempty"`
	// Indicates whether the vulnerability is easily exploited. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	RealRisk *bool `json:"RealRisk,omitempty" xml:"RealRisk,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Common Vulnerabilities and Exposures (CVE) IDs related to the vulnerability. Multiple CVE IDs are separated by commas (,).
	//
	// example:
	//
	// CVE-2017-7518,CVE-2017-12188
	Related *string `json:"Related,omitempty" xml:"Related,omitempty"`
	// The timestamp when the vulnerability was fixed. Unit: milliseconds. This parameter is returned only if you fix vulnerabilities in the Security Center console.
	//
	// example:
	//
	// 1541207563000
	RepairTs *int64 `json:"RepairTs,omitempty" xml:"RepairTs,omitempty"`
	// The code that indicates the vulnerability fixing result.
	//
	// example:
	//
	// 0
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// The message that indicates the vulnerability fixing result.
	//
	// example:
	//
	// timeout
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// The tag of this vulnerability. Valid values:
	//
	// 	- **AI**: AI-related components.
	//
	// example:
	//
	// AI
	RuleTag *string `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
	// The status of the vulnerability. Valid values:
	//
	// 	- **1**: unfixed.
	//
	// 	- **2**: fix failed.
	//
	// 	- **3**: rollback failed.
	//
	// 	- **4**: being fixed.
	//
	// 	- **5**: being rolled back.
	//
	// 	- **6**: being verified.
	//
	// 	- **7**: fixed.
	//
	// 	- **8**: fixed and to be restarted.
	//
	// 	- **9**: rolled back.
	//
	// 	- **10**: ignored.
	//
	// 	- **11**: rolled back and to be restarted.
	//
	// 	- **12**: not found.
	//
	// 	- **20**: expired.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag that is added to the vulnerability.
	//
	// example:
	//
	// oval
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability.
	//
	// 	- **sys**: Windows system vulnerability.
	//
	// 	- **cms**: Web-CMS vulnerability.
	//
	// 	- **emg**: urgent vulnerability.
	//
	// 	- **app**: application vulnerability.
	//
	// 	- **sca**: application vulnerability that is detected by using software component analysis.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the asset.
	//
	// example:
	//
	// 04c56617-23fc-43a5-ab9b-****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeVulListResponseBodyVulRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecords) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecords) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeVulListResponseBodyVulRecords) GetAuthVersion() *string {
	return s.AuthVersion
}

func (s *DescribeVulListResponseBodyVulRecords) GetBind() *bool {
	return s.Bind
}

func (s *DescribeVulListResponseBodyVulRecords) GetExtendContentJson() *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	return s.ExtendContentJson
}

func (s *DescribeVulListResponseBodyVulRecords) GetFirstTs() *int64 {
	return s.FirstTs
}

func (s *DescribeVulListResponseBodyVulRecords) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeVulListResponseBodyVulRecords) GetImage() *string {
	return s.Image
}

func (s *DescribeVulListResponseBodyVulRecords) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVulListResponseBodyVulRecords) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeVulListResponseBodyVulRecords) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeVulListResponseBodyVulRecords) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeVulListResponseBodyVulRecords) GetK8sClusterId() *string {
	return s.K8sClusterId
}

func (s *DescribeVulListResponseBodyVulRecords) GetK8sNamespace() *string {
	return s.K8sNamespace
}

func (s *DescribeVulListResponseBodyVulRecords) GetK8sNodeId() *string {
	return s.K8sNodeId
}

func (s *DescribeVulListResponseBodyVulRecords) GetK8sNodeName() *string {
	return s.K8sNodeName
}

func (s *DescribeVulListResponseBodyVulRecords) GetK8sPodName() *string {
	return s.K8sPodName
}

func (s *DescribeVulListResponseBodyVulRecords) GetLastTs() *int64 {
	return s.LastTs
}

func (s *DescribeVulListResponseBodyVulRecords) GetModifyTs() *int64 {
	return s.ModifyTs
}

func (s *DescribeVulListResponseBodyVulRecords) GetName() *string {
	return s.Name
}

func (s *DescribeVulListResponseBodyVulRecords) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeVulListResponseBodyVulRecords) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeVulListResponseBodyVulRecords) GetOnline() *bool {
	return s.Online
}

func (s *DescribeVulListResponseBodyVulRecords) GetOsName() *string {
	return s.OsName
}

func (s *DescribeVulListResponseBodyVulRecords) GetOsVersion() *string {
	return s.OsVersion
}

func (s *DescribeVulListResponseBodyVulRecords) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *DescribeVulListResponseBodyVulRecords) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeVulListResponseBodyVulRecords) GetRaspDefend() *int32 {
	return s.RaspDefend
}

func (s *DescribeVulListResponseBodyVulRecords) GetRaspStatus() *int32 {
	return s.RaspStatus
}

func (s *DescribeVulListResponseBodyVulRecords) GetRealRisk() *bool {
	return s.RealRisk
}

func (s *DescribeVulListResponseBodyVulRecords) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVulListResponseBodyVulRecords) GetRelated() *string {
	return s.Related
}

func (s *DescribeVulListResponseBodyVulRecords) GetRepairTs() *int64 {
	return s.RepairTs
}

func (s *DescribeVulListResponseBodyVulRecords) GetResultCode() *string {
	return s.ResultCode
}

func (s *DescribeVulListResponseBodyVulRecords) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DescribeVulListResponseBodyVulRecords) GetRuleTag() *string {
	return s.RuleTag
}

func (s *DescribeVulListResponseBodyVulRecords) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeVulListResponseBodyVulRecords) GetTag() *string {
	return s.Tag
}

func (s *DescribeVulListResponseBodyVulRecords) GetType() *string {
	return s.Type
}

func (s *DescribeVulListResponseBodyVulRecords) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeVulListResponseBodyVulRecords) SetAliasName(v string) *DescribeVulListResponseBodyVulRecords {
	s.AliasName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetAuthVersion(v string) *DescribeVulListResponseBodyVulRecords {
	s.AuthVersion = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetBind(v bool) *DescribeVulListResponseBodyVulRecords {
	s.Bind = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetExtendContentJson(v *DescribeVulListResponseBodyVulRecordsExtendContentJson) *DescribeVulListResponseBodyVulRecords {
	s.ExtendContentJson = v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetFirstTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.FirstTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetGroupId(v int64) *DescribeVulListResponseBodyVulRecords {
	s.GroupId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetImage(v string) *DescribeVulListResponseBodyVulRecords {
	s.Image = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetInstanceId(v string) *DescribeVulListResponseBodyVulRecords {
	s.InstanceId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetInstanceName(v string) *DescribeVulListResponseBodyVulRecords {
	s.InstanceName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetInternetIp(v string) *DescribeVulListResponseBodyVulRecords {
	s.InternetIp = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetIntranetIp(v string) *DescribeVulListResponseBodyVulRecords {
	s.IntranetIp = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetK8sClusterId(v string) *DescribeVulListResponseBodyVulRecords {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetK8sNamespace(v string) *DescribeVulListResponseBodyVulRecords {
	s.K8sNamespace = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetK8sNodeId(v string) *DescribeVulListResponseBodyVulRecords {
	s.K8sNodeId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetK8sNodeName(v string) *DescribeVulListResponseBodyVulRecords {
	s.K8sNodeName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetK8sPodName(v string) *DescribeVulListResponseBodyVulRecords {
	s.K8sPodName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetLastTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.LastTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetModifyTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.ModifyTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetName(v string) *DescribeVulListResponseBodyVulRecords {
	s.Name = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetNamespace(v string) *DescribeVulListResponseBodyVulRecords {
	s.Namespace = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetNecessity(v string) *DescribeVulListResponseBodyVulRecords {
	s.Necessity = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetOnline(v bool) *DescribeVulListResponseBodyVulRecords {
	s.Online = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetOsName(v string) *DescribeVulListResponseBodyVulRecords {
	s.OsName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetOsVersion(v string) *DescribeVulListResponseBodyVulRecords {
	s.OsVersion = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetPrimaryId(v int64) *DescribeVulListResponseBodyVulRecords {
	s.PrimaryId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetProgress(v int32) *DescribeVulListResponseBodyVulRecords {
	s.Progress = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRaspDefend(v int32) *DescribeVulListResponseBodyVulRecords {
	s.RaspDefend = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRaspStatus(v int32) *DescribeVulListResponseBodyVulRecords {
	s.RaspStatus = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRealRisk(v bool) *DescribeVulListResponseBodyVulRecords {
	s.RealRisk = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRegionId(v string) *DescribeVulListResponseBodyVulRecords {
	s.RegionId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRelated(v string) *DescribeVulListResponseBodyVulRecords {
	s.Related = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRepairTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.RepairTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetResultCode(v string) *DescribeVulListResponseBodyVulRecords {
	s.ResultCode = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetResultMessage(v string) *DescribeVulListResponseBodyVulRecords {
	s.ResultMessage = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRuleTag(v string) *DescribeVulListResponseBodyVulRecords {
	s.RuleTag = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetStatus(v int32) *DescribeVulListResponseBodyVulRecords {
	s.Status = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetTag(v string) *DescribeVulListResponseBodyVulRecords {
	s.Tag = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetType(v string) *DescribeVulListResponseBodyVulRecords {
	s.Type = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetUuid(v string) *DescribeVulListResponseBodyVulRecords {
	s.Uuid = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) Validate() error {
	if s.ExtendContentJson != nil {
		if err := s.ExtendContentJson.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVulListResponseBodyVulRecordsExtendContentJson struct {
	// The path to the package of the software that has the vulnerability.
	//
	// example:
	//
	// /roo/www/web
	AbsolutePath *string `json:"AbsolutePath,omitempty" xml:"AbsolutePath,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// RHSA-2019:0230-Important: polkit security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The description of the vulnerability.
	//
	// example:
	//
	// kernel version:5.10.84-10.2.al8.x86_64
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The returned message that indicates the urgent vulnerability.
	//
	// example:
	//
	// com.xxl.rpc.util.XxlRpcException: xxl-rpc request data is empty.\\n\\tat com.xxl.rpc.remoting.net.impl.servlet.serve"
	EmgProof *string `json:"EmgProof,omitempty" xml:"EmgProof,omitempty"`
	// The public IP address of the asset that is associated with the vulnerability.
	//
	// example:
	//
	// 1.2.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The timestamp when the vulnerability was last detected. Unit: milliseconds.
	//
	// example:
	//
	// 1620404763000
	LastTs *int64 `json:"LastTs,omitempty" xml:"LastTs,omitempty"`
	// Indicates whether the vulnerability needs to be fixed.
	Necessity *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity `json:"Necessity,omitempty" xml:"Necessity,omitempty" type:"Struct"`
	// The name of the operating system.
	//
	// example:
	//
	// centos
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The information about the operating system version.
	//
	// example:
	//
	// 7
	OsRelease *string `json:"OsRelease,omitempty" xml:"OsRelease,omitempty"`
	// The ID of the vulnerability.
	//
	// example:
	//
	// 111
	PrimaryId *int64 `json:"PrimaryId,omitempty" xml:"PrimaryId,omitempty"`
	// The information about RPM Package Manager (RPM) packages.
	RpmEntityList []*DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList `json:"RpmEntityList,omitempty" xml:"RpmEntityList,omitempty" type:"Repeated"`
	// The status of the vulnerability. Valid values:
	//
	// 	- **1**: unfixed.
	//
	// 	- **2**: fix failed.
	//
	// 	- 3: rollback failed.
	//
	// 	- **4**: being fixed.
	//
	// 	- **5**: being rolled back.
	//
	// 	- **6**: being verified.
	//
	// 	- **7**: fixed.
	//
	// 	- **8**: fixed and to be restarted.
	//
	// 	- **9**: rolled back.
	//
	// 	- **10**: ignored.
	//
	// 	- **11**: rolled back and to be restarted.
	//
	// 	- **12**: not found.
	//
	// 	- **20**: expired.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag that is added to the vulnerability.
	//
	// example:
	//
	// oval
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The URL of the vulnerability.
	//
	// example:
	//
	// http://39.99.XX.XX:30005/toLogin
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The CVE list.
	CveList []*string `json:"cveList,omitempty" xml:"cveList,omitempty" type:"Repeated"`
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJson) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJson) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetAbsolutePath() *string {
	return s.AbsolutePath
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetDescription() *string {
	return s.Description
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetEmgProof() *string {
	return s.EmgProof
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetIp() *string {
	return s.Ip
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetLastTs() *int64 {
	return s.LastTs
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetNecessity() *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	return s.Necessity
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetOs() *string {
	return s.Os
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetOsRelease() *string {
	return s.OsRelease
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetRpmEntityList() []*DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	return s.RpmEntityList
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetStatus() *string {
	return s.Status
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetTag() *string {
	return s.Tag
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetTarget() *string {
	return s.Target
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) GetCveList() []*string {
	return s.CveList
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetAbsolutePath(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.AbsolutePath = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetAliasName(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.AliasName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetDescription(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.Description = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetEmgProof(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.EmgProof = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetIp(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.Ip = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetLastTs(v int64) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.LastTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetNecessity(v *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.Necessity = v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetOs(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.Os = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetOsRelease(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.OsRelease = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetPrimaryId(v int64) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.PrimaryId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetRpmEntityList(v []*DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.RpmEntityList = v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetStatus(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.Status = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetTag(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.Tag = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetTarget(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.Target = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetCveList(v []*string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.CveList = v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) Validate() error {
	if s.Necessity != nil {
		if err := s.Necessity.Validate(); err != nil {
			return err
		}
	}
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

type DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity struct {
	// The asset importance score. Valid values:
	//
	// 	- **2**: important asset.
	//
	// 	- **1**: common asset.
	//
	// 	- **0**: test asset.
	//
	// example:
	//
	// 1
	AssetsFactor *string `json:"Assets_factor,omitempty" xml:"Assets_factor,omitempty"`
	// The Common Vulnerability Scoring System (CVSS) score.
	//
	// example:
	//
	// 7.8
	CvssFactor *string `json:"Cvss_factor,omitempty" xml:"Cvss_factor,omitempty"`
	// The environment score.
	//
	// example:
	//
	// 1.0
	EnviromentFactor *string `json:"Enviroment_factor,omitempty" xml:"Enviroment_factor,omitempty"`
	// Indicates whether the vulnerability priority score is calculated. Valid values:
	//
	// 	- **0**: no.
	//
	// 	- **1**: yes.
	//
	// example:
	//
	// 1
	IsCalc *string `json:"Is_calc,omitempty" xml:"Is_calc,omitempty"`
	// The status of the vulnerability priority score. Valid values:
	//
	// 	- **none**: No score is generated.
	//
	// 	- **pending**: The score is pending calculation.
	//
	// 	- **normal**: The calculation is normal.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time score.
	//
	// example:
	//
	// 1.0
	TimeFactor *string `json:"Time_factor,omitempty" xml:"Time_factor,omitempty"`
	// The vulnerability priority score.
	//
	// The following list describes scores and related fixing suggestions:
	//
	// 	- If the score is from **13.5 to 15**, the vulnerability is a high-risk vulnerability. You must fix the vulnerability at the earliest opportunity.
	//
	// 	- If the score is **greater than or equal to 7 but less than 13.5**, the vulnerability is a medium-risk vulnerability. You can fix the vulnerability at your convenience.
	//
	// 	- If the score is **less than 7**, the vulnerability is a low-risk vulnerability. You can ignore the vulnerability.
	//
	// example:
	//
	// 7.8
	TotalScore *string `json:"Total_score,omitempty" xml:"Total_score,omitempty"`
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) GetAssetsFactor() *string {
	return s.AssetsFactor
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) GetCvssFactor() *string {
	return s.CvssFactor
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) GetEnviromentFactor() *string {
	return s.EnviromentFactor
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) GetIsCalc() *string {
	return s.IsCalc
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) GetStatus() *string {
	return s.Status
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) GetTimeFactor() *string {
	return s.TimeFactor
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) GetTotalScore() *string {
	return s.TotalScore
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetAssetsFactor(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.AssetsFactor = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetCvssFactor(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.CvssFactor = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetEnviromentFactor(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.EnviromentFactor = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetIsCalc(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.IsCalc = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetStatus(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.Status = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetTimeFactor(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.TimeFactor = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetTotalScore(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.TotalScore = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) Validate() error {
	return dara.Validate(s)
}

type DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList struct {
	// The name of the container.
	//
	// example:
	//
	// k8s_67895c4_xxx
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The extended information about the software package that has the vulnerability.
	//
	// example:
	//
	// {"msg_no_lookups_configured_CVE_2021_44228": "false", "jndi_class_not_exist": "false"}
	ExtendField *string `json:"ExtendField,omitempty" xml:"ExtendField,omitempty"`
	// The complete version number.
	//
	// example:
	//
	// 3.10.0-693.2.2.el7
	FullVersion *string `json:"FullVersion,omitempty" xml:"FullVersion,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// registry_387ytb_xxx
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The reason why the vulnerability is detected.
	//
	// example:
	//
	// python-perf version less than 0:3.10.0-693.21.1.el7
	MatchDetail *string `json:"MatchDetail,omitempty" xml:"MatchDetail,omitempty"`
	// The rules that are used to detect the vulnerability.
	MatchList []*string `json:"MatchList,omitempty" xml:"MatchList,omitempty" type:"Repeated"`
	// The name of the RPM package.
	//
	// example:
	//
	// python-perf
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The path to the software that has the vulnerability.
	//
	// example:
	//
	// /usr/lib64/python2.7/site-packages
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The process ID.
	//
	// example:
	//
	// 8664
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The command that is used to fix the vulnerability.
	//
	// example:
	//
	// **	- update python-perf
	UpdateCmd *string `json:"UpdateCmd,omitempty" xml:"UpdateCmd,omitempty"`
	// The version number of the package of the software that has the vulnerability.
	//
	// example:
	//
	// 3.10.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetContainerName() *string {
	return s.ContainerName
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetExtendField() *string {
	return s.ExtendField
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetFullVersion() *string {
	return s.FullVersion
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetMatchDetail() *string {
	return s.MatchDetail
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetMatchList() []*string {
	return s.MatchList
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetName() *string {
	return s.Name
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetPath() *string {
	return s.Path
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetPid() *string {
	return s.Pid
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetUpdateCmd() *string {
	return s.UpdateCmd
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GetVersion() *string {
	return s.Version
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetContainerName(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.ContainerName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetExtendField(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.ExtendField = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetFullVersion(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.FullVersion = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetImageName(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.ImageName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetMatchDetail(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.MatchDetail = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetMatchList(v []*string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.MatchList = v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetName(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Name = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetPath(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Path = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetPid(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Pid = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetUpdateCmd(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.UpdateCmd = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetVersion(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Version = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) Validate() error {
	return dara.Validate(s)
}
