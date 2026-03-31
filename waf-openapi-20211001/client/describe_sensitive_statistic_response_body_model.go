// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeSensitiveStatisticResponseBodyData) *DescribeSensitiveStatisticResponseBody
	GetData() []*DescribeSensitiveStatisticResponseBodyData
	SetRequestId(v string) *DescribeSensitiveStatisticResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeSensitiveStatisticResponseBody
	GetTotalCount() *int64
}

type DescribeSensitiveStatisticResponseBody struct {
	// The statistics of the sensitive data.
	Data []*DescribeSensitiveStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeSensitiveStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveStatisticResponseBody) GetData() []*DescribeSensitiveStatisticResponseBodyData {
	return s.Data
}

func (s *DescribeSensitiveStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSensitiveStatisticResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSensitiveStatisticResponseBody) SetData(v []*DescribeSensitiveStatisticResponseBodyData) *DescribeSensitiveStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSensitiveStatisticResponseBody) SetRequestId(v string) *DescribeSensitiveStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSensitiveStatisticResponseBody) SetTotalCount(v int64) *DescribeSensitiveStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSensitiveStatisticResponseBody) Validate() error {
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

type DescribeSensitiveStatisticResponseBodyData struct {
	// The API.
	//
	// example:
	//
	// /api/login
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 10.50.11.**
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 169
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The domain name.
	//
	// example:
	//
	// a.****.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The type of the sensitive data.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported types of sensitive data.
	//
	// example:
	//
	// 1003
	SensitiveCode *string `json:"SensitiveCode,omitempty" xml:"SensitiveCode,omitempty"`
}

func (s DescribeSensitiveStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveStatisticResponseBodyData) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeSensitiveStatisticResponseBodyData) GetClientIP() *string {
	return s.ClientIP
}

func (s *DescribeSensitiveStatisticResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *DescribeSensitiveStatisticResponseBodyData) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeSensitiveStatisticResponseBodyData) GetSensitiveCode() *string {
	return s.SensitiveCode
}

func (s *DescribeSensitiveStatisticResponseBodyData) SetApiFormat(v string) *DescribeSensitiveStatisticResponseBodyData {
	s.ApiFormat = &v
	return s
}

func (s *DescribeSensitiveStatisticResponseBodyData) SetClientIP(v string) *DescribeSensitiveStatisticResponseBodyData {
	s.ClientIP = &v
	return s
}

func (s *DescribeSensitiveStatisticResponseBodyData) SetCount(v int64) *DescribeSensitiveStatisticResponseBodyData {
	s.Count = &v
	return s
}

func (s *DescribeSensitiveStatisticResponseBodyData) SetMatchedHost(v string) *DescribeSensitiveStatisticResponseBodyData {
	s.MatchedHost = &v
	return s
}

func (s *DescribeSensitiveStatisticResponseBodyData) SetSensitiveCode(v string) *DescribeSensitiveStatisticResponseBodyData {
	s.SensitiveCode = &v
	return s
}

func (s *DescribeSensitiveStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}
