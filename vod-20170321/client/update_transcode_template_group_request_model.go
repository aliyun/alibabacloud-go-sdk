// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTranscodeTemplateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLocked(v string) *UpdateTranscodeTemplateGroupRequest
	GetLocked() *string
	SetName(v string) *UpdateTranscodeTemplateGroupRequest
	GetName() *string
	SetTranscodeTemplateGroupId(v string) *UpdateTranscodeTemplateGroupRequest
	GetTranscodeTemplateGroupId() *string
	SetTranscodeTemplateList(v string) *UpdateTranscodeTemplateGroupRequest
	GetTranscodeTemplateList() *string
}

type UpdateTranscodeTemplateGroupRequest struct {
	// The lock status of the transcoding template group. Valid values:
	//
	// 	- **Enabled**: The transcoding template group is locked and cannot be modified.
	//
	// 	- **Disabled*	- (default): The transcoding template group is not locked.
	//
	// example:
	//
	// Disabled
	Locked *string `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The name of the transcoding template group.
	//
	// 	- The name cannot exceed 128 bytes.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// transcodetemplate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the transcoding template group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4c71a339fe*****52b4fa6f4527
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	// The configurations of the transcoding template. The value must be a JSON string. For more information about the data structure, see [TranscodeTemplate](~~52839#title-9mb-8o2-uu6~~).
	//
	// example:
	//
	// [{"Video":{"Bitrate":"400","Codec":"H.264","Fps":"30"},"Audio":{"Codec":"AAC","Bitrate":"64","Definition":"SD","EncryptType":"Private","Container":{"Format":"m3u8"},"PackageType":"HLSPackage"}}]
	TranscodeTemplateList *string `json:"TranscodeTemplateList,omitempty" xml:"TranscodeTemplateList,omitempty"`
}

func (s UpdateTranscodeTemplateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTranscodeTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateTranscodeTemplateGroupRequest) GetLocked() *string {
	return s.Locked
}

func (s *UpdateTranscodeTemplateGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTranscodeTemplateGroupRequest) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *UpdateTranscodeTemplateGroupRequest) GetTranscodeTemplateList() *string {
	return s.TranscodeTemplateList
}

func (s *UpdateTranscodeTemplateGroupRequest) SetLocked(v string) *UpdateTranscodeTemplateGroupRequest {
	s.Locked = &v
	return s
}

func (s *UpdateTranscodeTemplateGroupRequest) SetName(v string) *UpdateTranscodeTemplateGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateTranscodeTemplateGroupRequest) SetTranscodeTemplateGroupId(v string) *UpdateTranscodeTemplateGroupRequest {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *UpdateTranscodeTemplateGroupRequest) SetTranscodeTemplateList(v string) *UpdateTranscodeTemplateGroupRequest {
	s.TranscodeTemplateList = &v
	return s
}

func (s *UpdateTranscodeTemplateGroupRequest) Validate() error {
	return dara.Validate(s)
}
