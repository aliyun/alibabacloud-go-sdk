// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomFieldRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFieldId(v string) *EnableCustomFieldRequest
  GetFieldId() *string 
  SetInstanceId(v string) *EnableCustomFieldRequest
  GetInstanceId() *string 
}

type EnableCustomFieldRequest struct {
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

func (s EnableCustomFieldRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomFieldRequest) GoString() string {
  return s.String()
}

func (s *EnableCustomFieldRequest) GetFieldId() *string  {
  return s.FieldId
}

func (s *EnableCustomFieldRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableCustomFieldRequest) SetFieldId(v string) *EnableCustomFieldRequest {
  s.FieldId = &v
  return s
}

func (s *EnableCustomFieldRequest) SetInstanceId(v string) *EnableCustomFieldRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableCustomFieldRequest) Validate() error {
  return dara.Validate(s)
}

