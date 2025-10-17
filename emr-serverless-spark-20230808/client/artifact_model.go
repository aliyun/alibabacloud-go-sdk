// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArtifact interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *Artifact
	GetBizId() *string
	SetCatagoryBizId(v string) *Artifact
	GetCatagoryBizId() *string
	SetCreator(v int64) *Artifact
	GetCreator() *int64
	SetCredential(v *Credential) *Artifact
	GetCredential() *Credential
	SetFullPath(v []*string) *Artifact
	GetFullPath() []*string
	SetGmtCreated(v string) *Artifact
	GetGmtCreated() *string
	SetGmtModified(v string) *Artifact
	GetGmtModified() *string
	SetLocation(v string) *Artifact
	GetLocation() *string
	SetModifier(v int64) *Artifact
	GetModifier() *int64
	SetModifierName(v string) *Artifact
	GetModifierName() *string
	SetName(v string) *Artifact
	GetName() *string
}

type Artifact struct {
	// This parameter is required.
	BizId         *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	CatagoryBizId *string `json:"catagoryBizId,omitempty" xml:"catagoryBizId,omitempty"`
	// This parameter is required.
	Creator    *int64      `json:"creator,omitempty" xml:"creator,omitempty"`
	Credential *Credential `json:"credential,omitempty" xml:"credential,omitempty"`
	FullPath   []*string   `json:"fullPath,omitempty" xml:"fullPath,omitempty" type:"Repeated"`
	// This parameter is required.
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// This parameter is required.
	Modifier     *int64  `json:"modifier,omitempty" xml:"modifier,omitempty"`
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s Artifact) String() string {
	return dara.Prettify(s)
}

func (s Artifact) GoString() string {
	return s.String()
}

func (s *Artifact) GetBizId() *string {
	return s.BizId
}

func (s *Artifact) GetCatagoryBizId() *string {
	return s.CatagoryBizId
}

func (s *Artifact) GetCreator() *int64 {
	return s.Creator
}

func (s *Artifact) GetCredential() *Credential {
	return s.Credential
}

func (s *Artifact) GetFullPath() []*string {
	return s.FullPath
}

func (s *Artifact) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *Artifact) GetGmtModified() *string {
	return s.GmtModified
}

func (s *Artifact) GetLocation() *string {
	return s.Location
}

func (s *Artifact) GetModifier() *int64 {
	return s.Modifier
}

func (s *Artifact) GetModifierName() *string {
	return s.ModifierName
}

func (s *Artifact) GetName() *string {
	return s.Name
}

func (s *Artifact) SetBizId(v string) *Artifact {
	s.BizId = &v
	return s
}

func (s *Artifact) SetCatagoryBizId(v string) *Artifact {
	s.CatagoryBizId = &v
	return s
}

func (s *Artifact) SetCreator(v int64) *Artifact {
	s.Creator = &v
	return s
}

func (s *Artifact) SetCredential(v *Credential) *Artifact {
	s.Credential = v
	return s
}

func (s *Artifact) SetFullPath(v []*string) *Artifact {
	s.FullPath = v
	return s
}

func (s *Artifact) SetGmtCreated(v string) *Artifact {
	s.GmtCreated = &v
	return s
}

func (s *Artifact) SetGmtModified(v string) *Artifact {
	s.GmtModified = &v
	return s
}

func (s *Artifact) SetLocation(v string) *Artifact {
	s.Location = &v
	return s
}

func (s *Artifact) SetModifier(v int64) *Artifact {
	s.Modifier = &v
	return s
}

func (s *Artifact) SetModifierName(v string) *Artifact {
	s.ModifierName = &v
	return s
}

func (s *Artifact) SetName(v string) *Artifact {
	s.Name = &v
	return s
}

func (s *Artifact) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	return nil
}
