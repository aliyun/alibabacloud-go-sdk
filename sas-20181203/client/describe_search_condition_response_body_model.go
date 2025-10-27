// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSearchConditionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConditionList(v []*DescribeSearchConditionResponseBodyConditionList) *DescribeSearchConditionResponseBody
	GetConditionList() []*DescribeSearchConditionResponseBodyConditionList
	SetRequestId(v string) *DescribeSearchConditionResponseBody
	GetRequestId() *string
}

type DescribeSearchConditionResponseBody struct {
	// An array that consists of the filter conditions.
	ConditionList []*DescribeSearchConditionResponseBodyConditionList `json:"ConditionList,omitempty" xml:"ConditionList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3AEC47AF-8CFA-485E-AC9A-3A8ABC06EA7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSearchConditionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchConditionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSearchConditionResponseBody) GetConditionList() []*DescribeSearchConditionResponseBodyConditionList {
	return s.ConditionList
}

func (s *DescribeSearchConditionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSearchConditionResponseBody) SetConditionList(v []*DescribeSearchConditionResponseBodyConditionList) *DescribeSearchConditionResponseBody {
	s.ConditionList = v
	return s
}

func (s *DescribeSearchConditionResponseBody) SetRequestId(v string) *DescribeSearchConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSearchConditionResponseBody) Validate() error {
	if s.ConditionList != nil {
		for _, item := range s.ConditionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSearchConditionResponseBodyConditionList struct {
	// The type of the filter condition. Valid values:
	//
	// 	- **system**: default filter conditions.
	//
	// 	- **user**: custom filter conditions.
	//
	// example:
	//
	// system
	ConditionType *string `json:"ConditionType,omitempty" xml:"ConditionType,omitempty"`
	// The filter condition. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **filterParams**: the parameters of the filter condition. The value of this field is in the JSON format and contains the following fields:
	//
	//     	- **labelKey**: the key for rendering.
	//
	//     	- **label**: the display name.
	//
	//     	- **value**: the value of the filter condition. The value of this field is in the JSON format and contains the following fields:
	//
	//         	- **name**: the name of the filter item.
	//
	//         	- **value**: the value of the filter item.
	//
	// 	- **LogicalExp**: the logical relationship among the filter conditions. Valid values:
	//
	//     	- **AND**: The filter conditions are evaluated by using a logical **AND**.
	//
	//     	- **OR**: The filter conditions are evaluated by using a logical **OR**.
	//
	// >  If the value of **ConditionType*	- is **system**, **labelKey*	- is returned. The labelKey field is used only for internationalization rendering.
	//
	// example:
	//
	// {\\"filterParams\\":[{\\"labelKey\\":\\"a|b\\",\\"value\\":\\"{\\\\\\"name\\\\\\":\\\\\\"sadsasd\\\\\\",\\\\\\"value\\\\\\":\\\\\\"dasdsdas\\\\\\"}\\"}],\\"LogicalExp\\":\\"OR\\"}
	FilterConditions *string `json:"FilterConditions,omitempty" xml:"FilterConditions,omitempty"`
	// The filter condition name.
	//
	// example:
	//
	// StopMachine
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The key of the filter condition name.
	//
	// example:
	//
	// stop_machine
	NameKey *string `json:"NameKey,omitempty" xml:"NameKey,omitempty"`
}

func (s DescribeSearchConditionResponseBodyConditionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchConditionResponseBodyConditionList) GoString() string {
	return s.String()
}

func (s *DescribeSearchConditionResponseBodyConditionList) GetConditionType() *string {
	return s.ConditionType
}

func (s *DescribeSearchConditionResponseBodyConditionList) GetFilterConditions() *string {
	return s.FilterConditions
}

func (s *DescribeSearchConditionResponseBodyConditionList) GetName() *string {
	return s.Name
}

func (s *DescribeSearchConditionResponseBodyConditionList) GetNameKey() *string {
	return s.NameKey
}

func (s *DescribeSearchConditionResponseBodyConditionList) SetConditionType(v string) *DescribeSearchConditionResponseBodyConditionList {
	s.ConditionType = &v
	return s
}

func (s *DescribeSearchConditionResponseBodyConditionList) SetFilterConditions(v string) *DescribeSearchConditionResponseBodyConditionList {
	s.FilterConditions = &v
	return s
}

func (s *DescribeSearchConditionResponseBodyConditionList) SetName(v string) *DescribeSearchConditionResponseBodyConditionList {
	s.Name = &v
	return s
}

func (s *DescribeSearchConditionResponseBodyConditionList) SetNameKey(v string) *DescribeSearchConditionResponseBodyConditionList {
	s.NameKey = &v
	return s
}

func (s *DescribeSearchConditionResponseBodyConditionList) Validate() error {
	return dara.Validate(s)
}
