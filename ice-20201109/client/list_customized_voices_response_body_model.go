// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomizedVoicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListCustomizedVoicesResponseBodyData) *ListCustomizedVoicesResponseBody
	GetData() *ListCustomizedVoicesResponseBodyData
	SetRequestId(v string) *ListCustomizedVoicesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCustomizedVoicesResponseBody
	GetSuccess() *bool
}

type ListCustomizedVoicesResponseBody struct {
	// The data returned.
	Data *ListCustomizedVoicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCustomizedVoicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizedVoicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomizedVoicesResponseBody) GetData() *ListCustomizedVoicesResponseBodyData {
	return s.Data
}

func (s *ListCustomizedVoicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomizedVoicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCustomizedVoicesResponseBody) SetData(v *ListCustomizedVoicesResponseBodyData) *ListCustomizedVoicesResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomizedVoicesResponseBody) SetRequestId(v string) *ListCustomizedVoicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomizedVoicesResponseBody) SetSuccess(v bool) *ListCustomizedVoicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomizedVoicesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomizedVoicesResponseBodyData struct {
	// The queried personalized human voices.
	CustomizedVoiceList []*ListCustomizedVoicesResponseBodyDataCustomizedVoiceList `json:"CustomizedVoiceList,omitempty" xml:"CustomizedVoiceList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 41
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomizedVoicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizedVoicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomizedVoicesResponseBodyData) GetCustomizedVoiceList() []*ListCustomizedVoicesResponseBodyDataCustomizedVoiceList {
	return s.CustomizedVoiceList
}

func (s *ListCustomizedVoicesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomizedVoicesResponseBodyData) SetCustomizedVoiceList(v []*ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) *ListCustomizedVoicesResponseBodyData {
	s.CustomizedVoiceList = v
	return s
}

func (s *ListCustomizedVoicesResponseBodyData) SetTotalCount(v int32) *ListCustomizedVoicesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCustomizedVoicesResponseBodyData) Validate() error {
	if s.CustomizedVoiceList != nil {
		for _, item := range s.CustomizedVoiceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomizedVoicesResponseBodyDataCustomizedVoiceList struct {
	// The media asset ID of the sample audio file.
	//
	// example:
	//
	// ****4d5e829d498aaf966b119348****
	DemoAudioMediaId *string `json:"DemoAudioMediaId,omitempty" xml:"DemoAudioMediaId,omitempty"`
	// The gender. Valid values:
	//
	// 	- female
	//
	// 	- male
	//
	// example:
	//
	// male
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// The scenario. Valid values:
	//
	// 	- story
	//
	// 	- interaction
	//
	// 	- navigation
	//
	// example:
	//
	// story
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// 	- The voice type. Valid values:
	//
	//     	- Basic
	//
	//     	- Standard
	//
	// example:
	//
	// Standard
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The voice description.
	VoiceDesc *string `json:"VoiceDesc,omitempty" xml:"VoiceDesc,omitempty"`
	// The voice ID.
	//
	// example:
	//
	// xiaozhuan
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// The voice name.
	VoiceName *string `json:"VoiceName,omitempty" xml:"VoiceName,omitempty"`
}

func (s ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) GoString() string {
	return s.String()
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) GetDemoAudioMediaId() *string {
	return s.DemoAudioMediaId
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) GetGender() *string {
	return s.Gender
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) GetScenario() *string {
	return s.Scenario
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) GetType() *string {
	return s.Type
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) GetVoiceDesc() *string {
	return s.VoiceDesc
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) GetVoiceId() *string {
	return s.VoiceId
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) GetVoiceName() *string {
	return s.VoiceName
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) SetDemoAudioMediaId(v string) *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList {
	s.DemoAudioMediaId = &v
	return s
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) SetGender(v string) *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList {
	s.Gender = &v
	return s
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) SetScenario(v string) *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList {
	s.Scenario = &v
	return s
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) SetType(v string) *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList {
	s.Type = &v
	return s
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) SetVoiceDesc(v string) *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList {
	s.VoiceDesc = &v
	return s
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) SetVoiceId(v string) *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList {
	s.VoiceId = &v
	return s
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) SetVoiceName(v string) *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList {
	s.VoiceName = &v
	return s
}

func (s *ListCustomizedVoicesResponseBodyDataCustomizedVoiceList) Validate() error {
	return dara.Validate(s)
}
