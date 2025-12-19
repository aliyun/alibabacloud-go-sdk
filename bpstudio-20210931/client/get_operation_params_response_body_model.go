// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOperationParamsResponseBody
	GetCode() *string
	SetData(v *GetOperationParamsResponseBodyData) *GetOperationParamsResponseBody
	GetData() *GetOperationParamsResponseBodyData
	SetMessage(v string) *GetOperationParamsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOperationParamsResponseBody
	GetRequestId() *string
}

type GetOperationParamsResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetOperationParamsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// getBpmOperationParams errors
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F439A659-3B3D-50FB-A4BC-69FBF16413C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOperationParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOperationParamsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationParamsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOperationParamsResponseBody) GetData() *GetOperationParamsResponseBodyData {
	return s.Data
}

func (s *GetOperationParamsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOperationParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOperationParamsResponseBody) SetCode(v string) *GetOperationParamsResponseBody {
	s.Code = &v
	return s
}

func (s *GetOperationParamsResponseBody) SetData(v *GetOperationParamsResponseBodyData) *GetOperationParamsResponseBody {
	s.Data = v
	return s
}

func (s *GetOperationParamsResponseBody) SetMessage(v string) *GetOperationParamsResponseBody {
	s.Message = &v
	return s
}

func (s *GetOperationParamsResponseBody) SetRequestId(v string) *GetOperationParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOperationParamsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOperationParamsResponseBodyData struct {
	// example:
	//
	// {
	//
	//       "tags": {
	//
	//         "controlType": "KeyValue",
	//
	//         "display": "Label",
	//
	//         "orderBy": 0,
	//
	//         "associatedLabel": true,
	//
	//         "required": true
	//
	//       }
	//
	//     }
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// rename
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// {
	//
	//       "instance_id": {
	//
	//         "display": "ECS instance ID",
	//
	//         "dataType": "String",
	//
	//         "orderBy": 0,
	//
	//         "attributeName": "instance_id",
	//
	//         "enableVariable": "String",
	//
	//         "required": true
	//
	//       },
	//
	//       "instance_name": {
	//
	//         "display": "ECS name",
	//
	//         "dataType": "String",
	//
	//         "orderBy": 0,
	//
	//         "attributeName": "instance_name",
	//
	//         "required": true
	//
	//       },
	//
	//       "private_ip": {
	//
	//         "display": "Intranet IP",
	//
	//         "dataType": "String",
	//
	//         "orderBy": 0,
	//
	//         "value": "private_ip",
	//
	//         "required": true
	//
	//       }
	//
	//     }
	Outputs map[string]interface{} `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// example:
	//
	// {
	//
	//       "instanceId": "ResourceId",
	//
	//       "regionId": "region",
	//
	//       "appId": "appId"
	//
	//     }
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s GetOperationParamsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOperationParamsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOperationParamsResponseBodyData) GetAttributes() map[string]interface{} {
	return s.Attributes
}

func (s *GetOperationParamsResponseBodyData) GetOperation() *string {
	return s.Operation
}

func (s *GetOperationParamsResponseBodyData) GetOutputs() map[string]interface{} {
	return s.Outputs
}

func (s *GetOperationParamsResponseBodyData) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *GetOperationParamsResponseBodyData) GetService() *string {
	return s.Service
}

func (s *GetOperationParamsResponseBodyData) SetAttributes(v map[string]interface{}) *GetOperationParamsResponseBodyData {
	s.Attributes = v
	return s
}

func (s *GetOperationParamsResponseBodyData) SetOperation(v string) *GetOperationParamsResponseBodyData {
	s.Operation = &v
	return s
}

func (s *GetOperationParamsResponseBodyData) SetOutputs(v map[string]interface{}) *GetOperationParamsResponseBodyData {
	s.Outputs = v
	return s
}

func (s *GetOperationParamsResponseBodyData) SetProperties(v map[string]interface{}) *GetOperationParamsResponseBodyData {
	s.Properties = v
	return s
}

func (s *GetOperationParamsResponseBodyData) SetService(v string) *GetOperationParamsResponseBodyData {
	s.Service = &v
	return s
}

func (s *GetOperationParamsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
