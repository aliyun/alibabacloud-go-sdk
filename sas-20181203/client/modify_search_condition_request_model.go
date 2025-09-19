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
	// The filter condition. The value of this parameter is in the JSON format and is case-sensitive. The value contains the following fields:
	//
	// 	- **filterParams**: the filter-related parameters. The value is in the JSON format. Valid values:
	//
	//     	- **label**: the display name of the filter condition in the console.
	//
	//     	- **value**: the settings of the filter condition. The value is in the JSON format. The value contains the following fields:
	//
	//         	- **name**: the name of the field for filtering. For more information, see the value description of name.
	//
	//         	- **value**: the value of the field for filtering.
	//
	// 	- **LogicalExp**: the logical relationship among multiple filter conditions. Valid values:
	//
	//     	- **OR**
	//
	//     	- **AND**
	//
	// >  Value description of **name**:
	//
	// 	- If **Type*	- is set to **ecs**, you can call the [DescribeCriteria](~~DescribeCriteria~~) operation to query the supported filter conditions.
	//
	// 	- If **Type*	- is set to **cloud_product**, you can call the [GetCloudAssetCriteria](~~GetCloudAssetCriteria~~) operation to query the supported filter conditions.
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
	// The source IP address of the request. You do not need to specify this parameter. It is automatically obtained by the system.
	//
	// example:
	//
	// 27.223.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the asset. Default value: **ecs**. Valid values:
	//
	// 	- **ecs**: host
	//
	// 	- **cloud_product**: Alibaba Cloud service
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
