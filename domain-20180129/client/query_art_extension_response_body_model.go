// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryArtExtensionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDateOrPeriod(v string) *QueryArtExtensionResponseBody
	GetDateOrPeriod() *string
	SetDimensions(v string) *QueryArtExtensionResponseBody
	GetDimensions() *string
	SetFeatures(v string) *QueryArtExtensionResponseBody
	GetFeatures() *string
	SetInscriptionsAndMarkings(v string) *QueryArtExtensionResponseBody
	GetInscriptionsAndMarkings() *string
	SetMaker(v string) *QueryArtExtensionResponseBody
	GetMaker() *string
	SetMaterialsAndTechniques(v string) *QueryArtExtensionResponseBody
	GetMaterialsAndTechniques() *string
	SetObjectType(v string) *QueryArtExtensionResponseBody
	GetObjectType() *string
	SetReference(v string) *QueryArtExtensionResponseBody
	GetReference() *string
	SetRequestId(v string) *QueryArtExtensionResponseBody
	GetRequestId() *string
	SetSubject(v string) *QueryArtExtensionResponseBody
	GetSubject() *string
	SetTitle(v string) *QueryArtExtensionResponseBody
	GetTitle() *string
}

type QueryArtExtensionResponseBody struct {
	// example:
	//
	// 2019-10-01
	DateOrPeriod *string `json:"DateOrPeriod,omitempty" xml:"DateOrPeriod,omitempty"`
	// example:
	//
	// 20 cm
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
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
	// 814B2AF0-ED6F-4C13-B41C-8AC0B1023583
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// peace
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// example:
	//
	// Peace and friendship
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s QueryArtExtensionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryArtExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryArtExtensionResponseBody) GetDateOrPeriod() *string {
	return s.DateOrPeriod
}

func (s *QueryArtExtensionResponseBody) GetDimensions() *string {
	return s.Dimensions
}

func (s *QueryArtExtensionResponseBody) GetFeatures() *string {
	return s.Features
}

func (s *QueryArtExtensionResponseBody) GetInscriptionsAndMarkings() *string {
	return s.InscriptionsAndMarkings
}

func (s *QueryArtExtensionResponseBody) GetMaker() *string {
	return s.Maker
}

func (s *QueryArtExtensionResponseBody) GetMaterialsAndTechniques() *string {
	return s.MaterialsAndTechniques
}

func (s *QueryArtExtensionResponseBody) GetObjectType() *string {
	return s.ObjectType
}

func (s *QueryArtExtensionResponseBody) GetReference() *string {
	return s.Reference
}

func (s *QueryArtExtensionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryArtExtensionResponseBody) GetSubject() *string {
	return s.Subject
}

func (s *QueryArtExtensionResponseBody) GetTitle() *string {
	return s.Title
}

func (s *QueryArtExtensionResponseBody) SetDateOrPeriod(v string) *QueryArtExtensionResponseBody {
	s.DateOrPeriod = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetDimensions(v string) *QueryArtExtensionResponseBody {
	s.Dimensions = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetFeatures(v string) *QueryArtExtensionResponseBody {
	s.Features = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetInscriptionsAndMarkings(v string) *QueryArtExtensionResponseBody {
	s.InscriptionsAndMarkings = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetMaker(v string) *QueryArtExtensionResponseBody {
	s.Maker = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetMaterialsAndTechniques(v string) *QueryArtExtensionResponseBody {
	s.MaterialsAndTechniques = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetObjectType(v string) *QueryArtExtensionResponseBody {
	s.ObjectType = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetReference(v string) *QueryArtExtensionResponseBody {
	s.Reference = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetRequestId(v string) *QueryArtExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetSubject(v string) *QueryArtExtensionResponseBody {
	s.Subject = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetTitle(v string) *QueryArtExtensionResponseBody {
	s.Title = &v
	return s
}

func (s *QueryArtExtensionResponseBody) Validate() error {
	return dara.Validate(s)
}
