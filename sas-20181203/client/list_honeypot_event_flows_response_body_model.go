// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotEventFlowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHoneypotEventFlowsResponseBody
	GetCode() *string
	SetHoneypotEventFlows(v []*ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) *ListHoneypotEventFlowsResponseBody
	GetHoneypotEventFlows() []*ListHoneypotEventFlowsResponseBodyHoneypotEventFlows
	SetHttpStatusCode(v int32) *ListHoneypotEventFlowsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListHoneypotEventFlowsResponseBody
	GetMessage() *string
	SetPageInfo(v *ListHoneypotEventFlowsResponseBodyPageInfo) *ListHoneypotEventFlowsResponseBody
	GetPageInfo() *ListHoneypotEventFlowsResponseBodyPageInfo
	SetRequestId(v string) *ListHoneypotEventFlowsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHoneypotEventFlowsResponseBody
	GetSuccess() *bool
}

type ListHoneypotEventFlowsResponseBody struct {
	// The status code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The attack timelines.
	HoneypotEventFlows []*ListHoneypotEventFlowsResponseBodyHoneypotEventFlows `json:"HoneypotEventFlows,omitempty" xml:"HoneypotEventFlows,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListHoneypotEventFlowsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9F4E6157-9600-5588-86B9-38F09067****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHoneypotEventFlowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotEventFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHoneypotEventFlowsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHoneypotEventFlowsResponseBody) GetHoneypotEventFlows() []*ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	return s.HoneypotEventFlows
}

func (s *ListHoneypotEventFlowsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHoneypotEventFlowsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHoneypotEventFlowsResponseBody) GetPageInfo() *ListHoneypotEventFlowsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListHoneypotEventFlowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHoneypotEventFlowsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHoneypotEventFlowsResponseBody) SetCode(v string) *ListHoneypotEventFlowsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBody) SetHoneypotEventFlows(v []*ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) *ListHoneypotEventFlowsResponseBody {
	s.HoneypotEventFlows = v
	return s
}

func (s *ListHoneypotEventFlowsResponseBody) SetHttpStatusCode(v int32) *ListHoneypotEventFlowsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBody) SetMessage(v string) *ListHoneypotEventFlowsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBody) SetPageInfo(v *ListHoneypotEventFlowsResponseBodyPageInfo) *ListHoneypotEventFlowsResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListHoneypotEventFlowsResponseBody) SetRequestId(v string) *ListHoneypotEventFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBody) SetSuccess(v bool) *ListHoneypotEventFlowsResponseBody {
	s.Success = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBody) Validate() error {
	if s.HoneypotEventFlows != nil {
		for _, item := range s.HoneypotEventFlows {
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

type ListHoneypotEventFlowsResponseBodyHoneypotEventFlows struct {
	// The ID of the probe.
	//
	// example:
	//
	// d3c0dafa-5059-4eb0-8c28-7d40f58*****
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The name of the probe.
	//
	// example:
	//
	// hw-d***
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The ID of the container.
	//
	// example:
	//
	// eca09895****
	DockerId *string `json:"DockerId,omitempty" xml:"DockerId,omitempty"`
	// The destination IP address.
	//
	// example:
	//
	// 112.126.205.***
	DstIp *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 80
	DstPort *int32 `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The UUID of the connection in the attack.
	//
	// example:
	//
	// fd7f1ff4-0c4b-41cb-99ad-0724349d****
	EventConnection *string `json:"EventConnection,omitempty" xml:"EventConnection,omitempty"`
	// The extended information about the attack payload.
	//
	// example:
	//
	// {\\"payload\\":{\\"format\\":\\"line\\",\\"name\\":{\\"cn\\":\\"payload\\",\\"en\\":\\"payload\\"},\\"value\\":\\"\\"},\\"uid\\":{\\"format\\":\\"line\\",\\"name\\":{\\"cn\\":\\"\\",\\"en\\":\\"\\"},\\"uid\\":\\"5fa2ece9-aa08-4bbd-a272-5d27*********\\",\\"value\\":\\"\\"}}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The extension information about the virtual private cloud (VPC).
	//
	// example:
	//
	// {\\"vpc_id\\":\\"\\",\\"vpc_dest_port\\":\\"\\",\\"vpc_dest_ip\\":\\"\\"}
	Extra1 *string `json:"Extra1,omitempty" xml:"Extra1,omitempty"`
	// The Object Storage Service (OSS) URL of the file.
	//
	// example:
	//
	// https://pop-test-file-upload.oss-cn-beijing.aliyuncs.com/5626_26331*****
	FileOssUrl *string `json:"FileOssUrl,omitempty" xml:"FileOssUrl,omitempty"`
	// The timestamp when the intrusion event was first occurred.
	//
	// example:
	//
	// 1686621122000
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// The ID of the intrusion event. The value is a string.
	//
	// example:
	//
	// 19bec028-d98b-45c4-a4d9-cc3d593f****
	HoneypotEventId *string `json:"HoneypotEventId,omitempty" xml:"HoneypotEventId,omitempty"`
	// The ID of the honeypot.
	//
	// example:
	//
	// 911df9d6fe20451c059edbcffa1d1c33452f6a71e59d4826da067af224*****
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
	// The name of the honeypot.
	//
	// example:
	//
	// hw-zhi*****
	HoneypotName *string `json:"HoneypotName,omitempty" xml:"HoneypotName,omitempty"`
	// The timestamp when the intrusion event was last occurred.
	//
	// example:
	//
	// 1686622222000
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **2**: low
	//
	// 	- **3**: medium
	//
	// 	- **4**: high
	//
	// example:
	//
	// 4
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The ID of the intrusion event.
	//
	// example:
	//
	// 306527555
	SecurityEventId *int64 `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 121.41.48.***
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// The source media access control (MAC) address.
	//
	// example:
	//
	// 00:0C:29:CA:**:**
	SrcMac *string `json:"SrcMac,omitempty" xml:"SrcMac,omitempty"`
	// The source port number.
	//
	// example:
	//
	// 80
	SrcPort *int32 `json:"SrcPort,omitempty" xml:"SrcPort,omitempty"`
	// The handling status of the intrusion event. Valid values:
	//
	// 	- **1**: pending handling
	//
	// 	- **2**: ignored
	//
	// 	- **4**: confirmed
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the attack type.
	//
	// example:
	//
	// web_access
	TypeId *string `json:"TypeId,omitempty" xml:"TypeId,omitempty"`
	// The UUID of an attack in the intrusion event.
	//
	// example:
	//
	// 5fa2ece9-aa08-4bbd-a272-5d27d1c6*****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GoString() string {
	return s.String()
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetAgentId() *string {
	return s.AgentId
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetAgentName() *string {
	return s.AgentName
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetDockerId() *string {
	return s.DockerId
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetDstIp() *string {
	return s.DstIp
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetDstPort() *int32 {
	return s.DstPort
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetEventConnection() *string {
	return s.EventConnection
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetExtra() *string {
	return s.Extra
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetExtra1() *string {
	return s.Extra1
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetFileOssUrl() *string {
	return s.FileOssUrl
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetHoneypotEventId() *string {
	return s.HoneypotEventId
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetHoneypotName() *string {
	return s.HoneypotName
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetLastTime() *int64 {
	return s.LastTime
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetSecurityEventId() *int64 {
	return s.SecurityEventId
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetSrcIp() *string {
	return s.SrcIp
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetSrcMac() *string {
	return s.SrcMac
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetSrcPort() *int32 {
	return s.SrcPort
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetStatus() *int32 {
	return s.Status
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetTypeId() *string {
	return s.TypeId
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) GetUid() *string {
	return s.Uid
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetAgentId(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.AgentId = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetAgentName(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.AgentName = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetDockerId(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.DockerId = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetDstIp(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.DstIp = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetDstPort(v int32) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.DstPort = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetEventConnection(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.EventConnection = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetExtra(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.Extra = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetExtra1(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.Extra1 = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetFileOssUrl(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.FileOssUrl = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetFirstTime(v int64) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.FirstTime = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetHoneypotEventId(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.HoneypotEventId = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetHoneypotId(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.HoneypotId = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetHoneypotName(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.HoneypotName = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetLastTime(v int64) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.LastTime = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetRiskLevel(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.RiskLevel = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetSecurityEventId(v int64) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.SecurityEventId = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetSrcIp(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.SrcIp = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetSrcMac(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.SrcMac = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetSrcPort(v int32) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.SrcPort = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetStatus(v int32) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.Status = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetTypeId(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.TypeId = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) SetUid(v string) *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows {
	s.Uid = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyHoneypotEventFlows) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotEventFlowsResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 78
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHoneypotEventFlowsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotEventFlowsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListHoneypotEventFlowsResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListHoneypotEventFlowsResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotEventFlowsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotEventFlowsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHoneypotEventFlowsResponseBodyPageInfo) SetCount(v int32) *ListHoneypotEventFlowsResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyPageInfo) SetCurrentPage(v int32) *ListHoneypotEventFlowsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyPageInfo) SetPageSize(v int32) *ListHoneypotEventFlowsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyPageInfo) SetTotalCount(v int32) *ListHoneypotEventFlowsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListHoneypotEventFlowsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
