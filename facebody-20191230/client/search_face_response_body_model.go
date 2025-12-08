// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SearchFaceResponseBodyData) *SearchFaceResponseBody
	GetData() *SearchFaceResponseBodyData
	SetRequestId(v string) *SearchFaceResponseBody
	GetRequestId() *string
}

type SearchFaceResponseBody struct {
	Data *SearchFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 4159e64a-0fe8-436c-a8de-ee531555db3c
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchFaceResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBody) GetData() *SearchFaceResponseBodyData {
	return s.Data
}

func (s *SearchFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchFaceResponseBody) SetData(v *SearchFaceResponseBodyData) *SearchFaceResponseBody {
	s.Data = v
	return s
}

func (s *SearchFaceResponseBody) SetRequestId(v string) *SearchFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchFaceResponseBodyData struct {
	MatchList []*SearchFaceResponseBodyDataMatchList `json:"MatchList,omitempty" xml:"MatchList,omitempty" type:"Repeated"`
}

func (s SearchFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyData) GetMatchList() []*SearchFaceResponseBodyDataMatchList {
	return s.MatchList
}

func (s *SearchFaceResponseBodyData) SetMatchList(v []*SearchFaceResponseBodyDataMatchList) *SearchFaceResponseBodyData {
	s.MatchList = v
	return s
}

func (s *SearchFaceResponseBodyData) Validate() error {
	if s.MatchList != nil {
		for _, item := range s.MatchList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchFaceResponseBodyDataMatchList struct {
	FaceItems []*SearchFaceResponseBodyDataMatchListFaceItems `json:"FaceItems,omitempty" xml:"FaceItems,omitempty" type:"Repeated"`
	Location  *SearchFaceResponseBodyDataMatchListLocation    `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// example:
	//
	// 71.7349
	QualitieScore *float32 `json:"QualitieScore,omitempty" xml:"QualitieScore,omitempty"`
}

func (s SearchFaceResponseBodyDataMatchList) String() string {
	return dara.Prettify(s)
}

func (s SearchFaceResponseBodyDataMatchList) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyDataMatchList) GetFaceItems() []*SearchFaceResponseBodyDataMatchListFaceItems {
	return s.FaceItems
}

func (s *SearchFaceResponseBodyDataMatchList) GetLocation() *SearchFaceResponseBodyDataMatchListLocation {
	return s.Location
}

func (s *SearchFaceResponseBodyDataMatchList) GetQualitieScore() *float32 {
	return s.QualitieScore
}

func (s *SearchFaceResponseBodyDataMatchList) SetFaceItems(v []*SearchFaceResponseBodyDataMatchListFaceItems) *SearchFaceResponseBodyDataMatchList {
	s.FaceItems = v
	return s
}

func (s *SearchFaceResponseBodyDataMatchList) SetLocation(v *SearchFaceResponseBodyDataMatchListLocation) *SearchFaceResponseBodyDataMatchList {
	s.Location = v
	return s
}

func (s *SearchFaceResponseBodyDataMatchList) SetQualitieScore(v float32) *SearchFaceResponseBodyDataMatchList {
	s.QualitieScore = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchList) Validate() error {
	if s.FaceItems != nil {
		for _, item := range s.FaceItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchFaceResponseBodyDataMatchListFaceItems struct {
	// example:
	//
	// 36.820168
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// U1
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// example:
	//
	// 001
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	// example:
	//
	// 0.892133
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s SearchFaceResponseBodyDataMatchListFaceItems) String() string {
	return dara.Prettify(s)
}

func (s SearchFaceResponseBodyDataMatchListFaceItems) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) GetConfidence() *float32 {
	return s.Confidence
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) GetDbName() *string {
	return s.DbName
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) GetEntityId() *string {
	return s.EntityId
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) GetExtraData() *string {
	return s.ExtraData
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) GetFaceId() *string {
	return s.FaceId
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) GetScore() *float32 {
	return s.Score
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetConfidence(v float32) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.Confidence = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetDbName(v string) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.DbName = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetEntityId(v string) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.EntityId = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetExtraData(v string) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.ExtraData = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetFaceId(v string) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.FaceId = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetScore(v float32) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.Score = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) Validate() error {
	return dara.Validate(s)
}

type SearchFaceResponseBodyDataMatchListLocation struct {
	// example:
	//
	// 200
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 200
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 5
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 6
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SearchFaceResponseBodyDataMatchListLocation) String() string {
	return dara.Prettify(s)
}

func (s SearchFaceResponseBodyDataMatchListLocation) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyDataMatchListLocation) GetHeight() *int32 {
	return s.Height
}

func (s *SearchFaceResponseBodyDataMatchListLocation) GetWidth() *int32 {
	return s.Width
}

func (s *SearchFaceResponseBodyDataMatchListLocation) GetX() *int32 {
	return s.X
}

func (s *SearchFaceResponseBodyDataMatchListLocation) GetY() *int32 {
	return s.Y
}

func (s *SearchFaceResponseBodyDataMatchListLocation) SetHeight(v int32) *SearchFaceResponseBodyDataMatchListLocation {
	s.Height = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListLocation) SetWidth(v int32) *SearchFaceResponseBodyDataMatchListLocation {
	s.Width = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListLocation) SetX(v int32) *SearchFaceResponseBodyDataMatchListLocation {
	s.X = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListLocation) SetY(v int32) *SearchFaceResponseBodyDataMatchListLocation {
	s.Y = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListLocation) Validate() error {
	return dara.Validate(s)
}
