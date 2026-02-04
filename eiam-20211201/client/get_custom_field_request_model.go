// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFieldId(v string) *GetCustomFieldRequest
	GetFieldId() *string
	SetInstanceId(v string) *GetCustomFieldRequest
	GetInstanceId() *string
}

type GetCustomFieldRequest struct {
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

func (s GetCustomFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomFieldRequest) GoString() string {
	return s.String()
}

func (s *GetCustomFieldRequest) GetFieldId() *string {
	return s.FieldId
}

func (s *GetCustomFieldRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCustomFieldRequest) SetFieldId(v string) *GetCustomFieldRequest {
	s.FieldId = &v
	return s
}

func (s *GetCustomFieldRequest) SetInstanceId(v string) *GetCustomFieldRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCustomFieldRequest) Validate() error {
	return dara.Validate(s)
}
