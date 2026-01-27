// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *GetConnectionTicketRequest
	GetAccessType() *string
	SetClientId(v string) *GetConnectionTicketRequest
	GetClientId() *string
	SetClientOS(v string) *GetConnectionTicketRequest
	GetClientOS() *string
	SetClientType(v string) *GetConnectionTicketRequest
	GetClientType() *string
	SetClientVersion(v string) *GetConnectionTicketRequest
	GetClientVersion() *string
	SetCommandContent(v string) *GetConnectionTicketRequest
	GetCommandContent() *string
	SetDesktopId(v string) *GetConnectionTicketRequest
	GetDesktopId() *string
	SetLoginToken(v string) *GetConnectionTicketRequest
	GetLoginToken() *string
	SetOwnerId(v int64) *GetConnectionTicketRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetConnectionTicketRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetConnectionTicketRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetConnectionTicketRequest
	GetResourceOwnerId() *int64
	SetSessionId(v string) *GetConnectionTicketRequest
	GetSessionId() *string
	SetTag(v []*GetConnectionTicketRequestTag) *GetConnectionTicketRequest
	GetTag() []*GetConnectionTicketRequestTag
	SetTaskId(v string) *GetConnectionTicketRequest
	GetTaskId() *string
	SetTicketBlackList(v []*string) *GetConnectionTicketRequest
	GetTicketBlackList() []*string
	SetUuid(v string) *GetConnectionTicketRequest
	GetUuid() *string
}

type GetConnectionTicketRequest struct {
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// Windows_NT 10.0.18363 x64
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The type of the client.
	//
	// Valid values:
	//
	// 	- html5: web client
	//
	// 	- Linux: self-developed hardware terminal
	//
	// 	- android: Android client
	//
	// 	- windows: Windows client
	//
	// 	- ios: iOS client
	//
	// 	- macos: macOS client
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 2.1.0-R-20210731.151756
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// {
	//
	//       "startApplication": {
	//
	//             "startApplicationList": [
	//
	//                   {
	//
	//                         "sessionName": "",
	//
	//                         "appList": [
	//
	//                               {
	//
	//                                     "appPath": "C:\\\\Program Files\\\\Google\\\\Chrome\\\\Application\\\\chrome.exe",
	//
	//                                     "appParameter": "www.example.com www.example1.com"
	//
	//                               }
	//
	//                         ]
	//
	//                   }
	//
	//             ]
	//
	//       }
	//
	// }
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The cloud compute ID.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The tags. You can specify 1 to 20 tags.
	Tag []*GetConnectionTicketRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the cloud computer connection task.\\
	//
	// The first time you call the GetConnectionTicket operation, you do not need to configure this parameter. When you call the GetConnectionTicket operation later, set this parameter to the value of the `TaskId` parameter returned from the previous invocation.
	//
	// example:
	//
	// 2afbad19-778a-4fc5-9674-1f19c638****
	TaskId          *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TicketBlackList []*string `json:"TicketBlackList,omitempty" xml:"TicketBlackList,omitempty" type:"Repeated"`
	// example:
	//
	// 28c80e90-f71e-4c23-93d6-1225329cf949
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *GetConnectionTicketRequest) GetClientId() *string {
	return s.ClientId
}

func (s *GetConnectionTicketRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *GetConnectionTicketRequest) GetClientType() *string {
	return s.ClientType
}

func (s *GetConnectionTicketRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *GetConnectionTicketRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *GetConnectionTicketRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *GetConnectionTicketRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *GetConnectionTicketRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetConnectionTicketRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetConnectionTicketRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetConnectionTicketRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetConnectionTicketRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetConnectionTicketRequest) GetTag() []*GetConnectionTicketRequestTag {
	return s.Tag
}

func (s *GetConnectionTicketRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetConnectionTicketRequest) GetTicketBlackList() []*string {
	return s.TicketBlackList
}

func (s *GetConnectionTicketRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetConnectionTicketRequest) SetAccessType(v string) *GetConnectionTicketRequest {
	s.AccessType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientId(v string) *GetConnectionTicketRequest {
	s.ClientId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientOS(v string) *GetConnectionTicketRequest {
	s.ClientOS = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientType(v string) *GetConnectionTicketRequest {
	s.ClientType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientVersion(v string) *GetConnectionTicketRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetCommandContent(v string) *GetConnectionTicketRequest {
	s.CommandContent = &v
	return s
}

func (s *GetConnectionTicketRequest) SetDesktopId(v string) *GetConnectionTicketRequest {
	s.DesktopId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetLoginToken(v string) *GetConnectionTicketRequest {
	s.LoginToken = &v
	return s
}

func (s *GetConnectionTicketRequest) SetOwnerId(v int64) *GetConnectionTicketRequest {
	s.OwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetRegionId(v string) *GetConnectionTicketRequest {
	s.RegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceOwnerAccount(v string) *GetConnectionTicketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceOwnerId(v int64) *GetConnectionTicketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetSessionId(v string) *GetConnectionTicketRequest {
	s.SessionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTag(v []*GetConnectionTicketRequestTag) *GetConnectionTicketRequest {
	s.Tag = v
	return s
}

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTicketBlackList(v []*string) *GetConnectionTicketRequest {
	s.TicketBlackList = v
	return s
}

func (s *GetConnectionTicketRequest) SetUuid(v string) *GetConnectionTicketRequest {
	s.Uuid = &v
	return s
}

func (s *GetConnectionTicketRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetConnectionTicketRequestTag struct {
	// The tag key. If you specify the `Tag` parameter, you must also specify the `Key` parameter. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun` and contain only spaces.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify an empty string as a tag value. A tag value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConnectionTicketRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionTicketRequestTag) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetConnectionTicketRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetConnectionTicketRequestTag) SetKey(v string) *GetConnectionTicketRequestTag {
	s.Key = &v
	return s
}

func (s *GetConnectionTicketRequestTag) SetValue(v string) *GetConnectionTicketRequestTag {
	s.Value = &v
	return s
}

func (s *GetConnectionTicketRequestTag) Validate() error {
	return dara.Validate(s)
}
