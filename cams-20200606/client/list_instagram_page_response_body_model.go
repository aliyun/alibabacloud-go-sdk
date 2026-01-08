// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstagramPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListInstagramPageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListInstagramPageResponseBody
	GetCode() *string
	SetData(v *ListInstagramPageResponseBodyData) *ListInstagramPageResponseBody
	GetData() *ListInstagramPageResponseBodyData
	SetMessage(v string) *ListInstagramPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstagramPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstagramPageResponseBody
	GetSuccess() *bool
}

type ListInstagramPageResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListInstagramPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// dsfdsf-3jfj***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstagramPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstagramPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstagramPageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListInstagramPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstagramPageResponseBody) GetData() *ListInstagramPageResponseBodyData {
	return s.Data
}

func (s *ListInstagramPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstagramPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstagramPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstagramPageResponseBody) SetAccessDeniedDetail(v string) *ListInstagramPageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListInstagramPageResponseBody) SetCode(v string) *ListInstagramPageResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstagramPageResponseBody) SetData(v *ListInstagramPageResponseBodyData) *ListInstagramPageResponseBody {
	s.Data = v
	return s
}

func (s *ListInstagramPageResponseBody) SetMessage(v string) *ListInstagramPageResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstagramPageResponseBody) SetRequestId(v string) *ListInstagramPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstagramPageResponseBody) SetSuccess(v bool) *ListInstagramPageResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstagramPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstagramPageResponseBodyData struct {
	AgentInfo     *ListInstagramPageResponseBodyDataAgentInfo     `json:"AgentInfo,omitempty" xml:"AgentInfo,omitempty" type:"Struct"`
	BeebotInfo    *ListInstagramPageResponseBodyDataBeebotInfo    `json:"BeebotInfo,omitempty" xml:"BeebotInfo,omitempty" type:"Struct"`
	InstagramInfo *ListInstagramPageResponseBodyDataInstagramInfo `json:"InstagramInfo,omitempty" xml:"InstagramInfo,omitempty" type:"Struct"`
}

func (s ListInstagramPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstagramPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstagramPageResponseBodyData) GetAgentInfo() *ListInstagramPageResponseBodyDataAgentInfo {
	return s.AgentInfo
}

func (s *ListInstagramPageResponseBodyData) GetBeebotInfo() *ListInstagramPageResponseBodyDataBeebotInfo {
	return s.BeebotInfo
}

func (s *ListInstagramPageResponseBodyData) GetInstagramInfo() *ListInstagramPageResponseBodyDataInstagramInfo {
	return s.InstagramInfo
}

func (s *ListInstagramPageResponseBodyData) SetAgentInfo(v *ListInstagramPageResponseBodyDataAgentInfo) *ListInstagramPageResponseBodyData {
	s.AgentInfo = v
	return s
}

func (s *ListInstagramPageResponseBodyData) SetBeebotInfo(v *ListInstagramPageResponseBodyDataBeebotInfo) *ListInstagramPageResponseBodyData {
	s.BeebotInfo = v
	return s
}

func (s *ListInstagramPageResponseBodyData) SetInstagramInfo(v *ListInstagramPageResponseBodyDataInstagramInfo) *ListInstagramPageResponseBodyData {
	s.InstagramInfo = v
	return s
}

