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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The transcoding usage data returned at each interval.
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
	// The main streaming domain. This parameter is returned only when you add the domain key to the value of the Split parameter.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The transcoding length. Unit: minutes.
	//
	// example:
	//
	// 2000
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The frame rate of the transcoded stream. This parameter is returned only when you add the fps key to the value of the Split parameter.
	//
	// example:
	//
	// normal
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The region in which the domain name resides. Valid values:
	//
	// >  This parameter takes effect only when you set Split to region.
	//
	// 	- **cn-beijing**: China (Beijing)
	//
	// 	- **cn-shanghai**: China (Shanghai)
	//
	// 	- **cn-qingdao**: China (Qingdao)
	//
	// 	- **cn-shenzhen**: China (Shenzhen)
	//
	// 	- **ap-northeast-1**: Japan (Tokyo)
	//
	// 	- **ap-southeast-1**: Singapore
	//
	// 	- **ap-southeast-5**: Indonesia (Jakarta)
	//
	// 	- **eu-central-1**: Germany (Frankfurt)
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The resolution of the transcoded stream. This parameter is returned only when you add the resolution key to the value of the Split parameter. Valid values:
	//
	// 	- **2K**
	//
	// 	- **4K**
	//
	// 	- **LD**: low definition
	//
	// 	- **SD**: standard definition
	//
	// 	- **HD**: high definition
	//
	// 	- **def**: audio
	//
	// example:
	//
	// HD
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The transcoding type. Valid values:
	//
	// >  This parameter takes effect only if the request parameter Split is set to transcode_type.
	//
	// 	- **H264NBHD**: Narrowband HD™ transcoding based on H.264
	//
	// 	- **H265NBHD**: Narrowband HD™ transcoding based on H.265
	//
	// 	- **AUDIO**: audio transcoding
	//
	// example:
	//
	// H264STD
	TanscodeType *string `json:"TanscodeType,omitempty" xml:"TanscodeType,omitempty"`
	// The timestamp of the data entry.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
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
