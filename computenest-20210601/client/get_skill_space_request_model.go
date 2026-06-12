// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSkillSpaceId(v string) *GetSkillSpaceRequest
	GetSkillSpaceId() *string
}

type GetSkillSpaceRequest struct {
	// The ID of the SkillSpace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ss-11111
	SkillSpaceId *string `json:"SkillSpaceId,omitempty" xml:"SkillSpaceId,omitempty"`
}

func (s GetSkillSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillSpaceRequest) GoString() string {
	return s.String()
}

func (s *GetSkillSpaceRequest) GetSkillSpaceId() *string {
	return s.SkillSpaceId
}

func (s *GetSkillSpaceRequest) SetSkillSpaceId(v string) *GetSkillSpaceRequest {
	s.SkillSpaceId = &v
	return s
}

func (s *GetSkillSpaceRequest) Validate() error {
	return dara.Validate(s)
}
