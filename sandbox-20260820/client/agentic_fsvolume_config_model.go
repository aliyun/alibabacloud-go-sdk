// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticFSVolumeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointID(v string) *AgenticFSVolumeConfig
	GetAccessPointID() *string
	SetAgenticSpaceID(v string) *AgenticFSVolumeConfig
	GetAgenticSpaceID() *string
	SetFileSystemID(v string) *AgenticFSVolumeConfig
	GetFileSystemID() *string
	SetGroupID(v int32) *AgenticFSVolumeConfig
	GetGroupID() *int32
	SetServerAddr(v string) *AgenticFSVolumeConfig
	GetServerAddr() *string
	SetUserID(v int32) *AgenticFSVolumeConfig
	GetUserID() *int32
}

type AgenticFSVolumeConfig struct {
	// The access point ID.
	//
	// example:
	//
	// ap-xxxx
	AccessPointID *string `json:"accessPointID,omitempty" xml:"accessPointID,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// agentic-xxxx
	AgenticSpaceID *string `json:"agenticSpaceID,omitempty" xml:"agenticSpaceID,omitempty"`
	// The file system ID.
	//
	// example:
	//
	// 03204sl2qjiax4ocdt2
	FileSystemID *string `json:"fileSystemID,omitempty" xml:"fileSystemID,omitempty"`
	// The group ID for local mounting.
	//
	// example:
	//
	// 1000
	GroupID *int32 `json:"groupID,omitempty" xml:"groupID,omitempty"`
	// The AgenticFS access point address.
	//
	// example:
	//
	// ap-xxxx.03204sl2qjiax4ocdt2-ljs60.cn-shanghai.nas.aliyuncs.com
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
	// The user ID for local mounting.
	//
	// example:
	//
	// 1000
	UserID *int32 `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s AgenticFSVolumeConfig) String() string {
	return dara.Prettify(s)
}

func (s AgenticFSVolumeConfig) GoString() string {
	return s.String()
}

func (s *AgenticFSVolumeConfig) GetAccessPointID() *string {
	return s.AccessPointID
}

func (s *AgenticFSVolumeConfig) GetAgenticSpaceID() *string {
	return s.AgenticSpaceID
}

func (s *AgenticFSVolumeConfig) GetFileSystemID() *string {
	return s.FileSystemID
}

func (s *AgenticFSVolumeConfig) GetGroupID() *int32 {
	return s.GroupID
}

func (s *AgenticFSVolumeConfig) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *AgenticFSVolumeConfig) GetUserID() *int32 {
	return s.UserID
}

func (s *AgenticFSVolumeConfig) SetAccessPointID(v string) *AgenticFSVolumeConfig {
	s.AccessPointID = &v
	return s
}

func (s *AgenticFSVolumeConfig) SetAgenticSpaceID(v string) *AgenticFSVolumeConfig {
	s.AgenticSpaceID = &v
	return s
}

func (s *AgenticFSVolumeConfig) SetFileSystemID(v string) *AgenticFSVolumeConfig {
	s.FileSystemID = &v
	return s
}

func (s *AgenticFSVolumeConfig) SetGroupID(v int32) *AgenticFSVolumeConfig {
	s.GroupID = &v
	return s
}

func (s *AgenticFSVolumeConfig) SetServerAddr(v string) *AgenticFSVolumeConfig {
	s.ServerAddr = &v
	return s
}

func (s *AgenticFSVolumeConfig) SetUserID(v int32) *AgenticFSVolumeConfig {
	s.UserID = &v
	return s
}

func (s *AgenticFSVolumeConfig) Validate() error {
	return dara.Validate(s)
}
