// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRiskLevelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceList(v []*ListInstanceRiskLevelsRequestInstanceList) *ListInstanceRiskLevelsRequest
	GetInstanceList() []*ListInstanceRiskLevelsRequestInstanceList
}

type ListInstanceRiskLevelsRequest struct {
	// The instances.
	InstanceList []*ListInstanceRiskLevelsRequestInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
}

func (s ListInstanceRiskLevelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRiskLevelsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRiskLevelsRequest) GetInstanceList() []*ListInstanceRiskLevelsRequestInstanceList {
	return s.InstanceList
}

func (s *ListInstanceRiskLevelsRequest) SetInstanceList(v []*ListInstanceRiskLevelsRequestInstanceList) *ListInstanceRiskLevelsRequest {
	s.InstanceList = v
	return s
}

func (s *ListInstanceRiskLevelsRequest) Validate() error {
	if s.InstanceList != nil {
		for _, item := range s.InstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceRiskLevelsRequestInstanceList struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-m5efigezp50l2cmb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The serial number of the instance.
	//
	// example:
	//
	// f2d6e901-1004-4ca8-9dae-53ec04a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListInstanceRiskLevelsRequestInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRiskLevelsRequestInstanceList) GoString() string {
	return s.String()
}

func (s *ListInstanceRiskLevelsRequestInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceRiskLevelsRequestInstanceList) GetUuid() *string {
	return s.Uuid
}

func (s *ListInstanceRiskLevelsRequestInstanceList) SetInstanceId(v string) *ListInstanceRiskLevelsRequestInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceRiskLevelsRequestInstanceList) SetUuid(v string) *ListInstanceRiskLevelsRequestInstanceList {
	s.Uuid = &v
	return s
}

func (s *ListInstanceRiskLevelsRequestInstanceList) Validate() error {
	return dara.Validate(s)
}
