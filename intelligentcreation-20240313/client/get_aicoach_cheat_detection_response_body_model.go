// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachCheatDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheatId(v string) *GetAICoachCheatDetectionResponseBody
	GetCheatId() *string
	SetErrorCode(v string) *GetAICoachCheatDetectionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAICoachCheatDetectionResponseBody
	GetErrorMessage() *string
	SetGmtCreate(v string) *GetAICoachCheatDetectionResponseBody
	GetGmtCreate() *string
	SetImageCheat(v *GetAICoachCheatDetectionResponseBodyImageCheat) *GetAICoachCheatDetectionResponseBody
	GetImageCheat() *GetAICoachCheatDetectionResponseBodyImageCheat
	SetRequestId(v string) *GetAICoachCheatDetectionResponseBody
	GetRequestId() *string
	SetStatus(v int32) *GetAICoachCheatDetectionResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *GetAICoachCheatDetectionResponseBody
	GetSuccess() *bool
	SetVoiceCheat(v *GetAICoachCheatDetectionResponseBodyVoiceCheat) *GetAICoachCheatDetectionResponseBody
	GetVoiceCheat() *GetAICoachCheatDetectionResponseBodyVoiceCheat
}

type GetAICoachCheatDetectionResponseBody struct {
	// example:
	//
	// 1
	CheatId *string `json:"cheatId,omitempty" xml:"cheatId,omitempty"`
	// example:
	//
	// success
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 2025-02-24 12:00:00
	GmtCreate  *string                                         `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	ImageCheat *GetAICoachCheatDetectionResponseBodyImageCheat `json:"imageCheat,omitempty" xml:"imageCheat,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true
	//
	// example:
	//
	// True
	Success    *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
	VoiceCheat *GetAICoachCheatDetectionResponseBodyVoiceCheat `json:"voiceCheat,omitempty" xml:"voiceCheat,omitempty" type:"Struct"`
}

func (s GetAICoachCheatDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachCheatDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponseBody) GetCheatId() *string {
	return s.CheatId
}

func (s *GetAICoachCheatDetectionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAICoachCheatDetectionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAICoachCheatDetectionResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAICoachCheatDetectionResponseBody) GetImageCheat() *GetAICoachCheatDetectionResponseBodyImageCheat {
	return s.ImageCheat
}

func (s *GetAICoachCheatDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAICoachCheatDetectionResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *GetAICoachCheatDetectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAICoachCheatDetectionResponseBody) GetVoiceCheat() *GetAICoachCheatDetectionResponseBodyVoiceCheat {
	return s.VoiceCheat
}

func (s *GetAICoachCheatDetectionResponseBody) SetCheatId(v string) *GetAICoachCheatDetectionResponseBody {
	s.CheatId = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetErrorCode(v string) *GetAICoachCheatDetectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetErrorMessage(v string) *GetAICoachCheatDetectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetGmtCreate(v string) *GetAICoachCheatDetectionResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetImageCheat(v *GetAICoachCheatDetectionResponseBodyImageCheat) *GetAICoachCheatDetectionResponseBody {
	s.ImageCheat = v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetRequestId(v string) *GetAICoachCheatDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetStatus(v int32) *GetAICoachCheatDetectionResponseBody {
	s.Status = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetSuccess(v bool) *GetAICoachCheatDetectionResponseBody {
	s.Success = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetVoiceCheat(v *GetAICoachCheatDetectionResponseBodyVoiceCheat) *GetAICoachCheatDetectionResponseBody {
	s.VoiceCheat = v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) Validate() error {
	if s.ImageCheat != nil {
		if err := s.ImageCheat.Validate(); err != nil {
			return err
		}
	}
	if s.VoiceCheat != nil {
		if err := s.VoiceCheat.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAICoachCheatDetectionResponseBodyImageCheat struct {
	// example:
	//
	// demo
	Desc *string                                               `json:"desc,omitempty" xml:"desc,omitempty"`
	List []*GetAICoachCheatDetectionResponseBodyImageCheatList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetAICoachCheatDetectionResponseBodyImageCheat) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachCheatDetectionResponseBodyImageCheat) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheat) GetDesc() *string {
	return s.Desc
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheat) GetList() []*GetAICoachCheatDetectionResponseBodyImageCheatList {
	return s.List
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheat) GetStatus() *int32 {
	return s.Status
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheat) SetDesc(v string) *GetAICoachCheatDetectionResponseBodyImageCheat {
	s.Desc = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheat) SetList(v []*GetAICoachCheatDetectionResponseBodyImageCheatList) *GetAICoachCheatDetectionResponseBodyImageCheat {
	s.List = v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheat) SetStatus(v int32) *GetAICoachCheatDetectionResponseBodyImageCheat {
	s.Status = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheat) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAICoachCheatDetectionResponseBodyImageCheatList struct {
	// example:
	//
	// 2025-03-22 10:05:07
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// https://demo.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAICoachCheatDetectionResponseBodyImageCheatList) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachCheatDetectionResponseBodyImageCheatList) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheatList) GetTime() *string {
	return s.Time
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheatList) GetUrl() *string {
	return s.Url
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheatList) SetTime(v string) *GetAICoachCheatDetectionResponseBodyImageCheatList {
	s.Time = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheatList) SetUrl(v string) *GetAICoachCheatDetectionResponseBodyImageCheatList {
	s.Url = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheatList) Validate() error {
	return dara.Validate(s)
}

type GetAICoachCheatDetectionResponseBodyVoiceCheat struct {
	ComparisonList []*GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList `json:"comparisonList,omitempty" xml:"comparisonList,omitempty" type:"Repeated"`
	// example:
	//
	// demo
	Desc         *string                                                       `json:"desc,omitempty" xml:"desc,omitempty"`
	OriginalList []*GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList `json:"originalList,omitempty" xml:"originalList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetAICoachCheatDetectionResponseBodyVoiceCheat) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachCheatDetectionResponseBodyVoiceCheat) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) GetComparisonList() []*GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList {
	return s.ComparisonList
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) GetDesc() *string {
	return s.Desc
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) GetOriginalList() []*GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList {
	return s.OriginalList
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) GetStatus() *int32 {
	return s.Status
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) SetComparisonList(v []*GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) *GetAICoachCheatDetectionResponseBodyVoiceCheat {
	s.ComparisonList = v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) SetDesc(v string) *GetAICoachCheatDetectionResponseBodyVoiceCheat {
	s.Desc = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) SetOriginalList(v []*GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList) *GetAICoachCheatDetectionResponseBodyVoiceCheat {
	s.OriginalList = v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) SetStatus(v int32) *GetAICoachCheatDetectionResponseBodyVoiceCheat {
	s.Status = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) Validate() error {
	if s.ComparisonList != nil {
		for _, item := range s.ComparisonList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OriginalList != nil {
		for _, item := range s.OriginalList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList struct {
	// example:
	//
	// 2024-12-11 10:07:23
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// https://demo.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) GetTime() *string {
	return s.Time
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) GetUrl() *string {
	return s.Url
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) SetTime(v string) *GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList {
	s.Time = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) SetUrl(v string) *GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList {
	s.Url = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) Validate() error {
	return dara.Validate(s)
}

type GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList struct {
	// example:
	//
	// https://demo.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList) GetUrl() *string {
	return s.Url
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList) SetUrl(v string) *GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList {
	s.Url = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList) Validate() error {
	return dara.Validate(s)
}
