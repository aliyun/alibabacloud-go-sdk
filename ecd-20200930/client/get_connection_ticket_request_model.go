// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandContent(v string) *GetConnectionTicketRequest
	GetCommandContent() *string
	SetDesktopId(v string) *GetConnectionTicketRequest
	GetDesktopId() *string
	SetEndUserId(v string) *GetConnectionTicketRequest
	GetEndUserId() *string
	SetOwnerId(v int64) *GetConnectionTicketRequest
	GetOwnerId() *int64
	SetPassword(v string) *GetConnectionTicketRequest
	GetPassword() *string
	SetRegionId(v string) *GetConnectionTicketRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetConnectionTicketRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetConnectionTicketRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *GetConnectionTicketRequest
	GetTaskId() *string
	SetUuid(v string) *GetConnectionTicketRequest
	GetUuid() *string
}

type GetConnectionTicketRequest struct {
	// The command that you want to run to configure a custom application in user mode. After you obtain the credential, the application is automatically started. Parameter description in the command:
	//
	// 	- appPath: the path of the application startup file. Example: `"C:\\\\Program Files (x86)\\\\000\\\\000.exe"`. Use double slashes (\\\\\\) as the delimiter. Type of the parameter value: string.
	//
	// 	- appParameter: the startup arguments of the application. Example: `"meetingid 000 meetingname aaa"`. Separate multiple arguments with spaces. Type of the parameter value: string.
	//
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
	// The ID of the cloud computer for which you want to generate a connection credential. This parameter is required.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The ID of the end user of the cloud computer. The end user must be the current end user of the cloud computer.
	//
	// example:
	//
	// Alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the current end user of the cloud computer.
	//
	// example:
	//
	// Ab123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the cloud computer connection task.
	//
	// example:
	//
	// 2afbad19-778a-4fc5-9674-1f19c63862da
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The unique identifier of the client. If you use an Alibaba Cloud Workspace client, click **About*	- on the client logon page to view the identifier of the client.
	//
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

func (s *GetConnectionTicketRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *GetConnectionTicketRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *GetConnectionTicketRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *GetConnectionTicketRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetConnectionTicketRequest) GetPassword() *string {
	return s.Password
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

func (s *GetConnectionTicketRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetConnectionTicketRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetConnectionTicketRequest) SetCommandContent(v string) *GetConnectionTicketRequest {
	s.CommandContent = &v
	return s
}

func (s *GetConnectionTicketRequest) SetDesktopId(v string) *GetConnectionTicketRequest {
	s.DesktopId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetEndUserId(v string) *GetConnectionTicketRequest {
	s.EndUserId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetOwnerId(v int64) *GetConnectionTicketRequest {
	s.OwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetPassword(v string) *GetConnectionTicketRequest {
	s.Password = &v
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

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetUuid(v string) *GetConnectionTicketRequest {
	s.Uuid = &v
	return s
}

func (s *GetConnectionTicketRequest) Validate() error {
	return dara.Validate(s)
}
