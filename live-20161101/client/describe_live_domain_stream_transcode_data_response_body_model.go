// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainStreamTranscodeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveDomainStreamTranscodeDataResponseBody
	GetRequestId() *string
	SetTranscodeDataList(v *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataList) *DescribeLiveDomainStreamTranscodeDataResponseBody
	GetTranscodeDataList() *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataList
}

type DescribeLiveDomainStreamTranscodeDataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId         *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranscodeDataList *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataList `json:"TranscodeDataList,omitempty" xml:"TranscodeDataList,omitempty" type:"Struct"`
}

func (s DescribeLiveDomainStreamTranscodeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainStreamTranscodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBody) GetTranscodeDataList() *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataList {
	return s.TranscodeDataList
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBody) SetRequestId(v string) *DescribeLiveDomainStreamTranscodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBody) SetTranscodeDataList(v *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataList) *DescribeLiveDomainStreamTranscodeDataResponseBody {
	s.TranscodeDataList = v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBody) Validate() error {
	if s.TranscodeDataList != nil {
		if err := s.TranscodeDataList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataList struct {
	TranscodeData []*DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData `json:"TranscodeData,omitempty" xml:"TranscodeData,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataList) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataList) GetTranscodeData() []*DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData {
	return s.TranscodeData
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataList) SetTranscodeData(v []*DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataList {
	s.TranscodeData = v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataList) Validate() error {
	if s.TranscodeData != nil {
		for _, item := range s.TranscodeData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData struct {
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Fps          *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Resolution   *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	TanscodeType *string `json:"TanscodeType,omitempty" xml:"TanscodeType,omitempty"`
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) GetFps() *string {
	return s.Fps
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) GetResolution() *string {
	return s.Resolution
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) GetTanscodeType() *string {
	return s.TanscodeType
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) SetDomain(v string) *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData {
	s.Domain = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) SetDuration(v int32) *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData {
	s.Duration = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) SetFps(v string) *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData {
	s.Fps = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) SetRegion(v string) *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData {
	s.Region = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) SetResolution(v string) *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData {
	s.Resolution = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) SetTanscodeType(v string) *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData {
	s.TanscodeType = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) SetTimeStamp(v string) *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponseBodyTranscodeDataListTranscodeData) Validate() error {
	return dara.Validate(s)
}
