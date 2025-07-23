// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteOperationSyncShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *ExecuteOperationSyncShrinkRequest
  GetApplicationId() *string 
  SetAttributesShrink(v string) *ExecuteOperationSyncShrinkRequest
  GetAttributesShrink() *string 
  SetClientToken(v string) *ExecuteOperationSyncShrinkRequest
  GetClientToken() *string 
  SetOperation(v string) *ExecuteOperationSyncShrinkRequest
  GetOperation() *string 
  SetResourceGroupId(v string) *ExecuteOperationSyncShrinkRequest
  GetResourceGroupId() *string 
  SetServiceType(v string) *ExecuteOperationSyncShrinkRequest
  GetServiceType() *string 
}

type ExecuteOperationSyncShrinkRequest struct {
  // example:
  // 
  // BE68D71ZY5YYIU9R
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
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

func (s ExecuteOperationSyncShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteOperationSyncShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExecuteOperationSyncShrinkRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *ExecuteOperationSyncShrinkRequest) GetAttributesShrink() *string  {
  return s.AttributesShrink
}

func (s *ExecuteOperationSyncShrinkRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *ExecuteOperationSyncShrinkRequest) GetOperation() *string  {
  return s.Operation
}

func (s *ExecuteOperationSyncShrinkRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *ExecuteOperationSyncShrinkRequest) GetServiceType() *string  {
  return s.ServiceType
}

func (s *ExecuteOperationSyncShrinkRequest) SetApplicationId(v string) *ExecuteOperationSyncShrinkRequest {
  s.ApplicationId = &v
  return s
}

func (s *ExecuteOperationSyncShrinkRequest) SetAttributesShrink(v string) *ExecuteOperationSyncShrinkRequest {
  s.AttributesShrink = &v
  return s
}

func (s *ExecuteOperationSyncShrinkRequest) SetClientToken(v string) *ExecuteOperationSyncShrinkRequest {
  s.ClientToken = &v
  return s
}

func (s *ExecuteOperationSyncShrinkRequest) SetOperation(v string) *ExecuteOperationSyncShrinkRequest {
  s.Operation = &v
  return s
}

func (s *ExecuteOperationSyncShrinkRequest) SetResourceGroupId(v string) *ExecuteOperationSyncShrinkRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *ExecuteOperationSyncShrinkRequest) SetServiceType(v string) *ExecuteOperationSyncShrinkRequest {
  s.ServiceType = &v
  return s
}

func (s *ExecuteOperationSyncShrinkRequest) Validate() error {
  return dara.Validate(s)
}

