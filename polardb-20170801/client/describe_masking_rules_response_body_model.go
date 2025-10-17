// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMaskingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeMaskingRulesResponseBody
	GetDBClusterId() *string
	SetData(v *DescribeMaskingRulesResponseBodyData) *DescribeMaskingRulesResponseBody
	GetData() *DescribeMaskingRulesResponseBodyData
	SetMessage(v string) *DescribeMaskingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMaskingRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMaskingRulesResponseBody
	GetSuccess() *bool
}

type DescribeMaskingRulesResponseBody struct {
	// The ID of the cluster.
	//
	// example:
	//
	// pc-bp1s826a1up******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The result data that is returned.
	Data *DescribeMaskingRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned for the request.
	//
	// > If the request is successful, Successful is returned. If the request fails, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2F83D131-1C18-4599-889D-729A9D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid value:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMaskingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaskingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMaskingRulesResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeMaskingRulesResponseBody) GetData() *DescribeMaskingRulesResponseBodyData {
	return s.Data
}

func (s *DescribeMaskingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMaskingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMaskingRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMaskingRulesResponseBody) SetDBClusterId(v string) *DescribeMaskingRulesResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeMaskingRulesResponseBody) SetData(v *DescribeMaskingRulesResponseBodyData) *DescribeMaskingRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMaskingRulesResponseBody) SetMessage(v string) *DescribeMaskingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMaskingRulesResponseBody) SetRequestId(v string) *DescribeMaskingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMaskingRulesResponseBody) SetSuccess(v bool) *DescribeMaskingRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMaskingRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMaskingRulesResponseBodyData struct {
	// Details about the masking rules.
	RuleList []*string `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// The version of the masking rule. Valid values: v1 and v2. Default value: v1
	//
	// example:
	//
	// v1
	RuleVersion *string `json:"RuleVersion,omitempty" xml:"RuleVersion,omitempty"`
}

func (s DescribeMaskingRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaskingRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMaskingRulesResponseBodyData) GetRuleList() []*string {
	return s.RuleList
}

func (s *DescribeMaskingRulesResponseBodyData) GetRuleVersion() *string {
	return s.RuleVersion
}

func (s *DescribeMaskingRulesResponseBodyData) SetRuleList(v []*string) *DescribeMaskingRulesResponseBodyData {
	s.RuleList = v
	return s
}

func (s *DescribeMaskingRulesResponseBodyData) SetRuleVersion(v string) *DescribeMaskingRulesResponseBodyData {
	s.RuleVersion = &v
	return s
}

func (s *DescribeMaskingRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
