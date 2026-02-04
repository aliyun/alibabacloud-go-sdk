// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFieldId(v string) *DeleteCustomFieldRequest
	GetFieldId() *string
	SetInstanceId(v string) *DeleteCustomFieldRequest
	GetInstanceId() *string
}

type DeleteCustomFieldRequest struct {
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

func (s DeleteCustomFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomFieldRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomFieldRequest) GetFieldId() *string {
	return s.FieldId
}

func (s *DeleteCustomFieldRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCustomFieldRequest) SetFieldId(v string) *DeleteCustomFieldRequest {
	s.FieldId = &v
	return s
}

func (s *DeleteCustomFieldRequest) SetInstanceId(v string) *DeleteCustomFieldRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCustomFieldRequest) Validate() error {
	return dara.Validate(s)
}
