// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAvatarProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperateType(v string) *OperateAvatarProjectRequest
	GetOperateType() *string
	SetProjectId(v string) *OperateAvatarProjectRequest
	GetProjectId() *string
	SetResChannelNumber(v int32) *OperateAvatarProjectRequest
	GetResChannelNumber() *int32
	SetResType(v string) *OperateAvatarProjectRequest
	GetResType() *string
}

type OperateAvatarProjectRequest struct {
	// example:
	//
	// DELETE
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// example:
	//
	// 812907463682949120
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 1
	ResChannelNumber *int32 `json:"resChannelNumber,omitempty" xml:"resChannelNumber,omitempty"`
	// example:
	//
	// FREE
	ResType *string `json:"resType,omitempty" xml:"resType,omitempty"`
}

func (s OperateAvatarProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *OperateAvatarProjectRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *OperateAvatarProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *OperateAvatarProjectRequest) GetResChannelNumber() *int32 {
	return s.ResChannelNumber
}

func (s *OperateAvatarProjectRequest) GetResType() *string {
	return s.ResType
}

func (s *OperateAvatarProjectRequest) SetOperateType(v string) *OperateAvatarProjectRequest {
	s.OperateType = &v
	return s
}

func (s *OperateAvatarProjectRequest) SetProjectId(v string) *OperateAvatarProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *OperateAvatarProjectRequest) SetResChannelNumber(v int32) *OperateAvatarProjectRequest {
	s.ResChannelNumber = &v
	return s
}

func (s *OperateAvatarProjectRequest) SetResType(v string) *OperateAvatarProjectRequest {
	s.ResType = &v
	return s
}

func (s *OperateAvatarProjectRequest) Validate() error {
	return dara.Validate(s)
}
