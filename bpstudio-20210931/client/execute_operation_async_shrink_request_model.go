// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteOperationASyncShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *ExecuteOperationASyncShrinkRequest
  GetApplicationId() *string 
  SetAttributesShrink(v string) *ExecuteOperationASyncShrinkRequest
  GetAttributesShrink() *string 
  SetClientToken(v string) *ExecuteOperationASyncShrinkRequest
  GetClientToken() *string 
  SetOperation(v string) *ExecuteOperationASyncShrinkRequest
  GetOperation() *string 
  SetResourceGroupId(v string) *ExecuteOperationASyncShrinkRequest
  GetResourceGroupId() *string 
  SetServiceType(v string) *ExecuteOperationASyncShrinkRequest
  GetServiceType() *string 
}

type ExecuteOperationASyncShrinkRequest struct {
  // The ID of the Cloud Architect Design Tools (CADT) application.
  // 
  // example:
  // 
  // BG**********UQ
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // The parameters related to the action. Specify the parameters based on the value of Operation. The parameters are passed in the map format. The following examples show how to specify the parameters if you want to change the specifications of an Elastic Compute Service (ECS) instance:
  // 
  // 	- The following common parameters are required: change_type, regionId, instanceId, appId
  // 
  // 	- Example values for changing the instance type of the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_instance_type","instance_type":"ecs.hfr7.2xlarge","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
  // 
  // 	- Example values for stopping the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Stopped","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
  // 
  // 	- Example values for starting the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Running","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
  // 
  // 	- Example values for restarting the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Restart","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
  // 
  // Example of enumerating more than one set of parameters:
  // 
  // 	- { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{\\\\"change_type\\\\":\\\\"modify_instance_type\\\\",\\\\"instance_type\\\\":\\\\"ecs.hfr7.2xlarge\\\\",\\\\"instanceId\\\\":\\\\"i-xxxxxxxxx\\\\",\\\\"regionId\\\\":\\\\"cn-beijing\\\\",\\\\"appId\\\\":\\\\"xxxxxxxxxxxxx\\\\"}" }
  // 
  //     <!-- -->
  // 
  //     :
  // 
  //     <!-- -->
  // 
  //     { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{\\\\"change_type\\\\":\\\\"modify_instance_type\\\\",\\\\"instance_type\\\\":\\\\"ecs.hfr7.2xlarge\\\\",\\\\"instanceId\\\\":\\\\"i-xxxxxxxxx\\\\",\\\\"regionId\\\\":\\\\"cn-beijing\\\\",\\\\"appId\\\\":\\\\"xxxxxxxxxxxxx\\\\"}" }
  // 
  //     <!-- -->
  AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
  // The client token that is used to ensure the idempotence of the request.
  // 
  // example:
  // 
  // 1600765710019
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // This operation type is the operation type of modifying the product, some operation types are generic, and some are used alone. The following is an example of ECS deployment:
  // 
  // - The name of the ECS: rename
  // 
  // - Specification of ecs: modifyInstanceType
  // 
  // - Startup of ecs: modifyInstanceType
  // 
  // - Stop of ecs: modifyInstanceType
  // 
  // - Restart of ecs: modifyInstanceType
  // 
  // - Ecs Tag: addTags
  // 
  // - Deletion of ecs: ecsDelete
  // 
  // - Paid type for ecs: modifyPayType
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // queryTopo
  Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
  // Resource group ID, which is used to verify the permissions of the resource group
  // 
  // example:
  // 
  // ceshi
  ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
  // The type of the service. If you want to perform operations on an Elastic Compute Service (ECS) instance, set ServiceType to ecs.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ACK
  ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ExecuteOperationASyncShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteOperationASyncShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExecuteOperationASyncShrinkRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *ExecuteOperationASyncShrinkRequest) GetAttributesShrink() *string  {
  return s.AttributesShrink
}

func (s *ExecuteOperationASyncShrinkRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *ExecuteOperationASyncShrinkRequest) GetOperation() *string  {
  return s.Operation
}

func (s *ExecuteOperationASyncShrinkRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *ExecuteOperationASyncShrinkRequest) GetServiceType() *string  {
  return s.ServiceType
}

func (s *ExecuteOperationASyncShrinkRequest) SetApplicationId(v string) *ExecuteOperationASyncShrinkRequest {
  s.ApplicationId = &v
  return s
}

func (s *ExecuteOperationASyncShrinkRequest) SetAttributesShrink(v string) *ExecuteOperationASyncShrinkRequest {
  s.AttributesShrink = &v
  return s
}

func (s *ExecuteOperationASyncShrinkRequest) SetClientToken(v string) *ExecuteOperationASyncShrinkRequest {
  s.ClientToken = &v
  return s
}

func (s *ExecuteOperationASyncShrinkRequest) SetOperation(v string) *ExecuteOperationASyncShrinkRequest {
  s.Operation = &v
  return s
}

func (s *ExecuteOperationASyncShrinkRequest) SetResourceGroupId(v string) *ExecuteOperationASyncShrinkRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *ExecuteOperationASyncShrinkRequest) SetServiceType(v string) *ExecuteOperationASyncShrinkRequest {
  s.ServiceType = &v
  return s
}

func (s *ExecuteOperationASyncShrinkRequest) Validate() error {
  return dara.Validate(s)
}

