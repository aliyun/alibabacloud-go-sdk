// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAliasInput interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalVersionWeight(v map[string]*float32) *CreateAliasInput
	GetAdditionalVersionWeight() map[string]*float32
	SetAliasName(v string) *CreateAliasInput
	GetAliasName() *string
	SetDescription(v string) *CreateAliasInput
	GetDescription() *string
	SetVersionId(v string) *CreateAliasInput
	GetVersionId() *string
}

type CreateAliasInput struct {
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight" xml:"additionalVersionWeight"`
	// This parameter is required.
	//
	// example:
	//
	// prod
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// example:
	//
	// my alias
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s CreateAliasInput) String() string {
	return dara.Prettify(s)
}

func (s CreateAliasInput) GoString() string {
	return s.String()
}

func (s *CreateAliasInput) GetAdditionalVersionWeight() map[string]*float32 {
	return s.AdditionalVersionWeight
}

func (s *CreateAliasInput) GetAliasName() *string {
	return s.AliasName
}

func (s *CreateAliasInput) GetDescription() *string {
	return s.Description
}

func (s *CreateAliasInput) GetVersionId() *string {
	return s.VersionId
}

func (s *CreateAliasInput) SetAdditionalVersionWeight(v map[string]*float32) *CreateAliasInput {
	s.AdditionalVersionWeight = v
	return s
}

func (s *CreateAliasInput) SetAliasName(v string) *CreateAliasInput {
	s.AliasName = &v
	return s
}

func (s *CreateAliasInput) SetDescription(v string) *CreateAliasInput {
	s.Description = &v
	return s
}

func (s *CreateAliasInput) SetVersionId(v string) *CreateAliasInput {
	s.VersionId = &v
	return s
}

func (s *CreateAliasInput) Validate() error {
	return dara.Validate(s)
}
