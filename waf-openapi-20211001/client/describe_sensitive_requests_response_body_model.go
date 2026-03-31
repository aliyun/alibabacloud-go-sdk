// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveRequestsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeSensitiveRequestsResponseBodyData) *DescribeSensitiveRequestsResponseBody
	GetData() []*DescribeSensitiveRequestsResponseBodyData
	SetRequestId(v string) *DescribeSensitiveRequestsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeSensitiveRequestsResponseBody
	GetTotalCount() *int64
}

type DescribeSensitiveRequestsResponseBody struct {
	// The tracing results of the data.
	Data []*DescribeSensitiveRequestsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 26E46541-7AAB-5565-801D-F14DBDC5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSensitiveRequestsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveRequestsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveRequestsResponseBody) GetData() []*DescribeSensitiveRequestsResponseBodyData {
	return s.Data
}

func (s *DescribeSensitiveRequestsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSensitiveRequestsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSensitiveRequestsResponseBody) SetData(v []*DescribeSensitiveRequestsResponseBodyData) *DescribeSensitiveRequestsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSensitiveRequestsResponseBody) SetRequestId(v string) *DescribeSensitiveRequestsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSensitiveRequestsResponseBody) SetTotalCount(v int64) *DescribeSensitiveRequestsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSensitiveRequestsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSensitiveRequestsResponseBodyData struct {
	// The number of risks in the previous 30 days.
	//
	// example:
	//
	// 23
	AbnormalCount *int64 `json:"AbnormalCount,omitempty" xml:"AbnormalCount,omitempty"`
	// The API.
	//
	// example:
	//
	// /api/users/login
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The ID of the API.
	//
	// example:
	//
	// 09559c0d71ca2ffc996b81***836d8
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 103.118.55.**
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// The evaluation result. Valid values:
	//
	// 	- **leak**: Data leaks may occur.
	//
	// 	- **none**: No data leak can occur.
	//
	// example:
	//
	// leak
	DetectionResult *string `json:"DetectionResult,omitempty" xml:"DetectionResult,omitempty"`
	// The number of events in the previous 30 days.
	//
	// example:
	//
	// 679
	EventCount *int64 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The statistics of the sensitive data.
	InfoCount []*DescribeSensitiveRequestsResponseBodyDataInfoCount `json:"InfoCount,omitempty" xml:"InfoCount,omitempty" type:"Repeated"`
	// The domain name of the API.
	//
	// example:
	//
	// a.****.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The sensitive data.
	SensitiveList []*string `json:"SensitiveList,omitempty" xml:"SensitiveList,omitempty" type:"Repeated"`
}

func (s DescribeSensitiveRequestsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveRequestsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveRequestsResponseBodyData) GetAbnormalCount() *int64 {
	return s.AbnormalCount
}

func (s *DescribeSensitiveRequestsResponseBodyData) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeSensitiveRequestsResponseBodyData) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeSensitiveRequestsResponseBodyData) GetClientIP() *string {
	return s.ClientIP
}

func (s *DescribeSensitiveRequestsResponseBodyData) GetDetectionResult() *string {
	return s.DetectionResult
}

func (s *DescribeSensitiveRequestsResponseBodyData) GetEventCount() *int64 {
	return s.EventCount
}

func (s *DescribeSensitiveRequestsResponseBodyData) GetInfoCount() []*DescribeSensitiveRequestsResponseBodyDataInfoCount {
	return s.InfoCount
}

func (s *DescribeSensitiveRequestsResponseBodyData) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeSensitiveRequestsResponseBodyData) GetSensitiveList() []*string {
	return s.SensitiveList
}

func (s *DescribeSensitiveRequestsResponseBodyData) SetAbnormalCount(v int64) *DescribeSensitiveRequestsResponseBodyData {
	s.AbnormalCount = &v
	return s
}

func (s *DescribeSensitiveRequestsResponseBodyData) SetApiFormat(v string) *DescribeSensitiveRequestsResponseBodyData {
	s.ApiFormat = &v
	return s
}

func (s *DescribeSensitiveRequestsResponseBodyData) SetApiId(v string) *DescribeSensitiveRequestsResponseBodyData {
	s.ApiId = &v
	return s
}

func (s *DescribeSensitiveRequestsResponseBodyData) SetClientIP(v string) *DescribeSensitiveRequestsResponseBodyData {
	s.ClientIP = &v
	return s
}

func (s *DescribeSensitiveRequestsResponseBodyData) SetDetectionResult(v string) *DescribeSensitiveRequestsResponseBodyData {
	s.DetectionResult = &v
	return s
}

func (s *DescribeSensitiveRequestsResponseBodyData) SetEventCount(v int64) *DescribeSensitiveRequestsResponseBodyData {
	s.EventCount = &v
	return s
}

func (s *DescribeSensitiveRequestsResponseBodyData) SetInfoCount(v []*DescribeSensitiveRequestsResponseBodyDataInfoCount) *DescribeSensitiveRequestsResponseBodyData {
	s.InfoCount = v
	return s
}

func (s *DescribeSensitiveRequestsResponseBodyData) SetMatchedHost(v string) *DescribeSensitiveRequestsResponseBodyData {
	s.MatchedHost = &v
	return s
}

func (s *DescribeSensitiveRequestsResponseBodyData) SetSensitiveList(v []*string) *DescribeSensitiveRequestsResponseBodyData {
	s.SensitiveList = v
	return s
}

func (s *DescribeSensitiveRequestsResponseBodyData) Validate() error {
	if s.InfoCount != nil {
		for _, item := range s.InfoCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSensitiveRequestsResponseBodyDataInfoCount struct {
	// The type of the sensitive data.
	//
	// example:
	//
	// 1001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of sensitive data entries.
	//
	// example:
	//
	// 23
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeSensitiveRequestsResponseBodyDataInfoCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveRequestsResponseBodyDataInfoCount) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveRequestsResponseBodyDataInfoCount) GetCode() *string {
	return s.Code
}

func (s *DescribeSensitiveRequestsResponseBodyDataInfoCount) GetCount() *int64 {
	return s.Count
}

func (s *DescribeSensitiveRequestsResponseBodyDataInfoCount) SetCode(v string) *DescribeSensitiveRequestsResponseBodyDataInfoCount {
	s.Code = &v
	return s
}

func (s *DescribeSensitiveRequestsResponseBodyDataInfoCount) SetCount(v int64) *DescribeSensitiveRequestsResponseBodyDataInfoCount {
	s.Count = &v
	return s
}

func (s *DescribeSensitiveRequestsResponseBodyDataInfoCount) Validate() error {
	return dara.Validate(s)
}
