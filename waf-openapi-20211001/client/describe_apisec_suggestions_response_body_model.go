// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecSuggestionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecSuggestionsResponseBodyData) *DescribeApisecSuggestionsResponseBody
	GetData() []*DescribeApisecSuggestionsResponseBodyData
	SetRequestId(v string) *DescribeApisecSuggestionsResponseBody
	GetRequestId() *string
}

type DescribeApisecSuggestionsResponseBody struct {
	// The protection suggestions.
	Data []*DescribeApisecSuggestionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApisecSuggestionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSuggestionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecSuggestionsResponseBody) GetData() []*DescribeApisecSuggestionsResponseBodyData {
	return s.Data
}

func (s *DescribeApisecSuggestionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecSuggestionsResponseBody) SetData(v []*DescribeApisecSuggestionsResponseBodyData) *DescribeApisecSuggestionsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecSuggestionsResponseBody) SetRequestId(v string) *DescribeApisecSuggestionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecSuggestionsResponseBody) Validate() error {
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

type DescribeApisecSuggestionsResponseBodyData struct {
	// The API.
	//
	// example:
	//
	// /apisec/v1/saveinfo
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The domain name or IP address of the API.
	//
	// example:
	//
	// a.aliyun.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The rule ID of the protection suggestion.
	//
	// example:
	//
	// 15060a1f8fed40130b7c4a7bf8d8733b
	SuggestId *string `json:"SuggestId,omitempty" xml:"SuggestId,omitempty"`
	// The rule content of the protection suggestion. The value is a string that consists of multiple parameters in the JSON format. Valid values:
	//
	// 	- **event_tags**: event type
	//
	// 	- **black_iplist**: IP address blacklist
	//
	// 	- **ip_baseline**: IP address
	//
	// 	- **freq_baseline**: throttling frequency
	//
	// 	- **client_id_baseline**: client information
	//
	// 	- **country_baseline**: country information
	//
	// 	- **province_baseline**: province information
	//
	// 	- **sensitive_type**: sensitive information
	//
	// example:
	//
	// {
	//
	//     "rule": "ClientRule",
	//
	//     "client_id_baseline": ["Edge"]
	//
	// }
	SuggestRule *string `json:"SuggestRule,omitempty" xml:"SuggestRule,omitempty"`
	// The rule type of the protection suggestion. Valid values:
	//
	// 	- **BotRule**: bot management rules
	//
	// 	- **BlackIPRule**: IP address blacklist rules
	//
	// 	- **WhiteIPRule**: IP address whitelist rules
	//
	// 	- **RateLimitRule**: throttling rules
	//
	// 	- **ClientRule**: client rules
	//
	// 	- **GeoRule**: region-related rules
	//
	// 	- **SensitiveRule**: sensitive information rules
	//
	// 	- **UnauthRule**: authentication rules
	//
	// example:
	//
	// WhiteIPRule
	SuggestType *string `json:"SuggestType,omitempty" xml:"SuggestType,omitempty"`
}

func (s DescribeApisecSuggestionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSuggestionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecSuggestionsResponseBodyData) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeApisecSuggestionsResponseBodyData) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeApisecSuggestionsResponseBodyData) GetSuggestId() *string {
	return s.SuggestId
}

func (s *DescribeApisecSuggestionsResponseBodyData) GetSuggestRule() *string {
	return s.SuggestRule
}

func (s *DescribeApisecSuggestionsResponseBodyData) GetSuggestType() *string {
	return s.SuggestType
}

func (s *DescribeApisecSuggestionsResponseBodyData) SetApiFormat(v string) *DescribeApisecSuggestionsResponseBodyData {
	s.ApiFormat = &v
	return s
}

func (s *DescribeApisecSuggestionsResponseBodyData) SetMatchedHost(v string) *DescribeApisecSuggestionsResponseBodyData {
	s.MatchedHost = &v
	return s
}

func (s *DescribeApisecSuggestionsResponseBodyData) SetSuggestId(v string) *DescribeApisecSuggestionsResponseBodyData {
	s.SuggestId = &v
	return s
}

func (s *DescribeApisecSuggestionsResponseBodyData) SetSuggestRule(v string) *DescribeApisecSuggestionsResponseBodyData {
	s.SuggestRule = &v
	return s
}

func (s *DescribeApisecSuggestionsResponseBodyData) SetSuggestType(v string) *DescribeApisecSuggestionsResponseBodyData {
	s.SuggestType = &v
	return s
}

func (s *DescribeApisecSuggestionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
