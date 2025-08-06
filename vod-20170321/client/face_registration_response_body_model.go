// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceRegistrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistImageIds(v *FaceRegistrationResponseBodyNonExistImageIds) *FaceRegistrationResponseBody
	GetNonExistImageIds() *FaceRegistrationResponseBodyNonExistImageIds
	SetRegisteredPerson(v *FaceRegistrationResponseBodyRegisteredPerson) *FaceRegistrationResponseBody
	GetRegisteredPerson() *FaceRegistrationResponseBodyRegisteredPerson
	SetRequestId(v string) *FaceRegistrationResponseBody
	GetRequestId() *string
}

type FaceRegistrationResponseBody struct {
	NonExistImageIds *FaceRegistrationResponseBodyNonExistImageIds `json:"NonExistImageIds,omitempty" xml:"NonExistImageIds,omitempty" type:"Struct"`
	RegisteredPerson *FaceRegistrationResponseBodyRegisteredPerson `json:"RegisteredPerson,omitempty" xml:"RegisteredPerson,omitempty" type:"Struct"`
	RequestId        *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FaceRegistrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FaceRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *FaceRegistrationResponseBody) GetNonExistImageIds() *FaceRegistrationResponseBodyNonExistImageIds {
	return s.NonExistImageIds
}

func (s *FaceRegistrationResponseBody) GetRegisteredPerson() *FaceRegistrationResponseBodyRegisteredPerson {
	return s.RegisteredPerson
}

func (s *FaceRegistrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FaceRegistrationResponseBody) SetNonExistImageIds(v *FaceRegistrationResponseBodyNonExistImageIds) *FaceRegistrationResponseBody {
	s.NonExistImageIds = v
	return s
}

func (s *FaceRegistrationResponseBody) SetRegisteredPerson(v *FaceRegistrationResponseBodyRegisteredPerson) *FaceRegistrationResponseBody {
	s.RegisteredPerson = v
	return s
}

func (s *FaceRegistrationResponseBody) SetRequestId(v string) *FaceRegistrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceRegistrationResponseBody) Validate() error {
	return dara.Validate(s)
}

type FaceRegistrationResponseBodyNonExistImageIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s FaceRegistrationResponseBodyNonExistImageIds) String() string {
	return dara.Prettify(s)
}

func (s FaceRegistrationResponseBodyNonExistImageIds) GoString() string {
	return s.String()
}

func (s *FaceRegistrationResponseBodyNonExistImageIds) GetString_() []*string {
	return s.String_
}

func (s *FaceRegistrationResponseBodyNonExistImageIds) SetString_(v []*string) *FaceRegistrationResponseBodyNonExistImageIds {
	s.String_ = v
	return s
}

func (s *FaceRegistrationResponseBodyNonExistImageIds) Validate() error {
	return dara.Validate(s)
}

type FaceRegistrationResponseBodyRegisteredPerson struct {
	Faces      *FaceRegistrationResponseBodyRegisteredPersonFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Struct"`
	PersonId   *string                                            `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	PersonName *string                                            `json:"PersonName,omitempty" xml:"PersonName,omitempty"`
}

func (s FaceRegistrationResponseBodyRegisteredPerson) String() string {
	return dara.Prettify(s)
}

func (s FaceRegistrationResponseBodyRegisteredPerson) GoString() string {
	return s.String()
}

func (s *FaceRegistrationResponseBodyRegisteredPerson) GetFaces() *FaceRegistrationResponseBodyRegisteredPersonFaces {
	return s.Faces
}

func (s *FaceRegistrationResponseBodyRegisteredPerson) GetPersonId() *string {
	return s.PersonId
}

func (s *FaceRegistrationResponseBodyRegisteredPerson) GetPersonName() *string {
	return s.PersonName
}

func (s *FaceRegistrationResponseBodyRegisteredPerson) SetFaces(v *FaceRegistrationResponseBodyRegisteredPersonFaces) *FaceRegistrationResponseBodyRegisteredPerson {
	s.Faces = v
	return s
}

func (s *FaceRegistrationResponseBodyRegisteredPerson) SetPersonId(v string) *FaceRegistrationResponseBodyRegisteredPerson {
	s.PersonId = &v
	return s
}

func (s *FaceRegistrationResponseBodyRegisteredPerson) SetPersonName(v string) *FaceRegistrationResponseBodyRegisteredPerson {
	s.PersonName = &v
	return s
}

func (s *FaceRegistrationResponseBodyRegisteredPerson) Validate() error {
	return dara.Validate(s)
}

type FaceRegistrationResponseBodyRegisteredPersonFaces struct {
	Face []*FaceRegistrationResponseBodyRegisteredPersonFacesFace `json:"Face,omitempty" xml:"Face,omitempty" type:"Repeated"`
}

func (s FaceRegistrationResponseBodyRegisteredPersonFaces) String() string {
	return dara.Prettify(s)
}

func (s FaceRegistrationResponseBodyRegisteredPersonFaces) GoString() string {
	return s.String()
}

func (s *FaceRegistrationResponseBodyRegisteredPersonFaces) GetFace() []*FaceRegistrationResponseBodyRegisteredPersonFacesFace {
	return s.Face
}

func (s *FaceRegistrationResponseBodyRegisteredPersonFaces) SetFace(v []*FaceRegistrationResponseBodyRegisteredPersonFacesFace) *FaceRegistrationResponseBodyRegisteredPersonFaces {
	s.Face = v
	return s
}

func (s *FaceRegistrationResponseBodyRegisteredPersonFaces) Validate() error {
	return dara.Validate(s)
}

type FaceRegistrationResponseBodyRegisteredPersonFacesFace struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Quality *string `json:"Quality,omitempty" xml:"Quality,omitempty"`
	Target  *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s FaceRegistrationResponseBodyRegisteredPersonFacesFace) String() string {
	return dara.Prettify(s)
}

func (s FaceRegistrationResponseBodyRegisteredPersonFacesFace) GoString() string {
	return s.String()
}

func (s *FaceRegistrationResponseBodyRegisteredPersonFacesFace) GetImageId() *string {
	return s.ImageId
}

func (s *FaceRegistrationResponseBodyRegisteredPersonFacesFace) GetQuality() *string {
	return s.Quality
}

func (s *FaceRegistrationResponseBodyRegisteredPersonFacesFace) GetTarget() *string {
	return s.Target
}

func (s *FaceRegistrationResponseBodyRegisteredPersonFacesFace) SetImageId(v string) *FaceRegistrationResponseBodyRegisteredPersonFacesFace {
	s.ImageId = &v
	return s
}

func (s *FaceRegistrationResponseBodyRegisteredPersonFacesFace) SetQuality(v string) *FaceRegistrationResponseBodyRegisteredPersonFacesFace {
	s.Quality = &v
	return s
}

func (s *FaceRegistrationResponseBodyRegisteredPersonFacesFace) SetTarget(v string) *FaceRegistrationResponseBodyRegisteredPersonFacesFace {
	s.Target = &v
	return s
}

func (s *FaceRegistrationResponseBodyRegisteredPersonFacesFace) Validate() error {
	return dara.Validate(s)
}
