// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCustomFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFieldId(v string) *DisableCustomFieldRequest
	GetFieldId() *string
	SetInstanceId(v string) *DisableCustomFieldRequest
	GetInstanceId() *string
}

type DisableCustomFieldRequest struct {
	// fieldId
	//
	// This parameter is required.
	//
	// example:
	//
	// ufd_001
	FieldId *string `json:"FieldId,omitempty" xml:"FieldId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableCustomFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableCustomFieldRequest) GoString() string {
	return s.String()
}

func (s *DisableCustomFieldRequest) GetFieldId() *string {
	return s.FieldId
}

func (s *DisableCustomFieldRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableCustomFieldRequest) SetFieldId(v string) *DisableCustomFieldRequest {
	s.FieldId = &v
	return s
}

func (s *DisableCustomFieldRequest) SetInstanceId(v string) *DisableCustomFieldRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableCustomFieldRequest) Validate() error {
	return dara.Validate(s)
}
