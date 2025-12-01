// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulListByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVulListByIdResponseBody
	GetCode() *string
	SetData(v []*GetVulListByIdResponseBodyData) *GetVulListByIdResponseBody
	GetData() []*GetVulListByIdResponseBodyData
	SetHttpStatusCode(v int32) *GetVulListByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVulListByIdResponseBody
	GetMessage() *string
	SetPageInfo(v *GetVulListByIdResponseBodyPageInfo) *GetVulListByIdResponseBody
	GetPageInfo() *GetVulListByIdResponseBodyPageInfo
	SetRequestId(v string) *GetVulListByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVulListByIdResponseBody
	GetSuccess() *bool
}

type GetVulListByIdResponseBody struct {
	// API response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data []*GetVulListByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination information.
	PageInfo *GetVulListByIdResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// D38B3D2F-67FD-57FF-87D1-C431D2C70F76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Values: - **true**: Yes. - **false**: No.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVulListByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVulListByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulListByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVulListByIdResponseBody) GetData() []*GetVulListByIdResponseBodyData {
	return s.Data
}

func (s *GetVulListByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVulListByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVulListByIdResponseBody) GetPageInfo() *GetVulListByIdResponseBodyPageInfo {
	return s.PageInfo
}

func (s *GetVulListByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVulListByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVulListByIdResponseBody) SetCode(v string) *GetVulListByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetVulListByIdResponseBody) SetData(v []*GetVulListByIdResponseBodyData) *GetVulListByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetVulListByIdResponseBody) SetHttpStatusCode(v int32) *GetVulListByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVulListByIdResponseBody) SetMessage(v string) *GetVulListByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetVulListByIdResponseBody) SetPageInfo(v *GetVulListByIdResponseBodyPageInfo) *GetVulListByIdResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetVulListByIdResponseBody) SetRequestId(v string) *GetVulListByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulListByIdResponseBody) SetSuccess(v bool) *GetVulListByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetVulListByIdResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVulListByIdResponseBodyData struct {
	// Vulnerability Alias
	//
	// example:
	//
	// Tomcat websocket 拒绝服务漏洞利用代码披露（CVE-2020-13935）
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// Impact description
	EffectMsgDTOS []*GetVulListByIdResponseBodyDataEffectMsgDTOS `json:"EffectMsgDTOS,omitempty" xml:"EffectMsgDTOS,omitempty" type:"Repeated"`
	// Timestamp of the first time the vulnerability was detected
	//
	// example:
	//
	// 1620404763000
	FirstTs *string `json:"FirstTs,omitempty" xml:"FirstTs,omitempty"`
	// Instance name of the asset
	//
	// example:
	//
	// 凌星-CentOS
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Public IP of the asset
	//
	// example:
	//
	// 39.101.73.28
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Private IP of the asset
	//
	// example:
	//
	// 172.22.216.17
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// Timestamp of the last time the vulnerability was detected
	//
	// example:
	//
	// 1620404763000
	LastTs *string `json:"LastTs,omitempty" xml:"LastTs,omitempty"`
	// Vulnerability name
	//
	// example:
	//
	// SCA:ACSV-2020-111301
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Necessity level of vulnerability repair
	//
	// example:
	//
	// later,asap,nntf
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// List of associated CVEs for the vulnerability, separated by commas (,) if there are multiple values.
	//
	// example:
	//
	// CVE-2020-13935
	Related *string `json:"Related,omitempty" xml:"Related,omitempty"`
	// Repair command
	//
	// example:
	//
	// **	- update python-perf
	RepairCmd *string `json:"RepairCmd,omitempty" xml:"RepairCmd,omitempty"`
	// Timestamp of vulnerability repair
	//
	// example:
	//
	// 1541207563000
	RepairTs *string `json:"RepairTs,omitempty" xml:"RepairTs,omitempty"`
	// Vulnerability status:
	//
	// 1: Not fixed
	//
	// 2: Fix failed
	//
	// 3: Rollback failed
	//
	// 4: Fixing
	//
	// 5: Rolling back
	//
	// 6: Verifying
	//
	// 7: Fixed successfully
	//
	// 8: Fixed successfully, pending reboot
	//
	// 9: Rolled back successfully
	//
	// 10: Ignored
	//
	// 11: Rolled back successfully, pending reboot
	//
	// 12: Vulnerability does not exist
	//
	// 20: Expired
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Vulnerability tag
	//
	// example:
	//
	// Restart Required
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// UUID of the asset instance.
	//
	// example:
	//
	// hdm_5cf2eaf263c021b354877943f181956d
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetVulListByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVulListByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVulListByIdResponseBodyData) GetAliasName() *string {
	return s.AliasName
}

func (s *GetVulListByIdResponseBodyData) GetEffectMsgDTOS() []*GetVulListByIdResponseBodyDataEffectMsgDTOS {
	return s.EffectMsgDTOS
}

