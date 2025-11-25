// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCCRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeWebCCRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeWebCCRulesResponseBody
	GetTotalCount() *int64
	SetWebCCRules(v []*DescribeWebCCRulesResponseBodyWebCCRules) *DescribeWebCCRulesResponseBody
	GetWebCCRules() []*DescribeWebCCRulesResponseBodyWebCCRules
}

type DescribeWebCCRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EAED912D-909E-45F0-AF74-AC0CCDCAE314
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of custom frequency control rules.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The custom frequency control rule.
	WebCCRules []*DescribeWebCCRulesResponseBodyWebCCRules `json:"WebCCRules,omitempty" xml:"WebCCRules,omitempty" type:"Repeated"`
}

func (s DescribeWebCCRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebCCRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeWebCCRulesResponseBody) GetWebCCRules() []*DescribeWebCCRulesResponseBodyWebCCRules {
	return s.WebCCRules
}

func (s *DescribeWebCCRulesResponseBody) SetRequestId(v string) *DescribeWebCCRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebCCRulesResponseBody) SetTotalCount(v int64) *DescribeWebCCRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebCCRulesResponseBody) SetWebCCRules(v []*DescribeWebCCRulesResponseBodyWebCCRules) *DescribeWebCCRulesResponseBody {
	s.WebCCRules = v
	return s
}

func (s *DescribeWebCCRulesResponseBody) Validate() error {
	if s.WebCCRules != nil {
		for _, item := range s.WebCCRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebCCRulesResponseBodyWebCCRules struct {
	// The action triggered if the rule is matched. Valid values:
	//
	// 	- **close**: The requests that match the rule are blocked.
	//
	// 	- **captcha**: Completely Automated Public Turing test to tell Computers and Humans Apart (CAPTCHA) verification for the requests that match the rule is implemented.
	//
	// example:
	//
	// close
	Act *string `json:"Act,omitempty" xml:"Act,omitempty"`
	// The number of requests that are allowed from a single IP address. Valid values: **2*	- to **2000**.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The check interval. Valid values: **5*	- to **10800**. Unit: seconds.
	//
	// example:
	//
	// 5
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The match mode. Valid values:
	//
	// 	- **prefix**: prefix match.
	//
	// 	- **match**: exact match.
	//
	// example:
	//
	// prefix
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// wq
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The validity period. Valid values: **1*	- to **1440**. Unit: minutes.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The check path.
	//
	// example:
	//
	// /hello
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeWebCCRulesResponseBodyWebCCRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesResponseBodyWebCCRules) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) GetAct() *string {
	return s.Act
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) GetCount() *int32 {
	return s.Count
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) GetMode() *string {
	return s.Mode
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) GetName() *string {
	return s.Name
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) GetUri() *string {
	return s.Uri
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetAct(v string) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Act = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetCount(v int32) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Count = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetInterval(v int32) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Interval = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetMode(v string) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Mode = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetName(v string) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Name = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetTtl(v int32) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Ttl = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetUri(v string) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Uri = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) Validate() error {
	return dara.Validate(s)
}
