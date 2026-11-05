// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAssetSelectionCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddAssetSelectionCriteriaRequest
	GetClientToken() *string
	SetCriteria(v string) *AddAssetSelectionCriteriaRequest
	GetCriteria() *string
	SetCriteriaOperation(v string) *AddAssetSelectionCriteriaRequest
	GetCriteriaOperation() *string
	SetDryRun(v bool) *AddAssetSelectionCriteriaRequest
	GetDryRun() *bool
	SetSelectionKey(v string) *AddAssetSelectionCriteriaRequest
	GetSelectionKey() *string
	SetTargetOperationList(v []*AddAssetSelectionCriteriaRequestTargetOperationList) *AddAssetSelectionCriteriaRequest
	GetTargetOperationList() []*AddAssetSelectionCriteriaRequestTargetOperationList
}

type AddAssetSelectionCriteriaRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use different tokens for different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies the conditions for searching for assets. The value of this parameter is in JSON format. Note that the parameter value is case-sensitive.
	//
	// > You can search for assets by using conditions such as the instance ID, instance name, VPC ID, region, and public IP address. You can call the [DescribeCriteria](~~DescribeCriteria~~) operation to query the supported search conditions.
	//
	// example:
	//
	// {"LogicalExp":"AND","Criteria":"[{\\"name\\":\\"osType\\",\\"value\\":\\"linux\\",\\"logicalExp\\":\\"AND\\"},{\\"name\\":\\"alarmStatus\\",\\"value\\":\\"YES\\",\\"logicalExp\\":\\"AND\\"}]"}
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The operation type for the criteria. Valid values:
	//
	// - **add**: Adds assets.
	//
	// - **del**: Deletes assets.
	//
	// example:
	//
	// add
	CriteriaOperation *string `json:"CriteriaOperation,omitempty" xml:"CriteriaOperation,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - true: performs a dry run. The system checks the request but does not perform the actual operation.
	//
	// - false: performs the actual operation.
	//
	// Default value: false.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The unique identifier for this asset selection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5196d280-5bfa-496a-ba70-8a3935e3****
	SelectionKey *string `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
	// The asset list.
	TargetOperationList []*AddAssetSelectionCriteriaRequestTargetOperationList `json:"TargetOperationList,omitempty" xml:"TargetOperationList,omitempty" type:"Repeated"`
}

func (s AddAssetSelectionCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAssetSelectionCriteriaRequest) GoString() string {
	return s.String()
}

func (s *AddAssetSelectionCriteriaRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddAssetSelectionCriteriaRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *AddAssetSelectionCriteriaRequest) GetCriteriaOperation() *string {
	return s.CriteriaOperation
}

func (s *AddAssetSelectionCriteriaRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddAssetSelectionCriteriaRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *AddAssetSelectionCriteriaRequest) GetTargetOperationList() []*AddAssetSelectionCriteriaRequestTargetOperationList {
	return s.TargetOperationList
}

func (s *AddAssetSelectionCriteriaRequest) SetClientToken(v string) *AddAssetSelectionCriteriaRequest {
	s.ClientToken = &v
	return s
}

func (s *AddAssetSelectionCriteriaRequest) SetCriteria(v string) *AddAssetSelectionCriteriaRequest {
	s.Criteria = &v
	return s
}

func (s *AddAssetSelectionCriteriaRequest) SetCriteriaOperation(v string) *AddAssetSelectionCriteriaRequest {
	s.CriteriaOperation = &v
	return s
}

func (s *AddAssetSelectionCriteriaRequest) SetDryRun(v bool) *AddAssetSelectionCriteriaRequest {
	s.DryRun = &v
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
	// The operation type. Valid values:
	//
	// - **add**: Add.
	//
	// - **del**: Delete.
	//
	// example:
	//
	// del
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The asset ID. If you select assets by machine, the value is the UUID of the machine. If you select assets by group, the value is the group ID. If you select assets by VPC, the value is the VPC ID.
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
