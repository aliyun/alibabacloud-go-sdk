// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecRulesResponseBodyData) *DescribeApisecRulesResponseBody
	GetData() []*DescribeApisecRulesResponseBodyData
	SetRequestId(v string) *DescribeApisecRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeApisecRulesResponseBody
	GetTotalCount() *int64
}

type DescribeApisecRulesResponseBody struct {
	// The policies.
	Data []*DescribeApisecRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 6
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisecRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecRulesResponseBody) GetData() []*DescribeApisecRulesResponseBodyData {
	return s.Data
}

func (s *DescribeApisecRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeApisecRulesResponseBody) SetData(v []*DescribeApisecRulesResponseBodyData) *DescribeApisecRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecRulesResponseBody) SetRequestId(v string) *DescribeApisecRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecRulesResponseBody) SetTotalCount(v int64) *DescribeApisecRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisecRulesResponseBody) Validate() error {
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

type DescribeApisecRulesResponseBodyData struct {
	// The ID of the policy.
	//
	// example:
	//
	// 34933
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The details of the policy. The value is a string that consists of multiple parameters in the JSON format.
	//
	// example:
	//
	// {
	//
	//     "ext": "Date",
	//
	//     "regex": "-",
	//
	//     "code": "2009",
	//
	//     "level": "S1",
	//
	//     "origin": "default",
	//
	//     "name": "2009"
	//
	// }
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **risk**: risk detection
	//
	// 	- **event**: security event
	//
	// 	- **sensitive_word**: sensitive data
	//
	// 	- **auth_flag**: authentication credential
	//
	// 	- **api_tag**: business purpose
	//
	// 	- **desensitization**: data masking
	//
	// 	- **whitelist**: whitelist
	//
	// 	- **recognition**: API recognition
	//
	// 	- **offline_api**: lifecycle management
	//
	// example:
	//
	// risk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the policy was updated. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1721095301
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeApisecRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecRulesResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DescribeApisecRulesResponseBodyData) GetRule() *string {
	return s.Rule
}

func (s *DescribeApisecRulesResponseBodyData) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeApisecRulesResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DescribeApisecRulesResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeApisecRulesResponseBodyData) SetId(v int64) *DescribeApisecRulesResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeApisecRulesResponseBodyData) SetRule(v string) *DescribeApisecRulesResponseBodyData {
	s.Rule = &v
	return s
}

func (s *DescribeApisecRulesResponseBodyData) SetStatus(v int64) *DescribeApisecRulesResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeApisecRulesResponseBodyData) SetType(v string) *DescribeApisecRulesResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeApisecRulesResponseBodyData) SetUpdateTime(v int64) *DescribeApisecRulesResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *DescribeApisecRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
