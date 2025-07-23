// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteOperationASyncRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *ExecuteOperationASyncRequest
  GetApplicationId() *string 
  SetAttributes(v map[string]interface{}) *ExecuteOperationASyncRequest
  GetAttributes() map[string]interface{} 
  SetClientToken(v string) *ExecuteOperationASyncRequest
  GetClientToken() *string 
  SetOperation(v string) *ExecuteOperationASyncRequest
  GetOperation() *string 
  SetResourceGroupId(v string) *ExecuteOperationASyncRequest
  GetResourceGroupId() *string 
  SetServiceType(v string) *ExecuteOperationASyncRequest
  GetServiceType() *string 
}

type ExecuteOperationASyncRequest struct {
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
  Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
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

func (s ExecuteOperationASyncRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteOperationASyncRequest) GoString() string {
  return s.String()
}

func (s *ExecuteOperationASyncRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *ExecuteOperationASyncRequest) GetAttributes() map[string]interface{}  {
  return s.Attributes
}

func (s *ExecuteOperationASyncRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *ExecuteOperationASyncRequest) GetOperation() *string  {
  return s.Operation
}

func (s *ExecuteOperationASyncRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *ExecuteOperationASyncRequest) GetServiceType() *string  {
  return s.ServiceType
}

func (s *ExecuteOperationASyncRequest) SetApplicationId(v string) *ExecuteOperationASyncRequest {
  s.ApplicationId = &v
  return s
}

func (s *ExecuteOperationASyncRequest) SetAttributes(v map[string]interface{}) *ExecuteOperationASyncRequest {
  s.Attributes = v
  return s
}

func (s *ExecuteOperationASyncRequest) SetClientToken(v string) *ExecuteOperationASyncRequest {
  s.ClientToken = &v
  return s
}

func (s *ExecuteOperationASyncRequest) SetOperation(v string) *ExecuteOperationASyncRequest {
  s.Operation = &v
  return s
}

func (s *ExecuteOperationASyncRequest) SetResourceGroupId(v string) *ExecuteOperationASyncRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *ExecuteOperationASyncRequest) SetServiceType(v string) *ExecuteOperationASyncRequest {
  s.ServiceType = &v
  return s
}

func (s *ExecuteOperationASyncRequest) Validate() error {
  return dara.Validate(s)
}

