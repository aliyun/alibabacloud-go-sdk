// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForSaveArtExtensionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDateOrPeriod(v string) *SaveSingleTaskForSaveArtExtensionRequest
	GetDateOrPeriod() *string
	SetDimensions(v string) *SaveSingleTaskForSaveArtExtensionRequest
	GetDimensions() *string
	SetDomainName(v string) *SaveSingleTaskForSaveArtExtensionRequest
	GetDomainName() *string
	SetFeatures(v string) *SaveSingleTaskForSaveArtExtensionRequest
	GetFeatures() *string
	SetInscriptionsAndMarkings(v string) *SaveSingleTaskForSaveArtExtensionRequest
	GetInscriptionsAndMarkings() *string
	SetLang(v string) *SaveSingleTaskForSaveArtExtensionRequest
	GetLang() *string
	SetMaker(v string) *SaveSingleTaskForSaveArtExtensionRequest
	GetMaker() *string
	SetMaterialsAndTechniques(v string) *SaveSingleTaskForSaveArtExtensionRequest
	GetMaterialsAndTechniques() *string
	SetObjectType(v string) *SaveSingleTaskForSaveArtExtensionRequest
	GetObjectType() *string
	SetReference(v string) *SaveSingleTaskForSaveArtExtensionRequest
	GetReference() *string
	SetSubject(v string) *SaveSingleTaskForSaveArtExtensionRequest
	GetSubject() *string
	SetTitle(v string) *SaveSingleTaskForSaveArtExtensionRequest
	GetTitle() *string
}

type SaveSingleTaskForSaveArtExtensionRequest struct {
	// example:
	//
	// 2019-10-01
	DateOrPeriod *string `json:"DateOrPeriod,omitempty" xml:"DateOrPeriod,omitempty"`
	// example:
	//
	// 20 cm
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.art
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// iconicity
	Features *string `json:"Features,omitempty" xml:"Features,omitempty"`
	// example:
	//
	// realism
	InscriptionsAndMarkings *string `json:"InscriptionsAndMarkings,omitempty" xml:"InscriptionsAndMarkings,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// zhang san
	Maker *string `json:"Maker,omitempty" xml:"Maker,omitempty"`
	// example:
	//
	// silk
	MaterialsAndTechniques *string `json:"MaterialsAndTechniques,omitempty" xml:"MaterialsAndTechniques,omitempty"`
	// example:
	//
	// The embroidery
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// drawings
	Reference *string `json:"Reference,omitempty" xml:"Reference,omitempty"`
	// example:
	//
	// peace
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// example:
	//
	// Peace and friendship
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SaveSingleTaskForSaveArtExtensionRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForSaveArtExtensionRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) GetDateOrPeriod() *string {
	return s.DateOrPeriod
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) GetFeatures() *string {
	return s.Features
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) GetInscriptionsAndMarkings() *string {
	return s.InscriptionsAndMarkings
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) GetMaker() *string {
	return s.Maker
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) GetMaterialsAndTechniques() *string {
	return s.MaterialsAndTechniques
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) GetReference() *string {
	return s.Reference
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) GetSubject() *string {
	return s.Subject
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) GetTitle() *string {
	return s.Title
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetDateOrPeriod(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.DateOrPeriod = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetDimensions(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Dimensions = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetDomainName(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetFeatures(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Features = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetInscriptionsAndMarkings(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.InscriptionsAndMarkings = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetLang(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetMaker(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Maker = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetMaterialsAndTechniques(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.MaterialsAndTechniques = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetObjectType(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.ObjectType = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetReference(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Reference = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetSubject(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Subject = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetTitle(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Title = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) Validate() error {
	return dara.Validate(s)
}
