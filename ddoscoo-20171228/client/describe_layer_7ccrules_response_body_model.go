// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLayer7CCRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayer7CCRules(v []*DescribeLayer7CCRulesResponseBodyLayer7CCRules) *DescribeLayer7CCRulesResponseBody
	GetLayer7CCRules() []*DescribeLayer7CCRulesResponseBodyLayer7CCRules
	SetRequestId(v string) *DescribeLayer7CCRulesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeLayer7CCRulesResponseBody
	GetTotal() *int64
}

type DescribeLayer7CCRulesResponseBody struct {
	Layer7CCRules []*DescribeLayer7CCRulesResponseBodyLayer7CCRules `json:"Layer7CCRules,omitempty" xml:"Layer7CCRules,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLayer7CCRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer7CCRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLayer7CCRulesResponseBody) GetLayer7CCRules() []*DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	return s.Layer7CCRules
}

func (s *DescribeLayer7CCRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLayer7CCRulesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeLayer7CCRulesResponseBody) SetLayer7CCRules(v []*DescribeLayer7CCRulesResponseBodyLayer7CCRules) *DescribeLayer7CCRulesResponseBody {
	s.Layer7CCRules = v
	return s
}

func (s *DescribeLayer7CCRulesResponseBody) SetRequestId(v string) *DescribeLayer7CCRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBody) SetTotal(v int64) *DescribeLayer7CCRulesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLayer7CCRulesResponseBodyLayer7CCRules struct {
	// example:
	//
	// close
	Act *string `json:"Act,omitempty" xml:"Act,omitempty"`
	// example:
	//
	// 100
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// match
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// testCcRule1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1000
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// example:
	//
	// /a/b/c
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeLayer7CCRulesResponseBodyLayer7CCRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer7CCRulesResponseBodyLayer7CCRules) GoString() string {
	return s.String()
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) GetAct() *string {
	return s.Act
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) GetCount() *int32 {
	return s.Count
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) GetMode() *string {
	return s.Mode
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) GetName() *string {
	return s.Name
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) GetUri() *string {
	return s.Uri
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetAct(v string) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Act = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetCount(v int32) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Count = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetInterval(v int32) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Interval = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetMode(v string) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Mode = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetName(v string) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Name = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetTtl(v int32) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Ttl = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetUri(v string) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Uri = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) Validate() error {
	return dara.Validate(s)
}
