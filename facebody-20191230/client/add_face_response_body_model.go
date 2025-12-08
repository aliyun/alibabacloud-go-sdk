// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddFaceResponseBodyData) *AddFaceResponseBody
	GetData() *AddFaceResponseBodyData
	SetRequestId(v string) *AddFaceResponseBody
	GetRequestId() *string
}

type AddFaceResponseBody struct {
	Data *AddFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 2B93C43A-F824-40C8-AF79-844342B0F43A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceResponseBody) GetData() *AddFaceResponseBodyData {
	return s.Data
}

func (s *AddFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFaceResponseBody) SetData(v *AddFaceResponseBodyData) *AddFaceResponseBody {
	s.Data = v
	return s
}

func (s *AddFaceResponseBody) SetRequestId(v string) *AddFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFaceResponseBodyData struct {
	// example:
	//
	// 5
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	// example:
	//
	// 99.79581
	QualitieScore *float32 `json:"QualitieScore,omitempty" xml:"QualitieScore,omitempty"`
}

func (s AddFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceResponseBodyData) GetFaceId() *string {
	return s.FaceId
}

func (s *AddFaceResponseBodyData) GetQualitieScore() *float32 {
	return s.QualitieScore
}

func (s *AddFaceResponseBodyData) SetFaceId(v string) *AddFaceResponseBodyData {
	s.FaceId = &v
	return s
}

func (s *AddFaceResponseBodyData) SetQualitieScore(v float32) *AddFaceResponseBodyData {
	s.QualitieScore = &v
	return s
}

func (s *AddFaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
