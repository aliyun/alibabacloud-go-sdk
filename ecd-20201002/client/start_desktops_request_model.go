// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *StartDesktopsRequest
	GetClientId() *string
	SetClientOS(v string) *StartDesktopsRequest
	GetClientOS() *string
	SetClientToken(v string) *StartDesktopsRequest
	GetClientToken() *string
	SetClientVersion(v string) *StartDesktopsRequest
	GetClientVersion() *string
	SetDesktopId(v []*string) *StartDesktopsRequest
	GetDesktopId() []*string
	SetLoginToken(v string) *StartDesktopsRequest
	GetLoginToken() *string
	SetRegionId(v string) *StartDesktopsRequest
	GetRegionId() *string
	SetSessionId(v string) *StartDesktopsRequest
	GetSessionId() *string
	SetUuid(v string) *StartDesktopsRequest
	GetUuid() *string
}

type StartDesktopsRequest struct {
	// The ID of the Alibaba Cloud Workspace client (hereinafter referred to as WUYING client). The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The operating system (OS) of the device that run the client.
	//
	// example:
	//
	// Windows_NT 10.0.18363 x64
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 21e7be12-aa4f-4389-b3e1-82f4a1b5****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The client version. If you use a WUYING client, you can click **About*	- on the client logon page to view the version of the client.
	//
	// example:
	//
	// 2.1.0-R-20210731.151756
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The IDs of the cloud computers. You can specify the IDs of 1 to 20 cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-cg27ufmapab08****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by WUYING Workspace.
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
	// The UUID of the client.
	//
	// example:
	//
	// 71F6A700735E74A61161A53F0C47****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s StartDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StartDesktopsRequest) GetClientId() *string {
	return s.ClientId
}

func (s *StartDesktopsRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *StartDesktopsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartDesktopsRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *StartDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *StartDesktopsRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *StartDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartDesktopsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StartDesktopsRequest) GetUuid() *string {
	return s.Uuid
}

func (s *StartDesktopsRequest) SetClientId(v string) *StartDesktopsRequest {
	s.ClientId = &v
	return s
}

func (s *StartDesktopsRequest) SetClientOS(v string) *StartDesktopsRequest {
	s.ClientOS = &v
	return s
}

func (s *StartDesktopsRequest) SetClientToken(v string) *StartDesktopsRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDesktopsRequest) SetClientVersion(v string) *StartDesktopsRequest {
	s.ClientVersion = &v
	return s
}

func (s *StartDesktopsRequest) SetDesktopId(v []*string) *StartDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StartDesktopsRequest) SetLoginToken(v string) *StartDesktopsRequest {
	s.LoginToken = &v
	return s
}

func (s *StartDesktopsRequest) SetRegionId(v string) *StartDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StartDesktopsRequest) SetSessionId(v string) *StartDesktopsRequest {
	s.SessionId = &v
	return s
}

func (s *StartDesktopsRequest) SetUuid(v string) *StartDesktopsRequest {
	s.Uuid = &v
	return s
}

func (s *StartDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
