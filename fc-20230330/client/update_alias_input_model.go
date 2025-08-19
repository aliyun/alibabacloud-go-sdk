// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAliasInput interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalVersionWeight(v map[string]*float32) *UpdateAliasInput
	GetAdditionalVersionWeight() map[string]*float32
	SetDescription(v string) *UpdateAliasInput
	GetDescription() *string
	SetVersionId(v string) *UpdateAliasInput
	GetVersionId() *string
}

type UpdateAliasInput struct {
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight" xml:"additionalVersionWeight"`
	// example:
	//
	// my alias
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s UpdateAliasInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateAliasInput) GoString() string {
	return s.String()
}

func (s *UpdateAliasInput) GetAdditionalVersionWeight() map[string]*float32 {
	return s.AdditionalVersionWeight
}

func (s *UpdateAliasInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateAliasInput) GetVersionId() *string {
	return s.VersionId
}

func (s *UpdateAliasInput) SetAdditionalVersionWeight(v map[string]*float32) *UpdateAliasInput {
	s.AdditionalVersionWeight = v
	return s
}

func (s *UpdateAliasInput) SetDescription(v string) *UpdateAliasInput {
	s.Description = &v
	return s
}

func (s *UpdateAliasInput) SetVersionId(v string) *UpdateAliasInput {
	s.VersionId = &v
	return s
}

func (s *UpdateAliasInput) Validate() error {
	return dara.Validate(s)
}
