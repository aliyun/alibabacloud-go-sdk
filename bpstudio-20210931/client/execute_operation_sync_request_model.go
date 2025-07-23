// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteOperationSyncRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *ExecuteOperationSyncRequest
  GetApplicationId() *string 
  SetAttributes(v map[string]interface{}) *ExecuteOperationSyncRequest
  GetAttributes() map[string]interface{} 
  SetClientToken(v string) *ExecuteOperationSyncRequest
  GetClientToken() *string 
  SetOperation(v string) *ExecuteOperationSyncRequest
  GetOperation() *string 
  SetResourceGroupId(v string) *ExecuteOperationSyncRequest
  GetResourceGroupId() *string 
  SetServiceType(v string) *ExecuteOperationSyncRequest
  GetServiceType() *string 
}

type ExecuteOperationSyncRequest struct {
  // example:
  // 
  // BE68D71ZY5YYIU9R
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
  // example:
  // 
  // 1600765710019
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // rename
  Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
  // example:
  // 
  // rg-acfmyjt3c5om3hi
  ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // dds
  ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ExecuteOperationSyncRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteOperationSyncRequest) GoString() string {
  return s.String()
}

func (s *ExecuteOperationSyncRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *ExecuteOperationSyncRequest) GetAttributes() map[string]interface{}  {
  return s.Attributes
}

func (s *ExecuteOperationSyncRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *ExecuteOperationSyncRequest) GetOperation() *string  {
  return s.Operation
}

func (s *ExecuteOperationSyncRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *ExecuteOperationSyncRequest) GetServiceType() *string  {
  return s.ServiceType
}

func (s *ExecuteOperationSyncRequest) SetApplicationId(v string) *ExecuteOperationSyncRequest {
  s.ApplicationId = &v
  return s
}

func (s *ExecuteOperationSyncRequest) SetAttributes(v map[string]interface{}) *ExecuteOperationSyncRequest {
  s.Attributes = v
  return s
}

func (s *ExecuteOperationSyncRequest) SetClientToken(v string) *ExecuteOperationSyncRequest {
  s.ClientToken = &v
  return s
}

func (s *ExecuteOperationSyncRequest) SetOperation(v string) *ExecuteOperationSyncRequest {
  s.Operation = &v
  return s
}

func (s *ExecuteOperationSyncRequest) SetResourceGroupId(v string) *ExecuteOperationSyncRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *ExecuteOperationSyncRequest) SetServiceType(v string) *ExecuteOperationSyncRequest {
  s.ServiceType = &v
  return s
}

func (s *ExecuteOperationSyncRequest) Validate() error {
  return dara.Validate(s)
}

