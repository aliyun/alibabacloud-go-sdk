// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveApiStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeSensitiveApiStatisticResponseBodyData) *DescribeSensitiveApiStatisticResponseBody
	GetData() []*DescribeSensitiveApiStatisticResponseBodyData
	SetRequestId(v string) *DescribeSensitiveApiStatisticResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeSensitiveApiStatisticResponseBody
	GetTotalCount() *int64
}

type DescribeSensitiveApiStatisticResponseBody struct {
	// The statistics.
	Data []*DescribeSensitiveApiStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSensitiveApiStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveApiStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveApiStatisticResponseBody) GetData() []*DescribeSensitiveApiStatisticResponseBodyData {
	return s.Data
}

func (s *DescribeSensitiveApiStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSensitiveApiStatisticResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSensitiveApiStatisticResponseBody) SetData(v []*DescribeSensitiveApiStatisticResponseBodyData) *DescribeSensitiveApiStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSensitiveApiStatisticResponseBody) SetRequestId(v string) *DescribeSensitiveApiStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSensitiveApiStatisticResponseBody) SetTotalCount(v int64) *DescribeSensitiveApiStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSensitiveApiStatisticResponseBody) Validate() error {
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

type DescribeSensitiveApiStatisticResponseBodyData struct {
	// The number of personal information records involved in cross-border data transfer by domain name.
	//
	// example:
	//
	// 213
	InfoOutboundCount *int64 `json:"InfoOutboundCount,omitempty" xml:"InfoOutboundCount,omitempty"`
	// The domain name-related APIs.
	List []*DescribeSensitiveApiStatisticResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The domain name or IP address.
	//
	// example:
	//
	// www.***.top
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The number of sensitive personal information records involved in cross-border data transfer by domain name.
	//
	// example:
	//
	// 127
	SensitiveOutboundCount *int64 `json:"SensitiveOutboundCount,omitempty" xml:"SensitiveOutboundCount,omitempty"`
}

func (s DescribeSensitiveApiStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveApiStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveApiStatisticResponseBodyData) GetInfoOutboundCount() *int64 {
	return s.InfoOutboundCount
}

func (s *DescribeSensitiveApiStatisticResponseBodyData) GetList() []*DescribeSensitiveApiStatisticResponseBodyDataList {
	return s.List
}

func (s *DescribeSensitiveApiStatisticResponseBodyData) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeSensitiveApiStatisticResponseBodyData) GetSensitiveOutboundCount() *int64 {
	return s.SensitiveOutboundCount
}

func (s *DescribeSensitiveApiStatisticResponseBodyData) SetInfoOutboundCount(v int64) *DescribeSensitiveApiStatisticResponseBodyData {
	s.InfoOutboundCount = &v
	return s
}

func (s *DescribeSensitiveApiStatisticResponseBodyData) SetList(v []*DescribeSensitiveApiStatisticResponseBodyDataList) *DescribeSensitiveApiStatisticResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeSensitiveApiStatisticResponseBodyData) SetMatchedHost(v string) *DescribeSensitiveApiStatisticResponseBodyData {
	s.MatchedHost = &v
	return s
}

func (s *DescribeSensitiveApiStatisticResponseBodyData) SetSensitiveOutboundCount(v int64) *DescribeSensitiveApiStatisticResponseBodyData {
	s.SensitiveOutboundCount = &v
	return s
}

func (s *DescribeSensitiveApiStatisticResponseBodyData) Validate() error {
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

type DescribeSensitiveApiStatisticResponseBodyDataList struct {
	// The API.
	//
	// example:
	//
	// /api/login
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The ID of the API.
	//
	// example:
	//
	// d288137009c119a873d4c395****
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The number of personal information records involved in cross-border data transfer by API.
	//
	// example:
	//
	// 78
	InfoCount *int64 `json:"InfoCount,omitempty" xml:"InfoCount,omitempty"`
	// The types of sensitive data.
	SensitiveCode []*string `json:"SensitiveCode,omitempty" xml:"SensitiveCode,omitempty" type:"Repeated"`
	// The number of sensitive personal information records involved in cross-border data transfer by API.
	//
	// example:
	//
	// 55
	SensitiveCount *int64 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
}

func (s DescribeSensitiveApiStatisticResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveApiStatisticResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveApiStatisticResponseBodyDataList) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeSensitiveApiStatisticResponseBodyDataList) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeSensitiveApiStatisticResponseBodyDataList) GetInfoCount() *int64 {
	return s.InfoCount
}

func (s *DescribeSensitiveApiStatisticResponseBodyDataList) GetSensitiveCode() []*string {
	return s.SensitiveCode
}

func (s *DescribeSensitiveApiStatisticResponseBodyDataList) GetSensitiveCount() *int64 {
	return s.SensitiveCount
}

func (s *DescribeSensitiveApiStatisticResponseBodyDataList) SetApiFormat(v string) *DescribeSensitiveApiStatisticResponseBodyDataList {
	s.ApiFormat = &v
	return s
}

func (s *DescribeSensitiveApiStatisticResponseBodyDataList) SetApiId(v string) *DescribeSensitiveApiStatisticResponseBodyDataList {
	s.ApiId = &v
	return s
}

func (s *DescribeSensitiveApiStatisticResponseBodyDataList) SetInfoCount(v int64) *DescribeSensitiveApiStatisticResponseBodyDataList {
	s.InfoCount = &v
	return s
}

func (s *DescribeSensitiveApiStatisticResponseBodyDataList) SetSensitiveCode(v []*string) *DescribeSensitiveApiStatisticResponseBodyDataList {
	s.SensitiveCode = v
	return s
}

func (s *DescribeSensitiveApiStatisticResponseBodyDataList) SetSensitiveCount(v int64) *DescribeSensitiveApiStatisticResponseBodyDataList {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeSensitiveApiStatisticResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
