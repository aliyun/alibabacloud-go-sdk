// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeOutputsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOutputs(v []*ListNodeOutputsResponseBodyOutputs) *ListNodeOutputsResponseBody
	GetOutputs() []*ListNodeOutputsResponseBodyOutputs
	SetRequestId(v string) *ListNodeOutputsResponseBody
	GetRequestId() *string
}

type ListNodeOutputsResponseBody struct {
	Outputs []*ListNodeOutputsResponseBodyOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodeOutputsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodeOutputsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeOutputsResponseBody) GetOutputs() []*ListNodeOutputsResponseBodyOutputs {
	return s.Outputs
}

func (s *ListNodeOutputsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodeOutputsResponseBody) SetOutputs(v []*ListNodeOutputsResponseBodyOutputs) *ListNodeOutputsResponseBody {
	s.Outputs = v
	return s
}

func (s *ListNodeOutputsResponseBody) SetRequestId(v string) *ListNodeOutputsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeOutputsResponseBody) Validate() error {
	if s.Outputs != nil {
		for _, item := range s.Outputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodeOutputsResponseBodyOutputs struct {
	AlgoName    *string `json:"AlgoName,omitempty" xml:"AlgoName,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// MaxComputeTable
	LocationType *string `json:"LocationType,omitempty" xml:"LocationType,omitempty"`
	// example:
	//
	// node1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// outputModel
	OutputId *string `json:"OutputId,omitempty" xml:"OutputId,omitempty"`
	// example:
	//
	// 0
	OutputIndex *string `json:"OutputIndex,omitempty" xml:"OutputIndex,omitempty"`
	// example:
	//
	// Model
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// {
	//
	//    "table": "table_name",
	//
	//    "locationType": "MaxComputeTable"
	//
	// }
	Value map[string]interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeOutputsResponseBodyOutputs) String() string {
	return dara.Prettify(s)
}

func (s ListNodeOutputsResponseBodyOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeOutputsResponseBodyOutputs) GetAlgoName() *string {
	return s.AlgoName
}

func (s *ListNodeOutputsResponseBodyOutputs) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListNodeOutputsResponseBodyOutputs) GetLocationType() *string {
	return s.LocationType
}

func (s *ListNodeOutputsResponseBodyOutputs) GetNodeName() *string {
	return s.NodeName
}

func (s *ListNodeOutputsResponseBodyOutputs) GetOutputId() *string {
	return s.OutputId
}

func (s *ListNodeOutputsResponseBodyOutputs) GetOutputIndex() *string {
	return s.OutputIndex
}

func (s *ListNodeOutputsResponseBodyOutputs) GetType() *string {
	return s.Type
}

func (s *ListNodeOutputsResponseBodyOutputs) GetValue() map[string]interface{} {
	return s.Value
}

func (s *ListNodeOutputsResponseBodyOutputs) SetAlgoName(v string) *ListNodeOutputsResponseBodyOutputs {
	s.AlgoName = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetDisplayName(v string) *ListNodeOutputsResponseBodyOutputs {
	s.DisplayName = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetLocationType(v string) *ListNodeOutputsResponseBodyOutputs {
	s.LocationType = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetNodeName(v string) *ListNodeOutputsResponseBodyOutputs {
	s.NodeName = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetOutputId(v string) *ListNodeOutputsResponseBodyOutputs {
	s.OutputId = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetOutputIndex(v string) *ListNodeOutputsResponseBodyOutputs {
	s.OutputIndex = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetType(v string) *ListNodeOutputsResponseBodyOutputs {
	s.Type = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetValue(v map[string]interface{}) *ListNodeOutputsResponseBodyOutputs {
	s.Value = v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) Validate() error {
	return dara.Validate(s)
}
