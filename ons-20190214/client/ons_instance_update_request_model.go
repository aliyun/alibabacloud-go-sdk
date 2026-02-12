// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsInstanceUpdateRequest
	GetInstanceId() *string
	SetInstanceName(v string) *OnsInstanceUpdateRequest
	GetInstanceName() *string
	SetRemark(v string) *OnsInstanceUpdateRequest
	GetRemark() *string
}

type OnsInstanceUpdateRequest struct {
	// The ID of the instance whose name or description you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new name of the instance. The name must meet the following rules:
	//
	// 	- The name of the instance must be unique in the region where the instance is deployed.
	//
	// 	- The name must be 3 to 64 characters in length and can contain letters, digits, hyphens (-), underscores (_), and Chinese characters.
	//
	// 	- If you do not configure this parameter, the instance name remains unchanged.
	//
	// example:
	//
	// Updatedname
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The updated description of the instance. If you do not configure this parameter, the instance description remains unchanged.
	//
	// example:
	//
	// Updateddescription
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s OnsInstanceUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceUpdateRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceUpdateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsInstanceUpdateRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *OnsInstanceUpdateRequest) GetRemark() *string {
	return s.Remark
}

func (s *OnsInstanceUpdateRequest) SetInstanceId(v string) *OnsInstanceUpdateRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsInstanceUpdateRequest) SetInstanceName(v string) *OnsInstanceUpdateRequest {
	s.InstanceName = &v
	return s
}

func (s *OnsInstanceUpdateRequest) SetRemark(v string) *OnsInstanceUpdateRequest {
	s.Remark = &v
	return s
}

func (s *OnsInstanceUpdateRequest) Validate() error {
	return dara.Validate(s)
}
