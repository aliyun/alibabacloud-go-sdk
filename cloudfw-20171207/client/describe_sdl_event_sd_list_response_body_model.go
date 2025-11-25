// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlEventSdListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSdlEventSdListResponseBody
	GetRequestId() *string
	SetSdlEventSensitiveDataList(v []*DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) *DescribeSdlEventSdListResponseBody
	GetSdlEventSensitiveDataList() []*DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList
	SetTotalCount(v int64) *DescribeSdlEventSdListResponseBody
	GetTotalCount() *int64
}

type DescribeSdlEventSdListResponseBody struct {
	// example:
	//
	// 15FCCC52-1E23-57AE-B5EF-3E00A3******
	RequestId                 *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdlEventSensitiveDataList []*DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList `json:"SdlEventSensitiveDataList,omitempty" xml:"SdlEventSensitiveDataList,omitempty" type:"Repeated"`
	// example:
	//
	// 6
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSdlEventSdListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventSdListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventSdListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSdlEventSdListResponseBody) GetSdlEventSensitiveDataList() []*DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList {
	return s.SdlEventSensitiveDataList
}

func (s *DescribeSdlEventSdListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSdlEventSdListResponseBody) SetRequestId(v string) *DescribeSdlEventSdListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSdlEventSdListResponseBody) SetSdlEventSensitiveDataList(v []*DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) *DescribeSdlEventSdListResponseBody {
	s.SdlEventSensitiveDataList = v
	return s
}

func (s *DescribeSdlEventSdListResponseBody) SetTotalCount(v int64) *DescribeSdlEventSdListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSdlEventSdListResponseBody) Validate() error {
	if s.SdlEventSensitiveDataList != nil {
		for _, item := range s.SdlEventSensitiveDataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList struct {
	// example:
	//
	// sk-N***************************SxiJ
	SensitiveData *string `json:"SensitiveData,omitempty" xml:"SensitiveData,omitempty"`
	// example:
	//
	// 6
	SensitiveDataCnt *int64 `json:"SensitiveDataCnt,omitempty" xml:"SensitiveDataCnt,omitempty"`
	// example:
	//
	// S3
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	SensitiveType  *string `json:"SensitiveType,omitempty" xml:"SensitiveType,omitempty"`
	// example:
	//
	// 172.23.191.XXX
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// example:
	//
	// 1753928907
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) GetSensitiveData() *string {
	return s.SensitiveData
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) GetSensitiveDataCnt() *int64 {
	return s.SensitiveDataCnt
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) GetSensitiveType() *string {
	return s.SensitiveType
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) GetSrcIp() *string {
	return s.SrcIp
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) SetSensitiveData(v string) *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList {
	s.SensitiveData = &v
	return s
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) SetSensitiveDataCnt(v int64) *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList {
	s.SensitiveDataCnt = &v
	return s
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) SetSensitiveLevel(v string) *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList {
	s.SensitiveLevel = &v
	return s
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) SetSensitiveType(v string) *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList {
	s.SensitiveType = &v
	return s
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) SetSrcIp(v string) *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList {
	s.SrcIp = &v
	return s
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) SetStartTime(v int64) *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList {
	s.StartTime = &v
	return s
}

func (s *DescribeSdlEventSdListResponseBodySdlEventSensitiveDataList) Validate() error {
	return dara.Validate(s)
}
