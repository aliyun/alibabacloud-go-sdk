// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByHybridResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchMediaByHybridResponseBody
	GetCode() *string
	SetMediaList(v []*SearchMediaByHybridResponseBodyMediaList) *SearchMediaByHybridResponseBody
	GetMediaList() []*SearchMediaByHybridResponseBodyMediaList
	SetRequestId(v string) *SearchMediaByHybridResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SearchMediaByHybridResponseBody
	GetSuccess() *string
	SetTotal(v int64) *SearchMediaByHybridResponseBody
	GetTotal() *int64
}

type SearchMediaByHybridResponseBody struct {
	// example:
	//
	// 200
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	MediaList []*SearchMediaByHybridResponseBodyMediaList `json:"MediaList,omitempty" xml:"MediaList,omitempty" type:"Repeated"`
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchMediaByHybridResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByHybridResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMediaByHybridResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchMediaByHybridResponseBody) GetMediaList() []*SearchMediaByHybridResponseBodyMediaList {
	return s.MediaList
}

func (s *SearchMediaByHybridResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMediaByHybridResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SearchMediaByHybridResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *SearchMediaByHybridResponseBody) SetCode(v string) *SearchMediaByHybridResponseBody {
	s.Code = &v
	return s
}

func (s *SearchMediaByHybridResponseBody) SetMediaList(v []*SearchMediaByHybridResponseBodyMediaList) *SearchMediaByHybridResponseBody {
	s.MediaList = v
	return s
}

func (s *SearchMediaByHybridResponseBody) SetRequestId(v string) *SearchMediaByHybridResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMediaByHybridResponseBody) SetSuccess(v string) *SearchMediaByHybridResponseBody {
	s.Success = &v
	return s
}

func (s *SearchMediaByHybridResponseBody) SetTotal(v int64) *SearchMediaByHybridResponseBody {
	s.Total = &v
	return s
}

func (s *SearchMediaByHybridResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchMediaByHybridResponseBodyMediaList struct {
	ClipInfo []*SearchMediaByHybridResponseBodyMediaListClipInfo `json:"ClipInfo,omitempty" xml:"ClipInfo,omitempty" type:"Repeated"`
	// example:
	//
	// a18936e0e28771edb59ae6f6f47a****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s SearchMediaByHybridResponseBodyMediaList) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByHybridResponseBodyMediaList) GoString() string {
	return s.String()
}

func (s *SearchMediaByHybridResponseBodyMediaList) GetClipInfo() []*SearchMediaByHybridResponseBodyMediaListClipInfo {
	return s.ClipInfo
}

func (s *SearchMediaByHybridResponseBodyMediaList) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaByHybridResponseBodyMediaList) SetClipInfo(v []*SearchMediaByHybridResponseBodyMediaListClipInfo) *SearchMediaByHybridResponseBodyMediaList {
	s.ClipInfo = v
	return s
}

func (s *SearchMediaByHybridResponseBodyMediaList) SetMediaId(v string) *SearchMediaByHybridResponseBodyMediaList {
	s.MediaId = &v
	return s
}

func (s *SearchMediaByHybridResponseBodyMediaList) Validate() error {
	return dara.Validate(s)
}

type SearchMediaByHybridResponseBodyMediaListClipInfo struct {
	// example:
	//
	// 2
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 0.99
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 4
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SearchMediaByHybridResponseBodyMediaListClipInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByHybridResponseBodyMediaListClipInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaByHybridResponseBodyMediaListClipInfo) GetFrom() *float64 {
	return s.From
}

func (s *SearchMediaByHybridResponseBodyMediaListClipInfo) GetScore() *float64 {
	return s.Score
}

func (s *SearchMediaByHybridResponseBodyMediaListClipInfo) GetTo() *float64 {
	return s.To
}

func (s *SearchMediaByHybridResponseBodyMediaListClipInfo) SetFrom(v float64) *SearchMediaByHybridResponseBodyMediaListClipInfo {
	s.From = &v
	return s
}

func (s *SearchMediaByHybridResponseBodyMediaListClipInfo) SetScore(v float64) *SearchMediaByHybridResponseBodyMediaListClipInfo {
	s.Score = &v
	return s
}

func (s *SearchMediaByHybridResponseBodyMediaListClipInfo) SetTo(v float64) *SearchMediaByHybridResponseBodyMediaListClipInfo {
	s.To = &v
	return s
}

func (s *SearchMediaByHybridResponseBodyMediaListClipInfo) Validate() error {
	return dara.Validate(s)
}
