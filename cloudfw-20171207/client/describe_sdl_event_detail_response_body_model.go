// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlEventDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSdlEventDetailResponseBody
	GetRequestId() *string
	SetSdlEventDetailList(v []*DescribeSdlEventDetailResponseBodySdlEventDetailList) *DescribeSdlEventDetailResponseBody
	GetSdlEventDetailList() []*DescribeSdlEventDetailResponseBodySdlEventDetailList
	SetTotalCount(v int64) *DescribeSdlEventDetailResponseBody
	GetTotalCount() *int64
}

type DescribeSdlEventDetailResponseBody struct {
	// example:
	//
	// D19D8F70-D64B-5A95-905A-6073BF4A****
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdlEventDetailList []*DescribeSdlEventDetailResponseBodySdlEventDetailList `json:"SdlEventDetailList,omitempty" xml:"SdlEventDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSdlEventDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSdlEventDetailResponseBody) GetSdlEventDetailList() []*DescribeSdlEventDetailResponseBodySdlEventDetailList {
	return s.SdlEventDetailList
}

func (s *DescribeSdlEventDetailResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSdlEventDetailResponseBody) SetRequestId(v string) *DescribeSdlEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSdlEventDetailResponseBody) SetSdlEventDetailList(v []*DescribeSdlEventDetailResponseBodySdlEventDetailList) *DescribeSdlEventDetailResponseBody {
	s.SdlEventDetailList = v
	return s
}

func (s *DescribeSdlEventDetailResponseBody) SetTotalCount(v int64) *DescribeSdlEventDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSdlEventDetailResponseBody) Validate() error {
	if s.SdlEventDetailList != nil {
		for _, item := range s.SdlEventDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSdlEventDetailResponseBodySdlEventDetailList struct {
	// example:
	//
	// event-test
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// 10
	SensitiveDataCnt *int64 `json:"SensitiveDataCnt,omitempty" xml:"SensitiveDataCnt,omitempty"`
	// example:
	//
	// S3
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	SensitiveType  *string `json:"SensitiveType,omitempty" xml:"SensitiveType,omitempty"`
	// example:
	//
	// 1753928907
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSdlEventDetailResponseBodySdlEventDetailList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventDetailResponseBodySdlEventDetailList) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventDetailResponseBodySdlEventDetailList) GetEventName() *string {
	return s.EventName
}

func (s *DescribeSdlEventDetailResponseBodySdlEventDetailList) GetSensitiveDataCnt() *int64 {
	return s.SensitiveDataCnt
}

func (s *DescribeSdlEventDetailResponseBodySdlEventDetailList) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *DescribeSdlEventDetailResponseBodySdlEventDetailList) GetSensitiveType() *string {
	return s.SensitiveType
}

func (s *DescribeSdlEventDetailResponseBodySdlEventDetailList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSdlEventDetailResponseBodySdlEventDetailList) SetEventName(v string) *DescribeSdlEventDetailResponseBodySdlEventDetailList {
	s.EventName = &v
	return s
}

func (s *DescribeSdlEventDetailResponseBodySdlEventDetailList) SetSensitiveDataCnt(v int64) *DescribeSdlEventDetailResponseBodySdlEventDetailList {
	s.SensitiveDataCnt = &v
	return s
}

func (s *DescribeSdlEventDetailResponseBodySdlEventDetailList) SetSensitiveLevel(v string) *DescribeSdlEventDetailResponseBodySdlEventDetailList {
	s.SensitiveLevel = &v
	return s
}

func (s *DescribeSdlEventDetailResponseBodySdlEventDetailList) SetSensitiveType(v string) *DescribeSdlEventDetailResponseBodySdlEventDetailList {
	s.SensitiveType = &v
	return s
}

func (s *DescribeSdlEventDetailResponseBodySdlEventDetailList) SetStartTime(v int64) *DescribeSdlEventDetailResponseBodySdlEventDetailList {
	s.StartTime = &v
	return s
}

func (s *DescribeSdlEventDetailResponseBodySdlEventDetailList) Validate() error {
	return dara.Validate(s)
}
