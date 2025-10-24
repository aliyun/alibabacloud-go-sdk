// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveRequestLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeSensitiveRequestLogResponseBodyData) *DescribeSensitiveRequestLogResponseBody
	GetData() []*DescribeSensitiveRequestLogResponseBodyData
	SetRequestId(v string) *DescribeSensitiveRequestLogResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeSensitiveRequestLogResponseBody
	GetTotalCount() *int64
}

type DescribeSensitiveRequestLogResponseBody struct {
	// The access logs.
	Data []*DescribeSensitiveRequestLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 7
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSensitiveRequestLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveRequestLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveRequestLogResponseBody) GetData() []*DescribeSensitiveRequestLogResponseBodyData {
	return s.Data
}

func (s *DescribeSensitiveRequestLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSensitiveRequestLogResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSensitiveRequestLogResponseBody) SetData(v []*DescribeSensitiveRequestLogResponseBodyData) *DescribeSensitiveRequestLogResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSensitiveRequestLogResponseBody) SetRequestId(v string) *DescribeSensitiveRequestLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSensitiveRequestLogResponseBody) SetTotalCount(v int64) *DescribeSensitiveRequestLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSensitiveRequestLogResponseBody) Validate() error {
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

type DescribeSensitiveRequestLogResponseBodyData struct {
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
	// 197b52abcd81d6a8bd4***e477
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 103.118.55.**
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// The number of sensitive data records involved in cross-border data transfer.
	//
	// example:
	//
	// 12
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The domain name of the API.
	//
	// example:
	//
	// a.****.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// IP region, formatted as a region code.
	//
	// example:
	//
	// CN
	RemoteCountryId *string `json:"RemoteCountryId,omitempty" xml:"RemoteCountryId,omitempty"`
	// The time when the request was initiated. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1723392000
	RequestTime *int64 `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	// The details of sensitive data. The value is a string that consists of a JSON struct. The JSON struct contains key-value pairs. In a key-value pair, a key indicates the identifier of a sensitive data type, including built-in and custom types, and a value indicates specific sensitive data.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported sensitive data types.
	//
	// example:
	//
	// {
	//
	//           "1000":[
	//
	//               "90.88.49.19",
	//
	//               "90.88.49.19"
	//
	//           ],
	//
	//           "835436":[
	//
	//               "www.abc.com"
	//
	//           ]
	//
	// }
	SensitiveList *string `json:"SensitiveList,omitempty" xml:"SensitiveList,omitempty"`
	// The trace ID.
	//
	// example:
	//
	// 0a3d455b17027784870843933dce3d
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeSensitiveRequestLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveRequestLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveRequestLogResponseBodyData) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeSensitiveRequestLogResponseBodyData) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeSensitiveRequestLogResponseBodyData) GetClientIP() *string {
	return s.ClientIP
}

func (s *DescribeSensitiveRequestLogResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *DescribeSensitiveRequestLogResponseBodyData) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeSensitiveRequestLogResponseBodyData) GetRemoteCountryId() *string {
	return s.RemoteCountryId
}

func (s *DescribeSensitiveRequestLogResponseBodyData) GetRequestTime() *int64 {
	return s.RequestTime
}

func (s *DescribeSensitiveRequestLogResponseBodyData) GetSensitiveList() *string {
	return s.SensitiveList
}

func (s *DescribeSensitiveRequestLogResponseBodyData) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeSensitiveRequestLogResponseBodyData) SetApiFormat(v string) *DescribeSensitiveRequestLogResponseBodyData {
	s.ApiFormat = &v
	return s
}

func (s *DescribeSensitiveRequestLogResponseBodyData) SetApiId(v string) *DescribeSensitiveRequestLogResponseBodyData {
	s.ApiId = &v
	return s
}

func (s *DescribeSensitiveRequestLogResponseBodyData) SetClientIP(v string) *DescribeSensitiveRequestLogResponseBodyData {
	s.ClientIP = &v
	return s
}

func (s *DescribeSensitiveRequestLogResponseBodyData) SetCount(v int64) *DescribeSensitiveRequestLogResponseBodyData {
	s.Count = &v
	return s
}

func (s *DescribeSensitiveRequestLogResponseBodyData) SetMatchedHost(v string) *DescribeSensitiveRequestLogResponseBodyData {
	s.MatchedHost = &v
	return s
}

func (s *DescribeSensitiveRequestLogResponseBodyData) SetRemoteCountryId(v string) *DescribeSensitiveRequestLogResponseBodyData {
	s.RemoteCountryId = &v
	return s
}

func (s *DescribeSensitiveRequestLogResponseBodyData) SetRequestTime(v int64) *DescribeSensitiveRequestLogResponseBodyData {
	s.RequestTime = &v
	return s
}

func (s *DescribeSensitiveRequestLogResponseBodyData) SetSensitiveList(v string) *DescribeSensitiveRequestLogResponseBodyData {
	s.SensitiveList = &v
	return s
}

func (s *DescribeSensitiveRequestLogResponseBodyData) SetTraceId(v string) *DescribeSensitiveRequestLogResponseBodyData {
	s.TraceId = &v
	return s
}

func (s *DescribeSensitiveRequestLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
