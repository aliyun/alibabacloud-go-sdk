// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlias interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalVersionWeight(v map[string]*float32) *Alias
	GetAdditionalVersionWeight() map[string]*float32
	SetAliasName(v string) *Alias
	GetAliasName() *string
	SetCreatedTime(v string) *Alias
	GetCreatedTime() *string
	SetDescription(v string) *Alias
	GetDescription() *string
	SetLastModifiedTime(v string) *Alias
	GetLastModifiedTime() *string
	SetVersionId(v string) *Alias
	GetVersionId() *string
}

type Alias struct {
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight" xml:"additionalVersionWeight"`
	// example:
	//
	// prod
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// example:
	//
	// 2006-01-02T15:04:05Z07:00
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// alias for pre env
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2006-01-02T15:04:05Z07:00
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// example:
	//
	// 1
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s Alias) String() string {
	return dara.Prettify(s)
}

func (s Alias) GoString() string {
	return s.String()
}

func (s *Alias) GetAdditionalVersionWeight() map[string]*float32 {
	return s.AdditionalVersionWeight
}

func (s *Alias) GetAliasName() *string {
	return s.AliasName
}

func (s *Alias) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Alias) GetDescription() *string {
	return s.Description
}

func (s *Alias) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *Alias) GetVersionId() *string {
	return s.VersionId
}

func (s *Alias) SetAdditionalVersionWeight(v map[string]*float32) *Alias {
	s.AdditionalVersionWeight = v
	return s
}

func (s *Alias) SetAliasName(v string) *Alias {
	s.AliasName = &v
	return s
}

func (s *Alias) SetCreatedTime(v string) *Alias {
	s.CreatedTime = &v
	return s
}

func (s *Alias) SetDescription(v string) *Alias {
	s.Description = &v
	return s
}

func (s *Alias) SetLastModifiedTime(v string) *Alias {
	s.LastModifiedTime = &v
	return s
}

func (s *Alias) SetVersionId(v string) *Alias {
	s.VersionId = &v
	return s
}

func (s *Alias) Validate() error {
	return dara.Validate(s)
}
