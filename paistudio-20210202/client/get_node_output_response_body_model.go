// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgoName(v string) *GetNodeOutputResponseBody
	GetAlgoName() *string
	SetDisplayName(v string) *GetNodeOutputResponseBody
	GetDisplayName() *string
	SetLocationType(v string) *GetNodeOutputResponseBody
	GetLocationType() *string
	SetNodeName(v string) *GetNodeOutputResponseBody
	GetNodeName() *string
	SetRequestId(v string) *GetNodeOutputResponseBody
	GetRequestId() *string
	SetType(v string) *GetNodeOutputResponseBody
	GetType() *string
	SetValue(v map[string]interface{}) *GetNodeOutputResponseBody
	GetValue() map[string]interface{}
}

type GetNodeOutputResponseBody struct {
	AlgoName    *string `json:"AlgoName,omitempty" xml:"AlgoName,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// MaxComputeTable
	LocationType *string `json:"LocationType,omitempty" xml:"LocationType,omitempty"`
	NodeName     *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// 601FD8B1-78EB-5220-844C-92AC2EDAF7E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Model
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// {
	//
	//   "name": "model_flow_wayrh3k605s7i51wey_node_7n3tstbuhr36t0ukiz_model",
	//
	//   "modelType": "OfflineModel",
	//
	//   "labelCol": "_c2",
	//
	//   "features": "pm10,so2,co,no2",
	//
	//   "gmtCreateTime": "2021-01-21T17:12:35.232Z",
	//
	//   "gmtModifiedTime": "2021-01-21T17:12:35.232Z",
	//
	//   "parameters": {
	//
	//     "epsilon": "0.000001",
	//
	//     "enableSparse": "false",
	//
	//     "regularizedLevel": "1",
	//
	//     "roleArn": "true",
	//
	//     "maxIter": "100",
	//
	//     "regularizedType": "None",
	//
	//     "generatePmml": "true"
	//
	//   }
	//
	// }
	Value map[string]interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetNodeOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeOutputResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeOutputResponseBody) GetAlgoName() *string {
	return s.AlgoName
}

func (s *GetNodeOutputResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetNodeOutputResponseBody) GetLocationType() *string {
	return s.LocationType
}

func (s *GetNodeOutputResponseBody) GetNodeName() *string {
	return s.NodeName
}

func (s *GetNodeOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeOutputResponseBody) GetType() *string {
	return s.Type
}

func (s *GetNodeOutputResponseBody) GetValue() map[string]interface{} {
	return s.Value
}

func (s *GetNodeOutputResponseBody) SetAlgoName(v string) *GetNodeOutputResponseBody {
	s.AlgoName = &v
	return s
}

func (s *GetNodeOutputResponseBody) SetDisplayName(v string) *GetNodeOutputResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetNodeOutputResponseBody) SetLocationType(v string) *GetNodeOutputResponseBody {
	s.LocationType = &v
	return s
}

func (s *GetNodeOutputResponseBody) SetNodeName(v string) *GetNodeOutputResponseBody {
	s.NodeName = &v
	return s
}

func (s *GetNodeOutputResponseBody) SetRequestId(v string) *GetNodeOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeOutputResponseBody) SetType(v string) *GetNodeOutputResponseBody {
	s.Type = &v
	return s
}

func (s *GetNodeOutputResponseBody) SetValue(v map[string]interface{}) *GetNodeOutputResponseBody {
	s.Value = v
	return s
}

func (s *GetNodeOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
