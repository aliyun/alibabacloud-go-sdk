// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAssetSelectionCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *AddAssetSelectionCriteriaRequest
	GetCriteria() *string
	SetCriteriaOperation(v string) *AddAssetSelectionCriteriaRequest
	GetCriteriaOperation() *string
	SetSelectionKey(v string) *AddAssetSelectionCriteriaRequest
	GetSelectionKey() *string
	SetTargetOperationList(v []*AddAssetSelectionCriteriaRequestTargetOperationList) *AddAssetSelectionCriteriaRequest
	GetTargetOperationList() []*AddAssetSelectionCriteriaRequestTargetOperationList
}

type AddAssetSelectionCriteriaRequest struct {
	// The search conditions that are used to query assets. The value of this parameter is in the JSON format and is case-sensitive.
	//
	// > A search condition can be an instance ID, instance name, virtual private cloud (VPC) ID, region, or public IP address. You can call the [DescribeCriteria](~~DescribeCriteria~~) operation to query the supported search conditions.
	//
	// example:
	//
	// {"LogicalExp":"AND","Criteria":"[{\\"name\\":\\"osType\\",\\"value\\":\\"linux\\",\\"logicalExp\\":\\"AND\\"},{\\"name\\":\\"alarmStatus\\",\\"value\\":\\"YES\\",\\"logicalExp\\":\\"AND\\"}]"}
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The type of the operation on search conditions. Valid values:
	//
	// 	- **add**: adds assets.
	//
	// 	- **del**: deletes assets.
	//
	// example:
	//
	// add
	CriteriaOperation *string `json:"CriteriaOperation,omitempty" xml:"CriteriaOperation,omitempty"`
	// The unique ID of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5196d280-5bfa-496a-ba70-8a3935e3****
	SelectionKey *string `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
	// The list of assets.
	TargetOperationList []*AddAssetSelectionCriteriaRequestTargetOperationList `json:"TargetOperationList,omitempty" xml:"TargetOperationList,omitempty" type:"Repeated"`
}

func (s AddAssetSelectionCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAssetSelectionCriteriaRequest) GoString() string {
	return s.String()
}

func (s *AddAssetSelectionCriteriaRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *AddAssetSelectionCriteriaRequest) GetCriteriaOperation() *string {
	return s.CriteriaOperation
}

func (s *AddAssetSelectionCriteriaRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *AddAssetSelectionCriteriaRequest) GetTargetOperationList() []*AddAssetSelectionCriteriaRequestTargetOperationList {
	return s.TargetOperationList
}

func (s *AddAssetSelectionCriteriaRequest) SetCriteria(v string) *AddAssetSelectionCriteriaRequest {
	s.Criteria = &v
	return s
}

func (s *AddAssetSelectionCriteriaRequest) SetCriteriaOperation(v string) *AddAssetSelectionCriteriaRequest {
	s.CriteriaOperation = &v
	return s
}

func (s *AddAssetSelectionCriteriaRequest) SetSelectionKey(v string) *AddAssetSelectionCriteriaRequest {
	s.SelectionKey = &v
	return s
}

func (s *AddAssetSelectionCriteriaRequest) SetTargetOperationList(v []*AddAssetSelectionCriteriaRequestTargetOperationList) *AddAssetSelectionCriteriaRequest {
	s.TargetOperationList = v
	return s
}

func (s *AddAssetSelectionCriteriaRequest) Validate() error {
	if s.TargetOperationList != nil {
		for _, item := range s.TargetOperationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddAssetSelectionCriteriaRequestTargetOperationList struct {
	// The type of the operation. Valid values:
	//
	// 	- **add**
	//
	// 	- **del**
	//
	// example:
	//
	// del
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The ID of the asset.
	//
	// example:
	//
	// 1188****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s AddAssetSelectionCriteriaRequestTargetOperationList) String() string {
	return dara.Prettify(s)
}

func (s AddAssetSelectionCriteriaRequestTargetOperationList) GoString() string {
	return s.String()
}

func (s *AddAssetSelectionCriteriaRequestTargetOperationList) GetOperation() *string {
	return s.Operation
}

func (s *AddAssetSelectionCriteriaRequestTargetOperationList) GetTarget() *string {
	return s.Target
}

func (s *AddAssetSelectionCriteriaRequestTargetOperationList) SetOperation(v string) *AddAssetSelectionCriteriaRequestTargetOperationList {
	s.Operation = &v
	return s
}

func (s *AddAssetSelectionCriteriaRequestTargetOperationList) SetTarget(v string) *AddAssetSelectionCriteriaRequestTargetOperationList {
	s.Target = &v
	return s
}

func (s *AddAssetSelectionCriteriaRequestTargetOperationList) Validate() error {
	return dara.Validate(s)
}
