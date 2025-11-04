// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaClipByFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchMediaClipByFaceResponseBody
	GetCode() *string
	SetMediaClipList(v []*SearchMediaClipByFaceResponseBodyMediaClipList) *SearchMediaClipByFaceResponseBody
	GetMediaClipList() []*SearchMediaClipByFaceResponseBodyMediaClipList
	SetRequestId(v string) *SearchMediaClipByFaceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SearchMediaClipByFaceResponseBody
	GetSuccess() *string
	SetTotal(v int64) *SearchMediaClipByFaceResponseBody
	GetTotal() *int64
}

type SearchMediaClipByFaceResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The media asset clips that meet the requirements.
	MediaClipList []*SearchMediaClipByFaceResponseBodyMediaClipList `json:"MediaClipList,omitempty" xml:"MediaClipList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// E44FFACD-9E90-555A-A09A-6FD3B7335E39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of media asset clips that meet the conditions.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchMediaClipByFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaClipByFaceResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMediaClipByFaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchMediaClipByFaceResponseBody) GetMediaClipList() []*SearchMediaClipByFaceResponseBodyMediaClipList {
	return s.MediaClipList
}

func (s *SearchMediaClipByFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMediaClipByFaceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SearchMediaClipByFaceResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *SearchMediaClipByFaceResponseBody) SetCode(v string) *SearchMediaClipByFaceResponseBody {
	s.Code = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBody) SetMediaClipList(v []*SearchMediaClipByFaceResponseBodyMediaClipList) *SearchMediaClipByFaceResponseBody {
	s.MediaClipList = v
	return s
}

func (s *SearchMediaClipByFaceResponseBody) SetRequestId(v string) *SearchMediaClipByFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBody) SetSuccess(v string) *SearchMediaClipByFaceResponseBody {
	s.Success = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBody) SetTotal(v int64) *SearchMediaClipByFaceResponseBody {
	s.Total = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBody) Validate() error {
	if s.MediaClipList != nil {
		for _, item := range s.MediaClipList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMediaClipByFaceResponseBodyMediaClipList struct {
	// The type of the character. Valid values: celebrity sensitive politician custom unknown
	//
	// example:
	//
	// celebrity
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the entity, which is the same as the entity ID returned in tag analysis.
	//
	// example:
	//
	// 1031025****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The name of the entity.
	LabelName *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	// The information about clips related to the face.
	OccurrencesInfos []*SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos `json:"OccurrencesInfos,omitempty" xml:"OccurrencesInfos,omitempty" type:"Repeated"`
	// The score of the clip. The value is of the Float type. The value is in the range of [0,1].
	//
	// example:
	//
	// 0.99041677
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s SearchMediaClipByFaceResponseBodyMediaClipList) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaClipByFaceResponseBodyMediaClipList) GoString() string {
	return s.String()
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipList) GetCategory() *string {
	return s.Category
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipList) GetEntityId() *string {
	return s.EntityId
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipList) GetLabelName() *string {
	return s.LabelName
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipList) GetOccurrencesInfos() []*SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos {
	return s.OccurrencesInfos
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipList) GetScore() *float32 {
	return s.Score
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipList) SetCategory(v string) *SearchMediaClipByFaceResponseBodyMediaClipList {
	s.Category = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipList) SetEntityId(v string) *SearchMediaClipByFaceResponseBodyMediaClipList {
	s.EntityId = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipList) SetLabelName(v string) *SearchMediaClipByFaceResponseBodyMediaClipList {
	s.LabelName = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipList) SetOccurrencesInfos(v []*SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos) *SearchMediaClipByFaceResponseBodyMediaClipList {
	s.OccurrencesInfos = v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipList) SetScore(v float32) *SearchMediaClipByFaceResponseBodyMediaClipList {
	s.Score = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipList) Validate() error {
	if s.OccurrencesInfos != nil {
		for _, item := range s.OccurrencesInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos struct {
	// The end time of the clip. Unit: seconds. The value is of the Float type.
	//
	// example:
	//
	// 69.06635
	EndTime    *float32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Expression *string  `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The start time of the clip. Unit: seconds. The value is of the Float type.
	//
	// example:
	//
	// 61.066353
	StartTime *float32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The information about the face in the clip.
	TrackData []*SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData `json:"TrackData,omitempty" xml:"TrackData,omitempty" type:"Repeated"`
}

func (s SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos) GoString() string {
	return s.String()
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos) GetEndTime() *float32 {
	return s.EndTime
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos) GetExpression() *string {
	return s.Expression
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos) GetStartTime() *float32 {
	return s.StartTime
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos) GetTrackData() []*SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData {
	return s.TrackData
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos) SetEndTime(v float32) *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos {
	s.EndTime = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos) SetExpression(v string) *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos {
	s.Expression = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos) SetStartTime(v float32) *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos {
	s.StartTime = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos) SetTrackData(v []*SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData) *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos {
	s.TrackData = v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfos) Validate() error {
	if s.TrackData != nil {
		for _, item := range s.TrackData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData struct {
	// The coordinates of the face.
	BoxPosition *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition `json:"BoxPosition,omitempty" xml:"BoxPosition,omitempty" type:"Struct"`
	// The timestamp when the face appears in the clip. Unit: seconds. The value is of the Float type.
	//
	// example:
	//
	// 62.03302
	Timestamp *float32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData) GoString() string {
	return s.String()
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData) GetBoxPosition() *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition {
	return s.BoxPosition
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData) GetTimestamp() *float32 {
	return s.Timestamp
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData) SetBoxPosition(v *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition) *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData {
	s.BoxPosition = v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData) SetTimestamp(v float32) *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData {
	s.Timestamp = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackData) Validate() error {
	if s.BoxPosition != nil {
		if err := s.BoxPosition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition struct {
	// The height of the rectangle frame. Unit: pixels.
	//
	// example:
	//
	// 168
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the rectangle frame. Unit: pixels.
	//
	// example:
	//
	// 128
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The x-axis coordinate of the upper-left corner. Unit: pixels.
	//
	// example:
	//
	// 517
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-axis coordinate of the upper-left corner. Unit: pixels.
	//
	// example:
	//
	// 409
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition) GoString() string {
	return s.String()
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition) GetH() *int32 {
	return s.H
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition) GetW() *int32 {
	return s.W
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition) GetX() *int32 {
	return s.X
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition) GetY() *int32 {
	return s.Y
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition) SetH(v int32) *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition {
	s.H = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition) SetW(v int32) *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition {
	s.W = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition) SetX(v int32) *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition {
	s.X = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition) SetY(v int32) *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition {
	s.Y = &v
	return s
}

func (s *SearchMediaClipByFaceResponseBodyMediaClipListOccurrencesInfosTrackDataBoxPosition) Validate() error {
	return dara.Validate(s)
}
