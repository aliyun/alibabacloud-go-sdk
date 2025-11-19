// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIVideoTagResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAIVideoTagResultResponseBody
	GetRequestId() *string
	SetVideoTagResult(v *GetAIVideoTagResultResponseBodyVideoTagResult) *GetAIVideoTagResultResponseBody
	GetVideoTagResult() *GetAIVideoTagResultResponseBodyVideoTagResult
}

type GetAIVideoTagResultResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8829B4DB-AFD9-4FF6-12965DBFFA14****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	VideoTagResult *GetAIVideoTagResultResponseBodyVideoTagResult `json:"VideoTagResult,omitempty" xml:"VideoTagResult,omitempty" type:"Struct"`
}

func (s GetAIVideoTagResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAIVideoTagResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAIVideoTagResultResponseBody) GetVideoTagResult() *GetAIVideoTagResultResponseBodyVideoTagResult {
	return s.VideoTagResult
}

func (s *GetAIVideoTagResultResponseBody) SetRequestId(v string) *GetAIVideoTagResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIVideoTagResultResponseBody) SetVideoTagResult(v *GetAIVideoTagResultResponseBodyVideoTagResult) *GetAIVideoTagResultResponseBody {
	s.VideoTagResult = v
	return s
}

func (s *GetAIVideoTagResultResponseBody) Validate() error {
	if s.VideoTagResult != nil {
		if err := s.VideoTagResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAIVideoTagResultResponseBodyVideoTagResult struct {
	// The video categories.
	Category []*GetAIVideoTagResultResponseBodyVideoTagResultCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Repeated"`
	// The keyword tags.
	Keyword []*GetAIVideoTagResultResponseBodyVideoTagResultKeyword `json:"Keyword,omitempty" xml:"Keyword,omitempty" type:"Repeated"`
	// The location tags.
	Location []*GetAIVideoTagResultResponseBodyVideoTagResultLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
	// The figure tags.
	Person []*GetAIVideoTagResultResponseBodyVideoTagResultPerson `json:"Person,omitempty" xml:"Person,omitempty" type:"Repeated"`
	// The time tags.
	Time []*GetAIVideoTagResultResponseBodyVideoTagResultTime `json:"Time,omitempty" xml:"Time,omitempty" type:"Repeated"`
}

func (s GetAIVideoTagResultResponseBodyVideoTagResult) String() string {
	return dara.Prettify(s)
}

func (s GetAIVideoTagResultResponseBodyVideoTagResult) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) GetCategory() []*GetAIVideoTagResultResponseBodyVideoTagResultCategory {
	return s.Category
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) GetKeyword() []*GetAIVideoTagResultResponseBodyVideoTagResultKeyword {
	return s.Keyword
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) GetLocation() []*GetAIVideoTagResultResponseBodyVideoTagResultLocation {
	return s.Location
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) GetPerson() []*GetAIVideoTagResultResponseBodyVideoTagResultPerson {
	return s.Person
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) GetTime() []*GetAIVideoTagResultResponseBodyVideoTagResultTime {
	return s.Time
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) SetCategory(v []*GetAIVideoTagResultResponseBodyVideoTagResultCategory) *GetAIVideoTagResultResponseBodyVideoTagResult {
	s.Category = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) SetKeyword(v []*GetAIVideoTagResultResponseBodyVideoTagResultKeyword) *GetAIVideoTagResultResponseBodyVideoTagResult {
	s.Keyword = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) SetLocation(v []*GetAIVideoTagResultResponseBodyVideoTagResultLocation) *GetAIVideoTagResultResponseBodyVideoTagResult {
	s.Location = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) SetPerson(v []*GetAIVideoTagResultResponseBodyVideoTagResultPerson) *GetAIVideoTagResultResponseBodyVideoTagResult {
	s.Person = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) SetTime(v []*GetAIVideoTagResultResponseBodyVideoTagResultTime) *GetAIVideoTagResultResponseBodyVideoTagResult {
	s.Time = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) Validate() error {
	if s.Category != nil {
		for _, item := range s.Category {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Keyword != nil {
		for _, item := range s.Keyword {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Location != nil {
		for _, item := range s.Location {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Person != nil {
		for _, item := range s.Person {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Time != nil {
		for _, item := range s.Time {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAIVideoTagResultResponseBodyVideoTagResultCategory struct {
	// The tag string.
	//
	// example:
	//
	// Retouching
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultCategory) String() string {
	return dara.Prettify(s)
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultCategory) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultCategory) GetTag() *string {
	return s.Tag
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultCategory) SetTag(v string) *GetAIVideoTagResultResponseBodyVideoTagResultCategory {
	s.Tag = &v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultCategory) Validate() error {
	return dara.Validate(s)
}

type GetAIVideoTagResultResponseBodyVideoTagResultKeyword struct {
	// The tag string.
	//
	// example:
	//
	// Cushion
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The points in time when the tags are displayed. Unit: milliseconds.
	Times []*string `json:"Times,omitempty" xml:"Times,omitempty" type:"Repeated"`
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultKeyword) String() string {
	return dara.Prettify(s)
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultKeyword) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultKeyword) GetTag() *string {
	return s.Tag
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultKeyword) GetTimes() []*string {
	return s.Times
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultKeyword) SetTag(v string) *GetAIVideoTagResultResponseBodyVideoTagResultKeyword {
	s.Tag = &v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultKeyword) SetTimes(v []*string) *GetAIVideoTagResultResponseBodyVideoTagResultKeyword {
	s.Times = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultKeyword) Validate() error {
	return dara.Validate(s)
}

type GetAIVideoTagResultResponseBodyVideoTagResultLocation struct {
	// The tag string.
	//
	// example:
	//
	// Asia
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The points in time when the tags are displayed. Unit: milliseconds.
	Times []*string `json:"Times,omitempty" xml:"Times,omitempty" type:"Repeated"`
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultLocation) String() string {
	return dara.Prettify(s)
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultLocation) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultLocation) GetTag() *string {
	return s.Tag
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultLocation) GetTimes() []*string {
	return s.Times
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultLocation) SetTag(v string) *GetAIVideoTagResultResponseBodyVideoTagResultLocation {
	s.Tag = &v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultLocation) SetTimes(v []*string) *GetAIVideoTagResultResponseBodyVideoTagResultLocation {
	s.Times = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultLocation) Validate() error {
	return dara.Validate(s)
}

type GetAIVideoTagResultResponseBodyVideoTagResultPerson struct {
	// The URL of the profile photo.
	//
	// > This parameter is returned only when a figure tag was used.
	//
	// example:
	//
	// http://example.com/aivideotag/8829B4DB-AFD9-4F*****F6-12965DBFFA14/Index_****.jpg
	FaceUrl *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	// The tag string.
	//
	// example:
	//
	// John
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The points in time when the tags are displayed. Unit: milliseconds.
	Times []*string `json:"Times,omitempty" xml:"Times,omitempty" type:"Repeated"`
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultPerson) String() string {
	return dara.Prettify(s)
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultPerson) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultPerson) GetFaceUrl() *string {
	return s.FaceUrl
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultPerson) GetTag() *string {
	return s.Tag
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultPerson) GetTimes() []*string {
	return s.Times
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultPerson) SetFaceUrl(v string) *GetAIVideoTagResultResponseBodyVideoTagResultPerson {
	s.FaceUrl = &v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultPerson) SetTag(v string) *GetAIVideoTagResultResponseBodyVideoTagResultPerson {
	s.Tag = &v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultPerson) SetTimes(v []*string) *GetAIVideoTagResultResponseBodyVideoTagResultPerson {
	s.Times = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultPerson) Validate() error {
	return dara.Validate(s)
}

type GetAIVideoTagResultResponseBodyVideoTagResultTime struct {
	// The tag string.
	//
	// example:
	//
	// Milliseconds
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The points in time when the tags are displayed. Unit: milliseconds.
	Times []*string `json:"Times,omitempty" xml:"Times,omitempty" type:"Repeated"`
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultTime) String() string {
	return dara.Prettify(s)
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultTime) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultTime) GetTag() *string {
	return s.Tag
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultTime) GetTimes() []*string {
	return s.Times
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultTime) SetTag(v string) *GetAIVideoTagResultResponseBodyVideoTagResultTime {
	s.Tag = &v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultTime) SetTimes(v []*string) *GetAIVideoTagResultResponseBodyVideoTagResultTime {
	s.Times = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultTime) Validate() error {
	return dara.Validate(s)
}
