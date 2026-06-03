// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendCallWithFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilePath(v string) *SuspendCallWithFileRequest
	GetFilePath() *string
	SetGroupId(v string) *SuspendCallWithFileRequest
	GetGroupId() *string
	SetInstanceId(v string) *SuspendCallWithFileRequest
	GetInstanceId() *string
}

type SuspendCallWithFileRequest struct {
	// example:
	//
	// xxxx
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// 0197261b-30e6-467b-83d6-7f72af868b03
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// a4274627-265f-4e14-b2d6-4ee7d4f8593e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SuspendCallWithFileRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendCallWithFileRequest) GoString() string {
	return s.String()
}

func (s *SuspendCallWithFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *SuspendCallWithFileRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SuspendCallWithFileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SuspendCallWithFileRequest) SetFilePath(v string) *SuspendCallWithFileRequest {
	s.FilePath = &v
	return s
}

func (s *SuspendCallWithFileRequest) SetGroupId(v string) *SuspendCallWithFileRequest {
	s.GroupId = &v
	return s
}

func (s *SuspendCallWithFileRequest) SetInstanceId(v string) *SuspendCallWithFileRequest {
	s.InstanceId = &v
	return s
}

func (s *SuspendCallWithFileRequest) Validate() error {
	return dara.Validate(s)
}
