// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTranscodeTemplateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceDelGroup(v string) *DeleteTranscodeTemplateGroupRequest
	GetForceDelGroup() *string
	SetTranscodeTemplateGroupId(v string) *DeleteTranscodeTemplateGroupRequest
	GetTranscodeTemplateGroupId() *string
	SetTranscodeTemplateIds(v string) *DeleteTranscodeTemplateGroupRequest
	GetTranscodeTemplateIds() *string
}

type DeleteTranscodeTemplateGroupRequest struct {
	// Specifies whether to forcibly delete the transcoding template group. Valid values:
	//
	// 	- **true**: deletes the transcoding template group and all the transcoding templates in the group.
	//
	// 	- **false*	- (default): deletes only the specified transcoding templates from the transcoding template group.
	//
	// example:
	//
	// true
	ForceDelGroup *string `json:"ForceDelGroup,omitempty" xml:"ForceDelGroup,omitempty"`
	// The ID of the transcoding template group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4c71a339fec*****152b4fa6f4527
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	// The IDs of the transcoding templates that you want to delete.
	//
	// 	- Separate multiple IDs with commas (,).
	//
	// 	- You can specify a maximum of 10 IDs.
	//
	// 	- This parameter is required if you set ForceDelGroup to false or leave ForceDelGroup empty.
	//
	// example:
	//
	// ["613702defdc4*****6a3b94cace1129e","bfd6c90253a2*****7fc054d7c5825"]
	TranscodeTemplateIds *string `json:"TranscodeTemplateIds,omitempty" xml:"TranscodeTemplateIds,omitempty"`
}

func (s DeleteTranscodeTemplateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTranscodeTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteTranscodeTemplateGroupRequest) GetForceDelGroup() *string {
	return s.ForceDelGroup
}

func (s *DeleteTranscodeTemplateGroupRequest) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *DeleteTranscodeTemplateGroupRequest) GetTranscodeTemplateIds() *string {
	return s.TranscodeTemplateIds
}

func (s *DeleteTranscodeTemplateGroupRequest) SetForceDelGroup(v string) *DeleteTranscodeTemplateGroupRequest {
	s.ForceDelGroup = &v
	return s
}

func (s *DeleteTranscodeTemplateGroupRequest) SetTranscodeTemplateGroupId(v string) *DeleteTranscodeTemplateGroupRequest {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *DeleteTranscodeTemplateGroupRequest) SetTranscodeTemplateIds(v string) *DeleteTranscodeTemplateGroupRequest {
	s.TranscodeTemplateIds = &v
	return s
}

func (s *DeleteTranscodeTemplateGroupRequest) Validate() error {
	return dara.Validate(s)
}
