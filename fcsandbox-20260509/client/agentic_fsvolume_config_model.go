// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticFSVolumeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetGroupID(v int32) *AgenticFSVolumeConfig
	GetGroupID() *int32
	SetServerAddr(v string) *AgenticFSVolumeConfig
	GetServerAddr() *string
	SetUserID(v int32) *AgenticFSVolumeConfig
	GetUserID() *int32
}

type AgenticFSVolumeConfig struct {
	GroupID    *int32  `json:"groupID,omitempty" xml:"groupID,omitempty"`
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
	UserID     *int32  `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s AgenticFSVolumeConfig) String() string {
	return dara.Prettify(s)
}

func (s AgenticFSVolumeConfig) GoString() string {
	return s.String()
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
