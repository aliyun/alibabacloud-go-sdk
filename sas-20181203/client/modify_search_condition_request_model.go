// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySearchConditionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterConditions(v string) *ModifySearchConditionRequest
	GetFilterConditions() *string
	SetName(v string) *ModifySearchConditionRequest
	GetName() *string
	SetSourceIp(v string) *ModifySearchConditionRequest
	GetSourceIp() *string
	SetType(v string) *ModifySearchConditionRequest
	GetType() *string
}

type ModifySearchConditionRequest struct {
	// The filter conditions. This parameter is in JSON format. Pay attention to the letter case when you specify this parameter. The following fields are included:
	//
	// - **filterParams**: The filter parameters. This parameter is in JSON format. The following fields are included:
	//
	//   - **label**: The display name for the search in the console.
	//
	//   - **value**: The filter parameter condition. This parameter is in JSON format. The following fields are included:
	//
	//      - **name**: The filter condition field. For more information about the valid values of this field, see the description below.
	//
	//      - **value**: The value that corresponds to the filter condition field.
	//
	// - **LogicalExp**: The logical relationship between multiple filter conditions. Valid values:
	//
	//   - **OR**: or
	//
	//   - **AND**: and
	//
	// > Valid values of **name**:
	//
	// > - If **Type*	- is set to **ecs**, call the [DescribeCriteria](~~DescribeCriteria~~) operation to query the supported search conditions.
	//
	// > - If **Type*	- is set to **cloud_product**, call the [GetCloudAssetCriteria](~~GetCloudAssetCriteria~~) operation to query the supported search conditions.
	//
	// example:
	//
	// {
	//
	//     "filterParams": [
	//
	//         {
	//
	//             "label": "UUID：xxx",
	//
	//             "value": "{\\"name\\":\\"uuidList\\",\\"value\\":\\"xxx\\"}"
	//
	//         }
	//
	//     ],
	//
	//     "LogicalExp": "OR"
	//
	// }
	FilterConditions *string `json:"FilterConditions,omitempty" xml:"FilterConditions,omitempty"`
	// The name of the common filter condition.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. The system automatically obtains the value.
	//
	// example:
	//
	// 27.223.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The asset type. Default value: **ecs**. Valid values:
	//
	// - **ecs**: host asset
	//
	// - **cloud_product**: cloud service.
	//
	// example:
	//
	// ecs
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifySearchConditionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySearchConditionRequest) GoString() string {
	return s.String()
}

func (s *ModifySearchConditionRequest) GetFilterConditions() *string {
	return s.FilterConditions
}

func (s *ModifySearchConditionRequest) GetName() *string {
	return s.Name
}

func (s *ModifySearchConditionRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifySearchConditionRequest) GetType() *string {
	return s.Type
}

func (s *ModifySearchConditionRequest) SetFilterConditions(v string) *ModifySearchConditionRequest {
	s.FilterConditions = &v
	return s
}

func (s *ModifySearchConditionRequest) SetName(v string) *ModifySearchConditionRequest {
	s.Name = &v
	return s
}

func (s *ModifySearchConditionRequest) SetSourceIp(v string) *ModifySearchConditionRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifySearchConditionRequest) SetType(v string) *ModifySearchConditionRequest {
	s.Type = &v
	return s
}

func (s *ModifySearchConditionRequest) Validate() error {
	return dara.Validate(s)
}