func (s *GetVulListByIdResponseBodyData) GetFirstTs() *string {
	return s.FirstTs
}

func (s *GetVulListByIdResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetVulListByIdResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *GetVulListByIdResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *GetVulListByIdResponseBodyData) GetLastTs() *string {
	return s.LastTs
}

func (s *GetVulListByIdResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetVulListByIdResponseBodyData) GetNecessity() *string {
	return s.Necessity
}

func (s *GetVulListByIdResponseBodyData) GetRelated() *string {
	return s.Related
}

func (s *GetVulListByIdResponseBodyData) GetRepairCmd() *string {
	return s.RepairCmd
}

func (s *GetVulListByIdResponseBodyData) GetRepairTs() *string {
	return s.RepairTs
}

func (s *GetVulListByIdResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetVulListByIdResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *GetVulListByIdResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *GetVulListByIdResponseBodyData) SetAliasName(v string) *GetVulListByIdResponseBodyData {
	s.AliasName = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetEffectMsgDTOS(v []*GetVulListByIdResponseBodyDataEffectMsgDTOS) *GetVulListByIdResponseBodyData {
	s.EffectMsgDTOS = v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetFirstTs(v string) *GetVulListByIdResponseBodyData {
	s.FirstTs = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetInstanceName(v string) *GetVulListByIdResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetInternetIp(v string) *GetVulListByIdResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetIntranetIp(v string) *GetVulListByIdResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetLastTs(v string) *GetVulListByIdResponseBodyData {
	s.LastTs = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetName(v string) *GetVulListByIdResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetNecessity(v string) *GetVulListByIdResponseBodyData {
	s.Necessity = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetRelated(v string) *GetVulListByIdResponseBodyData {
	s.Related = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetRepairCmd(v string) *GetVulListByIdResponseBodyData {
	s.RepairCmd = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetRepairTs(v string) *GetVulListByIdResponseBodyData {
	s.RepairTs = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetStatus(v string) *GetVulListByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetTag(v string) *GetVulListByIdResponseBodyData {
	s.Tag = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetUuid(v string) *GetVulListByIdResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) Validate() error {
	if s.EffectMsgDTOS != nil {
		for _, item := range s.EffectMsgDTOS {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVulListByIdResponseBodyDataEffectMsgDTOS struct {
	// Hit
	//
	// example:
	//
	// fastjson(jar) extendField.safemode equals false
	MatchList *string `json:"MatchList,omitempty" xml:"MatchList,omitempty"`
	// Path
	//
	// example:
	//
	// /uat6/qry/enquiry/policy/yrtPolicyList
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Software name
	//
	// example:
	//
	// python-perf 3.10.0
	SoftName *string `json:"SoftName,omitempty" xml:"SoftName,omitempty"`
}

func (s GetVulListByIdResponseBodyDataEffectMsgDTOS) String() string {
	return dara.Prettify(s)
}

func (s GetVulListByIdResponseBodyDataEffectMsgDTOS) GoString() string {
	return s.String()
}

func (s *GetVulListByIdResponseBodyDataEffectMsgDTOS) GetMatchList() *string {
	return s.MatchList
}

func (s *GetVulListByIdResponseBodyDataEffectMsgDTOS) GetPath() *string {
	return s.Path
}

func (s *GetVulListByIdResponseBodyDataEffectMsgDTOS) GetSoftName() *string {
	return s.SoftName
}

func (s *GetVulListByIdResponseBodyDataEffectMsgDTOS) SetMatchList(v string) *GetVulListByIdResponseBodyDataEffectMsgDTOS {
	s.MatchList = &v
	return s
}

func (s *GetVulListByIdResponseBodyDataEffectMsgDTOS) SetPath(v string) *GetVulListByIdResponseBodyDataEffectMsgDTOS {
	s.Path = &v
	return s
}

func (s *GetVulListByIdResponseBodyDataEffectMsgDTOS) SetSoftName(v string) *GetVulListByIdResponseBodyDataEffectMsgDTOS {
	s.SoftName = &v
	return s
}

func (s *GetVulListByIdResponseBodyDataEffectMsgDTOS) Validate() error {
	return dara.Validate(s)
}

type GetVulListByIdResponseBodyPageInfo struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items per page in the returned data.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of records in the query result.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetVulListByIdResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVulListByIdResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetVulListByIdResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetVulListByIdResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetVulListByIdResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetVulListByIdResponseBodyPageInfo) SetCurrentPage(v int32) *GetVulListByIdResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetVulListByIdResponseBodyPageInfo) SetPageSize(v int32) *GetVulListByIdResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetVulListByIdResponseBodyPageInfo) SetTotalCount(v int32) *GetVulListByIdResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *GetVulListByIdResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
