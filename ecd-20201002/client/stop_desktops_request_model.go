// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *StopDesktopsRequest
	GetClientId() *string
	SetClientOS(v string) *StopDesktopsRequest
	GetClientOS() *string
	SetClientToken(v string) *StopDesktopsRequest
	GetClientToken() *string
	SetClientVersion(v string) *StopDesktopsRequest
	GetClientVersion() *string
	SetDesktopId(v []*string) *StopDesktopsRequest
	GetDesktopId() []*string
	SetLoginToken(v string) *StopDesktopsRequest
	GetLoginToken() *string
	SetOsUpdate(v bool) *StopDesktopsRequest
	GetOsUpdate() *bool
	SetRegionId(v string) *StopDesktopsRequest
	GetRegionId() *string
	SetSessionId(v string) *StopDesktopsRequest
	GetSessionId() *string
	SetSessionToken(v string) *StopDesktopsRequest
	GetSessionToken() *string
	SetUuid(v string) *StopDesktopsRequest
	GetUuid() *string
}

type StopDesktopsRequest struct {
	// A unique ID that the client generates.
	//
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The operating system of the device where the client runs.
	//
	// example:
	//
	// Windows_NT 10.0.18363 x64
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// A client token to ensure idempotence. Generate a unique ASCII string for each request. The string must be no more than 64 characters long. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 6ce412a8-399f-49f9-9518-66ee028a****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The client version. If you use the WUYING Workspace client, find this value in the **About*	- dialog box on the logon screen.
	//
	// example:
	//
	// 2.1.0-R-20210731.151756
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The IDs of the cloud computers. You can specify from 1 to 20 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The logon credential.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OsUpdate   *bool   `json:"OsUpdate,omitempty" xml:"OsUpdate,omitempty"`
	// The region ID. Call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to get a list of regions where WUYING Workspace is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The logon credential.
	//
	// example:
	//
	// 04b7b80a0b020715c5c1b4175fc4771698****9e2a759557a4624665fd53ae40
	SessionToken *string `json:"SessionToken,omitempty" xml:"SessionToken,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s StopDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StopDesktopsRequest) GetClientId() *string {
	return s.ClientId
}

func (s *StopDesktopsRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *StopDesktopsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopDesktopsRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *StopDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *StopDesktopsRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *StopDesktopsRequest) GetOsUpdate() *bool {
	return s.OsUpdate
}

func (s *StopDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopDesktopsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StopDesktopsRequest) GetSessionToken() *string {
	return s.SessionToken
}

func (s *StopDesktopsRequest) GetUuid() *string {
	return s.Uuid
}

func (s *StopDesktopsRequest) SetClientId(v string) *StopDesktopsRequest {
	s.ClientId = &v
	return s
}

func (s *StopDesktopsRequest) SetClientOS(v string) *StopDesktopsRequest {
	s.ClientOS = &v
	return s
}

func (s *StopDesktopsRequest) SetClientToken(v string) *StopDesktopsRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDesktopsRequest) SetClientVersion(v string) *StopDesktopsRequest {
	s.ClientVersion = &v
	return s
}

func (s *StopDesktopsRequest) SetDesktopId(v []*string) *StopDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StopDesktopsRequest) SetLoginToken(v string) *StopDesktopsRequest {
	s.LoginToken = &v
	return s
}

func (s *StopDesktopsRequest) SetOsUpdate(v bool) *StopDesktopsRequest {
	s.OsUpdate = &v
	return s
}

func (s *StopDesktopsRequest) SetRegionId(v string) *StopDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StopDesktopsRequest) SetSessionId(v string) *StopDesktopsRequest {
	s.SessionId = &v
	return s
}

func (s *StopDesktopsRequest) SetSessionToken(v string) *StopDesktopsRequest {
	s.SessionToken = &v
	return s
}

func (s *StopDesktopsRequest) SetUuid(v string) *StopDesktopsRequest {
	s.Uuid = &v
	return s
}

func (s *StopDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
