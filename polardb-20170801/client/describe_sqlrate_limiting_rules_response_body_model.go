// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLRateLimitingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeSQLRateLimitingRulesResponseBodyData) *DescribeSQLRateLimitingRulesResponseBody
	GetData() *DescribeSQLRateLimitingRulesResponseBodyData
	SetMaxResults(v int32) *DescribeSQLRateLimitingRulesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *DescribeSQLRateLimitingRulesResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeSQLRateLimitingRulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeSQLRateLimitingRulesResponseBody
	GetRequestId() *string
}

type DescribeSQLRateLimitingRulesResponseBody struct {
	Data *DescribeSQLRateLimitingRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// nextToken
	//
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh0vHYf39hc0J5qELgsazkBk
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 22C0ACF0-DD29-4B67-9190-B7A48C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSQLRateLimitingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLRateLimitingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLRateLimitingRulesResponseBody) GetData() *DescribeSQLRateLimitingRulesResponseBodyData {
	return s.Data
}

func (s *DescribeSQLRateLimitingRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSQLRateLimitingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSQLRateLimitingRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSQLRateLimitingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLRateLimitingRulesResponseBody) SetData(v *DescribeSQLRateLimitingRulesResponseBodyData) *DescribeSQLRateLimitingRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSQLRateLimitingRulesResponseBody) SetMaxResults(v int32) *DescribeSQLRateLimitingRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesResponseBody) SetMessage(v string) *DescribeSQLRateLimitingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesResponseBody) SetNextToken(v string) *DescribeSQLRateLimitingRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesResponseBody) SetRequestId(v string) *DescribeSQLRateLimitingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSQLRateLimitingRulesResponseBodyData struct {
	RuleList []*string `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeSQLRateLimitingRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLRateLimitingRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSQLRateLimitingRulesResponseBodyData) GetRuleList() []*string {
	return s.RuleList
}

func (s *DescribeSQLRateLimitingRulesResponseBodyData) SetRuleList(v []*string) *DescribeSQLRateLimitingRulesResponseBodyData {
	s.RuleList = v
	return s
}

func (s *DescribeSQLRateLimitingRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
