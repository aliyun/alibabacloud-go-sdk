// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClientFromBlackListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIP(v string) *RemoveClientFromBlackListRequest
	GetClientIP() *string
	SetClientToken(v string) *RemoveClientFromBlackListRequest
	GetClientToken() *string
	SetFileSystemId(v string) *RemoveClientFromBlackListRequest
	GetFileSystemId() *string
	SetRegionId(v string) *RemoveClientFromBlackListRequest
	GetRegionId() *string
}

type RemoveClientFromBlackListRequest struct {
	// The IP address of the client to be removed from the blacklist.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-00d91ca404a348****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The region ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveClientFromBlackListRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveClientFromBlackListRequest) GoString() string {
	return s.String()
}

func (s *RemoveClientFromBlackListRequest) GetClientIP() *string {
	return s.ClientIP
}

func (s *RemoveClientFromBlackListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveClientFromBlackListRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *RemoveClientFromBlackListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveClientFromBlackListRequest) SetClientIP(v string) *RemoveClientFromBlackListRequest {
	s.ClientIP = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) SetClientToken(v string) *RemoveClientFromBlackListRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) SetFileSystemId(v string) *RemoveClientFromBlackListRequest {
	s.FileSystemId = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) SetRegionId(v string) *RemoveClientFromBlackListRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) Validate() error {
	return dara.Validate(s)
}
