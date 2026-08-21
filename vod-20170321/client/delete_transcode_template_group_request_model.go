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
	// Indicates whether to force delete the entire transcoding template group. Valid values:
	//
	// - **true**: Force deletes the entire template group and all its transcoding templates.
	//
	// - **false*	- (default): Deletes only the specified transcoding templates.
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
	// - Separate multiple IDs with commas (,).
	//
	// - A maximum of 10 template IDs are supported.
	//
	// - This parameter is required if the ForceDelGroup parameter is empty or set to false.
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