func (s *ListInstagramPageResponseBodyData) Validate() error {
	if s.AgentInfo != nil {
		if err := s.AgentInfo.Validate(); err != nil {
			return err
		}
	}
	if s.BeebotInfo != nil {
		if err := s.BeebotInfo.Validate(); err != nil {
			return err
		}
	}
	if s.InstagramInfo != nil {
		if err := s.InstagramInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstagramPageResponseBodyDataAgentInfo struct {
	// example:
	//
	// aaa,bbb
	AgentKeywords *string `json:"AgentKeywords,omitempty" xml:"AgentKeywords,omitempty"`
	// example:
	//
	// Y
	AgentNoKeywords *string `json:"AgentNoKeywords,omitempty" xml:"AgentNoKeywords,omitempty"`
}

func (s ListInstagramPageResponseBodyDataAgentInfo) String() string {
	return dara.Prettify(s)
}

func (s ListInstagramPageResponseBodyDataAgentInfo) GoString() string {
	return s.String()
}

func (s *ListInstagramPageResponseBodyDataAgentInfo) GetAgentKeywords() *string {
	return s.AgentKeywords
}

func (s *ListInstagramPageResponseBodyDataAgentInfo) GetAgentNoKeywords() *string {
	return s.AgentNoKeywords
}

func (s *ListInstagramPageResponseBodyDataAgentInfo) SetAgentKeywords(v string) *ListInstagramPageResponseBodyDataAgentInfo {
	s.AgentKeywords = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataAgentInfo) SetAgentNoKeywords(v string) *ListInstagramPageResponseBodyDataAgentInfo {
	s.AgentNoKeywords = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataAgentInfo) Validate() error {
	return dara.Validate(s)
}

type ListInstagramPageResponseBodyDataBeebotInfo struct {
	// example:
	//
	// instance1
	BeebotInstanceId *string `json:"BeebotInstanceId,omitempty" xml:"BeebotInstanceId,omitempty"`
	// example:
	//
	// name1
	BeebotNamespaceId *string `json:"BeebotNamespaceId,omitempty" xml:"BeebotNamespaceId,omitempty"`
}

func (s ListInstagramPageResponseBodyDataBeebotInfo) String() string {
	return dara.Prettify(s)
}

func (s ListInstagramPageResponseBodyDataBeebotInfo) GoString() string {
	return s.String()
}

func (s *ListInstagramPageResponseBodyDataBeebotInfo) GetBeebotInstanceId() *string {
	return s.BeebotInstanceId
}

func (s *ListInstagramPageResponseBodyDataBeebotInfo) GetBeebotNamespaceId() *string {
	return s.BeebotNamespaceId
}

func (s *ListInstagramPageResponseBodyDataBeebotInfo) SetBeebotInstanceId(v string) *ListInstagramPageResponseBodyDataBeebotInfo {
	s.BeebotInstanceId = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataBeebotInfo) SetBeebotNamespaceId(v string) *ListInstagramPageResponseBodyDataBeebotInfo {
	s.BeebotNamespaceId = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataBeebotInfo) Validate() error {
	return dara.Validate(s)
}

type ListInstagramPageResponseBodyDataInstagramInfo struct {
	// example:
	//
	// 176546546464
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// matrryhtr
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// connected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// example:
	//
	// Y
	HttpFlag *string `json:"HttpFlag,omitempty" xml:"HttpFlag,omitempty"`
	// example:
	//
	// 213254324532523
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// example:
	//
	// iwhaclecloud-2
	PageName *string `json:"PageName,omitempty" xml:"PageName,omitempty"`
	// example:
	//
	// Y
	QueueFlag *string `json:"QueueFlag,omitempty" xml:"QueueFlag,omitempty"`
	// example:
	//
	// aaa
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// example:
	//
	// HTTP://WWW.***.COM
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty" xml:"StatusCallbackUrl,omitempty"`
	// example:
	//
	// AAA
	StatusQueue *string `json:"StatusQueue,omitempty" xml:"StatusQueue,omitempty"`
	// example:
	//
	// HTTP://WWW.***.COM
	UpCallbackUrl *string `json:"UpCallbackUrl,omitempty" xml:"UpCallbackUrl,omitempty"`
	// example:
	//
	// BBB
	UpQueue *string `json:"UpQueue,omitempty" xml:"UpQueue,omitempty"`
}

func (s ListInstagramPageResponseBodyDataInstagramInfo) String() string {
	return dara.Prettify(s)
}

func (s ListInstagramPageResponseBodyDataInstagramInfo) GoString() string {
	return s.String()
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) GetAccountId() *string {
	return s.AccountId
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) GetAccountName() *string {
	return s.AccountName
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) GetHttpFlag() *string {
	return s.HttpFlag
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) GetPageId() *string {
	return s.PageId
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) GetPageName() *string {
	return s.PageName
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) GetQueueFlag() *string {
	return s.QueueFlag
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) GetQueueName() *string {
	return s.QueueName
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) GetStatusCallbackUrl() *string {
	return s.StatusCallbackUrl
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) GetStatusQueue() *string {
	return s.StatusQueue
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) GetUpCallbackUrl() *string {
	return s.UpCallbackUrl
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) GetUpQueue() *string {
	return s.UpQueue
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) SetAccountId(v string) *ListInstagramPageResponseBodyDataInstagramInfo {
	s.AccountId = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) SetAccountName(v string) *ListInstagramPageResponseBodyDataInstagramInfo {
	s.AccountName = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) SetConnectionStatus(v string) *ListInstagramPageResponseBodyDataInstagramInfo {
	s.ConnectionStatus = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) SetHttpFlag(v string) *ListInstagramPageResponseBodyDataInstagramInfo {
	s.HttpFlag = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) SetPageId(v string) *ListInstagramPageResponseBodyDataInstagramInfo {
	s.PageId = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) SetPageName(v string) *ListInstagramPageResponseBodyDataInstagramInfo {
	s.PageName = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) SetQueueFlag(v string) *ListInstagramPageResponseBodyDataInstagramInfo {
	s.QueueFlag = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) SetQueueName(v string) *ListInstagramPageResponseBodyDataInstagramInfo {
	s.QueueName = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) SetStatusCallbackUrl(v string) *ListInstagramPageResponseBodyDataInstagramInfo {
	s.StatusCallbackUrl = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) SetStatusQueue(v string) *ListInstagramPageResponseBodyDataInstagramInfo {
	s.StatusQueue = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) SetUpCallbackUrl(v string) *ListInstagramPageResponseBodyDataInstagramInfo {
	s.UpCallbackUrl = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) SetUpQueue(v string) *ListInstagramPageResponseBodyDataInstagramInfo {
	s.UpQueue = &v
	return s
}

func (s *ListInstagramPageResponseBodyDataInstagramInfo) Validate() error {
	return dara.Validate(s)
}
