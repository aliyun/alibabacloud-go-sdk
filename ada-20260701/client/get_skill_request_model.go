// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetSkillRequest
	GetName() *string
	SetNetwork(v string) *GetSkillRequest
	GetNetwork() *string
	SetSkillVersion(v int64) *GetSkillRequest
	GetSkillVersion() *int64
}

type GetSkillRequest struct {
	// The Skill name.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-review
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type of the download URL. Valid values: public and internal. If omitted, no download URL is generated.
	//
	// example:
	//
	// public
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The release history version number to query. If omitted, the current Skill main record is returned.
	//
	// example:
	//
	// 2
	SkillVersion *int64 `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
}

func (s GetSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillRequest) GoString() string {
	return s.String()
}

func (s *GetSkillRequest) GetName() *string {
	return s.Name
}

func (s *GetSkillRequest) GetNetwork() *string {
	return s.Network
}

func (s *GetSkillRequest) GetSkillVersion() *int64 {
	return s.SkillVersion
}

func (s *GetSkillRequest) SetName(v string) *GetSkillRequest {
	s.Name = &v
	return s
}

func (s *GetSkillRequest) SetNetwork(v string) *GetSkillRequest {
	s.Network = &v
	return s
}

func (s *GetSkillRequest) SetSkillVersion(v int64) *GetSkillRequest {
	s.SkillVersion = &v
	return s
}

func (s *GetSkillRequest) Validate() error {
	return dara.Validate(s)
}
