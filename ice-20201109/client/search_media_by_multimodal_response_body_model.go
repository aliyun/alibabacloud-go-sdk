// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByMultimodalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchMediaByMultimodalResponseBody
	GetCode() *string
	SetMediaList(v []*SearchMediaByMultimodalResponseBodyMediaList) *SearchMediaByMultimodalResponseBody
	GetMediaList() []*SearchMediaByMultimodalResponseBodyMediaList
	SetRequestId(v string) *SearchMediaByMultimodalResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SearchMediaByMultimodalResponseBody
	GetSuccess() *string
	SetTotal(v int64) *SearchMediaByMultimodalResponseBody
	GetTotal() *int64
}

type SearchMediaByMultimodalResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The media assets that contain the specified content.
	MediaList []*SearchMediaByMultimodalResponseBodyMediaList `json:"MediaList,omitempty" xml:"MediaList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of data records that meet the specified filter condition.
	//
	// example:
	//
	// 20
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchMediaByMultimodalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByMultimodalResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMediaByMultimodalResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchMediaByMultimodalResponseBody) GetMediaList() []*SearchMediaByMultimodalResponseBodyMediaList {
	return s.MediaList
}

func (s *SearchMediaByMultimodalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMediaByMultimodalResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SearchMediaByMultimodalResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *SearchMediaByMultimodalResponseBody) SetCode(v string) *SearchMediaByMultimodalResponseBody {
	s.Code = &v
	return s
}

func (s *SearchMediaByMultimodalResponseBody) SetMediaList(v []*SearchMediaByMultimodalResponseBodyMediaList) *SearchMediaByMultimodalResponseBody {
	s.MediaList = v
	return s
}

func (s *SearchMediaByMultimodalResponseBody) SetRequestId(v string) *SearchMediaByMultimodalResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMediaByMultimodalResponseBody) SetSuccess(v string) *SearchMediaByMultimodalResponseBody {
	s.Success = &v
	return s
}

func (s *SearchMediaByMultimodalResponseBody) SetTotal(v int64) *SearchMediaByMultimodalResponseBody {
	s.Total = &v
	return s
}

func (s *SearchMediaByMultimodalResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchMediaByMultimodalResponseBodyMediaList struct {
	// The information about the clip.
	ClipInfo []*SearchMediaByMultimodalResponseBodyMediaListClipInfo `json:"ClipInfo,omitempty" xml:"ClipInfo,omitempty" type:"Repeated"`
	// The ID of the media asset.
	//
	// example:
	//
	// a18936e0e28771edb59ae6f6f47a****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s SearchMediaByMultimodalResponseBodyMediaList) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByMultimodalResponseBodyMediaList) GoString() string {
	return s.String()
}

func (s *SearchMediaByMultimodalResponseBodyMediaList) GetClipInfo() []*SearchMediaByMultimodalResponseBodyMediaListClipInfo {
	return s.ClipInfo
}

func (s *SearchMediaByMultimodalResponseBodyMediaList) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaByMultimodalResponseBodyMediaList) SetClipInfo(v []*SearchMediaByMultimodalResponseBodyMediaListClipInfo) *SearchMediaByMultimodalResponseBodyMediaList {
	s.ClipInfo = v
	return s
}

func (s *SearchMediaByMultimodalResponseBodyMediaList) SetMediaId(v string) *SearchMediaByMultimodalResponseBodyMediaList {
	s.MediaId = &v
	return s
}

func (s *SearchMediaByMultimodalResponseBodyMediaList) Validate() error {
	return dara.Validate(s)
}

type SearchMediaByMultimodalResponseBodyMediaListClipInfo struct {
	// The start time of the clip.
	//
	// example:
	//
	// 2
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// The score.
	//
	// example:
	//
	// 1.2
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The end time of the clip.
	//
	// example:
	//
	// 4
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SearchMediaByMultimodalResponseBodyMediaListClipInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByMultimodalResponseBodyMediaListClipInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaByMultimodalResponseBodyMediaListClipInfo) GetFrom() *float64 {
	return s.From
}

func (s *SearchMediaByMultimodalResponseBodyMediaListClipInfo) GetScore() *float64 {
	return s.Score
}

func (s *SearchMediaByMultimodalResponseBodyMediaListClipInfo) GetTo() *float64 {
	return s.To
}

func (s *SearchMediaByMultimodalResponseBodyMediaListClipInfo) SetFrom(v float64) *SearchMediaByMultimodalResponseBodyMediaListClipInfo {
	s.From = &v
	return s
}

func (s *SearchMediaByMultimodalResponseBodyMediaListClipInfo) SetScore(v float64) *SearchMediaByMultimodalResponseBodyMediaListClipInfo {
	s.Score = &v
	return s
}

func (s *SearchMediaByMultimodalResponseBodyMediaListClipInfo) SetTo(v float64) *SearchMediaByMultimodalResponseBodyMediaListClipInfo {
	s.To = &v
	return s
}

func (s *SearchMediaByMultimodalResponseBodyMediaListClipInfo) Validate() error {
	return dara.Validate(s)
}
