// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRiskNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceList(v []*ListInstanceRiskNumRequestInstanceList) *ListInstanceRiskNumRequest
	GetInstanceList() []*ListInstanceRiskNumRequestInstanceList
}

type ListInstanceRiskNumRequest struct {
	// The instances.
	InstanceList []*ListInstanceRiskNumRequestInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
}

func (s ListInstanceRiskNumRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRiskNumRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRiskNumRequest) GetInstanceList() []*ListInstanceRiskNumRequestInstanceList {
	return s.InstanceList
}

func (s *ListInstanceRiskNumRequest) SetInstanceList(v []*ListInstanceRiskNumRequestInstanceList) *ListInstanceRiskNumRequest {
	s.InstanceList = v
	return s
}

func (s *ListInstanceRiskNumRequest) Validate() error {
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

type ListInstanceRiskNumRequestInstanceList struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-wz9fdluqx20mp2x7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UUID of the instance.
	//
	// example:
	//
	// f2d6e901-1004-4ca8-9dae-53ec04a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListInstanceRiskNumRequestInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRiskNumRequestInstanceList) GoString() string {
	return s.String()
}

func (s *ListInstanceRiskNumRequestInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceRiskNumRequestInstanceList) GetUuid() *string {
	return s.Uuid
}

func (s *ListInstanceRiskNumRequestInstanceList) SetInstanceId(v string) *ListInstanceRiskNumRequestInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceRiskNumRequestInstanceList) SetUuid(v string) *ListInstanceRiskNumRequestInstanceList {
	s.Uuid = &v
	return s
}

func (s *ListInstanceRiskNumRequestInstanceList) Validate() error {
	return dara.Validate(s)
}
